package gobizfly

import (
	"fmt"
	"github.com/bizflycloud/gobizfly/testlib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"strings"
	"testing"
)

func TestListMachineBackupDirectories(t *testing.T) {
	setup()
	defer teardown()
	var cb cloudBackupService
	mux.HandleFunc(testlib.CloudBackupURL(strings.Join([]string{cb.itemMachinePath("123"), "directories"}, "/")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			resp := `{

  "directories": [

    {

      "activated": true,

      "created_at": "2021-12-03T08:23:05.122373+00:00",

      "deleted": false,

      "deleted_at": null,

      "description": null,

      "id": "1a5c093a-bbbb-4e0e-9fd0-b943f09cb282",

      "machine_id": "7b2e1783-9066-4bf6-bb09-84c5d2d3165f",

      "name": "backup_dir_test_1",

      "path": "/home/vinh/cho-meo-bo-ngua",

      "policies": [

        {

          "created_at": "2021-12-07T03:58:07.911829+00:00",

          "description": null,

          "id": "7a5c22f6-e52b-4bb5-8b49-eed9ee142feb",

          "limit_download": 456,

          "limit_upload": null,

          "name": "policy1",

          "policy_type": "AUTO",

          "retention_days": 4,

          "retentions": 2,

          "schedule_pattern": "0 0 5 5 5",

          "tenant_id": "2c21d5f1-7237-4d01-bc13-036d47753099",

          "updated_at": "2021-12-27T07:37:59.960788+00:00"

        }

      ],

      "quota": null,

      "size": 0,

      "tenant_id": "2c21d5f1-7237-4d01-bc13-036d47753099",

      "updated_at": "2021-12-03T08:23:05.122373+00:00"

    },

    {

      "activated": true,

      "created_at": "2021-11-01T04:21:57.905101+00:00",

      "deleted": false,

      "deleted_at": null,

      "description": null,

      "id": "326bf60e-7065-478a-b4b2-9d59cf8b63aa",

      "machine_id": "7b2e1783-9066-4bf6-bb09-84c5d2d3165f",

      "name": "backup_dir_test",

      "path": "/home/vinh/testbackup/testbackup4",

      "policies": [],

      "quota": null,

      "size": 0,

      "tenant_id": "2c21d5f1-7237-4d01-bc13-036d47753099",

      "updated_at": "2021-11-01T04:21:57.905101+00:00"

    }

  ],

  "total": 2

}`
			_, _ = fmt.Fprint(writer, resp)
		})
	machines, err := client.CloudBackup.ListMachineBackupDirectories(ctx, "123")
	require.NoError(t, err)
	assert.Equal(t, "/home/vinh/testbackup/testbackup4", machines[1].Path)
}

func TestActionDirectory(t *testing.T) {
	setup()
	defer teardown()
	var cb *cloudBackupService
	mux.HandleFunc(testlib.CloudBackupURL(strings.Join([]string{cb.itemMachinePath("123"), "directories", "action"}, "/")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
		})
	err := client.CloudBackup.ActionDirectory(ctx, "123", &CloudBackupStateDirectoryAction{
		Action:            "active",
		BackupDirectories: []string{"123"},
	})
	require.NoError(t, err)
}

func TestCreateBackupDirectory(t *testing.T) {
	setup()
	defer teardown()
	var cb *cloudBackupService
	mux.HandleFunc(testlib.CloudBackupURL(strings.Join([]string{cb.itemMachinePath("123"), "directories"}, "/")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			resp := `{

  "created_at": "2020-06-16T15:14:50.875699",

  "description": null,

  "id": "e0a69ae7-fada-4426-beca-704f67e6d46c",

  "machine_id": "e5e0c123-f7f8-425d-bb31-bc514c6e48ef",

  "name": "backup web",

  "path": "var/www/html",

  "quota": null,

  "size": 0,

  "tenant_id": "98dbba57-60c7-4c22-8087-8313997ea776",

  "updated_at": "2020-06-16T15:1"

}`

			_, _ = fmt.Fprint(writer, resp)
		})
	directory, err := client.CloudBackup.CreateBackupDirectory(ctx, "123", &CloudBackupCreateDirectoryPayload{
		Name: "test",
		Path: "/home/vinh/cho-meo-bo-ngua",
	})
	require.NoError(t, err)
	assert.Equal(t, directory.Name, "backup web")
}

func TestGetBackupDirectory(t *testing.T) {
	setup()
	defer teardown()

	var cb *cloudBackupService

	mux.HandleFunc(testlib.CloudBackupURL(strings.Join([]string{cb.itemMachinePath("123"), "directories", "456"}, "/")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			resp := `{

  "activated": true,

  "created_at": "2021-12-03T08:23:05.122373+00:00",

  "deleted": false,

  "deleted_at": null,

  "description": null,

  "id": "1a5c093a-bbbb-4e0e-9fd0-b943f09cb282",

  "machine": {

    "agent_version": "version: dev, commit: , built: ",

    "created_at": "2021-11-01T04:21:22.368979+00:00",

    "deleted": false,

    "encryption": false,

    "host_name": "vinhvn",

    "id": "7b2e1783-9066-4bf6-bb09-84c5d2d3165f",

    "ip_address": "192.168.18.13",

    "machine_storage_size": 1879048192,

    "name": "machine123",

    "operation_status": true,

    "os_version": "Arch Linux\n",

    "status": "OFFLINE",

    "tenant_id": "2c21d5f1-7237-4d01-bc13-036d47753099",

    "updated_at": "2022-01-13T04:05:18.838997+00:00"

  },

  "name": "backup_dir_test_1",

  "path": "/home/vinh/cho-meo-bo-ngua",

  "quota": null,

  "size": 0,

  "tenant_id": "2c21d5f1-7237-4d01-bc13-036d47753099",

  "updated_at": "2021-12-03T08:23:05.122373+00:00"

}`
			_, _ = fmt.Fprint(writer, resp)
		})
	directory, err := client.CloudBackup.GetBackupDirectory(ctx, "123", "456")
	require.NoError(t, err)
	assert.Equal(t, "/home/vinh/cho-meo-bo-ngua", directory.Path)
}

