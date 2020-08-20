// This file is part of gobizfly
//
// Copyright (C) 2020  BizFly Cloud
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>

package gobizfly

import (
	"fmt"
	"github.com/bizflycloud/gobizfly/testlib"
	"net/http"
	"testing"

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

	client.keystoneToken = "yyy"
	lbs, err := client.LoadBalancer.List(ctx, &ListOptions{})
	require.NoError(t, err)
	assert.Len(t, lbs, 0)
	assert.Equal(t, "xxx", client.keystoneToken)
}
