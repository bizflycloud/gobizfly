package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
)

type CloudBackupActivity struct {
	Action            string `json:"action"`
	BackupDirectoryId string `json:"backup_directory_id"`
	CreatedAt         string `json:"created_at"`
	Id                string `json:"id"`
	Extra             string `json:"extra,omitempty"`
	MachineId         string `json:"machine_id"`
	Message           string `json:"message"`
	PolicyId          string `json:"policy_id,omitempty"`
	Reason            string `json:"reason,omitempty"`
	RecoveryPoint     string `json:"recovery_point,omitempty"`
	Status            string `json:"status"`
	TenantId          string `json:"tenant_id"`
	UpdatedAt         string `json:"updated_at"`
	UserId            string `json:"user_id"`
}

func (cb *cloudBackupService) CloudBackupActivity(ctx context.Context) ([]*CloudBackupActivity, error) {
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
