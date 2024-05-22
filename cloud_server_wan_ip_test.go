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

func TestWanIPList(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.CloudServerURL(wanIpPath), func(writer http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		resp := `[
  {
    "id": "0c538770-7d2d-448d-92e8-bc55da5a2f08",
    "name": "",
    "network_id": "c8812eaa-6ed8-4ef7-8e20-575d2a23e7b1",
    "tenant_id": "ebbed256d9414b0598719c42dc17e837",
    "mac_address": "fa:16:3e:e8:a7:73",
    "admin_state_up": true,
    "status": "DOWN",
    "device_id": "",
    "device_owner": "",
    "fixed_ips": [
      {
        "subnet_id": "7969d050-f95c-4b27-b310-352a99002159",
        "ip_address": "10.3.247.7",
        "ip_version": 4
      }
    ],
    "allowed_address_pairs": [],
    "extra_dhcp_opts": [],
    "security_groups": [
      "f7f9ed60-46c4-471b-98d5-3e48006bb9ef"
    ],
    "description": "",
    "binding:vnic_type": "normal",
    "port_security_enabled": true,
    "qos_policy_id": null,
    "tags": [],
    "created_at": "2021-09-14T08:21:44Z",
    "updated_at": "2021-09-14T08:21:44Z",
    "revision_number": 1,
    "project_id": "ebbed256d9414b0598719c42dc17e837",
    "network_name": "STAGING_EXT_DIRECTNET_0",
    "bandwidth": 1000000,
    "billing_type": "free",
    "availability_zone": "HN2"
  }
]`
		_, _ = fmt.Fprint(writer, resp)
	})
	wanIps, err := client.CloudServer.PublicNetworkInterfaces().List(ctx)
	require.NoError(t, err)
	assert.Equal(t, 1, len(wanIps))
	assert.Equal(t, "free", wanIps[0].BillingType)
}

func TestWanIPCreate(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.CloudServerURL(wanIpPath), func(writer http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		var payload *CreatePublicNetworkInterfacePayload
		err := json.NewDecoder(r.Body).Decode(&payload)
		require.NoError(t, err)
		assert.Equal(t, "HN1", payload.AvailabilityZone)
		resp := `{
  "id": "ceebf0de-1fc2-4a08-a200-1906a30abe7e",
  "name": "demo_4",
  "network_id": "c8812eaa-6ed8-4ef7-8e20-575d2a23e7b1",
  "tenant_id": "ebbed256d9414b0598719c42dc17e837",
  "mac_address": "fa:16:3e:43:82:b4",
  "admin_state_up": true,
  "status": "DOWN",
  "device_id": "",
  "device_owner": "",
  "fixed_ips": [
    {
      "subnet_id": "7969d050-f95c-4b27-b310-352a99002159",
      "ip_address": "10.3.247.148"
    }
  ],
  "allowed_address_pairs": [],
  "extra_dhcp_opts": [],
  "security_groups": [
    "f7f9ed60-46c4-471b-98d5-3e48006bb9ef"
  ],
  "description": "",
  "binding:vnic_type": "normal",
  "port_security_enabled": true,
  "qos_policy_id": null,
  "tags": [
    "paid"
  ],
  "created_at": "2021-09-14T11:39:35Z",
  "updated_at": "2021-09-14T11:39:36Z",
  "revision_number": 2,
  "project_id": "ebbed256d9414b0598719c42dc17e837",
  "bandwidth": 100000,
  "billing_type": "paid",
  "attached_server": {},
  "availability_zone": "HN2"
}`
		_, _ = fmt.Fprint(writer, resp)
	})
	wanIp, err := client.CloudServer.PublicNetworkInterfaces().Create(ctx, &CreatePublicNetworkInterfacePayload{
		AvailabilityZone: "HN1",
		Name:             "test_123",
	})
	require.NoError(t, err)
	assert.Equal(t, "paid", wanIp.BillingType)
	assert.Equal(t, "HN2", wanIp.AvailabilityZone)
}

