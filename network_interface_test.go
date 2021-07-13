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

func TestNetworkInterfacesList(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.CloudServerURL(networkInterfacePath), func(writer http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		resp := `[
    {
        "id": "f8f78df1-43f1-4c73-9f4c-7d64fecb3b34",
        "name": "test-1",
        "network_id": "99b82e5d-98c3-403f-9ff5-8b5b940e3665",
        "tenant_id": "7c759790478644f88e7a58fca8dc6658",
        "mac_address": "fa:16:3e:98:25:bc",
        "admin_state_up": true,
        "status": "DOWN",
        "device_id": "",
        "device_owner": "",
        "fixed_ips": [
            {
                "subnet_id": "cb287da4-a2ee-4ad8-8287-4f070c359ae9",
                "ip_address": "10.20.1.80"
            }
        ],
        "allowed_address_pairs": [],
        "extra_dhcp_opts": [],
        "security_groups": [
            "e87c2f6c-32ec-43d7-b27e-804c4b463238"
        ],
        "description": "",
        "binding:vnic_type": "normal",
        "port_security_enabled": true,
        "qos_policy_id": null,
        "tags": [],
        "created_at": "2021-07-12T08:55:31Z",
        "updated_at": "2021-07-12T08:55:31Z",
        "revision_number": 1,
        "project_id": "7c759790478644f88e7a58fca8dc6658",
        "attached_server": {},
        "firewalls": []
    },
    {
        "id": "f8f78df1-43f1-4c73-9f4c-7d64fecb3b35",
        "name": "test-2",
        "network_id": "99b82e5d-98c3-403f-9ff5-8b5b940e3665",
        "tenant_id": "7c759790478644f88e7a58fca8dc6658",
        "mac_address": "fa:16:3e:98:25:bc",
        "admin_state_up": true,
        "status": "DOWN",
        "device_id": "",
        "device_owner": "",
        "fixed_ips": [
            {
                "subnet_id": "cb287da4-a2ee-4ad8-8287-4f070c359ae9",
                "ip_address": "10.20.1.80"
            }
        ],
        "allowed_address_pairs": [],
        "extra_dhcp_opts": [],
        "security_groups": [
            "e87c2f6c-32ec-43d7-b27e-804c4b463238"
        ],
        "description": "",
        "binding:vnic_type": "normal",
        "port_security_enabled": true,
        "qos_policy_id": null,
        "tags": [],
        "created_at": "2021-07-12T08:55:31Z",
        "updated_at": "2021-07-12T08:55:31Z",
        "revision_number": 1,
        "project_id": "7c759790478644f88e7a58fca8dc6658",
        "attached_server": {},
        "firewalls": []
    },
    {
        "id": "f8f78df1-43f1-4c73-9f4c-7d64fecb3b36",
        "name": "test-3",
        "network_id": "99b82e5d-98c3-403f-9ff5-8b5b940e3665",
        "tenant_id": "7c759790478644f88e7a58fca8dc6658",
        "mac_address": "fa:16:3e:98:25:bc",
        "admin_state_up": true,
        "status": "DOWN",
        "device_id": "",
        "device_owner": "",
        "fixed_ips": [
            {
                "subnet_id": "cb287da4-a2ee-4ad8-8287-4f070c359ae9",
                "ip_address": "10.20.1.80"
            }
        ],
        "allowed_address_pairs": [],
        "extra_dhcp_opts": [],
        "security_groups": [
            "e87c2f6c-32ec-43d7-b27e-804c4b463238"
        ],
        "description": "",
        "binding:vnic_type": "normal",
        "port_security_enabled": true,
        "qos_policy_id": null,
        "tags": [],
        "created_at": "2021-07-12T08:55:31Z",
        "updated_at": "2021-07-12T08:55:31Z",
        "revision_number": 1,
        "project_id": "7c759790478644f88e7a58fca8dc6658",
        "attached_server": {},
        "firewalls": []
    }
]
`
		_, _ = fmt.Fprint(writer, resp)
	})
	networkInterfaces, err := client.NetworkInterface.ListNetworkInterfaces(ctx, &ListNetworkInterfacesOptions{})
	require.NoError(t, err)
	assert.Len(t, networkInterfaces, 3)
	assert.Equal(t, 1, networkInterfaces[0].RevisionNumber)
}

func TestVPCNetworkInterfaceCreate(t *testing.T) {
	setup()
	defer teardown()
	var n networkInterfaceService
	mux.HandleFunc(testlib.CloudServerURL(n.resourceCreateNetworkInterfacePath("99b82e5d-98c3-403f-9ff5-8b5b940e3665")), func(writer http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		var payload *CreateNetworkInterfacePayload
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		assert.Equal(t, "test", payload.Name)
		resp := `{
            "id": "f8f78df1-43f1-4c73-9f4c-7d64fecb3b34",
            "name": "test",
            "network_id": "99b82e5d-98c3-403f-9ff5-8b5b940e3665",
            "tenant_id": "7c759790478644f88e7a58fca8dc6658",
            "mac_address": "fa:16:3e:98:25:bc",
            "admin_state_up": true,
            "status": "DOWN",
            "device_id": "",
            "device_owner": "",
            "fixed_ips": [
                {
                    "subnet_id": "cb287da4-a2ee-4ad8-8287-4f070c359ae9",
                    "ip_address": "10.20.1.80"
                }
            ],
            "allowed_address_pairs": [],
            "extra_dhcp_opts": [],
            "security_groups": [
                "e87c2f6c-32ec-43d7-b27e-804c4b463238"
            ],
            "description": "",
            "binding:vnic_type": "normal",
            "port_security_enabled": true,
            "qos_policy_id": null,
            "tags": [],
            "created_at": "2021-07-12T08:55:31Z",
            "updated_at": "2021-07-12T08:55:31Z",
            "revision_number": 1,
            "project_id": "7c759790478644f88e7a58fca8dc6658",
            "attached_server": {},
            "firewalls": []
        }`
		_, _ = fmt.Fprint(writer, resp)
	})
	networkInterface, err := client.NetworkInterface.CreateNetworkInterface(ctx, "99b82e5d-98c3-403f-9ff5-8b5b940e3665", &CreateNetworkInterfacePayload{
		Name: "test",
	})
	require.NoError(t, err)
	assert.Equal(t, "test", networkInterface.Name)
	assert.Equal(t, true, networkInterface.AdminStateUp)
}

