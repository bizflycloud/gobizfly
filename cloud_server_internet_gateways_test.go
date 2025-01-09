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

func TestInternetGatewayList(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.CloudServerURL(igwPath), func(writer http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		resp := `{
			"internet_gateways": [
				{
					"id": "caf4c25a-dacb-41e6-8e9d-d455c14a579f",
					"name": "igw03",
					"description": "Internet gateway 03",
					"tenant_id": "de6ca11139a249f6a9fed5bc63cda2bb",
					"project_id": "de6ca11139a249f6a9fed5bc63cda2bb",
					"status": "ACTIVE",
					"availability_zones": [
						"HN1"
					],
					"availability_zone_hints": [
						"HN1"
					],
					"flavor_id": null,
					"tags": [
						"IGW"
					],
					"created_at": "2025-01-09T02:48:15Z",
					"updated_at": "2025-01-09T02:48:21Z",
					"external_gateway_info": {
						"network_id": "c958ebb3-38c5-427e-9036-a8789c6f78f9",
						"enable_snat": true,
						"external_fixed_ips": [
							{
								"subnet_id": "43d9596d-4639-4a19-adf5-9c81f29652d5",
								"ip_address": "103.107.180.104"
							}
						]
					},
					"interfaces_info": [
						{
							"subnet_id": "deea0cb8-41a2-4069-b450-17d689d83454",
							"port_id": "57e6ccfb-64a5-4ed2-815e-073ddd8adc56",
							"ip_address": "10.20.2.254",
							"network_id": "aa6f8cd0-98de-42ab-aa3d-5617d3fa66d2",
							"network_info": {
								"id": "aa6f8cd0-98de-42ab-aa3d-5617d3fa66d2",
								"name": "ducnv",
								"description": "",
								"status": "ACTIVE"
							}
						}
					]
				}
			],
			"meta": {
				"next_page": null,
				"prev_page": null
			}
		}`
		_, _ = fmt.Fprint(writer, resp)
	})
	resp, err := client.CloudServer.InternetGateways().List(ctx, ListInternetGatewayOpts{})
	require.NoError(t, err)
	IGWs := resp.InternetGateways
	assert.Len(t, IGWs, 1)
	assert.Equal(t, "caf4c25a-dacb-41e6-8e9d-d455c14a579f", IGWs[0].ID)
	assert.Equal(t, "igw03", IGWs[0].Name)
	assert.Equal(t, "Internet gateway 03", IGWs[0].Description)
	assert.Equal(t, "de6ca11139a249f6a9fed5bc63cda2bb", IGWs[0].ProjectID)
	assert.Equal(t, "ACTIVE", IGWs[0].Status)
	assert.Equal(t, []string{"HN1"}, IGWs[0].AvailabilityZones)
	assert.Equal(t, []string{"HN1"}, IGWs[0].AvailabilityZoneHints)
	assert.Equal(t, []string{"IGW"}, IGWs[0].Tags)
	assert.Equal(t, "aa6f8cd0-98de-42ab-aa3d-5617d3fa66d2", IGWs[0].InterfacesInfo[0].NetworkID)
}

