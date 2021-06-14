// This file is part of gobizfly

package gobizfly

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/bizflycloud/gobizfly/testlib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSnapshotList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(testlib.CloudServerURL(snapshotPath), func(writer http.ResponseWriter, request *http.Request) {
		require.Equal(t, http.MethodGet, request.Method)
		resp := `
[
   {
      "status":"available",
      "volume_type_id":"ec6fb900-1ae0-4e9e-90e0-53a6063f95e1",
      "description":null,
      "updated_at":"2020-05-06T16:45:14.000000",
      "volume_id":"c4e6bf65-32d8-4ef3-bbd3-3cc9676f8246",
      "id":"586fd3ae-597c-4acc-aab0-713e80245b28",
      "size":20,
      "os-extended-snapshot-attributes:progress":"100%",
      "name":"timtro_rootdisk@06-05-2020",
      "os-extended-snapshot-attributes:project_id":"1e7f10a9850b45b488a3f0417ccb60e0",
      "created_at":"2020-05-06T16:45:12.000000",
      "metadata":{
         "category":"premium",
         "type":"backup",
         "job_id":"5cb143884096dd75cf0ebcb1"
      },
      "is_using_autoscale":false,
      "volume":{
         "attachments":[

         ],
         "availability_zone":"HN1",
         "encrypted":false,
         "updated_at":"2020-02-04T05:29:37.000000",
         "replication_status":null,
         "snapshot_id":null,
         "id":"c4e6bf65-32d8-4ef3-bbd3-3cc9676f8246",
         "size":20,
         "user_id":"fe17f866b86e4646af83313d967ac7db",
         "os-vol-tenant-attr:tenant_id":"1e7f10a9850b45b488a3f0417ccb60e0",
         "metadata":{
            "category":"premium",
            "readonly":"False",
            "expired_time":"31/12/2019",
            "policy":"free",
            "os_type":"Ubuntu 14.04",
            "attached_mode":"rw"
         },
         "status":"in-use",
         "volume_image_metadata":{
            "checksum":"7f0470e228554a0e16bfc0438ac50521",
            "min_ram":"0",
            "disk_format":"raw",
            "image_name":"Ubuntu-14.04-Server-amd64-LTS",
            "image_id":"83aeff78-14ad-498d-976b-ed15bc5fa5ac",
            "container_format":"bare",
            "min_disk":"0",
            "size":"2097152000"
         },
         "description":null,
         "multiattach":false,
         "source_volid":null,
         "consistencygroup_id":null,
         "name":"timtro_rootdisk",
         "bootable":true,
         "created_at":"2018-05-04T01:51:38.000000",
         "volume_type":"HDD",
         "attached_type":"rootdisk",
         "type":"HDD",
         "category":"premium"
      }
   },
   {
      "status":"available",
      "volume_type_id":"53d70bc4-90bf-4c35-834d-b3a0794a6453",
      "description":null,
      "updated_at":"2020-05-05T02:33:44.000000",
      "volume_id":"f51d3c52-9c95-4b68-8cfd-467e33bcf0b8",
      "id":"99adfcb1-c46d-405d-b982-85f97c764088",
      "size":20,
      "os-extended-snapshot-attributes:progress":"100%",
      "name":"snapshot-9-33-5",
      "os-extended-snapshot-attributes:project_id":"1e7f10a9850b45b488a3f0417ccb60e0",
      "created_at":"2020-05-05T02:33:42.000000",
      "metadata":{
         "category":"premium"
      },
      "is_using_autoscale":false,
      "volume":{
         "attachments":[

         ],
         "availability_zone":"HN1",
         "encrypted":false,
         "updated_at":"2020-04-27T07:00:15.000000",
         "replication_status":null,
         "snapshot_id":null,
         "id":"f51d3c52-9c95-4b68-8cfd-467e33bcf0b8",
         "size":20,
         "user_id":"eaf6c51f66614c59b1bcd0fd5951bbe7",
         "os-vol-tenant-attr:tenant_id":"1e7f10a9850b45b488a3f0417ccb60e0",
         "metadata":{
            "category":"premium",
            "attached_mode":"rw"
         },
         "status":"in-use",
         "volume_image_metadata":{
            "image_location":"snapshot",
            "image_state":"available",
            "container_format":"bare",
            "min_ram":"0",
            "image_name":"Ubuntu 18.04",
            "boot_roles":"admin",
            "image_id":"cbf5f34b-751b-42a5-830f-6b2324f61d5a",
            "min_disk":"5",
            "base_image_ref":"e410a263-b265-492d-9bb1-cd8e75fc3e92",
            "size":"5368709120",
            "instance_uuid":"b1405592-3cc0-4361-a597-e631ff7ba6b2",
            "user_id":"5676103832f14c129306bf525ec7b2de",
            "image_type":"image",
            "checksum":"a8181813ef91e1c2e1c506fc1643d8a4",
            "disk_format":"raw"
         },
         "description":null,
         "multiattach":false,
         "source_volid":null,
         "consistencygroup_id":null,
         "name":"prtgadmin2_rootdisk",
         "bootable":true,
         "created_at":"2020-04-27T06:59:57.000000",
         "volume_type":"SSD",
         "attached_type":"rootdisk",
         "type":"HDD",
         "category":"premium"
      }
   }
]
`
		_, _ = fmt.Fprint(writer, resp)
	})
	snapshots, err := client.Snapshot.List(ctx, &ListSnasphotsOptions{})
	require.NoError(t, err)
	assert.Len(t, snapshots, 2)
	volume := snapshots[0]
	assert.Equal(t, "586fd3ae-597c-4acc-aab0-713e80245b28", volume.Id)
}

func TestSnapshotGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(strings.Join([]string{testlib.CloudServerURL(snapshotPath), "d5d79b3f-d0cd-4535-b0d3-27d8ec2d62f5"}, "/"),
		func(writer http.ResponseWriter, request *http.Request) {
			require.Equal(t, http.MethodGet, request.Method)
			resp := `
{
   "status":"available",
   "volume_type_id":"b74a5c89-7293-4069-b87b-cf6867fc574a",
   "description":null,
   "updated_at":"2020-05-07T02:22:24.000000",
   "volume_id":"9d9b4c3f-f442-484f-844f-7418cab34e33",
   "id":"d5d79b3f-d0cd-4535-b0d3-27d8ec2d62f5",
   "size":20,
   "os-extended-snapshot-attributes:progress":"100%",
   "name":"ducpx-snapshot-test",
   "os-extended-snapshot-attributes:project_id":"1e7f10a9850b45b488a3f0417ccb60e0",
   "created_at":"2020-05-07T02:22:23.000000",
   "metadata":{
      "category":"basic"
   },
   "is_using_autoscale":false,
   "volume":{
      "attachments":[
         {
            "server_id":"903c3316-6c5b-42d1-adee-8da361eed707",
            "attachment_id":"cf98712b-61d7-47ed-987d-0a56f403d793",
            "attached_at":"2020-05-06T08:11:25.000000",
            "host_name":"thor-compute-023",
            "volume_id":"9d9b4c3f-f442-484f-844f-7418cab34e33",
            "device":"/dev/vda",
            "id":"9d9b4c3f-f442-484f-844f-7418cab34e33"
         }
      ],
      "links":[
         {
            "href":"https://dont.find.me.vccloud.vn:8776/v3/1e7f10a9850b45b488a3f0417ccb60e0/volumes/9d9b4c3f-f442-484f-844f-7418cab34e33",
            "rel":"self"
         },
         {
            "href":"https://dont.find.me.vccloud.vn:8776/1e7f10a9850b45b488a3f0417ccb60e0/volumes/9d9b4c3f-f442-484f-844f-7418cab34e33",
            "rel":"bookmark"
         }
      ],
      "availability_zone":"HN1",
      "encrypted":false,
      "updated_at":"2020-05-06T08:11:31.000000",
      "replication_status":null,
      "snapshot_id":null,
      "id":"9d9b4c3f-f442-484f-844f-7418cab34e33",
      "size":20,
      "user_id":"eaf6c51f66614c59b1bcd0fd5951bbe7",
      "os-vol-tenant-attr:tenant_id":"1e7f10a9850b45b488a3f0417ccb60e0",
      "metadata":{
         "category":"basic",
         "attached_mode":"rw"
      },
      "status":"in-use",
      "volume_image_metadata":{
         "image_location":"snapshot",
         "image_state":"available",
         "container_format":"bare",
         "min_ram":"0",
         "image_name":"CentOS 7.7@2019-11-25-4e809d99",
         "boot_roles":"admin",
         "image_id":"662e8b68-7888-4b07-94c8-351e1c553992",
         "owner_user_name":"duylkops",
         "min_disk":"5",
         "base_image_ref":"9a0f31e3-c43d-4fc2-ae1c-cc6ebde571fa",
         "size":"5368709120",
         "instance_uuid":"c9e7ad24-9f01-41d7-bee3-177952582e1e",
         "user_id":"5676103832f14c129306bf525ec7b2de",
         "image_type":"image",
         "checksum":"4fe6f312c1578dacbcd8bc12523174a8",
         "disk_format":"raw",
         "owner_project_name":"Packer-Images",
         "owner_id":"159c53f12fc24afc88c945e9bc6cc57d"
      },
      "description":null,
      "multiattach":false,
      "source_volid":null,
      "consistencygroup_id":null,
      "name":"ducpx-devgolang_rootdisk",
      "bootable":true,
      "created_at":"2020-05-06T08:10:29.000000",
      "volume_type":"BASIC_HDD",
      "attached_type":"rootdisk",
      "type":"HDD",
      "category":"premium"
   }
}
`
			_, _ = fmt.Fprint(writer, resp)
		})
	snapshot, err := client.Snapshot.Get(ctx, "d5d79b3f-d0cd-4535-b0d3-27d8ec2d62f5")
	require.NoError(t, err)
	require.Equal(t, "d5d79b3f-d0cd-4535-b0d3-27d8ec2d62f5", snapshot.Id, "check snapshot id")
	require.Equal(t, "ducpx-snapshot-test", snapshot.Name, "check snapshot name")
}

