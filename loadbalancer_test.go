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

func TestLoadBalancerList(t *testing.T) {
	setup()
	defer teardown()

	var l cloudLoadBalancerService
	mux.HandleFunc(testlib.LoadBalancerURL(l.resourcePath()), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
    "loadbalancers": [
        {
            "network_type": "external",
            "vip_network_id": "9f36fce7-e2c5-44aa-824b-b83c2dca47f6",
            "flavor_id": "",
            "updated_at": "2018-09-18T03:45:30",
            "name": "sapd-test",
            "type": "small",
            "provider": "amphora",
            "id": "ae8e2072-31fb-464a-8285-bc2f2a6bab4d",
            "vip_qos_policy_id": "94c75cb1-ffe9-4dba-8f37-a375fc10462d",
            "tenant_id": "1e7f10a9850b45b488a3f0417ccb60e0",
            "provisioning_status": "ACTIVE",
            "vip_port_id": "59b5004b-baa7-463d-bab8-409883ce2458",
            "created_at": "2018-09-18T03:43:29",
            "listeners": [
                {
                    "id": "5482c4a4-f822-46d0-9af3-026f7579d653"
                }
            ],
            "vip_address": "103.56.156.127",
            "pools": [
                {
                    "id": "1fb271b2-a77e-4afc-8ec6-c6bc110f4c75"
                }
            ],
            "project_id": "1e7f10a9850b45b488a3f0417ccb60e0",
            "admin_state_up": true,
            "description": "",
            "vip_subnet_id": "bbad9d0a-09ee-4053-a4f8-9cd8e7ea5e86",
            "operating_status": "ONLINE"
        }
    ],
    "loadbalancers_links": []
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	lbs, err := client.CloudLoadBalancer.List(ctx, &ListOptions{})
	require.NoError(t, err)
	assert.Len(t, lbs, 1)
	lb := lbs[0]
	assert.Equal(t, "ae8e2072-31fb-464a-8285-bc2f2a6bab4d", lb.ID)
}

func TestLoadBalancerCreate(t *testing.T) {
	setup()
	defer teardown()

	var l cloudLoadBalancerService
	mux.HandleFunc(testlib.LoadBalancerURL(l.resourcePath()), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		var payload struct {
			LoadBalancer *LoadBalancerCreateRequest `json:"loadbalancer"`
		}
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		assert.Equal(t, "Test Create LB", payload.LoadBalancer.Description)
		assert.Equal(t, "LB", payload.LoadBalancer.Name)

		resp := `
{
    "loadbalancer": {
        "updated_at": null,
        "vip_address": "45.124.94.216",
        "provisioning_status": "PENDING_CREATE",
        "vip_network_id": "180784e0-045d-40bb-adec-fdc3e9d3a32e",
        "vip_port_id": "7ef6fac8-1a0a-4255-b21c-03d36b1def73",
        "id": "e389f5eb-07b5-486b-be4d-4d4d1299f0ab",
        "admin_state_up": true,
        "listeners": [],
        "pools": [],
        "vip_qos_policy_id": "3b70c2d2-5a1f-44e9-9d2f-12aaa2369228",
        "operating_status": "OFFLINE",
        "flavor_id": "",
        "vip_subnet_id": "75da4441-db7c-4bdb-8ef5-b690c2fa9432",
        "project_id": "3063ff46d451438dbd19b5b4e48b6aa5",
        "name": "tsd",
        "tenant_id": "3063ff46d451438dbd19b5b4e48b6aa5",
        "description": "",
        "nova_flavor_id": "f4d23537-8a87-4c32-bb0b-60328e6f4374",
        "created_at": "2019-11-25T04:20:28",
        "provider": "amphora"
    }
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	lb, err := client.CloudLoadBalancer.Create(ctx, &LoadBalancerCreateRequest{
		Description: "Test Create LB",
		Name:        "LB",
	})
	require.NoError(t, err)
	assert.Equal(t, "e389f5eb-07b5-486b-be4d-4d4d1299f0ab", lb.ID)
	assert.Equal(t, "PENDING_CREATE", lb.ProvisioningStatus)
	assert.Equal(t, "OFFLINE", lb.OperatingStatus)
	assert.Equal(t, "amphora", lb.Provider)
}

func TestLoadBalancerGet(t *testing.T) {
	setup()
	defer teardown()

	var l cloudLoadBalancerService
	mux.HandleFunc(testlib.LoadBalancerURL(l.itemPath("ae8e2072-31fb-464a-8285-bc2f2a6bab4d")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
    "created_at": "2018-09-18T03:43:29",
    "tenant_id": "1e7f10a9850b45b488a3f0417ccb60e0",
    "type": "small",
    "pools": [
        {
            "id": "1fb271b2-a77e-4afc-8ec6-c6bc110f4c75"
        }
    ],
    "provisioning_status": "ACTIVE",
    "operating_status": "ONLINE",
    "name": "sapd-test",
    "admin_state_up": true,
    "vip_port_id": "59b5004b-baa7-463d-bab8-409883ce2458",
    "vip_address": "103.56.156.127",
    "network_type": "external",
    "vip_network_id": "9f36fce7-e2c5-44aa-824b-b83c2dca47f6",
    "vip_qos_policy_id": "94c75cb1-ffe9-4dba-8f37-a375fc10462d",
    "project_id": "1e7f10a9850b45b488a3f0417ccb60e0",
    "vip_subnet_id": "bbad9d0a-09ee-4053-a4f8-9cd8e7ea5e86",
    "listeners": [
        {
            "id": "5482c4a4-f822-46d0-9af3-026f7579d653"
        }
    ],
    "updated_at": "2018-09-18T03:45:30",
    "provider": "amphora",
    "description": "",
    "flavor_id": "",
    "id": "ae8e2072-31fb-464a-8285-bc2f2a6bab4d"
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	lb, err := client.CloudLoadBalancer.Get(ctx, "ae8e2072-31fb-464a-8285-bc2f2a6bab4d")
	require.NoError(t, err)
	assert.Equal(t, "ae8e2072-31fb-464a-8285-bc2f2a6bab4d", lb.ID)
	assert.Equal(t, "ACTIVE", lb.ProvisioningStatus)
	assert.Equal(t, "ONLINE", lb.OperatingStatus)
}

func TestLoadBalancerDelete(t *testing.T) {
	setup()
	defer teardown()

	var l cloudLoadBalancerService
	mux.HandleFunc(testlib.LoadBalancerURL(l.itemPath("ae8e2072-31fb-464a-8285-bc2f2a6bab4d")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
		w.WriteHeader(http.StatusNoContent)
	})

	require.NoError(t, client.CloudLoadBalancer.Delete(ctx, &LoadBalancerDeleteRequest{ID: "ae8e2072-31fb-464a-8285-bc2f2a6bab4d"}))
}

func TestLoadBalancerUpdate(t *testing.T) {
	setup()
	defer teardown()

	var l cloudLoadBalancerService
	mux.HandleFunc(testlib.LoadBalancerURL(l.itemPath("ae8e2072-31fb-464a-8285-bc2f2a6bab4d")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPut, r.Method)
		var payload struct {
			LoadBalancer *LoadBalancerUpdateRequest `json:"loadbalancer"`
		}
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		assert.Equal(t, "Temporarily disabled load balancer", *payload.LoadBalancer.Description)
		assert.Equal(t, "disabled_load_balancer", *payload.LoadBalancer.Name)
		assert.True(t, *payload.LoadBalancer.AdminStateUp)

		resp := `
{
    "loadbalancer": {
        "description": "Temporarily disabled load balancer",
        "project_id": "e3cd678b11784734bc366148aa37580e",
        "provisioning_status": "PENDING_UPDATE",
        "flavor_id": "",
        "vip_subnet_id": "d4af86e1-0051-488c-b7a0-527f97490c9a",
        "vip_address": "203.0.113.50",
        "vip_network_id": "d0d217df-3958-4fbf-a3c2-8dad2908c709",
        "vip_port_id": "b4ca07d1-a31e-43e2-891a-7d14f419f342",
        "provider": "octavia",
        "created_at": "2017-02-28T00:41:44",
        "updated_at": "2017-02-28T00:43:30",
        "id": "8b6fc468-07d5-4d8b-a0b9-695060e72c31",
        "operating_status": "ONLINE",
        "name": "disabled_load_balancer"
    }
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	adminStateUp := true
	desc := "Temporarily disabled load balancer"
	name := "disabled_load_balancer"
	lb, err := client.CloudLoadBalancer.Update(ctx, "ae8e2072-31fb-464a-8285-bc2f2a6bab4d", &LoadBalancerUpdateRequest{
		Description:  &desc,
		Name:         &name,
		AdminStateUp: &adminStateUp,
	})
	require.NoError(t, err)
	require.Equal(t, "8b6fc468-07d5-4d8b-a0b9-695060e72c31", lb.ID)
}

func TestListenerList(t *testing.T) {
	setup()
	defer teardown()

	var l cloudLoadBalancerListenerResource
	mux.HandleFunc(testlib.LoadBalancerURL(l.resourcePath("ae8e2072-31fb-464a-8285-bc2f2a6bab4d")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
    "listeners_links": [],
    "listeners": [
        {
            "timeout_tcp_inspect": 0,
            "insert_headers": {
                "X-Forwarded-Port": "true",
                "X-Forwarded-For": "true"
            },
            "protocol": "HTTP",
            "default_pool_id": "1fb271b2-a77e-4afc-8ec6-c6bc110f4c75",
            "provisioning_status": "ACTIVE",
            "loadbalancers": [
                {
                    "id": "ae8e2072-31fb-464a-8285-bc2f2a6bab4d"
                }
            ],
            "protocol_port": 80,
            "operating_status": "ONLINE",
            "created_at": "2018-09-18T03:43:31",
            "admin_state_up": true,
            "default_tls_container_ref": null,
            "tenant_id": "1e7f10a9850b45b488a3f0417ccb60e0",
            "l7policies": [],
            "timeout_member_connect": 5000,
            "timeout_member_data": 50000,
            "project_id": "1e7f10a9850b45b488a3f0417ccb60e0",
            "updated_at": "2018-09-18T03:45:33",
            "connection_limit": -1,
            "name": "Default Listener",
            "sni_container_refs": [],
            "description": "",
            "timeout_client_data": 50000,
            "id": "5482c4a4-f822-46d0-9af3-026f7579d653"
        }
    ]
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	listeners, err := client.CloudLoadBalancer.Listeners().List(ctx, "ae8e2072-31fb-464a-8285-bc2f2a6bab4d", &ListOptions{})
	require.NoError(t, err)
	assert.Len(t, listeners, 1)
	listener := listeners[0]
	assert.Equal(t, "5482c4a4-f822-46d0-9af3-026f7579d653", listener.ID)
}

func TestListenerCreate(t *testing.T) {
	setup()
	defer teardown()

	var l cloudLoadBalancerListenerResource
	mux.HandleFunc(testlib.LoadBalancerURL(l.resourcePath("ae8e2072-31fb-464a-8285-bc2f2a6bab4d")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		var payload struct {
			CloudLoadBalancerListener *CloudLoadBalancerListenerCreateRequest `json:"listener"`
		}
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		assert.Equal(t, "Test Create CloudLoadBalancerListener", *payload.CloudLoadBalancerListener.Description)
		assert.Equal(t, "CloudLoadBalancerListener", *payload.CloudLoadBalancerListener.Name)

		resp := `
{
	"listener": {
		"timeout_tcp_inspect": 0,
		"insert_headers": {
			"X-Forwarded-Port": "true",
			"X-Forwarded-For": "true"
		},
		"protocol": "HTTP",
		"default_pool_id": "1fb271b2-a77e-4afc-8ec6-c6bc110f4c75",
		"provisioning_status": "ACTIVE",
		"loadbalancers": [
			{
				"id": "ae8e2072-31fb-464a-8285-bc2f2a6bab4d"
			}
		],
		"protocol_port": 80,
		"operating_status": "ONLINE",
		"created_at": "2018-09-18T03:43:31",
		"admin_state_up": true,
		"default_tls_container_ref": null,
		"tenant_id": "1e7f10a9850b45b488a3f0417ccb60e0",
		"l7policies": [],
		"timeout_member_connect": 5000,
		"timeout_member_data": 50000,
		"project_id": "1e7f10a9850b45b488a3f0417ccb60e0",
		"updated_at": "2018-09-18T03:45:33",
		"connection_limit": -1,
		"name": "Default Listener",
		"sni_container_refs": null,
		"description": "",
		"timeout_client_data": 50000,
		"id": "5482c4a4-f822-46d0-9af3-026f7579d653"
	}
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	name := "CloudLoadBalancerListener"
	desc := "Test Create CloudLoadBalancerListener"
	listener, err := client.CloudLoadBalancer.Listeners().Create(ctx, "ae8e2072-31fb-464a-8285-bc2f2a6bab4d", &CloudLoadBalancerListenerCreateRequest{
		Description: &desc,
		Name:        &name,
	})
	require.NoError(t, err)
	assert.Equal(t, "5482c4a4-f822-46d0-9af3-026f7579d653", listener.ID)
	assert.Equal(t, "HTTP", listener.Protocol)
	assert.Nil(t, listener.DefaultTLSContainerRef)
}

func TestListenerGet(t *testing.T) {
	setup()
	defer teardown()

	var l cloudLoadBalancerListenerResource
	mux.HandleFunc(testlib.LoadBalancerURL(l.itemPath("5482c4a4-f822-46d0-9af3-026f7579d653")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
    "timeout_tcp_inspect": 0,
    "insert_headers": {
        "X-Forwarded-Port": "true",
        "X-Forwarded-For": "true"
    },
    "protocol": "HTTP",
    "default_pool_id": "1fb271b2-a77e-4afc-8ec6-c6bc110f4c75",
    "provisioning_status": "ACTIVE",
    "loadbalancers": [
        {
            "id": "ae8e2072-31fb-464a-8285-bc2f2a6bab4d"
        }
    ],
    "protocol_port": 80,
    "operating_status": "ONLINE",
    "created_at": "2018-09-18T03:43:31",
    "admin_state_up": true,
    "default_tls_container_ref": null,
    "tenant_id": "1e7f10a9850b45b488a3f0417ccb60e0",
    "l7policies": [],
    "timeout_member_connect": 5000,
    "timeout_member_data": 50000,
    "project_id": "1e7f10a9850b45b488a3f0417ccb60e0",
    "updated_at": "2018-09-18T03:45:33",
    "connection_limit": -1,
    "name": "Default Listener",
    "sni_container_refs": [],
    "description": "",
    "timeout_client_data": 50000,
    "id": "5482c4a4-f822-46d0-9af3-026f7579d653"
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	listener, err := client.CloudLoadBalancer.Listeners().Get(ctx, "5482c4a4-f822-46d0-9af3-026f7579d653")
	require.NoError(t, err)
	assert.Equal(t, "5482c4a4-f822-46d0-9af3-026f7579d653", listener.ID)
}

func TestListenerUpdate(t *testing.T) {
	setup()
	defer teardown()

	var l cloudLoadBalancerListenerResource
	mux.HandleFunc(testlib.LoadBalancerURL(l.itemPath("023f2e34-7806-443b-bfae-16c324569a3d")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPut, r.Method)

		var payload struct {
			CloudLoadBalancerListener *CloudLoadBalancerListenerUpdateRequest `json:"listener"`
		}
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		assert.Equal(t, "Test Update CloudLoadBalancerListener", *payload.CloudLoadBalancerListener.Description)
		assert.Equal(t, "ListenerUpdated", *payload.CloudLoadBalancerListener.Name)

		resp := `
{
    "listener": {
        "description": "An updated great TLS listener",
        "admin_state_up": true,
        "project_id": "e3cd678b11784734bc366148aa37580e",
        "protocol": "TERMINATED_HTTPS",
        "protocol_port": 443,
        "provisioning_status": "PENDING_UPDATE",
        "default_tls_container_ref": "http://198.51.100.10:9311/v1/containers/a570068c-d295-4780-91d4-3046a325db51",
        "loadbalancers": [
            {
                "id": "607226db-27ef-4d41-ae89-f2a800e9c2db"
            }
        ],
        "insert_headers": {
            "X-Forwarded-Port": "true",
            "X-Forwarded-For": "false"
        },
        "created_at": "2017-02-28T00:42:44",
        "updated_at": "2017-02-28T00:44:30",
        "id": "023f2e34-7806-443b-bfae-16c324569a3d",
        "operating_status": "OFFLINE",
        "default_pool_id": "ddb2b28f-89e9-45d3-a329-a359c3e39e4a",
        "sni_container_refs": [
            "http://198.51.100.10:9311/v1/containers/a570068c-d295-4780-91d4-3046a325db51",
            "http://198.51.100.10:9311/v1/containers/aaebb31e-7761-4826-8cb4-2b829caca3ee"
        ],
        "l7policies": [
            {
                "id": "5e618272-339d-4a80-8d14-dbc093091bb1"
            }
        ],
        "name": "great_updated_tls_listener",
        "timeout_client_data": 100000,
        "timeout_member_connect": 1000,
        "timeout_member_data": 100000,
        "timeout_tcp_inspect": 5
    }
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	name := "ListenerUpdated"
	desc := "Test Update CloudLoadBalancerListener"
	_, err := client.CloudLoadBalancer.Listeners().Update(ctx, "023f2e34-7806-443b-bfae-16c324569a3d", &CloudLoadBalancerListenerUpdateRequest{
		Name:        &name,
		Description: &desc,
	})
	require.NoError(t, err)
}

func TestListenerDelete(t *testing.T) {
	setup()
	defer teardown()

	var l cloudLoadBalancerListenerResource
	mux.HandleFunc(testlib.LoadBalancerURL(l.itemPath("023f2e34-7806-443b-bfae-16c324569a3d")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
		w.WriteHeader(http.StatusNoContent)
	})

	require.NoError(t, client.CloudLoadBalancer.Listeners().Delete(ctx, "023f2e34-7806-443b-bfae-16c324569a3d"))
}

func TestMemberList(t *testing.T) {
	setup()
	defer teardown()

	var m cloudLoadBalancerMemberResource
	mux.HandleFunc(testlib.LoadBalancerURL(m.resourcePath("023f2e34-7806-443b-bfae-16c324569a3d")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
    "members": [
        {
            "backup": false,
            "created_at": "2018-09-18T07:25:04",
            "weight": 1,
            "address": "10.6.169.102",
            "monitor_port": null,
            "subnet_id": "2f866d94-8218-4c9f-8c96-358837e63e6e",
            "protocol_port": 80,
            "project_id": "1e7f10a9850b45b488a3f0417ccb60e0",
            "provisioning_status": "ACTIVE",
            "monitor_address": null,
            "operating_status": "ONLINE",
            "updated_at": "2018-09-18T07:25:21",
            "name": "sapd-lemp-8",
            "admin_state_up": true,
            "tenant_id": "1e7f10a9850b45b488a3f0417ccb60e0",
            "id": "0b9b1602-fb7a-4f9e-ac2e-99f2d4f7b494"
        },
        {
            "backup": false,
            "created_at": "2018-09-18T07:25:22",
            "weight": 1,
            "address": "10.6.169.31",
            "monitor_port": null,
            "subnet_id": "2f866d94-8218-4c9f-8c96-358837e63e6e",
            "protocol_port": 80,
            "project_id": "1e7f10a9850b45b488a3f0417ccb60e0",
            "provisioning_status": "ACTIVE",
            "monitor_address": null,
            "operating_status": "ONLINE",
            "updated_at": "2018-09-18T07:25:27",
            "name": "sapd-lemp-11",
            "admin_state_up": true,
            "tenant_id": "1e7f10a9850b45b488a3f0417ccb60e0",
            "id": "54277bf2-68ea-4ddd-87ee-6bf4c91850a5"
        }
    ],
    "members_links": []
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	members, err := client.CloudLoadBalancer.Members().List(ctx, "023f2e34-7806-443b-bfae-16c324569a3d", &ListOptions{})
	require.NoError(t, err)
	assert.Len(t, members, 2)
}

func TestMemberGet(t *testing.T) {
	setup()
	defer teardown()

	var m cloudLoadBalancerMemberResource
	mux.HandleFunc(testlib.LoadBalancerURL(m.itemPath("023f2e34-7806-443b-bfae-16c324569a3d", "0b9b1602-fb7a-4f9e-ac2e-99f2d4f7b494")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
    "backup": false,
    "created_at": "2018-09-18T07:25:04",
    "weight": 1,
    "address": "10.6.169.102",
    "monitor_port": null,
    "subnet_id": "2f866d94-8218-4c9f-8c96-358837e63e6e",
    "protocol_port": 80,
    "project_id": "1e7f10a9850b45b488a3f0417ccb60e0",
    "provisioning_status": "ACTIVE",
    "monitor_address": null,
    "operating_status": "ONLINE",
    "updated_at": "2018-09-18T07:25:21",
    "name": "sapd-lemp-8",
    "admin_state_up": true,
    "tenant_id": "1e7f10a9850b45b488a3f0417ccb60e0",
    "id": "0b9b1602-fb7a-4f9e-ac2e-99f2d4f7b494"
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	member, err := client.CloudLoadBalancer.Members().Get(ctx, "023f2e34-7806-443b-bfae-16c324569a3d", "0b9b1602-fb7a-4f9e-ac2e-99f2d4f7b494")
	require.NoError(t, err)
	assert.Equal(t, "0b9b1602-fb7a-4f9e-ac2e-99f2d4f7b494", member.ID)
}

func TestMemberUpdate(t *testing.T) {
	setup()
	defer teardown()

	var m cloudLoadBalancerMemberResource
	mux.HandleFunc(testlib.LoadBalancerURL(m.itemPath("023f2e34-7806-443b-bfae-16c324569a3d", "957a1ace-1bd2-449b-8455-820b6e4b63f3")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPut, r.Method)

		resp := `
{
    "member": {
        "monitor_port": 8080,
        "project_id": "e3cd678b11784734bc366148aa37580e",
        "name": "web-server-1",
        "weight": 20,
        "backup": false,
        "admin_state_up": true,
        "subnet_id": "bbb35f84-35cc-4b2f-84c2-a6a29bba68aa",
        "created_at": "2017-05-11T17:21:34",
        "provisioning_status": "PENDING_UPDATE",
        "monitor_address": null,
        "updated_at": "2017-05-11T17:21:37",
        "address": "192.0.2.16",
        "protocol_port": 80,
        "id": "957a1ace-1bd2-449b-8455-820b6e4b63f3",
        "operating_status": "NO_MONITOR"
    }
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	name := "MemberUpdated"
	_, err := client.CloudLoadBalancer.Members().Update(ctx, "023f2e34-7806-443b-bfae-16c324569a3d", "957a1ace-1bd2-449b-8455-820b6e4b63f3", &CloudLoadBalancerMemberUpdateRequest{
		Name: name,
	})
	require.NoError(t, err)
}

func TestMemberDelete(t *testing.T) {
	setup()
	defer teardown()

	var m cloudLoadBalancerMemberResource
	mux.HandleFunc(testlib.LoadBalancerURL(m.itemPath("023f2e34-7806-443b-bfae-16c324569a3d", "957a1ace-1bd2-449b-8455-820b6e4b63f3")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
		w.WriteHeader(http.StatusNoContent)
	})

	require.NoError(t, client.CloudLoadBalancer.Members().Delete(ctx, "023f2e34-7806-443b-bfae-16c324569a3d", "957a1ace-1bd2-449b-8455-820b6e4b63f3"))
}

func TestMemberCreate(t *testing.T) {
	setup()
	defer teardown()

	var m cloudLoadBalancerMemberResource
	mux.HandleFunc(testlib.LoadBalancerURL(m.resourcePath("023f2e34-7806-443b-bfae-16c324569a3d")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		var payload struct {
			Member *CloudLoadBalancerMemberCreateRequest `json:"member"`
		}
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		assert.Equal(t, "web-server-1", payload.Member.Name)
		assert.Equal(t, 80, payload.Member.ProtocolPort)
		assert.Equal(t, "192.0.2.16", payload.Member.Address)

		resp := `
{
    "member": {
        "monitor_port": 8080,
        "project_id": "e3cd678b11784734bc366148aa37580e",
        "name": "web-server-1",
        "weight": 20,
        "backup": false,
        "admin_state_up": true,
        "subnet_id": "bbb35f84-35cc-4b2f-84c2-a6a29bba68aa",
        "created_at": "2017-05-11T17:21:34",
        "provisioning_status": "ACTIVE",
        "monitor_address": null,
        "updated_at": "2017-05-11T17:21:37",
        "address": "192.0.2.16",
        "protocol_port": 80,
        "id": "957a1ace-1bd2-449b-8455-820b6e4b63f3",
        "operating_status": "NO_MONITOR",
        "tags": ["test_tag"]
    }
}
`
		_, _ = fmt.Fprint(w, resp)
	})
	mcr := CloudLoadBalancerMemberCreateRequest{
		Name:         "web-server-1",
		Address:      "192.0.2.16",
		ProtocolPort: 80,
	}
	response, err := client.CloudLoadBalancer.Members().Create(ctx, "023f2e34-7806-443b-bfae-16c324569a3d", &mcr)
	require.NoError(t, err)
	assert.Equal(t, "957a1ace-1bd2-449b-8455-820b6e4b63f3", response.ID)
	assert.Equal(t, "192.0.2.16", response.Address)
	assert.Equal(t, 80, response.ProtocolPort)

}
func TestPoolList(t *testing.T) {
	setup()
	defer teardown()

	var p cloudLoadBalancerPoolResource
	mux.HandleFunc(testlib.LoadBalancerURL(p.resourcePath("ae8e2072-31fb-464a-8285-bc2f2a6bab4d")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
    "pools": [
        {
            "healthmonitor_id": "a1546f51-aa64-442a-a338-886561834a4c",
            "created_at": "2018-09-18T03:43:31",
            "protocol": "HTTP",
            "admin_state_up": true,
            "members": [],
            "loadbalancers": [
                {
                    "id": "ae8e2072-31fb-464a-8285-bc2f2a6bab4d"
                }
            ],
            "project_id": "1e7f10a9850b45b488a3f0417ccb60e0",
            "provisioning_status": "ACTIVE",
            "session_persistence": null,
            "listeners": [
                {
                    "id": "5482c4a4-f822-46d0-9af3-026f7579d653"
                }
            ],
            "operating_status": "ONLINE",
            "updated_at": "2018-09-18T03:45:33",
            "name": "Default",
            "lb_algorithm": "ROUND_ROBIN",
            "description": "",
            "tenant_id": "1e7f10a9850b45b488a3f0417ccb60e0",
            "id": "1fb271b2-a77e-4afc-8ec6-c6bc110f4c75"
        }
    ],
    "pools_links": []
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	pools, err := client.CloudLoadBalancer.Pools().List(ctx, "ae8e2072-31fb-464a-8285-bc2f2a6bab4d", &ListOptions{})
	require.NoError(t, err)
	assert.Len(t, pools, 1)
	pool := pools[0]
	assert.Equal(t, "1fb271b2-a77e-4afc-8ec6-c6bc110f4c75", pool.ID)
}

func TestPoolCreate(t *testing.T) {
	setup()
	defer teardown()

	var p cloudLoadBalancerPoolResource
	mux.HandleFunc(testlib.LoadBalancerURL(p.resourcePath("ae8e2072-31fb-464a-8285-bc2f2a6bab4d")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		var payload struct {
			CloudLoadBalancerPool *CloudLoadBalancerPoolCreateRequest `json:"pool"`
		}
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		assert.Equal(t, "CloudLoadBalancerPool", *payload.CloudLoadBalancerPool.Name)
		assert.NotNil(t, payload.CloudLoadBalancerPool.SessionPersistence)
		assert.Equal(t, "ROUND_ROBIN", payload.CloudLoadBalancerPool.LBAlgorithm)

		resp := `
{
	"pool": {
		"healthmonitor_id": "a1546f51-aa64-442a-a338-886561834a4c",
		"created_at": "2018-09-18T03:43:31",
		"protocol": "HTTP",
		"admin_state_up": true,
		"members": [],
		"loadbalancers": [
			{
				"id": "ae8e2072-31fb-464a-8285-bc2f2a6bab4d"
			}
		],
		"project_id": "1e7f10a9850b45b488a3f0417ccb60e0",
		"provisioning_status": "ACTIVE",
		"session_persistence": null,
		"listeners": [
			{
				"id": "5482c4a4-f822-46d0-9af3-026f7579d653"
			}
		],
		"operating_status": "ONLINE",
		"updated_at": "2018-09-18T03:45:33",
		"healthmonitor": {
			"url_path": "/",
			"created_at": "2018-09-18T03:43:31",
			"tenant_id": "1e7f10a9850b45b488a3f0417ccb60e0",
			"type": "HTTP",
			"delay": 5,
			"max_retries": 3,
			"pools": [
				{
					"id": "1fb271b2-a77e-4afc-8ec6-c6bc110f4c75"
				}
			],
			"provisioning_status": "ACTIVE",
			"http_method": "GET",
			"operating_status": "OFFLINE",
			"updated_at": "2018-09-18T03:45:30",
			"name": "",
			"admin_state_up": true,
			"max_retries_down": 3,
			"timeout": 5,
			"project_id": "1e7f10a9850b45b488a3f0417ccb60e0",
			"expected_codes": "200",
			"id": "a1546f51-aa64-442a-a338-886561834a4c"
		},
		"name": "Default",
		"lb_algorithm": "ROUND_ROBIN",
		"description": "",
		"tenant_id": "1e7f10a9850b45b488a3f0417ccb60e0",
		"id": "1fb271b2-a77e-4afc-8ec6-c6bc110f4c75"
	}
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	name := "CloudLoadBalancerPool"
	pool, err := client.CloudLoadBalancer.Pools().Create(ctx, "ae8e2072-31fb-464a-8285-bc2f2a6bab4d", &CloudLoadBalancerPoolCreateRequest{
		LBAlgorithm: "ROUND_ROBIN",
		Name:        &name,
		SessionPersistence: &SessionPersistence{
			Type:                   "Test",
			CookieName:             nil,
			PersistenceTimeout:     nil,
			PersistenceGranularity: nil,
		},
	})
	require.NoError(t, err)
	assert.Equal(t, "1fb271b2-a77e-4afc-8ec6-c6bc110f4c75", pool.ID)
	assert.Equal(t, "HTTP", pool.Protocol)
}

func TestPoolGet(t *testing.T) {
	setup()
	defer teardown()

	var p cloudLoadBalancerPoolResource
	mux.HandleFunc(testlib.LoadBalancerURL(p.itemPath("1fb271b2-a77e-4afc-8ec6-c6bc110f4c75")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
    "healthmonitor_id": "a1546f51-aa64-442a-a338-886561834a4c",
    "created_at": "2018-09-18T03:43:31",
    "protocol": "HTTP",
    "admin_state_up": true,
    "members": [],
    "loadbalancers": [
        {
            "id": "ae8e2072-31fb-464a-8285-bc2f2a6bab4d"
        }
    ],
    "project_id": "1e7f10a9850b45b488a3f0417ccb60e0",
    "provisioning_status": "ACTIVE",
    "session_persistence": null,
    "listeners": [
        {
            "id": "5482c4a4-f822-46d0-9af3-026f7579d653"
        }
    ],
    "operating_status": "ONLINE",
    "updated_at": "2018-09-18T03:45:33",
    "healthmonitor": {
        "url_path": "/",
        "created_at": "2018-09-18T03:43:31",
        "tenant_id": "1e7f10a9850b45b488a3f0417ccb60e0",
        "type": "HTTP",
        "delay": 5,
        "max_retries": 3,
        "pools": [
            {
                "id": "1fb271b2-a77e-4afc-8ec6-c6bc110f4c75"
            }
        ],
        "provisioning_status": "ACTIVE",
        "http_method": "GET",
        "operating_status": "OFFLINE",
        "updated_at": "2018-09-18T03:45:30",
        "name": "",
        "admin_state_up": true,
        "max_retries_down": 3,
        "timeout": 5,
        "project_id": "1e7f10a9850b45b488a3f0417ccb60e0",
        "expected_codes": "200",
        "id": "a1546f51-aa64-442a-a338-886561834a4c"
    },
    "name": "Default",
    "lb_algorithm": "ROUND_ROBIN",
    "description": "",
    "tenant_id": "1e7f10a9850b45b488a3f0417ccb60e0",
    "id": "1fb271b2-a77e-4afc-8ec6-c6bc110f4c75"
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	pool, err := client.CloudLoadBalancer.Pools().Get(ctx, "1fb271b2-a77e-4afc-8ec6-c6bc110f4c75")
	require.NoError(t, err)
	assert.Equal(t, "1fb271b2-a77e-4afc-8ec6-c6bc110f4c75", pool.ID)
}

func TestPoolUpdate(t *testing.T) {
	setup()
	defer teardown()

	var p cloudLoadBalancerPoolResource
	mux.HandleFunc(testlib.LoadBalancerURL(p.itemPath("4029d267-3983-4224-a3d0-afb3fe16a2cd")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPut, r.Method)

		var payload struct {
			CloudLoadBalancerPool *CloudLoadBalancerPoolUpdateRequest `json:"pool"`
		}
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		assert.Equal(t, "Test Update CloudLoadBalancerPool", *payload.CloudLoadBalancerPool.Description)
		assert.Equal(t, "PoolUpdated", *payload.CloudLoadBalancerPool.Name)

		resp := `
{
    "pool": {
        "lb_algorithm": "LEAST_CONNECTIONS",
        "protocol": "HTTP",
        "description": "Super Least Connections CloudLoadBalancerPool",
        "admin_state_up": true,
        "loadbalancers": [
            {
                "id": "607226db-27ef-4d41-ae89-f2a800e9c2db"
            }
        ],
        "created_at": "2017-05-10T18:14:44",
        "provisioning_status": "PENDING_UPDATE",
        "updated_at": "2017-05-10T23:08:12",
        "session_persistence": {
            "cookie_name": null,
            "type": "SOURCE_IP"
        },
        "listeners": [
            {
                "id": "023f2e34-7806-443b-bfae-16c324569a3d"
            }
        ],
        "members": [],
        "healthmonitor_id": null,
        "project_id": "e3cd678b11784734bc366148aa37580e",
        "id": "4029d267-3983-4224-a3d0-afb3fe16a2cd",
        "operating_status": "ONLINE",
        "name": "super-least-conn-pool"
    }
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	name := "PoolUpdated"
	desc := "Test Update CloudLoadBalancerPool"
	_, err := client.CloudLoadBalancer.Pools().Update(ctx, "4029d267-3983-4224-a3d0-afb3fe16a2cd", &CloudLoadBalancerPoolUpdateRequest{
		Name:        &name,
		Description: &desc,
	})
	require.NoError(t, err)
}

func TestPoolDelete(t *testing.T) {
	setup()
	defer teardown()
	var p cloudLoadBalancerPoolResource

	mux.HandleFunc(testlib.LoadBalancerURL(p.itemPath("023f2e34-7806-443b-bfae-16c324569a3d")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
		w.WriteHeader(http.StatusNoContent)
	})

	require.NoError(t, client.CloudLoadBalancer.Pools().Delete(ctx, "023f2e34-7806-443b-bfae-16c324569a3d"))
}

func TestHealthMonitorCreate(t *testing.T) {
	setup()
	defer teardown()

	poolID := "4029d267-3983-4224-a3d0-afb3fe16a2cd"

	mux.HandleFunc(testlib.LoadBalancerURL("/"+strings.Join([]string{"pool", poolID, "healthmonitor"}, "/")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		var payload struct {
			CloudLoadBalancerHealthMonitor *CloudLoadBalancerHealthMonitorCreateRequest `json:"healthmonitor"`
		}
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		assert.Equal(t, "super-pool-health-monitor", payload.CloudLoadBalancerHealthMonitor.Name)
		assert.Equal(t, "HTTP", payload.CloudLoadBalancerHealthMonitor.Type)
		assert.Equal(t, "200", payload.CloudLoadBalancerHealthMonitor.ExpectedCodes)

		resp := `
{
    "healthmonitor": {
        "project_id": "e3cd678b11784734bc366148aa37580e",
        "name": "super-pool-health-monitor",
        "admin_state_up": true,
        "pools": [
            {
                "id": "4029d267-3983-4224-a3d0-afb3fe16a2cd"
            }
        ],
        "created_at": "2017-05-11T23:53:47",
        "provisioning_status": "ACTIVE",
        "updated_at": "2017-05-11T23:53:47",
        "delay": 10,
        "expected_codes": "200",
        "max_retries": 1,
        "http_method": "GET",
        "timeout": 5,
        "max_retries_down": 3,
        "url_path": "/",
        "type": "HTTP",
        "id": "8ed3c5ac-6efa-420c-bedb-99ba14e58db5",
        "operating_status": "ONLINE",
        "tags": ["test_tag"],
        "http_version": 1.1,
        "domain_name": "testlab.com"
    }
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	hmcr := CloudLoadBalancerHealthMonitorCreateRequest{
		Name:           "super-pool-health-monitor",
		Type:           "HTTP",
		Delay:          10,
		MaxRetries:     1,
		MaxRetriesDown: 3,
		TimeOut:        5,
		HTTPMethod:     "GET",
		URLPath:        "/",
		ExpectedCodes:  "200",
	}

	hm, err := client.CloudLoadBalancer.HealthMonitors().Create(ctx, "4029d267-3983-4224-a3d0-afb3fe16a2cd", &hmcr)
	require.NoError(t, err)
	assert.Equal(t, "8ed3c5ac-6efa-420c-bedb-99ba14e58db5", hm.ID)
	assert.Equal(t, "HTTP", hm.Type)
	assert.Equal(t, "GET", hm.HTTPMethod)
}

func TestHealthMonitorGet(t *testing.T) {
	setup()
	defer teardown()

	var h cloudLoadBalancerHealthMonitorResource
	mux.HandleFunc(testlib.LoadBalancerURL(h.itemPath("06052618-d756-4cf4-8e68-cfe33151eab2")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
    "max_retries": 3,
    "operating_status": "ONLINE",
    "http_version": null,
    "id": "06052618-d756-4cf4-8e68-cfe33151eab2",
    "timeout": 5,
    "url_path": "/",
    "project_id": "17a1c3c952c84b3e84a82ddd48364938",
    "http_method": "GET",
    "domain_name": null,
    "admin_state_up": true,
    "delay": 5,
    "type": "HTTP",
    "created_at": "2020-05-23T02:41:52",
    "pools": [
        {
            "id": "745ca8f4-af18-49be-a2f8-9fb39600a66c"
        }
    ],
    "name": "",
    "updated_at": "2020-05-25T16:59:55",
    "expected_codes": "200-409",
    "provisioning_status": "ACTIVE",
    "max_retries_down": 3,
    "tags": [],
    "tenant_id": "17a1c3c952c84b3e84a82ddd48364938"
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	hm, err := client.CloudLoadBalancer.HealthMonitors().Get(ctx, "06052618-d756-4cf4-8e68-cfe33151eab2")
	require.NoError(t, err)
	assert.Equal(t, "06052618-d756-4cf4-8e68-cfe33151eab2", hm.ID)
}

func TestHealthMonitorUpdate(t *testing.T) {
	setup()
	defer teardown()

	var hm cloudLoadBalancerHealthMonitorResource
	mux.HandleFunc(testlib.LoadBalancerURL(hm.itemPath("8ed3c5ac-6efa-420c-bedb-99ba14e58db5")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPut, r.Method)

		var payload struct {
			CloudLoadBalancerHealthMonitor *CloudLoadBalancerHealthMonitorUpdateRequest `json:"healthmonitor"`
		}
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		assert.Equal(t, "super-pool-health-monitor-updated", payload.CloudLoadBalancerHealthMonitor.Name)

		resp := `
{
    "healthmonitor": {
        "project_id": "e3cd678b11784734bc366148aa37580e",
        "name": "super-pool-health-monitor-updated",
        "admin_state_up": true,
        "pools": [
            {
                "id": "4029d267-3983-4224-a3d0-afb3fe16a2cd"
            }
        ],
        "created_at": "2017-05-11T23:53:47",
        "provisioning_status": "PENDING_UPDATE",
        "updated_at": "2017-05-11T23:53:47",
        "delay": 5,
        "expected_codes": "200",
        "max_retries": 2,
        "http_method": "HEAD",
        "timeout": 2,
        "max_retries_down": 2,
        "url_path": "/index.html",
        "type": "HTTP",
        "id": "8ed3c5ac-6efa-420c-bedb-99ba14e58db5",
        "operating_status": "ONLINE",
        "tags": ["updated_tag"],
        "http_version": 1.1,
        "domain_name": null
    }
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	method := "HED"
	_, err := client.CloudLoadBalancer.HealthMonitors().Update(ctx, "8ed3c5ac-6efa-420c-bedb-99ba14e58db5", &CloudLoadBalancerHealthMonitorUpdateRequest{
		Name:       "super-pool-health-monitor-updated",
		HTTPMethod: &method,
	})
	require.NoError(t, err)
}

func TestHealthMonitorDelete(t *testing.T) {
	setup()
	defer teardown()
	var hm cloudLoadBalancerHealthMonitorResource

	mux.HandleFunc(testlib.LoadBalancerURL(hm.itemPath("06052618-d756-4cf4-8e68-cfe33151eab2")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
		w.WriteHeader(http.StatusNoContent)
	})

	require.NoError(t, client.CloudLoadBalancer.HealthMonitors().Delete(ctx, "06052618-d756-4cf4-8e68-cfe33151eab2"))
}

func TestBatchUpdateMember(t *testing.T) {
	setup()
	defer teardown()
	var member cloudLoadBalancerMemberResource
	mux.HandleFunc(testlib.LoadBalancerURL(member.resourcePath("06052618-d756-4cf4-8e68-cfe33151eab2")), func(w http.ResponseWriter, r *http.Request) {
		var payload CloudLoadBalancerBatchMemberUpdateRequest
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))

		require.Equal(t, http.MethodPut, r.Method)
		w.WriteHeader(http.StatusNoContent)
	})
	var members = CloudLoadBalancerBatchMemberUpdateRequest{
		Members: []CloudLoadBalancerExtendMemberUpdateRequest{
			{
				CloudLoadBalancerMemberUpdateRequest: CloudLoadBalancerMemberUpdateRequest{
					Name:           "test_members1",
					Weight:         1,
					MonitorAddress: "12.12.123.1",
					MonitorPort:    90,
					Backup:         false,
				},
				Address:      "123.123.123.123",
				ProtocolPort: 111,
			},
			{
				CloudLoadBalancerMemberUpdateRequest: CloudLoadBalancerMemberUpdateRequest{
					Name:           "test_member2",
					Weight:         2,
					MonitorAddress: "112.12.123.1",
					MonitorPort:    90,
				},
				ProtocolPort: 80,
				Address:      "1.1.11.1",
			},
		},
	}
	err := client.CloudLoadBalancer.Members().BatchUpdate(ctx, "06052618-d756-4cf4-8e68-cfe33151eab2", &members)
	require.NoError(t, err)
}

func TestL7PolicyCreate(t *testing.T) {
	setup()
	defer teardown()
	listenerID := "33188fce-15a1-4ef5-8587-f1cc2e1e31de"
	createPath := strings.Join([]string{listenerPath, listenerID, "l7policy"}, "/")
	mux.HandleFunc(testlib.LoadBalancerURL(createPath), func(w http.ResponseWriter, r *http.Request) {
		var payload struct {
			L7Policy CreateL7PolicyRequest `json:"l7policy"`
		}
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		assert.Equal(t, http.MethodPost, r.Method)

		resp := `{
			"l7policy": {
			  	"id": "00ec77c0-dd79-44cc-9476-9d2db4bfaef8",
			  	"name": "ducnv",
			  	"description": "",
			  	"provisioning_status": "PENDING_CREATE",
			  	"operating_status": "OFFLINE",
			  	"admin_state_up": true,
			  	"project_id": "7e96a5c1a73542f599bace3b038399e3",
			  	"action": "REDIRECT_TO_URL",
			  	"listener_id": "33188fce-15a1-4ef5-8587-f1cc2e1e31de",
			  	"redirect_pool_id": null,
			  	"redirect_url": "http://localhost.vn/api",
			  	"redirect_prefix": null,
			  	"position": 1,
			  	"rules": [],
			  	"created_at": "2024-03-24T16:35:04",
			  	"updated_at": null,
			  	"tags": [],
			  	"redirect_http_code": 302,
			  	"tenant_id": "7e96a5c1a73542f599bace3b038399e3"
			}
		}`
		_, _ = fmt.Fprint(w, resp)
	})
	rule := L7PolicyRuleRequest{
		Invert:      true,
		Type:        "HOST_NAME",
		CompareType: "EQUAL_TO",
		Value:       "localhost.vn",
	}
	rules := []L7PolicyRuleRequest{rule}
	createPolicyPayload := CreateL7PolicyRequest{
		Action:      "REDIRECT_TO_URL",
		Name:        "ducnv",
		Position:    "1",
		RedirectURL: "http://localhost.vn/api",
		Rules:       rules,
	}
	resp, err := client.CloudLoadBalancer.L7Policies().Create(ctx, "33188fce-15a1-4ef5-8587-f1cc2e1e31de", &createPolicyPayload)
	require.NoError(t, err)
	assert.Equal(t, "00ec77c0-dd79-44cc-9476-9d2db4bfaef8", resp.ID)
	assert.Equal(t, "ducnv", resp.Name)
	assert.Equal(t, "REDIRECT_TO_URL", resp.Action)
	assert.Equal(t, "33188fce-15a1-4ef5-8587-f1cc2e1e31de", resp.ListenerID)
	assert.Equal(t, "http://localhost.vn/api", *resp.RedirectURL)
	assert.Equal(t, 1, resp.Position)
}

func TestL7PolicyGet(t *testing.T) {
	setup()
	defer teardown()
	var policy cloudLoadBalancerL7PolicyResource
	mux.HandleFunc(testlib.LoadBalancerURL(policy.itemPath("00ec77c0-dd79-44cc-9476-9d2db4bfaef8")), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)

		resp := `{
			"id": "00ec77c0-dd79-44cc-9476-9d2db4bfaef8",
			"name": "ducnv",
			"description": "",
			"provisioning_status": "ACTIVE",
			"operating_status": "ONLINE",
			"admin_state_up": true,
			"project_id": "7e96a5c1a73542f599bace3b038399e3",
			"action": "REDIRECT_TO_URL",
			"listener_id": "33188fce-15a1-4ef5-8587-f1cc2e1e31de",
			"redirect_pool_id": null,
			"redirect_url": "http://localhost.vn/api",
			"redirect_prefix": null,
			"position": 1,
			"rules": [
				{
					"id": "ab92aa21-10c6-490a-8d7e-f57086c32692"
				}
			],
			"created_at": "2024-03-24T16:35:04",
			"updated_at": "2024-03-24T16:35:10",
			"tags": [],
			"redirect_http_code": 302,
			"tenant_id": "7e96a5c1a73542f599bace3b038399e3"
		}`
		_, _ = fmt.Fprint(w, resp)
	})
	resp, err := client.CloudLoadBalancer.L7Policies().Get(ctx, "00ec77c0-dd79-44cc-9476-9d2db4bfaef8")
	require.NoError(t, err)
	assert.Equal(t, "00ec77c0-dd79-44cc-9476-9d2db4bfaef8", resp.ID)
	assert.Equal(t, "ducnv", resp.Name)
	assert.Equal(t, "REDIRECT_TO_URL", resp.Action)
	assert.Equal(t, "33188fce-15a1-4ef5-8587-f1cc2e1e31de", resp.ListenerID)
	assert.Equal(t, "http://localhost.vn/api", *resp.RedirectURL)
	assert.Equal(t, 1, resp.Position)
}

func TestL7PolicyUpdate(t *testing.T) {
	setup()
	defer teardown()
	var policy cloudLoadBalancerL7PolicyResource
	mux.HandleFunc(testlib.LoadBalancerURL(policy.itemPath("00ec77c0-dd79-44cc-9476-9d2db4bfaef8")), func(w http.ResponseWriter, r *http.Request) {
		var payload struct {
			L7Plicy UpdateL7PolicyRequest `json:"l7policy"`
		}
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		assert.Equal(t, http.MethodPut, r.Method)

		resp := `{
			"l7policy": {
				"id": "00ec77c0-dd79-44cc-9476-9d2db4bfaef8",
				"name": "ducnv-policy",
				"description": "",
				"provisioning_status": "PENDING_UPDATE",
				"operating_status": "ONLINE",
				"admin_state_up": true,
				"project_id": "7e96a5c1a73542f599bace3b038399e3",
				"action": "REDIRECT_TO_URL",
				"listener_id": "33188fce-15a1-4ef5-8587-f1cc2e1e31de",
				"redirect_pool_id": null,
				"redirect_url": "http://localhost.vn/api",
				"redirect_prefix": null,
				"position": 1,
				"rules": [
					{
					"id": "ab92aa21-10c6-490a-8d7e-f57086c32692"
					}
				],
				"created_at": "2024-03-24T16:35:04",
				"updated_at": "2024-03-24T17:01:49",
				"tags": [],
				"redirect_http_code": 302,
				"tenant_id": "7e96a5c1a73542f599bace3b038399e3"
			}
		}`
		_, _ = fmt.Fprint(w, resp)
	})
	rule := UpdateL7PolicyRuleRequest{
		ID: "ab92aa21-10c6-490a-8d7e-f57086c32692",
		L7PolicyRuleRequest: L7PolicyRuleRequest{
			Invert:      true,
			Type:        "HOST_NAME",
			CompareType: "EQUAL_TO",
			Value:       "localhost.vn",
		},
	}
	rules := []UpdateL7PolicyRuleRequest{rule}
	redirectURL := "http://localhost.vn/api"
	updatePolicyPayload := UpdateL7PolicyRequest{
		Action:      "REDIRECT_TO_URL",
		Name:        "ducnv-policy",
		Position:    1,
		RedirectURL: &redirectURL,
		Rules:       rules,
	}
	resp, err := client.CloudLoadBalancer.L7Policies().Update(ctx, "00ec77c0-dd79-44cc-9476-9d2db4bfaef8", &updatePolicyPayload)
	require.NoError(t, err)
	assert.Equal(t, "00ec77c0-dd79-44cc-9476-9d2db4bfaef8", resp.ID)
	assert.Equal(t, "ducnv-policy", resp.Name)
	assert.Equal(t, "REDIRECT_TO_URL", resp.Action)
	assert.Equal(t, "33188fce-15a1-4ef5-8587-f1cc2e1e31de", resp.ListenerID)
	assert.Equal(t, "http://localhost.vn/api", *resp.RedirectURL)
	assert.Equal(t, 1, resp.Position)
}

func TestL7PolicyDelete(t *testing.T) {
	setup()
	defer teardown()
	var policy cloudLoadBalancerL7PolicyResource
	mux.HandleFunc(testlib.LoadBalancerURL(policy.itemPath("00ec77c0-dd79-44cc-9476-9d2db4bfaef8")), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)

		resp := `{
			"message": "Deleting l7 policy"
		}`
		_, _ = fmt.Fprint(w, resp)
	})
	err := client.CloudLoadBalancer.L7Policies().Delete(ctx, "00ec77c0-dd79-44cc-9476-9d2db4bfaef8")
	require.NoError(t, err)
}

func TestL7PolicyRulesList(t *testing.T) {
	setup()
	defer teardown()
	var policy cloudLoadBalancerL7PolicyResource
	path := strings.Join([]string{policy.itemPath("00ec77c0-dd79-44cc-9476-9d2db4bfaef8"), "rules"}, "/")
	mux.HandleFunc(testlib.LoadBalancerURL(path), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)

		resp := `{
			"rules": [
				{
					"id": "b575f1b7-ef8e-4eef-a2d9-b26addf266c0",
					"type": "HOST_NAME",
					"compare_type": "EQUAL_TO",
					"key": null,
					"value": "localhost.vn",
					"invert": true,
					"provisioning_status": "ACTIVE",
					"operating_status": "ONLINE",
					"created_at": "2024-03-24T17:18:24",
					"updated_at": "2024-03-24T17:18:27",
					"project_id": "7e96a5c1a73542f599bace3b038399e3",
					"admin_state_up": true,
					"tags": [],
					"tenant_id": "7e96a5c1a73542f599bace3b038399e3"
				}
			],
			"rules_links": []
		}`
		_, _ = fmt.Fprint(w, resp)
	})
	resp, err := client.CloudLoadBalancer.L7Policies().ListL7PolicyRules(ctx, "00ec77c0-dd79-44cc-9476-9d2db4bfaef8")
	require.NoError(t, err)
	firstRule := resp[0]
	assert.Equal(t, "b575f1b7-ef8e-4eef-a2d9-b26addf266c0", firstRule.ID)
	assert.Equal(t, "HOST_NAME", firstRule.Type)
	assert.Equal(t, "EQUAL_TO", firstRule.CompareType)
	assert.Equal(t, "localhost.vn", firstRule.Value)
	assert.Equal(t, true, firstRule.Invert)
}

func TestL7PolicyRuleCreate(t *testing.T) {
	setup()
	defer teardown()
	var policy cloudLoadBalancerL7PolicyResource
	path := strings.Join([]string{policy.itemPath("00ec77c0-dd79-44cc-9476-9d2db4bfaef8"), "rules"}, "/")
	mux.HandleFunc(testlib.LoadBalancerURL(path), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		var payload struct {
			Rule L7PolicyRuleRequest `json:"rule"`
		}
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		resp := `{
			"rule": {
				"id": "a1509ea5-f4cc-4ae4-86f2-96cf0400a7d5",
				"type": "HOST_NAME",
				"compare_type": "EQUAL_TO",
				"key": "",
				"value": "duc.nv",
				"invert": true,
				"provisioning_status": "PENDING_CREATE",
				"operating_status": "OFFLINE",
				"created_at": "2024-03-24T17:23:11",
				"updated_at": null,
				"project_id": "7e96a5c1a73542f599bace3b038399e3",
				"admin_state_up": true,
				"tags": [],
				"tenant_id": "7e96a5c1a73542f599bace3b038399e3"
			}
		}`
		_, _ = fmt.Fprint(w, resp)
	})
	payload := L7PolicyRuleRequest{
		Invert:      true,
		Type:        "HOST_NAME",
		CompareType: "EQUAL_TO",
		Value:       "duc.nv",
	}
	resp, err := client.CloudLoadBalancer.L7Policies().CreateL7PolicyRule(ctx, "00ec77c0-dd79-44cc-9476-9d2db4bfaef8", payload)
	require.NoError(t, err)
	assert.Equal(t, "a1509ea5-f4cc-4ae4-86f2-96cf0400a7d5", resp.ID)
	assert.Equal(t, "HOST_NAME", resp.Type)
	assert.Equal(t, "EQUAL_TO", resp.CompareType)
	assert.Equal(t, "duc.nv", resp.Value)
	assert.Equal(t, true, resp.Invert)
}
