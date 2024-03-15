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

func TestDomainList(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.CDNURL(strings.Join([]string{usersPath, domainPath}, "")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
    "total": 1,
    "pages": 1,
    "next": "/users/0b722b886f0d43f49e69e4648684c0b7/domains?page=1&limit=50",
    "prev": "/users/0b722b886f0d43f49e69e4648684c0b7/domains?page=1&limit=50",
    "results": [
        {
            "certificate": null,
            "cname": "",
            "slug": "autopro",
            "pagespeed": 0,
            "upstream_proto": "http",
            "ddos_protection": 0,
            "status": "ACTIVE",
            "created_at": "2016-06-06T00:00:00+00:00",
            "domain_id": "2a535a9d-b963-4148-a522-e68c10b3d337",
            "domain": "autopro.com.vn",
            "type": "CNAME"
        }
    ]
}
`
		_, _ = fmt.Fprint(w, resp)
	})
	resp, err := client.CDN.List(ctx, &ListOptions{})
	require.NoError(t, err)
	assert.Len(t, resp.Domains, 1)
	assert.Equal(t, resp.Total, 1)
	assert.Equal(t, resp.Domains[0].DomainID, "2a535a9d-b963-4148-a522-e68c10b3d337")
}

func TestDomainGet(t *testing.T) {
	setup()
	defer teardown()
	var c cdnService
	mux.HandleFunc(testlib.CDNURL(c.itemPath("3a9e0da6-3c2c-4856-b120-b54e944b352f")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
    "domain": {
        "cname": "",
        "upstream_proto": "http",
        "certificate": null,
        "secure_link": 0,
        "created_at": "2016-07-19T00:00:00+00:00",
        "last_24h_usage": 0,
        "id": 2755,
        "upstream_addrs": "192.168.6.37,192.168.6.89,10.5.20.105,10.5.20.107,10.5.20.110,10.5.20.92",
        "status": "ACTIVE",
        "upstream_host": "cafefcdn.com",
        "slug": "cafefcdn",
        "ddos_protection": 0,
        "secret_key": "27e516352c8d58d0ed1597ee08843d26",
        "domain_cdn": "cafefcdn.edge.vccloud.vn",
        "origin_addrs": [
            {
                "type": "ip",
                "host": "192.168.6.37"
            }
        ],
        "domain_id": "3a9e0da6-3c2c-4856-b120-b54e944b352f",
        "domain": "cafefcdn.com",
        "host_id": "99700105af1edb30ac17c3cf9ea6a165",
        "pagespeed": 1,
        "type": "REWRITE",
        "user": 57825
    }
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	domain, err := client.CDN.Get(ctx, "3a9e0da6-3c2c-4856-b120-b54e944b352f")
	require.NoError(t, err)
	assert.Equal(t, "cafefcdn.com", domain.Domain)
}

func TestDomainCreate(t *testing.T) {
	setup()
	defer teardown()
	var c cdnService
	mux.HandleFunc(testlib.CDNURL(c.resourcePath()), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		var payload *CreateDomainPayload
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		assert.Equal(t, "cdn.monkidia.com", payload.Domain)
		resp := `
{
    "message": "domain created",
    "domain": {
        "domain_id": "53e220ff-9aab-4cd3-935b-36ba8b20ded8",
        "domain": "cdn.monkidia.com",
        "slug": "cdn-monkidia"
    }
}
`
		_, _ = fmt.Fprint(w, resp)
	})
	resp, err := client.CDN.Create(ctx, &CreateDomainPayload{
		Domain: "cdn.monkidia.com",
		Origin: &Origin{
			Name:          "monkidia",
			UpstreamHost:  "www.huylvt.com",
			UpstreamAddrs: "www.huylvt.com",
			UpstreamProto: "http",
		},
	})
	require.NoError(t, err)
	require.Equal(t, "domain created", resp.Message)
	require.Equal(t, "cdn-monkidia", resp.Domain.Slug)
}

func TestDomainUpdate(t *testing.T) {
	setup()
	defer teardown()
	var c cdnService
	mux.HandleFunc(testlib.CDNURL(c.itemPath("3a9e0da6-3c2c-4856-b120-b54e944b352f")), func(writer http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPut, r.Method)
		var payload *UpdateDomainPayload
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		assert.Equal(t, "cafefcdn.com", payload.Origin.UpstreamHost)
		resp := `{
    "message": "Domain updated",
    "domain": {
        "cname": "",
        "upstream_proto": "http",
        "certificate": null,
        "secure_link": 0,
        "created_at": "2016-07-19T00:00:00+00:00",
        "last_24h_usage": 0,
        "id": 2755,
        "upstream_addrs": "192.168.6.37,192.168.6.89,10.5.20.105,10.5.20.107,10.5.20.110,10.5.20.92",
        "status": "ACTIVE",
        "upstream_host": "cafefcdn.com",
        "slug": "cafefcdn",
        "ddos_protection": 0,
        "secret_key": "27e516352c8d58d0ed1597ee08843d26",
        "domain_cdn": "cafefcdn.edge.vccloud.vn",
        "origin_addrs": [
            {
                "type": "ip",
                "host": "192.168.6.37"
            },
            {
                "type": "ip",
                "host": "192.168.6.89"
            },
            {
                "type": "ip",
                "host": "10.5.20.105"
            },
            {
                "type": "ip",
                "host": "10.5.20.107"
            },
            {
                "type": "ip",
                "host": "10.5.20.110"
            },
            {
                "type": "ip",
                "host": "10.5.20.92"
            }
        ],
        "domain_id": "3a9e0da6-3c2c-4856-b120-b54e944b352f",
        "domain": "cafefcdn.com",
        "host_id": "99700105af1edb30ac17c3cf9ea6a165",
        "pagespeed": 1,
        "type": "REWRITE",
        "user": 57825
    }
}
`
		_, _ = fmt.Fprint(writer, resp)
	})
	resp, err := client.CDN.Update(ctx, "3a9e0da6-3c2c-4856-b120-b54e944b352f", &UpdateDomainPayload{
		Origin: &Origin{
			Name:         "cafefcdn",
			UpstreamHost: "cafefcdn.com",
		},
	})
	require.NoError(t, err)
	assert.Equal(t, "99700105af1edb30ac17c3cf9ea6a165", resp.Domain.HostID)
	assert.Equal(t, "Domain updated", resp.Message)
}

func TestDeleteDomain(t *testing.T) {
	setup()
	defer teardown()
	var c cdnService
	mux.HandleFunc(testlib.CDNURL(c.itemPath("3a9e0da6-3c2c-4856-b120-b54e944b352f")),
		func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, http.MethodDelete, r.Method)
		})
	require.NoError(t, client.CDN.Delete(ctx, "3a9e0da6-3c2c-4856-b120-b54e944b352f"))
}

func TestDeleteCache(t *testing.T) {
	setup()
	defer teardown()
	var c cdnService
	mux.HandleFunc(testlib.CDNURL(c.itemPath("3a9e0da6-3c2c-4856-b120-b54e944b352f")),
		func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, http.MethodDelete, r.Method)
			var payload *Files
			require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		})
	files := Files{
		Files: []string{"/css/style.js", "/js/script.js", "/images/logo.jpg"}}
	require.NoError(t, client.CDN.DeleteCache(ctx, "3a9e0da6-3c2c-4856-b120-b54e944b352f", &files))
}
