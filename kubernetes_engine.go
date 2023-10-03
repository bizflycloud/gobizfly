// This file is part of gobizfly

package gobizfly

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	clusterPath           = "/_"
	kubeConfig            = "kubeconfig"
	k8sVersion            = "/k8s_versions"
	workerConfig          = "/worker_config"
	clusterJoinEverywhere = "/engine/cluster_join_everywhere"
	nodeEverywhere        = "_/node_everywhere"
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
	GetKubeConfig(ctx context.Context, clusterUID string) (string, error)
	GetKubernetesVersion(ctx context.Context) (*KubernetesVersionResponse, error)
	GetWorkerConfig(ctx context.Context, wclo *WorkerConfigsListOptions) (*WorkerConfigs, error)
	AddClusterEverywhere(ctx context.Context, id string, cjer *clusterJoinEverywhereRequest) (*clusterJoinEverywhereResponse, error)
	GetEverywhere(ctx context.Context, id string) (*EverywhereNode, error)
}

// KubernetesVersionResponse represents the get versions from the Kubernetes Engine API
type KubernetesVersionResponse struct {
	ControllerVersions []ControllerVersion `json:"controller_versions"`
	WorkerVersion      []string            `json:"worker_versions"`
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

// GetKubeConfig - Get Kubernetes config from the given cluster
func (c *kubernetesEngineService) GetKubeConfig(ctx context.Context, clusterUID string) (string, error) {
	req, err := c.client.NewRequest(ctx, http.MethodGet, kubernetesServiceName, strings.Join([]string{c.itemPath(clusterUID), kubeConfig}, "/"), nil)
	if err != nil {
		return "", err
	}
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

// GetKubernetesVersion - Get Kubernetes version from the Kubernetes Engine API
func (c *kubernetesEngineService) GetKubernetesVersion(ctx context.Context) (*KubernetesVersionResponse, error) {
	req, err := c.client.NewRequest(ctx, http.MethodGet, kubernetesServiceName, k8sVersion, nil)
	if err != nil {
		return nil, err
	}
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

type WorkerConfig struct {
	ID                string `json:"id" yaml:"id"`
	Version           string `json:"version" yaml:"version"`
	Everywhere        bool   `json:"everywhere" yaml:"everywhere"`
	Nvidiadevice      bool   `json:"nvidiadevice" yaml:"nvidiadevice"`
	CniVersion        string `json:"CNI_VERSION" yaml:"CNI_VERSION"`
	RuncVersion       string `json:"RUNC_VERSION" yaml:"RUNC_VERSION"`
	ContainerdVersion string `json:"CONTAINERD_VERSION" yaml:"CONTAINERD_VERSION"`
	KubeVersion       string `json:"KUBE_VERSION" yaml:"KUBE_VERSION"`
}

type WorkerConfigs struct {
	WorkerConfigs []WorkerConfig `json:"worker_configs" yaml:"worker_configs"`
	Page          int            `json:"page" yaml:"page"`
	Limit         int            `json:"limit" yaml:"limit"`
	Total         int            `json:"total" yaml:"total"`
}

type WorkerConfigsListOptions struct {
	Page       string `url:"page,omitempty"`
	Limit      string `url:"limit,omitempty"`
	Total      string `url:"total,omitempty"`
	Everywhere string `url:"everywhere,omitempty"`
	Version    string `url:"version,omitempty"`
}

func (c *kubernetesEngineService) GetWorkerConfig(ctx context.Context, wclo *WorkerConfigsListOptions) (*WorkerConfigs, error) {
	var workerConfigs *WorkerConfigs
	req, err := c.client.NewRequest(ctx, http.MethodGet, kubernetesServiceName, workerConfig, nil)
	if err != nil {
		return nil, err
	}

	params := req.URL.Query()
	if wclo != nil {
		if wclo.Page != "" {
			params.Add("page", wclo.Page)
		}
		if wclo.Limit != "" {
			params.Add("limit", wclo.Limit)
		}
		if wclo.Total != "" {
			params.Add("total", wclo.Total)
		}
		if wclo.Everywhere != "" {
			params.Add("everywhere", wclo.Everywhere)
		}
		if wclo.Version != "" {
			params.Add("version", wclo.Version)
		}
	}
	req.URL.RawQuery = params.Encode()

	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(workerConfigs); err != nil {
		return nil, err
	}
	return workerConfigs, nil
}
