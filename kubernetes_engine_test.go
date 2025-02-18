// This file is part of gobizfly

package gobizfly

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/bizflycloud/gobizfly/testlib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClusterList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(testlib.K8sURL(clusterPath+"/"), func(writer http.ResponseWriter, r *http.Request) {
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
      "package": {
        "id": "6609972809ba00eb5adc95e6",
        "name": "STANDARD - 1",
        "max_nodes": 50,
        "memory": 8,
        "high_availability": true,
        "sla": "99.99",
        "price": {
            "amount": 1386000,
            "billing_cycle": "month"
        },
        "specify": "standard"
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
  "package": {
    "id": "6609972809ba00eb5adc95e6",
    "name": "STANDARD - 1",
    "max_nodes": 50,
    "memory": 8,
    "high_availability": true,
    "sla": "99.99",
    "price": {
        "amount": 1386000,
        "billing_cycle": "month"
    },
    "specify": "standard"
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
  "cluster": {
    "uid": "nkieuavgndkdxwsqc",
    "name": "my-kubernetes-cluster-1",
    "version": {
      "id": "5f6425f3d0d3befd40e7a31f",
      "name": "v1.18.6-bke-5f6425f3",
      "description": "Kubernetes v1.18.6 on Bizfly Cloud",
      "kubernetes_version": "v1.18.6"
    },
    "package": {
      "id": "6609972809ba00eb5adc95e6",
      "name": "STANDARD - 1",
      "max_nodes": 50,
      "memory": 8,
      "high_availability": true,
      "sla": "99.99",
      "price": {
          "amount": 1386000,
          "billing_cycle": "month"
      },
      "specify": "standard"
    },
    "private_network_id": "727caa8c-1ed1-4302-b659-5a92864dcdef",
    "auto_upgrade": true,
    "tags": [
      "string"
    ],
    "provision_status": "PROVISIONING",
    "cluster_status": "PROVISIONING",
    "created_at": "2021-02-22T12:53:09.361Z",
    "created_by": "thanhpm@vccloud.vn",
    "worker_pools_count": 3,
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
        "created_at": "2021-02-22T12:53:09.362Z"
      }
    ]
  }
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	cluster, err := client.KubernetesEngine.Create(ctx, &ClusterCreateRequest{
		Name:        "my-kubernetes-cluster-1",
		Version:     "v1.18.6-5f7d3a91",
		AutoUpgrade: true,
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
	assert.Equal(t, true, cluster.AutoUpgrade)
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

func TestClusterUpdate(t *testing.T) {
	setup()
	defer teardown()

	var c kubernetesEngineService
	mux.HandleFunc(testlib.K8sURL(c.itemPath("zdckkshu1he44fsb")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPatch, r.Method)
		var payload *UpdateClusterRequest
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		assert.Equal(t, true, *payload.AutoUpgrade)
		resp := `{
      "uid": "zdckkshu1he44fsb",
      "name": "test--cluster",
      "version": {
        "id": "6528edeac362ab8b6db3f795",
        "name": "v1.28.2-6528edea",
        "description": null,
        "kubernetes_version": "v1.28.2",
        "local_dns_supported": true,
        "cni_plugin_supported": [
          "cilium",
          "kube-router"
        ]
      },
      "package": {
        "id": "6609972809ba00eb5adc95e6",
        "name": "STANDARD - 1",
        "max_nodes": 50,
        "memory": 8,
        "high_availability": true,
        "sla": "99.99",
        "price": {
            "amount": 1386000,
            "billing_cycle": "month"
        },
        "specify": "standard"
      },
      "provision_type": "standard",
      "private_network_id": "2d70b001-5634-4dbe-8c3e-fe75a742a1ec",
      "private_subnet_id": "1b9816a6-b647-4970-92ff-4996d39f5192",
      "access_mode": "PUBLIC",
      "access_policies": [],
      "auto_upgrade": true,
      "tags": [],
      "provision_status": "PROVISIONING",
      "cluster_status": "PROVISIONING",
      "created_at": "2024-03-24T17:41:02.314000",
      "created_by": "ducnguyenvan99@bizflycloud.vn",
      "worker_pools_count": 1,
      "migrations_count": 0,
      "kubelet_communication": "INTERNAL",
      "upgrade_time": {
        "day": -1,
        "time": "17:00:00"
      },
      "force_upgrade": true,
      "local_dns": false,
      "bcr_integrated": false,
      "message": "",
      "cni_plugin": "kube-router",
      "package": {
        "id": "65dc09c51d722ebf15c48b91",
        "name": "STANDARD - 1",
        "max_nodes": 50,
        "memory": 8,
        "high_availability": true,
        "sla": "99.99",
        "price": {
          "amount": 1260000,
          "billing_cycle": "month"
        },
        "specify": "standard"
      },
      "nodes_count": 1,
      "worker_pools": [
        {
          "id": "660065ae0a45091b62ee6eef",
          "name": "pool-az4i4byf",
          "version": "v1.28.2",
          "provision_type": "standard",
          "flavor": "nix.2c_2g",
          "profile_type": "premium",
          "volume_type": "PREMIUM-SSD1",
          "volume_size": 50,
          "availability_zone": "HN1",
          "desired_size": 1,
          "enable_autoscaling": false,
          "min_size": 1,
          "max_size": 1,
          "network_plan": "free_datatransfer",
          "billing_plan": "on_demand",
          "tags": [],
          "provision_status": "PENDING_PROVISION",
          "launch_config_id": null,
          "autoscaling_group_id": null,
          "created_at": "2024-03-24T17:41:02.268000",
          "labels": {},
          "taints": [],
          "flavor_detail": {
            "name": "nix.2c_2g",
            "vcpus": 2,
            "ram": 2048,
            "gpu": null,
            "category": "premium"
          }
        }
      ],
      "stats": {
        "worker_pools": 1,
        "total_cpu": null,
        "total_memory": null,
        "total_disk": {}
      }
    }`
		_, _ = fmt.Fprint(w, resp)
	})
	autoUpgrade := true
	payload := UpdateClusterRequest{
		AutoUpgrade: &autoUpgrade,
	}
	cluster, err := client.KubernetesEngine.UpdateCluster(ctx, "zdckkshu1he44fsb", &payload)
	require.NoError(t, err)
	assert.Equal(t, true, cluster.AutoUpgrade)
}

func TestGetUpgradeClusterVersion(t *testing.T) {
	setup()
	defer teardown()

	var c kubernetesEngineService
	path := strings.Join([]string{"zdckkshu1he44fsb", "upgrade"}, "/")
	mux.HandleFunc(testlib.K8sURL(c.itemPath(path)), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `{
      "upgrade": {
        "upgrade_to": "v1.28.2",
        "is_latest": true,
        "force_upgrade_time": null
      }
    }`
		_, _ = fmt.Fprint(w, resp)
	})
	resp, err := client.KubernetesEngine.GetUpgradeClusterVersion(ctx, "zdckkshu1he44fsb")
	require.NoError(t, err)
	assert.Equal(t, true, resp.IsLatest)
	assert.Equal(t, "v1.28.2", resp.UpgradeTo)
}

func TestUpgradeClusterVersion(t *testing.T) {
	setup()
	defer teardown()

	var c kubernetesEngineService
	path := strings.Join([]string{"zdckkshu1he44fsb", "upgrade"}, "/")
	mux.HandleFunc(testlib.K8sURL(c.itemPath(path)), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
	})
	err := client.KubernetesEngine.UpgradeClusterVersion(ctx, "zdckkshu1he44fsb", nil)
	require.NoError(t, err)
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
  "package": {
    "id": "6609972809ba00eb5adc95e6",
    "name": "STANDARD - 1",
    "max_nodes": 50,
    "memory": 8,
    "high_availability": true,
    "sla": "99.99",
    "price": {
        "amount": 1386000,
        "billing_cycle": "month"
    },
    "specify": "standard"
  },
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

func TestGetKubeConfig(t *testing.T) {
	setup()
	defer teardown()
	var c kubernetesEngineService
	mux.HandleFunc(testlib.K8sURL(strings.Join([]string{c.itemPath("xfbxsws38dcs8o94"), "kubeconfig"}, "/")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURuVENDQW9XZ0F3SUJBZ0lRZUhpWGZ0M0xRRVc1T2N5ZVZibVlvREFOQmdrcWhraUc5dzBCQVFzRkFEQjQKTVFzd0NRWURWUVFHRXdKV1RqRU9NQXdHQTFVRUNBd0ZTR0Z1YjJreERqQU1CZ05WQkFjTUJVaGhibTlwTVJRdwpFZ1lEVlFRS0RBdENhWHBHYkhsRGJHOTFaREVlTUJ3R0ExVUVDd3dWUW1sNlJteDVRMnh2ZFdSTGRXSmxjbTVsCmRHVnpNUk13RVFZRFZRUUREQXBMZFdKbGNtNWxkR1Z6TUI0WERUSXhNREl5TXpBeE1Ea3lORm9YRFRReE1ESXgKT0RBeE1Ea3lORm93ZURFTE1Ba0dBMVVFQmhNQ1ZrNHhEakFNQmdOVkJBZ01CVWhoYm05cE1RNHdEQVlEVlFRSApEQVZJWVc1dmFURVVNQklHQTFVRUNnd0xRbWw2Um14NVEyeHZkV1F4SGpBY0JnTlZCQXNNRlVKcGVrWnNlVU5zCmIzVmtTM1ZpWlhKdVpYUmxjekVUTUJFR0ExVUVBd3dLUzNWaVpYSnVaWFJsY3pDQ0FTSXdEUVlKS29aSWh2Y04KQVFFQkJRQURnZ0VQQURDQ0FRb0NnZ0VCQUxiVTI2K2tVblluN1RDMWkwY2w4WnBqVEx1eUFDZUp6SWV6Y0xXLwpjTlJnTGxIMlVkclB4eVZWWWFpQ3hHeCtMd01ET0xpSVh1SUNRRlJ3ZEFvL3l5OHErVWk2NTk4bndmQ05BblJlCjc2M2NidjFyV0N2d0hsRjlXSk9vYTd0MDNuSEhnbmsydm12dEpxaFF3VlkyWkJFZm51U1lIR0dvb21HMllEa0MKTko4VGNLdWVpb3l0QW1nbVkwdjZNdTlRVWx3ZytyRVdseDF3YUhlS1l1eHc4NDRGV3Y5bzBhaTBYMFNwUG84dwppRXROVDU3Uy9nSSt6SnExNktlRTRsVWlpZlpFTUxQVHlKRHJ5OHhUMzdXdkNRZmFjanBhVXowSHowQkdsZU1QCkd2WENQR0IrUXJZM29oZCsyOUxxSjZwTGwwQXU3ZzBPWUhDWlRJKytWVVlueDVFQ0F3RUFBYU1qTUNFd0RnWUQKVlIwUEFRSC9CQVFEQWdFR01BOEdBMVVkRXdFQi93UUZNQU1CQWY4d0RRWUpLb1pJaHZjTkFRRUxCUUFEZ2dFQgpBQ29PK09ZMkxpbE5pdFIreXc1MklLL3lBMVA0ZHVEVm5uSDhoQjFWQnpjMkE1OVE1OGc1cHUyS1BvZnJjak1lCjdqUzk1WVhOaGhuRG56alRZRGVJdGdFQVVJbTZES3lWK2M0VW1VZ2Q2d1MyeW9vYk5FSTdSK0c1cCtlSTQxb0QKYjVIcTlwUzlCTzJsbTgwRHNFM1lSSFJHaXBPTnRGNEswemlUZXl5ZHdIOHFDR05jaWpkY2NyMkRzNi9zblhyZQpLWVVHaFJ6VHhudTNMZTNPMWNrdVBBMmg1UGpiL1dmUGdXNkgrK0lRTG8wWnMzNzdEQlRnYVBHZG9lVGJ3eXdtCkE2Z2ZmbDhQbEdMU2RJeXFpb3RJSTBBd3F2M3RuQVR4VjhnV0ljczRMRDZFQnlwSTE3cnNNeHhVZkFwWkh5cFUKdWRpNEZ1QXVhQ1ZCWHlTTmYwUkozT3M9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0=
    server: https://xfbxsws38dcs8o94.api.k8saas.bizflycloud.vn
  name: bke-hanoi-tung491-test-k8s_23
contexts:
- context:
    cluster: bke-hanoi-tung491-test-k8s_23
    user: admin
  name: bke-hanoi-tung491-test-k8s_23
current-context: bke-hanoi-tung491-test-k8s_23
kind: Config
preferences: {}
users:
- name: admin
  user:
    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUR0ekNDQXArZ0F3SUJBZ0lRTmRlQWhsRUpRcWVIWmNJbHRENFcxakFOQmdrcWhraUc5dzBCQVFzRkFEQjQKTVFzd0NRWURWUVFHRXdKV1RqRU9NQXdHQTFVRUNBd0ZTR0Z1YjJreERqQU1CZ05WQkFjTUJVaGhibTlwTVJRdwpFZ1lEVlFRS0RBdENhWHBHYkhsRGJHOTFaREVlTUJ3R0ExVUVDd3dWUW1sNlJteDVRMnh2ZFdSTGRXSmxjbTVsCmRHVnpNUk13RVFZRFZRUUREQXBMZFdKbGNtNWxkR1Z6TUI0WERUSXhNRE14TWpBek1EZ3hNbG9YRFRJeE1ETXgKT1RBek1EZ3hNbG93ZGpFTE1Ba0dBMVVFQmhNQ1ZrNHhEakFNQmdOVkJBZ01CVWhoYm05cE1RNHdEQVlEVlFRSApEQVZJWVc1dmFURVhNQlVHQTFVRUNnd09jM2x6ZEdWdE9tMWhjM1JsY25NeEhqQWNCZ05WQkFzTUZVSnBla1pzCmVVTnNiM1ZrUzNWaVpYSnVaWFJsY3pFT01Bd0dBMVVFQXd3RllXUnRhVzR3Z2dFaU1BMEdDU3FHU0liM0RRRUIKQVFVQUE0SUJEd0F3Z2dFS0FvSUJBUURIL1p1WmN1Z2hkWWRCc1ZqMjJJU0NBVkxVd1JZR3l5UElYalQwZERwUApRSmJwRzdzc2hVeUdmOGxpL1NiZUw5N3Z6VnMvekVROFpsMGhFYlJBRzlVL21qY1JFbFFHVzhSK2JnRm43dXhmClJzamZFTlJ5b1UxZlJKSVluSy9MMGpyTlErdTZqUkZENlNheTdDM0ZyaWRVNUJJTXdLcDRQamYyQlB1aFoxdjMKZWF6UmZ4KzY1Ly8yMXRWdjNJS200NFFxcHhjdGtrSkRKamI4YlRYeC91ajZVdGIvSXBqUHNBTjliR1FKU3FydQpRckI1bTI2Y1VHVHNONHZNWTBRV0ZSa3B4a21aMXQrQjBYc3ZQZklldDRJOFRTa08wa09sTUlRZDJNbmtGeGdGCnNhWDA3R21Xcm9vSVI3MW91RjlBNUNwM2NZcm1sRnI0bFljNUkrQlV0VE96QWdNQkFBR2pQekE5TUE0R0ExVWQKRHdFQi93UUVBd0lGb0RBZEJnTlZIU1VFRmpBVUJnZ3JCZ0VGQlFjREFRWUlLd1lCQlFVSEF3SXdEQVlEVlIwVApBUUgvQkFJd0FEQU5CZ2txaGtpRzl3MEJBUXNGQUFPQ0FRRUFETnpzWVdFV00xbFgzUUNUNndORG5yVmdzUW9XCi9Va01KaXZOTmdGUWI0L3g1c0lWckZPQUN4cEdOdktwOTFmamN3Y3ZVSWxlYXNZeHFKdWQ0Vkw3bCt6Q203SVQKM21WajdrWGhxaUtsVGlxQnoxVG1jZFNsN1o0RE5mUGFWVDNxYVAwc0FyaGtFaU4rWVlheUZ0MjNtdGlZU0FkTAppV3hQUVJmK3pJbzl6dUoxTUZCc05JeDB6emJMSkM3Y2FOenM0Q2RTMWRSTEI3OEFsRTRtOCtWaGFLS0Y2SUthCjUyb0Z4ZmZjbzh5ZVhYcW0xTHVEeUN1L05Yd0RCRVUrV2wwKzV6TzUyWTBPRjZGQ2QwRWlQZ0hSUzAxejVQMC8KZTdCZTAwcjZSUUpKTkI0UjF6N1VoK0hlOUVRbTQvVTBsOUZxK3RxUFV4QXdtQlcxN0QvRXFFbEk0UT09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
    client-key-data: LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1JSUV2Z0lCQURBTkJna3Foa2lHOXcwQkFRRUZBQVNDQktnd2dnU2tBZ0VBQW9JQkFRREgvWnVaY3VnaGRZZEIKc1ZqMjJJU0NBVkxVd1JZR3l5UElYalQwZERwUFFKYnBHN3NzaFV5R2Y4bGkvU2JlTDk3dnpWcy96RVE4WmwwaApFYlJBRzlVL21qY1JFbFFHVzhSK2JnRm43dXhmUnNqZkVOUnlvVTFmUkpJWW5LL0wwanJOUSt1NmpSRkQ2U2F5CjdDM0ZyaWRVNUJJTXdLcDRQamYyQlB1aFoxdjNlYXpSZngrNjUvLzIxdFZ2M0lLbTQ0UXFweGN0a2tKREpqYjgKYlRYeC91ajZVdGIvSXBqUHNBTjliR1FKU3FydVFyQjVtMjZjVUdUc040dk1ZMFFXRlJrcHhrbVoxdCtCMFhzdgpQZklldDRJOFRTa08wa09sTUlRZDJNbmtGeGdGc2FYMDdHbVdyb29JUjcxb3VGOUE1Q3AzY1lybWxGcjRsWWM1CkkrQlV0VE96QWdNQkFBRUNnZ0VBVjhFRGhzaXg3UVNhTGd3NHdrL3RqUEl4dTJOaVcrYkZNOFdLclAxWEhMRjEKeHFIQmR0NmkzcDJ4NjNxemxHa2pCTXh5VHNNOTZkYnM1SGJWUmhBd2VYRWMycVBWTk5rTmxvQ0VvMnRtVXNSSApuZ0hQaHVFYWgwUWFheXhOd3p6alNuQ1VQazVxRmdkM1VLbHJ5RU1MeFNjeWVHQU9MU2IzL1Q3Z2YwbFFSSDFiCjFLZ3g3TVQyUURQeEdLclJkSStpbHprendHRzZLQzdIdGJnRDBkRjdHbVo4cW51TW5IVHdYb0JmaWZNdG1wcUUKNjVXVkl0Y05FMllHZHNWakxoSU9nMkt2TmhyaktiVm1rdUUydHNhbml2QlVDbkUyVE50YmwrNjcwSWZPSEJXbQpYVkoyQVB3a2twVHVvVlBYOHpwL3dxa25ZQXluS2NkeXg5cG5VR2loMlFLQmdRRHFIVWEwcjBEc3Y1UXFrS3NTClA4Z3ovUW9SRmZFblExdm1YcVlFakp4dVIyRm1ZOXVUSlFSVzU1aXlaajdBYmxiSGpoWlllclVSR0JtVTBpWUcKMmp1Wmt1V1ZFd210Mll5KzRLbmZLT2JWaU5NbFJPUUFTUGkraU5XZWVDc01lVFVjY05ScXFPck1wT3krZVVRZwptNkNDdDZyNzZSclRUK2xIWWVhdm5yMG4xUUtCZ1FEYXI3S21HQkF6WXdvbEZaOHZIcVJqT3VmNitpaXB2eGI5Cmh4L2ZKbmhsckZDTTJETzRxL0tYS2FaaHBkNzluVmVXQVJGTVhBNktUQW1VZWdnRnU4ZmQ2LzFmc005ZlpES0oKZVJhVmFiRUlqckduVXVwTGlDVHVSNzlZblpjNExTZTQyNmc1M2JzQWRhZGlNVDd4NCt3MzJPaEF3ODVmN0xLbgphanJ0ZVd6NVp3S0JnSGRzZmIvMzBsK3lqb3QwQnNBbFp5UVdCVWVYOE04OWppaWl5WDl5bHUydVhlSVVPRk1FClJBVnMyTGpRYlZ4T0xOaFpBODhZc1RyS0YycVNGTEhVS3lqNUJVSVpWd3VtK1NQNWlNMzhtRnYvRXU1bENRV1kKTThOR0crcGRsR0FsaUZFOHdTNnpnaXJvU3BnVFZneG9OdVhYZVZKTm85QjlhQnR3dG5PSnZ5WU5Bb0dCQUp3bgphNHk0a0JEeGZwUCtmWDE3QnUwb2FlL0g0M05hVlFOU0Vvc1lnRTR4bmc4RWJ1SkdQZUo4eGliaDkzbm5lVnhPCmhOaWV2MjgzWG52Y0s1QlVoeUpMV2RDVGczQmRMczBGWHYvdnlZOFB2WUY2Ym56aXlXUXdiVXpNc3VkVkx4RU0KSUhLNWhzZU1PNnFjK1pKbUt3Mng0QjRtODEyQnVneGJpWnA2NHpxdkFvR0JBTXVrcjVSSFdsL1YzNGY0Sk40aQpnaW5ZZVBJbG1XRDBDb2ZHSVUyTU1SNWhrbE9EZ2lnbzlOTzhjQTlUTVRsRDVDSitETG80TkZSSXFkTjhrb1UzCkJsWWdqdjY5dlVMak53TzdLVGdmT09YUUF6YUI3L2ZuR1llbnZzd0xzQnBRbWlqOFZuWDhzSklUbndXOU51cm4KVkRUTE1seDk3T3VqUEV5NVZQY2xoSk1TCi0tLS0tRU5EIFBSSVZBVEUgS0VZLS0tLS0K
`
		_, _ = fmt.Fprint(w, resp)
	})
	resp, err := client.KubernetesEngine.GetKubeConfig(ctx, "xfbxsws38dcs8o94", &GetKubeConfigOptions{})
	require.NoError(t, err)
	require.Equal(t, resp, `
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURuVENDQW9XZ0F3SUJBZ0lRZUhpWGZ0M0xRRVc1T2N5ZVZibVlvREFOQmdrcWhraUc5dzBCQVFzRkFEQjQKTVFzd0NRWURWUVFHRXdKV1RqRU9NQXdHQTFVRUNBd0ZTR0Z1YjJreERqQU1CZ05WQkFjTUJVaGhibTlwTVJRdwpFZ1lEVlFRS0RBdENhWHBHYkhsRGJHOTFaREVlTUJ3R0ExVUVDd3dWUW1sNlJteDVRMnh2ZFdSTGRXSmxjbTVsCmRHVnpNUk13RVFZRFZRUUREQXBMZFdKbGNtNWxkR1Z6TUI0WERUSXhNREl5TXpBeE1Ea3lORm9YRFRReE1ESXgKT0RBeE1Ea3lORm93ZURFTE1Ba0dBMVVFQmhNQ1ZrNHhEakFNQmdOVkJBZ01CVWhoYm05cE1RNHdEQVlEVlFRSApEQVZJWVc1dmFURVVNQklHQTFVRUNnd0xRbWw2Um14NVEyeHZkV1F4SGpBY0JnTlZCQXNNRlVKcGVrWnNlVU5zCmIzVmtTM1ZpWlhKdVpYUmxjekVUTUJFR0ExVUVBd3dLUzNWaVpYSnVaWFJsY3pDQ0FTSXdEUVlKS29aSWh2Y04KQVFFQkJRQURnZ0VQQURDQ0FRb0NnZ0VCQUxiVTI2K2tVblluN1RDMWkwY2w4WnBqVEx1eUFDZUp6SWV6Y0xXLwpjTlJnTGxIMlVkclB4eVZWWWFpQ3hHeCtMd01ET0xpSVh1SUNRRlJ3ZEFvL3l5OHErVWk2NTk4bndmQ05BblJlCjc2M2NidjFyV0N2d0hsRjlXSk9vYTd0MDNuSEhnbmsydm12dEpxaFF3VlkyWkJFZm51U1lIR0dvb21HMllEa0MKTko4VGNLdWVpb3l0QW1nbVkwdjZNdTlRVWx3ZytyRVdseDF3YUhlS1l1eHc4NDRGV3Y5bzBhaTBYMFNwUG84dwppRXROVDU3Uy9nSSt6SnExNktlRTRsVWlpZlpFTUxQVHlKRHJ5OHhUMzdXdkNRZmFjanBhVXowSHowQkdsZU1QCkd2WENQR0IrUXJZM29oZCsyOUxxSjZwTGwwQXU3ZzBPWUhDWlRJKytWVVlueDVFQ0F3RUFBYU1qTUNFd0RnWUQKVlIwUEFRSC9CQVFEQWdFR01BOEdBMVVkRXdFQi93UUZNQU1CQWY4d0RRWUpLb1pJaHZjTkFRRUxCUUFEZ2dFQgpBQ29PK09ZMkxpbE5pdFIreXc1MklLL3lBMVA0ZHVEVm5uSDhoQjFWQnpjMkE1OVE1OGc1cHUyS1BvZnJjak1lCjdqUzk1WVhOaGhuRG56alRZRGVJdGdFQVVJbTZES3lWK2M0VW1VZ2Q2d1MyeW9vYk5FSTdSK0c1cCtlSTQxb0QKYjVIcTlwUzlCTzJsbTgwRHNFM1lSSFJHaXBPTnRGNEswemlUZXl5ZHdIOHFDR05jaWpkY2NyMkRzNi9zblhyZQpLWVVHaFJ6VHhudTNMZTNPMWNrdVBBMmg1UGpiL1dmUGdXNkgrK0lRTG8wWnMzNzdEQlRnYVBHZG9lVGJ3eXdtCkE2Z2ZmbDhQbEdMU2RJeXFpb3RJSTBBd3F2M3RuQVR4VjhnV0ljczRMRDZFQnlwSTE3cnNNeHhVZkFwWkh5cFUKdWRpNEZ1QXVhQ1ZCWHlTTmYwUkozT3M9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0=
    server: https://xfbxsws38dcs8o94.api.k8saas.bizflycloud.vn
  name: bke-hanoi-tung491-test-k8s_23
contexts:
- context:
    cluster: bke-hanoi-tung491-test-k8s_23
    user: admin
  name: bke-hanoi-tung491-test-k8s_23
current-context: bke-hanoi-tung491-test-k8s_23
kind: Config
preferences: {}
users:
- name: admin
  user:
    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUR0ekNDQXArZ0F3SUJBZ0lRTmRlQWhsRUpRcWVIWmNJbHRENFcxakFOQmdrcWhraUc5dzBCQVFzRkFEQjQKTVFzd0NRWURWUVFHRXdKV1RqRU9NQXdHQTFVRUNBd0ZTR0Z1YjJreERqQU1CZ05WQkFjTUJVaGhibTlwTVJRdwpFZ1lEVlFRS0RBdENhWHBHYkhsRGJHOTFaREVlTUJ3R0ExVUVDd3dWUW1sNlJteDVRMnh2ZFdSTGRXSmxjbTVsCmRHVnpNUk13RVFZRFZRUUREQXBMZFdKbGNtNWxkR1Z6TUI0WERUSXhNRE14TWpBek1EZ3hNbG9YRFRJeE1ETXgKT1RBek1EZ3hNbG93ZGpFTE1Ba0dBMVVFQmhNQ1ZrNHhEakFNQmdOVkJBZ01CVWhoYm05cE1RNHdEQVlEVlFRSApEQVZJWVc1dmFURVhNQlVHQTFVRUNnd09jM2x6ZEdWdE9tMWhjM1JsY25NeEhqQWNCZ05WQkFzTUZVSnBla1pzCmVVTnNiM1ZrUzNWaVpYSnVaWFJsY3pFT01Bd0dBMVVFQXd3RllXUnRhVzR3Z2dFaU1BMEdDU3FHU0liM0RRRUIKQVFVQUE0SUJEd0F3Z2dFS0FvSUJBUURIL1p1WmN1Z2hkWWRCc1ZqMjJJU0NBVkxVd1JZR3l5UElYalQwZERwUApRSmJwRzdzc2hVeUdmOGxpL1NiZUw5N3Z6VnMvekVROFpsMGhFYlJBRzlVL21qY1JFbFFHVzhSK2JnRm43dXhmClJzamZFTlJ5b1UxZlJKSVluSy9MMGpyTlErdTZqUkZENlNheTdDM0ZyaWRVNUJJTXdLcDRQamYyQlB1aFoxdjMKZWF6UmZ4KzY1Ly8yMXRWdjNJS200NFFxcHhjdGtrSkRKamI4YlRYeC91ajZVdGIvSXBqUHNBTjliR1FKU3FydQpRckI1bTI2Y1VHVHNONHZNWTBRV0ZSa3B4a21aMXQrQjBYc3ZQZklldDRJOFRTa08wa09sTUlRZDJNbmtGeGdGCnNhWDA3R21Xcm9vSVI3MW91RjlBNUNwM2NZcm1sRnI0bFljNUkrQlV0VE96QWdNQkFBR2pQekE5TUE0R0ExVWQKRHdFQi93UUVBd0lGb0RBZEJnTlZIU1VFRmpBVUJnZ3JCZ0VGQlFjREFRWUlLd1lCQlFVSEF3SXdEQVlEVlIwVApBUUgvQkFJd0FEQU5CZ2txaGtpRzl3MEJBUXNGQUFPQ0FRRUFETnpzWVdFV00xbFgzUUNUNndORG5yVmdzUW9XCi9Va01KaXZOTmdGUWI0L3g1c0lWckZPQUN4cEdOdktwOTFmamN3Y3ZVSWxlYXNZeHFKdWQ0Vkw3bCt6Q203SVQKM21WajdrWGhxaUtsVGlxQnoxVG1jZFNsN1o0RE5mUGFWVDNxYVAwc0FyaGtFaU4rWVlheUZ0MjNtdGlZU0FkTAppV3hQUVJmK3pJbzl6dUoxTUZCc05JeDB6emJMSkM3Y2FOenM0Q2RTMWRSTEI3OEFsRTRtOCtWaGFLS0Y2SUthCjUyb0Z4ZmZjbzh5ZVhYcW0xTHVEeUN1L05Yd0RCRVUrV2wwKzV6TzUyWTBPRjZGQ2QwRWlQZ0hSUzAxejVQMC8KZTdCZTAwcjZSUUpKTkI0UjF6N1VoK0hlOUVRbTQvVTBsOUZxK3RxUFV4QXdtQlcxN0QvRXFFbEk0UT09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
    client-key-data: LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1JSUV2Z0lCQURBTkJna3Foa2lHOXcwQkFRRUZBQVNDQktnd2dnU2tBZ0VBQW9JQkFRREgvWnVaY3VnaGRZZEIKc1ZqMjJJU0NBVkxVd1JZR3l5UElYalQwZERwUFFKYnBHN3NzaFV5R2Y4bGkvU2JlTDk3dnpWcy96RVE4WmwwaApFYlJBRzlVL21qY1JFbFFHVzhSK2JnRm43dXhmUnNqZkVOUnlvVTFmUkpJWW5LL0wwanJOUSt1NmpSRkQ2U2F5CjdDM0ZyaWRVNUJJTXdLcDRQamYyQlB1aFoxdjNlYXpSZngrNjUvLzIxdFZ2M0lLbTQ0UXFweGN0a2tKREpqYjgKYlRYeC91ajZVdGIvSXBqUHNBTjliR1FKU3FydVFyQjVtMjZjVUdUc040dk1ZMFFXRlJrcHhrbVoxdCtCMFhzdgpQZklldDRJOFRTa08wa09sTUlRZDJNbmtGeGdGc2FYMDdHbVdyb29JUjcxb3VGOUE1Q3AzY1lybWxGcjRsWWM1CkkrQlV0VE96QWdNQkFBRUNnZ0VBVjhFRGhzaXg3UVNhTGd3NHdrL3RqUEl4dTJOaVcrYkZNOFdLclAxWEhMRjEKeHFIQmR0NmkzcDJ4NjNxemxHa2pCTXh5VHNNOTZkYnM1SGJWUmhBd2VYRWMycVBWTk5rTmxvQ0VvMnRtVXNSSApuZ0hQaHVFYWgwUWFheXhOd3p6alNuQ1VQazVxRmdkM1VLbHJ5RU1MeFNjeWVHQU9MU2IzL1Q3Z2YwbFFSSDFiCjFLZ3g3TVQyUURQeEdLclJkSStpbHprendHRzZLQzdIdGJnRDBkRjdHbVo4cW51TW5IVHdYb0JmaWZNdG1wcUUKNjVXVkl0Y05FMllHZHNWakxoSU9nMkt2TmhyaktiVm1rdUUydHNhbml2QlVDbkUyVE50YmwrNjcwSWZPSEJXbQpYVkoyQVB3a2twVHVvVlBYOHpwL3dxa25ZQXluS2NkeXg5cG5VR2loMlFLQmdRRHFIVWEwcjBEc3Y1UXFrS3NTClA4Z3ovUW9SRmZFblExdm1YcVlFakp4dVIyRm1ZOXVUSlFSVzU1aXlaajdBYmxiSGpoWlllclVSR0JtVTBpWUcKMmp1Wmt1V1ZFd210Mll5KzRLbmZLT2JWaU5NbFJPUUFTUGkraU5XZWVDc01lVFVjY05ScXFPck1wT3krZVVRZwptNkNDdDZyNzZSclRUK2xIWWVhdm5yMG4xUUtCZ1FEYXI3S21HQkF6WXdvbEZaOHZIcVJqT3VmNitpaXB2eGI5Cmh4L2ZKbmhsckZDTTJETzRxL0tYS2FaaHBkNzluVmVXQVJGTVhBNktUQW1VZWdnRnU4ZmQ2LzFmc005ZlpES0oKZVJhVmFiRUlqckduVXVwTGlDVHVSNzlZblpjNExTZTQyNmc1M2JzQWRhZGlNVDd4NCt3MzJPaEF3ODVmN0xLbgphanJ0ZVd6NVp3S0JnSGRzZmIvMzBsK3lqb3QwQnNBbFp5UVdCVWVYOE04OWppaWl5WDl5bHUydVhlSVVPRk1FClJBVnMyTGpRYlZ4T0xOaFpBODhZc1RyS0YycVNGTEhVS3lqNUJVSVpWd3VtK1NQNWlNMzhtRnYvRXU1bENRV1kKTThOR0crcGRsR0FsaUZFOHdTNnpnaXJvU3BnVFZneG9OdVhYZVZKTm85QjlhQnR3dG5PSnZ5WU5Bb0dCQUp3bgphNHk0a0JEeGZwUCtmWDE3QnUwb2FlL0g0M05hVlFOU0Vvc1lnRTR4bmc4RWJ1SkdQZUo4eGliaDkzbm5lVnhPCmhOaWV2MjgzWG52Y0s1QlVoeUpMV2RDVGczQmRMczBGWHYvdnlZOFB2WUY2Ym56aXlXUXdiVXpNc3VkVkx4RU0KSUhLNWhzZU1PNnFjK1pKbUt3Mng0QjRtODEyQnVneGJpWnA2NHpxdkFvR0JBTXVrcjVSSFdsL1YzNGY0Sk40aQpnaW5ZZVBJbG1XRDBDb2ZHSVUyTU1SNWhrbE9EZ2lnbzlOTzhjQTlUTVRsRDVDSitETG80TkZSSXFkTjhrb1UzCkJsWWdqdjY5dlVMak53TzdLVGdmT09YUUF6YUI3L2ZuR1llbnZzd0xzQnBRbWlqOFZuWDhzSklUbndXOU51cm4KVkRUTE1seDk3T3VqUEV5NVZQY2xoSk1TCi0tLS0tRU5EIFBSSVZBVEUgS0VZLS0tLS0K
`)
}

func TestGetClusterInfo(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.K8sURL(clusterInfo+"/"+"6515297b220963774dd304b0"), func(writer http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
    "shoot_uid": "n5s4coxhy30zwa0r",
    "pool_name": "pool-zd5n544x",
    "k8s_version": "v1.25.10",
    "worker_config": {
        "id": "65026d77ae05b43e97caf9a6",
        "version": "v1.25.10",
        "everywhere": true,
        "nvidiadevice": false,
        "CNI_VERSION": "v1.1.1",
        "RUNC_VERSION": "v1.1.4",
        "CONTAINERD_VERSION": "1.6.10",
        "KUBE_VERSION": "v1.25.10"
    }
}
`
		_, _ = fmt.Fprint(writer, resp)
	})
	cluster, err := client.KubernetesEngine.GetClusterInfo(ctx, "6515297b220963774dd304b0")
	require.NoError(t, err)
	assert.Equal(t, "n5s4coxhy30zwa0r", cluster.ShootUid)
	assert.Equal(t, "v1.25.10", cluster.K8sVersion)
}

func TestAddClusterEverywhere(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.K8sURL(clusterJoinEverywhere+"/"+"n5s4coxhy30zwa0r"), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		var payload *ClusterJoinEverywhereRequest
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		assert.Equal(t, "pool-zd5n544x-n5s4coxhy30zwa0r-go-test3", payload.Hostname)

		resp := `
{
    "apiserver": "https://n5s4coxhy30zwa0r.api.hn.bke-staging.bfcplatform.vn",
    "cluster_dns": "10.93.0.10",
    "cluster_cidr": "10.200.0.0/16",
    "cloud_provider": "external",
    "certificate": {
        "ca.pem": "-----BEGIN CERTIFICATE-----\nCACert\n-----END CERTIFICATE-----",
        "client-key.pem": "-----BEGIN PRIVATE KEY-----\nClientKey\n-----END PRIVATE KEY-----\n",
        "client.pem": "-----BEGIN CERTIFICATE-----\nClient\n-----END CERTIFICATE-----\n"
    },
    "max_pods": 110,
    "uuid": "fd4c01a4-9e35-48d8-a45a-7b724cf008ad",
    "reserved": {
        "system_reserved": {
            "cpu": "20m",
            "memory": "100Mi"
        },
        "kube_reserved": {
            "cpu": "50m",
            "memory": "394Mi"
        }
    }
}`
		_, _ = fmt.Fprint(w, resp)
	})

	cluster, err := client.KubernetesEngine.AddClusterEverywhere(ctx, "n5s4coxhy30zwa0r", &ClusterJoinEverywhereRequest{
		Hostname:    "pool-zd5n544x-n5s4coxhy30zwa0r-go-test3",
		IPAddresses: []string{"45.124.95.244"},
		Capacity: EverywhereNodeCapacity{
			Cores:    2,
			MemoryKB: 2024868,
		},
	})
	require.NoError(t, err)
	assert.Equal(t, "https://n5s4coxhy30zwa0r.api.hn.bke-staging.bfcplatform.vn", cluster.APIServer)
}

