package gobizfly

import (
	"fmt"
	"github.com/bizflycloud/gobizfly/testlib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestListStorageVault(t *testing.T) {
	setup()
	defer teardown()
	var cb cloudBackupService
	mux.HandleFunc(testlib.CloudBackupURL(cb.storageVaultsPath()),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			resp := `{

  "storage_vaults": [

    {

      "created_at": "2021-10-29T02:09:53.356265+00:00",

      "credential_type": "DEFAULT",

      "data_usage": "80783504 Kb",

      "deleted": false,

      "encryption_key": "2c21d5f1-7237-4d01-bc13-036d47753099",

      "id": "7bef2ea3-3eab-473a-a188-02952290b560",

      "name": "2c21d5f1-7237-4d01-bc13-036d47753099",

      "secret_ref": null,

      "storage_bucket": "endeavour-dev-2c21d5f1-7237-4d01-bc13-036d47753099",

      "storage_vault_type": "S3",

      "tenant_id": "2c21d5f1-7237-4d01-bc13-036d47753099",

      "updated_at": "2021-10-29T02:09:53.356265+00:00"

    }

  ],

  "total": 1

}`
			fmt.Fprint(writer, resp)
		})
	vaults, err := client.CloudBackup.ListStorageVaults(ctx)
	require.NoError(t, err)
	assert.Equal(t, "80783504 Kb", vaults[0].DataUsage)
}

func TestGetStorageVaults(t *testing.T) {
	setup()
	defer teardown()

	var cb *cloudBackupService
	mux.HandleFunc(testlib.CloudBackupURL(cb.itemStorageVaultPath("123")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			resp := `{

  "created_at": "2021-10-29T02:09:53.356265+00:00",

  "credential_type": "DEFAULT",

  "data_usage": "80783504 Kb",

  "deleted": false,

  "encryption_key": "2c21d5f1-7237-4d01-bc13-036d47753099",

  "id": "7bef2ea3-3eab-473a-a188-02952290b560",

  "name": "2c21d5f1-7237-4d01-bc13-036d47753099",

  "secret_ref": null,

  "storage_bucket": "endeavour-dev-2c21d5f1-7237-4d01-bc13-036d47753099",

  "storage_vault_type": "S3",

  "tenant_id": "2c21d5f1-7237-4d01-bc13-036d47753099",

  "updated_at": "2021-10-29T02:09:53.356265+00:00"

}`
			fmt.Fprint(writer, resp)
		})
	vault, err := client.CloudBackup.GetStorageVault(ctx, "123")
	require.NoError(t, err)
	assert.Equal(t, "S3", vault.StorageVaultType)
}

func TestCreateStorageVault(t *testing.T) {
	setup()
	defer teardown()

	var cb *cloudBackupService
	mux.HandleFunc(testlib.CloudBackupURL(cb.storageVaultsPath()),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			resp := `{

  "created_at": "2021-10-29T02:09:53.356265+00:00",

  "credential_type": "DEFAULT",

  "data_usage": "80783504 Kb",

  "deleted": false,

  "encryption_key": "2c21d5f1-7237-4d01-bc13-036d47753099",

  "id": "7bef2ea3-3eab-473a-a188-02952290b560",

  "name": "2c21d5f1-7237-4d01-bc13-036d47753099",

  "secret_ref": null,

  "storage_bucket": "endeavour-dev-2c21d5f1-7237-4d01-bc13-036d47753099",

  "storage_vault_type": "S3",

  "tenant_id": "2c21d5f1-7237-4d01-bc13-036d47753099",

  "updated_at": "2021-10-29T02:09:53.356265+00:00"

}`
			fmt.Fprint(writer, resp)
		})
	vault, err := client.CloudBackup.CreateStorageVault(ctx, &CloudBackupCreateStorageVaultPayload{
		AwsAccessKeyId:     "key_id",
		AwsSecretAccessKey: "access_key",
		EndpointUrl:        "/test",
		StorageVaultType:   "S3",
		Name:               "custom",
		StorageBucket:      "bucket_name",
		CredentialType:     "CUSTOM",
	})
	require.NoError(t, err)
	assert.Equal(t, "DEFAULT", vault.CredentialType)
}
