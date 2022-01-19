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

func TestListTenantRecoveryPoints(t *testing.T) {
	setup()
	defer teardown()
	var cp cloudBackupService
	mux.HandleFunc(testlib.CloudBackupURL(cp.recoveryPointPath()),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			resp := `[
{
"backup_directory": {
"activated": true,
"created_at": "2020-08-26T15:27:45.174584+07:00",
"deleted": true,
"description": null,
"id": "93f2ae07-b7ad-4fe7-8f8b-08d0831b73bd",
"machine": {
"agent_version": null,
"created_at": "2020-08-26T14:42:44.122828+07:00",
"deleted": true,
"encryption": false,
"host_name": null,
"id": "90f59741-060b-4381-8ad6-f56117f3d2ee",
"ip_address": null,
"name": null,
"os_machine_id": null,
"os_version": null,
"status": "OFFLINE",
"tenant_id": "97ef441a-3edf-4ab7-847a-7afe17964e8f",
"updated_at": "2020-09-10T16:50:31.635576+07:00"
},
"name": null,
"path": "var/as/henlib/dev/nginx",
"quota": null,
"size": 0,
"updated_at": "2020-09-10T16:50:31.594500+07:00"
},
"created_at": "2020-09-10T16:49:01.900761+07:00",
"id": "8df81afc-88ec-4960-9c40-944bb6f3ca65",
"name": "20/22/5",
"recovery_point_type": "INITIAL_REPLICA",
"status": "CREATED",
"updated_at": "2020-09-10T16:49:01.900761+07:00"
}
]`
			fmt.Fprint(writer, resp)
		})
	recoveryPoints, err := client.CloudBackup.ListTenantRecoveryPoints(ctx)
	require.NoError(t, err)
	assert.Equal(t, 1, len(recoveryPoints))
	assert.Equal(t, true, recoveryPoints[0].BackupDirectory.Activated)
}

func TestDeleteMultipleRecoveryPoint(t *testing.T) {
	setup()
	defer teardown()
	var cp cloudBackupService
	mux.HandleFunc(testlib.CloudBackupURL(cp.recoveryPointPath()),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, r.Method, http.MethodDelete)
		})
	deleteRecoveryPoint := []string{"123", "456", "789"}
	err := client.CloudBackup.DeleteMultipleRecoveryPoints(ctx, CloudBackupDeleteMultipleRecoveryPointPayload{
		RecoveryPointIds: deleteRecoveryPoint,
	})
	require.NoError(t, err)
}

func TestListDirectoryRecoveryPoints(t *testing.T) {
	setup()
	defer teardown()
	var cp cloudBackupService
	mux.HandleFunc(testlib.CloudBackupURL(cp.machineDirectoryRecoveryPointPath("123", "456")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			resp := `{
    "recovery_points": [
        {
            "backup_directory": {
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
                "name": "zui zui",
                "path": "/home/vinh/cho-meo-bo-ngua",
                "quota": null,
                "size": 0,
                "updated_at": "2022-01-14T02:59:30.198157+00:00"
            },
            "created_at": "2022-01-12T01:47:43.334811+00:00",
            "id": "521fa734-6348-4960-9cbc-c3305e0d8710",
            "index_hash": null,
            "local_size": 0,
            "method": "Policy_Type.MANUAL",
            "name": "test-file-backup",
            "progress": null,
            "recovery_point_type": "INITIAL_REPLICA",
            "status": "CREATED",
            "storage_size": 0,
            "total_files": 0,
            "updated_at": "2022-01-12T01:47:43.334811+00:00"
        },
        {
            "backup_directory": {
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
                "name": "zui zui",
                "path": "/home/vinh/cho-meo-bo-ngua",
                "quota": null,
                "size": 0,
                "updated_at": "2022-01-14T02:59:30.198157+00:00"
            },
            "created_at": "2022-01-12T01:47:29.303432+00:00",
            "id": "80f6d3ee-15c4-4f64-bcc7-482a30e01dc6",
            "index_hash": null,
            "local_size": 0,
            "method": "Policy_Type.MANUAL",
            "name": "backup-manual-01/12/2022",
            "progress": null,
            "recovery_point_type": "INITIAL_REPLICA",
            "status": "CREATED",
            "storage_size": 0,
            "total_files": 0,
            "updated_at": "2022-01-12T01:47:29.303432+00:00"
        }
],
    "total": 2
}]`
			fmt.Fprint(writer, resp)
		})
	recoveryPoints, err := client.CloudBackup.ListDirectoryRecoveryPoints(ctx, "123", "456")
	require.NoError(t, err)
	assert.Equal(t, 2, len(recoveryPoints))
	assert.Equal(t, false, recoveryPoints[0].BackupDirectory.Deleted)
}

