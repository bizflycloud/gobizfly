package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

// CloudBackupCreatePolicyPayload represents the payload for creating a backup policy
type CloudBackupCreatePolicyPayload struct {
	Name            string `json:"name"`
	StorageType     string `json:"storage_type"`
	SchedulePattern string `json:"schedule_pattern"`
	RetentionDays   int    `json:"retention_days"`
	Description     string `json:"description,omitempty"`
}

// CloudBackupPolicy represents a backup policy
type CloudBackupPolicy struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	SchedulePattern string `json:"schedule_pattern"`
	RetentionDays   int    `json:"retention_days"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	TenantID        string `json:"tenant_id"`
	Description     string `json:"description"`
	LimitDownload   int    `json:"limit_download,omitempty"`
	LimitUpload     int    `json:"limit_upload,omitempty"`
	PolicyType      string `json:"policy_type"`
	Retentions      int    `json:"retentions"`
}

// CloudBackupPatchPolicyPayload represents the payload for patching a backup policy
type CloudBackupPatchPolicyPayload struct {
	Name            string `json:"name,omitempty"`
	SchedulePattern string `json:"schedule_pattern,omitempty"`
	RetentionDays   int    `json:"retention_days,omitempty"`
}

// CloudBackupFullPolicy represents a backup policy
type CloudBackupFullPolicy struct {
	CloudBackupPolicy
	RetentionDays     int      `json:"retention_days"`
	RetentionHours    int      `json:"retention_hours"`
	RetentionWeeks    int      `json:"retention_weeks"`
	RetentionMonths   int      `json:"retention_months"`
	BackupDirectories []string `json:"backup_directories"`
}

// CloudBackupActionPolicyDirectoryPayload represents the payload for creating a backup policy
type CloudBackupActionPolicyDirectoryPayload struct {
	Action      string `json:"action"`
	DirectoryID string `json:"directory_id"`
}

// CreatePolicy creates a backup policy
func (cb *cloudBackupService) CreatePolicy(ctx context.Context, payload *CloudBackupCreatePolicyPayload) (*CloudBackupPolicy, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodPost, cloudBackupServiceName,
		cb.policyPath(), payload)
	if err != nil {
		return nil, err
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	var policy *CloudBackupPolicy
	if err := json.NewDecoder(resp.Body).Decode(&policy); err != nil {
		return nil, err
	}
	return policy, nil
}

// GetBackupDirectoryPolicy gets a backup policy
func (cb *cloudBackupService) GetBackupDirectoryPolicy(ctx context.Context, machineID string, directoryID string) (*CloudBackupPolicy, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodGet, cloudBackupServiceName,
		strings.Join([]string{cb.itemMachinePath(machineID), "directories", directoryID, "policies"}, "/"), nil)
	if err != nil {
		return nil, err
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	var policy *CloudBackupPolicy
	if err := json.NewDecoder(resp.Body).Decode(&policy); err != nil {
		return nil, err
	}
	return policy, nil
}

// GetPolicy gets a backup policy
func (cb *cloudBackupService) GetPolicy(ctx context.Context, policyID string) (*CloudBackupPolicy, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodGet, cloudBackupServiceName,
		cb.itemPolicyPath(policyID), nil)
	if err != nil {
		return nil, err
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	var policy *CloudBackupPolicy
	if err := json.NewDecoder(resp.Body).Decode(&policy); err != nil {
		return nil, err
	}
	return policy, nil
}

// PatchPolicy patches a backup policy
func (cb *cloudBackupService) PatchPolicy(ctx context.Context, policyID string, payload *CloudBackupPatchPolicyPayload) (*CloudBackupPolicy, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodPatch, cloudBackupServiceName,
		cb.itemPolicyPath(policyID), payload)
	if err != nil {
		return nil, err
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	var policy *CloudBackupPolicy
	if err := json.NewDecoder(resp.Body).Decode(&policy); err != nil {
		return nil, err
	}
	return policy, nil
}

// DeletePolicy deletes a backup policy
func (cb *cloudBackupService) DeletePolicy(ctx context.Context, policyID string) error {
	req, err := cb.client.NewRequest(ctx, http.MethodDelete, cloudBackupServiceName,
		cb.itemPolicyPath(policyID), nil)
	if err != nil {
		return err
	}
	_, err = cb.client.Do(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

// ListAppliedPolicyDirectories lists the directories that have a backup policy applied
func (cb *cloudBackupService) ListAppliedPolicyDirectories(ctx context.Context, policyID string) ([]*CloudBackupDirectory, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodGet, cloudBackupServiceName,
		strings.Join([]string{cb.itemPolicyPath(policyID), "directories"}, "/"), nil)
	if err != nil {
		return nil, err
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	var directories []*CloudBackupDirectory
	if err := json.NewDecoder(resp.Body).Decode(&directories); err != nil {
		return nil, err
	}
	return directories, nil
}

// ActionPolicyDirectory applies an action to a backup policy
func (cb *cloudBackupService) ActionPolicyDirectory(ctx context.Context, policyID string, payload *CloudBackupActionPolicyDirectoryPayload) error {
	req, err := cb.client.NewRequest(ctx, http.MethodPost, cloudBackupServiceName,
		strings.Join([]string{cb.itemPolicyPath(policyID), "action"}, "/"), payload)
	if err != nil {
		return err
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return err
	}
	return resp.Body.Close()
}
