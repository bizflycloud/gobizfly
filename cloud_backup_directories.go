package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

type CloudBackupDirectory struct {
	Activated   bool               `json:"activated"`
	CreatedAt   string             `json:"created_at"`
	Deleted     bool               `json:"deleted"`
	Description string             `json:"description"`
	Id          string             `json:"id"`
	Machine     CloudBackupMachine `json:"machine"`
	Name        string             `json:"name"`
	Path        string             `json:"path"`
	Quota       string             `json:"quota"`
	Size        int                `json:"size"`
	UpdatedAt   string             `json:"updated_at"`
}

type CloudBackupActionMultipleDirectoriesPayload struct {
	Action            string   `json:"action"`
	BackupDirectories []string `json:"backup_directories"`
}

func (cb *cloudBackupService) ActionDirectory(ctx context.Context, machineId string, payload *CloudBackupStateDirectoryAction) error {
	req, err := cb.client.NewRequest(ctx, http.MethodPost, cloudBackupServiceName,
		strings.Join([]string{cb.itemMachinePath(machineId), "directories", "action"}, "/"), payload)
	if err != nil {
		return err
	}
	_, err = cb.client.Do(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (cb *cloudBackupService) ListMachineBackupDirectories(ctx context.Context, machineId string) ([]*CloudBackupDirectory, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodGet, cloudBackupServiceName,
		strings.Join([]string{cb.itemMachinePath(machineId), "directories"}, "/"), nil)
	if err != nil {
		return nil, err
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var data struct {
		Directories []*CloudBackupDirectory `json:"directories"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data.Directories, nil
}

func (cb *cloudBackupService) CreateBackupDirectory(ctx context.Context, machineId string, payload *CloudBackupCreateDirectoryPayload) (*CloudBackupDirectory, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodPost, cloudBackupServiceName,
		strings.Join([]string{cb.itemMachinePath(machineId), "directories"}, "/"), payload)
	if err != nil {
		return nil, err
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var directory *CloudBackupDirectory
	if err := json.NewDecoder(resp.Body).Decode(&directory); err != nil {
		return nil, err
	}
	return directory, nil
}

func (cb *cloudBackupService) GetBackupDirectory(ctx context.Context, machineId string, directoryId string) (*CloudBackupDirectory, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodGet, cloudBackupServiceName,
		strings.Join([]string{cb.itemMachinePath(machineId), "directories", directoryId}, "/"), nil)
	if err != nil {
		return nil, err
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var directory *CloudBackupDirectory
	if err := json.NewDecoder(resp.Body).Decode(&directory); err != nil {
		return nil, err
	}
	return directory, nil
}

func (cb *cloudBackupService) PatchBackupDirectory(ctx context.Context, machineId string, directoryId string, payload *CloudBackupPatchDirectoryPayload) (*CloudBackupDirectory, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodPatch, cloudBackupServiceName,
		strings.Join([]string{cb.itemMachinePath(machineId), "directories", directoryId}, "/"), payload)
	if err != nil {
		return nil, err
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var directory *CloudBackupDirectory
	if err := json.NewDecoder(resp.Body).Decode(&directory); err != nil {
		return nil, err
	}
	return directory, nil
}

func (cb *cloudBackupService) DeleteBackupDirectory(ctx context.Context, machineId string, directoryId string, payload *CloudBackupDeleteDirectoryPayload) error {
	req, err := cb.client.NewRequest(ctx, http.MethodDelete, cloudBackupServiceName,
		strings.Join([]string{cb.itemMachinePath(machineId), "directories", directoryId}, "/"), payload)
	if err != nil {
		return err
	}
	_, err = cb.client.Do(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (cb *cloudBackupService) ListTenantDirectories(ctx context.Context) ([]*CloudBackupDirectory, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodGet, cloudBackupServiceName,
		cb.directoryPath(), nil)
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

func (cb *cloudBackupService) ActionBackupDirectory(ctx context.Context, machineId string, directoryId string, payload *CloudBackupActionDirectoryPayload) error {
	req, err := cb.client.NewRequest(ctx, http.MethodPost, cloudBackupServiceName,
		strings.Join([]string{cb.itemMachinePath(machineId), "directories", directoryId, "action"}, "/"), payload)
	if err != nil {
		return err
	}
	_, err = cb.client.Do(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (cb *cloudBackupService) DeleteMultipleDirectories(ctx context.Context, machineId string, payload *CloudBackupDeleteMultipleDirectoriesPayload) error {
	req, err := cb.client.NewRequest(ctx, http.MethodDelete, cloudBackupServiceName,
		strings.Join([]string{cb.itemMachinePath(machineId), "directories", "action"}, "/"), payload)
	if err != nil {
		return err
	}
	_, err = cb.client.Do(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (cb *cloudBackupService) ActionMultipleDirectories(ctx context.Context, machineId string, payload *CloudBackupActionMultipleDirectoriesPayload) error {
	req, err := cb.client.NewRequest(ctx, http.MethodPost, cloudBackupServiceName,
		strings.Join([]string{cb.itemMachinePath(machineId), "directories", "action"}, "/"), payload)
	if err != nil {
		return err
	}
	_, err = cb.client.Do(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
