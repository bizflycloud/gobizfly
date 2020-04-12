// Copyright 2020 The BizFly Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gobizfly

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	volumeBasePath = "/iaas-cloud/api/volumes"
)

var _ VolumeService = (*volume)(nil)

// VolumeService is an interface to interact with BizFly API Volume endpoint.
type VolumeService interface {
	List(ctx context.Context, opts *ListOptions) ([]*Volume, error)
	Create(ctx context.Context, req *VolumeCreateRequest) (*Volume, error)
	Get(ctx context.Context, id string) (*Volume, error)
	Delete(ctx context.Context, id string) error
	ExtendVolume(ctx context.Context, id string, newsize int) (*Task, error)
	Attach(ctx context.Context, id string, serverID string) (*VolumeAttachDetachResponse, error)
	Detach(ctx context.Context, id string, serverID string) (*VolumeAttachDetachResponse, error)
	Restore(ctx context.Context, id string, snapshotID string) (*Task, error)
}

// VolumeCreateRequest represents create new volume request payload.
type VolumeCreateRequest struct {
	Name             string `json:"name"`
	Size             int    `json:"size"`
	VolumeType       string `json:"volume_type"`
	AvailabilityZone string `json:"availability_zone"`
	SnapshotID       string `json:"snapshot_id,omitempty"`
}

// VolumeAttachment contains volume attachment information.
type VolumeAttachment struct {
	ServerID     string `json:"server_id"`
	AttachmentID string `json:"attachment_id"`
	VolumeID     string `json:"volume_id"`
	Device       string `json:"device"`
	ID           string `json:"id"`
}

// Volume contains volume information.
type Volume struct {
	ID               string            `json:"id"`
	Size             int               `json:"size"`
	AttachedType     string            `json:"attached_type"`
	Name             string            `json:"name"`
	VolumeType       string            `json:"volume_type"`
	Description      string            `json:"description"`
	SnapshotID       string            `json:"snapshot_id"`
	Bootable         string            `json:"bootable"`
	AvailabilityZone string            `json:"availability_zone"`
	Status           string            `json:"status"`
	UserID           string            `json:"user_id"`
	ProjectID        string            `json:"os-vol-tenant-attr:tenant_id"`
	CreatedAt        string            `json:"created_at"`
	UpdatedAt        string            `json:"updated_at"`
	Metadata         map[string]string `json:"metadata"`
	Attachments      []Server          `json:"attachments"`
}

type volume struct {
	client *Client
}

// List lists all volumes of users.
func (v *volume) List(ctx context.Context, opts *ListOptions) ([]*Volume, error) {
	req, err := v.client.NewRequest(ctx, http.MethodGet, volumeBasePath, nil)
	if err != nil {
		return nil, err
	}
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
func (v *volume) Create(ctx context.Context, vcr *VolumeCreateRequest) (*Volume, error) {
	req, err := v.client.NewRequest(ctx, http.MethodPost, volumeBasePath, &vcr)
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
func (v *volume) Get(ctx context.Context, id string) (*Volume, error) {
	req, err := v.client.NewRequest(ctx, http.MethodGet, volumeBasePath+"/"+id, nil)
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
func (v *volume) Delete(ctx context.Context, id string) error {
	req, err := v.client.NewRequest(ctx, http.MethodDelete, volumeBasePath+"/"+id, nil)

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

func (v *volume) itemActionPath(id string) string {
	return strings.Join([]string{volumeBasePath, id, "action"}, "/")
}

// ExtendVolume extends capacity of a volume.
func (v *volume) ExtendVolume(ctx context.Context, id string, newsize int) (*Task, error) {
	var payload = &VolumeAction{
		Type:    "extend",
		NewSize: newsize}

	req, err := v.client.NewRequest(ctx, http.MethodPost, v.itemActionPath(id), payload)
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
func (v *volume) Attach(ctx context.Context, id string, serverID string) (*VolumeAttachDetachResponse, error) {
	var payload = &VolumeAction{
		Type:     "attach",
		ServerID: serverID}

	req, err := v.client.NewRequest(ctx, http.MethodPost, v.itemActionPath(id), payload)
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
func (v *volume) Detach(ctx context.Context, id string, serverID string) (*VolumeAttachDetachResponse, error) {
	var payload = &VolumeAction{
		Type:     "detach",
		ServerID: serverID}

	req, err := v.client.NewRequest(ctx, http.MethodPost, v.itemActionPath(id), payload)
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
func (v *volume) Restore(ctx context.Context, id string, snapshotID string) (*Task, error) {
	var payload = &VolumeAction{
		Type:       "restore",
		SnapshotID: snapshotID}

	req, err := v.client.NewRequest(ctx, http.MethodPost, v.itemActionPath(id), payload)
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
