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

func TestKafkaClusterList(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.KafkaURL("/clusters"), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Equal(t, "name=cluster-1", r.URL.RawQuery)
		resp := `{
	"success": true,
	"error_code": 0,
	"message": "success",
	"data": [
	    {
	        "id": "6bdd08ee-02e3-4230-a12e-a1aecd8fac02",
	        "name": "cluster-1",
	        "version_id": "4e7d3d70-ccd3-4e6e-af0b-88722037d33d",
	        "nodes": 3,
	        "flavor": "2c_4g",
	        "volume_size": 10,
	        "status": "Active",
	        "created_at": "2025-12-19T16:15:29+07:00",
	        "availability_zone": "HN1",
	        "vpc_network_id": "",
	        "public_access": false,
	        "obs": {
	            "dashboard_url": "",
	            "obs_id": "",
	            "query_log": false
	        },
	        "project_id": ""
	    }
	],
	"metadata": {
	    "has_next": false,
	    "current_page": 1,
	    "has_previous": 0,
	    "total": 1,
	    "previous_page": 0,
	    "next_page": 0,
	    "page": 1
	}
}`
		_, _ = fmt.Fprint(w, resp)
	})

	clusters, err := client.Kafka.List(ctx, &KafkaClusterListOptions{
		Name: "cluster-1",
	})
	require.NoError(t, err)
	cluster := clusters[0]
	assert.Equal(t, "6bdd08ee-02e3-4230-a12e-a1aecd8fac02", cluster.ID)
	assert.Equal(t, "cluster-1", cluster.Name)
	assert.Equal(t, "4e7d3d70-ccd3-4e6e-af0b-88722037d33d", cluster.KafkaVersion)
	assert.Equal(t, 3, cluster.Nodes)
	assert.Equal(t, "2c_4g", cluster.Flavor)
	assert.Equal(t, 10, cluster.VolumeSize)
	assert.Equal(t, "Active", cluster.Status)
	assert.Equal(t, "2025-12-19T16:15:29+07:00", cluster.CreatedAt)
	assert.Equal(t, "HN1", cluster.AvailabilityZone)
	assert.Equal(t, "", cluster.VPCNetworkID)
	assert.Equal(t, false, cluster.PublicAccess)
	assert.Equal(t, "", cluster.OBS.DashboardURL)
	assert.Equal(t, "", cluster.OBS.OBSID)
	assert.Equal(t, false, cluster.OBS.QueryLog)
	assert.Equal(t, "", cluster.ProjectID)
}
