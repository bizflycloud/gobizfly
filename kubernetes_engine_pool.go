package gobizfly

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// WorkerPool represents worker pool information
type WorkerPool struct {
	Name              string            `json:"name"                         yaml:"name"`
	ProvisionType     string            `json:"provision_type,omitempty"     yaml:"provision_type,omitempty"`
	Version           string            `json:"version,omitempty"            yaml:"version,omitempty"`
	Flavor            string            `json:"flavor"                       yaml:"flavor"`
	ProfileType       string            `json:"profile_type"                 yaml:"profile_type"`
	VolumeType        string            `json:"volume_type"                  yaml:"volume_type"`
	VolumeSize        int               `json:"volume_size"                  yaml:"volume_size"`
	AvailabilityZone  string            `json:"availability_zone"            yaml:"availability_zone"`
	DesiredSize       int               `json:"desired_size"                 yaml:"desired_size"`
	EnableAutoScaling bool              `json:"enable_autoscaling,omitempty" yaml:"enable_autoscaling,omitempty"`
	MinSize           int               `json:"min_size,omitempty"           yaml:"min_size,omitempty"`
	MaxSize           int               `json:"max_size,omitempty"           yaml:"max_size,omitempty"`
	Tags              []string          `json:"tags,omitempty"               yaml:"tags,omitempty"`
	Taints            []Taint           `json:"taints,omitempty"             yaml:"taints,omitempty"`
	Labels            map[string]string `json:"labels,omitempty"             yaml:"labels,omitempty"`
	NetworkPlan       string            `json:"network_plan,omitempty"`
	BillingPlan       string            `json:"billing_plan,omitempty"`
}

type Taint struct {
	Effect string `json:"effect"          yaml:"effect"`
	Key    string `json:"key"             yaml:"labels,omitempty"`
	Value  string `json:"value,omitempty" yaml:"value,omitempty"`
}

// ExtendedWorkerPool represents worker pool information with addition fields
type ExtendedWorkerPool struct {
	WorkerPool
	UID                string `json:"id"                   yaml:"id"`
	ProvisionStatus    string `json:"provision_status"     yaml:"provision_status"`
	LaunchConfigID     string `json:"launch_config_id"     yaml:"launch_config_id"`
	AutoScalingGroupID string `json:"autoscaling_group_id" yaml:"autoscaling_group_id"`
	CreatedAt          string `json:"created_at"           yaml:"created_at"`
	ShootID            string `json:"shoot_id" yaml:"shoot_id"`
}

// ExtendedWorkerPools is a list of ExtendedWorkerPool
type ExtendedWorkerPools struct {
	WorkerPools []ExtendedWorkerPool `json:"worker_pools" yaml:"worker_pools"`
}

// AddWorkerPoolsRequest represents the request body to add worker pools
type AddWorkerPoolsRequest struct {
	WorkerPools []WorkerPool `json:"worker_pools" yaml:"worker_pools"`
}

// PoolNode represents node information in a worker pool
type PoolNode struct {
	ID           string   `json:"id"            yaml:"id"`
	Name         string   `json:"name"          yaml:"name"`
	PhysicalID   string   `json:"physical_id"   yaml:"physical_id"`
	IPAddresses  []string `json:"ip_addresses"  yaml:"ip_addresses"`
	Status       string   `json:"status"        yaml:"status"`
	StatusReason string   `json:"status_reason" yaml:"status_reason"`
}

// WorkerPoolWithNodes represents the worker pool information with additional node information
type WorkerPoolWithNodes struct {
	ExtendedWorkerPool
	Nodes []PoolNode `json:"nodes" yaml:"nodes"`
}

