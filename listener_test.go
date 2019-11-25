package gobizfly

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

func TestListenerList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(loadBalancerPath+"/ae8e2072-31fb-464a-8285-bc2f2a6bab4d/listeners", func(w http.ResponseWriter, r *http.Request) {
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

	listeners, err := client.Listener.List(ctx, "ae8e2072-31fb-464a-8285-bc2f2a6bab4d", &ListOptions{})
	require.NoError(t, err)
	assert.Len(t, listeners, 1)
	listener := listeners[0]
	assert.Equal(t, "5482c4a4-f822-46d0-9af3-026f7579d653", listener.ID)
}

func TestListenerCreate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(loadBalancerPath+"/ae8e2072-31fb-464a-8285-bc2f2a6bab4d/listeners", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		var payload struct {
			Listener *ListenerCreateRequest `json:"listener"`
		}
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		assert.Equal(t, "Test Create Listener", *payload.Listener.Description)
		assert.Equal(t, "Listener", *payload.Listener.Name)

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

	name := "Listener"
	desc := "Test Create Listener"
	listener, err := client.Listener.Create(ctx, "ae8e2072-31fb-464a-8285-bc2f2a6bab4d", &ListenerCreateRequest{
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

	mux.HandleFunc(listenerPath+"/5482c4a4-f822-46d0-9af3-026f7579d653", func(w http.ResponseWriter, r *http.Request) {
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

	listener, err := client.Listener.Get(ctx, "5482c4a4-f822-46d0-9af3-026f7579d653")
	require.NoError(t, err)
	assert.Equal(t, "5482c4a4-f822-46d0-9af3-026f7579d653", listener.ID)
}

func TestListenerUpdate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(listenerPath+"/023f2e34-7806-443b-bfae-16c324569a3d", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPut, r.Method)

		var payload struct {
			Listener *ListenerUpdateRequest `json:"listener"`
		}
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		assert.Equal(t, "Test Update Listener", *payload.Listener.Description)
		assert.Equal(t, "ListenerUpdated", *payload.Listener.Name)

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
	desc := "Test Update Listener"
	_, err := client.Listener.Update(ctx, "023f2e34-7806-443b-bfae-16c324569a3d", &ListenerUpdateRequest{
		Name:        &name,
		Description: &desc,
	})
	require.NoError(t, err)
}

func TestListenerDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(listenerPath+"/023f2e34-7806-443b-bfae-16c324569a3d", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
		w.WriteHeader(http.StatusNoContent)
	})

	require.NoError(t, client.Listener.Delete(ctx, "023f2e34-7806-443b-bfae-16c324569a3d"))
}
