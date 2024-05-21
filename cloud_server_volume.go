// This file is part of gobizfly

package gobizfly

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const (
	volumeBasePath      = "/volumes"
	volumeTypesBasePath = "/volume-types"
)

var _ VolumeService = (*cloudServerVolumeResource)(nil)

// VolumeService is an interface to interact with Bizfly API Volume endpoint.
type VolumeService interface {
	List(ctx context.Context, opts *VolumeListOptions) ([]*Volume, error)
	Create(ctx context.Context, req *VolumeCreateRequest) (*Volume, error)
	Get(ctx context.Context, id string) (*Volume, error)
	Delete(ctx context.Context, id string) error
	ExtendVolume(ctx context.Context, id string, newsize int) (*Task, error)
	Attach(ctx context.Context, id string, serverID string) (*VolumeAttachDetachResponse, error)
	Detach(ctx context.Context, id string, serverID string) (*VolumeAttachDetachResponse, error)
	Restore(ctx context.Context, id string, snapshotID string) (*Task, error)
	Patch(ctx context.Context, id string, req *VolumePatchRequest) (*Volume, error)
	ListVolumeTypes(ctx context.Context, opts *ListVolumeTypesOptions) ([]*VolumeType, error)
}

// VolumeListOptions represents options to list volumes.
// Name is the filter to list volumes by name.
// Size is the filter to list volumes by size.
// Status is the filter to list volumes by status.
// AvailabilityZone is the filter to list volumes by availability zone.
// Category is the filter to list volumes by category.
// BillingPlan is the filter to list volumes by billing plan.
// Bootable is the filter to list volumes by bootable.
type VolumeListOptions struct {
	Name             string `json:"name"`
	Size             int    `json:"size"`
	Status           string `json:"status"`
	AvailabilityZone string `json:"availability_zone"`
	Category         string `json:"category"`
	BillingPlan      string `json:"billing_plan"`
	Bootable         *bool  `json:"bootable"`
}

// VolumeCreateRequest represents create new volume request payload.
type VolumeCreateRequest struct {
	Name             string `json:"name"`
	Description      string `json:"description"`
	Size             int    `json:"size"`
	VolumeType       string `json:"volume_type"`
	VolumeCategory   string `json:"category"`
	AvailabilityZone string `json:"availability_zone"`
	SnapshotID       string `json:"snapshot_id,omitempty"`
	ServerID         string `json:"instance_uuid,omitempty"`
	BillingPlan      string `json:"billing_plan,omitempty"`
}

type VolumePatchRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// VolumeAttachment contains volume attachment information.
type VolumeAttachment struct {
	Server       Server `json:"server"`
	ServerID     string `json:"server_id"`
	AttachmentID string `json:"attachment_id"`
	VolumeID     string `json:"volume_id"`
	Device       string `json:"device"`
	ID           string `json:"id"`
}

type VolumeImageMetadata struct {
	ContainerFormat string `json:"container_format"`
	DiskFormat      string `json:"disk_format"`
	ImageID         string `json:"image_id"`
	ImageName       string `json:"image_name"`
	ImageType       string `json:"image_type"`
	MinDisk         string `json:"min_disk"`
	MinRam          string `json:"min_ram"`
	Size            string `json:"size"`
}

// Volume contains volume information.
type Volume struct {
	ID               string              `json:"id"`
	Size             int                 `json:"size"`
	AttachedType     string              `json:"attached_type"`
	Name             string              `json:"name"`
	Type             string              `json:"type"`
	VolumeType       string              `json:"volume_type"`
	Description      string              `json:"description"`
	SnapshotID       string              `json:"snapshot_id"`
	Bootable         bool                `json:"bootable"`
	AvailabilityZone string              `json:"availability_zone"`
	Status           string              `json:"status"`
	UserID           string              `json:"user_id"`
	ProjectID        string              `json:"os-vol-tenant-attr:tenant_id"`
	CreatedAt        string              `json:"created_at"`
	UpdatedAt        string              `json:"updated_at"`
	Metadata         map[string]string   `json:"metadata"`
	Attachments      []VolumeAttachment  `json:"attachments"`
	Category         string              `json:"category"`
	BillingPlan      string              `json:"billing_plan"`
	Encrypted        bool                `json:"encrypted"`
	ImageMetadata    VolumeImageMetadata `json:"volume_image_metadata"`
}

