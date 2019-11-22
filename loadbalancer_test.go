package gobizfly

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

func TestLoadBalancerList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(loadBalancerPath, func(w http.ResponseWriter, r *http.Request) {
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

	lbs, err := client.LoadBalancer.List(ctx, &ListOptions{})
	require.NoError(t, err)
	assert.Len(t, lbs, 1)
	lb := lbs[0]
	assert.Equal(t, "ae8e2072-31fb-464a-8285-bc2f2a6bab4d", lb.ID)
}

func TestLoadBalancerCreate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(loadBalancerPath, func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		resp := `
{
    "loadbalancer": {
        "description": "My favorite load balancer",
        "admin_state_up": true,
        "project_id": "e3cd678b11784734bc366148aa37580e",
        "provisioning_status": "PENDING_CREATE",
        "flavor_id": "",
        "vip_subnet_id": "d4af86e1-0051-488c-b7a0-527f97490c9a",
        "vip_address": "203.0.113.50",
        "vip_network_id": "d0d217df-3958-4fbf-a3c2-8dad2908c709",
        "vip_port_id": "b4ca07d1-a31e-43e2-891a-7d14f419f342",
        "provider": "octavia",
        "created_at": "2017-02-28T00:41:44",
        "updated_at": "2017-02-28T00:43:30",
        "id": "607226db-27ef-4d41-ae89-f2a800e9c2db",
        "operating_status": "OFFLINE",
        "name": "best_load_balancer",
        "vip_qos_policy_id": "ec4f78ca-8da8-4e99-8a1a-e3b94595a7a3",
        "tags": ["test_tag"]
    }
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	lb, err := client.LoadBalancer.Create(ctx, &LoadBalancerCreateRequest{})
	require.NoError(t, err)
	assert.Equal(t, "607226db-27ef-4d41-ae89-f2a800e9c2db", lb.ID)
	assert.Equal(t, "PENDING_CREATE", lb.ProvisioningStatus)
}

func TestLoadBalancerCGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(loadBalancerPath+"/ae8e2072-31fb-464a-8285-bc2f2a6bab4d", func(w http.ResponseWriter, r *http.Request) {
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

	lb, err := client.LoadBalancer.Get(ctx, "ae8e2072-31fb-464a-8285-bc2f2a6bab4d")
	require.NoError(t, err)
	assert.Equal(t, "ae8e2072-31fb-464a-8285-bc2f2a6bab4d", lb.ID)
	assert.Equal(t, "ACTIVE", lb.ProvisioningStatus)
	assert.Equal(t, "ONLINE", lb.OperatingStatus)
}

func TestLoadBalancerDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(loadBalancerPath+"/ae8e2072-31fb-464a-8285-bc2f2a6bab4d", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
		w.WriteHeader(http.StatusNoContent)
	})

	require.NoError(t, client.LoadBalancer.Delete(ctx, &LoadBalancerDeleteRequest{ID: "ae8e2072-31fb-464a-8285-bc2f2a6bab4d"}))
}