func TestRecoveryPointAction(t *testing.T) {
	setup()
	defer teardown()
	var cb cloudBackupService
	mux.HandleFunc(testlib.CloudBackupURL(strings.Join([]string{cb.itemRecoveryPointPath("123"), "action"}, "/")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPatch, r.Method)
			resp := `{

  "backup_directory": {

    "activated": true,

    "created_at": "2020-08-26T15:27:45.174584+07:00",

    "deleted": true,

    "description": null,

    "id": "93f2ae07-b7ad-4fe7-8f8b-08d0831b73bd",

    "machine": {

      "agent_version": null,

      "created_at": "2020-08-26T14:42:44.122828+07:00",

      "deleted": true,

      "encryption": false,

      "host_name": null,

      "id": "90f59741-060b-4381-8ad6-f56117f3d2ee",

      "ip_address": null,

      "name": null,

      "os_machine_id": null,

      "os_version": null,

      "status": "OFFLINE",

      "tenant_id": "97ef441a-3edf-4ab7-847a-7afe17964e8f",

      "updated_at": "2020-09-10T16:50:31.635576+07:00"

    },

    "name": null,

    "path": "var/as/henlib/dev/nginx",

    "quota": null,

    "size": 0,

    "updated_at": "2020-09-10T16:50:31.594500+07:00"

  },

  "created_at": "2020-09-10T16:49:01.900761+07:00",

  "id": "8df81afc-88ec-4960-9c40-944bb6f3ca65",

  "name": "20/22/5",

  "recovery_point_type": "INITIAL_REPLICA",

  "status": "CREATED",

  "updated_at": "2020-09-10T16:49:01.900761+07:00"

}`
			fmt.Fprint(writer, resp)
		})
	recoveryPoint, err := client.CloudBackup.RecoveryPointAction(ctx, "123", &CloudBackupRecoveryPointActionPayload{
		Action: "convert_to_manual_recovery_point",
	})
	require.NoError(t, err)
	assert.Equal(t, "", recoveryPoint.BackupDirectory.Machine.Name)
}

func TestGetRecoveryPoint(t *testing.T) {
	setup()
	defer teardown()

	var cb cloudBackupService
	mux.HandleFunc(testlib.CloudBackupURL(cb.itemRecoveryPointPath("123")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			resp := `{

  "backup_directory": {

    "activated": true,

    "created_at": "2020-08-26T15:27:45.174584+07:00",

    "deleted": true,

    "description": null,

    "id": "93f2ae07-b7ad-4fe7-8f8b-08d0831b73bd",

    "machine": {

      "agent_version": null,

      "created_at": "2020-08-26T14:42:44.122828+07:00",

      "deleted": true,

      "encryption": false,

      "host_name": null,

      "id": "90f59741-060b-4381-8ad6-f56117f3d2ee",

      "ip_address": null,

      "name": null,

      "os_machine_id": null,

      "os_version": null,

      "status": "OFFLINE",

      "tenant_id": "97ef441a-3edf-4ab7-847a-7afe17964e8f",

      "updated_at": "2020-09-10T16:50:31.635576+07:00"

    },

    "name": null,

    "path": "var/as/henlib/dev/nginx",

    "quota": null,

    "size": 0,

    "updated_at": "2020-09-10T16:50:31.594500+07:00"

  },

  "created_at": "2020-09-10T16:49:01.900761+07:00",

  "id": "8df81afc-88ec-4960-9c40-944bb6f3ca65",

  "name": "20/22/5",

  "recovery_point_type": "INITIAL_REPLICA",

  "status": "CREATED",

  "updated_at": "2020-09-10T16:49:01.900761+07:00"

}`
			fmt.Fprint(writer, resp)
		})
	recoveryPoint, err := client.CloudBackup.GetRecoveryPoint(ctx, "123")
	require.NoError(t, err)
	assert.Equal(t, "CREATED", recoveryPoint.Status)
}

