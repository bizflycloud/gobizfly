// This file is part of gobizfly

package gobizfly

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"

	"github.com/bizflycloud/gobizfly/testlib"
)

func TestVolumeList(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.CloudServerURL(volumeBasePath), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
[
	{
		"status": "in-use",
		"user_id": "55d38aecb1034c06b99c1c87fb6f0740",
		"attachments": [
			{
				"server_id": "2b2628b1-0d11-4fd7-8d63-4ec24ff493ea",
				"attachment_id": "9c2ba971-dd59-4527-a441-26dcd3174d68",
				"attached_at": "2020-04-11T17:51:47.000000",
				"volume_id": "7b099bbb-21e9-48f9-8cec-4076d78fa201",
				"device": "/dev/vdb",
				"id": "7b099bbb-21e9-48f9-8cec-4076d78fa201",
				"server": {
					"OS-EXT-STS:task_state": null,
					"addresses": {
						"priv_sapd@vccloud.vn": [
							{
								"OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:98:c0:3f",
								"version": 4,
								"addr": "10.20.165.43",
								"OS-EXT-IPS:type": "fixed"
							}
						],
						"EXT_DIRECTNET_8": [
							{
								"OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:17:87:67",
								"version": 4,
								"addr": "103.107.182.201",
								"OS-EXT-IPS:type": "fixed"
							}
						]
					},
					"OS-EXT-STS:vm_state": "active",
					"OS-SRV-USG:launched_at": "2020-04-11T15:42:53.000000",
					"flavor": {
						"name": "nix.2c_2g",
						"ram": 2048,
						"OS-FLV-DISABLED:disabled": false,
						"vcpus": 2,
						"swap": "",
						"os-flavor-access:is_public": true,
						"rxtx_factor": 1.0,
						"OS-FLV-EXT-DATA:ephemeral": 0,
						"disk": 0,
						"id": "be7dab73-2c87-4d59-a2fd-49e4f7845310"
					},
					"id": "2b2628b1-0d11-4fd7-8d63-4ec24ff493ea",
					"security_groups": [
						{
							"name": "default"
						},
						{
							"name": "default"
						}
					],
					"user_id": "55d38aecb1034c06b99c1c87fb6f0740",
					"OS-DCF:diskConfig": "MANUAL",
					"accessIPv4": "",
					"accessIPv6": "",
					"progress": 0,
					"OS-EXT-STS:power_state": 1,
					"OS-EXT-AZ:availability_zone": "HN1",
					"config_drive": "",
					"status": "ACTIVE",
					"updated": "2020-04-11T15:42:54Z",
					"hostId": "023ca0b3e91f594819943b6e70ff0d0436734e19df2a2147a8eb1333",
					"OS-SRV-USG:terminated_at": null,
					"key_name": "sapd1",
					"name": "ceph-15",
					"created": "2020-04-11T15:42:35Z",
					"tenant_id": "17a1c3c952c84b3e84a82ddd48364938",
					"os-extended-volumes:volumes_attached": [
						{
							"id": "018e0772-1930-4329-a08c-b07422aa9fc1"
						},
						{
							"id": "7b099bbb-21e9-48f9-8cec-4076d78fa201"
						},
						{
							"id": "af2516f4-002d-41bc-932d-4ebdc5d49107"
						}
					],
					"metadata": {
						"category": "premium",
						"os_type": "CentOS 7.7"
					},
					"ipv6": false
				}
			}
		],
		"availability_zone": "HN1",
		"bootable": false,
		"encrypted": false,
		"created_at": "2020-04-11T17:51:41.000000",
		"description": null,
		"os-vol-tenant-attr:tenant_id": "17a1c3c952c84b3e84a82ddd48364938",
		"updated_at": "2020-04-11T17:51:47.000000",
		"type": "HDD",
		"name": "sapd-vol-1",
		"replication_status": null,
		"consistencygroup_id": null,
		"source_volid": null,
		"snapshot_id": null,
		"multiattach": false,
		"metadata": {
			"category": "premium",
			"attached_mode": "rw"
		},
		"id": "7b099bbb-21e9-48f9-8cec-4076d78fa201",
		"size": 20,
		"attached_type": "datadisk",
		"category": "premium"
	}
]
`
		_, _ = fmt.Fprint(w, resp)
	})

	volumes, err := client.Volume.List(ctx, &ListOptions{})
	require.NoError(t, err)
	assert.Len(t, volumes, 1)
	volume := volumes[0]
	assert.Equal(t, "7b099bbb-21e9-48f9-8cec-4076d78fa201", volume.ID)
}

func TestVolumeGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(testlib.CloudServerURL(volumeBasePath+"/7b099bbb-21e9-48f9-8cec-4076d78fa201"), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
	"status": "in-use",
	"user_id": "55d38aecb1034c06b99c1c87fb6f0740",
	"attachments": [
		{
			"server_id": "2b2628b1-0d11-4fd7-8d63-4ec24ff493ea",
			"attachment_id": "9c2ba971-dd59-4527-a441-26dcd3174d68",
			"attached_at": "2020-04-11T17:51:47.000000",
			"volume_id": "7b099bbb-21e9-48f9-8cec-4076d78fa201",
			"device": "/dev/vdb",
			"id": "7b099bbb-21e9-48f9-8cec-4076d78fa201",
			"server": {
				"OS-EXT-STS:task_state": null,
				"addresses": {
					"priv_sapd@vccloud.vn": [
						{
							"OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:98:c0:3f",
							"version": 4,
							"addr": "10.20.165.43",
							"OS-EXT-IPS:type": "fixed"
						}
					],
					"EXT_DIRECTNET_8": [
						{
							"OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:17:87:67",
							"version": 4,
							"addr": "103.107.182.201",
							"OS-EXT-IPS:type": "fixed"
						}
					]
				},
				"OS-EXT-STS:vm_state": "active",
				"OS-SRV-USG:launched_at": "2020-04-11T15:42:53.000000",
				"flavor": {
					"name": "nix.2c_2g",
					"ram": 2048,
					"OS-FLV-DISABLED:disabled": false,
					"vcpus": 2,
					"swap": "",
					"os-flavor-access:is_public": true,
					"rxtx_factor": 1.0,
					"OS-FLV-EXT-DATA:ephemeral": 0,
					"disk": 0,
					"id": "be7dab73-2c87-4d59-a2fd-49e4f7845310"
				},
				"id": "2b2628b1-0d11-4fd7-8d63-4ec24ff493ea",
				"security_groups": [
					{
						"name": "default"
					},
					{
						"name": "default"
					}
				],
				"user_id": "55d38aecb1034c06b99c1c87fb6f0740",
				"OS-DCF:diskConfig": "MANUAL",
				"accessIPv4": "",
				"accessIPv6": "",
				"progress": 0,
				"OS-EXT-STS:power_state": 1,
				"OS-EXT-AZ:availability_zone": "HN1",
				"config_drive": "",
				"status": "ACTIVE",
				"updated": "2020-04-11T15:42:54Z",
				"hostId": "023ca0b3e91f594819943b6e70ff0d0436734e19df2a2147a8eb1333",
				"OS-SRV-USG:terminated_at": null,
				"key_name": "sapd1",
				"name": "ceph-15",
				"created": "2020-04-11T15:42:35Z",
				"tenant_id": "17a1c3c952c84b3e84a82ddd48364938",
				"os-extended-volumes:volumes_attached": [
					{
						"id": "018e0772-1930-4329-a08c-b07422aa9fc1"
					},
					{
						"id": "7b099bbb-21e9-48f9-8cec-4076d78fa201"
					},
					{
						"id": "af2516f4-002d-41bc-932d-4ebdc5d49107"
					}
				],
				"metadata": {
					"category": "premium",
					"os_type": "CentOS 7.7"
				},
				"ipv6": false
			}
		}
	],
	"availability_zone": "HN1",
	"bootable": false,
	"encrypted": false,
	"created_at": "2020-04-11T17:51:41.000000",
	"description": null,
	"os-vol-tenant-attr:tenant_id": "17a1c3c952c84b3e84a82ddd48364938",
	"updated_at": "2020-04-11T17:51:47.000000",
	"type": "HDD",
	"name": "sapd-vol-1",
	"replication_status": null,
	"consistencygroup_id": null,
	"source_volid": null,
	"snapshot_id": null,
	"multiattach": false,
	"metadata": {
		"category": "premium",
		"attached_mode": "rw"
	},
	"id": "7b099bbb-21e9-48f9-8cec-4076d78fa201",
	"size": 20,
	"attached_type": "datadisk",
	"category": "premium"
}		
`
		_, _ = fmt.Fprint(w, resp)
	})

	volume, err := client.Volume.Get(ctx, "7b099bbb-21e9-48f9-8cec-4076d78fa201")
	require.NoError(t, err)
	assert.Equal(t, "7b099bbb-21e9-48f9-8cec-4076d78fa201", volume.ID)
	assert.Equal(t, "sapd-vol-1", volume.Name)
	assert.Equal(t, "HDD", volume.VolumeType)
}

