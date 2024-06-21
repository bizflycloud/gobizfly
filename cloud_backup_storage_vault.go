package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
)

// CloudBackupStorageVault represents a Cloud Backup Storage Vault.
type CloudBackupStorageVault struct {
	CreatedAt        string `json:"created_at"`
	CredentialType   string `json:"credential_type"`
	DataUsage        string `json:"data_usage"`
	Deleted          bool   `json:"deleted"`
	EncryptionKey    string `json:"encryption_key"`
	ID               string `json:"id"`
	Name             string `json:"name"`
	SecretRef        string `json:"secret_ref"`
	StorageBucket    string `json:"storage_bucket"`
	StorageVaultType string `json:"storage_vault_type"`
	TenantID         string `json:"tenant_id"`
	UpdatedAt        string `json:"updated_at"`
}

// CloudBackupCreateStorageVaultPayload represents the payload for creating a Cloud Backup Storage Vault.
type CloudBackupCreateStorageVaultPayload struct {
	AwsAccessKeyID     string `json:"aws_access_key_id"`
	AwsSecretAccessKey string `json:"aws_secret_access_key"`
	EndpointURL        string `json:"endpoint_url"`
	StorageVaultType   string `json:"storage_vault_type"`
	Name               string `json:"name"`
	StorageBucket      string `json:"storage_bucket"`
	CredentialType     string `json:"credential_type"`
}

// ListStorageVaults lists all Cloud Backup Storage Vaults.
func (cb *cloudBackupService) ListStorageVaults(ctx context.Context) ([]*CloudBackupStorageVault, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodGet, cloudBackupServiceName,
		cb.storageVaultsPath(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	var data struct {
		StorageVaults []*CloudBackupStorageVault `json:"storage_vaults"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data.StorageVaults, nil
}

// GetStorageVault gets a Cloud Backup Storage Vault.
func (cb *cloudBackupService) GetStorageVault(ctx context.Context, valutID string) (*CloudBackupStorageVault, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodGet, cloudBackupServiceName,
		cb.itemStorageVaultPath(valutID), nil)
	if err != nil {
		return nil, err
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var vault *CloudBackupStorageVault
	if err := json.NewDecoder(resp.Body).Decode(&vault); err != nil {
		return nil, err
	}
	return vault, nil
}

// CreateStorageVault creates a Cloud Backup Storage Vault.
func (cb *cloudBackupService) CreateStorageVault(ctx context.Context, payload *CloudBackupCreateStorageVaultPayload) (*CloudBackupStorageVault, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodPost, cloudBackupServiceName,
		cb.storageVaultsPath(), payload)
	if err != nil {
		return nil, err
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var vault *CloudBackupStorageVault
	if err := json.NewDecoder(resp.Body).Decode(&vault); err != nil {
		return nil, err
	}
	return vault, nil
}
