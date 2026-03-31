// This file is part of gobizfly

package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
)

// FileStorageAccessRule represents an access rule for a share.
type FileStorageAccessRule struct {
	ID          string `json:"id"`
	AccessTo    string `json:"access_to"`
	AccessLevel string `json:"access_level"`
	AccessType  string `json:"access_type"`
	State       string `json:"state"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// ManageAccessRequest represents the payload for managing access rules.
type ManageAccessRequest struct {
	Access []ShareAccessRule `json:"access"`
}

// FileStorageAccessStatus represents the access status for a share.
type FileStorageAccessStatus struct {
	Status string `json:"status"`
}

// GetAccessRules returns all access rules for a share.
func (fs *fileStorageService) GetAccessRules(ctx context.Context, shareID string) ([]*FileStorageAccessRule, error) {
	req, err := fs.client.NewRequest(ctx, http.MethodGet, fileStorageServiceName, fs.shareAccessPath(shareID), nil)
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
	var data []*FileStorageAccessRule
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

// ManageAccessRules replaces all existing access rules for a share.
func (fs *fileStorageService) ManageAccessRules(ctx context.Context, shareID string, manageReq *ManageAccessRequest) ([]*FileStorageAccessRule, error) {
	req, err := fs.client.NewRequest(ctx, http.MethodPost, fileStorageServiceName, fs.shareAccessPath(shareID), manageReq)
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
	var data []*FileStorageAccessRule
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

// DeleteAccessRule deletes a specific access rule from a share.
func (fs *fileStorageService) DeleteAccessRule(ctx context.Context, shareID string, ruleID string) error {
	req, err := fs.client.NewRequest(ctx, http.MethodDelete, fileStorageServiceName, fs.shareAccessRulePath(shareID, ruleID), nil)
	if err != nil {
		return err
	}
	resp, err := fs.client.Do(ctx, req)
	if err != nil {
		return err
	}
	return resp.Body.Close()
}

// GetAccessStatus returns the access status for a share.
func (fs *fileStorageService) GetAccessStatus(ctx context.Context, shareID string) (*FileStorageAccessStatus, error) {
	req, err := fs.client.NewRequest(ctx, http.MethodGet, fileStorageServiceName, fs.shareAccessStatusPath(shareID), nil)
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
	var data *FileStorageAccessStatus
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}
