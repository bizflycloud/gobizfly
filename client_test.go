// This file is part of gobizfly

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
	testRegion := "HN"
	client, err = NewClient(WithAPIUrl(serverTest.URL), WithRegionName(testRegion))
	services := []*Service{
		{Name: "Cloud Server",
			CanonicalName: serverServiceName,
			ServiceUrl:    serverTest.URL + "/iaas-cloud/api",
			Region:        testRegion},
		{
			Name:          "Load Balancer",
			CanonicalName: loadBalancerServiceName,
			ServiceUrl:    serverTest.URL + "/api/loadbalancers",
			Region:        testRegion,
		},
		{
			Name:          "CloudWatcher",
			CanonicalName: cloudwatcherServiceName,
			ServiceUrl:    serverTest.URL + "/api/alert",
			Region:        testRegion,
		},
		{
			Name:          "Auto Scaling",
			CanonicalName: autoScalingServiceName,
			ServiceUrl:    serverTest.URL + "/api/auto-scaling",
			Region:        testRegion,
		},
		{
			Name:          "Accounts",
			CanonicalName: accountName,
			ServiceUrl:    serverTest.URL + "/api/account",
			Region:        testRegion,
		},
		{
			Name:          "Auth",
			CanonicalName: authServiceName,
			ServiceUrl:    serverTest.URL + "/api",
			Region:        testRegion,
		},
		{
			Name:          "Kubernetes",
			CanonicalName: kubernetesServiceName,
			ServiceUrl:    serverTest.URL + "/api/kubernetes-engine",
			Region:        testRegion,
		},
		{
			Name:          "Container Registry",
			CanonicalName: containerRegistryName,
			ServiceUrl:    serverTest.URL + "/api/container-registry",
			Region:        testRegion,
		},
		{
			Name:          "CDN",
			CanonicalName: cdnName,
			ServiceUrl:    serverTest.URL + "/api/cdn",
			Region:        testRegion,
		},
		{
			Name:          "DNS",
			CanonicalName: dnsName,
			ServiceUrl:    serverTest.URL + "/api/dns",
			Region:        testRegion,
		},
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
