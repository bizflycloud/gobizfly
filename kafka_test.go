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

func TestKafkaListVersion(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.KafkaURL("/versions"), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `{
			"data": [
				{"id": "v1", "name": "Kafka 2.8", "description": "Stable version"},
				{"id": "v2", "name": "Kafka 3.0", "description": "Latest version"}
			]
		}`
		_, _ = fmt.Fprint(w, resp)
	})

	versions, err := client.Kafka.ListVersion(ctx, &KafkaVersionListOptions{})
	require.NoError(t, err)
	assert.Len(t, versions, 2)
	assert.Equal(t, "v1", versions[0].ID)
	assert.Equal(t, "Kafka 2.8", versions[0].Name)
	assert.Equal(t, "v2", versions[1].ID)
	assert.Equal(t, "Kafka 3.0", versions[1].Name)
}

// Add more tests for Create, Delete, ListFlavor, Resize, AddNode as needed.

func TestKafkaGet(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.KafkaURL("/clusters/cluster-id"), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `{"data": {"id": "cluster-id", "name": "test-cluster", "version_id": "v1", "nodes": [], "flavor": "2c_4g", "volume_size": 10, "status": "Active", "created_at": "2025-12-19T16:15:29+07:00", "availability_zone": "HN1", "public_access": false, "obs": {"dashboard_url": "", "obs_id": "", "query_log": false}, "project_id": ""}}`
		_, _ = fmt.Fprint(w, resp)
	})

	cluster, err := client.Kafka.Get(ctx, "cluster-id")
	require.NoError(t, err)
	assert.Equal(t, "cluster-id", cluster.ID)
	assert.Equal(t, "test-cluster", cluster.Name)
}

func TestKafkaCreate(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.KafkaURL("/clusters"), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		resp := `{"data": "task-id-123"}`
		_, _ = fmt.Fprint(w, resp)
	})
	req := &KafkaInitClusterRequest{
		ClusterName: "test-cluster",
		VersionID:   "v1",
		Nodes:       3,
		Flavor:      "2c_4g",
		VolumeSize:  10,
	}
	task, err := client.Kafka.Create(ctx, req)
	require.NoError(t, err)
	assert.Equal(t, "task-id-123", task.TaskID)
}

func TestKafkaDelete(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.KafkaURL("/clusters/cluster-id"), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
		resp := `{"data": "task-id-del"}`
		_, _ = fmt.Fprint(w, resp)
	})
	task, err := client.Kafka.Delete(ctx, "cluster-id")
	require.NoError(t, err)
	assert.Equal(t, "task-id-del", task.TaskID)
}

func TestKafkaListFlavor(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.KafkaURL("/flavors"), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `{"data": [{"id": "f1", "name": "2c_4g", "vcpus": 2, "ram": 4, "disk": 100, "is_default": true}]}`
		_, _ = fmt.Fprint(w, resp)
	})

	flavors, err := client.Kafka.ListFlavor(ctx, &KafkaFlavorListOptions{})
	require.NoError(t, err)
	assert.Len(t, flavors, 1)
	assert.Equal(t, "f1", flavors[0].ID)
	assert.Equal(t, "2c_4g", flavors[0].Name)
}

func TestKafkaResize(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.KafkaURL("/clusters/cluster-id"), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPut, r.Method)
		resp := `{"data": "task-id-resize"}`
		_, _ = fmt.Fprint(w, resp)
	})
	req := &KafkaResizeClusterRequest{Flavor: "4c_8g", Type: "flavor"}
	task, err := client.Kafka.Resize(ctx, "cluster-id", req)
	require.NoError(t, err)
	assert.Equal(t, "task-id-resize", task.TaskID)
}

func TestKafkaAddNode(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.KafkaURL("/clusters/cluster-id/resize"), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		resp := `{"data": "task-id-addnode"}`
		_, _ = fmt.Fprint(w, resp)
	})
	req := &KafkaAddNodeRequest{Nodes: 1, Type: "increase"}
	task, err := client.Kafka.AddNode(ctx, "cluster-id", req)
	require.NoError(t, err)
	assert.Equal(t, "task-id-addnode", task.TaskID)
}