func TestSnapshotCreate(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.CloudServerURL(snapshotPath), func(writer http.ResponseWriter, request *http.Request) {
		require.Equal(t, http.MethodPost, request.Method)
		var snapshot *SnapshotCreateRequest
		require.NoError(t, json.NewDecoder(request.Body).Decode(&snapshot))
		assert.Equal(t, "ducpx-test-create-snapshot", snapshot.Name)
		assert.Equal(t, true, snapshot.Force)
		assert.Equal(t, "c4e6bf65-32d8-4ef3-bbd3-3cc9676f8246", snapshot.VolumeId)

		resp := `
{
   "status":"available",
   "volume_type_id":"ec6fb900-1ae0-4e9e-90e0-53a6063f95e1",
   "description":null,
   "updated_at":"2020-05-06T16:45:14.000000",
   "volume_id":"c4e6bf65-32d8-4ef3-bbd3-3cc9676f8246",
   "id":"586fd3ae-597c-4acc-aab0-713e80245b28",
   "size":20,
   "os-extended-snapshot-attributes:progress":"100%",
   "name":"ducpx-test-create-snapshot",
   "os-extended-snapshot-attributes:project_id":"1e7f10a9850b45b488a3f0417ccb60e0",
   "created_at":"2020-05-06T16:45:12.000000",
   "metadata":{
      "category":"premium",
      "type":"backup",
      "job_id":"5cb143884096dd75cf0ebcb1"
   },
   "is_using_autoscale":false,
   "volume":{
      "attachments":[

      ],
      "availability_zone":"HN1",
      "encrypted":false,
      "updated_at":"2020-02-04T05:29:37.000000",
      "replication_status":null,
      "snapshot_id":null,
      "id":"c4e6bf65-32d8-4ef3-bbd3-3cc9676f8246",
      "size":20,
      "user_id":"fe17f866b86e4646af83313d967ac7db",
      "os-vol-tenant-attr:tenant_id":"1e7f10a9850b45b488a3f0417ccb60e0",
      "metadata":{
         "category":"premium",
         "readonly":"False",
         "expired_time":"31/12/2019",
         "policy":"free",
         "os_type":"Ubuntu 14.04",
         "attached_mode":"rw"
      },
      "status":"in-use",
      "volume_image_metadata":{
         "checksum":"7f0470e228554a0e16bfc0438ac50521",
         "min_ram":"0",
         "disk_format":"raw",
         "image_name":"Ubuntu-14.04-Server-amd64-LTS",
         "image_id":"83aeff78-14ad-498d-976b-ed15bc5fa5ac",
         "container_format":"bare",
         "min_disk":"0",
         "size":"2097152000"
      },
      "description":null,
      "multiattach":false,
      "source_volid":null,
      "consistencygroup_id":null,
      "name":"timtro_rootdisk",
      "bootable":true,
      "created_at":"2018-05-04T01:51:38.000000",
      "volume_type":"HDD",
      "attached_type":"rootdisk",
      "type":"HDD",
      "category":"premium"
   }
}
`

		_, _ = fmt.Fprint(writer, resp)
	})

	snapshot, err := client.Snapshot.Create(ctx, &SnapshotCreateRequest{
		Name:     "ducpx-test-create-snapshot",
		VolumeId: "c4e6bf65-32d8-4ef3-bbd3-3cc9676f8246",
		Force:    true,
	})

	require.NoError(t, err)
	assert.Equal(t, "586fd3ae-597c-4acc-aab0-713e80245b28", snapshot.Id)
	assert.Equal(t, "ducpx-test-create-snapshot", snapshot.Name)
	assert.Equal(t, 20, snapshot.Size)
	assert.Equal(t, "c4e6bf65-32d8-4ef3-bbd3-3cc9676f8246", snapshot.VolumeId)
}

func TestSnapshotDelete(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(strings.Join([]string{testlib.CloudServerURL(snapshotPath), "d5d79b3f-d0cd-4535-b0d3-27d8ec2d62f5"}, "/"),
		func(writer http.ResponseWriter, request *http.Request) {
			require.Equal(t, http.MethodDelete, request.Method)
			writer.WriteHeader(http.StatusNoContent)
		})
	require.NoError(t, client.Snapshot.Delete(ctx, "d5d79b3f-d0cd-4535-b0d3-27d8ec2d62f5"))
}
