package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

// CloudBackupRecoveryPoint represents a Cloud Backup Recovery Point.
type CloudBackupRecoveryPoint struct {
	RecoveryPointType string `json:"recovery_point_type"`
	Status            string `json:"status"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	Name              string `json:"name"`
	ID                string `json:"id"`
}

// CloudBackupFile represents a Cloud Backup File.
type CloudBackupFile struct {
	ID           string `json:"id"`
	ItemName     string `json:"item_name"`
	Size         int    `json:"size"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	ContentType  string `json:"content_type"`
	ETag         string `json:"eTag"`
	ItemType     string `json:"item_type"`
	LastModified string `json:"last_modified"`
	Mode         string `json:"mode"`
	Status       string `json:"status"`
}

// CloudBackupRecoveryPointActionPayload represents a Cloud Backup Recovery Point Action Payload.
type CloudBackupRecoveryPointActionPayload struct {
	Action string `json:"action"`
}

// CloudBackupMachineRecoveryPoint represents a Cloud Backup Recovery Point Action.
type CloudBackupMachineRecoveryPoint struct {
	BackupDirectory   CloudBackupDirectory `json:"backup_directory"`
	CreatedAt         string               `json:"created_at"`
	ID                string               `json:"id"`
	Name              string               `json:"name"`
	RecoveryPointType string               `json:"recovery_point_type"`
	Status            string               `json:"status"`
	UpdatedAt         string               `json:"updated_at"`
}

// CloudBackupExtendedRecoveryPoint represents a Cloud Backup Extended Recovery Point.
type CloudBackupExtendedRecoveryPoint struct {
	CloudBackupMachineRecoveryPoint
	IndexHash   string `json:"index_hash,omitempty"`
	LocalSize   int    `json:"local_size,omitempty"`
	Method      string `json:"method,omitempty"`
	Progress    string `json:"progress,omitempty"`
	StorageSize int    `json:"storage_size,omitempty"`
	TotalFiles  int    `json:"total_files,omitempty"`
}

// CloudBackupRecoveryPointItem represents a Cloud Backup Recovery Point Item.
type CloudBackupRecoveryPointItem struct {
	AccessMode  string `json:"access_mode"`
	AccessTime  string `json:"access_time"`
	ChangeTime  string `json:"change_time"`
	ContentType string `json:"content_type"`
	CreatedAt   string `json:"created_at"`
	Gid         string `json:"gid"`
	ID          string `json:"id"`
	IsDir       bool   `json:"is_dir"`
	ItemName    string `json:"item_name"`
	ItemType    string `json:"item_type"`
	Mode        string `json:"mode"`
	ModifyTime  string `json:"modify_time"`
	RealName    string `json:"real_name"`
	Size        string `json:"size"`
	Status      string `json:"status"`
	SymlinkPath string `json:"symlink_path"`
	Uid         string `json:"uid"`
	UpdatedAt   string `json:"updated_at"`
}

// CloudBackupRestoreRecoveryPointPayload represents a Cloud Backup Restore Recovery Point Payload.
type CloudBackupRestoreRecoveryPointPayload struct {
	RecoveryPointID string `json:"recovery_point_id"`
	Path            string `json:"path"`
}

// CloudBackupStateDirectoryAction represents a Cloud Backup State Directory Action.
type CloudBackupStateDirectoryAction struct {
	Action            string   `json:"action"`
	BackupDirectories []string `json:"backup_directories"`
}

// CloudBackupCreateDirectoryPayload represents a Cloud Backup Create Directory Payload.
type CloudBackupCreateDirectoryPayload struct {
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	Path        string   `json:"path"`
	Policies    []string `json:"policies,omitempty"`
}

// CloudBackupPatchDirectoryPayload represents a Cloud Backup Patch Directory Payload.
type CloudBackupPatchDirectoryPayload struct {
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Policies    []string `json:"policies,omitempty"`
}

// CloudBackupDeleteDirectoryPayload represents a Cloud Backup Delete Directory Payload.
type CloudBackupDeleteDirectoryPayload struct {
	Keep bool `json:"keep"`
}

// CloudBackupActionDirectoryPayload represents a Cloud Backup Action Directory Payload.
type CloudBackupActionDirectoryPayload struct {
	Action      string `json:"action"`
	StorageType string `json:"storage_type,omitempty"`
	Name        string `json:"name,omitempty"`
}

// CloudBackupDeleteMultipleRecoveryPointPayload represents a Cloud Backup Delete Multiple Recovery Point Payload.
type CloudBackupDeleteMultipleRecoveryPointPayload struct {
	RecoveryPointIDs []string `json:"recovery_point_ids"`
}

// CloudBackupDeleteMultipleDirectoriesPayload represents a Cloud Backup Delete Multiple Recovery Point Payload.
type CloudBackupDeleteMultipleDirectoriesPayload struct {
	BackupDirectories []string `json:"backup_directories"`
	Keep              bool     `json:"keep"`
}

