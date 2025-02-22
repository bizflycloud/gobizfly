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

// EverywhereNode represents a Kubernetes everywhere node
type EverywhereNode struct {
	ID        string `json:"id"         yaml:"id"`
	Shoot     string `json:"shoot"      yaml:"shoot"`
	PoolID    string `json:"pool_id"    yaml:"pool_id"`
	NodeName  string `json:"node_name"  yaml:"node_name"`
	PublicIP  string `json:"public_ip"  yaml:"public_ip"`
	PrivateIP string `json:"private_ip" yaml:"private_ip"`
	Region    string `json:"region"     yaml:"region"`
	UUID      string `json:"uuid"       yaml:"uuid"`
	CreatedAt string `json:"created_at" yaml:"created_at"`
	UpdatedAt string `json:"updated_at" yaml:"updated_at"`
	Deleted   bool   `json:"deleted"    yaml:"deleted"`
}

// ClusterCreateRequest represents the request body for creating a Kubernetes cluster
type ClusterCreateRequest struct {
	Name          string       `json:"name"                       yaml:"name"`
	Version       string       `json:"version"                    yaml:"version"`
	ProvisionType string       `json:"provision_type,omitempty"   yaml:"provision_type,omitempty"`
	Package       string       `json:"package,omitempty"          yaml:"package,omitempty"`
	AutoUpgrade   bool         `json:"auto_upgrade,omitempty"     yaml:"auto_upgrade,omitempty"`
	VPCNetworkID  string       `json:"private_network_id"         yaml:"private_network_id"`
	EnableCloud   bool         `json:"enable_cloud,omitempty"     yaml:"enable_cloud,omitempty"`
	Tags          []string     `json:"tags,omitempty"             yaml:"tags,omitempty"`
	LocalDNS      bool         `json:"local_dns"`
	CNIPlugin     string       `json:"cni_plugin,omitempty"`
	WorkerPools   []WorkerPool `json:"worker_pools"           yaml:"worker_pools"`
}

// ControllerVersion represents the version of the controller
type ControllerVersion struct {
	ID          string `json:"id,omitempty"       yaml:"id,omitempty"`
	Name        string `json:"name,omitempty"     yaml:"name,omitempty"`
	Description string `json:"description"        yaml:"description"`
	K8SVersion  string `json:"kubernetes_version" yaml:"kubernetes_version"`
}

// Clusters represents list of a Kubernetes clusters
type Clusters struct {
	Clusters_ []Cluster `json:"clusters" yaml:"clusters"`
}

// Cluster represents a Kubernetes cluster
type Cluster struct {
	UID              string            `json:"uid"                yaml:"uid"`
	Name             string            `json:"name"               yaml:"name"`
	Version          ControllerVersion `json:"version"            yaml:"version"`
	ClusterPackage   KubernetesPackage `json:"package"            yaml:"package"`
	VPCNetworkID     string            `json:"private_network_id" yaml:"private_network_id"`
	SubnetID         string            `json:"private_subnet_id"`
	AutoUpgrade      bool              `json:"auto_upgrade"       yaml:"auto_upgrade"`
	Tags             []string          `json:"tags"               yaml:"tags"`
	ProvisionStatus  string            `json:"provision_status"   yaml:"provision_status"`
	ClusterStatus    string            `json:"cluster_status"     yaml:"cluster_status"`
	CreatedAt        string            `json:"created_at"         yaml:"created_at"`
	CreatedBy        string            `json:"created_by"         yaml:"created_by"`
	WorkerPoolsCount int               `json:"worker_pools_count" yaml:"worker_pools_count"`
	ProvisionType    string            `json:"provision_type"     yaml:"provision_type"`
	AccessMode       string            `json:"access_mode"`
	CNIPlugin        string            `json:"cni_plugin"`
	ForceUpgrade     bool              `json:"force_upgrade"`
	LocalDNS         bool              `json:"local_dns"`
	BcrIntegrated    bool              `json:"bcr_integrated"`
}

// ExtendedCluster represents a Kubernetes cluster with additional worker pools' information
type ExtendedCluster struct {
	Cluster
	WorkerPools []ExtendedWorkerPool `json:"worker_pools" yaml:"worker_pools"`
}

// ClusterStat represents the statistic information of a Kubernetes cluster
type ClusterStat struct {
	WorkerPoolCount int `json:"worker_pools" yaml:"worker_pools"`
	TotalCPU        int `json:"total_cpu"    yaml:"total_cpu"`
	TotalMemory     int `json:"total_memory" yaml:"total_memory"`
}

// UpgradeVersionTime represents the upgrade version time of a Kubernetes cluster
type UpgradeVersionTime struct {
	Day  int    `json:"day"`
	Time string `json:"time"`
}

