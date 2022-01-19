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

func TestListTenantPolicies(t *testing.T) {
	setup()
	defer teardown()
	var cb *cloudBackupService
	mux.HandleFunc(testlib.CloudBackupURL(cb.policyPath()),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			resp := `{

  "policies": [

    {

      "backup_directories": [],

      "created_at": "2021-12-30T08:39:48.636780+00:00",

      "description": null,

      "id": "97a31ace-50fb-4cdc-b6bf-b401778d97f5",

      "limit_download": null,

      "limit_upload": null,

      "name": "policy1",

      "policy_type": "AUTO",

      "retention_days": null,

      "retentions": 1,

      "schedule_pattern": "0 0 5 5 0",

      "tenant_id": "2c21d5f1-7237-4d01-bc13-036d47753099",

      "updated_at": "2021-12-30T08:39:48.636780+00:00"

    },

    {

      "backup_directories": [],

      "created_at": "2021-12-30T08:39:31.078026+00:00",

      "description": null,

      "id": "5fb109f2-95f8-441d-b4c5-6566c3870036",

      "limit_download": 456,

      "limit_upload": 123,

      "name": "policy1",

      "policy_type": "AUTO",

      "retention_days": null,

      "retentions": 1,

      "schedule_pattern": "0 0 5 5 0",

      "tenant_id": "2c21d5f1-7237-4d01-bc13-036d47753099",

      "updated_at": "2021-12-30T08:39:31.078026+00:00"

    },

    {

      "backup_directories": [],

      "created_at": "2021-12-30T08:39:16.625138+00:00",

      "description": null,

      "id": "80285da5-8a9a-4686-a440-86e0816dc1a9",

      "limit_download": 456,

      "limit_upload": 123,

      "name": "policy1",

      "policy_type": "AUTO",

      "retention_days": null,

      "retentions": 1,

      "schedule_pattern": "0 0 5 5 5",

      "tenant_id": "2c21d5f1-7237-4d01-bc13-036d47753099",

      "updated_at": "2021-12-30T08:39:16.625138+00:00"

    }

  ],

  "total": 20

}`
			fmt.Fprint(writer, resp)
		})
	policies, err := client.CloudBackup.ListTenantPolicies(ctx)
	require.NoError(t, err)
	assert.Equal(t, "policy1", policies[0].Name)
}

func TestCreatePolicy(t *testing.T) {
	setup()
	defer teardown()
	var cb *cloudBackupService
	mux.HandleFunc(testlib.CloudBackupURL(cb.policyPath()),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			resp := `{

  "created_at": "2021-12-30T08:39:48.636780+00:00",

  "description": null,

  "id": "97a31ace-50fb-4cdc-b6bf-b401778d97f5",

  "limit_download": null,

  "limit_upload": null,

  "name": "policy1",

  "policy_type": "AUTO",

  "retention_days": null,

  "retentions": 1,

  "schedule_pattern": "0 0 5 5 0",

  "tenant_id": "2c21d5f1-7237-4d01-bc13-036d47753099",

  "updated_at": "2021-12-30T08:39:48.636780+00:00"

}`
			fmt.Fprint(writer, resp)
		},
	)
	policy, err := client.CloudBackup.CreatePolicy(ctx, &CreatePolicyPayload{
		Name:            "policy1",
		SchedulePattern: "* * * *",
		StorageType:     "123",
		RetentionDays:   12,
		Description:     "123123",
	})
	require.NoError(t, err)
	assert.Equal(t, "policy1", policy.Name)
}

