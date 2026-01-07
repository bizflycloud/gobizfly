package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

// CustomImage represents a custom image information
type CustomImage struct {
	Name            string     `json:"name"`
	Description     string     `json:"description"`
	DiskFormat      string     `json:"disk_format"`
	ContainerFormat string     `json:"container_format"`
	Visibility      string     `json:"visibility"`
	Size            int        `json:"size"`
	VirtualSize     int        `json:"virtual_size"`
	Status          string     `json:"status"`
	Checksum        string     `json:"checksum"`
	Protected       bool       `json:"protected"`
	MinRam          int        `json:"min_ram"`
	MinDisk         int        `json:"min_disk"`
	Owner           string     `json:"owner"`
	OSHidden        bool       `json:"os_hidden"`
	OSHashAlgo      string     `json:"os_hash_algo"`
	OSHashValue     string     `json:"os_hash_value"`
	ID              string     `json:"id"`
	CreatedAt       string     `json:"created_at"`
	UpdatedAt       string     `json:"updated_at"`
	Locations       []Location `json:"locations"`
	DirectURL       string     `json:"direct_url"`
	Tags            []string   `json:"tags"`
	File            string     `json:"file"`
	Schema          string     `json:"schema"`
	BillingPlan     string     `json:"billing_plan"`
}

// CreateCustomImageResp represents the request body when creating custom image
type CreateCustomImageResp struct {
	Image     CustomImage `json:"image"`
	Success   bool        `json:"success"`
	Token     string      `json:"token,omitempty"`
	UploadURI string      `json:"upload_uri,omitempty"`
}

// CustomImageGetResp represents the response body when getting custom image
type CustomImageGetResp struct {
	Image CustomImage `json:"image"`
	Token string      `json:"token"`
}

// itemCustomImagePath represents the path to get a custom image
func (s *cloudServerCustomOSImageResource) itemCustomImagePath(id string) string {
	return strings.Join([]string{customImagePath, id}, "/")
}

// osDistributionVersion represents the os distribution version
type osDistributionVersion struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

// osImageResponse represents the response body when getting os image
type osImageResponse struct {
	OSDistribution string                  `json:"os"`
	Version        []osDistributionVersion `json:"versions"`
}

type cloudServerOSImageResource struct {
	client *Client
}

type cloudServerCustomOSImageResource struct {
	client *Client
}

func (cs *cloudServerService) OSImages() *cloudServerOSImageResource {
	return &cloudServerOSImageResource{client: cs.client}
}

func (cs *cloudServerService) CustomImages() *cloudServerCustomOSImageResource {
	return &cloudServerCustomOSImageResource{client: cs.client}
}

// Get list server os images
func (s *cloudServerOSImageResource) List(ctx context.Context) ([]osImageResponse, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, serverServiceName, osImagePath, nil)

	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("os_images", "True")
	req.URL.RawQuery = q.Encode()

	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	var respPayload struct {
		OSImages []osImageResponse `json:"os_images"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respPayload); err != nil {
		return nil, err
	}
	return respPayload.OSImages, nil
}

// List - List custom images
func (s *cloudServerCustomOSImageResource) List(ctx context.Context) ([]*CustomImage, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, serverServiceName, customImagesPath, nil)
	if err != nil {
		return nil, err
	}
	var data struct {
		Images []*CustomImage `json:"images"`
	}
	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data.Images, nil
}

// Create - Create custom image
func (s *cloudServerCustomOSImageResource) Create(ctx context.Context, cipl *CreateCustomImagePayload) (*CreateCustomImageResp, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, serverServiceName, strings.Join([]string{customImagePath, "upload"}, "/"), cipl)
	if err != nil {
		return nil, err
	}
	var data *CreateCustomImageResp
	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

// Delete - Delete a custom image by id
func (s *cloudServerCustomOSImageResource) Delete(ctx context.Context, imageID string) error {
	req, err := s.client.NewRequest(ctx, http.MethodDelete, serverServiceName, s.itemCustomImagePath(imageID), nil)

	if err != nil {
		return err
	}
	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return err
	}
	return resp.Body.Close()
}

// Get - Get a custom image by id
func (s *cloudServerCustomOSImageResource) Get(ctx context.Context, imageID string) (*CustomImageGetResp, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, serverServiceName, strings.Join([]string{customImagePath, imageID}, "/"), nil)
	if err != nil {
		return nil, err
	}
	var data *CustomImageGetResp
	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}