// FullCluster represents a Kubernetes cluster with additional worker pools' data and its statistic information
type FullCluster struct {
	ExtendedCluster
	UpgradeTime UpgradeVersionTime `json:"upgrade_time"`
	Stat        ClusterStat        `json:"stat" yaml:"stat"`
}

// UpdateClusterRequest represents the request body for update a Kubernetes cluster
type UpdateClusterRequest struct {
	AccessPolicies *[]string          `json:"access_policies,omitempty"`
	AutoUpgrade    *bool              `json:"auto_upgrade,omitempty"`
	BcrIntegrated  *bool              `json:"bcr_integrated,omitempty"`
	UpgradeTime    UpgradeVersionTime `json:"upgrade_time,omitempty"`
}

type AddonStatusResponse struct {
	Type    string `json:"type"`
	Version string `json:"version"`
	Status  string `json:"status"`
}

// UpgradeClusterVersionRequest represents the request body for upgrade version a Kubernetes cluster
type UpgradeClusterVersionRequest struct {
	ControlPlaneOnly string `json:"control_plane_only,omitempty"`
}

// UpgradeClusterVersionResponse represents the response of the request for upgrade version a Kubernetes cluster
type UpgradeClusterVersionResponse struct {
	IsLatest  bool   `json:"is_latest"`
	UpgradeTo string `json:"upgrade_to"`
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

func (c *kubernetesEngineService) UpdateCluster(ctx context.Context, id string, payload *UpdateClusterRequest) (*ExtendedCluster, error) {
	var detailCluster ExtendedCluster
	req, err := c.client.NewRequest(ctx, http.MethodPatch, kubernetesServiceName, c.itemPath(id), payload)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&detailCluster); err != nil {
		return nil, err
	}
	return &detailCluster, nil
}

func (c *kubernetesEngineService) GetUpgradeClusterVersion(ctx context.Context, id string) (*UpgradeClusterVersionResponse, error) {
	var upgradeClusterVersion struct {
		Upgrade UpgradeClusterVersionResponse
	}
	path := strings.Join([]string{id, "upgrade"}, "/")
	req, err := c.client.NewRequest(ctx, http.MethodGet, kubernetesServiceName, c.itemPath(path), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&upgradeClusterVersion); err != nil {
		return nil, err
	}
	return &upgradeClusterVersion.Upgrade, nil
}

func (c *kubernetesEngineService) UpgradeClusterVersion(ctx context.Context, id string, payload *UpgradeClusterVersionRequest) error {
	path := strings.Join([]string{id, "upgrade"}, "/")
	req, err := c.client.NewRequest(ctx, http.MethodPost, kubernetesServiceName, c.itemPath(path), payload)
	if err != nil {
		return err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return err
	}
	return resp.Body.Close()
}

func (c *kubernetesEngineService) UpgradePackage(ctx context.Context, id string, payload *UpgradePackageRequest) error {
	path := strings.Join([]string{id, "package", "upgrade"}, "/")
	req, err := c.client.NewRequest(ctx, http.MethodPost, kubernetesServiceName, c.itemPath(path), payload)
	if err != nil {
		return err
	}

	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func (c *kubernetesEngineService) GetDashboardURL(ctx context.Context, id string) (string, error) {
	path := strings.Join([]string{id, "dashboard"}, "/")
	req, err := c.client.NewRequest(ctx, http.MethodGet, kubernetesServiceName, c.itemPath(path), nil)
	if err != nil {
		return "", err
	}

	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var data *DashboardURLResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}
	return data.URL, nil
}

func (c *kubernetesEngineService) InstallAddon(ctx context.Context, id string, addonType string) error {
	path := strings.Join([]string{id, "add_ons", addonType}, "/")
	payload := map[string]string{
		"action": "install",
	}
	req, err := c.client.NewRequest(ctx, http.MethodPost, kubernetesServiceName, c.itemPath(path), payload)
	if err != nil {
		return err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (c *kubernetesEngineService) UninstallAddon(ctx context.Context, id string, addonType string) error {
	path := strings.Join([]string{id, "add_ons", addonType}, "/")
	payload := map[string]string{
		"action": "uninstall",
	}
	req, err := c.client.NewRequest(ctx, http.MethodPost, kubernetesServiceName, c.itemPath(path), payload)
	if err != nil {
		return err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (c *kubernetesEngineService) GetAddonStatus(ctx context.Context, id string, addonType string) (*AddonStatusResponse, error) {
	path := strings.Join([]string{id, "add_ons", addonType}, "/") + "?status_only=true"
	req, err := c.client.NewRequest(ctx, http.MethodGet, kubernetesServiceName, c.itemPath(path), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var addonStatus AddonStatusResponse
	if err := json.NewDecoder(resp.Body).Decode(&addonStatus); err != nil {
		return nil, err
	}
	return &addonStatus, nil
}