// UpdateWorkerPoolRequest represents the request body to update worker pool
type UpdateWorkerPoolRequest struct {
	DesiredSize       int               `json:"desired_size,omitempty"       yaml:"desired_size,omitempty"`
	EnableAutoScaling bool              `json:"enable_autoscaling,omitempty" yaml:"enable_autoscaling,omitempty"`
	MinSize           int               `json:"min_size,omitempty"           yaml:"min_size,omitempty"`
	MaxSize           int               `json:"max_size,omitempty"           yaml:"max_size,omitempty"`
	UpdateStrategy    string            `json:"update_strategy,omitempty"    yaml:"update_strategy,omitempty"`
	Taints            []Taint           `json:"taints,omitempty"             yaml:"taints,omitempty"`
	Labels            map[string]string `json:"labels,omitempty"             yaml:"labels,omitempty"`
}

// AddWorkerPools represents the request body to add worker pools into cluster
func (c *kubernetesEngineService) AddWorkerPools(
	ctx context.Context,
	id string,
	awp *AddWorkerPoolsRequest,
) ([]*ExtendedWorkerPool, error) {
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
		Pools []*ExtendedWorkerPool `json:"worker_pools" yaml:"worker_pools"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, err
	}
	return respData.Pools, nil
}

// RecycleNode - Delete a node and replace by a new one with the same configuration
func (c *kubernetesEngineService) RecycleNode(
	ctx context.Context,
	clusterUID string,
	poolID string,
	nodePhysicalID string,
) error {
	req, err := c.client.NewRequest(
		ctx,
		http.MethodPut,
		kubernetesServiceName,
		strings.Join([]string{clusterPath, clusterUID, poolID, nodePhysicalID}, "/"),
		nil,
	)
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

// DeleteClusterWorkerPool - Delete a worker pool in the given cluster
func (c *kubernetesEngineService) DeleteClusterWorkerPool(ctx context.Context, clusterUID string, PoolID string) error {
	req, err := c.client.NewRequest(
		ctx,
		http.MethodDelete,
		kubernetesServiceName,
		strings.Join([]string{clusterPath, clusterUID, PoolID}, "/"),
		nil,
	)
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

// GetClusterWorkerPool - Get a cluster worker pool with additional node information
func (c *kubernetesEngineService) GetClusterWorkerPool(
	ctx context.Context,
	clusterUID string,
	PoolID string,
) (*WorkerPoolWithNodes, error) {
	var pool *WorkerPoolWithNodes
	req, err := c.client.NewRequest(
		ctx,
		http.MethodGet,
		kubernetesServiceName,
		strings.Join([]string{clusterPath, clusterUID, PoolID}, "/"),
		nil,
	)
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

// UpdateClusterWorkerPool - Update a worker pool in the given cluster
func (c *kubernetesEngineService) UpdateClusterWorkerPool(
	ctx context.Context,
	clusterUID string,
	PoolID string,
	uwp *UpdateWorkerPoolRequest,
) error {
	req, err := c.client.NewRequest(
		ctx,
		http.MethodPatch,
		kubernetesServiceName,
		strings.Join([]string{clusterPath, clusterUID, PoolID}, "/"),
		&uwp,
	)
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

// DeleteClusterWorkerPoolNode - Delete a node in the given worker pool
func (c *kubernetesEngineService) DeleteClusterWorkerPoolNode(
	ctx context.Context,
	clusterUID string,
	PoolID string,
	NodeID string,
) error {
	req, err := c.client.NewRequest(
		ctx,
		http.MethodDelete,
		kubernetesServiceName,
		strings.Join([]string{clusterPath, clusterUID, PoolID, NodeID}, "/")+"?force=true",
		nil,
	)
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

// GetDetailWorkerPool - Get a cluster worker pool
func (c *kubernetesEngineService) GetDetailWorkerPool(
	ctx context.Context,
	PoolID string,
) (*WorkerPoolWithNodes, error) {
	var pool *WorkerPoolWithNodes
	req, err := c.client.NewRequest(
		ctx,
		http.MethodGet,
		kubernetesServiceName,
		strings.Join([]string{workerPoolPath, PoolID}, "/"),
		nil,
	)
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
