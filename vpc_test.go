// This file is part of gobizfly

package gobizfly

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/bizflycloud/gobizfly/testlib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVPCList(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.CloudServerURL(vpcPath), func(writer http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		resp := `[
    {
        "id": "0e03c7c5-267b-41f9-baa7-c4d2f2283d50",
        "name": "asdfasdf",
        "tenant_id": "ebbed256d9414b0598719c42dc17e837",
        "admin_state_up": true,
        "mtu": 1500,
        "status": "ACTIVE",
        "subnets": [
            {
                "id": "cf24149a-8ba9-4445-84a9-99b27258cf23",
                "name": "asdfasdf",
                "tenant_id": "ebbed256d9414b0598719c42dc17e837",
                "network_id": "0e03c7c5-267b-41f9-baa7-c4d2f2283d50",
                "ip_version": 4,
                "subnetpool_id": null,
                "enable_dhcp": true,
                "ipv6_ra_mode": null,
                "ipv6_address_mode": null,
                "gateway_ip": null,
                "cidr": "10.108.16.0/20",
                "allocation_pools": [
                    {
                        "start": "10.108.16.1",
                        "end": "10.108.31.254"
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
                "created_at": "2021-01-28T02:50:50Z",
                "updated_at": "2021-01-28T02:50:50Z",
                "revision_number": 0,
                "project_id": "ebbed256d9414b0598719c42dc17e837"
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
        "description": "asdf",
        "port_security_enabled": true,
        "qos_policy_id": "ab305294-df0d-4f22-88d7-3b7f06167bb0",
        "tags": [
            "default-vpc-network"
        ],
        "created_at": "2021-01-28T02:50:49Z",
        "updated_at": "2021-01-28T02:56:35Z",
        "revision_number": 5,
        "project_id": "ebbed256d9414b0598719c42dc17e837",
        "is_default": true
    },
    {
        "id": "54d26e73-89ea-4e46-ab82-c14d4692d8b2",
        "name": "Airflow",
        "tenant_id": "ebbed256d9414b0598719c42dc17e837",
        "admin_state_up": true,
        "mtu": 1500,
        "status": "ACTIVE",
        "subnets": [
            {
                "id": "b4af9e10-e1cf-4ba1-bf11-98ce46440607",
                "name": "Airflow",
                "tenant_id": "ebbed256d9414b0598719c42dc17e837",
                "network_id": "54d26e73-89ea-4e46-ab82-c14d4692d8b2",
                "ip_version": 4,
                "subnetpool_id": null,
                "enable_dhcp": true,
                "ipv6_ra_mode": null,
                "ipv6_address_mode": null,
                "gateway_ip": null,
                "cidr": "10.23.237.0/24",
                "allocation_pools": [
                    {
                        "start": "10.23.237.1",
                        "end": "10.23.237.254"
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
                "created_at": "2021-01-20T09:07:29Z",
                "updated_at": "2021-01-20T09:07:29Z",
                "revision_number": 0,
                "project_id": "ebbed256d9414b0598719c42dc17e837"
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
        "created_at": "2021-01-20T09:07:24Z",
        "updated_at": "2021-01-28T02:23:41Z",
        "revision_number": 5,
        "project_id": "ebbed256d9414b0598719c42dc17e837",
        "is_default": false
    },
    {
        "id": "7262bf39-14b1-4f06-aae6-2e62944bb124",
        "name": "test_vpc",
        "tenant_id": "ebbed256d9414b0598719c42dc17e837",
        "admin_state_up": true,
        "mtu": 1500,
        "status": "ACTIVE",
        "subnets": [
            {
                "id": "1bedd484-0ef0-4097-aaee-a07163b0cfd5",
                "name": "test_vpc",
                "tenant_id": "ebbed256d9414b0598719c42dc17e837",
                "network_id": "7262bf39-14b1-4f06-aae6-2e62944bb124",
                "ip_version": 4,
                "subnetpool_id": null,
                "enable_dhcp": true,
                "ipv6_ra_mode": null,
                "ipv6_address_mode": null,
                "gateway_ip": null,
                "cidr": "10.108.16.0/20",
                "allocation_pools": [
                    {
                        "start": "10.108.16.1",
                        "end": "10.108.31.254"
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
                "created_at": "2021-01-28T02:32:12Z",
                "updated_at": "2021-01-28T02:32:12Z",
                "revision_number": 0,
                "project_id": "ebbed256d9414b0598719c42dc17e837"
            }
        ],
        "shared": false,
        "availability_zone_hints": [
            "HN1",
            "HN2"
        ],
        "availability_zones": [
            "HN2"
        ],
        "ipv4_address_scope": null,
        "ipv6_address_scope": null,
        "router:external": false,
        "description": "asdf",
        "port_security_enabled": true,
        "qos_policy_id": "ab305294-df0d-4f22-88d7-3b7f06167bb0",
        "tags": [],
        "created_at": "2021-01-28T02:32:11Z",
        "updated_at": "2021-01-28T02:32:12Z",
        "revision_number": 2,
        "project_id": "ebbed256d9414b0598719c42dc17e837",
        "is_default": false
    },
    {
        "id": "79423750-628d-481b-b3bd-c34065c40585",
        "name": "priv_vctest_devcs_tung491@vccloud.vn",
        "tenant_id": "ebbed256d9414b0598719c42dc17e837",
        "admin_state_up": true,
        "mtu": 1500,
        "status": "ACTIVE",
        "subnets": [
            {
                "id": "c908f63c-ede7-4eb7-90f3-f8b3d2c976ec",
                "name": "priv_subnet_vctest_devcs_tung491@vccloud.vn",
                "tenant_id": "ebbed256d9414b0598719c42dc17e837",
                "network_id": "79423750-628d-481b-b3bd-c34065c40585",
                "ip_version": 4,
                "subnetpool_id": null,
                "enable_dhcp": true,
                "ipv6_ra_mode": null,
                "ipv6_address_mode": null,
                "gateway_ip": null,
                "cidr": "10.26.53.0/24",
                "allocation_pools": [
                    {
                        "start": "10.26.53.1",
                        "end": "10.26.53.254"
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
                "created_at": "2021-01-09T03:10:46Z",
                "updated_at": "2021-01-09T03:10:46Z",
                "revision_number": 0,
                "project_id": "ebbed256d9414b0598719c42dc17e837"
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
        "tags": [
            "default-vpc-network"
        ],
        "created_at": "2021-01-09T03:10:38Z",
        "updated_at": "2021-01-29T02:34:09Z",
        "revision_number": 8,
        "project_id": "ebbed256d9414b0598719c42dc17e837",
        "is_default": true
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
		resp := `
{
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
	mux.HandleFunc(testlib.CloudServerURL(v.itemPath("0e03c7c5-267b-41f9-baa7-c4d2f2283d50")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			resp := `
{
        "id": "0e03c7c5-267b-41f9-baa7-c4d2f2283d50",
        "name": "asdfasdf",
        "tenant_id": "ebbed256d9414b0598719c42dc17e837",
        "admin_state_up": true,
        "mtu": 1500,
        "status": "ACTIVE",
        "subnets": [
            {
                "id": "cf24149a-8ba9-4445-84a9-99b27258cf23",
                "name": "asdfasdf",
                "tenant_id": "ebbed256d9414b0598719c42dc17e837",
                "network_id": "0e03c7c5-267b-41f9-baa7-c4d2f2283d50",
                "ip_version": 4,
                "subnetpool_id": null,
                "enable_dhcp": true,
                "ipv6_ra_mode": null,
                "ipv6_address_mode": null,
                "gateway_ip": null,
                "cidr": "10.108.16.0/20",
                "allocation_pools": [
                    {
                        "start": "10.108.16.1",
                        "end": "10.108.31.254"
                    }
                ],
                "host_routes": [
        {
            "destination": "192.168.211.110/32",
            "nexthop": "10.241.10.129"
        }
],
                "dns_nameservers": [
                    "103.92.35.110",
                    "208.67.222.222"
                ],
                "description": "",
                "service_types": [],
                "tags": [],
                "created_at": "2021-01-28T02:50:50Z",
                "updated_at": "2021-01-28T02:50:50Z",
                "revision_number": 0,
                "project_id": "ebbed256d9414b0598719c42dc17e837"
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
        "description": "asdf",
        "port_security_enabled": true,
        "qos_policy_id": "ab305294-df0d-4f22-88d7-3b7f06167bb0",
        "tags": [
            "default-vpc-network"
        ],
        "created_at": "2021-01-28T02:50:49Z",
        "updated_at": "2021-01-28T02:56:35Z",
        "revision_number": 5,
        "project_id": "ebbed256d9414b0598719c42dc17e837",
        "ip_availability": [
            {
                "subnet_id": "cf24149a-8ba9-4445-84a9-99b27258cf23",
                "ip_version": 4,
                "cidr": "10.108.16.0/20",
                "subnet_name": "asdfasdf",
                "used_ips": 2,
                "total_ips": 4094
            }
        ],
        "ports": [],
        "is_default": true
}
`
			_, _ = fmt.Fprint(writer, resp)
		})
	vpc, err := client.VPC.Get(ctx, "0e03c7c5-267b-41f9-baa7-c4d2f2283d50")
	require.NoError(t, err)
	assert.Equal(t, "2021-01-28T02:50:49Z", vpc.CreatedAt)
	assert.Equal(t, "10.241.10.129", vpc.Subnets[0].HostRoutes[0].NextHop)
}

func TestVPCUpdate(t *testing.T) {
	setup()
	defer teardown()
	var v vpcService
	mux.HandleFunc(testlib.CloudServerURL(v.itemPath("0e03c7c5-267b-41f9-baa7-c4d2f2283d50")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPut, r.Method)
			var payload *UpdateVPCPayload
			require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
			assert.Equal(t, "test_update", payload.Name)
			resp := `{
        "id": "0e03c7c5-267b-41f9-baa7-c4d2f2283d50",
        "name": "test_update",
        "tenant_id": "ebbed256d9414b0598719c42dc17e837",
        "admin_state_up": true,
        "mtu": 1500,
        "status": "ACTIVE",
        "subnets": [
            {
                "id": "cf24149a-8ba9-4445-84a9-99b27258cf23",
                "name": "asdfasdf",
                "tenant_id": "ebbed256d9414b0598719c42dc17e837",
                "network_id": "0e03c7c5-267b-41f9-baa7-c4d2f2283d50",
                "ip_version": 4,
                "subnetpool_id": null,
                "enable_dhcp": true,
                "ipv6_ra_mode": null,
                "ipv6_address_mode": null,
                "gateway_ip": null,
                "cidr": "10.108.16.0/20",
                "allocation_pools": [
                    {
                        "start": "10.108.16.1",
                        "end": "10.108.31.254"
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
                "created_at": "2021-01-28T02:50:50Z",
                "updated_at": "2021-01-28T02:50:50Z",
                "revision_number": 0,
                "project_id": "ebbed256d9414b0598719c42dc17e837"
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
        "description": "asdf",
        "port_security_enabled": true,
        "qos_policy_id": "ab305294-df0d-4f22-88d7-3b7f06167bb0",
        "tags": [
            "default-vpc-network"
        ],
        "created_at": "2021-01-28T02:50:49Z",
        "updated_at": "2021-01-28T02:56:35Z",
        "revision_number": 5,
        "project_id": "ebbed256d9414b0598719c42dc17e837",
        "ip_availability": [
            {
                "subnet_id": "cf24149a-8ba9-4445-84a9-99b27258cf23",
                "ip_version": 4,
                "cidr": "10.108.16.0/20",
                "subnet_name": "asdfasdf",
                "used_ips": 2,
                "total_ips": 4094
            }
        ],
        "ports": [],
        "is_default": false
}
`
			_, _ = fmt.Fprint(writer, resp)
		})
	vpc, err := client.VPC.Update(ctx, "0e03c7c5-267b-41f9-baa7-c4d2f2283d50", &UpdateVPCPayload{
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
