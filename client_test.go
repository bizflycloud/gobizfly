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
	testRegion := "HaNoi"
	client, err = NewClient(WithAPIURL(serverTest.URL), WithRegionName(testRegion))
	services := []*Service{
		{Name: "Cloud Server",
			CanonicalName: serverServiceName,
			ServiceURL:    serverTest.URL + "/iaas-cloud/api",
			Region:        testRegion},
		{
			Name:          "Load Balancer",
			CanonicalName: loadBalancerServiceName,
			ServiceURL:    serverTest.URL + "/api/loadbalancers",
			Region:        testRegion,
		},
		{
			Name:          "Simple Storage",
			CanonicalName: simpleStorageServiceName,
			ServiceURL:    serverTest.URL + "/api/simple-storage",
			Region:        testRegion,
		},
		{
			Name:          "CloudWatcher",
			CanonicalName: cloudwatcherServiceName,
			ServiceURL:    serverTest.URL + "/api/alert",
			Region:        testRegion,
		},
		{
			Name:          "Auto Scaling",
			CanonicalName: autoScalingServiceName,
			ServiceURL:    serverTest.URL + "/api/auto-scaling",
			Region:        testRegion,
		},
		{
			Name:          "Accounts",
			CanonicalName: accountName,
			ServiceURL:    serverTest.URL + "/api/account",
			Region:        testRegion,
		},
		{
			Name:          "Auth",
			CanonicalName: authServiceName,
			ServiceURL:    serverTest.URL + "/api",
			Region:        testRegion,
		},
		{
			Name:          "Kubernetes",
			CanonicalName: kubernetesServiceName,
			ServiceURL:    serverTest.URL + "/api/kubernetes-engine",
			Region:        testRegion,
		},
		{
			Name:          "Container Registry",
			CanonicalName: containerRegistryName,
			ServiceURL:    serverTest.URL + "/api/container-registry",
			Region:        testRegion,
		},
		{
			Name:          "CDN",
			CanonicalName: cdnName,
			ServiceURL:    serverTest.URL + "/api/cdn",
			Region:        testRegion,
		},
		{
			Name:          "DNS",
			CanonicalName: dnsName,
			ServiceURL:    serverTest.URL + "/api/dns",
			Region:        testRegion,
		},
		{
			Name:          "Cloud Backup",
			CanonicalName: cloudBackupServiceName,
			ServiceURL:    serverTest.URL + "/api/cloud-backup",
			Region:        testRegion,
		},
		{
			Name:          "Database",
			CanonicalName: databaseServiceName,
			ServiceURL:    serverTest.URL + "/api/cloud-database",
			Region:        testRegion,
		},
		{
			Name:          "KMS",
			CanonicalName: kmsServiceName,
			ServiceURL:    serverTest.URL + "/api/ssl",
			Region:        testRegion,
		},
		{
			Name:          "Kafka",
			CanonicalName: kafkaServiceName,
			ServiceURL:    serverTest.URL + "/api/kafka/v1",
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