func TestVolumeDelete(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.CloudServerURL(volumeBasePath+"/7b099bbb-21e9-48f9-8cec-4076d78fa201"), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
		w.WriteHeader(http.StatusNoContent)
	})
	require.NoError(t, client.Volume.Delete(ctx, "7b099bbb-21e9-48f9-8cec-4076d78fa201"))
}

func TestVolumeCreate(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.CloudServerURL(volumeBasePath), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		var volume *VolumeCreateRequest
		require.NoError(t, json.NewDecoder(r.Body).Decode(&volume))
		assert.Equal(t, "sapd-test-goclient", volume.Name)
		assert.Equal(t, 20, volume.Size)
		assert.Equal(t, "SSD", volume.VolumeType)
		assert.Equal(t, "premium", volume.VolumeCategory)

		resp := `
{
	"status": "available",
	"user_id": "55d38aecb1034c06b99c1c87fb6f0740",
	"attachments": [],
	"availability_zone": "HN1",
	"bootable": false,
	"encrypted": false,
	"created_at": "2020-04-12T08:45:45.000000",
	"description": null,
	"os-vol-tenant-attr:tenant_id": "17a1c3c952c84b3e84a82ddd48364938",
	"updated_at": "2020-04-12T08:45:46.000000",
	"volume_type": "SSD",
	"name": "sapd-test-goclient",
	"replication_status": null,
	"consistencygroup_id": null,
	"source_volid": null,
	"snapshot_id": null,
	"multiattach": false,
	"metadata": {
		"category": "premium"
	},
	"id": "af2516f4-002d-41bc-932d-4ebdc5d49107",
	"size": 20
}`

		_, _ = fmt.Fprint(w, resp)
	})

	volume, err := client.Volume.Create(ctx, &VolumeCreateRequest{
		Name:             "sapd-test-goclient",
		Size:             20,
		VolumeType:       "SSD",
		AvailabilityZone: "HN1",
		VolumeCategory:   "premium",
	})

	require.NoError(t, err)
	assert.Equal(t, "af2516f4-002d-41bc-932d-4ebdc5d49107", volume.ID)
	assert.Equal(t, 20, volume.Size)
	assert.Equal(t, "", volume.SnapshotID)
	assert.Equal(t, "HN1", volume.AvailabilityZone)
}

