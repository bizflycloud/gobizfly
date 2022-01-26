package gobizfly

import (
	"encoding/json"
	"fmt"
	"github.com/bizflycloud/gobizfly/testlib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestBackupList(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.CloudServerURL(backupPath), func(writer http.ResponseWriter, request *http.Request) {
		assert.Equal(t, http.MethodGet, request.Method)
		resp := `[
  {
    "_id": "6016290450b0b5a3708df3dc",
    "created_at": "2021-01-31T10:50:28.332000",
    "next_run_at": "2021-02-15T02:00:00.000000",
    "billing_plan": "saving_plan",
    "options": {
      "frequency": "1440",
      "size": "2"
    },
    "resource_id": "36089786-0b73-4787-923d-cbb0a2a34377",
    "resource_type": "volume",
    "scheduled_hour": 2,
    "tenant_id": "ebbed256d9414b0598719c42dc17e837",
    "type": "volume_snapshot",
    "updated_at": "2021-01-31T10:50:28.332000"
  }
]`
		_, _ = fmt.Fprint(writer, resp)
	})
	backups, err := client.ScheduledVolumeBackup.List(ctx)
	require.NoError(t, err)
	assert.Equal(t, 1, len(backups))
	assert.Equal(t, "6016290450b0b5a3708df3dc", backups[0].ID)
	assert.Equal(t, "2021-01-31T10:50:28.332000", backups[0].CreatedAt)
}

func TestBackupCreate(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.CloudServerURL(backupPath), func(writer http.ResponseWriter, request *http.Request) {
		assert.Equal(t, http.MethodPost, request.Method)
		var payload *CreateBackupPayload
		err := json.NewDecoder(request.Body).Decode(&payload)
		require.NoError(t, err)
		assert.Equal(t, "36089786-0b73-4787-923d-cbb0a2a34377", payload.ResourceID)
		resp := `{
    "message": "Create backup successfully",
    "code": "OK",
    "trace_id": "33f1753adbc93449cc23f5d00042d041",
    "transaction_id": "7b0a88b54b58eca5",
    "request_id": "req-16d2d50f-2d22-4e97-853e-2b379007d839",
    "data": {
        "_id": "61d66d04f42f2f2e4a831fbe",
        "type": "volume_snapshot",
        "resource_type": "volume",
        "resource_id": "ef1cf35d-43a4-4da6-8d1e-70b1e67cca32",
        "options": {
            "frequency": "1440",
            "size": "3"
        },
        "tenant_id": "ebbed256d9414b0598719c42dc17e837",
        "created_at": "2022-01-06T11:16:04.223000",
        "scheduled_hour": 2,
        "updated_at": "2022-01-06T11:16:04.223000",
        "next_run_at": "2022-01-07T02:00:00.000000",
        "billing_plan": "on_demand",
        "volume_id": "ef1cf35d-43a4-4da6-8d1e-70b1e67cca32",
        "snapshots": [
            {
                "id": "ab66c6ed-8974-4400-aa5b-ac4c24fce383",
                "created_at": "2022-01-05T18:45:14.000000",
                "updated_at": "2022-01-05T18:45:15.000000",
                "name": "test7@06-01-2022",
                "description": null,
                "volume_id": "ef1cf35d-43a4-4da6-8d1e-70b1e67cca32",
                "volume_type_id": "53d70bc4-90bf-4c35-834d-b3a0794a6453",
                "status": "available",
                "size": 40,
                "metadata": {
                    "category": "premium",
                    "type": "backup",
                    "job_id": "61a841b01941a3333ad7bfea"
                },
                "os-extended-snapshot-attributes:project_id": "ebbed256d9414b0598719c42dc17e837",
                "os-extended-snapshot-attributes:progress": "100%",
                "snapshot_type": "PREMIUM_SSD",
                "type": "SSD",
                "category": "premium",
                "billing_plan": "saving_plan"
            },
            {
                "id": "3ff93879-9952-4a3e-9f13-503f804d08fc",
                "created_at": "2022-01-04T18:45:20.000000",
                "updated_at": "2022-01-04T18:45:22.000000",
                "name": "test7@05-01-2022",
                "description": null,
                "volume_id": "ef1cf35d-43a4-4da6-8d1e-70b1e67cca32",
                "volume_type_id": "53d70bc4-90bf-4c35-834d-b3a0794a6453",
                "status": "available",
                "size": 40,
                "metadata": {
                    "category": "premium",
                    "type": "backup",
                    "job_id": "61a841b01941a3333ad7bfea"
                },
                "os-extended-snapshot-attributes:project_id": "ebbed256d9414b0598719c42dc17e837",
                "os-extended-snapshot-attributes:progress": "100%",
                "snapshot_type": "PREMIUM_SSD",
                "type": "SSD",
                "category": "premium",
                "billing_plan": "saving_plan"
            },
            {
                "id": "68fb9c52-8047-4993-ab6d-fb965d2b397d",
                "created_at": "2022-01-03T18:45:13.000000",
                "updated_at": "2022-01-03T18:45:14.000000",
                "name": "test7@04-01-2022",
                "description": null,
                "volume_id": "ef1cf35d-43a4-4da6-8d1e-70b1e67cca32",
                "volume_type_id": "53d70bc4-90bf-4c35-834d-b3a0794a6453",
                "status": "available",
                "size": 40,
                "metadata": {
                    "category": "premium",
                    "type": "backup",
                    "job_id": "61a841b01941a3333ad7bfea"
                },
                "os-extended-snapshot-attributes:project_id": "ebbed256d9414b0598719c42dc17e837",
                "os-extended-snapshot-attributes:progress": "100%",
                "snapshot_type": "PREMIUM_SSD",
                "type": "SSD",
                "category": "premium",
                "billing_plan": "saving_plan"
            }
        ],
        "volume": {
            "id": "ef1cf35d-43a4-4da6-8d1e-70b1e67cca32",
            "status": "available",
            "size": 40,
            "availability_zone": "HN1",
            "created_at": "2021-12-02T03:45:41.000000",
            "updated_at": "2021-12-02T03:45:42.000000",
            "attachments": [],
            "name": "test7",
            "description": "",
            "volume_type": "PREMIUM_SSD",
            "snapshot_id": null,
            "source_volid": null,
            "metadata": {
                "category": "premium",
                "billing_plan": "on_demand"
            },
            "user_id": "7156c45b82cb4fabba997a90b032c0de",
            "bootable": false,
            "encrypted": false,
            "replication_status": null,
            "consistencygroup_id": null,
            "multiattach": false,
            "os-vol-tenant-attr:tenant_id": "ebbed256d9414b0598719c42dc17e837",
            "billing_plan": "on_demand",
            "type": "SSD",
            "category": "premium",
            "region_name": "HaNoi",
            "attached_type": "datadisk"
        }
    }
}`
		_, _ = fmt.Fprint(writer, resp)
	})
	backup, err := client.ScheduledVolumeBackup.Create(ctx, &CreateBackupPayload{
		ResourceID: "36089786-0b73-4787-923d-cbb0a2a34377",
		Frequency:  "1440",
		Hour:       2,
		Size:       "2",
	})
	require.NoError(t, err)
	assert.Equal(t, "ef1cf35d-43a4-4da6-8d1e-70b1e67cca32", backup.ResourceID)
	assert.Equal(t, "on_demand", backup.BillingPlan)
	assert.Equal(t, "1440", backup.Options.Frequency)
}

