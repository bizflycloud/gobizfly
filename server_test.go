// This file is part of gobizfly
//
// Copyright (C) 2020  BizFly Cloud
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>

package gobizfly

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServerGet(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(serverBasePath, func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
[
	{
		"OS-EXT-STS:task_state": null,
		"addresses": {
			"EXT_DIRECTNET_4": [
				{
					"OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:c3:5f:16",
					"version": 4,
					"addr": "103.56.158.222",
					"OS-EXT-IPS:type": "fixed"
				}
			],
			"priv_sapd@vccloud.vn": [
				{
					"OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:aa:ea:bd",
					"version": 4,
					"addr": "10.20.165.11",
					"OS-EXT-IPS:type": "fixed"
				}
			]
		},
		"OS-EXT-STS:vm_state": "active",
		"OS-SRV-USG:launched_at": "2020-04-08T09:36:25.000000",
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
		"id": "c0f541d1-385a-4b0f-8c9a-5bd583475477",
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
		"updated": "2020-04-08T09:36:26Z",
		"hostId": "5af42401f6f37e199d7d73a6081a83bc49ee1859b6e836a7990c0907",
		"OS-SRV-USG:terminated_at": null,
		"key_name": "sapd1",
		"name": "meeting-now-1",
		"created": "2020-04-08T09:36:08Z",
		"tenant_id": "17a1c3c952c84b3e84a82ddd48364938",
		"os-extended-volumes:volumes_attached": [
			{
				"attachments": [
					{
						"server_id": "c0f541d1-385a-4b0f-8c9a-5bd583475477",
						"attachment_id": "58072f06-d697-466d-8515-26a9cb823938",
						"attached_at": "2020-04-08T09:36:15.000000",
						"host_name": "thor-compute-005",
						"volume_id": "71b9caeb-1df3-4a60-8741-fdea426fed4c",
						"device": "/dev/vda",
						"id": "71b9caeb-1df3-4a60-8741-fdea426fed4c"
					}
				],
				"availability_zone": "HN1",
				"encrypted": false,
				"updated_at": "2020-04-08T09:36:16.000000",
				"replication_status": null,
				"snapshot_id": null,
				"id": "71b9caeb-1df3-4a60-8741-fdea426fed4c",
				"size": 20,
				"user_id": "55d38aecb1034c06b99c1c87fb6f0740",
				"os-vol-tenant-attr:tenant_id": "17a1c3c952c84b3e84a82ddd48364938",
				"metadata": {
					"category": "premium",
					"attached_mode": "rw"
				},
				"status": "in-use",
				"volume_image_metadata": {
					"image_location": "snapshot",
					"image_state": "available",
					"container_format": "bare",
					"min_ram": "0",
					"image_name": "Jitsi Ubuntu 18.04",
					"boot_roles": "admin",
					"image_id": "62637635-ed3f-41da-9655-438bf38ceba4",
					"owner_user_name": "duylkops",
					"min_disk": "5",
					"base_image_ref": "e410a263-b265-492d-9bb1-cd8e75fc3e92",
					"size": "5368709120",
					"instance_uuid": "4b8a24a0-36aa-47e7-9ab0-d470240b4461",
					"user_id": "5676103832f14c129306bf525ec7b2de",
					"image_type": "image",
					"checksum": "72f0c680a49eb86a1fc416b03aae63d0",
					"disk_format": "raw",
					"owner_project_name": "Packer-Images",
					"owner_id": "159c53f12fc24afc88c945e9bc6cc57d"
				},
				"description": null,
				"multiattach": false,
				"source_volid": null,
				"consistencygroup_id": null,
				"name": "meeting-now-1_rootdisk",
				"bootable": "true",
				"created_at": "2020-04-08T09:35:53.000000",
				"volume_type": "HDD",
				"attached_type": "rootdisk",
				"snapshots": []
			}
		],
		"metadata": {
			"category": "premium",
			"os_type": "Jitsi Ubuntu 18.04",
			"prebuild_app_name": "Jitsi",
			"service": "prebuild_app"
		},
		"ipv6": false,
		"category": "premium"
	}
]
		`
		_, _ = fmt.Fprint(w, resp)
	})

	servers, err := client.Server.List(ctx, &ListOptions{})
	require.NoError(t, err)
	server := servers[0]
	assert.Equal(t, "c0f541d1-385a-4b0f-8c9a-5bd583475477", server.ID)
	assert.Equal(t, "meeting-now-1", server.Name)
	assert.Equal(t, "sapd1", server.KeyName)
	assert.Equal(t, "ACTIVE", server.Status)
	assert.Equal(t, false, server.IPv6)
	assert.Equal(t, "premium", server.Category)

}

func TestServerCreate(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(serverBasePath, func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		var scr *ServerCreateRequest
		payload := []*ServerCreateRequest{scr}
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		assert.Equal(t, "sapd123", payload[0].Name)
		assert.Equal(t, "image", payload[0].OS.Type)
		assert.Equal(t, "2c_2g", payload[0].FlavorName)
		assert.Equal(t, "HDD", payload[0].RootDisk.Type)
		assert.Equal(t, 40, payload[0].RootDisk.Size)
		resp := `
{
	"task_id": "71b9caeb-1df3-4a60-8741-fdea426fed4c"
}
`
		_, _ = fmt.Fprint(w, resp)
	})
	scr := &ServerCreateRequest{
		Name:             "sapd123",
		FlavorName:       "2c_2g",
		SSHKey:           "sapd1",
		Password:         true,
		RootDisk:         &ServerDisk{40, "HDD"},
		Type:             "premium",
		AvailabilityZone: "HN1",
		OS:               &ServerOS{"cbf5f34b-751b-42a5-830f-6b2324f61d5a", "image"},
	}
	task, err := client.Server.Create(ctx, scr)
	require.NoError(t, err)
	assert.Equal(t, "71b9caeb-1df3-4a60-8741-fdea426fed4c", task.TaskID)

}

func TestServerDelete(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(serverBasePath+"/"+"c0f541d1-385a-4b0f-8c9a-5bd583475477", func(w http.ResponseWriter, r *http.Request) {

		resp := `test
		`
		_, _ = fmt.Fprint(w, resp)
	})

	err := client.Server.Delete(ctx, "c0f541d1-385a-4b0f-8c9a-5bd583475477")
	require.NoError(t, err)

}

func TestServerSoftReboot(t *testing.T) {
	setup()
	defer teardown()
	var svr *server
	mux.HandleFunc(svr.itemActionPath("6768e664-7e3e-11ea-ba40-ffdde7ae9a5b"), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		var sa *ServerAction
		require.NoError(t, json.NewDecoder(r.Body).Decode(&sa))
		assert.Equal(t, "soft_reboot", sa.Action)
		resp := `
{
	"message": "Soft reboot server th\u00e0nh c\u00f4ng"
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	response, err := client.Server.SoftReboot(ctx, "6768e664-7e3e-11ea-ba40-ffdde7ae9a5b")
	require.NoError(t, err)
	assert.Equal(t, "Soft reboot server th\u00e0nh c\u00f4ng", response.Message)

}

func TestServerHardReboot(t *testing.T) {
	setup()
	defer teardown()
	var svr *server
	mux.HandleFunc(svr.itemActionPath("6768e664-7e3e-11ea-ba40-ffdde7ae9a5b"), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		var sa *ServerAction
		require.NoError(t, json.NewDecoder(r.Body).Decode(&sa))
		assert.Equal(t, "hard_reboot", sa.Action)
		resp := `
{
	"message": "Hard reboot server th\u00e0nh c\u00f4ng"
}`
		_, _ = fmt.Fprint(w, resp)
	})

	response, err := client.Server.HardReboot(ctx, "6768e664-7e3e-11ea-ba40-ffdde7ae9a5b")
	require.NoError(t, err)
	assert.Equal(t, "Hard reboot server th\u00e0nh c\u00f4ng", response.Message)

}

func TestServerStart(t *testing.T) {
	setup()
	defer teardown()
	var svr *server
	mux.HandleFunc(svr.itemActionPath("5767c20e-fba4-4b23-8045-31e641d10d57"), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		var sa *ServerAction
		require.NoError(t, json.NewDecoder(r.Body).Decode(&sa))
		assert.Equal(t, "start", sa.Action)
		resp := `
{
	"OS-EXT-STS:task_state": null,
	"addresses": {
		"priv_sapd@vccloud.vn": [
			{
				"OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:49:27:4b",
				"version": 4,
				"addr": "10.20.165.22",
				"OS-EXT-IPS:type": "fixed"
			}
		],
		"EXT_DIRECTNET_9": [
			{
				"OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:46:ed:7d",
				"version": 4,
				"addr": "103.107.183.146",
				"OS-EXT-IPS:type": "fixed"
			}
		]
	},
	"OS-EXT-STS:vm_state": "stopped",
	"OS-SRV-USG:launched_at": "2020-04-14T10:31:17.000000",
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
	"id": "5767c20e-fba4-4b23-8045-31e641d10d57",
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
	"OS-EXT-STS:power_state": 4,
	"OS-EXT-AZ:availability_zone": "HN1",
	"config_drive": "",
	"status": "SHUTOFF",
	"updated": "2020-04-14T11:02:39Z",
	"hostId": "74ca4ef173ad2fd2e875a30ee7f594072ba1367ac3d963532f2430a1",
	"OS-SRV-USG:terminated_at": null,
	"key_name": "sapd1",
	"name": "sapd12345x",
	"created": "2020-04-14T10:30:59Z",
	"tenant_id": "17a1c3c952c84b3e84a82ddd48364938",
	"os-extended-volumes:volumes_attached": [
		{
			"id": "ef173de9-d587-4570-b9a9-9baf760c8b85"
		}
	],
	"metadata": {
		"category": "premium",
		"os_type": "Ubuntu 18.04"
	},
	"ipv6": false
}`
		_, _ = fmt.Fprint(w, resp)
	})

	server, err := client.Server.Start(ctx, "5767c20e-fba4-4b23-8045-31e641d10d57")
	require.NoError(t, err)
	assert.Equal(t, "5767c20e-fba4-4b23-8045-31e641d10d57", server.ID)

}

