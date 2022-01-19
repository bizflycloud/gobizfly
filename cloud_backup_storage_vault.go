package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
)

type StorageVault struct {
	CreatedAt        string `json:"created_at"`
	CredentialType   string `json:"credential_type"`
	DataUsage        string `json:"data_usage"`
	Deleted          bool   `json:"deleted"`
	EncryptionKey    string `json:"encryption_key"`
	Id               string `json:"id"`
	Name             string `json:"name"`
	SecretRef        string `json:"secret_ref"`
	StorageBucket    string `json:"storage_bucket"`
	StorageVaultType string `json:"storage_vault_type"`
	TenantId         string `json:"tenant_id"`
	UpdatedAt        string `json:"updated_at"`
}

type CreateStorageVaultPayload struct {
	AwsAccessKeyId     string `json:"aws_access_key_id"`
	AwsSecretAccessKey string `json:"aws_secret_access_key"`
	EndpointUrl        string `json:"endpoint_url"`
	StorageVaultType   string `json:"storage_vault_type"`
	Name               string `json:"name"`
	StorageBucket      string `json:"storage_bucket"`
	CredentialType     string `json:"credential_type"`
}

func (cb *cloudBackupService) ListStorageVaults(ctx context.Context) ([]*StorageVault, error) {
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
		StorageVaults []*StorageVault `json:"storage_vaults"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data.StorageVaults, nil
}

func (cb *cloudBackupService) GetStorageVault(ctx context.Context, valutId string) (*StorageVault, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodGet, cloudBackupServiceName,
		cb.itemStorageVaultPath(valutId), nil)
	if err != nil {
		return nil, err
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var vault *StorageVault
	if err := json.NewDecoder(resp.Body).Decode(&vault); err != nil {
		return nil, err
	}
	return vault, nil
}

func (cb *cloudBackupService) CreateStorageVault(ctx context.Context, payload *CreateStorageVaultPayload) (*StorageVault, error) {
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
	var vault *StorageVault
	if err := json.NewDecoder(resp.Body).Decode(&vault); err != nil {
		return nil, err
	}
	return vault, nil
}