// ListTenantRecoveryPoints lists all recovery points belonging to a tenant.
func (cb *cloudBackupService) ListTenantRecoveryPoints(ctx context.Context) ([]*CloudBackupMachineRecoveryPoint, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodGet, cloudBackupServiceName, cb.recoveryPointPath(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var recoveryPoints []*CloudBackupMachineRecoveryPoint
	if err := json.NewDecoder(resp.Body).Decode(&recoveryPoints); err != nil {
		return nil, err
	}
	return recoveryPoints, nil
}

// DeleteMultipleRecoveryPoints deletes multiple recovery points.
func (cb *cloudBackupService) DeleteMultipleRecoveryPoints(ctx context.Context, payload CloudBackupDeleteMultipleRecoveryPointPayload) error {
	req, err := cb.client.NewRequest(ctx, http.MethodDelete, cloudBackupServiceName, cb.recoveryPointPath(), payload)
	if err != nil {
		return err
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

// ListDirectoryRecoveryPoints lists all recovery points belonging to a directory.
func (cb *cloudBackupService) ListDirectoryRecoveryPoints(ctx context.Context, machineID string, directoryID string) ([]*CloudBackupMachineRecoveryPoint, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodGet, cloudBackupServiceName,
		cb.machineDirectoryRecoveryPointPath(machineID, directoryID), nil)
	if err != nil {
		return nil, err
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var data struct {
		RecoveryPoints []*CloudBackupMachineRecoveryPoint `json:"recovery_points"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data.RecoveryPoints, nil
}

// ListRecoveryPointFiles lists all files belonging to a recovery point.
func (cb *cloudBackupService) ListRecoveryPointFiles(ctx context.Context, recoveryPointID string) ([]*CloudBackupFile, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodGet, cloudBackupServiceName,
		strings.Join([]string{cb.itemRecoveryPointPath(recoveryPointID), "files"}, "/"), nil)
	if err != nil {
		return nil, err
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var files []*CloudBackupFile
	if err := json.NewDecoder(resp.Body).Decode(&files); err != nil {
		return nil, err
	}
	return files, nil
}

// RecoveryPointAction performs an action on a recovery point.
func (cb *cloudBackupService) RecoveryPointAction(ctx context.Context, recoveryPointID string, payload *CloudBackupRecoveryPointActionPayload) (*CloudBackupMachineRecoveryPoint, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodPatch, cloudBackupServiceName,
		strings.Join([]string{cb.itemRecoveryPointPath(recoveryPointID), "action"}, "/"), payload)
	if err != nil {
		return nil, err
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var recoveryPoint *CloudBackupMachineRecoveryPoint
	if err = json.NewDecoder(resp.Body).Decode(&recoveryPoint); err != nil {
		return nil, err
	}
	return recoveryPoint, nil
}

// ListMachineRecoveryPoints lists all recovery points belonging to a machine.
func (cb *cloudBackupService) ListMachineRecoveryPoints(ctx context.Context, machineID string) ([]*CloudBackupExtendedRecoveryPoint, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodGet, cloudBackupServiceName,
		cb.itemMachinePath(machineID), nil)
	if err != nil {
		return nil, err
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var recoveryPoints struct {
		RecoveryPoints []*CloudBackupExtendedRecoveryPoint `json:"recovery_points"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&recoveryPoints); err != nil {
		return nil, err
	}
	return recoveryPoints.RecoveryPoints, nil
}

// GetRecoveryPoint gets a recovery point.
func (cb *cloudBackupService) GetRecoveryPoint(ctx context.Context, recoveryPointID string) (*CloudBackupMachineRecoveryPoint, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodGet, cloudBackupServiceName,
		cb.itemRecoveryPointPath(recoveryPointID), nil)
	if err != nil {
		return nil, err
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var recoveryPoint *CloudBackupMachineRecoveryPoint
	if err := json.NewDecoder(resp.Body).Decode(&recoveryPoint); err != nil {
		return nil, err
	}
	return recoveryPoint, nil
}

// DeleteRecoveryPoint deletes a recovery point.
func (cb *cloudBackupService) DeleteRecoveryPoint(ctx context.Context, recoveryPointID string) error {
	req, err := cb.client.NewRequest(ctx, http.MethodDelete, cloudBackupServiceName,
		cb.itemRecoveryPointPath(recoveryPointID), nil)
	if err != nil {
		return err
	}
	_, err = cb.client.Do(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

// ListRecoveryPointItems lists all items belonging to a recovery point.
func (cb *cloudBackupService) ListRecoveryPointItems(ctx context.Context, recoveryPointID string) ([]*CloudBackupRecoveryPointItem, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodGet, cloudBackupServiceName,
		strings.Join([]string{cb.itemRecoveryPointPath(recoveryPointID), "items"}, "/"), nil)
	if err != nil {
		return nil, err
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var data struct {
		Items []*CloudBackupRecoveryPointItem `json:"items"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data.Items, nil
}

// RestoreRecoveryPoint restores a recovery point.
func (cb *cloudBackupService) RestoreRecoveryPoint(ctx context.Context, recoveryPointID string, payload *CloudBackupRestoreRecoveryPointPayload) error {
	req, err := cb.client.NewRequest(ctx, http.MethodPost, cloudBackupServiceName,
		strings.Join([]string{cb.itemRecoveryPointPath(recoveryPointID), "action"}, "/"), payload)
	if err != nil {
		return err
	}
	_, err = cb.client.Do(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