func TestCloudBackupService_GetBackupDirectoryPolicy(t *testing.T) {
	setup()
	defer teardown()
	var cb *cloudBackupService
	mux.HandleFunc(testlib.CloudBackupURL(cb.itemPolicyPath("123")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			resp := `{

  "created_at": "2021-12-30T08:39:48.636780+00:00",

  "description": null,

  "id": "97a31ace-50fb-4cdc-b6bf-b401778d97f5",

  "limit_download": null,

  "limit_upload": null,

  "name": "policy1",

  "policy_type": "AUTO",

  "retention_days": null,

  "retentions": 1,

  "schedule_pattern": "0 0 5 5 0",

  "tenant_id": "2c21d5f1-7237-4d01-bc13-036d47753099",

  "updated_at": "2021-12-30T08:39:48.636780+00:00"

}`
			fmt.Fprint(writer, resp)
		})
	policy, err := client.CloudBackup.GetPolicy(ctx, "123")
	require.NoError(t, err)
	assert.Equal(t, "AUTO", policy.PolicyType)
}

func TestPatchPolicy(t *testing.T) {
	setup()
	defer teardown()
	var cb *cloudBackupService
	mux.HandleFunc(testlib.CloudBackupURL(cb.itemPolicyPath("123")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPatch, r.Method)
			resp := `{

  "created_at": "2021-12-30T08:39:48.636780+00:00",

  "description": null,

  "id": "97a31ace-50fb-4cdc-b6bf-b401778d97f5",

  "limit_download": null,

  "limit_upload": null,

  "name": "policy1",

  "policy_type": "AUTO",

  "retention_days": null,

  "retentions": 1,

  "schedule_pattern": "0 0 5 5 0",

  "tenant_id": "2c21d5f1-7237-4d01-bc13-036d47753099",

  "updated_at": "2021-12-30T08:39:48.636780+00:00"

}`
			fmt.Fprint(writer, resp)
		})
	policy, err := client.CloudBackup.PatchPolicy(ctx, "123", &PatchPolicyPayload{
		Name: "123213123",
	})
	require.NoError(t, err)
	assert.Equal(t, "0 0 5 5 0", policy.SchedulePattern)
}

func TestDeletePolicy(t *testing.T) {
	setup()
	defer teardown()

	var cb cloudBackupService

	mux.HandleFunc(testlib.CloudBackupURL(cb.itemPolicyPath("123")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodDelete, r.Method)
		})
	err := client.CloudBackup.DeletePolicy(ctx, "123")
	require.NoError(t, err)
}

func TestListAppliedPolicyDirectories(t *testing.T) {
	setup()
	defer teardown()
	var cb *cloudBackupService
	mux.HandleFunc(testlib.CloudBackupURL(strings.Join([]string{cb.itemPolicyPath("123"), "directories"}, "/")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			resp := `[

  {

    "activated": true,

    "created_at": "2021-11-04T01:50:18.859043+00:00",

    "deleted": false,

    "deleted_at": null,

    "description": null,

    "id": "a247b36f-bfc6-457e-87b7-541c152588a8",

    "machine_id": "7c3f6165-408b-49ca-9d03-57494980b758",

    "name": "backup_dir_test",

    "path": "/home/vinh/cho-meo-bo-ngua",

    "quota": null,

    "size": 0,

    "tenant_id": "2c21d5f1-7237-4d01-bc13-036d47753099",

    "updated_at": "2021-11-04T01:50:18.859043+00:00"

  }

]`
			fmt.Fprint(writer, resp)
		})
	directories, err := client.CloudBackup.ListAppliedPolicyDirectories(ctx, "123")
	require.NoError(t, err)
	assert.Equal(t, false, directories[0].Deleted)
}

func TestCloudBackupService_ActionPolicyDirectory(t *testing.T) {
	setup()
	defer teardown()
	var cb *cloudBackupService
	mux.HandleFunc(testlib.CloudBackupURL(strings.Join([]string{cb.itemPolicyPath("123"), "action"}, "/")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
		})
	err := client.CloudBackup.ActionPolicyDirectory(ctx, "123", &ActionPolicyDirectoryPayload{
		Action:      "apply_directory",
		DirectoryId: "123123213123",
	})
	require.NoError(t, err)
}