func TestInternetGatewayGet(t *testing.T) {
	setup()
	defer teardown()
	id := "caf4c25a-dacb-41e6-8e9d-d455c14a579f"
	url := fmt.Sprintf("%s/%s", testlib.CloudServerURL(igwPath), id)
	mux.HandleFunc(url, func(writer http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		resp := `{
					"id": "caf4c25a-dacb-41e6-8e9d-d455c14a579f",
					"name": "igw03",
					"description": "Internet gateway 03",
					"tenant_id": "de6ca11139a249f6a9fed5bc63cda2bb",
					"project_id": "de6ca11139a249f6a9fed5bc63cda2bb",
					"status": "ACTIVE",
					"availability_zones": [
						"HN1"
					],
					"availability_zone_hints": [
						"HN1"
					],
					"flavor_id": null,
					"tags": [
						"IGW"
					],
					"created_at": "2025-01-09T02:48:15Z",
					"updated_at": "2025-01-09T02:48:21Z",
					"external_gateway_info": {
						"network_id": "c958ebb3-38c5-427e-9036-a8789c6f78f9",
						"enable_snat": true,
						"external_fixed_ips": [
							{
								"subnet_id": "43d9596d-4639-4a19-adf5-9c81f29652d5",
								"ip_address": "103.107.180.104"
							}
						]
					},
					"interfaces_info": [
						{
							"subnet_id": "deea0cb8-41a2-4069-b450-17d689d83454",
							"port_id": "57e6ccfb-64a5-4ed2-815e-073ddd8adc56",
							"ip_address": "10.20.2.254",
							"network_id": "aa6f8cd0-98de-42ab-aa3d-5617d3fa66d2",
							"network_info": {
								"id": "aa6f8cd0-98de-42ab-aa3d-5617d3fa66d2",
								"name": "ducnv",
								"description": "",
								"status": "ACTIVE"
							}
						}
					]
				}`
		_, _ = fmt.Fprint(writer, resp)
	})

	resp, err := client.CloudServer.InternetGateways().Get(ctx, id)
	require.NoError(t, err)
	assert.Equal(t, "caf4c25a-dacb-41e6-8e9d-d455c14a579f", resp.ID)
	assert.Equal(t, "igw03", resp.Name)
	assert.Equal(t, "Internet gateway 03", resp.Description)
	assert.Equal(t, "de6ca11139a249f6a9fed5bc63cda2bb", resp.ProjectID)
	assert.Equal(t, "ACTIVE", resp.Status)
	assert.Equal(t, []string{"HN1"}, resp.AvailabilityZones)
	assert.Equal(t, []string{"HN1"}, resp.AvailabilityZoneHints)
	assert.Equal(t, []string{"IGW"}, resp.Tags)
	assert.Equal(t, "aa6f8cd0-98de-42ab-aa3d-5617d3fa66d2", resp.InterfacesInfo[0].NetworkID)
}

func TestInternetGatewayCreate(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.CloudServerURL(igwPath), func(writer http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		var (
			payload CreateInternetGatewayPayload
		)
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		assert.Equal(t, "igw03", payload.Name)
		desc := "Internet gateway 03"
		assert.Equal(t, desc, *payload.Description)
		assert.Equal(t, []string{"aa6f8cd0-98de-42ab-aa3d-5617d3fa66d2"}, *payload.NetworkIDs)
		resp := `{
					"id": "caf4c25a-dacb-41e6-8e9d-d455c14a579f",
					"name": "igw03",
					"description": "Internet gateway 03",
					"tenant_id": "de6ca11139a249f6a9fed5bc63cda2bb",
					"project_id": "de6ca11139a249f6a9fed5bc63cda2bb",
					"status": "ACTIVE",
					"availability_zones": [
						"HN1"
					],
					"availability_zone_hints": [
						"HN1"
					],
					"flavor_id": null,
					"tags": [
						"IGW"
					],
					"created_at": "2025-01-09T02:48:15Z",
					"updated_at": "2025-01-09T02:48:21Z",
					"external_gateway_info": {
						"network_id": "c958ebb3-38c5-427e-9036-a8789c6f78f9",
						"enable_snat": true,
						"external_fixed_ips": [
							{
								"subnet_id": "43d9596d-4639-4a19-adf5-9c81f29652d5",
								"ip_address": "103.107.180.104"
							}
						]
					},
					"interfaces_info": [
						{
							"subnet_id": "deea0cb8-41a2-4069-b450-17d689d83454",
							"port_id": "57e6ccfb-64a5-4ed2-815e-073ddd8adc56",
							"ip_address": "10.20.2.254",
							"network_id": "aa6f8cd0-98de-42ab-aa3d-5617d3fa66d2",
							"network_info": {
								"id": "aa6f8cd0-98de-42ab-aa3d-5617d3fa66d2",
								"name": "ducnv",
								"description": "",
								"status": "ACTIVE"
							}
						}
					]
				}`
		_, _ = fmt.Fprint(writer, resp)
	})

	desc := "Internet gateway 03"
	netIDs := []string{"aa6f8cd0-98de-42ab-aa3d-5617d3fa66d2"}
	payload := CreateInternetGatewayPayload{
		Name:        "igw03",
		Description: &desc,
		NetworkIDs:  &netIDs,
	}
	resp, err := client.CloudServer.InternetGateways().Create(ctx, payload)
	require.NoError(t, err)
	assert.Equal(t, "caf4c25a-dacb-41e6-8e9d-d455c14a579f", resp.ID)
	assert.Equal(t, "igw03", resp.Name)
	assert.Equal(t, "Internet gateway 03", resp.Description)
	assert.Equal(t, "de6ca11139a249f6a9fed5bc63cda2bb", resp.ProjectID)
	assert.Equal(t, "ACTIVE", resp.Status)
	assert.Equal(t, []string{"HN1"}, resp.AvailabilityZones)
	assert.Equal(t, []string{"HN1"}, resp.AvailabilityZoneHints)
	assert.Equal(t, []string{"IGW"}, resp.Tags)
	assert.Equal(t, "aa6f8cd0-98de-42ab-aa3d-5617d3fa66d2", resp.InterfacesInfo[0].NetworkID)
}