func TestVolumeExtend(t *testing.T) {
	setup()
	defer teardown()

	var v volume
	mux.HandleFunc(testlib.CloudServerURL(v.itemActionPath("4cb94590-c4a2-4a37-90d6-30064f68d19e")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		resp := `
{
	"task_id": "3c414504-8fba-4b70-bdcb-a5b44a2ae406"
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	resp, err := client.Volume.ExtendVolume(ctx, "4cb94590-c4a2-4a37-90d6-30064f68d19e", 30)
	require.NoError(t, err)
	assert.Equal(t, "3c414504-8fba-4b70-bdcb-a5b44a2ae406", resp.TaskID)
}

func TestVolumeRestore(t *testing.T) {
	setup()
	defer teardown()

	var v volume
	mux.HandleFunc(testlib.CloudServerURL(v.itemActionPath("4cb94590-c4a2-4a37-90d6-30064f68d19e")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		resp := `
{
	"task_id": "3c414504-8fba-4b70-bdcb-a5b44a2ae406"
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	resp, err := client.Volume.Restore(ctx, "4cb94590-c4a2-4a37-90d6-30064f68d19e", "38466e4e-7cca-11ea-a78b-9794b3babf27")
	require.NoError(t, err)
	assert.Equal(t, "3c414504-8fba-4b70-bdcb-a5b44a2ae406", resp.TaskID)
}

func TestVolumeAttach(t *testing.T) {
	setup()
	defer teardown()

	var v volume
	mux.HandleFunc(testlib.CloudServerURL(v.itemActionPath("894f0e66-4571-4fea-9766-5fc615aec4a5")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		resp := `
{
	"message": "Attach successfully",
	"volume_detail": {
		"status": "available",
		"user_id": "55d38aecb1034c06b99c1c87fb6f0740",
		"attachments": [],
		"availability_zone": "HN1",
		"bootable": false,
		"encrypted": false,
		"created_at": "2020-04-12T14:38:42.000000",
		"description": null,
		"os-vol-tenant-attr:tenant_id": "17a1c3c952c84b3e84a82ddd48364938",
		"updated_at": "2020-04-12T14:39:08.000000",
		"volume_type": "HDD",
		"name": "sapd-test-2",
		"replication_status": null,
		"consistencygroup_id": null,
		"source_volid": null,
		"snapshot_id": null,
		"multiattach": false,
		"metadata": {
			"category": "premium"
		},
		"id": "894f0e66-4571-4fea-9766-5fc615aec4a5",
		"size": 20
	}
}`
		_, _ = fmt.Fprint(w, resp)
	})

	resp, err := client.Volume.Attach(ctx, "894f0e66-4571-4fea-9766-5fc615aec4a5", "abc3f3cc-7ccb-11ea-945f-7be40572932e")
	require.NoError(t, err)
	assert.Equal(t, "894f0e66-4571-4fea-9766-5fc615aec4a5", resp.VolumeDetail.ID)
	assert.Equal(t, "Attach successfully", resp.Message)
}

func TestVolumeDetach(t *testing.T) {
	setup()
	defer teardown()

	var v volume
	mux.HandleFunc(testlib.CloudServerURL(v.itemActionPath("894f0e66-4571-4fea-9766-5fc615aec4a5")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		resp := `
{
	"message": "Detach successfully",
	"volume_detail": {
		"status": "available",
		"user_id": "55d38aecb1034c06b99c1c87fb6f0740",
		"attachments": [],
		"availability_zone": "HN1",
		"bootable": false,
		"encrypted": false,
		"created_at": "2020-04-12T14:38:42.000000",
		"description": null,
		"os-vol-tenant-attr:tenant_id": "17a1c3c952c84b3e84a82ddd48364938",
		"updated_at": "2020-04-12T14:39:08.000000",
		"volume_type": "HDD",
		"name": "sapd-test-2",
		"replication_status": null,
		"consistencygroup_id": null,
		"source_volid": null,
		"snapshot_id": null,
		"multiattach": false,
		"metadata": {
			"category": "premium"
		},
		"id": "894f0e66-4571-4fea-9766-5fc615aec4a5",
		"size": 20
	}
}`
		_, _ = fmt.Fprint(w, resp)
	})

	resp, err := client.Volume.Detach(ctx, "894f0e66-4571-4fea-9766-5fc615aec4a5", "abc3f3cc-7ccb-11ea-945f-7be40572932e")
	require.NoError(t, err)
	assert.Equal(t, "894f0e66-4571-4fea-9766-5fc615aec4a5", resp.VolumeDetail.ID)
	assert.Equal(t, "Detach successfully", resp.Message)
}

func TestPatchVolume(t *testing.T) {
	setup()
	defer teardown()

	var v volume
	mux.HandleFunc(testlib.CloudServerURL(v.itemPath("894f0e66-4571-4fea-9766-5fc615aec4a5")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPatch, r.Method)
		resp := `
{
	"status": "in-use",
	"user_id": "894f0e66-4571-4fea-9766-5fc615aec4a5",
	"attachments": [
		{
			"server_id": "2b2628b1-0d11-4fd7-8d63-4ec24ff493ea",
			"attachment_id": "9c2ba971-dd59-4527-a441-26dcd3174d68",
			"attached_at": "2020-04-11T17:51:47.000000",
			"volume_id": "7b099bbb-21e9-48f9-8cec-4076d78fa201",
			"device": "/dev/vdb",
			"id": "7b099bbb-21e9-48f9-8cec-4076d78fa201",
			"server": {
				"OS-EXT-STS:task_state": null,
				"addresses": {
					"priv_sapd@vccloud.vn": [
						{
							"OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:98:c0:3f",
							"version": 4,
							"addr": "10.20.165.43",
							"OS-EXT-IPS:type": "fixed"
						}
					],
					"EXT_DIRECTNET_8": [
						{
							"OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:17:87:67",
							"version": 4,
							"addr": "103.107.182.201",
							"OS-EXT-IPS:type": "fixed"
						}
					]
				},
				"OS-EXT-STS:vm_state": "active",
				"OS-SRV-USG:launched_at": "2020-04-11T15:42:53.000000",
				"flavor": {
					"name": "nix.2c_2g",
					"ram": 2048,
					"OS-FLV-DISABLED:disabled": false,
					"vcpus": 2,
					"swap": "",
					"os-flavor-access:is_public": true,
					"rxtx_factor": 1.0,
					"OS-FLV-EXT-DATA:ephemeral": 0,
					"disk": 0,
					"id": "be7dab73-2c87-4d59-a2fd-49e4f7845310"
				},
				"id": "2b2628b1-0d11-4fd7-8d63-4ec24ff493ea",
				"security_groups": [
					{
						"name": "default"
					},
					{
						"name": "default"
					}
				],
				"user_id": "55d38aecb1034c06b99c1c87fb6f0740",
				"OS-DCF:diskConfig": "MANUAL",
				"accessIPv4": "",
				"accessIPv6": "",
				"progress": 0,
				"OS-EXT-STS:power_state": 1,
				"OS-EXT-AZ:availability_zone": "HN1",
				"config_drive": "",
				"status": "ACTIVE",
				"updated": "2020-04-11T15:42:54Z",
				"hostId": "023ca0b3e91f594819943b6e70ff0d0436734e19df2a2147a8eb1333",
				"OS-SRV-USG:terminated_at": null,
				"key_name": "sapd1",
				"name": "ceph-15",
				"created": "2020-04-11T15:42:35Z",
				"tenant_id": "17a1c3c952c84b3e84a82ddd48364938",
				"os-extended-volumes:volumes_attached": [
					{
						"id": "018e0772-1930-4329-a08c-b07422aa9fc1"
					},
					{
						"id": "7b099bbb-21e9-48f9-8cec-4076d78fa201"
					},
					{
						"id": "af2516f4-002d-41bc-932d-4ebdc5d49107"
					}
				],
				"metadata": {
					"category": "premium",
					"os_type": "CentOS 7.7"
				},
				"ipv6": false
			}
		}
	],
	"availability_zone": "HN1",
	"bootable": false,
	"encrypted": false,
	"created_at": "2020-04-11T17:51:41.000000",
	"description": "test_description",
	"os-vol-tenant-attr:tenant_id": "17a1c3c952c84b3e84a82ddd48364938",
	"updated_at": "2020-04-11T17:51:47.000000",
	"type": "HDD",
	"name": "sapd-vol-1",
	"replication_status": null,
	"consistencygroup_id": null,
	"source_volid": null,
	"snapshot_id": null,
	"multiattach": false,
	"metadata": {
		"category": "premium",
		"attached_mode": "rw"
	},
	"id": "7b099bbb-21e9-48f9-8cec-4076d78fa201",
	"size": 20,
	"attached_type": "datadisk",
	"category": "premium"
}		
`
		_, _ = fmt.Fprint(w, resp)
	})
	patchRequest := &VolumePatchRequest{
		Description: "test_description",
	}
	resp, err := client.Volume.Patch(ctx, "894f0e66-4571-4fea-9766-5fc615aec4a5", patchRequest)
	require.NoError(t, err)
	require.Equal(t, resp.Description, "test_description")

	}
