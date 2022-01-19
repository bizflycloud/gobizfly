package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

type RecoveryPoint struct {
	RecoveryPointType string `json:"recovery_point_type"`
	Status            string `json:"status"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	Name              string `json:"name"`
	Id                string `json:"id"`
}

type File struct {
	Id           string `json:"id"`
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

type RecoveryPointActionPayload struct {
	Action string `json:"action"`
}
type MachineRecoveryPoint struct {
	BackupDirectory   BackupDirectory `json:"backup_directory"`
	CreatedAt         string          `json:"created_at"`
	Id                string          `json:"id"`
	Name              string          `json:"name"`
	RecoveryPointType string          `json:"recovery_point_type"`
	Status            string          `json:"status"`
	UpdatedAt         string          `json:"updated_at"`
}

type ExtendedRecoveryPoint struct {
	MachineRecoveryPoint
	IndexHash   string `json:"index_hash,omitempty"`
	LocalSize   int    `json:"local_size,omitempty"`
	Method      string `json:"method,omitempty"`
	Progress    string `json:"progress,omitempty"`
	StorageSize int    `json:"storage_size,omitempty"`
	TotalFiles  int    `json:"total_files,omitempty"`
}

type RecoveryPointItem struct {
	AccessMode  string `json:"access_mode"`
	AccessTime  string `json:"access_time"`
	ChangeTime  string `json:"change_time"`
	ContentType string `json:"content_type"`
	CreatedAt   string `json:"created_at"`
	Gid         string `json:"gid"`
	Id          string `json:"id"`
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

type RestoreRecoveryPointPayload struct {
	RecoveryPointId string `json:"recovery_point_id"`
	Path            string `json:"path"`
}

type StateDirectoryAction struct {
	Action            string   `json:"action"`
	BackupDirectories []string `json:"backup_directories"`
}

type CreateDirectoryPayload struct {
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	Path        string   `json:"path"`
	Policies    []string `json:"policies,omitempty"`
}

type PatchDirectoryPayload struct {
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Policies    []string `json:"policies,omitempty"`
}

type DeleteDirectoryPayload struct {
	Keep bool `json:"keep"`
}

type ActionDirectoryPayload struct {
	Action      string `json:"action"`
	StorageType string `json:"storage_type,omitempty"`
	Name        string `json:"name,omitempty"`
}

type DeleteMultipleRecoveryPointPayload struct {
	RecoveryPointIds []string `json:"recovery_point_ids"`
}

type DeleteMultipleDirectoriesPayload struct {
	BackupDirectories []string `json:"backup_directories"`
	Keep              bool     `json:"keep"`
}

func (cb *cloudBackupService) ListTenantRecoveryPoints(ctx context.Context) ([]*MachineRecoveryPoint, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodGet, cloudBackupServiceName, cb.recoveryPointPath(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var recoveryPoints []*MachineRecoveryPoint
	if err := json.NewDecoder(resp.Body).Decode(&recoveryPoints); err != nil {
		return nil, err
	}
	return recoveryPoints, nil
}

func (cb *cloudBackupService) DeleteMultipleRecoveryPoints(ctx context.Context, payload DeleteMultipleRecoveryPointPayload) error {
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

func (cb *cloudBackupService) ListDirectoryRecoveryPoints(ctx context.Context, machineId string, directoryId string) ([]*MachineRecoveryPoint, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodGet, cloudBackupServiceName,
		cb.machineDirectoryRecoveryPointPath(machineId, directoryId), nil)
	if err != nil {
		return nil, err
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var data struct {
		RecoveryPoints []*MachineRecoveryPoint `json:"recovery_points"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data.RecoveryPoints, nil
}

func (cb *cloudBackupService) ListRecoveryPointFiles(ctx context.Context, recoveryPointId string) ([]*File, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodGet, cloudBackupServiceName,
		strings.Join([]string{cb.itemRecoveryPointPath(recoveryPointId), "files"}, "/"), nil)
	if err != nil {
		return nil, err
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var files []*File
	if err := json.NewDecoder(resp.Body).Decode(&files); err != nil {
		return nil, err
	}
	return files, nil
}

func (cb *cloudBackupService) RecoveryPointAction(ctx context.Context, recoveryPointId string, payload *RecoveryPointActionPayload) (*MachineRecoveryPoint, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodPatch, cloudBackupServiceName,
		strings.Join([]string{cb.itemRecoveryPointPath(recoveryPointId), "action"}, "/"), payload)
	if err != nil {
		return nil, err
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var recoveryPoint *MachineRecoveryPoint
	if err = json.NewDecoder(resp.Body).Decode(&recoveryPoint); err != nil {
		return nil, err
	}
	return recoveryPoint, nil
}

func (cb *cloudBackupService) ListMachineRecoveryPoints(ctx context.Context, machineId string) ([]*ExtendedRecoveryPoint, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodGet, cloudBackupServiceName,
		cb.itemMachinePath(machineId), nil)
	if err != nil {
		return nil, err
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var recoveryPoints struct {
		RecoveryPoints []*ExtendedRecoveryPoint `json:"recovery_points"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&recoveryPoints); err != nil {
		return nil, err
	}
	return recoveryPoints.RecoveryPoints, nil
}

func (cb *cloudBackupService) GetRecoveryPoint(ctx context.Context, recoveryPointId string) (*MachineRecoveryPoint, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodGet, cloudBackupServiceName,
		cb.itemRecoveryPointPath(recoveryPointId), nil)
	if err != nil {
		return nil, err
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var recoveryPoint *MachineRecoveryPoint
	if err := json.NewDecoder(resp.Body).Decode(&recoveryPoint); err != nil {
		return nil, err
	}
	return recoveryPoint, nil
}

func (cb *cloudBackupService) DeleteRecoveryPoint(ctx context.Context, recoveryPointId string) error {
	req, err := cb.client.NewRequest(ctx, http.MethodDelete, cloudBackupServiceName,
		cb.itemRecoveryPointPath(recoveryPointId), nil)
	if err != nil {
		return err
	}
	_, err = cb.client.Do(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (cb *cloudBackupService) ListRecoveryPointItems(ctx context.Context, recoveryPointId string) ([]*RecoveryPointItem, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodGet, cloudBackupServiceName,
		strings.Join([]string{cb.itemRecoveryPointPath(recoveryPointId), "items"}, "/"), nil)
	if err != nil {
		return nil, err
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var data struct {
		Items []*RecoveryPointItem `json:"items"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data.Items, nil
}

func (cb *cloudBackupService) RestoreRecoveryPoint(ctx context.Context, recoveryPointId string, payload *RestoreRecoveryPointPayload) error {
	req, err := cb.client.NewRequest(ctx, http.MethodPost, cloudBackupServiceName,
		strings.Join([]string{cb.itemRecoveryPointPath(recoveryPointId), "action"}, "/"), payload)
	if err != nil {
		return err
	}
	_, err = cb.client.Do(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
