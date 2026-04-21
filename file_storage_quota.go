// This file is part of gobizfly

package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
)

// QuotaRequest represents the payload for setting or updating a quota.
type QuotaRequest struct {
	ProjectID      string `json:"project_id"`
	MaxShares      int    `json:"max_shares"`
	MaxTotalSizeGB int    `json:"max_total_size_gb"`
}

// QuotaResponse represents the quota information for a project.
type QuotaResponse struct {
	ProjectID      string `json:"project_id"`
	MaxShares      int    `json:"max_shares"`
	MaxTotalSizeGB int    `json:"max_total_size_gb"`
}

// FileStorageRegion represents a region for file storage.
type FileStorageRegion struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

// GetQuota returns the quota for a project.
func (fs *fileStorageService) GetQuota(ctx context.Context, projectID string) (*QuotaResponse, error) {
	req, err := fs.client.NewRequest(ctx, http.MethodGet, fileStorageServiceName, fs.quotaProjectPath(projectID), nil)
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
	var data *QuotaResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

// UpdateQuota sets or updates the quota for a project.
func (fs *fileStorageService) UpdateQuota(ctx context.Context, quotaReq *QuotaRequest) (*QuotaResponse, error) {
	req, err := fs.client.NewRequest(ctx, http.MethodPost, fileStorageServiceName, fs.quotaPath(), quotaReq)
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
	var data *QuotaResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

// ListRegions returns a list of available regions for file storage.
func (fs *fileStorageService) ListRegions(ctx context.Context) ([]*FileStorageRegion, error) {
	req, err := fs.client.NewRequest(ctx, http.MethodGet, fileStorageServiceName, fs.regionsPath(), nil)
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
	var data []*FileStorageRegion
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}
