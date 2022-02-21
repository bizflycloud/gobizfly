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
  "cluster": {
    "uid": "nkieuavgndkdxwsqc",
    "name": "my-kubernetes-cluster-1",
    "version": {
      "id": "5f6425f3d0d3befd40e7a31f",
      "name": "v1.18.6-bke-5f6425f3",
      "description": "Kubernetes v1.18.6 on BizFly Cloud",
      "kubernetes_version": "v1.18.6"
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
	resp, err := client.KubernetesEngine.GetKubeConfig(ctx, "xfbxsws38dcs8o94")
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