func TestServerStop(t *testing.T) {
	setup()
	defer teardown()
	var svr *server
	mux.HandleFunc(svr.itemActionPath("5767c20e-fba4-4b23-8045-31e641d10d57"), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		var sa *ServerAction
		require.NoError(t, json.NewDecoder(r.Body).Decode(&sa))
		assert.Equal(t, "stop", sa.Action)
		resp := `
{
	"OS-EXT-STS:task_state": null,
	"addresses": {
		"priv_sapd@vccloud.vn": [
			{
				"OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:49:27:4b",
				"version": 4,
				"addr": "10.20.165.22",
				"OS-EXT-IPS:type": "fixed"
			}
		],
		"EXT_DIRECTNET_9": [
			{
				"OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:46:ed:7d",
				"version": 4,
				"addr": "103.107.183.146",
				"OS-EXT-IPS:type": "fixed"
			}
		]
	},
	"OS-EXT-STS:vm_state": "stopped",
	"OS-SRV-USG:launched_at": "2020-04-14T10:31:17.000000",
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
	"id": "5767c20e-fba4-4b23-8045-31e641d10d57",
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
	"OS-EXT-STS:power_state": 4,
	"OS-EXT-AZ:availability_zone": "HN1",
	"config_drive": "",
	"status": "SHUTOFF",
	"updated": "2020-04-14T11:02:39Z",
	"hostId": "74ca4ef173ad2fd2e875a30ee7f594072ba1367ac3d963532f2430a1",
	"OS-SRV-USG:terminated_at": null,
	"key_name": "sapd1",
	"name": "sapd12345x",
	"created": "2020-04-14T10:30:59Z",
	"tenant_id": "17a1c3c952c84b3e84a82ddd48364938",
	"os-extended-volumes:volumes_attached": [
		{
			"id": "ef173de9-d587-4570-b9a9-9baf760c8b85"
		}
	],
	"metadata": {
		"category": "premium",
		"os_type": "Ubuntu 18.04"
	},
	"ipv6": false
}`
		_, _ = fmt.Fprint(w, resp)
	})

	server, err := client.Server.Stop(ctx, "5767c20e-fba4-4b23-8045-31e641d10d57")
	require.NoError(t, err)
	assert.Equal(t, "5767c20e-fba4-4b23-8045-31e641d10d57", server.ID)

}

