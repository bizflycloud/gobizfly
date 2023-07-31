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

func TestZoneList(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.DNSURL(zonesPath), func(writer http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `{
  "zones": [
    {
      "name": "tesrt2.vn",
      "id": "81b3881a-5f11-432c-b4ec-2f57557d1c8c",
      "deleted": 0,
      "created_at": "2020-10-19 09:54:46.718000",
      "update_at": "2020-10-19 09:54:46.718000",
      "tenant_id": "333e2300b3c644cc93359a41a07c2321",
      "nameserver": [
        "ns-captain.dev.bizflycloud.vn.",
        "ns-batman.dev.bizflycloud.vn."
      ],
      "ttl": 3600,
      "active": false,
      "records_set": [
        "f0611990-2b1c-4234-98de-318ebfd1deed",
        "ec3e301d-cd83-44e2-a589-4ee559792d75"
      ]
    }
  ],
  "_meta": {
    "max_results": 1,
    "total": 10,
    "page": 1
  }
}`
		_, _ = fmt.Fprint(writer, resp)
	})
	resp, err := client.DNS.ListZones(ctx, &ListOptions{})
	require.NoError(t, err)
	assert.Equal(t, 1, resp.Meta.MaxResults)
	assert.Equal(t, 3600, resp.Zones[0].TTL)
}

func TestCreateZone(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.DNSURL(zonesPath), func(writer http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		var payload *CreateZonePayload
		assert.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		resp := `{
    "id": "5f01ffb6-ddb5-4541-b978-87a7eed71058",
    "name": "ddddafs.com",
    "deleted": 0,
    "created_at": "2021-07-05 17:45:55.704000",
    "updated_at": "2021-07-05 17:45:55.704000",
    "tenant_id": "bc8d2790fc9a46949818b942c0a824de",
    "serial": 1625481955,
    "nameserver": [
        "ns4.bizflycloud.vn.",
        "ns5.bizflycloud.vn.",
        "ns6.bizflycloud.vn."
    ],
    "ttl": 3600,
    "active": false,
    "record_set": [
        {
            "id": "47b9147c-7332-4dd8-b56e-56c3a30d6bd5",
            "name": "ddddafs.com",
            "type": "NS",
            "ttl": 3600,
            "data": [
                "ns4.bizflycloud.vn.",
                "ns5.bizflycloud.vn.",
                "ns6.bizflycloud.vn."
            ],
            "routing_policy_data": {}
        },
        {
            "id": "30ac3108-ac5e-43e8-9eea-1743d6770aa5",
            "name": "ddddafs.com",
            "type": "SOA",
            "ttl": 3600,
            "data": [
                "ns4.bizflycloud.vn."
            ],
            "routing_policy_data": {}
        }
    ]
}
`
		_, _ = fmt.Fprint(writer, resp)
	})
	zone, err := client.DNS.CreateZone(ctx, &CreateZonePayload{
		Name: "ddddafs.com",
	})
	require.NoError(t, err)
	assert.Equal(t, 0, zone.Deleted)
}

func TestGetZone(t *testing.T) {
	setup()
	defer teardown()
	var d dnsService
	mux.HandleFunc(testlib.DNSURL(d.zoneItemPath("81b3881a-5f11-432c-b4ec-2f57557d1c8c")), func(writer http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		resp := `{
  "name": "tesrt2.vn",
  "id": "81b3881a-5f11-432c-b4ec-2f57557d1c8c",
  "deleted": 0,
  "created_at": "2020-10-19 09:54:46.718000",
  "update_at": "2020-10-19 09:54:46.718000",
  "tenant_id": "333e2300b3c644cc93359a41a07c2321",
  "nameserver": [
    "ns-captain.dev.bizflycloud.vn.",
    "ns-batman.dev.bizflycloud.vn."
  ],
  "ttl": 3600,
  "active": false,
  "records_set": [
    {
      "id": "ec3e301d-cd83-44e2-a589-4ee559792d75",
      "name": "testsdx.vn",
      "type": "NS",
      "ttl": "3600",
      "data": [
        "ns-captain.dev.bizflycloud.vn.",
        "ns-bat.dev.bizflycloud.vn."
      ],
      "routing_policy_data": {}
    },
    {
      "id": "f0611990-2b1c-4234-98de-318ebfd1deed",
      "name": "testsdx.vn",
      "type": "SOA",
      "ttl": 3600,
      "data": [
        "ns-captain.dev.bizflycloud.vn."
      ],
      "routing_policy_data": {}
    }
  ]
}`
		_, _ = fmt.Fprint(writer, resp)
	})
	zone, err := client.DNS.GetZone(ctx, "81b3881a-5f11-432c-b4ec-2f57557d1c8c")
	require.NoError(t, err)
	assert.Equal(t, "333e2300b3c644cc93359a41a07c2321", zone.TenantId)
}

func TestDeleteZone(t *testing.T) {
	setup()
	defer teardown()
	var d dnsService
	mux.HandleFunc(testlib.DNSURL(d.zoneItemPath("81b3881a-5f11-432c-b4ec-2f57557d1c8c")), func(writer http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
	})
	require.NoError(t, client.DNS.DeleteZone(ctx, "81b3881a-5f11-432c-b4ec-2f57557d1c8c"))
}

