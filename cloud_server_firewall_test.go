// This file is part of gobizfly

package gobizfly

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/bizflycloud/gobizfly/testlib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFirewallList(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.CloudServerURL(firewallBasePath), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
[
    {
        "id": "b1a01fdf-492b-48aa-9b7b-399509cbb5e4",
        "name": "sapd-fw-hcm",
        "tenant_id": "17a1c3c952c84b3e84a82ddd48364938",
        "description": "sapd-fw-hcm created 2020-09-08 14:33:24.107718",
        "tags": [],
        "created_at": "2020-09-08T07:33:24Z",
        "updated_at": "2020-09-08T07:33:25Z",
        "revision_number": 7,
        "project_id": "17a1c3c952c84b3e84a82ddd48364938",
        "servers_count": 2,
        "servers": [
            "b6c43097-27b2-4f09-8498-b7f4067909a0",
            "7371145f-bed6-4dbd-bf5a-21b48ab1102f"
        ],
        "rules_count": 4,
        "inbound": [
            {
                "id": "4ad34f53-e82f-40a5-9c79-c21695145f5c",
                "tenant_id": "17a1c3c952c84b3e84a82ddd48364938",
                "security_group_id": "b1a01fdf-492b-48aa-9b7b-399509cbb5e4",
                "ethertype": "IPv4",
                "direction": "ingress",
                "protocol": "tcp",
                "port_range_min": 22,
                "port_range_max": 22,
                "remote_ip_prefix": "0.0.0.0/0",
                "remote_group_id": null,
                "description": "SSH",
                "tags": [],
                "created_at": "2020-09-08T07:33:24Z",
                "updated_at": "2020-09-08T07:33:24Z",
                "revision_number": 0,
                "project_id": "17a1c3c952c84b3e84a82ddd48364938",
                "type": "SSH",
                "cidr": "0.0.0.0/0",
                "port_range": "22"
            },
            {
                "id": "a4782655-d73c-4132-9f0f-032382bb55c9",
                "tenant_id": "17a1c3c952c84b3e84a82ddd48364938",
                "security_group_id": "b1a01fdf-492b-48aa-9b7b-399509cbb5e4",
                "ethertype": "IPv4",
                "direction": "ingress",
                "protocol": "icmp",
                "port_range_min": null,
                "port_range_max": null,
                "remote_ip_prefix": "0.0.0.0/0",
                "remote_group_id": null,
                "description": "PING",
                "tags": [],
                "created_at": "2020-09-08T07:33:24Z",
                "updated_at": "2020-09-08T07:33:24Z",
                "revision_number": 0,
                "project_id": "17a1c3c952c84b3e84a82ddd48364938",
                "type": "PING",
                "cidr": "0.0.0.0/0",
                "port_range": null
            },
            {
                "id": "e12a161e-b008-47d5-b427-45dc1a24bc94",
                "tenant_id": "17a1c3c952c84b3e84a82ddd48364938",
                "security_group_id": "b1a01fdf-492b-48aa-9b7b-399509cbb5e4",
                "ethertype": "IPv4",
                "direction": "ingress",
                "protocol": "tcp",
                "port_range_min": 3389,
                "port_range_max": 3389,
                "remote_ip_prefix": "0.0.0.0/0",
                "remote_group_id": null,
                "description": "RDP",
                "tags": [],
                "created_at": "2020-09-08T07:33:24Z",
                "updated_at": "2020-09-08T07:33:24Z",
                "revision_number": 0,
                "project_id": "17a1c3c952c84b3e84a82ddd48364938",
                "type": "RDP",
                "cidr": "0.0.0.0/0",
                "port_range": "3389"
            }
        ],
        "outbound": [
            {
                "id": "eb832c28-ba61-4260-acf2-09afdd3c5943",
                "tenant_id": "17a1c3c952c84b3e84a82ddd48364938",
                "security_group_id": "b1a01fdf-492b-48aa-9b7b-399509cbb5e4",
                "ethertype": "IPv4",
                "direction": "egress",
                "protocol": null,
                "port_range_min": null,
                "port_range_max": null,
                "remote_ip_prefix": null,
                "remote_group_id": null,
                "description": "ALL OUT",
                "tags": [],
                "created_at": "2020-09-08T07:33:25Z",
                "updated_at": "2020-09-08T07:33:25Z",
                "revision_number": 0,
                "project_id": "17a1c3c952c84b3e84a82ddd48364938",
                "type": "ALL OUT",
                "cidr": null,
                "port_range": null
            }
        ]
    }
]
`
		_, _ = fmt.Fprint(w, resp)
	})

	fws, err := client.Firewall.List(ctx, &ListOptions{})
	require.NoError(t, err)
	assert.Len(t, fws, 1)
	fw := fws[0]
	assert.Equal(t, "b1a01fdf-492b-48aa-9b7b-399509cbb5e4", fw.ID)
	assert.Equal(t, "sapd-fw-hcm", fw.Name)
	assert.Equal(t, 4, fw.RulesCount)
	assert.Equal(t, 2, len(fw.Servers))
}

func TestFirewallGet(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.CloudServerURL(firewallBasePath+"/b1a01fdf-492b-48aa-9b7b-399509cbb5e4"), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
    "id": "b1a01fdf-492b-48aa-9b7b-399509cbb5e4",
    "name": "sapd-fw-hcm",
    "tenant_id": "17a1c3c952c84b3e84a82ddd48364938",
    "description": "sapd-fw-hcm created 2020-09-08 14:33:24.107718",
    "tags": [],
    "created_at": "2020-09-08T07:33:24Z",
    "updated_at": "2020-09-08T07:33:25Z",
    "revision_number": 7,
    "project_id": "17a1c3c952c84b3e84a82ddd48364938",
    "servers_count": 2,
    "servers": [
        {
            "id": "b6c43097-27b2-4f09-8498-b7f4067909a0",
            "name": "ubuntu-2vcpu-2gb-01",
            "status": "ACTIVE",
            "tenant_id": "17a1c3c952c84b3e84a82ddd48364938",
            "user_id": "55d38aecb1034c06b99c1c87fb6f0740",
            "metadata": {
                "category": "dedicated",
                "os_type": "Ubuntu 20.04 THOR_TRANSFER"
            },
            "hostId": "ca09528b7a51ae54755143e60c57bdca5d68d3475d77f6759f728c9e",
            "flavor": {
                "id": "d470460e-d803-41ff-998c-a013ded4a8c8",
                "name": "2c_2g_dedicated",
                "ram": 2048,
                "disk": 0,
                "swap": "",
                "OS-FLV-EXT-DATA:ephemeral": 0,
                "OS-FLV-DISABLED:disabled": false,
                "vcpus": 2,
                "os-flavor-access:is_public": true,
                "rxtx_factor": 1.0
            },
            "created": "2020-09-05T03:06:10Z",
            "updated": "2020-09-05T03:06:24Z",
            "addresses": {
                "priv_sapd@vccloud.vn": [
                    {
                        "version": 4,
                        "addr": "10.20.1.156",
                        "OS-EXT-IPS:type": "fixed",
                        "OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:9c:76:9f"
                    }
                ],
                "EXT_DIRECTNET_1": [
                    {
                        "version": 4,
                        "addr": "103.153.73.214",
                        "OS-EXT-IPS:type": "fixed",
                        "OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:09:4f:71"
                    }
                ]
            },
            "accessIPv4": "",
            "accessIPv6": "",
            "OS-DCF:diskConfig": "MANUAL",
            "progress": 0,
            "OS-EXT-AZ:availability_zone": "HCM1",
            "config_drive": "",
            "key_name": "sapd1",
            "OS-SRV-USG:launched_at": "2020-09-05T03:05:26.000000",
            "OS-SRV-USG:terminated_at": null,
            "OS-EXT-STS:task_state": null,
            "OS-EXT-STS:vm_state": "active",
            "OS-EXT-STS:power_state": 1,
            "os-extended-volumes:volumes_attached": [
                {
                    "id": "86850fe7-55ad-40dd-adde-1bba13845006"
                }
            ],
            "ip_addresses": {
                "LAN": [
                    {
                        "version": 4,
                        "addr": "10.20.1.156",
                        "OS-EXT-IPS:type": "fixed",
                        "OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:9c:76:9f"
                    }
                ],
                "WAN_V4": [
                    {
                        "version": 4,
                        "addr": "103.153.73.214",
                        "OS-EXT-IPS:type": "fixed",
                        "OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:09:4f:71"
                    }
                ],
                "WAN_V6": []
            },
            "zone_name": "HCM1",
            "region_name": "HoChiMinh",
            "category": "dedicated"
        },
        {
            "id": "7371145f-bed6-4dbd-bf5a-21b48ab1102f",
            "name": "ubuntu-2vcpu-4gb-01",
            "status": "ACTIVE",
            "tenant_id": "17a1c3c952c84b3e84a82ddd48364938",
            "user_id": "55d38aecb1034c06b99c1c87fb6f0740",
            "metadata": {
                "category": "premium",
                "os_type": "Ubuntu 20.04 THOR_TRANSFER"
            },
            "hostId": "787f9ae05e40d07996722c23f3a7c0d6edc827f4ca2288343d5e4e45",
            "flavor": {
                "id": "0c3a8e65-6653-4a99-851d-524f5d22950e",
                "name": "nix.2c_4g",
                "ram": 4096,
                "disk": 0,
                "swap": "",
                "OS-FLV-EXT-DATA:ephemeral": 0,
                "OS-FLV-DISABLED:disabled": false,
                "vcpus": 2,
                "os-flavor-access:is_public": true,
                "rxtx_factor": 1.0
            },
            "created": "2020-09-04T09:27:36Z",
            "updated": "2020-09-04T09:27:53Z",
            "addresses": {
                "priv_sapd@vccloud.vn": [
                    {
                        "version": 4,
                        "addr": "10.20.1.161",
                        "OS-EXT-IPS:type": "fixed",
                        "OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:dd:b3:7d"
                    }
                ],
                "EXT_DIRECTNET_1": [
                    {
                        "version": 4,
                        "addr": "103.153.73.52",
                        "OS-EXT-IPS:type": "fixed",
                        "OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:64:62:65"
                    }
                ]
            },
            "accessIPv4": "",
            "accessIPv6": "",
            "OS-DCF:diskConfig": "MANUAL",
            "progress": 0,
            "OS-EXT-AZ:availability_zone": "HCM1",
            "config_drive": "",
            "key_name": "sapd1",
            "OS-SRV-USG:launched_at": "2020-09-04T09:27:28.000000",
            "OS-SRV-USG:terminated_at": null,
            "OS-EXT-STS:task_state": null,
            "OS-EXT-STS:vm_state": "active",
            "OS-EXT-STS:power_state": 1,
            "os-extended-volumes:volumes_attached": [
                {
                    "id": "e0004420-d783-43e9-a05a-bbc56b56ad18"
                }
            ],
            "ip_addresses": {
                "LAN": [
                    {
                        "version": 4,
                        "addr": "10.20.1.161",
                        "OS-EXT-IPS:type": "fixed",
                        "OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:dd:b3:7d"
                    }
                ],
                "WAN_V4": [
                    {
                        "version": 4,
                        "addr": "103.153.73.52",
                        "OS-EXT-IPS:type": "fixed",
                        "OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:64:62:65"
                    }
                ],
                "WAN_V6": []
            },
            "zone_name": "HCM1",
            "region_name": "HoChiMinh",
            "category": "premium"
        }
    ],
    "rules_count": 4,
    "inbound": [
        {
            "id": "4ad34f53-e82f-40a5-9c79-c21695145f5c",
            "tenant_id": "17a1c3c952c84b3e84a82ddd48364938",
            "security_group_id": "b1a01fdf-492b-48aa-9b7b-399509cbb5e4",
            "ethertype": "IPv4",
            "direction": "ingress",
            "protocol": "tcp",
            "port_range_min": 22,
            "port_range_max": 22,
            "remote_ip_prefix": "0.0.0.0/0",
            "remote_group_id": null,
            "description": "SSH",
            "tags": [],
            "created_at": "2020-09-08T07:33:24Z",
            "updated_at": "2020-09-08T07:33:24Z",
            "revision_number": 0,
            "project_id": "17a1c3c952c84b3e84a82ddd48364938",
            "type": "SSH",
            "cidr": "0.0.0.0/0",
            "port_range": "22"
        },
        {
            "id": "a4782655-d73c-4132-9f0f-032382bb55c9",
            "tenant_id": "17a1c3c952c84b3e84a82ddd48364938",
            "security_group_id": "b1a01fdf-492b-48aa-9b7b-399509cbb5e4",
            "ethertype": "IPv4",
            "direction": "ingress",
            "protocol": "icmp",
            "port_range_min": null,
            "port_range_max": null,
            "remote_ip_prefix": "0.0.0.0/0",
            "remote_group_id": null,
            "description": "PING",
            "tags": [],
            "created_at": "2020-09-08T07:33:24Z",
            "updated_at": "2020-09-08T07:33:24Z",
            "revision_number": 0,
            "project_id": "17a1c3c952c84b3e84a82ddd48364938",
            "type": "PING",
            "cidr": "0.0.0.0/0",
            "port_range": null
        },
        {
            "id": "e12a161e-b008-47d5-b427-45dc1a24bc94",
            "tenant_id": "17a1c3c952c84b3e84a82ddd48364938",
            "security_group_id": "b1a01fdf-492b-48aa-9b7b-399509cbb5e4",
            "ethertype": "IPv4",
            "direction": "ingress",
            "protocol": "tcp",
            "port_range_min": 3389,
            "port_range_max": 3389,
            "remote_ip_prefix": "0.0.0.0/0",
            "remote_group_id": null,
            "description": "RDP",
            "tags": [],
            "created_at": "2020-09-08T07:33:24Z",
            "updated_at": "2020-09-08T07:33:24Z",
            "revision_number": 0,
            "project_id": "17a1c3c952c84b3e84a82ddd48364938",
            "type": "RDP",
            "cidr": "0.0.0.0/0",
            "port_range": "3389"
        }
    ],
    "outbound": [
        {
            "id": "eb832c28-ba61-4260-acf2-09afdd3c5943",
            "tenant_id": "17a1c3c952c84b3e84a82ddd48364938",
            "security_group_id": "b1a01fdf-492b-48aa-9b7b-399509cbb5e4",
            "ethertype": "IPv4",
            "direction": "egress",
            "protocol": null,
            "port_range_min": null,
            "port_range_max": null,
            "remote_ip_prefix": null,
            "remote_group_id": null,
            "description": "ALL OUT",
            "tags": [],
            "created_at": "2020-09-08T07:33:25Z",
            "updated_at": "2020-09-08T07:33:25Z",
            "revision_number": 0,
            "project_id": "17a1c3c952c84b3e84a82ddd48364938",
            "type": "ALL OUT",
            "cidr": null,
            "port_range": null
        }
    ]
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	fw, err := client.Firewall.Get(ctx, "b1a01fdf-492b-48aa-9b7b-399509cbb5e4")
	require.NoError(t, err)
	assert.Equal(t, "b1a01fdf-492b-48aa-9b7b-399509cbb5e4", fw.ID)
	assert.Equal(t, "sapd-fw-hcm", fw.Name)
	assert.Equal(t, 4, fw.RulesCount)
	assert.Equal(t, 2, len(fw.Servers))
}

func TestDeleteFirewall(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.CloudServerURL(firewallBasePath+"/b1a01fdf-492b-48aa-9b7b-399509cbb5e4"), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
		resp := `
{
    "message": "Delete firewall successfully"
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	resp, err := client.Firewall.Delete(ctx, "b1a01fdf-492b-48aa-9b7b-399509cbb5e4")
	require.NoError(t, err)
	assert.Equal(t, "Delete firewall successfully", resp.Message)
}

func TestFirewallCreate(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.CloudServerURL(firewallBasePath), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		resp := `
{
    "id": "48dc460b-3ea7-4cb3-bc5d-1d41f297dcd0",
    "name": "bizflycloud-firewall",
    "tenant_id": "17a1c3c952c84b3e84a82ddd48364938",
    "description": "bizflycloud-firewall created 2020-09-09 12:04:31.414243",
    "tags": [],
    "created_at": "2020-09-09T05:04:31Z",
    "updated_at": "2020-09-09T05:04:32Z",
    "revision_number": 8,
    "project_id": "17a1c3c952c84b3e84a82ddd48364938",
    "servers_count": 0,
    "servers": [
		{
            "id": "6ee4ed07-ba59-4dd9-a0ca-cbc21861ca4b",
            "name": "sapd-tf-server-2",
            "status": "ACTIVE",
            "tenant_id": "17a1c3c952c84b3e84a82ddd48364938",
            "user_id": "55d38aecb1034c06b99c1c87fb6f0740",
            "metadata": {
                "category": "premium",
                "os_type": "Ubuntu 18.04"
            },
            "hostId": "84bf6a43769bd4c041c75c1737c38d4b8df5faf6916f9f603b4702d8",
            "flavor": {
                "id": "f4d23537-8a87-4c32-bb0b-60328e6f4374",
                "name": "nix.4c_2g",
                "ram": 2048,
                "disk": 0,
                "swap": "",
                "OS-FLV-EXT-DATA:ephemeral": 0,
                "OS-FLV-DISABLED:disabled": false,
                "vcpus": 4,
                "os-flavor-access:is_public": true,
                "rxtx_factor": 1.0
            },
            "created": "2020-10-07T04:41:12Z",
            "updated": "2020-10-07T04:41:31Z",
            "addresses": {
                "priv_sapd@vccloud.vn": [
                    {
                        "version": 4,
                        "addr": "10.20.165.89",
                        "OS-EXT-IPS:type": "fixed",
                        "OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:f0:b2:cd"
                    }
                ],
                "EXT_DIRECTNET_11": [
                    {
                        "version": 4,
                        "addr": "123.30.234.210",
                        "OS-EXT-IPS:type": "fixed",
                        "OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:10:22:0a"
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
            "OS-SRV-USG:launched_at": "2020-10-07T04:41:30.000000",
            "OS-SRV-USG:terminated_at": null,
            "OS-EXT-STS:task_state": null,
            "OS-EXT-STS:vm_state": "active",
            "OS-EXT-STS:power_state": 1,
            "os-extended-volumes:volumes_attached": [
                {
                    "id": "364ee097-190d-482a-ba8f-87ef75d07e46"
                }
            ],
            "ip_addresses": {
                "LAN": [
                    {
                        "version": 4,
                        "addr": "10.20.165.89",
                        "OS-EXT-IPS:type": "fixed",
                        "OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:f0:b2:cd"
                    }
                ],
                "WAN_V4": [
                    {
                        "version": 4,
                        "addr": "123.30.234.210",
                        "OS-EXT-IPS:type": "fixed",
                        "OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:10:22:0a"
                    }
                ],
                "WAN_V6": []
            },
            "zone_name": "HN1",
            "region_name": "HaNoi",
            "category": "premium"
        }
	],
    "rules_count": 5,
    "inbound": [
        {
            "id": "25de18c5-eccb-4ce8-b550-5cd612eafe52",
            "tenant_id": "17a1c3c952c84b3e84a82ddd48364938",
            "security_group_id": "48dc460b-3ea7-4cb3-bc5d-1d41f297dcd0",
            "ethertype": "IPv4",
            "direction": "ingress",
            "protocol": "tcp",
            "port_range_min": 22,
            "port_range_max": 22,
            "remote_ip_prefix": "0.0.0.0/0",
            "remote_group_id": null,
            "description": "SSH",
            "tags": [],
            "created_at": "2020-09-09T05:04:31Z",
            "updated_at": "2020-09-09T05:04:31Z",
            "revision_number": 0,
            "project_id": "17a1c3c952c84b3e84a82ddd48364938",
            "type": "SSH",
            "cidr": "0.0.0.0/0",
            "port_range": "22"
        },
        {
            "id": "3a091f59-285c-4ac0-8b16-adfe8c0a7faa",
            "tenant_id": "17a1c3c952c84b3e84a82ddd48364938",
            "security_group_id": "48dc460b-3ea7-4cb3-bc5d-1d41f297dcd0",
            "ethertype": "IPv6",
            "direction": "ingress",
            "protocol": "tcp",
            "port_range_min": 22,
            "port_range_max": 22,
            "remote_ip_prefix": "2001:db8:85a3::8a2e:370:7334/128",
            "remote_group_id": null,
            "description": "SSH",
            "tags": [],
            "created_at": "2020-09-09T05:04:32Z",
            "updated_at": "2020-09-09T05:04:32Z",
            "revision_number": 0,
            "project_id": "17a1c3c952c84b3e84a82ddd48364938",
            "type": "SSH",
            "cidr": "2001:db8:85a3::8a2e:370:7334/128",
            "port_range": "22"
        },
        {
            "id": "c4fcee86-d699-4d6f-8a80-b78093a47959",
            "tenant_id": "17a1c3c952c84b3e84a82ddd48364938",
            "security_group_id": "48dc460b-3ea7-4cb3-bc5d-1d41f297dcd0",
            "ethertype": "IPv4",
            "direction": "ingress",
            "protocol": "tcp",
            "port_range_min": 80,
            "port_range_max": 80,
            "remote_ip_prefix": "192.168.17.5/32",
            "remote_group_id": null,
            "description": "HTTP",
            "tags": [],
            "created_at": "2020-09-09T05:04:32Z",
            "updated_at": "2020-09-09T05:04:32Z",
            "revision_number": 0,
            "project_id": "17a1c3c952c84b3e84a82ddd48364938",
            "type": "HTTP",
            "cidr": "192.168.17.5/32",
            "port_range": "80"
        }
    ],
    "outbound": [
        {
            "id": "03e28458-052a-41b3-98ad-f6d192e11635",
            "tenant_id": "17a1c3c952c84b3e84a82ddd48364938",
            "security_group_id": "48dc460b-3ea7-4cb3-bc5d-1d41f297dcd0",
            "ethertype": "IPv4",
            "direction": "egress",
            "protocol": "tcp",
            "port_range_min": 1,
            "port_range_max": 255,
            "remote_ip_prefix": "192.168.0.0/28",
            "remote_group_id": null,
            "description": "CUSTOM",
            "tags": [],
            "created_at": "2020-09-09T05:04:32Z",
            "updated_at": "2020-09-09T05:04:32Z",
            "revision_number": 0,
            "project_id": "17a1c3c952c84b3e84a82ddd48364938",
            "type": "CUSTOM",
            "cidr": "192.168.0.0/28",
            "port_range": "1-255"
        },
        {
            "id": "b36c8249-5aa3-4d02-95c5-7e1487117d00",
            "tenant_id": "17a1c3c952c84b3e84a82ddd48364938",
            "security_group_id": "48dc460b-3ea7-4cb3-bc5d-1d41f297dcd0",
            "ethertype": "IPv6",
            "direction": "egress",
            "protocol": "ipv6-icmp",
            "port_range_min": null,
            "port_range_max": null,
            "remote_ip_prefix": "::/0",
            "remote_group_id": null,
            "description": "PING",
            "tags": [],
            "created_at": "2020-09-09T05:04:32Z",
            "updated_at": "2020-09-09T05:04:32Z",
            "revision_number": 0,
            "project_id": "17a1c3c952c84b3e84a82ddd48364938",
            "type": "PING",
            "cidr": "::/0",
            "port_range": null
        }
    ]
}
`
		_, _ = fmt.Fprint(w, resp)
	})
	fcr := FirewallRequestPayload{
		Name: "bizflycloud-firewall",
		InBound: []FirewallRuleCreateRequest{
			{
				Type:      "SSH",
				Protocol:  "TCP",
				PortRange: "22",
				CIDR:      "0.0.0.0/0",
			},
			{
				Type:      "HTTP",
				Protocol:  "TCP",
				PortRange: "80",
				CIDR:      "192.168.17.5",
			},
			{
				Type:      "SSH",
				Protocol:  "TCP",
				PortRange: "22",
				CIDR:      "2001:0db8:85a3:0000:0000:8a2e:0370:7334/128",
			},
		},
		OutBound: []FirewallRuleCreateRequest{
			{
				Type:     "PING",
				Protocol: "ICMP",
				CIDR:     "::/0",
			},
			{
				Type:      "CUSTOM",
				Protocol:  "TCP",
				PortRange: "1-255",
				CIDR:      "192.168.0.0/28",
			},
		},
		Targets: []string{
			"6ee4ed07-ba59-4dd9-a0ca-cbc21861ca4b",
		},
	}
	fw, err := client.Firewall.Create(ctx, &fcr)
	require.NoError(t, err)
	assert.Equal(t, "48dc460b-3ea7-4cb3-bc5d-1d41f297dcd0", fw.ID)
	assert.Equal(t, "192.168.0.0/28", fw.OutBound[0].CIDR)
}

