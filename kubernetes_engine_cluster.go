package gobizfly

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// EverywhereNode represents a Kubernetes everywhere node
type EverywhereNode struct {
	ID              string            `json:"id" yaml:"id"`
	Shoot           string            `json:"shoot" yaml:"shoot"`
	PoolID          string 			  `json:"pool_id" yaml:"pool_id"`
	NodeName     	string            `json:"node_name" yaml:"node_name"`
	PublicIP      	string            `json:"public_ip" yaml:"public_ip"`
	UUID            string            `json:"uuid" yaml:"uuid"`
	CreatedAt       string            `json:"created_at" yaml:"created_at"`
	UpdatedAt       string            `json:"updated_at" yaml:"updated_at"`
}

// ClusterCreateRequest represents the request body for creating a Kubernetes cluster
type ClusterCreateRequest struct {
	Name         string       `json:"name" yaml:"name"`
	Version      string       `json:"version" yaml:"version"`
	AutoUpgrade  bool         `json:"auto_upgrade,omitempty" yaml:"auto_upgrade,omitempty"`
	VPCNetworkID string       `json:"private_network_id" yaml:"private_network_id"`
	EnableCloud  bool         `json:"enable_cloud,omitempty" yaml:"enable_cloud,omitempty"`
	Tags         []string     `json:"tags,omitempty" yaml:"tags,omitempty"`
	WorkerPools  []WorkerPool `json:"worker_pools" yaml:"worker_pools"`
}

// ControllerVersion represents the version of the controller
type ControllerVersion struct {
	ID          string `json:"id,omitempty" yaml:"id,omitempty"`
	Name        string `json:"name,omitempty" yaml:"name,omitempty"`
	Description string `json:"description" yaml:"description"`
	K8SVersion  string `json:"kubernetes_version" yaml:"kubernetes_version"`
}

// Clusters represents list of a Kubernetes clusters
type Clusters struct {
	Clusters_ []Cluster `json:"clusters" yaml:"clusters"`
}

// Cluster represents a Kubernetes cluster
type Cluster struct {
	UID              string            `json:"uid" yaml:"uid"`
	Name             string            `json:"name" yaml:"name"`
	Version          ControllerVersion `json:"version" yaml:"version"`
	VPCNetworkID     string            `json:"private_network_id" yaml:"private_network_id"`
	AutoUpgrade      bool              `json:"auto_upgrade" yaml:"auto_upgrade"`
	Tags             []string          `json:"tags" yaml:"tags"`
	ProvisionStatus  string            `json:"provision_status" yaml:"provision_status"`
	ClusterStatus    string            `json:"cluster_status" yaml:"cluster_status"`
	CreatedAt        string            `json:"created_at" yaml:"created_at"`
	CreatedBy        string            `json:"created_by" yaml:"created_by"`
	WorkerPoolsCount int               `json:"worker_pools_count" yaml:"worker_pools_count"`
	ProvisionType	 string			   `json:"provision_type" yaml:"provision_type"`
}

// ExtendedCluster represents a Kubernetes cluster with additional worker pools' information
type ExtendedCluster struct {
	Cluster
	WorkerPools []ExtendedWorkerPool `json:"worker_pools" yaml:"worker_pools"`
}

// ClusterStat represents the statistic information of a Kubernetes cluster
type ClusterStat struct {
	WorkerPoolCount int `json:"worker_pools" yaml:"worker_pools"`
	TotalCPU        int `json:"total_cpu" yaml:"total_cpu"`
	TotalMemory     int `json:"total_memory" yaml:"total_memory"`
}

// FullCluster represents a Kubernetes cluster with additional worker pools' data and its statistic information
type FullCluster struct {
	ExtendedCluster
	Stat ClusterStat `json:"stat" yaml:"stat"`
}

// List returns a list of Kubernetes clusters
func (c *kubernetesEngineService) List(ctx context.Context, opts *ListOptions) ([]*Cluster, error) {
	req, err := c.client.NewRequest(ctx, http.MethodGet, kubernetesServiceName, c.resourcePath(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var data struct {
		Clusters []*Cluster `json:"clusters" yaml:"clusters"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data.Clusters, nil
}

// Create - create a Kubernetes cluster with worker pools' information
func (c *kubernetesEngineService) Create(ctx context.Context, clcr *ClusterCreateRequest) (*ExtendedCluster, error) {
	var data struct {
		Cluster *ExtendedCluster `json:"cluster" yaml:"cluster"`
	}
	req, err := c.client.NewRequest(ctx, http.MethodPost, kubernetesServiceName, c.resourcePath(), &clcr)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data.Cluster, nil
}

// Get - Get a cluster with additional worker pools' and statistic information
func (c *kubernetesEngineService) Get(ctx context.Context, id string) (*FullCluster, error) {
	var cluster *FullCluster
	req, err := c.client.NewRequest(ctx, http.MethodGet, kubernetesServiceName, c.itemPath(id), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&cluster); err != nil {
		return nil, err
	}
	return cluster, nil
}

// Delete - Delete a Kubernetes cluster
func (c *kubernetesEngineService) Delete(ctx context.Context, id string) error {
	req, err := c.client.NewRequest(ctx, http.MethodDelete, kubernetesServiceName, c.itemPath(id), nil)
	if err != nil {
		return err
	}
	resp, err := c.client.Do(ctx, req)

	if err != nil {
		fmt.Println("error send req")
		return err
	}
	_, _ = io.Copy(ioutil.Discard, resp.Body)
	return resp.Body.Close()
}

func (c *kubernetesEngineService) GetEverywhere(ctx context.Context, id string) (*EverywhereNode, error) {
	var everywhereNode *EverywhereNode
	req, err := c.client.NewRequest(ctx, http.MethodGet, kubernetesServiceName, c.EverywherePath(id), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&everywhereNode); err != nil {
		return nil, err
	}
	return everywhereNode, nil
}