func TestGetNodeEverywhereByUUID(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.K8sURL(nodeEverywhere+"/"+"eecae8cd-e6e7-4ad9-b72e-89ac6a40dcbd"), func(writer http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
    "id": "651d3eced6ff4acf9de8b97d",
    "shoot": "n5s4coxhy30zwa0r",
    "pool_id": "6515297b220963774dd304b0",
    "node_name": "pool-zd5n544x-n5s4coxhy30zwa0r-go-test-3",
    "public_ip": "139.59.102.178",
    "uuid": "eecae8cd-e6e7-4ad9-b72e-89ac6a40dcbd",
    "created_at": "2023-10-04T10:30:38.229000",
    "updated_at": "2023-10-04T10:30:38.229000"
}
`
		_, _ = fmt.Fprint(writer, resp)
	})
	cluster, err := client.KubernetesEngine.GetEverywhere(ctx, "eecae8cd-e6e7-4ad9-b72e-89ac6a40dcbd")
	require.NoError(t, err)
	assert.Equal(t, "n5s4coxhy30zwa0r", cluster.Shoot)
	assert.Equal(t, "6515297b220963774dd304b0", cluster.PoolID)
}

func TestPackageStandardList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(testlib.K8sURL("/package/"), func(writer http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		// Validate the query parameter
		query := r.URL.Query()
		specify := query.Get("specify")
		require.Equal(t, "standard", specify, "The specify query parameter should be 'standard'")

		resp := `
{
  "packages":[
      {
        "id":"6609972809ba00eb5adc95e6",
        "name":"STANDARD - 1",
        "max_nodes":50,
        "memory":8,
        "high_availability":true,
        "sla":"99.99",
        "price":{
            "amount":1386000,
            "billing_cycle":"month"
        },
        "specify":"standard"
      },
      {
        "id":"65dd8d678de21c6b5c4da1e5",
        "name":"STANDARD - 0",
        "max_nodes":10,
        "memory":2,
        "high_availability":false,
        "sla":"99.00",
        "price":{
            "amount":0,
            "billing_cycle":"month"
        },
        "specify":"standard"
      }
  ],
  "page":1,
  "limit":10,
  "total":2
}
`
		_, _ = fmt.Fprint(writer, resp)
	})
	packages, err := client.KubernetesEngine.GetPackages(ctx, "standard")
	require.NoError(t, err)
	assert.Len(t, packages.Packages, 2)
	assert.LessOrEqual(t, "6609972809ba00eb5adc95e6", packages.Packages[0].ID)
}

func TestGetDetailWorkerPool(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.K8sURL(strings.Join([]string{workerPoolPath, "67b36fed16bd9672f0f01e78"}, "/")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
    {
      "id": "67b36fed16bd9672f0f01e78",
      "name": "pool-69645",
      "version": "v1.29.13",
      "provision_type": "standard",
      "flavor": "nix.2c_2g",
      "profile_type": "premium",
      "volume_type": "PREMIUM-HDD1",
      "volume_size": 40,
      "availability_zone": "HN1",
      "desired_size": 2,
      "enable_autoscaling": true,
      "min_size": 2,
      "max_size": 4,
      "network_plan": "free_datatransfer",
      "billing_plan": "on_demand",
      "tags": [
          "pool_tag",
          ""
      ],
      "provision_status": "PROVISIONED",
      "launch_config_id": "dde59999-2fa8-4463-9aa1-4896e5954344",
      "autoscaling_group_id": "c97d5376-c668-4ef9-8de5-8a92c291aee8",
      "created_at": "2025-02-17T17:20:45.405000",
      "labels": {
          "UpdateLabel": "UpdateLabelVal1"
      },
      "taints": [
          {
              "effect": "NoSchedule",
              "key": "UpdateTaint1",
              "value": "UpdateTaintVal1"
          }
      ],
      "auto_repair": false,
      "flavor_detail": {
          "name": "nix.2c_2g",
          "vcpus": 2,
          "ram": 2048,
          "gpu": null,
          "category": "premium"
      },
      "nodes": [
          {
              "id": "e365b68a-a365-49a8-b5b0-2b1fcdb6e250",
              "name": "pool-69645-x2d030sk95w69eia-node-Cv4N54hV",
              "physical_id": "de9cb2f0-c8f0-4bf0-a6ea-cf83d6b55ef7",
              "ip_addresses": [
                  "10.20.2.149"
              ],
              "status": "ACTIVE",
              "status_reason": "Creation succeeded"
          },
          {
              "id": "92000ba7-fb0e-4ae3-91d5-94d23760bd99",
              "name": "pool-69645-x2d030sk95w69eia-node-YPN8sXXf",
              "physical_id": "bc98b260-e047-4fee-a1c0-4dcf5f3d794b",
              "ip_addresses": [
                  "10.20.2.224"
              ],
              "status": "ACTIVE",
              "status_reason": "Creation succeeded"
          }
      ],
      "everywhere_nodes": [],
      "shoot_id": "x2d030sk95w69eia"
  }
`
		_, _ = fmt.Fprint(w, resp)
	})
	workerpool, err := client.KubernetesEngine.GetDetailWorkerPool(ctx, "67b36fed16bd9672f0f01e78")
	require.NoError(t, err)
	assert.Equal(t, "pool-69645", workerpool.Name)
	assert.Equal(t, "67b36fed16bd9672f0f01e78", workerpool.UID)
	assert.Equal(t, "x2d030sk95w69eia", workerpool.ShootID)
	assert.Equal(t, "PREMIUM-HDD1", workerpool.VolumeType)
	assert.Equal(t, 40, workerpool.VolumeSize)
	assert.Equal(t, 2, workerpool.MinSize)
	assert.Equal(t, 4, workerpool.MaxSize)
	assert.Equal(t, 2, workerpool.DesiredSize)
	assert.Equal(t, "free_datatransfer", workerpool.NetworkPlan)
	assert.Equal(t, "on_demand", workerpool.BillingPlan)
	assert.Equal(t, "HN1", workerpool.AvailabilityZone)
	assert.Equal(t, "premium", workerpool.ProfileType)
	assert.Equal(t, "nix.2c_2g", workerpool.Flavor)
	assert.Equal(t, "v1.29.13", workerpool.Version)
	assert.Equal(t, "standard", workerpool.ProvisionType)
}
