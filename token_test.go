// Copyright 2019 The BizFly Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gobizfly

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTokenAuthPassword(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(tokenPath, func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		resp := `
{
    "token": "xxx",
    "expires_at": "2019-11-22T15:39:54.000000Z"
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	tok, err := client.Token.Create(ctx, &TokenCreateRequest{AuthMethod: "password", Username: "foo@bizflycloud.vn", Password: "xxx"})
	require.NoError(t, err)
	require.Equal(t, "xxx", tok.KeystoneToken)
}

func TestTokenAuthApplicationCredential(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(tokenPath, func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		resp := `
{
    "token": "xxx",
    "expires_at": "2019-11-22T15:39:54.000000Z"
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	tok, err := client.Token.Create(ctx, &TokenCreateRequest{AuthMethod: "application_credential", AppCredID: "174b36fd6c9e4a1da2e7c7dbddb89c69", AppCredSecret: "xxx"})
	require.NoError(t, err)
	require.Equal(t, "xxx", tok.KeystoneToken)
}

func TestRetryWhenTokenExpired(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(tokenPath, func(w http.ResponseWriter, r *http.Request) {
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
	mux.HandleFunc(l.resourcePath(), func(w http.ResponseWriter, r *http.Request) {
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