// VolumeType contains volume type information.
type VolumeType struct {
	Name              string   `json:"name"`
	Category          string   `json:"category"`
	Type              string   `json:"type"`
	AvailabilityZones []string `json:"availability_zones"`
}

// ListVolumeTypesOptions contains options for listing volume types.
type ListVolumeTypesOptions struct {
	Category         string `json:"category,omitempty"`
	AvailabilityZone string `json:"availability_zone,omitempty"`
}

type cloudServerVolumeResource struct {
	client *Client
}

func (cs *cloudServerService) Volumes() *cloudServerVolumeResource {
	return &cloudServerVolumeResource{client: cs.client}
}

// List lists all volumes of users.
func (v *cloudServerVolumeResource) List(ctx context.Context, opts *VolumeListOptions) ([]*Volume, error) {
	req, err := v.client.NewRequest(ctx, http.MethodGet, serverServiceName, volumeBasePath, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	if opts != nil {
		if opts.Name != "" {
			q.Add("name", opts.Name)
		}
		if opts.Size != 0 {
			q.Add("size", strconv.Itoa(opts.Size))
		}
		if opts.Status != "" {
			q.Add("status", opts.Status)
		}
		if opts.AvailabilityZone != "" {
			q.Add("availability_zone", opts.AvailabilityZone)
		}
		if opts.Category != "" {
			q.Add("category", opts.Category)
		}
		if opts.BillingPlan != "" {
			q.Add("billing_plan", opts.BillingPlan)
		}
		if opts.Bootable != nil {
			q.Add("bootable", strconv.FormatBool(*opts.Bootable))
		}
	}
	req.URL.RawQuery = q.Encode()
	resp, err := v.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var volumes []*Volume

	if err := json.NewDecoder(resp.Body).Decode(&volumes); err != nil {
		return nil, err
	}
	return volumes, nil
}

// Create creates a new volume.
func (v *cloudServerVolumeResource) Create(ctx context.Context, vcr *VolumeCreateRequest) (*Volume, error) {
	req, err := v.client.NewRequest(ctx, http.MethodPost, serverServiceName, volumeBasePath, &vcr)
	if err != nil {
		return nil, err
	}
	resp, err := v.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var volumeRespData *Volume

	if err := json.NewDecoder(resp.Body).Decode(&volumeRespData); err != nil {
		return nil, err
	}
	return volumeRespData, nil
}

// Get gets information of a volume.
func (v *cloudServerVolumeResource) Get(ctx context.Context, id string) (*Volume, error) {
	req, err := v.client.NewRequest(ctx, http.MethodGet, serverServiceName, volumeBasePath+"/"+id, nil)
	if err != nil {
		return nil, err
	}

	resp, err := v.client.Do(ctx, req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var volume *Volume

	if err := json.NewDecoder(resp.Body).Decode(&volume); err != nil {
		return nil, err
	}
	return volume, nil
}

// Delete deletes a volume.
func (v *cloudServerVolumeResource) Delete(ctx context.Context, id string) error {
	req, err := v.client.NewRequest(ctx, http.MethodDelete, serverServiceName, volumeBasePath+"/"+id, nil)

	if err != nil {
		return err
	}

	resp, err := v.client.Do(ctx, req)
	if err != nil {
		return err
	}
	_, _ = io.Copy(ioutil.Discard, resp.Body)

	return resp.Body.Close()
}

// VolumeAction represents volume action request payload.
type VolumeAction struct {
	Type       string `json:"type"`
	NewSize    int    `json:"new_size,omitempty"`
	ServerID   string `json:"instance_uuid,omitempty"`
	SnapshotID string `json:"snapshot_id,omitempty"`
}

// Task represents task response when perform an action.
type Task struct {
	TaskID string `json:"task_id"`
}

// VolumeAttachDetachResponse contains information when detach or attach a volume from/to a server.
type VolumeAttachDetachResponse struct {
	Message      string `json:"message"`
	VolumeDetail Volume `json:"volume_detail"`
}

func (v *cloudServerVolumeResource) itemActionPath(id string) string {
	return strings.Join([]string{volumeBasePath, id, "action"}, "/")
}

func (v *cloudServerVolumeResource) itemPath(id string) string {
	return strings.Join([]string{volumeBasePath, id}, "/")
}

// ExtendVolume extends capacity of a volume.
func (v *cloudServerVolumeResource) ExtendVolume(ctx context.Context, id string, newsize int) (*Task, error) {
	var payload = &VolumeAction{
		Type:    "extend",
		NewSize: newsize}

	req, err := v.client.NewRequest(ctx, http.MethodPost, serverServiceName, v.itemActionPath(id), payload)
	if err != nil {
		return nil, err
	}

	resp, err := v.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	t := &Task{}
	if err := json.NewDecoder(resp.Body).Decode(t); err != nil {
		return nil, err
	}
	return t, nil
}

// Attach attaches a volume to a server.
func (v *cloudServerVolumeResource) Attach(ctx context.Context, id string, serverID string) (*VolumeAttachDetachResponse, error) {
	var payload = &VolumeAction{
		Type:     "attach",
		ServerID: serverID}

	req, err := v.client.NewRequest(ctx, http.MethodPost, serverServiceName, v.itemActionPath(id), payload)
	if err != nil {
		return nil, err
	}

	resp, err := v.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	volumeAttachResponse := &VolumeAttachDetachResponse{}
	if err := json.NewDecoder(resp.Body).Decode(volumeAttachResponse); err != nil {
		return nil, err
	}
	return volumeAttachResponse, nil
}

// Detach detaches a volume from a server.
func (v *cloudServerVolumeResource) Detach(ctx context.Context, id string, serverID string) (*VolumeAttachDetachResponse, error) {
	var payload = &VolumeAction{
		Type:     "detach",
		ServerID: serverID}

	req, err := v.client.NewRequest(ctx, http.MethodPost, serverServiceName, v.itemActionPath(id), payload)
	if err != nil {
		return nil, err
	}

	resp, err := v.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	r := &VolumeAttachDetachResponse{}
	if err := json.NewDecoder(resp.Body).Decode(r); err != nil {
		return nil, err
	}
	return r, nil
}

// Restore restores a volume from a snapshot.
func (v *cloudServerVolumeResource) Restore(ctx context.Context, id string, snapshotID string) (*Task, error) {
	var payload = &VolumeAction{
		Type:       "restore_volume",
		SnapshotID: snapshotID}

	req, err := v.client.NewRequest(ctx, http.MethodPost, serverServiceName, v.itemActionPath(id), payload)
	if err != nil {
		return nil, err
	}

	resp, err := v.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	t := &Task{}
	if err := json.NewDecoder(resp.Body).Decode(t); err != nil {
		return nil, err
	}
	return t, nil
}

// Patch partially updates a volume.
func (v *cloudServerVolumeResource) Patch(ctx context.Context, id string, bpr *VolumePatchRequest) (*Volume, error) {
	req, err := v.client.NewRequest(ctx, http.MethodPatch, serverServiceName, v.itemPath(id), bpr)
	if err != nil {
		return nil, err
	}
	resp, err := v.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var volume *Volume
	if err := json.NewDecoder(resp.Body).Decode(&volume); err != nil {
		return nil, err
	}
	return volume, nil
}

// ListVolumeTypes lists volume types.
func (v *cloudServerVolumeResource) ListVolumeTypes(ctx context.Context, opts *ListVolumeTypesOptions) ([]*VolumeType, error) {
	req, err := v.client.NewRequest(ctx, http.MethodGet, serverServiceName, volumeTypesBasePath, nil)
	if err != nil {
		return nil, err
	}
	// Set query parameters from options.
	query := req.URL.Query()
	if opts != nil {
		if opts.Category != "" {
			query.Set("category", opts.Category)
		}
		if opts.AvailabilityZone != "" {
			query.Set("availability_zone", opts.AvailabilityZone)
		}
	}
	req.URL.RawQuery = query.Encode()
	resp, err := v.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var volumeTypesResp struct {
		VolumeTypes []*VolumeType `json:"volume_types"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&volumeTypesResp); err != nil {
		return nil, err
	}
	return volumeTypesResp.VolumeTypes, nil
}
