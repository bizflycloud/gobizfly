package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
)

// CloudBackupActivity represents a Cloud Backup Activity data.
type CloudBackupActivity struct {
	Action            string `json:"action"`
	BackupDirectoryID string `json:"backup_directory_id"`
	CreatedAt         string `json:"created_at"`
	ID                string `json:"id"`
	Extra             string `json:"extra,omitempty"`
	MachineID         string `json:"machine_id"`
	Message           string `json:"message"`
	PolicyID          string `json:"policy_id,omitempty"`
	Reason            string `json:"reason,omitempty"`
	RecoveryPoint     string `json:"recovery_point,omitempty"`
	Status            string `json:"status"`
	TenantID          string `json:"tenant_id"`
	UpdatedAt         string `json:"updated_at"`
	UserID            string `json:"user_id"`
}

// CloudBackupListActivities - List Cloud Backup Activities
func (cb *cloudBackupService) CloudBackupListActivities(ctx context.Context) ([]*CloudBackupActivity, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodGet, cloudBackupServiceName, cb.activityPath(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var data struct {
		ActivityData []*CloudBackupActivity `json:"activities"`
	}
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data.ActivityData, nil
}