func TestServerRebuild(t *testing.T) {
	setup()
	defer teardown()
	var svr *server
	mux.HandleFunc(svr.itemActionPath("5767c20e-fba4-4b23-8045-31e641d10d57"), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		var sa *ServerAction
		require.NoError(t, json.NewDecoder(r.Body).Decode(&sa))
		assert.Equal(t, "rebuild", sa.Action)
		assert.Equal(t, "263457d0-7e40-11ea-99fe-3b298a7e3e62", sa.ImageID)
		resp := `
{
	"task_id": "f188d844-7e3f-11ea-a878-17c5949416eb"
}
		`
		_, _ = fmt.Fprint(w, resp)
	})

	task, err := client.Server.Rebuild(ctx, "5767c20e-fba4-4b23-8045-31e641d10d57", "263457d0-7e40-11ea-99fe-3b298a7e3e62")
	require.NoError(t, err)
	assert.Equal(t, "f188d844-7e3f-11ea-a878-17c5949416eb", task.TaskID)

}

func TestServerGetVNC(t *testing.T) {
	setup()
	defer teardown()
	var svr *server
	mux.HandleFunc(svr.itemActionPath("5767c20e-fba4-4b23-8045-31e641d10d57"), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		var sa *ServerAction
		require.NoError(t, json.NewDecoder(r.Body).Decode(&sa))
		assert.Equal(t, "get_vnc", sa.Action)
		assert.Equal(t, "novnc", sa.ConsoleType)
		resp := `
{
	"console": {
		"url": "https://hn-1.vccloud.vn:6080/vnc_auto.html?token=d2f12bd2-631c-4e97-950f-f8c6b3fce1cb",
		"type": "novnc"
	}
}`
		_, _ = fmt.Fprint(w, resp)
	})

	console, err := client.Server.GetVNC(ctx, "5767c20e-fba4-4b23-8045-31e641d10d57")
	require.NoError(t, err)
	assert.Equal(t, "novnc", console.Type)
	assert.Equal(t, "https://hn-1.vccloud.vn:6080/vnc_auto.html?token=d2f12bd2-631c-4e97-950f-f8c6b3fce1cb", console.URL)

}