func TestCreateRecord(t *testing.T) {
	setup()
	defer teardown()
	var d dnsService
	mux.HandleFunc(testlib.DNSURL(strings.Join([]string{d.zoneItemPath("5e6d0d98-ddd4-4d2f-87b0-aa60261cafd4"), "record"}, "/")), func(writer http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		var payload *BaseCreateRecordPayload
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		resp := `{
    "record": {
        "id": "5e6d0d98-ddd4-4d2f-87b0-aa60261cafd4",
        "name": "adfs",
        "type": "A",
        "deleted": 0,
        "create_at": "2021-07-05 17:52:27.240000",
        "update_at": "2021-07-05 17:52:27.240000",
        "ttl": 300,
        "tenant_id": "bc8d2790fc9a46949818b942c0a824de",
        "zone_id": "5f01ffb6-ddb5-4541-b978-87a7eed71058",
        "data": [
            "1.1.1.1"
        ],
        "routing_policy_data": {}
    }
}
`
		_, _ = fmt.Fprint(writer, resp)
	})
	payload := CreateNormalRecordPayload{
		BaseCreateRecordPayload: BaseCreateRecordPayload{
			Name: "testsdx.vn",
			Type: "A",
			TTL:  300,
		},
		Data: []string{"10.5.23.1", "20.1.1.1"},
	}
	resp, err := client.DNS.CreateRecord(ctx, "5e6d0d98-ddd4-4d2f-87b0-aa60261cafd4", &payload)
	require.NoError(t, err)
	assert.Equal(t, "5e6d0d98-ddd4-4d2f-87b0-aa60261cafd4", resp.ID)
}

func TestGetRecord(t *testing.T) {
	setup()
	defer teardown()
	var d dnsService
	mux.HandleFunc(testlib.DNSURL(d.recordItemPath("5e6d0d98-ddd4-4d2f-87b0-aa60261cafd4")), func(writer http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		resp := `{
    "record": {
        "id": "5e6d0d98-ddd4-4d2f-87b0-aa60261cafd4",
        "name": "adfs",
        "type": "A",
        "deleted": 0,
        "create_at": "2021-07-05 17:52:27.240000",
        "update_at": "2021-07-05 17:52:27.240000",
        "ttl": 300,
        "tenant_id": "bc8d2790fc9a46949818b942c0a824de",
        "zone_id": "5f01ffb6-ddb5-4541-b978-87a7eed71058",
        "data": [
            "1.1.1.1"
        ],
        "routing_policy_data": {}
    }
}
`
		_, _ = fmt.Fprint(writer, resp)
	})
	record, err := client.DNS.GetRecord(ctx, "5e6d0d98-ddd4-4d2f-87b0-aa60261cafd4")
	require.NoError(t, err)
	assert.Equal(t, 300, record.TTL)
}

func TestUpdateRecord(t *testing.T) {
	setup()
	defer teardown()
	var d dnsService
	mux.HandleFunc(testlib.DNSURL(d.recordItemPath("0ed9f98b-7991-4d49-929f-801f246d21f3")), func(writer http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		var payload *UpdateMXRecordPayload
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		resp := `{
  "id": "0ed9f98b-7991-4d49-929f-801f246d21f3",
  "name": "mx",
  "deleted": 0,
  "create_at": "2020-10-19 11:09:56.909000",
  "update_at": "2020-10-19 11:09:56.909000",
  "tenant_id": "333e2300b3c644cc93359a41a07c2321",
  "zone_id": "48d6ce71-43ed-45d3-9ab3-747dd08f500f",
  "type": "MX",
  "ttl": 300,
  "data": [
    {
      "value": "imap1.vccloud.vn",
      "priority": 20
    },
    {
      "value": "imap2.vccloud.vn",
      "priority": 2
    }
  ],
  "routing_policy_data": {}
}`
		_, _ = fmt.Fprint(writer, resp)
	})
	payload := UpdateMXRecordPayload{
		BaseUpdateRecordPayload: BaseUpdateRecordPayload{
			Name: "mx",
			Type: "MX",
			TTL:  300,
		},
		Data: []MXData{
			MXData{
				Value:    "imap1.vccloud.vn",
				Priority: 20,
			},
			MXData{
				Value:    "imap2.vccloud.vn",
				Priority: 2,
			},
		},
	}
	zone, err := client.DNS.UpdateRecord(ctx, "0ed9f98b-7991-4d49-929f-801f246d21f3", &payload)
	require.NoError(t, err)
	assert.Equal(t, "mx", zone.Name)
}

func TestDeleteRecord(t *testing.T) {
	setup()
	defer teardown()
	var d dnsService
	mux.HandleFunc(testlib.DNSURL(d.recordItemPath("81b3881a-5f11-432c-b4ec-2f57557d1c8c")), func(writer http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
	})
	require.NoError(t, client.DNS.DeleteRecord(ctx, "81b3881a-5f11-432c-b4ec-2f57557d1c8c"))
}
