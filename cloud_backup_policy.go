package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

type CloudBackupCreatePolicyPayload struct {
	Name            string `json:"name"`
	StorageType     string `json:"storage_type"`
	SchedulePattern string `json:"schedule_pattern"`
	RetentionDays   int    `json:"retention_days"`
	Description     string `json:"description,omitempty"`
}

type CloudBackupPolicy struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	SchedulePattern string `json:"schedule_pattern"`
	RetentionDays   int    `json:"retention_days"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	TenantId        string `json:"tenant_id"`
	Description     string `json:"description"`
	LimitDownload   int    `json:"limit_download,omitempty"`
	LimitUpload     int    `json:"limit_upload,omitempty"`
	PolicyType      string `json:"policy_type"`
	Retentions      int    `json:"retentions"`
}

type CloudBackupPatchPolicyPayload struct {
	Name            string `json:"name,omitempty"`
	SchedulePattern string `json:"schedule_pattern,omitempty"`
	RetentionDays   int    `json:"retention_days,omitempty"`
}

type CloudBackupFullPolicy struct {
	CloudBackupPolicy
	RetentionDays     int      `json:"retention_days"`
	RetentionHours    int      `json:"retention_hours"`
	RetentionWeeks    int      `json:"retention_weeks"`
	RetentionMonths   int      `json:"retention_months"`
	BackupDirectories []string `json:"backup_directories"`
}

type CloudBackupActionPolicyDirectoryPayload struct {
	Action      string `json:"action"`
	DirectoryId string `json:"directory_id"`
}

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
	defer resp.Body.Close()
	var policy *CloudBackupPolicy
	if err := json.NewDecoder(resp.Body).Decode(&policy); err != nil {
		return nil, err
	}
	return policy, nil
}

func (cb *cloudBackupService) GetBackupDirectoryPolicy(ctx context.Context, machineId string, directoryId string) (*CloudBackupPolicy, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodGet, cloudBackupServiceName,
		strings.Join([]string{cb.itemMachinePath(machineId), "directories", directoryId, "policies"}, "/"), nil)
	if err != nil {
		return nil, err
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var policy *CloudBackupPolicy
	if err := json.NewDecoder(resp.Body).Decode(&policy); err != nil {
		return nil, err
	}
	return policy, nil
}

func (cb *cloudBackupService) GetPolicy(ctx context.Context, policyId string) (*CloudBackupPolicy, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodGet, cloudBackupServiceName,
		cb.itemPolicyPath(policyId), nil)
	if err != nil {
		return nil, err
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var policy *CloudBackupPolicy
	if err := json.NewDecoder(resp.Body).Decode(&policy); err != nil {
		return nil, err
	}
	return policy, nil
}

func (cb *cloudBackupService) PatchPolicy(ctx context.Context, policyId string, payload *CloudBackupPatchPolicyPayload) (*CloudBackupPolicy, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodPatch, cloudBackupServiceName,
		cb.itemPolicyPath(policyId), payload)
	if err != nil {
		return nil, err
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var policy *CloudBackupPolicy
	if err := json.NewDecoder(resp.Body).Decode(&policy); err != nil {
		return nil, err
	}
	return policy, nil
}

func (cb *cloudBackupService) DeletePolicy(ctx context.Context, policyId string) error {
	req, err := cb.client.NewRequest(ctx, http.MethodDelete, cloudBackupServiceName,
		cb.itemPolicyPath(policyId), nil)
	if err != nil {
		return err
	}
	_, err = cb.client.Do(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (cb *cloudBackupService) ListAppliedPolicyDirectories(ctx context.Context, policyId string) ([]*CloudBackupDirectory, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodGet, cloudBackupServiceName,
		strings.Join([]string{cb.itemPolicyPath(policyId), "directories"}, "/"), nil)
	if err != nil {
		return nil, err
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var directories []*CloudBackupDirectory
	if err := json.NewDecoder(resp.Body).Decode(&directories); err != nil {
		return nil, err
	}
	return directories, nil
}

func (cb *cloudBackupService) ActionPolicyDirectory(ctx context.Context, policyId string, payload *CloudBackupActionPolicyDirectoryPayload) error {
	req, err := cb.client.NewRequest(ctx, http.MethodPost, cloudBackupServiceName,
		strings.Join([]string{cb.itemPolicyPath(policyId), "action"}, "/"), payload)
	if err != nil {
		return err
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return err
	}
	return resp.Body.Close()
}