func TestServerResize(t *testing.T) {
	setup()
	defer teardown()
	var svr *server
	mux.HandleFunc(svr.itemActionPath("5767c20e-fba4-4b23-8045-31e641d10d57"), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		var sa *ServerAction
		require.NoError(t, json.NewDecoder(r.Body).Decode(&sa))
		assert.Equal(t, "resize", sa.Action)
		assert.Equal(t, "4c_4g", sa.FlavorName)
		resp := `
{
	"task_id": "6ac1c3aa-7e41-11ea-a8b0-9b7b1be3dcee"
}
		`
		_, _ = fmt.Fprint(w, resp)
	})

	task, err := client.Server.Resize(ctx, "5767c20e-fba4-4b23-8045-31e641d10d57", "4c_4g")
	require.NoError(t, err)
	assert.Equal(t, "6ac1c3aa-7e41-11ea-a8b0-9b7b1be3dcee", task.TaskID)

}

func TestServerFlavorList(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(flavorPath, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		//TODO assert query paramter
		resp := `
[
    {
        "_id": "5d7f58903c4c0127da9896ae",
        "name": "1c_1g"
    },
    {
        "_id": "5d7f58903c4c0127da9896b5",
        "name": "2c_4g"
    }
]
`
		_, _ = fmt.Fprint(w, resp)
	})

	flavors, err := client.Server.ListFlavors(ctx)
	require.NoError(t, err)
	assert.Equal(t, "1c_1g", flavors[0].Name)
}

func TestOSImageList(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(osImagePath, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		resp := `
{
    "os_images": [
        {
            "os": "Ubuntu",
            "versions": [
                {
                    "name": "18.04 x64",
                    "id": "d9513b4e-60c4-45c6-a8e0-0d814a7c0799"
                },
                {
                    "name": "16.04 x64",
                    "id": "ffb85594-a23a-4bc6-bcd5-02039e6d0a03"
                },
                {
                    "name": "14.04 x64",
                    "id": "83aeff78-14ad-498d-976b-ed15bc5fa5ac"
                }
            ]
        },
        {
            "os": "CentOS",
            "versions": [
                {
                    "name": "8.0 x64",
                    "id": "a992ab85-e8ea-497a-9eba-29a37f7c3151"
                },
                {
                    "name": "7.7 x64",
                    "id": "eba5353f-8524-4d29-ac68-984b1c80e693"
                },
                {
                    "name": "6.10 x64",
                    "id": "dd3b7e3d-f8ab-4856-9bb1-9737828bd1b1"
                },
                {
                    "name": "6.8 x64",
                    "id": "f9dbf562-a637-498b-8af4-b4e7aaf20cdd"
                }
            ]
        },
        {
            "os": "Debian",
            "versions": [
                {
                    "name": "9 x64",
                    "id": "dddc47ec-da9c-4010-ae16-86272bd192eb"
                },
                {
                    "name": "10 x64",
                    "id": "28333e61-70b6-4bbf-9cbf-33b57194c389"
                }
            ]
        },
        {
            "os": "Windows",
            "versions": [
                {
                    "name": "2019 Standard",
                    "id": "9fabfae0-a06d-44cd-a2ae-c8d38e7ab5be"
                },
                {
                    "name": "2016 Standard",
                    "id": "9e09a71c-ceed-4f00-aeea-ed3c7d391807"
                },
                {
                    "name": "2016 Datacenter",
                    "id": "91dff6d8-e26b-4b9a-9299-f74ee8a3de02"
                },
                {
                    "name": "2012 R2 Standard",
                    "id": "ff9bacce-bce9-4ea1-9a31-93b14b788c48"
                },
                {
                    "name": "2012 R2 Datacenter",
                    "id": "e18e42bd-4141-45f4-b4c1-8e9e82fd9e87"
                },
                {
                    "name": "2008 R2 Enterprise",
                    "id": "c900de6e-dc7b-4ab9-8e13-12fdcf5b0f84"
                }
            ]
        }
    ]
}
`
		_, _ = fmt.Fprint(w, resp)
	})
	osImages, err := client.Server.ListOSImages(ctx)
	require.NoError(t, err)
	assert.Equal(t, "d9513b4e-60c4-45c6-a8e0-0d814a7c0799", osImages[0].Version[0].ID)
}
