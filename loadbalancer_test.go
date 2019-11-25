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

	lb, err := client.LoadBalancer.Create(ctx, &LoadBalancerCreateRequest{})
	require.NoError(t, err)
	assert.Equal(t, "e389f5eb-07b5-486b-be4d-4d4d1299f0ab", lb.ID)
	assert.Equal(t, "PENDING_CREATE", lb.ProvisioningStatus)
	assert.Equal(t, "OFFLINE", lb.OperatingStatus)
	assert.Equal(t, "amphora", lb.Provider)
}

func TestLoadBalancerGet(t *testing.T) {
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

func TestLoadBalancerUpdate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(loadBalancerPath+"/ae8e2072-31fb-464a-8285-bc2f2a6bab4d", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPut, r.Method)
		_, _ = fmt.Fprint(w, `{}`)
	})

	// TODO(cuonglm): add real test data when clarify Update request payload with @sapd
	_, err := client.LoadBalancer.Update(ctx, "ae8e2072-31fb-464a-8285-bc2f2a6bab4d", &LoadBalancerUpdateRequest{})
	require.NoError(t, err)
}
