// Copyright 2019 The BizFly Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gobizfly

import (
	"context"
	"net/http"
	"net/http/httptest"
)

var (
	ctx    = context.TODO()
	server *httptest.Server
	mux    *http.ServeMux
	client *Client
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	var err error
	client, err = NewClient(WithAPIUrl(server.URL))
	if err != nil {
		panic(err)
	}
}

func teardown() {
	server.Close()
}
