package gobizfly

import (
	"fmt"
	"github.com/bizflycloud/gobizfly/testlib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestListActivities(t *testing.T) {
	setup()
	defer teardown()
	var cp cloudBackupService
	mux.HandleFunc(testlib.CloudBackupURL(cp.activityPath()),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			resp := `{
  "activities": [
    {
      "action": "DELETE_BACKUP_DIRECTORY",
      "backup_directory_id": "326bf60e-7065-478a-b4b2-9d59cf8b63aa",
      "created_at": "2022-01-14T03:03:46.650294+00:00",
      "extra": "{\"keep\": true}",
      "id": "a2a17cee-8aab-4365-a0a4-86e371e1fb15",
      "machine_id": "7b2e1783-9066-4bf6-bb09-84c5d2d3165f",
      "message": "Delete backup directory /home/vinh/testbackup/testbackup4 without delete data",
      "policy_id": null,
      "reason": null,
      "recovery_point": null,
      "status": "COMPLETED",
      "tenant_id": "2c21d5f1-7237-4d01-bc13-036d47753099",
      "updated_at": "2022-01-14T03:03:46.650294+00:00",
      "user_id": "2c21d5f1-7237-4d01-bc13-036d47753099"
    },
    {
      "action": "UPDATE_POLICY_BACKUP_DIRECTORY",
      "backup_directory_id": "1a5c093a-bbbb-4e0e-9fd0-b943f09cb282",
      "created_at": "2022-01-14T02:59:30.283774+00:00",
      "extra": null,
      "id": "390be083-a9f3-4d2a-a310-e280e73172eb",
      "machine_id": "7b2e1783-9066-4bf6-bb09-84c5d2d3165f",
      "message": "Update policies in backup directory /home/vinh/cho-meo-bo-ngua",
      "policy_id": null,
      "reason": null,
      "recovery_point": null,
      "status": "COMPLETED",
      "tenant_id": "2c21d5f1-7237-4d01-bc13-036d47753099",
      "updated_at": "2022-01-14T02:59:30.283774+00:00",
      "user_id": "2c21d5f1-7237-4d01-bc13-036d47753099"
    }
  ]
}`
			_, _ = fmt.Fprint(writer, resp)
		})
	activities, err := client.CloudBackup.CloudBackupListActivities(ctx)
	require.NoError(t, err)
	require.Equal(t, 2, len(activities))
	assert.Equal(t, activities[0].BackupDirectoryId, "326bf60e-7065-478a-b4b2-9d59cf8b63aa")
}
