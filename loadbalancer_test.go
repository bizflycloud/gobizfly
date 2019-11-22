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