func TestVPCNetworkInterfaceGet(t *testing.T) {
	setup()
	defer teardown()
	var n networkInterfaceService
	mux.HandleFunc(testlib.CloudServerURL(n.resourceNetworkInterfacePath("99b82e5d-98c3-403f-9ff5-8b5b940e3665", "f8f78df1-43f1-4c73-9f4c-7d64fecb3b34")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			resp := `{
                "id": "f8f78df1-43f1-4c73-9f4c-7d64fecb3b34",
                "name": "test_update",
                "network_id": "99b82e5d-98c3-403f-9ff5-8b5b940e3665",
                "tenant_id": "7c759790478644f88e7a58fca8dc6658",
                "mac_address": "fa:16:3e:98:25:bc",
                "admin_state_up": true,
                "status": "DOWN",
                "device_id": "",
                "device_owner": "",
                "fixed_ips": [
                    {
                        "subnet_id": "cb287da4-a2ee-4ad8-8287-4f070c359ae9",
                        "ip_address": "10.20.1.80"
                    }
                ],
                "allowed_address_pairs": [],
                "extra_dhcp_opts": [],
                "security_groups": [
                    "e87c2f6c-32ec-43d7-b27e-804c4b463238"
                ],
                "description": "",
                "binding:vnic_type": "normal",
                "port_security_enabled": true,
                "qos_policy_id": null,
                "tags": [],
                "created_at": "2021-07-12T08:55:31Z",
                "updated_at": "2021-07-12T08:55:31Z",
                "revision_number": 1,
                "project_id": "7c759790478644f88e7a58fca8dc6658",
                "attached_server": {},
                "firewalls": []
            }`
			_, _ = fmt.Fprint(writer, resp)
		})
	vpc, err := client.NetworkInterface.GetNetworkInterface(ctx, "99b82e5d-98c3-403f-9ff5-8b5b940e3665", "f8f78df1-43f1-4c73-9f4c-7d64fecb3b34")
	require.NoError(t, err)
	assert.Equal(t, "2021-07-12T08:55:31Z", vpc.CreatedAt)
}

func TestVPCNetworkInterfaceUpdate(t *testing.T) {
	setup()
	defer teardown()
	var n networkInterfaceService
	mux.HandleFunc(testlib.CloudServerURL(n.resourceNetworkInterfacePath("99b82e5d-98c3-403f-9ff5-8b5b940e3665", "f8f78df1-43f1-4c73-9f4c-7d64fecb3b34")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPut, r.Method)
			var payload *UpdateNetworkInterfacePayload
			require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
			assert.Equal(t, "test_update", payload.Name)
			resp := `{
                "id": "f8f78df1-43f1-4c73-9f4c-7d64fecb3b34",
                "name": "test_update",
                "network_id": "99b82e5d-98c3-403f-9ff5-8b5b940e3665",
                "tenant_id": "7c759790478644f88e7a58fca8dc6658",
                "mac_address": "fa:16:3e:98:25:bc",
                "admin_state_up": true,
                "status": "DOWN",
                "device_id": "",
                "device_owner": "",
                "fixed_ips": [
                    {
                        "subnet_id": "cb287da4-a2ee-4ad8-8287-4f070c359ae9",
                        "ip_address": "10.20.1.80"
                    }
                ],
                "allowed_address_pairs": [],
                "extra_dhcp_opts": [],
                "security_groups": [
                    "e87c2f6c-32ec-43d7-b27e-804c4b463238"
                ],
                "description": "",
                "binding:vnic_type": "normal",
                "port_security_enabled": true,
                "qos_policy_id": null,
                "tags": [],
                "created_at": "2021-07-12T08:55:31Z",
                "updated_at": "2021-07-12T08:55:31Z",
                "revision_number": 1,
                "project_id": "7c759790478644f88e7a58fca8dc6658",
                "attached_server": {},
                "firewalls": []
            }`
			_, _ = fmt.Fprint(writer, resp)
		})
	networkInterface, err := client.NetworkInterface.UpdateNetworkInterface(ctx, "99b82e5d-98c3-403f-9ff5-8b5b940e3665", "f8f78df1-43f1-4c73-9f4c-7d64fecb3b34", &UpdateNetworkInterfacePayload{
		Name: "test_update",
	})
	require.NoError(t, err)
	assert.Equal(t, "test_update", networkInterface.Name)
}

func TestNetworkInterfaceDelete(t *testing.T) {
	setup()
	defer teardown()
	var n networkInterfaceService
	mux.HandleFunc(testlib.CloudServerURL(n.resourceNetworkInterfacePath("99b82e5d-98c3-403f-9ff5-8b5b940e3665", "f8f78df1-43f1-4c73-9f4c-7d64fecb3b34")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodDelete, r.Method)
		})
	require.NoError(t, client.NetworkInterface.DeleteNetworkInterface(ctx, "99b82e5d-98c3-403f-9ff5-8b5b940e3665", "f8f78df1-43f1-4c73-9f4c-7d64fecb3b34"))
}
