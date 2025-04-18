// This file is part of gobizfly

package gobizfly

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	clusterPath           = "/_"
	kubeConfig            = "kubeconfig"
	k8sVersion            = "/k8s_versions"
	clusterInfo           = "/engine/cluster_info"
	clusterJoinEverywhere = "/engine/cluster_join_everywhere"
	nodeEverywhere        = "/_/node_everywhere"
	k8sPackages           = "/package/"
	workerPoolPath        = "/worker_pool"
)

var _ KubernetesEngineService = (*kubernetesEngineService)(nil)

type kubernetesEngineService struct {
	client *Client
}

type KubernetesEngineService interface {
	List(ctx context.Context, opts *ListOptions) ([]*Cluster, error)
	Create(ctx context.Context, req *ClusterCreateRequest) (*ExtendedCluster, error)
	Get(ctx context.Context, id string) (*FullCluster, error)
	Delete(ctx context.Context, id string) error
	AddWorkerPools(ctx context.Context, id string, awp *AddWorkerPoolsRequest) ([]*ExtendedWorkerPool, error)
	RecycleNode(ctx context.Context, clusterUID string, PoolID string, NodePhysicalID string) error
	DeleteClusterWorkerPool(ctx context.Context, clusterUID string, PoolID string) error
	GetClusterWorkerPool(ctx context.Context, clusterUID string, PoolID string) (*WorkerPoolWithNodes, error)
	UpdateClusterWorkerPool(ctx context.Context, clusterUID string, PoolID string, uwp *UpdateWorkerPoolRequest) error
	DeleteClusterWorkerPoolNode(ctx context.Context, clusterUID string, PoolID string, NodeID string) error
	ForcedDeleteClusterWorkerPoolNode(ctx context.Context, clusterUID string, PoolID string, NodeID string) error
	GetKubeConfig(ctx context.Context, clusterUID string, opts *GetKubeConfigOptions) (string, error)
	GetKubernetesVersion(ctx context.Context, opts GetKubernetesVersionOpts) (*KubernetesVersionResponse, error)
	GetClusterInfo(ctx context.Context, pool_id string) (*ClusterInfoResponse, error)
	AddClusterEverywhere(ctx context.Context, id string, cjer *ClusterJoinEverywhereRequest) (*ClusterJoinEverywhereResponse, error)
	GetEverywhere(ctx context.Context, id string) (*EverywhereNode, error)
	UpdateCluster(ctx context.Context, id string, payload *UpdateClusterRequest) (*ExtendedCluster, error)
	UpgradeClusterVersion(ctx context.Context, id string, payload *UpgradeClusterVersionRequest) error
	GetUpgradeClusterVersion(ctx context.Context, id string) (*UpgradeClusterVersionResponse, error)
	GetPackages(ctx context.Context, provisionType string) (*KubernetesPackagesResponse, error)
	GetDetailWorkerPool(ctx context.Context, PoolID string) (*WorkerPoolWithNodes, error)
	UpgradePackage(ctx context.Context, id string, payload *UpgradePackageRequest) error
	GetDashboardURL(ctx context.Context, id string) (string, error)
	InstallAddon(ctx context.Context, id string, addonType string) error
	UninstallAddon(ctx context.Context, id string, addonType string) error
	GetAddonStatus(ctx context.Context, id string, addonType string) (*AddonStatusResponse, error)
}

// KubernetesVersionResponse represents the get versions from the Kubernetes Engine API
type KubernetesVersionResponse struct {
	ControllerVersions []ControllerVersion `json:"controller_versions"`
	WorkerVersion      []string            `json:"worker_versions"`
}

type KubernetesPackagesResponse struct {
	Packages []KubernetesPackage `json:"packages"`
}

type KubernetesPackage struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// DashboardURLResponse đại diện cho response khi lấy URL của Kubernetes dashboard
type DashboardURLResponse struct {
	URL string `json:"url"`
}

func (c *kubernetesEngineService) resourcePath() string {
	return clusterPath + "/"
}

func (c *kubernetesEngineService) itemPath(id string) string {
	return strings.Join([]string{clusterPath, id}, "/")
}

func (c *kubernetesEngineService) EverywherePath(id string) string {
	return strings.Join([]string{nodeEverywhere, id}, "/")
}

type GetKubeConfigOptions struct {
	ExpiteTime string `json:"expire_time,omitempty"`
}

type UpgradePackageRequest struct {
	NewPackage string `json:"new_package"`
}

// GetKubeConfig - Get Kubernetes config from the given cluster
func (c *kubernetesEngineService) GetKubeConfig(ctx context.Context, clusterUID string, opts *GetKubeConfigOptions) (string, error) {
	req, err := c.client.NewRequest(ctx, http.MethodGet, kubernetesServiceName, strings.Join([]string{c.itemPath(clusterUID), kubeConfig}, "/"), nil)
	if err != nil {
		return "", err
	}
	params := req.URL.Query()
	if opts != nil {
		if opts.ExpiteTime != "" {
			params.Add("expire_time", opts.ExpiteTime)
		}
	}
	req.URL.RawQuery = params.Encode()

	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return "", nil
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	return bodyString, nil
}

type GetKubernetesVersionOpts struct {
	All *bool
}

// GetKubernetesVersion - Get Kubernetes version from the Kubernetes Engine API
func (c *kubernetesEngineService) GetKubernetesVersion(ctx context.Context, opts GetKubernetesVersionOpts) (*KubernetesVersionResponse, error) {
	req, err := c.client.NewRequest(ctx, http.MethodGet, kubernetesServiceName, k8sVersion, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	if opts.All != nil {
		all := strconv.FormatBool(*opts.All)
		q.Add("all", all)
	}
	req.URL.RawQuery = q.Encode()
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data *KubernetesVersionResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

// GetKubernetesVersion - Get Kubernetes Engine Package from the Kubernetes Engine API
func (c *kubernetesEngineService) GetPackages(ctx context.Context, provisionType string) (*KubernetesPackagesResponse, error) {
	req, err := c.client.NewRequest(ctx, http.MethodGet, kubernetesServiceName, k8sPackages, nil)
	if err != nil {
		return nil, err
	}
	params := req.URL.Query()
	params.Add("specify", provisionType)

	req.URL.RawQuery = params.Encode()
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data *KubernetesPackagesResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}
