package gobizfly

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

func TestPoolList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(loadBalancerPath+"/ae8e2072-31fb-464a-8285-bc2f2a6bab4d/pools", func(w http.ResponseWriter, r *http.Request) {
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

	pools, err := client.Pool.List(ctx, "ae8e2072-31fb-464a-8285-bc2f2a6bab4d", &ListOptions{})
	require.NoError(t, err)
	assert.Len(t, pools, 1)
	pool := pools[0]
	assert.Equal(t, "1fb271b2-a77e-4afc-8ec6-c6bc110f4c75", pool.ID)
}

func TestPoolCreate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(loadBalancerPath+"/ae8e2072-31fb-464a-8285-bc2f2a6bab4d/pools", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		var payload struct {
			Pool *PoolCreateRequest `json:"pool"`
		}
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		assert.Equal(t, "Test Create Pool", *payload.Pool.Description)
		assert.Equal(t, "Pool", *payload.Pool.Name)
		assert.NotNil(t, payload.Pool.SessionPersistence)
		assert.Equal(t, "ROUND_ROBIN", payload.Pool.LBAlgorithm)

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

	name := "Pool"
	desc := "Test Create Pool"
	pool, err := client.Pool.Create(ctx, "ae8e2072-31fb-464a-8285-bc2f2a6bab4d", &PoolCreateRequest{
		LBAlgorithm: "ROUND_ROBIN",
		Description: &desc,
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

	mux.HandleFunc(poolPath+"/1fb271b2-a77e-4afc-8ec6-c6bc110f4c75", func(w http.ResponseWriter, r *http.Request) {
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

	pool, err := client.Pool.Get(ctx, "1fb271b2-a77e-4afc-8ec6-c6bc110f4c75")
	require.NoError(t, err)
	assert.Equal(t, "1fb271b2-a77e-4afc-8ec6-c6bc110f4c75", pool.ID)
}

func TestPoolUpdate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(poolPath+"/4029d267-3983-4224-a3d0-afb3fe16a2cd", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPut, r.Method)

		var payload struct {
			Pool *PoolUpdateRequest `json:"pool"`
		}
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		assert.Equal(t, "Test Update Pool", *payload.Pool.Description)
		assert.Equal(t, "PoolUpdated", *payload.Pool.Name)

		resp := `
{
    "pool": {
        "lb_algorithm": "LEAST_CONNECTIONS",
        "protocol": "HTTP",
        "description": "Super Least Connections Pool",
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
	desc := "Test Update Pool"
	_, err := client.Pool.Update(ctx, "4029d267-3983-4224-a3d0-afb3fe16a2cd", &PoolUpdateRequest{
		Name:        &name,
		Description: &desc,
	})
	require.NoError(t, err)
}

func TestPoolDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(poolPath+"/023f2e34-7806-443b-bfae-16c324569a3d", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
		w.WriteHeader(http.StatusNoContent)
	})

	require.NoError(t, client.Pool.Delete(ctx, "023f2e34-7806-443b-bfae-16c324569a3d"))
}
