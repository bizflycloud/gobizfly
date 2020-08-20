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
	client, err = NewClient(WithAPIUrl(serverTest.URL), WithRegionName("HN"))

	services := []*Service{
		&Service{Name: "Cloud Server",
			CanonicalName: "cloud_server",
			ServiceUrl: "https://manage.bizflycloud.vn/iaas-cloud/api",
			Region: "HN"},
			&Service{
			Name: "Load Balancer",
			CanonicalName: "load_balancer",
			ServiceUrl: "https://manage.bizflycloud.vn/api/loadbalancers"},
	}
	client.services = services
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
