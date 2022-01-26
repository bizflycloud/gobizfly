package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
)

const (
	backupPath = "/backup"
)

var _ ScheduledVolumeBackup = (*scheduledVolumeBackup)(nil)

type scheduledVolumeBackup struct {
	client *Client
}

type ScheduledVolumeBackup interface {
	Create(ctx context.Context, payload *CreateBackupPayload) (*ExtendedBackup, error)
	Get(ctx context.Context, backupID string) (*ExtendedBackup, error)
	List(ctx context.Context) ([]*Backup, error)
	Delete(ctx context.Context, backupID string) error
	Update(ctx context.Context, backupID string, payload *UpdateBackupPayload) (*ExtendedBackup, error)
}

type BackupOption struct {
	Frequency string `json:"frequency"`
	Size      string `json:"size"`
}

type Backup struct {
	ID            string       `json:"_id"`
	CreatedAt     string       `json:"created_at"`
	NextRunAt     string       `json:"next_run_at"`
	BillingPlan   string       `json:"billing_plan"`
	Options       BackupOption `json:"options"`
	ResourceID    string       `json:"resource_id"`
	ResourceType  string       `json:"resource_type"`
	ScheduledHour int          `json:"scheduled_hour"`
	TenantID      string       `json:"tenant_id"`
	Type          string       `json:"type"`
	UpdatedAt     string       `json:"updated_at"`
}

type ExtendedBackup struct {
	Backup
	Snapshots []Snapshot `json:"snapshots"`
	VolumeId  string     `json:"volume_id"`
	Volume    Volume     `json:"volume"`
}

type CreateBackupPayload struct {
	ResourceID string `json:"resource_id"`
	Frequency  string `json:"frequency"`
	Size       string `json:"size"`
	Hour       int    `json:"hour,omitempty"`
}

type UpdateBackupPayload struct {
	Frequency string `json:"frequency,omitempty"`
	Size      string `json:"size,omitempty"`
	Hour      int    `json:"hour,omitempty"`
}

func (b scheduledVolumeBackup) resourcePath() string {
	return backupPath
}

func (b scheduledVolumeBackup) itemPath(id string) string {
	return b.resourcePath() + "/" + id
}

func (b scheduledVolumeBackup) Create(ctx context.Context, payload *CreateBackupPayload) (*ExtendedBackup, error) {
	req, err := b.client.NewRequest(ctx, http.MethodPost, serverServiceName, b.resourcePath(), payload)
	if err != nil {
		return nil, err
	}
	var dataResponse struct {
		Data *ExtendedBackup `json:"data"`
	}
	resp, err := b.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&dataResponse); err != nil {
		return nil, err
	}
	return dataResponse.Data, nil
}

func (b scheduledVolumeBackup) Get(ctx context.Context, backupID string) (*ExtendedBackup, error) {
	req, err := b.client.NewRequest(ctx, http.MethodGet, serverServiceName, b.itemPath(backupID), nil)
	if err != nil {
		return nil, err
	}
	var backup *ExtendedBackup
	resp, err := b.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&backup); err != nil {
		return nil, err
	}
	return backup, nil
}

func (b scheduledVolumeBackup) List(ctx context.Context) ([]*Backup, error) {
	req, err := b.client.NewRequest(ctx, http.MethodGet, serverServiceName, b.resourcePath(), nil)
	if err != nil {
		return nil, err
	}
	var backups []*Backup
	resp, err := b.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&backups); err != nil {
		return nil, err
	}
	return backups, nil
}

func (b scheduledVolumeBackup) Delete(ctx context.Context, backupID string) error {
	req, err := b.client.NewRequest(ctx, http.MethodDelete, serverServiceName, b.itemPath(backupID), nil)
	if err != nil {
		return err
	}
	resp, err := b.client.Do(ctx, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (b scheduledVolumeBackup) Update(ctx context.Context, backupID string, payload *UpdateBackupPayload) (*ExtendedBackup, error) {
	req, err := b.client.NewRequest(ctx, http.MethodPut, serverServiceName, b.itemPath(backupID), payload)
	if err != nil {
		return nil, err
	}
	var backup *ExtendedBackup
	resp, err := b.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&backup); err != nil {
		return nil, err
	}
	return backup, nil
}