func TestInternetGatewayUpdate(t *testing.T) {
	setup()
	defer teardown()
	id := "caf4c25a-dacb-41e6-8e9d-d455c14a579f"
	url := fmt.Sprintf("%s/%s", testlib.CloudServerURL(igwPath), id)
	mux.HandleFunc(url, func(writer http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		var (
			payload UpdateInternetGatewayPayload
		)
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		assert.Equal(t, "igw03", payload.Name)
		desc := "Internet gateway 03"
		assert.Equal(t, desc, payload.Description)
		assert.Equal(t, []string{"aa6f8cd0-98de-42ab-aa3d-5617d3fa66d2"}, payload.NetworkIDs)
		resp := `{
					"id": "caf4c25a-dacb-41e6-8e9d-d455c14a579f",
					"name": "igw03",
					"description": "Internet gateway 03",
					"tenant_id": "de6ca11139a249f6a9fed5bc63cda2bb",
					"project_id": "de6ca11139a249f6a9fed5bc63cda2bb",
					"status": "ACTIVE",
					"availability_zones": [
						"HN1"
					],
					"availability_zone_hints": [
						"HN1"
					],
					"flavor_id": null,
					"tags": [
						"IGW"
					],
					"created_at": "2025-01-09T02:48:15Z",
					"updated_at": "2025-01-09T02:48:21Z",
					"external_gateway_info": {
						"network_id": "c958ebb3-38c5-427e-9036-a8789c6f78f9",
						"enable_snat": true,
						"external_fixed_ips": [
							{
								"subnet_id": "43d9596d-4639-4a19-adf5-9c81f29652d5",
								"ip_address": "103.107.180.104"
							}
						]
					},
					"interfaces_info": [
						{
							"subnet_id": "deea0cb8-41a2-4069-b450-17d689d83454",
							"port_id": "57e6ccfb-64a5-4ed2-815e-073ddd8adc56",
							"ip_address": "10.20.2.254",
							"network_id": "aa6f8cd0-98de-42ab-aa3d-5617d3fa66d2",
							"network_info": {
								"id": "aa6f8cd0-98de-42ab-aa3d-5617d3fa66d2",
								"name": "ducnv",
								"description": "",
								"status": "ACTIVE"
							}
						}
					]
				}`
		_, _ = fmt.Fprint(writer, resp)
	})

	desc := "Internet gateway 03"
	netIDs := []string{"aa6f8cd0-98de-42ab-aa3d-5617d3fa66d2"}
	payload := UpdateInternetGatewayPayload{
		Name:        "igw03",
		Description: desc,
		NetworkIDs:  netIDs,
	}
	resp, err := client.CloudServer.InternetGateways().Update(ctx, id, payload)
	require.NoError(t, err)
	assert.Equal(t, "caf4c25a-dacb-41e6-8e9d-d455c14a579f", resp.ID)
	assert.Equal(t, "igw03", resp.Name)
	assert.Equal(t, "Internet gateway 03", resp.Description)
	assert.Equal(t, "de6ca11139a249f6a9fed5bc63cda2bb", resp.ProjectID)
	assert.Equal(t, "ACTIVE", resp.Status)
	assert.Equal(t, []string{"HN1"}, resp.AvailabilityZones)
	assert.Equal(t, []string{"HN1"}, resp.AvailabilityZoneHints)
	assert.Equal(t, []string{"IGW"}, resp.Tags)
	assert.Equal(t, "aa6f8cd0-98de-42ab-aa3d-5617d3fa66d2", resp.InterfacesInfo[0].NetworkID)
}

func TestInternetGatewayDelete(t *testing.T) {
	setup()
	defer teardown()
	id := "caf4c25a-dacb-41e6-8e9d-d455c14a579f"
	url := fmt.Sprintf("%s/%s", testlib.CloudServerURL(igwPath), id)
	mux.HandleFunc(url, func(writer http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		resp := `{
					"message": "OK"
				}`
		_, _ = fmt.Fprint(writer, resp)
	})

	err := client.CloudServer.InternetGateways().Delete(ctx, id)
	require.NoError(t, err)
}
