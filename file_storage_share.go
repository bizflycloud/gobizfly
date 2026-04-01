// This file is part of gobizfly

package gobizfly

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Share represents a file storage share.
type Share struct {
	ID            string            `json:"id"`
	Name          string            `json:"name"`
	Size          int               `json:"size"`
	ShareProtocol string            `json:"share_protocol"`
	Description   string            `json:"description"`
	NetworkID     string            `json:"network_id"`
	SubnetID      string            `json:"subnet_id"`
	ShareType     string            `json:"share_type"`
	Zone          string            `json:"zone"`
	Status        string            `json:"status"`
	ExportLocation string           `json:"export_location"`
	CreatedAt     string            `json:"created_at"`
	UpdatedAt     string            `json:"updated_at"`
	Metadata      map[string]string `json:"metadata,omitempty"`
}

// CreateShareRequest represents the payload for creating a new share.
type CreateShareRequest struct {
	Name          string             `json:"name"`
	Size          int                `json:"size"`
	ShareProtocol string            `json:"share_protocol,omitempty"`
	Description   string            `json:"description,omitempty"`
	NetworkID     string            `json:"network_id,omitempty"`
	SubnetID      string            `json:"subnet_id,omitempty"`
	ShareType     string            `json:"share_type,omitempty"`
	Access        []ShareAccessRule  `json:"access,omitempty"`
	Zone          string            `json:"zone,omitempty"`
}

// ShareAccessRule represents an access rule in create/manage requests.
type ShareAccessRule struct {
	AccessTo    string `json:"access_to"`
	AccessLevel string `json:"access_level,omitempty"`
	AccessType  string `json:"access_type,omitempty"`
}

// ResizeShareRequest represents the payload for resizing a share.
type ResizeShareRequest struct {
	NewSize int `json:"new_size"`
}

// List returns a list of all shares for the project.
func (fs *fileStorageService) List(ctx context.Context) ([]*Share, error) {
	req, err := fs.client.NewRequest(ctx, http.MethodGet, fileStorageServiceName, fs.resourcePath(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := fs.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	var wrapper struct {
		FileStorages []*Share `json:"filestorages"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&wrapper); err != nil {
		return nil, err
	}
	return wrapper.FileStorages, nil
}

// Create creates a new share.
func (fs *fileStorageService) Create(ctx context.Context, createReq *CreateShareRequest) (*Share, error) {
	req, err := fs.client.NewRequest(ctx, http.MethodPost, fileStorageServiceName, fs.resourcePath(), createReq)
	if err != nil {
		return nil, err
	}
	resp, err := fs.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	var wrapper struct {
		FileStorage *Share `json:"filestorage"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&wrapper); err != nil {
		return nil, err
	}
	return wrapper.FileStorage, nil
}

// Get returns details of a specific share.
func (fs *fileStorageService) Get(ctx context.Context, shareID string) (*Share, error) {
	req, err := fs.client.NewRequest(ctx, http.MethodGet, fileStorageServiceName, fs.shareItemPath(shareID), nil)
	if err != nil {
		return nil, err
	}
	resp, err := fs.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	var wrapper struct {
		FileStorage *Share `json:"filestorage"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&wrapper); err != nil {
		return nil, err
	}
	return wrapper.FileStorage, nil
}

// Delete deletes a share. If force is true, the share is forcefully deleted.
func (fs *fileStorageService) Delete(ctx context.Context, shareID string, force bool) error {
	req, err := fs.client.NewRequest(ctx, http.MethodDelete, fileStorageServiceName, fs.shareItemPath(shareID), nil)
	if err != nil {
		return err
	}
	if force {
		q := req.URL.Query()
		q.Set("force", fmt.Sprintf("%t", force))
		req.URL.RawQuery = q.Encode()
	}
	resp, err := fs.client.Do(ctx, req)
	if err != nil {
		return err
	}
	return resp.Body.Close()
}

// Resize resizes a share (extend or shrink).
func (fs *fileStorageService) Resize(ctx context.Context, shareID string, resizeReq *ResizeShareRequest) (*Share, error) {
	req, err := fs.client.NewRequest(ctx, http.MethodPost, fileStorageServiceName, fs.shareResizePath(shareID), resizeReq)
	if err != nil {
		return nil, err
	}
	resp, err := fs.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	var wrapper struct {
		FileStorage *Share `json:"filestorage"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&wrapper); err != nil {
		return nil, err
	}
	return wrapper.FileStorage, nil
}