func TestUpdateFirewall(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.CloudServerURL(firewallBasePath+"/48dc460b-3ea7-4cb3-bc5d-1d41f297dcd0"), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPatch, r.Method)
		resp := `
{
    "id": "48dc460b-3ea7-4cb3-bc5d-1d41f297dcd0",
    "name": "bizflycloud-firewall",
    "tenant_id": "17a1c3c952c84b3e84a82ddd48364938",
    "description": "bizflycloud-firewall created 2020-09-09 12:04:31.414243",
    "tags": [],
    "created_at": "2020-09-09T05:04:31Z",
    "updated_at": "2020-09-09T06:27:35Z",
    "revision_number": 13,
    "project_id": "17a1c3c952c84b3e84a82ddd48364938",
    "servers_count": 1,
    "servers": [
        {
            "id": "b6c43097-27b2-4f09-8498-b7f4067909a0",
            "name": "ubuntu-2vcpu-2gb-01",
            "status": "ACTIVE",
            "tenant_id": "17a1c3c952c84b3e84a82ddd48364938",
            "user_id": "55d38aecb1034c06b99c1c87fb6f0740",
            "metadata": {
                "category": "dedicated",
                "os_type": "Ubuntu 20.04 THOR_TRANSFER"
            },
            "hostId": "ca09528b7a51ae54755143e60c57bdca5d68d3475d77f6759f728c9e",
            "flavor": {
                "id": "d470460e-d803-41ff-998c-a013ded4a8c8",
                "name": "2c_2g_dedicated",
                "ram": 2048,
                "disk": 0,
                "swap": "",
                "OS-FLV-EXT-DATA:ephemeral": 0,
                "OS-FLV-DISABLED:disabled": false,
                "vcpus": 2,
                "os-flavor-access:is_public": true,
                "rxtx_factor": 1.0
            },
            "created": "2020-09-05T03:06:10Z",
            "updated": "2020-09-05T03:06:24Z",
            "addresses": {
                "priv_sapd@vccloud.vn": [
                    {
                        "version": 4,
                        "addr": "10.20.1.156",
                        "OS-EXT-IPS:type": "fixed",
                        "OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:9c:76:9f"
                    }
                ],
                "EXT_DIRECTNET_1": [
                    {
                        "version": 4,
                        "addr": "103.153.73.214",
                        "OS-EXT-IPS:type": "fixed",
                        "OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:09:4f:71"
                    }
                ]
            },
            "accessIPv4": "",
            "accessIPv6": "",
            "OS-DCF:diskConfig": "MANUAL",
            "progress": 0,
            "OS-EXT-AZ:availability_zone": "HCM1",
            "config_drive": "",
            "key_name": "sapd1",
            "OS-SRV-USG:launched_at": "2020-09-05T03:05:26.000000",
            "OS-SRV-USG:terminated_at": null,
            "OS-EXT-STS:task_state": null,
            "OS-EXT-STS:vm_state": "active",
            "OS-EXT-STS:power_state": 1,
            "os-extended-volumes:volumes_attached": [
                {
                    "id": "86850fe7-55ad-40dd-adde-1bba13845006"
                }
            ],
            "ip_addresses": {
                "LAN": [
                    {
                        "version": 4,
                        "addr": "10.20.1.156",
                        "OS-EXT-IPS:type": "fixed",
                        "OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:9c:76:9f"
                    }
                ],
                "WAN_V4": [
                    {
                        "version": 4,
                        "addr": "103.153.73.214",
                        "OS-EXT-IPS:type": "fixed",
                        "OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:09:4f:71"
                    }
                ],
                "WAN_V6": []
            },
            "zone_name": "HCM1",
            "region_name": "HoChiMinh",
            "category": "dedicated"
        }
    ],
    "rules_count": 0,
    "inbound": [],
    "outbound": []
}
`
		_, _ = fmt.Fprint(w, resp)
	})
	fur := FirewallRequestPayload{
		Targets:  []string{"b6c43097-27b2-4f09-8498-b7f4067909a0"},
		InBound:  []FirewallRuleCreateRequest{},
		OutBound: []FirewallRuleCreateRequest{},
	}
	fw, err := client.Firewall.Update(ctx, "48dc460b-3ea7-4cb3-bc5d-1d41f297dcd0", &fur)
	require.NoError(t, err)
	assert.Equal(t, "48dc460b-3ea7-4cb3-bc5d-1d41f297dcd0", fw.ID)
	assert.Equal(t, "b6c43097-27b2-4f09-8498-b7f4067909a0", fw.Servers[0].ID)
}

