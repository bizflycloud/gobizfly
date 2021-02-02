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
	"encoding/json"
	"fmt"
	"github.com/bizflycloud/gobizfly/testlib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"strings"
	"testing"
)

func TestClusterList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(testlib.K8sURL(clusterPath), func(writer http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
  "clusters": [
    {
      "uid": "ji84wqtzr77ogo6b",
      "name": "deploy_test",
      "version": {
        "id": "5f7d3a91d857155ad4993a32",
        "name": "v1.18.6-5f7d3a91",
        "description": null,
        "kubernetes_version": "v1.18.6"
      },
      "auto_upgrade": true,
      "tags": [],
      "provision_status": "PROVISIONED",
      "cluster_status": "PROVISIONED",
      "created_at": "2020-10-25T15:47:22.611000",
      "created_by": "svtt.tungds@vccloud.vn",
      "worker_pools_count": 1
    }
  ]
}
`
		_, _ = fmt.Fprint(writer, resp)
	})
	clusters, err := client.KubernetesEngine.List(ctx, &ListOptions{})
	require.NoError(t, err)
	assert.Len(t, clusters, 1)
	assert.LessOrEqual(t, "ji84wqtzr77ogo6b", clusters[0].UID)
}

func TestClusterGet(t *testing.T) {
	setup()
	defer teardown()
	var c kubernetesEngineService
	mux.HandleFunc(testlib.K8sURL(c.itemPath("ji84wqtzr77ogo6b")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
  "uid": "ji84wqtzr77ogo6b",
  "name": "deploy_test",
  "version": {
    "id": "5f7d3a91d857155ad4993a32",
    "name": "v1.18.6-5f7d3a91",
    "description": null,
    "kubernetes_version": "v1.18.6"
  },
  "auto_upgrade": true,
  "tags": [],
  "provision_status": "PROVISIONED",
  "cluster_status": "PROVISIONED",
  "created_at": "2020-10-25T15:47:22.611000",
  "created_by": "svtt.tungds@vccloud.vn",
  "worker_pools_count": 1,
  "worker_pools": [
    {
      "id": "5f959e0ac0b18944d0a4f13a",
      "name": "pool-curdcqn2",
      "version": "v1.18.6",
      "flavor": "3c_6g_enterprise",
      "flavor_detail": {
        "name": "3c_6g_enterprise",
        "vcpus": 3,
        "ram": 6144
      },
      "profile_type": "enterprise",
      "volume_type": "ENTERPRISE-SSD1",
      "volume_size": 50,
      "availability_zone": "HN1",
      "desired_size": 1,
      "enable_autoscaling": false,
      "min_size": 1,
      "max_size": 1,
      "tags": [],
      "provision_status": "PROVISIONED",
      "launch_config_id": "5a5c240f-e161-47c3-9214-b5ce7d2cda66",
      "autoscaling_group_id": "524cc8ce-4f88-47ce-ae16-a3b45e9dd89b",
      "created_at": "2020-10-25T15:47:22.608000"
    }
  ],
  "stats": {
    "worker_pools": 1,
    "total_cpu": null,
    "total_memory": null,
    "total_disk": {}
  }
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	cluster, err := client.KubernetesEngine.Get(ctx, "ji84wqtzr77ogo6b")
	require.NoError(t, err)
	assert.Equal(t, "deploy_test", cluster.Name)
}

func TestClusterCreate(t *testing.T) {
	setup()
	defer teardown()

	var c kubernetesEngineService
	mux.HandleFunc(testlib.K8sURL(c.resourcePath()), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		var payload *ClusterCreateRequest
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		assert.Equal(t, "my-kubernetes-cluster-1", payload.Name)

		resp := `
{
  "name": "my-kubernetes-cluster-1",
  "version": {"id": "5f7d3a91d857155ad4993a32", "name": "v1.18.6-5f7d3a91", "description": null, "kubernetes_version": "v1.18.6"},
  "auto_upgrade": false,
  "enable_cloud": true,
  "tags": [
    "string"
  ],
  "worker_pools": [
    {
      "name": "my-first-pool",
      "version": "v1.18.0",
      "flavor": "8c_8g",
      "profile_type": "premium",
      "volume_type": "SSD",
      "volume_size": 40,
      "availability_zone": "HN2",
      "desired_size": 1,
      "enable_autoscaling": true,
      "min_size": 1,
      "max_size": 3,
      "tags": [
        "string"
      ]
    }
  ]
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	cluster, err := client.KubernetesEngine.Create(ctx, &ClusterCreateRequest{
		Name: "my-kubernetes-cluster-1",
		Version: ControllerVersion{
			Name:       "v1.18.6-5f7d3a91",
			ID:         "5f7d3a91d857155ad4993a32",
			K8SVersion: "v1.18.6",
		},
		AutoUpgrade: false,
		EnableCloud: true,
		Tags:        []string{"string"},
		WorkerPools: []WorkerPool{
			{
				Name:              "my-first-pool",
				Version:           "v1.18.0",
				Flavor:            "8c_8g",
				ProfileType:       "premium",
				VolumeType:        "SSD",
				VolumeSize:        40,
				AvailabilityZone:  "HN2",
				DesiredSize:       1,
				EnableAutoScaling: true,
				MinSize:           1,
				MaxSize:           3,
				Tags:              []string{"string"},
			},
		},
	})
	require.NoError(t, err)
	assert.Equal(t, false, cluster.AutoUpgrade)
}

func TestClusterDelete(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(strings.Join([]string{testlib.K8sURL(clusterPath), "ji84wqtzr77ogo6b"}, "/"),
		func(writer http.ResponseWriter, request *http.Request) {
			require.Equal(t, http.MethodDelete, request.Method)
			writer.WriteHeader(http.StatusNoContent)
		})
	require.NoError(t, client.KubernetesEngine.Delete(ctx, "ji84wqtzr77ogo6b"))
}

func TestAddWorkerPool(t *testing.T) {
	setup()
	defer teardown()

	var c kubernetesEngineService
	mux.HandleFunc(testlib.K8sURL(c.itemPath("ji84wqtzr77ogo6b")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		var payload struct {
			WorkerPools *[]WorkerPool `json:"worker_pools"`
		}
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		resp := `
{
  "worker_pools": [
    {
      "id": "5eaf853b449fea8e4d0852e5",
      "name": "my-first-pool",
      "version": "v1.18.0",
      "flavor": "6c_6g",
      "flavor_detail": {
        "name": "6c_6g",
        "vcpus": 6,
        "ram": 6114
      },
      "profile_type": "premium",
      "volume_type": "SSD",
      "volume_size": 40,
      "availability_zone": "HN2",
      "desired_size": 1,
      "enable_autoscaling": true,
      "min_size": 1,
      "max_size": 3,
      "tags": [
        "string"
      ],
      "provision_status": "PROVISIONING",
      "launch_config_id": "2b1a00f1-28cd-4daf-b886-c0f07b401ed4",
      "autoscaling_group_id": "31e8465b-7275-4055-aeba-e3984453a223",
      "created_at": "2020-11-25T16:32:28.546Z"
    }
  ]
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	workerpools, err := client.KubernetesEngine.AddWorkerPools(ctx, "ji84wqtzr77ogo6b", &AddWorkerPoolsRequest{
		WorkerPools: []WorkerPool{
			{
				Name:              "my-first-pool",
				Version:           "v1.18.0",
				Flavor:            "8c_8g",
				ProfileType:       "premium",
				VolumeType:        "SSD",
				VolumeSize:        40,
				AvailabilityZone:  "HN2",
				DesiredSize:       1,
				EnableAutoScaling: true,
				MinSize:           1,
				MaxSize:           3,
				Tags:              []string{"string"},
			},
		},
	})

	require.NoError(t, err)
	assert.Equal(t, "my-first-pool", workerpools[0].Name)
}

func TestRecycleNode(t *testing.T) {
	setup()
	defer teardown()
	var c kubernetesEngineService
	mux.HandleFunc(testlib.K8sURL(strings.Join([]string{c.itemPath("ji84wqtzr77ogo6b"), "ji84wqtzr77ogo6b"}, "/")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPut, r.Method)
	})
}

func TestDeleteClusterWorkerPool(t *testing.T) {
	setup()
	defer teardown()
	var c kubernetesEngineService
	mux.HandleFunc(testlib.K8sURL(strings.Join([]string{c.itemPath("ji84wqtzr77ogo6b"), "ji84wqtzr77ogo6b"}, "/")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
	})
	err := client.KubernetesEngine.DeleteClusterWorkerPool(ctx, "ji84wqtzr77ogo6b", "ji84wqtzr77ogo6b")
	require.NoError(t, err)
}

func TestGetClusterWorkerPool(t *testing.T) {
	setup()
	defer teardown()
	var c kubernetesEngineService
	mux.HandleFunc(testlib.K8sURL(strings.Join([]string{c.itemPath("ji84wqtzr77ogo6b"), "5f959e0ac0b18944d0a4f13a"}, "/")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
  "id": "5f959e0ac0b18944d0a4f13a",
  "name": "pool-curdcqn2",
  "version": "v1.18.6",
  "flavor": "3c_6g_enterprise",
  "flavor_detail": {
    "name": "3c_6g_enterprise",
    "vcpus": 3,
    "ram": 6144
  },
  "profile_type": "enterprise",
  "volume_type": "ENTERPRISE-SSD1",
  "volume_size": 50,
  "availability_zone": "HN1",
  "desired_size": 1,
  "enable_autoscaling": false,
  "min_size": 1,
  "max_size": 1,
  "tags": [],
  "provision_status": "PROVISIONED",
  "launch_config_id": "5a5c240f-e161-47c3-9214-b5ce7d2cda66",
  "autoscaling_group_id": "524cc8ce-4f88-47ce-ae16-a3b45e9dd89b",
  "created_at": "2020-10-25T15:47:22.608000",
  "nodes": []
}
`
		_, _ = fmt.Fprint(w, resp)
	})
	workerpool, err := client.KubernetesEngine.GetClusterWorkerPool(ctx, "ji84wqtzr77ogo6b", "5f959e0ac0b18944d0a4f13a")
	require.NoError(t, err)
	assert.Equal(t, "pool-curdcqn2", workerpool.Name)
}

func TestDeleteClusterWorkerNode(t *testing.T) {
	setup()
	defer teardown()
	var c kubernetesEngineService
	mux.HandleFunc(testlib.K8sURL(strings.Join([]string{c.itemPath("ji84wqtzr77ogo6b"), "5f959e0ac0b18944d0a4f13a", "5f959e0ac0b18944d0a4f13a"}, "/")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
	})
	err := client.KubernetesEngine.DeleteClusterWorkerPoolNode(ctx, "ji84wqtzr77ogo6b", "5f959e0ac0b18944d0a4f13a", "5f959e0ac0b18944d0a4f13a")
	require.NoError(t, err)
}

func TestUpdateClusterWorkerPool(t *testing.T) {
	setup()
	defer teardown()
	var c kubernetesEngineService
	mux.HandleFunc(testlib.K8sURL(strings.Join([]string{c.itemPath("ji84wqtzr77ogo6b"), "5f959e0ac0b18944d0a4f13a"}, "/")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPatch, r.Method)
	})

	err := client.KubernetesEngine.UpdateClusterWorkerPool(ctx, "ji84wqtzr77ogo6b", "5f959e0ac0b18944d0a4f13a", &UpdateWorkerPoolRequest{
		DesiredSize:       1,
		EnableAutoScaling: true,
		MinSize:           4,
		MaxSize:           5,
	})
	require.NoError(t, err)
}