func TestWanIPGet(t *testing.T) {
	setup()
	defer teardown()
	var w cloudServerPublicNetworkInterfaceResource
	mux.HandleFunc(testlib.CloudServerURL(w.itemPath("123")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			resp := `{
  "id": "ceebf0de-1fc2-4a08-a200-1906a30abe7e",
  "name": "demo_4",
  "network_id": "c8812eaa-6ed8-4ef7-8e20-575d2a23e7b1",
  "tenant_id": "ebbed256d9414b0598719c42dc17e837",
  "mac_address": "fa:16:3e:43:82:b4",
  "admin_state_up": true,
  "status": "DOWN",
  "device_id": "",
  "device_owner": "",
  "fixed_ips": [
    {
      "subnet_id": "7969d050-f95c-4b27-b310-352a99002159",
      "ip_address": "10.3.247.148"
    }
  ],
  "allowed_address_pairs": [],
  "extra_dhcp_opts": [],
  "security_groups": [
    "f7f9ed60-46c4-471b-98d5-3e48006bb9ef"
  ],
  "description": "",
  "binding:vnic_type": "normal",
  "port_security_enabled": true,
  "qos_policy_id": null,
  "tags": [
    "paid"
  ],
  "created_at": "2021-09-14T11:39:35Z",
  "updated_at": "2021-09-14T11:39:36Z",
  "revision_number": 2,
  "project_id": "ebbed256d9414b0598719c42dc17e837",
  "bandwidth": 100000,
  "billing_type": "paid",
  "attached_server": {},
  "availability_zone": "HN2"
}`
			_, _ = fmt.Fprint(writer, resp)
		})
	wanIp, err := client.CloudServer.PublicNetworkInterfaces().Get(ctx, "123")
	require.NoError(t, err)
	assert.Equal(t, "paid", wanIp.BillingType)
	assert.Equal(t, 100000, wanIp.Bandwidth)
}

func TestWanIpDelete(t *testing.T) {
	setup()
	defer teardown()
	var w cloudServerPublicNetworkInterfaceResource
	mux.HandleFunc(testlib.CloudServerURL(w.itemPath("f8f78df1-43f1-4c73-9f4c-7d64fecb3b34")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodDelete, r.Method)
		})
	require.NoError(t, client.CloudServer.PublicNetworkInterfaces().Delete(ctx, "f8f78df1-43f1-4c73-9f4c-7d64fecb3b34"))
}

func TestWanIPAttachServer(t *testing.T) {
	setup()
	defer teardown()
	var w cloudServerPublicNetworkInterfaceResource
	mux.HandleFunc(testlib.CloudServerURL(w.actionPath("ceebf0de-1fc2-4a08-a200-1906a30abe7e")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
		})
	payload := ActionPublicNetworkInterfacePayload{
		Action:   "attach_server",
		ServerId: "123",
	}
	err := client.CloudServer.PublicNetworkInterfaces().Action(ctx, "ceebf0de-1fc2-4a08-a200-1906a30abe7e", &payload)
	require.NoError(t, err)
}

func TestWanIPDetachServer(t *testing.T) {
	setup()
	defer teardown()
	var w cloudServerPublicNetworkInterfaceResource
	mux.HandleFunc(testlib.CloudServerURL(w.actionPath("ceebf0de-1fc2-4a08-a200-1906a30abe7e")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
		})
	payload := ActionPublicNetworkInterfacePayload{
		Action: "detach_server",
	}
	err := client.CloudServer.PublicNetworkInterfaces().Action(ctx, "ceebf0de-1fc2-4a08-a200-1906a30abe7e", &payload)
	require.NoError(t, err)
}

func TestWanIPConvertToPaid(t *testing.T) {
	setup()
	defer teardown()
	var w cloudServerPublicNetworkInterfaceResource
	mux.HandleFunc(testlib.CloudServerURL(w.actionPath("ceebf0de-1fc2-4a08-a200-1906a30abe7e")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
		})
	payload := ActionPublicNetworkInterfacePayload{
		Action: "convert_to_paid",
	}
	err := client.CloudServer.PublicNetworkInterfaces().Action(ctx, "ceebf0de-1fc2-4a08-a200-1906a30abe7e", &payload)
	require.NoError(t, err)
}