func TestRemoveServerFirewall(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.CloudServerURL(firewallBasePath+"/48dc460b-3ea7-4cb3-bc5d-1d41f297dcd0/servers"), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
		resp := `
{
    "id": "48dc460b-3ea7-4cb3-bc5d-1d41f297dcd0",
    "name": "bizflycloud-firewall",
    "tenant_id": "17a1c3c952c84b3e84a82ddd48364938",
    "description": "bizflycloud-firewall created 2020-09-09 12:04:31.414243",
    "tags": [],
    "created_at": "2020-09-09T05:04:31Z",
    "updated_at": "2020-09-09T06:27:35Z",
    "revision_number": 13,
    "project_id": "17a1c3c952c84b3e84a82ddd48364938",
    "servers_count": 0,
    "servers": [],
    "rules_count": 0,
    "inbound": [],
    "outbound": []
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	frsr := FirewallRemoveServerRequest{
		Servers: []string{"b6c43097-27b2-4f09-8498-b7f4067909a0"},
	}
	fw, err := client.Firewall.RemoveServer(ctx, "48dc460b-3ea7-4cb3-bc5d-1d41f297dcd0", &frsr)
	require.NoError(t, err)
	assert.Len(t, fw.Servers, 0)
	assert.Equal(t, "48dc460b-3ea7-4cb3-bc5d-1d41f297dcd0", fw.ID)
}

func TestDeleteFirewallRule(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.CloudServerURL(firewallBasePath+"/48dc460b-3ea7-4cb3-bc5d-1d41f297dcd0"), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
		resp := `
{
	"message": "Deleted Firewall Rule"
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	resp, err := client.Firewall.DeleteRule(ctx, "48dc460b-3ea7-4cb3-bc5d-1d41f297dcd0")
	require.NoError(t, err)
	assert.Equal(t, "Deleted Firewall Rule", resp.Message)
}

func TestCreateFirewallRule(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.CloudServerURL(firewallBasePath+"/f09bfc4b-92a9-468a-b41d-6ba8d4bd7552/rules"), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		resp := `
{
    "security_group_rule": {
        "id": "45c37e35-86ff-4fc6-9285-0bb9f5110db8",
        "tenant_id": "a7d1e56edcac40d0896d2b97f414afc5",
        "security_group_id": "f09bfc4b-92a9-468a-b41d-6ba8d4bd7552",
        "ethertype": "IPv4",
        "direction": "ingress",
        "protocol": "udp",
        "port_range_min": 80,
        "port_range_max": 90,
        "remote_ip_prefix": "0.0.0.0/0",
        "remote_group_id": null,
        "description": "CUSTOM",
        "created_at": "2020-09-11T07:27:21Z",
        "updated_at": "2020-09-11T07:27:21Z",
        "revision_number": 0,
        "project_id": "a7d1e56edcac40d0896d2b97f414afc5"
    }
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	var fsrcr = FirewallSingleRuleCreateRequest{
		Direction: "ingress",
		FirewallRuleCreateRequest: FirewallRuleCreateRequest{
			CIDR:      "0.0.0.0/0",
			Protocol:  "UDP",
			Type:      "CUSTOM",
			PortRange: "80-90",
		},
	}
	resp, err := client.Firewall.CreateRule(ctx, "f09bfc4b-92a9-468a-b41d-6ba8d4bd7552", &fsrcr)
	require.NoError(t, err)
	assert.Equal(t, "45c37e35-86ff-4fc6-9285-0bb9f5110db8", resp.ID)
	assert.Equal(t, "ingress", resp.Direction)
	assert.Equal(t, "IPv4", resp.EtherType)
	assert.Equal(t, "udp", resp.Protocol)
}