func TestCloudBackupService_PatchBackupDirectory(t *testing.T) {
	setup()
	defer teardown()
	var cb cloudBackupService

	mux.HandleFunc(testlib.CloudBackupURL(strings.Join([]string{cb.itemMachinePath("123"), "directories", "456"}, "/")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPatch, r.Method)
			resp := `{

  "created_at": "2020-06-16T15:14:50.875699",

  "description": null,

  "id": "e0a69ae7-fada-4426-beca-704f67e6d46c",

  "machine_id": "e5e0c123-f7f8-425d-bb31-bc514c6e48ef",

  "name": "backup web",

  "path": "var/www/html",

  "quota": null,

  "size": 0,

  "tenant_id": "98dbba57-60c7-4c22-8087-8313997ea776",

  "updated_at": "2020-06-16T15:1"

}`
			_, _ = fmt.Fprint(writer, resp)
		})
	directory, err := client.CloudBackup.PatchBackupDirectory(ctx, "123", "456", &CloudBackupPatchDirectoryPayload{
		Name: "test_123",
	})
	require.NoError(t, err)
	assert.Equal(t, "var/www/html", directory.Path)
}

func TestCloudBackupService_DeleteBackupDirectory(t *testing.T) {
	setup()
	defer teardown()

	var cb cloudBackupService
	mux.HandleFunc(testlib.CloudBackupURL(strings.Join([]string{cb.itemMachinePath("123"), "directories", "456"}, "/")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodDelete, r.Method)
		})
	err := client.CloudBackup.DeleteBackupDirectory(ctx, "123", "456", &CloudBackupDeleteDirectoryPayload{
		Keep: true,
	})
	require.NoError(t, err)
}

func TestCloudBackupService_ListTenantDirectories(t *testing.T) {
	setup()
	defer teardown()
	var cb cloudBackupService
	mux.HandleFunc(testlib.CloudBackupURL(cb.directoryPath()),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			resp := `[

  {

    "created_at": "2020-06-16T11:53:36.864994",

    "description": null,

    "id": "3237174f-9a9b-4304-a72e-4ebdaa15b6a2",

    "machine_id": "e5e0c123-f7f8-425d-bb31-bc514c6e48ef",

    "name": "backup config",

    "path": "etc/nginx",

    "quota": null,

    "size": 0,

    "tenant_id": "98dbba57-60c7-4c22-8087-8313997ea776",

    "updated_at": "2020-06-16T11:5"

  },

  {

    "created_at": "2020-06-16T15:14:50.875699",

    "description": null,

    "id": "e0a69ae7-fada-4426-beca-704f67e6d46c",

    "machine_id": "e5e0c123-f7f8-425d-bb31-bc514c6e48ef",

    "name": "backup web",

    "path": "etc/httpd",

    "quota": null,

    "size": 0,

    "tenant_id": "98dbba57-60c7-4c22-8087-8313997ea776",

    "updated_at": "2020-06-16T15:1"

  }

]`
			_, _ = fmt.Fprint(writer, resp)
		})
	directories, err := client.CloudBackup.ListTenantDirectories(ctx)
	require.NoError(t, err)
	assert.Equal(t, "etc/httpd", directories[1].Path)
}

func TestCloudBackupService_ActionBackupDirectory(t *testing.T) {
	setup()
	defer teardown()
	var cb cloudBackupService
	mux.HandleFunc(testlib.CloudBackupURL(strings.Join([]string{cb.itemMachinePath("123"), "directories", "456", "action"}, "/")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
		})
	err := client.CloudBackup.ActionBackupDirectory(ctx, "123", "456", &CloudBackupActionDirectoryPayload{
		Action: "deactive",
	})
	require.NoError(t, err)
}

func TestCloudBackupService_DeleteMultipleDirectories(t *testing.T) {
	setup()
	defer teardown()
	var cb cloudBackupService
	mux.HandleFunc(testlib.CloudBackupURL(strings.Join([]string{cb.itemMachinePath("123"), "directories", "action"}, "/")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodDelete, r.Method)
		})
	err := client.CloudBackup.DeleteMultipleDirectories(ctx, "123", &CloudBackupDeleteMultipleDirectoriesPayload{
		BackupDirectories: []string{"123", "456"},
		Keep:              true,
	})
	require.NoError(t, err)
}

func TestCloudBackupService_ActionMultipleDirectories(t *testing.T) {
	setup()
	defer teardown()
	var cb cloudBackupService
	mux.HandleFunc(testlib.CloudBackupURL(strings.Join([]string{cb.itemMachinePath("123"), "directories", "action"}, "/")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
		})
	err := client.CloudBackup.ActionMultipleDirectories(ctx, "123", &CloudBackupActionMultipleDirectoriesPayload{
		BackupDirectories: []string{"123", "456"},
		Action:            "deactive",
	})
	require.NoError(t, err)
}
