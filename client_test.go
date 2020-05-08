// Copyright 2019 The BizFly Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gobizfly

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	ctx        = context.TODO()
	serverTest *httptest.Server
	mux        *http.ServeMux
	client     *Client
)

func setup() {
	mux = http.NewServeMux()
	serverTest = httptest.NewServer(mux)

	var err error
	client, err = NewClient(WithAPIUrl(serverTest.URL))
	if err != nil {
		panic(err)
	}
}

func teardown() {
	serverTest.Close()
}

func TestErrFromStatus(t *testing.T) {
	err := errorFromStatus(404, "Volume not found")
	if !errors.Is(err, ErrNotFound) {
		t.Errorf("Error")
	}
	err = errorFromStatus(403, "Permission denied")
	if !errors.Is(err, ErrPermissionDenied) {
		t.Errorf("Error")
	}
	err = errorFromStatus(400, "Client error")
	if !errors.Is(err, ErrCommon) {
		t.Errorf("Error")
	}
}
