// This file is part of gobizfly

package gobizfly

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/bizflycloud/gobizfly/testlib"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestToken(t *testing.T) {

	tests := []struct {
		name          string
		tcr           *TokenCreateRequest
		expectedToken string
	}{
		{"Auth Password", &TokenCreateRequest{AuthMethod: "password", Username: "foo@bizflycloud.vn", Password: "xxx"}, "auth-password-token"},
		{"Auth Application Secret", &TokenCreateRequest{AuthMethod: "application_credential", AppCredID: "174b36fd6c9e4a1da2e7c7dbddb89c69", AppCredSecret: "foo"}, "auth-app-token"},
		{"Auth Password with Project Name", &TokenCreateRequest{AuthMethod: "password", Username: "foo@bizflycloud.vn", Password: "xxx", ProjectName: "testIAM"}, "auth-password_IAM"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			setup()
			defer teardown()
			mux.HandleFunc(testlib.AuthURL(tokenPath), func(w http.ResponseWriter, r *http.Request) {
				require.Equal(t, http.MethodPost, r.Method)
				resp := fmt.Sprintf(`
{
    "token": "%s",
    "expires_at": "2019-11-22T15:39:54.000000Z"
}
`, tc.expectedToken)
				_, _ = fmt.Fprint(w, resp)
			})

			mux.HandleFunc(serviceUrl, func(w http.ResponseWriter, r *http.Request) {
				require.Equal(t, http.MethodGet, r.Method)
				resp := `
{
  "services": [
    {
      "canonical_name": "cloud_server", 
      "code": "CS", 
      "description": "Cloud Server", 
      "enabled": true, 
      "icon": "https://manage.bizflycloud.vn/iaas-cloud/api", 
      "id": 2, 
      "name": "Cloud Server", 
      "region": "HN", 
      "service_url": "https://manage.bizflycloud.vn/iaas-cloud/api"
    }
  ]
}`
				_, _ = fmt.Fprint(w, resp)
			})

			tok, err := client.Token.Create(ctx, tc.tcr)
			require.NoError(t, err)
			require.Equal(t, tc.expectedToken, tok.KeystoneToken)
		})
	}
}

func TestRetryWhenTokenExpired(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(testlib.AuthURL(tokenPath), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		resp := `
{
    "token": "xxx",
    "expires_at": "2019-11-22T15:39:54.000000Z"
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	var l loadbalancer
	mux.HandleFunc(testlib.LoadBalancerURL(l.resourcePath()), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		token := r.Header.Get("X-Auth-Token")
		if token != "xxx" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		resp := `
{
    "loadbalancers": []
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	mux.HandleFunc(serviceUrl, func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
  "services": [
    {
      "canonical_name": "cloud_server", 
      "code": "CS", 
      "description": "Cloud Server", 
      "enabled": true, 
      "icon": "https://manage.bizflycloud.vn/iaas-cloud/api", 
      "id": 2, 
      "name": "Cloud Server", 
      "region": "HN", 
      "service_url": "https://manage.bizflycloud.vn/iaas-cloud/api"
    }
  ]
}`
		_, _ = fmt.Fprint(w, resp)
	})

	client.keystoneToken = "yyy"
	lbs, err := client.LoadBalancer.List(ctx, &ListOptions{})
	require.NoError(t, err)
	assert.Len(t, lbs, 0)
	assert.Equal(t, "xxx", client.keystoneToken)
}
