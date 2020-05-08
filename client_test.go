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
	tests := []struct {
		statusCode int
		msg        string
		err        error
	}{
		{http.StatusBadRequest, "Volume not found", ErrCommon},
		{http.StatusNotFound, "Permission denied", ErrNotFound},
		{http.StatusForbidden, "Generic error", ErrPermissionDenied},
	}

	for _, tc := range tests {
		if err := errorFromStatus(tc.statusCode, tc.msg); !errors.Is(err, tc.err) {
			t.Errorf("unexpected error, want: %v, got: %v", tc.err, err)
		}
	}
}