func TestListMachineRecoveryPoints(t *testing.T) {
	setup()
	defer teardown()

	var cb cloudBackupService
	mux.HandleFunc(testlib.CloudBackupURL(cb.itemMachinePath("123")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			resp := `{

  "recovery_points": [

    {

      "backup_directory": {

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

        "name": "zui zui",

        "path": "/home/vinh/cho-meo-bo-ngua",

        "quota": null,

        "size": 0,

        "updated_at": "2022-01-14T02:59:30.198157+00:00"

      },

      "created_at": "2022-01-12T01:47:43.334811+00:00",

      "id": "521fa734-6348-4960-9cbc-c3305e0d8710",

      "index_hash": null,

      "local_size": 0,

      "method": "Policy_Type.MANUAL",

      "name": "test-file-backup",

      "progress": null,

      "recovery_point_type": "INITIAL_REPLICA",

      "status": "CREATED",

      "storage_size": 0,

      "total_files": 0,

      "updated_at": "2022-01-12T01:47:43.334811+00:00"

    }

  ],

  "total": 80

}`
			fmt.Fprint(writer, resp)
		})
	recoveryPoints, err := client.CloudBackup.ListMachineRecoveryPoints(ctx, "123")
	require.NoError(t, err)
	assert.Equal(t, 0, recoveryPoints[0].StorageSize)
}

func TestDeleteRecoveryPoint(t *testing.T) {
	setup()
	defer teardown()
	var cb cloudBackupService
	mux.HandleFunc(testlib.CloudBackupURL(cb.itemRecoveryPointPath("123")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodDelete, r.Method)
		})
	err := client.CloudBackup.DeleteRecoveryPoint(ctx, "123")
	require.NoError(t, err)
}

func TestListRecoveryPointItems(t *testing.T) {
	setup()
	defer teardown()
	var cb cloudBackupService
	mux.HandleFunc(testlib.CloudBackupURL(cb.itemRecoveryPointPath("123")+"/items"),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			resp := `{

  "items": [

    {

      "access_mode": null,

      "access_time": null,

      "change_time": null,

      "content_type": null,

      "created_at": "2021-12-29T02:21:23.691510+00:00",

      "gid": null,

      "id": "f0cbac64-f75e-4455-aa7e-82c42559cf66",

      "is_dir": null,

      "item_name": "cho-meo-bo-ngua",

      "item_type": null,

      "mode": null,

      "modify_time": null,

      "real_name": null,

      "size": null,

      "status": null,

      "symlink_path": null,

      "uid": null,

      "updated_at": "2021-12-29T02:21:23.691510+00:00"

    },

    {

      "access_mode": null,

      "access_time": null,

      "change_time": null,

      "content_type": null,

      "created_at": "2021-12-29T02:21:23.691510+00:00",

      "gid": null,

      "id": "c4903d3d-fc9e-42af-acb0-bf3dd68bf279",

      "is_dir": null,

      "item_name": "cho-meo-bo-ngua/con-cho-pug.jpg",

      "item_type": null,

      "mode": null,

      "modify_time": null,

      "real_name": null,

      "size": null,

      "status": null,

      "symlink_path": null,

      "uid": null,

      "updated_at": "2021-12-29T02:21:23.691510+00:00"

    }

  ],

  "total": 2

}
`
			fmt.Fprint(writer, resp)
		})
	items, err := client.CloudBackup.ListRecoveryPointItems(ctx, "123")
	require.NoError(t, err)
	assert.Equal(t, 2, len(items))
}