func TestBackupGet(t *testing.T) {
	setup()
	defer teardown()
	var b scheduledVolumeBackup
	mux.HandleFunc(testlib.CloudServerURL(b.itemPath("6016290450b0b5a3708df3dc")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			resp := `{
    "_id": "61a841b01941a3333ad7bfea",
    "type": "volume_snapshot",
    "resource_type": "volume",
    "resource_id": "ef1cf35d-43a4-4da6-8d1e-70b1e67cca32",
    "options": {
        "frequency": "1440",
        "size": "3"
    },
    "tenant_id": "ebbed256d9414b0598719c42dc17e837",
    "created_at": "2021-12-02T10:46:56.424000",
    "scheduled_hour": 2,
    "updated_at": "2021-12-02T10:46:56.424000",
    "next_run_at": "2022-01-07T02:00:00.000000",
    "billing_plan": "on_demand",
    "volume_id": "ef1cf35d-43a4-4da6-8d1e-70b1e67cca32",
    "snapshots": [
        {
            "id": "ab66c6ed-8974-4400-aa5b-ac4c24fce383",
            "created_at": "2022-01-05T18:45:14.000000",
            "updated_at": "2022-01-05T18:45:15.000000",
            "name": "test7@06-01-2022",
            "description": null,
            "volume_id": "ef1cf35d-43a4-4da6-8d1e-70b1e67cca32",
            "volume_type_id": "53d70bc4-90bf-4c35-834d-b3a0794a6453",
            "status": "available",
            "size": 40,
            "metadata": {
                "category": "premium",
                "type": "backup",
                "job_id": "61a841b01941a3333ad7bfea"
            },
            "os-extended-snapshot-attributes:project_id": "ebbed256d9414b0598719c42dc17e837",
            "os-extended-snapshot-attributes:progress": "100%",
            "snapshot_type": "PREMIUM_SSD",
            "type": "SSD",
            "category": "premium",
            "billing_plan": "saving_plan"
        },
        {
            "id": "3ff93879-9952-4a3e-9f13-503f804d08fc",
            "created_at": "2022-01-04T18:45:20.000000",
            "updated_at": "2022-01-04T18:45:22.000000",
            "name": "test7@05-01-2022",
            "description": null,
            "volume_id": "ef1cf35d-43a4-4da6-8d1e-70b1e67cca32",
            "volume_type_id": "53d70bc4-90bf-4c35-834d-b3a0794a6453",
            "status": "available",
            "size": 40,
            "metadata": {
                "category": "premium",
                "type": "backup",
                "job_id": "61a841b01941a3333ad7bfea"
            },
            "os-extended-snapshot-attributes:project_id": "ebbed256d9414b0598719c42dc17e837",
            "os-extended-snapshot-attributes:progress": "100%",
            "snapshot_type": "PREMIUM_SSD",
            "type": "SSD",
            "category": "premium",
            "billing_plan": "saving_plan"
        },
        {
            "id": "68fb9c52-8047-4993-ab6d-fb965d2b397d",
            "created_at": "2022-01-03T18:45:13.000000",
            "updated_at": "2022-01-03T18:45:14.000000",
            "name": "test7@04-01-2022",
            "description": null,
            "volume_id": "ef1cf35d-43a4-4da6-8d1e-70b1e67cca32",
            "volume_type_id": "53d70bc4-90bf-4c35-834d-b3a0794a6453",
            "status": "available",
            "size": 40,
            "metadata": {
                "category": "premium",
                "type": "backup",
                "job_id": "61a841b01941a3333ad7bfea"
            },
            "os-extended-snapshot-attributes:project_id": "ebbed256d9414b0598719c42dc17e837",
            "os-extended-snapshot-attributes:progress": "100%",
            "snapshot_type": "PREMIUM_SSD",
            "type": "SSD",
            "category": "premium",
            "billing_plan": "saving_plan"
        }
    ],
    "volume": {
        "id": "ef1cf35d-43a4-4da6-8d1e-70b1e67cca32",
        "status": "available",
        "size": 40,
        "availability_zone": "HN1",
        "created_at": "2021-12-02T03:45:41.000000",
        "updated_at": "2021-12-02T03:45:42.000000",
        "attachments": [],
        "name": "test7",
        "description": "",
        "volume_type": "PREMIUM_SSD",
        "snapshot_id": null,
        "source_volid": null,
        "metadata": {
            "category": "premium",
            "billing_plan": "on_demand"
        },
        "user_id": "7156c45b82cb4fabba997a90b032c0de",
        "bootable": false,
        "encrypted": false,
        "replication_status": null,
        "consistencygroup_id": null,
        "multiattach": false,
        "os-vol-tenant-attr:tenant_id": "ebbed256d9414b0598719c42dc17e837",
        "billing_plan": "on_demand",
        "type": "SSD",
        "category": "premium",
        "region_name": "HaNoi",
        "attached_type": "datadisk"
    }
}`
			_, _ = fmt.Fprint(writer, resp)
		})
	backup, err := client.ScheduledVolumeBackup.Get(ctx, "6016290450b0b5a3708df3dc")
	require.NoError(t, err)
	assert.Equal(t, "on_demand", backup.BillingPlan)
	assert.Equal(t, "1440", backup.Options.Frequency)
	assert.Equal(t, 2, backup.ScheduledHour)
}

