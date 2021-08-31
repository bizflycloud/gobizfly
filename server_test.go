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

func TestServerGet(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.CloudServerURL("/servers"), func(w http.ResponseWriter, r *http.Request) {
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
		"category": "premium",
		"region_name": "HaNoi"
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
	mux.HandleFunc(testlib.CloudServerURL("/servers"), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		var scr *ServerCreateRequest
		payload := []*ServerCreateRequest{scr}
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		assert.Equal(t, "sapd123", payload[0].Name)
		assert.Equal(t, "image", payload[0].OS.Type)
		assert.Equal(t, "2c_2g", payload[0].FlavorName)
		assert.Equal(t, "HDD", payload[0].RootDisk.Type)
		assert.Equal(t, 40, payload[0].RootDisk.Size)
		assert.Equal(t, "free_datatransfer", payload[0].NetworkPlan)
		assert.Equal(t, []string{"123", "456", "678"}, payload[0].Firewalls)
		assert.Equal(t, []string{"123", "456"}, payload[0].NetworkInterface)
		resp := `
{
	"task_id": [
		"71b9caeb-1df3-4a60-8741-fdea426fed4c"
	]
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
		Firewalls: []string{"123", "456", "678"},
		NetworkPlan: "free_datatransfer",
		NetworkInterface: []string{"123", "456"},
	}
	task, err := client.Server.Create(ctx, scr)
	require.NoError(t, err)
	assert.Equal(t, "71b9caeb-1df3-4a60-8741-fdea426fed4c", task.Task[0])

}

func TestServerDelete(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.CloudServerURL("/servers")+"/"+"c0f541d1-385a-4b0f-8c9a-5bd583475477", func(w http.ResponseWriter, r *http.Request) {

		resp := `test
		`
		_, _ = fmt.Fprint(w, resp)
	})

	err := client.Server.Delete(ctx, "c0f541d1-385a-4b0f-8c9a-5bd583475477", []string{})
	require.NoError(t, err)

}

func TestServerSoftReboot(t *testing.T) {
	setup()
	defer teardown()
	var svr *server
	mux.HandleFunc(testlib.CloudServerURL(svr.itemActionPath("6768e664-7e3e-11ea-ba40-ffdde7ae9a5b")), func(w http.ResponseWriter, r *http.Request) {
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
	mux.HandleFunc(testlib.CloudServerURL(svr.itemActionPath("6768e664-7e3e-11ea-ba40-ffdde7ae9a5b")), func(w http.ResponseWriter, r *http.Request) {
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
	mux.HandleFunc(testlib.CloudServerURL(svr.itemActionPath("5767c20e-fba4-4b23-8045-31e641d10d57")), func(w http.ResponseWriter, r *http.Request) {
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
	"ipv6": false,
	"category": "premium",
	"region_name": "HaNoi"
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
	mux.HandleFunc(testlib.CloudServerURL(svr.itemActionPath("5767c20e-fba4-4b23-8045-31e641d10d57")), func(w http.ResponseWriter, r *http.Request) {
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
	"ipv6": false,
	"category": "premium",
	"region_name": "HaNoi"
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
	mux.HandleFunc(testlib.CloudServerURL(svr.itemActionPath("5767c20e-fba4-4b23-8045-31e641d10d57")), func(w http.ResponseWriter, r *http.Request) {
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
	mux.HandleFunc(testlib.CloudServerURL(svr.itemActionPath("5767c20e-fba4-4b23-8045-31e641d10d57")), func(w http.ResponseWriter, r *http.Request) {
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
	mux.HandleFunc(testlib.CloudServerURL(svr.itemActionPath("5767c20e-fba4-4b23-8045-31e641d10d57")), func(w http.ResponseWriter, r *http.Request) {
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
	mux.HandleFunc(testlib.CloudServerURL(flavorPath), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
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
	mux.HandleFunc(testlib.CloudServerURL(osImagePath), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "True", r.URL.Query().Get("os_images"))
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

func TestGetServerTaskResponseNotReady(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.CloudServerURL(taskPath+"/7b1759dd-6e52-4799-b1ed-6441cbec1efb"), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		resp := `
{
    "ready": false,
    "result": {
        "action": "create_server",
        "progress": 100
    }
}

`
		_, _ = fmt.Fprint(w, resp)
	})

	resp, err := client.Server.GetTask(ctx, "7b1759dd-6e52-4799-b1ed-6441cbec1efb")
	require.NoError(t, err)
	assert.Equal(t, false, resp.Ready)
}

func TestGetServerTaskResponseReady(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.CloudServerURL(taskPath+"/7b1759dd-6e52-4799-b1ed-6441cbec1efb"), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		resp := `
{
    "ready": true,
    "result": {
        "id": "366d5fa3-49d2-4c0d-bde5-f542bddb212a",
        "name": "multi-server-TNnUw4Im",
        "status": "ACTIVE",
        "tenant_id": "17a1c3c952c84b3e84a82ddd48364938",
        "user_id": "55d38aecb1034c06b99c1c87fb6f0740",
        "metadata": {
            "category": "premium",
            "os_type": "Ubuntu 18.04"
        },
        "hostId": "8b22a6409fb25479713094dd4ebd424fdf237fb372fb38a2b5417e19",
        "flavor": {
            "id": "be7dab73-2c87-4d59-a2fd-49e4f7845310",
            "name": "nix.2c_2g",
            "ram": 2048,
            "disk": 0,
            "swap": "",
            "OS-FLV-EXT-DATA:ephemeral": 0,
            "OS-FLV-DISABLED:disabled": false,
            "vcpus": 2,
            "os-flavor-access:is_public": true,
            "rxtx_factor": 1.0
        },
        "created": "2020-09-22T09:48:18Z",
        "updated": "2020-09-22T09:48:35Z",
        "addresses": {
            "priv_sapd@vccloud.vn": [
                {
                    "version": 4,
                    "addr": "10.20.165.67",
                    "OS-EXT-IPS:type": "fixed",
                    "OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:a5:01:9b"
                }
            ],
            "EXT_DIRECTNET_2": [
                {
                    "version": 4,
                    "addr": "103.56.156.109",
                    "OS-EXT-IPS:type": "fixed",
                    "OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:1c:25:8e"
                }
            ]
        },
        "accessIPv4": "",
        "accessIPv6": "",
        "OS-DCF:diskConfig": "MANUAL",
        "progress": 0,
        "OS-EXT-AZ:availability_zone": "HN1",
        "config_drive": "",
        "key_name": "sapd1",
        "OS-SRV-USG:launched_at": "2020-09-22T09:48:35.000000",
        "OS-SRV-USG:terminated_at": null,
        "security_groups": [
            {
                "name": "default"
            },
            {
                "name": "default"
            }
        ],
        "OS-EXT-STS:task_state": null,
        "OS-EXT-STS:vm_state": "active",
        "OS-EXT-STS:power_state": 1,
        "os-extended-volumes:volumes_attached": [
            {
                "id": "4ca864a0-b546-4373-bda5-f338656d23d7",
                "status": "in-use",
                "size": 20,
                "availability_zone": "HN1",
                "created_at": "2020-09-22T09:47:46.000000",
                "updated_at": "2020-09-22T09:48:24.000000",
                "attachments": [
                    {
                        "id": "4ca864a0-b546-4373-bda5-f338656d23d7",
                        "attachment_id": "ca7872d6-94e6-43ad-8165-8de0acac4ea8",
                        "volume_id": "4ca864a0-b546-4373-bda5-f338656d23d7",
                        "server_id": "366d5fa3-49d2-4c0d-bde5-f542bddb212a",
                        "host_name": "thor-compute-033",
                        "device": "/dev/vda",
                        "attached_at": "2020-09-22T09:48:24.000000"
                    }
                ],
                "name": "multi-server-TNnUw4Im_rootdisk",
                "description": null,
                "volume_type": "PREMIUM_HDD",
                "snapshot_id": null,
                "source_volid": null,
                "metadata": {
                    "category": "premium"
                },
                "user_id": "55d38aecb1034c06b99c1c87fb6f0740",
                "bootable": "true",
                "encrypted": false,
                "replication_status": null,
                "consistencygroup_id": null,
                "multiattach": false,
                "os-vol-tenant-attr:tenant_id": "17a1c3c952c84b3e84a82ddd48364938",
                "volume_image_metadata": {
                    "base_image_ref": "e410a263-b265-492d-9bb1-cd8e75fc3e92",
                    "boot_roles": "member,reader,admin",
                    "image_location": "snapshot",
                    "image_state": "available",
                    "image_type": "image",
                    "instance_uuid": "498059da-ce21-498b-990b-c0f44a95cc3d",
                    "owner_id": "159c53f12fc24afc88c945e9bc6cc57d",
                    "owner_project_name": "Packer-Images",
                    "owner_user_name": "image-builder",
                    "user_id": "5676103832f14c129306bf525ec7b2de",
                    "image_id": "82774ccd-f0f7-4858-9367-fd1cd819d8a9",
                    "image_name": "Ubuntu 18.04",
                    "checksum": "dc6b626e835035be71436b35c97c330d",
                    "container_format": "bare",
                    "disk_format": "raw",
                    "min_disk": "5",
                    "min_ram": "0",
                    "size": "5368709120"
                },
                "attached_type": "rootdisk",
                "type": "HDD",
                "category": "premium",
                "snapshots": []
            }
        ],
        "category": "premium",
        "ip_addresses": {
            "LAN": [
                {
                    "version": 4,
                    "addr": "10.20.165.67",
                    "OS-EXT-IPS:type": "fixed",
                    "OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:a5:01:9b"
                }
            ],
            "WAN_V4": [
                {
                    "version": 4,
                    "addr": "103.56.156.109",
                    "OS-EXT-IPS:type": "fixed",
                    "OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:1c:25:8e"
                }
            ],
            "WAN_V6": []
        },
        "zone_name": "HN1",
        "region_name": "HaNoi",
        "autoscale_service": {},
        "ipv6": false,
        "success": true
    }
}


`
		_, _ = fmt.Fprint(w, resp)
	})

	resp, err := client.Server.GetTask(ctx, "7b1759dd-6e52-4799-b1ed-6441cbec1efb")
	require.NoError(t, err)
	assert.Equal(t, true, resp.Ready)
	assert.Equal(t, "366d5fa3-49d2-4c0d-bde5-f542bddb212a", resp.Result.Server.ID)
}

func TestServerChangeCategory(t *testing.T) {
	setup()
	defer teardown()
	var svr *server
	mux.HandleFunc(testlib.CloudServerURL(svr.itemActionPath("5767c20e-fba4-4b23-8045-31e641d10d57")), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		var sa *ServerAction
		require.NoError(t, json.NewDecoder(r.Body).Decode(&sa))
		assert.Equal(t, "change_type", sa.Action)
		assert.Equal(t, "enterprise", sa.NewType)
		resp := `
{
	"task_id": "f188d844-7e3f-11ea-a878-17c5949416eb"
}
		`
		_, _ = fmt.Fprint(w, resp)
	})

	task, err := client.Server.ChangeCategory(ctx, "5767c20e-fba4-4b23-8045-31e641d10d57", "enterprise")
	require.NoError(t, err)
	assert.Equal(t, "f188d844-7e3f-11ea-a878-17c5949416eb", task.TaskID)

}

func TestServerAddVPC(t *testing.T) {
	setup()
	defer teardown()
	var s *server
	mux.HandleFunc(testlib.CloudServerURL(s.itemActionPath("04c13e91-ede3-41b8-8824-7d3541f33b5a")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			resp := `{
    "id": "04c13e91-ede3-41b8-8824-7d3541f33b5a",
    "name": "asdfasdfasdf",
    "status": "ACTIVE",
    "tenant_id": "ebbed256d9414b0598719c42dc17e837",
    "user_id": "7156c45b82cb4fabba997a90b032c0de",
    "metadata": {
        "network_plan": "free_datatransfer",
        "category": "premium",
        "os_type": "CentOS-7.0[64-bit-version]"
    },
    "hostId": "0829fc8b8e0750802ae393a41998c4ed2ed5068254c6a96e0571dbaf",
    "flavor": {
        "id": "be7dab73-2c87-4d59-a2fd-49e4f7845310",
        "name": "2c_2g",
        "ram": 2048,
        "disk": 0,
        "swap": "",
        "OS-FLV-EXT-DATA:ephemeral": 0,
        "OS-FLV-DISABLED:disabled": false,
        "vcpus": 2,
        "os-flavor-access:is_public": true,
        "rxtx_factor": 1.0
    },
    "created": "2021-01-25T04:36:44Z",
    "updated": "2021-01-25T04:37:18Z",
    "addresses": {
        "EXT_DIRECTNET_4": [
            {
                "version": 4,
                "addr": "103.56.158.174",
                "OS-EXT-IPS:type": "fixed",
                "OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:be:98:50"
            }
        ],
        "Airflow": [
            {
                "version": 4,
                "addr": "10.23.237.44",
                "OS-EXT-IPS:type": "fixed",
                "OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:af:96:6a",
                "network_name": "Airflow"
            }
        ],
        "priv_vctest_devcs_tung491@vccloud.vn": [
            {
                "version": 4,
                "addr": "10.26.53.174",
                "OS-EXT-IPS:type": "fixed",
                "OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:dd:54:a4",
                "network_name": "priv_vctest_devcs_tung491@vccloud.vn"
            }
        ],
        "asdfasdf": [
            {
                "version": 4,
                "addr": "10.108.18.48",
                "OS-EXT-IPS:type": "fixed",
                "OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:75:63:bc",
                "network_name": "asdfasdf"
            }
        ]
    },
    "accessIPv4": "",
    "accessIPv6": "",
    "OS-DCF:diskConfig": "MANUAL",
    "progress": 0,
    "OS-EXT-AZ:availability_zone": "HN1",
    "config_drive": "",
    "key_name": null,
    "OS-SRV-USG:launched_at": "2021-01-25T04:37:17.000000",
    "OS-SRV-USG:terminated_at": null,
    "security_groups": [
        {
            "name": "default"
        },
        {
            "name": "default"
        },
        {
            "name": "default"
        },
        {
            "name": "default"
        }
    ],
    "OS-EXT-STS:task_state": null,
    "OS-EXT-STS:vm_state": "active",
    "OS-EXT-STS:power_state": 1,
    "os-extended-volumes:volumes_attached": [
        {
            "id": "b3a1486e-bc9a-4e41-908b-83a3079ba304",
            "status": "in-use",
            "size": 40,
            "availability_zone": "HN1",
            "created_at": "2021-01-25T04:35:30.000000",
            "updated_at": "2021-01-25T04:36:49.000000",
            "attachments": [
                {
                    "id": "b3a1486e-bc9a-4e41-908b-83a3079ba304",
                    "attachment_id": "a69f07c1-2455-478a-8d38-f0dcf27af00e",
                    "volume_id": "b3a1486e-bc9a-4e41-908b-83a3079ba304",
                    "server_id": "04c13e91-ede3-41b8-8824-7d3541f33b5a",
                    "host_name": "thor-compute-016",
                    "device": "/dev/vda",
                    "attached_at": "2021-01-25T04:36:49.000000"
                }
            ],
            "name": "asdfasdfasdf_rootdisk",
            "description": null,
            "volume_type": "PREMIUM_SSD",
            "snapshot_id": null,
            "source_volid": null,
            "metadata": {
                "category": "premium"
            },
            "user_id": "7156c45b82cb4fabba997a90b032c0de",
            "bootable": "true",
            "encrypted": false,
            "replication_status": null,
            "consistencygroup_id": null,
            "multiattach": false,
            "os-vol-tenant-attr:tenant_id": "ebbed256d9414b0598719c42dc17e837",
            "volume_image_metadata": {
                "signature_verified": "False",
                "image_id": "9a0f31e3-c43d-4fc2-ae1c-cc6ebde571fa",
                "image_name": "CentOS-7.0[64-bit-version]",
                "checksum": "07d74cc43fd0a3b4531673f70a3b686f",
                "container_format": "bare",
                "disk_format": "raw",
                "min_disk": "0",
                "min_ram": "0",
                "size": "4194304000"
            },
            "attached_type": "rootdisk",
            "type": "SSD",
            "category": "premium",
            "snapshots": []
        },
        {
            "id": "7878d049-2838-43a7-a5fe-8d867ec33d6b",
            "status": "in-use",
            "size": 50,
            "availability_zone": "HN1",
            "created_at": "2021-01-25T04:37:34.000000",
            "updated_at": "2021-01-25T04:37:40.000000",
            "attachments": [
                {
                    "id": "7878d049-2838-43a7-a5fe-8d867ec33d6b",
                    "attachment_id": "02ec6bce-7500-46b7-b9b7-8498bca3f72e",
                    "volume_id": "7878d049-2838-43a7-a5fe-8d867ec33d6b",
                    "server_id": "04c13e91-ede3-41b8-8824-7d3541f33b5a",
                    "host_name": "thor-compute-016",
                    "device": "/dev/vdb",
                    "attached_at": "2021-01-25T04:37:39.000000"
                }
            ],
            "name": "asdfasdfasdf_datadisk",
            "description": null,
            "volume_type": "PREMIUM_SSD",
            "snapshot_id": null,
            "source_volid": null,
            "metadata": {
                "category": "premium"
            },
            "user_id": "7156c45b82cb4fabba997a90b032c0de",
            "bootable": "false",
            "encrypted": false,
            "replication_status": null,
            "consistencygroup_id": null,
            "multiattach": false,
            "os-vol-tenant-attr:tenant_id": "ebbed256d9414b0598719c42dc17e837",
            "attached_type": "datadisk",
            "type": "SSD",
            "category": "premium",
            "snapshots": []
        }
    ],
    "locked": false,
    "category": "premium",
    "ip_addresses": {
        "LAN": [
            {
                "version": 4,
                "addr": "10.23.237.44",
                "OS-EXT-IPS:type": "fixed",
                "OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:af:96:6a",
                "network_name": "Airflow"
            },
            {
                "version": 4,
                "addr": "10.26.53.174",
                "OS-EXT-IPS:type": "fixed",
                "OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:dd:54:a4",
                "network_name": "priv_vctest_devcs_tung491@vccloud.vn"
            },
            {
                "version": 4,
                "addr": "10.108.18.48",
                "OS-EXT-IPS:type": "fixed",
                "OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:75:63:bc",
                "network_name": "asdfasdf"
            }
        ],
        "WAN_V4": [
            {
                "version": 4,
                "addr": "103.56.158.174",
                "OS-EXT-IPS:type": "fixed",
                "OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:be:98:50"
            }
        ],
        "WAN_V6": []
    },
    "zone_name": "HN1",
    "region_name": "HaNoi",
    "autoscale_service": {},
    "ipv6": false,
    "network_plan": "free_datatransfer",
    "flavor_name": "2c_2g"
}
`
			_, _ = fmt.Fprint(writer, resp)
		})
	server, err := client.Server.AddVPC(ctx, "04c13e91-ede3-41b8-8824-7d3541f33b5a", []string{"04c13e91-ede3-41b8-8824-7d3541f33b5a"})
	require.NoError(t, err)
	assert.Len(t, server.IPAddresses.LanAddresses, 3)
}

func TestServerRemoveVPC(t *testing.T) {
	setup()
	defer teardown()
	var s *server
	mux.HandleFunc(testlib.CloudServerURL(s.itemActionPath("04c13e91-ede3-41b8-8824-7d3541f33b5a")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			resp := `{
    "id": "04c13e91-ede3-41b8-8824-7d3541f33b5a",
    "name": "asdfasdfasdf",
    "status": "ACTIVE",
    "tenant_id": "ebbed256d9414b0598719c42dc17e837",
    "user_id": "7156c45b82cb4fabba997a90b032c0de",
    "metadata": {
        "network_plan": "free_datatransfer",
        "category": "premium",
        "os_type": "CentOS-7.0[64-bit-version]"
    },
    "hostId": "0829fc8b8e0750802ae393a41998c4ed2ed5068254c6a96e0571dbaf",
    "flavor": {
        "id": "be7dab73-2c87-4d59-a2fd-49e4f7845310",
        "name": "2c_2g",
        "ram": 2048,
        "disk": 0,
        "swap": "",
        "OS-FLV-EXT-DATA:ephemeral": 0,
        "OS-FLV-DISABLED:disabled": false,
        "vcpus": 2,
        "os-flavor-access:is_public": true,
        "rxtx_factor": 1.0
    },
    "created": "2021-01-25T04:36:44Z",
    "updated": "2021-01-25T04:37:18Z",
    "addresses": {
        "EXT_DIRECTNET_4": [
            {
                "version": 4,
                "addr": "103.56.158.174",
                "OS-EXT-IPS:type": "fixed",
                "OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:be:98:50"
            }
        ],
        "Airflow": [
            {
                "version": 4,
                "addr": "10.23.237.44",
                "OS-EXT-IPS:type": "fixed",
                "OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:af:96:6a",
                "network_name": "Airflow"
            }
        ],
        "priv_vctest_devcs_tung491@vccloud.vn": [
            {
                "version": 4,
                "addr": "10.26.53.174",
                "OS-EXT-IPS:type": "fixed",
                "OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:dd:54:a4",
                "network_name": "priv_vctest_devcs_tung491@vccloud.vn"
            }
        ],
        "asdfasdf": [
            {
                "version": 4,
                "addr": "10.108.18.48",
                "OS-EXT-IPS:type": "fixed",
                "OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:75:63:bc",
                "network_name": "asdfasdf"
            }
        ]
    },
    "accessIPv4": "",
    "accessIPv6": "",
    "OS-DCF:diskConfig": "MANUAL",
    "progress": 0,
    "OS-EXT-AZ:availability_zone": "HN1",
    "config_drive": "",
    "key_name": null,
    "OS-SRV-USG:launched_at": "2021-01-25T04:37:17.000000",
    "OS-SRV-USG:terminated_at": null,
    "security_groups": [
        {
            "name": "default"
        },
        {
            "name": "default"
        },
        {
            "name": "default"
        },
        {
            "name": "default"
        }
    ],
    "OS-EXT-STS:task_state": null,
    "OS-EXT-STS:vm_state": "active",
    "OS-EXT-STS:power_state": 1,
    "os-extended-volumes:volumes_attached": [
        {
            "id": "b3a1486e-bc9a-4e41-908b-83a3079ba304",
            "status": "in-use",
            "size": 40,
            "availability_zone": "HN1",
            "created_at": "2021-01-25T04:35:30.000000",
            "updated_at": "2021-01-25T04:36:49.000000",
            "attachments": [
                {
                    "id": "b3a1486e-bc9a-4e41-908b-83a3079ba304",
                    "attachment_id": "a69f07c1-2455-478a-8d38-f0dcf27af00e",
                    "volume_id": "b3a1486e-bc9a-4e41-908b-83a3079ba304",
                    "server_id": "04c13e91-ede3-41b8-8824-7d3541f33b5a",
                    "host_name": "thor-compute-016",
                    "device": "/dev/vda",
                    "attached_at": "2021-01-25T04:36:49.000000"
                }
            ],
            "name": "asdfasdfasdf_rootdisk",
            "description": null,
            "volume_type": "PREMIUM_SSD",
            "snapshot_id": null,
            "source_volid": null,
            "metadata": {
                "category": "premium"
            },
            "user_id": "7156c45b82cb4fabba997a90b032c0de",
            "bootable": "true",
            "encrypted": false,
            "replication_status": null,
            "consistencygroup_id": null,
            "multiattach": false,
            "os-vol-tenant-attr:tenant_id": "ebbed256d9414b0598719c42dc17e837",
            "volume_image_metadata": {
                "signature_verified": "False",
                "image_id": "9a0f31e3-c43d-4fc2-ae1c-cc6ebde571fa",
                "image_name": "CentOS-7.0[64-bit-version]",
                "checksum": "07d74cc43fd0a3b4531673f70a3b686f",
                "container_format": "bare",
                "disk_format": "raw",
                "min_disk": "0",
                "min_ram": "0",
                "size": "4194304000"
            },
            "attached_type": "rootdisk",
            "type": "SSD",
            "category": "premium",
            "snapshots": []
        },
        {
            "id": "7878d049-2838-43a7-a5fe-8d867ec33d6b",
            "status": "in-use",
            "size": 50,
            "availability_zone": "HN1",
            "created_at": "2021-01-25T04:37:34.000000",
            "updated_at": "2021-01-25T04:37:40.000000",
            "attachments": [
                {
                    "id": "7878d049-2838-43a7-a5fe-8d867ec33d6b",
                    "attachment_id": "02ec6bce-7500-46b7-b9b7-8498bca3f72e",
                    "volume_id": "7878d049-2838-43a7-a5fe-8d867ec33d6b",
                    "server_id": "04c13e91-ede3-41b8-8824-7d3541f33b5a",
                    "host_name": "thor-compute-016",
                    "device": "/dev/vdb",
                    "attached_at": "2021-01-25T04:37:39.000000"
                }
            ],
            "name": "asdfasdfasdf_datadisk",
            "description": null,
            "volume_type": "PREMIUM_SSD",
            "snapshot_id": null,
            "source_volid": null,
            "metadata": {
                "category": "premium"
            },
            "user_id": "7156c45b82cb4fabba997a90b032c0de",
            "bootable": "false",
            "encrypted": false,
            "replication_status": null,
            "consistencygroup_id": null,
            "multiattach": false,
            "os-vol-tenant-attr:tenant_id": "ebbed256d9414b0598719c42dc17e837",
            "attached_type": "datadisk",
            "type": "SSD",
            "category": "premium",
            "snapshots": []
        }
    ],
    "locked": false,
    "category": "premium",
    "ip_addresses": {
        "LAN": [
            {
                "version": 4,
                "addr": "10.23.237.44",
                "OS-EXT-IPS:type": "fixed",
                "OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:af:96:6a",
                "network_name": "Airflow"
            },
            {
                "version": 4,
                "addr": "10.26.53.174",
                "OS-EXT-IPS:type": "fixed",
                "OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:dd:54:a4",
                "network_name": "priv_vctest_devcs_tung491@vccloud.vn"
            }
        ],
        "WAN_V4": [
            {
                "version": 4,
                "addr": "103.56.158.174",
                "OS-EXT-IPS:type": "fixed",
                "OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:be:98:50"
            }
        ],
        "WAN_V6": []
    },
    "zone_name": "HN1",
    "region_name": "HaNoi",
    "autoscale_service": {},
    "ipv6": false,
    "network_plan": "free_datatransfer",
    "flavor_name": "2c_2g"
}
`
			_, _ = fmt.Fprint(writer, resp)
		})
	server, err := client.Server.RemoveVPC(ctx, "04c13e91-ede3-41b8-8824-7d3541f33b5a", []string{"04c13e91-ede3-41b8-8824-7d3541f33b5a"})
	require.NoError(t, err)
	assert.Len(t, server.IPAddresses.LanAddresses, 2)
}

func TestCustomImageList(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.CloudServerURL(customImagesPath), func(writer http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `{
    "images": [
        {
            "description": "",
            "name": "Cloud-server.postman_collection.json",
            "disk_format": "raw",
            "container_format": "bare",
            "visibility": "private",
            "size": 40838,
            "virtual_size": null,
            "status": "active",
            "checksum": "d9d256a9a06ae5eb18579987a74ec8fb",
            "protected": false,
            "min_ram": 0,
            "min_disk": 0,
            "owner": "bc8d2790fc9a46949818b942c0a824de",
            "os_hidden": false,
            "os_hash_algo": "sha512",
            "os_hash_value": "a7572c56048cd67c20f23d908bd04d29a86ea53bbebe3e011e8d4a3a216ff7faa07810e480ea05409543ba0a0345f5b31ac136aa0f05788d34dbb210ae586049",
            "id": "13e73169-d92c-4ec0-b839-1a22a093cbc4",
            "created_at": "2021-02-05T02:03:00Z",
            "updated_at": "2021-02-05T02:03:13Z",
            "locations": [
                {
                    "url": "rbd://46b21e69-6559-4eed-aba9-540ab496a77d/glances/13e73169-d92c-4ec0-b839-1a22a093cbc4/snap",
                    "metadata": {}
                }
            ],
            "direct_url": "rbd://46b21e69-6559-4eed-aba9-540ab496a77d/glances/13e73169-d92c-4ec0-b839-1a22a093cbc4/snap",
            "tags": [],
            "file": "https://hn-1.bizflycloud.vn:9292/v2/images/13e73169-d92c-4ec0-b839-1a22a093cbc4/file",
            "schema": "/v2/schemas/image"
        }
    ]
}
`
		_, _ = fmt.Fprint(writer, resp)
	})
	images, err := client.Server.ListCustomImages(ctx)
	require.NoError(t, err)
	assert.Equal(t, "/v2/schemas/image", images[0].Schema)
}

func TestCustomImageCreate(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.CloudServerURL(strings.Join([]string{customImagePath, "upload"}, "/")), func(writer http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		var payload *CreateCustomImagePayload
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		resp := `{
    "image": {
        "description": "",
        "name": "Cloud-server.postman_collection.jsonafasds",
        "disk_format": "raw",
        "container_format": "bare",
        "visibility": "private",
        "size": null,
        "virtual_size": null,
        "status": "queued",
        "checksum": null,
        "protected": false,
        "min_ram": 0,
        "min_disk": 0,
        "owner": "bc8d2790fc9a46949818b942c0a824de",
        "os_hidden": false,
        "os_hash_algo": null,
        "os_hash_value": null,
        "id": "0b5ae9ed-7cfb-454b-a5cc-df0bba693532",
        "created_at": "2021-02-05T02:05:11Z",
        "updated_at": "2021-02-05T02:05:11Z",
        "locations": [],
        "tags": [],
        "file": "/v2/images/0b5ae9ed-7cfb-454b-a5cc-df0bba693532/file",
        "schema": "/v2/schemas/image"
    },
    "success": true,
    "token": "gAAAAABgHKajmGPDwMPFTqpoZLQ34D7shjMQI-_zTsANuXpcybC7786K-wE1FPegZa71cst1GZNp2NXejugv8jBZLOhUqaHuR_HtZWLQWy4RIXwY-w5m-YSWzHJb4aCK1gmPP9ZoN7i3oylr7THieBjA3t_BsRj-jvgWCtjStoevYPOQ-zwz7JY",
    "upload_uri": "https://hn-1.bizflycloud.vn:9292/v2/images/0b5ae9ed-7cfb-454b-a5cc-df0bba693532/file"
}
`
		_, _ = fmt.Fprint(writer, resp)
	})
	resp, err := client.Server.CreateCustomImage(ctx, &CreateCustomImagePayload{
		Name:       "Cloud-server.postman_collection.jsonafasds",
		DiskFormat: "raw",
	})
	require.NoError(t, err)
	assert.Equal(t, "bare", resp.Image.ContainerFormat)
}

func TestCustomImageDelete(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.CloudServerURL(strings.Join([]string{customImagePath, "0b5ae9ed-7cfb-454b-a5cc-df0bba693532"}, "/")),
		func(writer http.ResponseWriter, r *http.Request) {
			require.Equal(t, http.MethodDelete, r.Method)
		})
	err := client.Server.DeleteCustomImage(ctx, "0b5ae9ed-7cfb-454b-a5cc-df0bba693532")
	require.NoError(t, err)
}

func TestCustomImageDownload(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.CloudServerURL(strings.Join([]string{customImagePath, "809cd001-07e3-4e00-a5c5-2d8fa0516d90"}, "/")),
		func(writer http.ResponseWriter, r *http.Request) {
			require.Equal(t, http.MethodGet, r.Method)
			resp := `{
    "image": {
        "description": "",
        "name": "large-file",
        "disk_format": "raw",
        "container_format": "bare",
        "visibility": "private",
        "size": 5561008296,
        "virtual_size": null,
        "status": "active",
        "checksum": "f4c392cfbbecc6656b547b1b55afcd16",
        "protected": false,
        "min_ram": 0,
        "min_disk": 0,
        "owner": "bc8d2790fc9a46949818b942c0a824de",
        "os_hidden": false,
        "os_hash_algo": "sha512",
        "os_hash_value": "dc44107ab48eaa6d4ae60c6e82acdf0098f129805e61b8cc00c2df7e3568e3407ba2b999bd387c2b0ac18fe9faf82b21b5a109eb27bc932d3d85803a521e4ef0",
        "id": "809cd001-07e3-4e00-a5c5-2d8fa0516d90",
        "created_at": "2021-02-08T14:20:05Z",
        "updated_at": "2021-02-08T14:31:16Z",
        "locations": [
            {
                "url": "rbd://46b21e69-6559-4eed-aba9-540ab496a77d/glances/809cd001-07e3-4e00-a5c5-2d8fa0516d90/snap",
                "metadata": {}
            }
        ],
        "direct_url": "rbd://46b21e69-6559-4eed-aba9-540ab496a77d/glances/809cd001-07e3-4e00-a5c5-2d8fa0516d90/snap",
        "tags": [],
        "file": "https://hn-1.bizflycloud.vn:9292/v2/images/809cd001-07e3-4e00-a5c5-2d8fa0516d90/file",
        "schema": "/v2/schemas/image"
    },
    "token": "gAAAAABgQY1kuIcMMK17B1nWBMd5dHTcitsChmDv1WU4xIGgHph09_Do4wDjY8V5XTyTYWgeYVaN5cNJG2In3oyvSm6uYlugYk6nxC-XiD81rp-8zxvSPTDd2jlaKfrrFgd2HJH5rfpV5iQlMMps52vGpxFrxFa-2ppK8BQbye63nKK_e4LwahE"
}
`
			_, _ = fmt.Fprint(writer, resp)
		})
	image, err := client.Server.GetCustomImage(ctx, "809cd001-07e3-4e00-a5c5-2d8fa0516d90")
	require.NoError(t, err)
	assert.Equal(t, image.Token, "gAAAAABgQY1kuIcMMK17B1nWBMd5dHTcitsChmDv1WU4xIGgHph09_Do4wDjY8V5XTyTYWgeYVaN5cNJG2In3oyvSm6uYlugYk6nxC-XiD81rp-8zxvSPTDd2jlaKfrrFgd2HJH5rfpV5iQlMMps52vGpxFrxFa-2ppK8BQbye63nKK_e4LwahE")
}
