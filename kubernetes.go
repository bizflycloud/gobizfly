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
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	clusterPath = "/_"
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
}

type WorkerPool struct {
	Name              string   `json:"name"`
	Version           string   `json:"version,omitempty"`
	Flavor            string   `json:"flavor"`
	ProfileType       string   `json:"profile_type"`
	VolumeType        string   `json:"volume_type"`
	VolumeSize        int      `json:"volume_size"`
	AvailabilityZone  string   `json:"availability_zone"`
	DesiredSize       int      `json:"desired_size"`
	EnableAutoScaling bool     `json:"enable_autoscaling,omitempty"`
	MinSize           int      `json:"min_size"`
	MaxSize           int      `json:"max_size"`
	Tags              []string `json:"tags,omitempty"`
}

type ControllerVersion struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description"`
	K8SVersion  string `json:"kubernetes_version"`
}

type Clusters struct {
	Clusters_ []Cluster `json:"clusters"`
}

type Cluster struct {
	UID              string            `json:"uid"`
	Name             string            `json:"name"`
	Version          ControllerVersion `json:"version"`
	PrivateNetworkID string            `json:"private_network_id"`
	AutoUpgrade      bool              `json:"auto_upgrade"`
	Tags             []string          `json:"tags"`
	ProvisionStatus  string            `json:"provision_status"`
	ClusterStatus    string            `json:"cluster_status"`
	CreatedAt        string            `json:"created_at"`
	CreatedBy        string            `json:"created_by"`
	WorkerPoolsCount int               `json:"worker_pools_count"`
}

type ExtendedCluster struct {
	Cluster
	WorkerPools []ExtendedWorkerPool `json:"worker_pools"`
}

type ClusterStat struct {
	WorkerPoolCount int `json:"worker_pools"`
	TotalCPU        int `json:"total_cpu"`
	TotalMemory     int `json:"total_memory"`
}

type FullCluster struct {
	ExtendedCluster
	Stat ClusterStat `json:"stat"`
}

type ExtendedWorkerPool struct {
	WorkerPool
	UID                string `json:"id"`
	ProvisionStatus    string `json:"provision_status"`
	LaunchConfigID     string `json:"launch_config_id"`
	AutoScalingGroupID string `json:"autoscaling_group_id"`
	CreatedAt          string `json:"created_at"`
}

type ExtendedWorkerPools struct {
	WorkerPools []ExtendedWorkerPool `json:"worker_pools"`
}

type ClusterCreateRequest struct {
	Name             string       `json:"name"`
	Version          string       `json:"version"`
	AutoUpgrade      bool         `json:"auto_upgrade,omitempty"`
	PrivateNetworkID string       `json:"private_network_id"`
	EnableCloud      bool         `json:"enable_cloud,omitempty"`
	Tags             []string     `json:"tags,omitempty"`
	WorkerPools      []WorkerPool `json:"worker_pools"`
}

type AddWorkerPoolsRequest struct {
	WorkerPools []WorkerPool `json:"worker_pools"`
}

type PoolNode struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	PhysicalID   string   `json:"physical_id"`
	IPAddresses  []string `json:"ip_addresses"`
	Status       string   `json:"status"`
	StatusReason string   `json:"status_reason"`
}

type WorkerPoolWithNodes struct {
	ExtendedWorkerPool
	Nodes []PoolNode `json:"nodes"`
}

type UpdateWorkerPoolRequest struct {
	DesiredSize       int  `json:"desired_size"`
	EnableAutoScaling bool `json:"enable_autoscaling"`
	MinSize           int  `json:"min_size"`
	MaxSize           int  `json:"max_size"`
}

func (c *kubernetesEngineService) resourcePath() string {
	return clusterPath + "/"
}

func (c *kubernetesEngineService) itemPath(id string) string {
	return strings.Join([]string{clusterPath, id}, "/")
}

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
		Clusters []*Cluster `json:"clusters"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data.Clusters, nil
}

func (c *kubernetesEngineService) Create(ctx context.Context, clcr *ClusterCreateRequest) (*ExtendedCluster, error) {
	var data struct {
		Cluster *ExtendedCluster `json:"cluster"`
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

func (c *kubernetesEngineService) AddWorkerPools(ctx context.Context, id string, awp *AddWorkerPoolsRequest) ([]*ExtendedWorkerPool, error) {
	req, err := c.client.NewRequest(ctx, http.MethodPost, kubernetesServiceName, c.itemPath(id), &awp)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var respData struct {
		Pools []*ExtendedWorkerPool `json:"worker_pools"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, err
	}
	return respData.Pools, nil
}

func (c *kubernetesEngineService) RecycleNode(ctx context.Context, clusterUID string, poolID string, nodePhysicalID string) error {
	req, err := c.client.NewRequest(ctx, http.MethodPut, kubernetesServiceName, strings.Join([]string{clusterPath, clusterUID, poolID, nodePhysicalID}, "/"), nil)
	if err != nil {
		return err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return err
	}
	_, _ = io.Copy(ioutil.Discard, resp.Body)
	return resp.Body.Close()
}

func (c *kubernetesEngineService) DeleteClusterWorkerPool(ctx context.Context, clusterUID string, PoolID string) error {
	req, err := c.client.NewRequest(ctx, http.MethodDelete, kubernetesServiceName, strings.Join([]string{clusterPath, clusterUID, PoolID}, "/"), nil)
	if err != nil {
		return err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return err
	}
	_, _ = io.Copy(ioutil.Discard, resp.Body)
	return resp.Body.Close()
}

func (c *kubernetesEngineService) GetClusterWorkerPool(ctx context.Context, clusterUID string, PoolID string) (*WorkerPoolWithNodes, error) {
	var pool *WorkerPoolWithNodes
	req, err := c.client.NewRequest(ctx, http.MethodGet, kubernetesServiceName, strings.Join([]string{clusterPath, clusterUID, PoolID}, "/"), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&pool); err != nil {
		return nil, err
	}
	return pool, nil
}

func (c *kubernetesEngineService) UpdateClusterWorkerPool(ctx context.Context, clusterUID string, PoolID string, uwp *UpdateWorkerPoolRequest) error {
	req, err := c.client.NewRequest(ctx, http.MethodPatch, kubernetesServiceName, strings.Join([]string{clusterPath, clusterUID, PoolID}, "/"), &uwp)
	if err != nil {
		return err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return err
	}
	_, _ = io.Copy(ioutil.Discard, resp.Body)
	return resp.Body.Close()
}

func (c *kubernetesEngineService) DeleteClusterWorkerPoolNode(ctx context.Context, clusterUID string, PoolID string, NodeID string) error {
	req, err := c.client.NewRequest(ctx, http.MethodDelete, kubernetesServiceName, strings.Join([]string{clusterPath, clusterUID, PoolID, NodeID}, "/"), nil)
	if err != nil {
		return err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return err
	}
	_, _ = io.Copy(ioutil.Discard, resp.Body)
	return resp.Body.Close()
}