func TestBackupDelete(t *testing.T) {
	setup()
	defer teardown()
	var b scheduledVolumeBackup
	mux.HandleFunc(testlib.CloudServerURL(b.itemPath("6016290450b0b5a3708df3dc")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodDelete, r.Method)
		})
	require.NoError(t, client.ScheduledVolumeBackup.Delete(ctx, "6016290450b0b5a3708df3dc"))
}

func TestBackupUpdate(t *testing.T) {
	setup()
	defer teardown()
	var b scheduledVolumeBackup
	mux.HandleFunc(testlib.CloudServerURL(b.itemPath("6016290450b0b5a3708df3dc")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPut, r.Method)
			resp := `{
  "_id": "6016290450b0b5a3708df3dc",
  "created_at": "2021-01-31T10:50:28.332000",
  "next_run_at": "2021-02-15T02:00:00.000000",
  "options": {
    "frequency": "2880",
    "size": "2"
  },
  "resource_id": "36089786-0b73-4787-923d-cbb0a2a34377",
  "resource_type": "volume",
  "scheduled_hour": 2,
  "billing_plan": "saving_plan",
  "snapshots": [
    {
      "category": "premium",
      "created_at": "2021-02-13T18:48:12.000000",
      "description": null,
      "id": "5159247d-7196-4c0a-8568-ab096cae4f35",
      "metadata": {
        "category": "premium",
        "job_id": "6016290450b0b5a3708df3dc",
        "type": "backup"
      },
      "name": "asfdasdf234sadf_rootdisk@13-02-2021",
      "os-extended-snapshot-attributes:progress": "100%",
      "os-extended-snapshot-attributes:project_id": "ebbed256d9414b0598719c42dc17e837",
      "size": 40,
      "snapshot_type": "PREMIUM_SSD",
      "status": "available",
      "type": "SSD",
      "updated_at": "2021-02-13T18:48:13.000000",
      "volume_id": "36089786-0b73-4787-923d-cbb0a2a34377",
      "volume_type_id": "53d70bc4-90bf-4c35-834d-b3a0794a6453"
    },
    {
      "category": "premium",
      "created_at": "2021-02-12T18:48:03.000000",
      "description": null,
      "id": "9a79e708-b2ea-4232-8dbc-d514de29f852",
      "metadata": {
        "category": "premium",
        "job_id": "6016290450b0b5a3708df3dc",
        "type": "backup"
      },
      "name": "asfdasdf234sadf_rootdisk@12-02-2021",
      "os-extended-snapshot-attributes:progress": "100%",
      "os-extended-snapshot-attributes:project_id": "ebbed256d9414b0598719c42dc17e837",
      "size": 40,
      "snapshot_type": "PREMIUM_SSD",
      "status": "available",
      "type": "SSD",
      "updated_at": "2021-02-12T18:48:04.000000",
      "volume_id": "36089786-0b73-4787-923d-cbb0a2a34377",
      "volume_type_id": "53d70bc4-90bf-4c35-834d-b3a0794a6453"
    }
  ],
  "tenant_id": "ebbed256d9414b0598719c42dc17e837",
  "type": "volume_snapshot",
  "updated_at": "2021-01-31T10:50:28.332000",
  "volume": {
    "attached_type": "datadisk",
    "attachments": [],
    "availability_zone": "HN1",
    "bootable": true,
    "category": "premium",
    "consistencygroup_id": null,
    "created_at": "2021-01-25T04:11:21.000000",
    "description": null,
    "encrypted": false,
    "id": "36089786-0b73-4787-923d-cbb0a2a34377",
    "metadata": {
      "category": "premium"
    },
    "multiattach": false,
    "name": "asfdasdf234sadf_rootdisk",
    "os-vol-tenant-attr:tenant_id": "ebbed256d9414b0598719c42dc17e837",
    "region_name": "HaNoi",
    "replication_status": null,
    "size": 40,
    "snapshot_id": null,
    "source_volid": null,
    "status": "available",
    "type": "SSD",
    "updated_at": "2021-01-25T04:12:31.000000",
    "user_id": "7156c45b82cb4fabba997a90b032c0de",
    "volume_image_metadata": {
      "checksum": "07d74cc43fd0a3b4531673f70a3b686f",
      "container_format": "bare",
      "disk_format": "raw",
      "image_id": "9a0f31e3-c43d-4fc2-ae1c-cc6ebde571fa",
      "image_name": "CentOS-7.0[64-bit-version]",
      "min_disk": "0",
      "min_ram": "0",
      "signature_verified": "False",
      "size": "4194304000"
    },
    "volume_type": "PREMIUM_SSD"
  },
  "volume_id": "36089786-0b73-4787-923d-cbb0a2a34377"
}`
			_, _ = fmt.Fprint(writer, resp)
		})
	backup, err := client.ScheduledVolumeBackup.Update(ctx, "6016290450b0b5a3708df3dc", &UpdateBackupPayload{
		Frequency: "2880",
	})
	require.NoError(t, err)
	assert.Equal(t, "2880", backup.Options.Frequency)
}
