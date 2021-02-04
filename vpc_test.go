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
	"github.com/bizflycloud/gobizfly/testlib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestVPCList(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.CloudServerURL(vpcPath), func(writer http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		resp := `[
    {
        "id": "41f40298-8d05-4889-9672-f03cfccd719f",
        "name": "vxsCFZ",
        "tenant_id": "bc8d2790fc9a46949818b942c0a824de",
        "admin_state_up": true,
        "mtu": 1500,
        "status": "ACTIVE",
        "subnets": [
            {
                "id": "d05bbc5d-8933-4dcf-af98-f0d20179f471",
                "name": "vxsCFZ",
                "tenant_id": "bc8d2790fc9a46949818b942c0a824de",
                "network_id": "41f40298-8d05-4889-9672-f03cfccd719f",
                "ip_version": 4,
                "subnetpool_id": null,
                "enable_dhcp": true,
                "ipv6_ra_mode": null,
                "ipv6_address_mode": null,
                "gateway_ip": null,
                "cidr": "10.21.148.0/24",
                "allocation_pools": [
                    {
                        "start": "10.21.148.1",
                        "end": "10.21.148.254"
                    }
                ],
                "host_routes": [],
                "dns_nameservers": [
                    "103.92.35.110",
                    "208.67.222.222"
                ],
                "description": "",
                "service_types": [],
                "tags": [],
                "created_at": "2021-02-03T06:59:41Z",
                "updated_at": "2021-02-03T06:59:41Z",
                "revision_number": 0,
                "project_id": "bc8d2790fc9a46949818b942c0a824de"
            }
        ],
        "shared": false,
        "availability_zone_hints": [
            "HN1",
            "HN2"
        ],
        "availability_zones": [
            "HN1"
        ],
        "ipv4_address_scope": null,
        "ipv6_address_scope": null,
        "router:external": false,
        "description": "",
        "port_security_enabled": true,
        "qos_policy_id": "ab305294-df0d-4f22-88d7-3b7f06167bb0",
        "tags": [],
        "created_at": "2021-02-03T06:59:35Z",
        "updated_at": "2021-02-03T06:59:41Z",
        "revision_number": 2,
        "project_id": "bc8d2790fc9a46949818b942c0a824de",
        "is_default": false
    },
    {
        "id": "b5b52090-3ca0-4954-8300-a0ea998f5ab7",
        "name": "Airflow",
        "tenant_id": "bc8d2790fc9a46949818b942c0a824de",
        "admin_state_up": true,
        "mtu": 1500,
        "status": "ACTIVE",
        "subnets": [
            {
                "id": "a62ae59e-8efa-4d46-9d37-a87bef6de3f0",
                "name": "Airflow",
                "tenant_id": "bc8d2790fc9a46949818b942c0a824de",
                "network_id": "b5b52090-3ca0-4954-8300-a0ea998f5ab7",
                "ip_version": 4,
                "subnetpool_id": null,
                "enable_dhcp": true,
                "ipv6_ra_mode": null,
                "ipv6_address_mode": null,
                "gateway_ip": null,
                "cidr": "10.26.86.0/24",
                "allocation_pools": [
                    {
                        "start": "10.26.86.1",
                        "end": "10.26.86.254"
                    }
                ],
                "host_routes": [],
                "dns_nameservers": [
                    "103.92.35.110",
                    "208.67.222.222"
                ],
                "description": "",
                "service_types": [],
                "tags": [],
                "created_at": "2021-01-20T14:28:24Z",
                "updated_at": "2021-01-20T14:28:24Z",
                "revision_number": 0,
                "project_id": "bc8d2790fc9a46949818b942c0a824de"
            }
        ],
        "shared": false,
        "availability_zone_hints": [
            "HN1",
            "HN2"
        ],
        "availability_zones": [
            "HN1"
        ],
        "ipv4_address_scope": null,
        "ipv6_address_scope": null,
        "router:external": false,
        "description": "",
        "port_security_enabled": true,
        "qos_policy_id": "ab305294-df0d-4f22-88d7-3b7f06167bb0",
        "tags": [],
        "created_at": "2021-01-20T14:28:19Z",
        "updated_at": "2021-01-20T14:28:24Z",
        "revision_number": 2,
        "project_id": "bc8d2790fc9a46949818b942c0a824de",
        "is_default": false
    },
    {
        "id": "bc65f5ae-59a5-4405-80a5-a23aa08194bb",
        "name": "priv_svtt.tungds@vccloud.vn",
        "tenant_id": "bc8d2790fc9a46949818b942c0a824de",
        "admin_state_up": true,
        "mtu": 1500,
        "status": "ACTIVE",
        "subnets": [
            {
                "id": "75532b4a-6f84-414f-aa64-bf635d449589",
                "name": "priv_subnet_svtt.tungds@vccloud.vn",
                "tenant_id": "bc8d2790fc9a46949818b942c0a824de",
                "network_id": "bc65f5ae-59a5-4405-80a5-a23aa08194bb",
                "ip_version": 4,
                "subnetpool_id": null,
                "enable_dhcp": true,
                "ipv6_ra_mode": null,
                "ipv6_address_mode": null,
                "gateway_ip": null,
                "cidr": "10.26.118.0/24",
                "allocation_pools": [
                    {
                        "start": "10.26.118.1",
                        "end": "10.26.118.254"
                    }
                ],
                "host_routes": [],
                "dns_nameservers": [
                    "123.31.11.89",
                    "208.67.222.222"
                ],
                "description": "",
                "service_types": [],
                "tags": [],
                "created_at": "2020-10-05T03:08:04Z",
                "updated_at": "2020-10-05T03:08:04Z",
                "revision_number": 0,
                "project_id": "bc8d2790fc9a46949818b942c0a824de"
            }
        ],
        "shared": false,
        "availability_zone_hints": [
            "HN1",
            "HN2"
        ],
        "availability_zones": [
            "HN1"
        ],
        "ipv4_address_scope": null,
        "ipv6_address_scope": null,
        "router:external": false,
        "description": "",
        "port_security_enabled": true,
        "qos_policy_id": "ab305294-df0d-4f22-88d7-3b7f06167bb0",
        "tags": [],
        "created_at": "2020-10-05T03:07:58Z",
        "updated_at": "2021-02-03T06:53:07Z",
        "revision_number": 6,
        "project_id": "bc8d2790fc9a46949818b942c0a824de",
        "is_default": false
    },
    {
        "id": "d9b58855-f44e-4958-9e18-624efe0ebc07",
        "name": "kdsfasfd",
        "tenant_id": "bc8d2790fc9a46949818b942c0a824de",
        "admin_state_up": true,
        "mtu": 1500,
        "status": "ACTIVE",
        "subnets": [
            {
                "id": "bfb5cbf3-c26e-406e-a3a4-94735bf5ff8d",
                "name": "kdsfasfd",
                "tenant_id": "bc8d2790fc9a46949818b942c0a824de",
                "network_id": "d9b58855-f44e-4958-9e18-624efe0ebc07",
                "ip_version": 4,
                "subnetpool_id": null,
                "enable_dhcp": true,
                "ipv6_ra_mode": null,
                "ipv6_address_mode": null,
                "gateway_ip": null,
                "cidr": "10.26.200.0/24",
                "allocation_pools": [
                    {
                        "start": "10.26.200.1",
                        "end": "10.26.200.254"
                    }
                ],
                "host_routes": [],
                "dns_nameservers": [
                    "103.92.35.110",
                    "208.67.222.222"
                ],
                "description": "",
                "service_types": [],
                "tags": [],
                "created_at": "2021-02-02T04:33:06Z",
                "updated_at": "2021-02-02T04:33:06Z",
                "revision_number": 0,
                "project_id": "bc8d2790fc9a46949818b942c0a824de"
            }
        ],
        "shared": false,
        "availability_zone_hints": [
            "HN1",
            "HN2"
        ],
        "availability_zones": [
            "HN1"
        ],
        "ipv4_address_scope": null,
        "ipv6_address_scope": null,
        "router:external": false,
        "description": "",
        "port_security_enabled": true,
        "qos_policy_id": "ab305294-df0d-4f22-88d7-3b7f06167bb0",
        "tags": [],
        "created_at": "2021-02-02T04:33:01Z",
        "updated_at": "2021-02-02T04:33:06Z",
        "revision_number": 2,
        "project_id": "bc8d2790fc9a46949818b942c0a824de",
        "is_default": false
    }
]
`
		_, _ = fmt.Fprint(writer, resp)
	})
	vpcs, err := client.VPC.List(ctx)
	require.NoError(t, err)
	assert.Len(t, vpcs, 4)
	assert.Equal(t, 1500, vpcs[0].MTU)
}

func TestVPCCreate(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.CloudServerURL(vpcPath), func(writer http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		var payload *CreateVPCPayload
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		assert.Equal(t, "test", payload.Name)
		resp := `{
    "id": "41f40298-8d05-4889-9672-f03cfccd719f",
    "name": "test",
    "tenant_id": "bc8d2790fc9a46949818b942c0a824de",
    "admin_state_up": true,
    "mtu": 1500,
    "status": "ACTIVE",
    "subnets": [],
    "shared": false,
    "project_id": "bc8d2790fc9a46949818b942c0a824de",
    "port_security_enabled": true,
    "qos_policy_id": "ab305294-df0d-4f22-88d7-3b7f06167bb0",
    "router:external": false,
    "provider:network_type": "gre",
    "provider:physical_network": null,
    "provider:segmentation_id": 2393,
    "availability_zone_hints": [
        "HN1",
        "HN2"
    ],
    "is_default": true,
    "availability_zones": [],
    "ipv4_address_scope": null,
    "ipv6_address_scope": null,
    "description": "",
    "tags": [],
    "created_at": "2021-02-03T06:59:35Z",
    "updated_at": "2021-02-03T06:59:36Z",
    "revision_number": 1
}
`
		_, _ = fmt.Fprint(writer, resp)
	})
	vpc, err := client.VPC.Create(ctx, &CreateVPCPayload{
		Name: "test",
	})
	require.NoError(t, err)
	assert.Equal(t, "test", vpc.Name)
	assert.Equal(t, true, vpc.IsDefault)
}

func TestVPCGet(t *testing.T) {
	setup()
	defer teardown()
	var v vpcService
	mux.HandleFunc(testlib.CloudServerURL(v.itemPath("41f40298-8d05-4889-9672-f03cfccd719f")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			resp := `
{
    "network": {
        "id": "41f40298-8d05-4889-9672-f03cfccd719f",
        "name": "test1",
        "tenant_id": "bc8d2790fc9a46949818b942c0a824de",
        "admin_state_up": true,
        "mtu": 1500,
        "status": "ACTIVE",
        "subnets": [
            {
                "id": "d05bbc5d-8933-4dcf-af98-f0d20179f471",
                "name": "vxsCFZ",
                "tenant_id": "bc8d2790fc9a46949818b942c0a824de",
                "network_id": "41f40298-8d05-4889-9672-f03cfccd719f",
                "ip_version": 4,
                "subnetpool_id": null,
                "enable_dhcp": true,
                "ipv6_ra_mode": null,
                "ipv6_address_mode": null,
                "gateway_ip": null,
                "cidr": "10.21.148.0/24",
                "allocation_pools": [
                    {
                        "start": "10.21.148.1",
                        "end": "10.21.148.254"
                    }
                ],
                "host_routes": [],
                "dns_nameservers": [
                    "103.92.35.110",
                    "208.67.222.222"
                ],
                "description": "",
                "service_types": [],
                "tags": [],
                "created_at": "2021-02-03T06:59:41Z",
                "updated_at": "2021-02-03T06:59:41Z",
                "revision_number": 0,
                "project_id": "bc8d2790fc9a46949818b942c0a824de"
            }
        ],
        "shared": false,
        "availability_zone_hints": [
            "HN1",
            "HN2"
        ],
        "availability_zones": [
            "HN1"
        ],
        "ipv4_address_scope": null,
        "ipv6_address_scope": null,
        "router:external": false,
        "description": "",
        "port_security_enabled": true,
        "qos_policy_id": "ab305294-df0d-4f22-88d7-3b7f06167bb0",
        "tags": [],
        "created_at": "2021-02-03T06:59:35Z",
        "updated_at": "2021-02-03T07:28:06Z",
        "revision_number": 4,
        "project_id": "bc8d2790fc9a46949818b942c0a824de",
        "ip_availability": [
            {
                "subnet_id": "d05bbc5d-8933-4dcf-af98-f0d20179f471",
                "ip_version": 4,
                "cidr": "10.21.148.0/24",
                "subnet_name": "vxsCFZ",
                "used_ips": 2,
                "total_ips": 254
            }
        ],
        "ports": [],
        "is_default": false
    }
}
`
			_, _ = fmt.Fprint(writer, resp)
		})
	vpc, err := client.VPC.Get(ctx, "41f40298-8d05-4889-9672-f03cfccd719f")
	require.NoError(t, err)
	assert.Equal(t, "2021-02-03T06:59:35Z", vpc.CreatedAt)
}

func TestVPCUpdate(t *testing.T) {
	setup()
	defer teardown()
	var v vpcService
	mux.HandleFunc(testlib.CloudServerURL(v.itemPath("41f40298-8d05-4889-9672-f03cfccd719f")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPut, r.Method)
			var payload *UpdateVPCPayload
			require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
			assert.Equal(t, "test_update", payload.Name)
			resp := `{
    "network": {
        "id": "41f40298-8d05-4889-9672-f03cfccd719f",
        "name": "test1",
        "tenant_id": "bc8d2790fc9a46949818b942c0a824de",
        "admin_state_up": true,
        "mtu": 1500,
        "status": "ACTIVE",
        "subnets": [
            {
                "id": "d05bbc5d-8933-4dcf-af98-f0d20179f471",
                "name": "vxsCFZ",
                "tenant_id": "bc8d2790fc9a46949818b942c0a824de",
                "network_id": "41f40298-8d05-4889-9672-f03cfccd719f",
                "ip_version": 4,
                "subnetpool_id": null,
                "enable_dhcp": true,
                "ipv6_ra_mode": null,
                "ipv6_address_mode": null,
                "gateway_ip": null,
                "cidr": "10.21.148.0/24",
                "allocation_pools": [
                    {
                        "start": "10.21.148.1",
                        "end": "10.21.148.254"
                    }
                ],
                "host_routes": [],
                "dns_nameservers": [
                    "103.92.35.110",
                    "208.67.222.222"
                ],
                "description": "",
                "service_types": [],
                "tags": [],
                "created_at": "2021-02-03T06:59:41Z",
                "updated_at": "2021-02-03T06:59:41Z",
                "revision_number": 0,
                "project_id": "bc8d2790fc9a46949818b942c0a824de"
            }
        ],
        "shared": false,
        "availability_zone_hints": [
            "HN1",
            "HN2"
        ],
        "availability_zones": [
            "HN1"
        ],
        "ipv4_address_scope": null,
        "ipv6_address_scope": null,
        "router:external": false,
        "description": "",
        "port_security_enabled": true,
        "qos_policy_id": "ab305294-df0d-4f22-88d7-3b7f06167bb0",
        "tags": [],
        "created_at": "2021-02-03T06:59:35Z",
        "updated_at": "2021-02-03T07:28:06Z",
        "revision_number": 4,
        "project_id": "bc8d2790fc9a46949818b942c0a824de",
        "ip_availability": [
            {
                "subnet_id": "d05bbc5d-8933-4dcf-af98-f0d20179f471",
                "ip_version": 4,
                "cidr": "10.21.148.0/24",
                "subnet_name": "vxsCFZ",
                "used_ips": 2,
                "total_ips": 254
            }
        ],
        "ports": [],
        "is_default": false
    }
}
`
			_, _ = fmt.Fprint(writer, resp)
		})
	vpc, err := client.VPC.Update(ctx, "41f40298-8d05-4889-9672-f03cfccd719f", &UpdateVPCPayload{
		Name:      "test_update",
		IsDefault: false,
	})
	require.NoError(t, err)
	assert.Equal(t, false, vpc.IsDefault)
}

func TestVPCDelete(t *testing.T) {
	setup()
	defer teardown()
	var v vpcService
	mux.HandleFunc(testlib.CloudServerURL(v.itemPath("0e03c7c5-267b-41f9-baa7-c4d2f2283d50")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodDelete, r.Method)
		})
	require.NoError(t, client.VPC.Delete(ctx, "0e03c7c5-267b-41f9-baa7-c4d2f2283d50"))
}
