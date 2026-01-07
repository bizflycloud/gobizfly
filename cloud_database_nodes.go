// This file is part of gobizfly

package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
)

const (
	cloudDatabaseNodesResourcePath = "/nodes"
)

type cloudDatabaseNodes struct {
	client *Client
}

// CloudDatabaseNodeCreate contains detail to create database node payload.
type CloudDatabaseNodeCreate struct {
	Configuration string                          `json:"configuration,omitempty"`
	Name          string                          `json:"name,omitempty"`
	ReplicaOf     string                          `json:"replica_of,omitempty" validate:"required"`
	Replicas      *CloudDatabaseReplicaNodeCreate `json:"replicas,omitempty"`
	Role          string                          `json:"role,omitempty"`
	Secondaries   *CloudDatabaseReplicaNodeCreate `json:"secondaries,omitempty"`
	Suggestion    bool                            `json:"suggestion,omitempty"`
}

// CloudDatabaseNodeCreateResponse contains response from Create Node API.
// Flavor can be string or object depending on API version.
type CloudDatabaseNodeCreateResponse struct {
	ID               string                 `json:"id"`
	Name             string                 `json:"name"`
	Status           string                 `json:"status"`
	OperatingStatus  string                 `json:"operating_status"`
	Flavor           interface{}            `json:"flavor"` 
	Datastore        CloudDatabaseDatastore `json:"datastore"`
	Region           string                 `json:"region"`
	Volume           CloudDatabaseVolume    `json:"volume"`
	ReplicaOf        interface{}            `json:"replica_of"` 
	CreatedAt        string                 `json:"created"`
	UpdatedAt        string                 `json:"updated"`
	TaskID           string                 `json:"task_id"`
	Message          string                 `json:"message"`
	AvailabilityZone string                 `json:"availability_zone"`
	Description      string                 `json:"description"`
	DNS              CloudDatabaseDNS       `json:"dns"`
	EnableFailover   bool                   `json:"enable_failover"`
	InstanceID       string                 `json:"instance_id"`
	NodeType         string                 `json:"node_type"`
	RegionName       string                 `json:"region_name"`
	Role             string                 `json:"role"`
}

// CloudDatabaseNode contains detail of a database node.
type CloudDatabaseNode struct {
	Addresses        CloudDatabaseAddresses `json:"addresses"`
	AvailabilityZone string                 `json:"availability_zone"`
	CreatedAt        string                 `json:"created_at"`
	Datastore        CloudDatabaseDatastore `json:"datastore"`
	Description      string                 `json:"description"`
	DNS              CloudDatabaseDNS       `json:"dns"`
	EnableFailover   bool                   `json:"enable_failover"`
	Flavor           string                 `json:"flavor"`
	ID               string                 `json:"id"`
	InstanceID       string                 `json:"instance_id"`
	Message          string                 `json:"message"`
	Name             string                 `json:"name"`
	NodeType         string                 `json:"node_type"`
	OperatingStatus  string                 `json:"operating_status"`
	RegionName       string                 `json:"region_name"`
	ReplicaOf        string                 `json:"replica_of"`
	Replicas         []CloudDatabaseNode    `json:"replicas"`
	Role             string                 `json:"role"`
	Status           string                 `json:"status"`
	TaskID           string                 `json:"task_id"`
	Volume           CloudDatabaseVolume    `json:"volume"`
}

func (db *cloudDatabaseService) Nodes() *cloudDatabaseNodes {
	return &cloudDatabaseNodes{client: db.client}
}

// CloudDatabase Node Resource Path
func (no *cloudDatabaseNodes) resourcePath(nodeID string) string {
	return cloudDatabaseNodesResourcePath + "/" + nodeID
}

// CloudDatabase Node Resource Action Path
func (no *cloudDatabaseNodes) resourceActionPath(nodeID string) string {
	return cloudDatabaseNodesResourcePath + "/" + nodeID + "/action"
}

// CloudDatabase Node resourceType Path
func (no *cloudDatabaseNodes) resourceTypePath(nodeID string, resourceType string) string {
	return cloudDatabaseNodesResourcePath + "/" + nodeID + "/" + resourceType
}

// List all Node.
func (no *cloudDatabaseNodes) List(ctx context.Context, opts *CloudDatabaseListOption) ([]*CloudDatabaseNode, error) {
	req, err := no.client.NewRequest(ctx, http.MethodGet, databaseServiceName, cloudDatabaseNodesResourcePath, nil)
	if err != nil {
		return nil, err
	}

	AddParamsListOption(req, opts)

	resp, err := no.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()
	var data struct {
		Nodes []*CloudDatabaseNode `json:"nodes"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Nodes, nil
}

// ListBackups List all backup in nodes.
func (no *cloudDatabaseNodes) ListBackups(ctx context.Context, nodeID string, opts *CloudDatabaseListOption) ([]*CloudDatabaseBackup, error) {
	req, err := no.client.NewRequest(ctx, http.MethodGet, databaseServiceName, no.resourceTypePath(nodeID, backups), nil)
	if err != nil {
		return nil, err
	}

	AddParamsListOption(req, opts)

	resp, err := no.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()
	var data struct {
		Backups []*CloudDatabaseBackup `json:"backups"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Backups, nil
}

// ListBackupSchedules List all schedule in nodes.
func (no *cloudDatabaseNodes) ListBackupSchedules(ctx context.Context, nodeID string, opts *CloudDatabaseListOption) ([]*CloudDatabaseBackupSchedule, error) {
	req, err := no.client.NewRequest(ctx, http.MethodGet, databaseServiceName, no.resourceTypePath(nodeID, schedules), nil)
	if err != nil {
		return nil, err
	}

	AddParamsListOption(req, opts)

	resp, err := no.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()
	var data struct {
		BackupSchedules []*CloudDatabaseBackupSchedule `json:"schedules"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.BackupSchedules, nil
}

// Create a new replica or secondary nodes
func (no *cloudDatabaseNodes) Create(ctx context.Context, icr *CloudDatabaseNodeCreate) (*CloudDatabaseNodeCreateResponse, error) {
	req, err := no.client.NewRequest(ctx, http.MethodPost, databaseServiceName, cloudDatabaseNodesResourcePath, &icr)
	if err != nil {
		return nil, err
	}

	resp, err := no.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()
	var nodeResp *CloudDatabaseNodeCreateResponse

	if err := json.NewDecoder(resp.Body).Decode(&nodeResp); err != nil {
		return nil, err
	}

	return nodeResp, nil
}

// CreateSuggestion get suggestion when create a new replica or secondary nodes.
func (no *cloudDatabaseNodes) CreateSuggestion(ctx context.Context, icr *CloudDatabaseNodeCreate) (*CloudDatabaseSuggestion, error) {
	// set true for get suggestion
	icr.Suggestion = true

	req, err := no.client.NewRequest(ctx, http.MethodPost, databaseServiceName, cloudDatabaseNodesResourcePath, &icr)
	if err != nil {
		return nil, err
	}

	resp, err := no.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()
	var suggestions *CloudDatabaseSuggestion

	if err := json.NewDecoder(resp.Body).Decode(&suggestions); err != nil {
		return nil, err
	}

	return suggestions, nil
}

// Get a node.
func (no *cloudDatabaseNodes) Get(ctx context.Context, nodeID string) (*CloudDatabaseNode, error) {
	req, err := no.client.NewRequest(ctx, http.MethodGet, databaseServiceName, no.resourcePath(nodeID), nil)
	if err != nil {
		return nil, err
	}

	resp, err := no.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()
	var nodes *CloudDatabaseNode

	if err := json.NewDecoder(resp.Body).Decode(&nodes); err != nil {
		return nil, err
	}

	return nodes, nil
}

// Action with a nodes.
func (no *cloudDatabaseNodes) Action(ctx context.Context, nodeID string, nar *CloudDatabaseAction) (*CloudDatabaseMessageResponse, error) {
	if nar.Action == resizeFlavor && nar.FlavorName == "" {
		return nil, ErrRequireFlavorName
	}

	if nar.Action == resizeVolume && nar.NewSize == 0 {
		return nil, ErrRequireNewSize
	}

	req, err := no.client.NewRequest(ctx, http.MethodPost, databaseServiceName, no.resourceActionPath(nodeID), &nar)
	if err != nil {
		return nil, err
	}

	resp, err := no.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()
	var dmr *CloudDatabaseMessageResponse

	if err := json.NewDecoder(resp.Body).Decode(&dmr); err != nil {
		return nil, err
	}

	return dmr, nil
}

// ActionSuggestion get suggestion when action with a nodes.
func (no *cloudDatabaseNodes) ActionSuggestion(ctx context.Context, nodeID string, nar *CloudDatabaseAction) (*CloudDatabaseSuggestion, error) {
	if nar.Action == resizeFlavor && nar.FlavorName == "" {
		return nil, ErrRequireFlavorName
	}

	if nar.Action == resizeVolume && nar.NewSize == 0 {
		return nil, ErrRequireNewSize
	}

	// set true for get suggestion
	nar.Suggestion = true

	req, err := no.client.NewRequest(ctx, http.MethodPost, databaseServiceName, no.resourceActionPath(nodeID), &nar)
	if err != nil {
		return nil, err
	}

	resp, err := no.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()
	var suggestions *CloudDatabaseSuggestion

	if err := json.NewDecoder(resp.Body).Decode(&suggestions); err != nil {
		return nil, err
	}

	return suggestions, nil
}

// ResizeFlavor with an instances.
func (no *cloudDatabaseNodes) ResizeFlavor(ctx context.Context, nodeID string, flavorName string) (*CloudDatabaseMessageResponse, error) {
	var iar = CloudDatabaseAction{Action: resizeFlavor, FlavorName: flavorName}

	req, err := no.client.NewRequest(ctx, http.MethodPost, databaseServiceName, no.resourceActionPath(nodeID), &iar)
	if err != nil {
		return nil, err
	}

	resp, err := no.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()
	var dmr *CloudDatabaseMessageResponse

	if err := json.NewDecoder(resp.Body).Decode(&dmr); err != nil {
		return nil, err
	}

	return dmr, nil
}

// ResizeFlavorSuggestion get suggestion when resize an instances.
func (no *cloudDatabaseNodes) ResizeFlavorSuggestion(ctx context.Context, nodeID string, flavorName string) (*CloudDatabaseSuggestion, error) {
	var iar = CloudDatabaseAction{Action: resizeFlavor, FlavorName: flavorName, Suggestion: true}

	req, err := no.client.NewRequest(ctx, http.MethodPost, databaseServiceName, no.resourceActionPath(nodeID), &iar)
	if err != nil {
		return nil, err
	}

	resp, err := no.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()
	var suggestions *CloudDatabaseSuggestion

	if err := json.NewDecoder(resp.Body).Decode(&suggestions); err != nil {
		return nil, err
	}

	return suggestions, nil
}

// ResizeVolume with an instances.
func (no *cloudDatabaseNodes) ResizeVolume(ctx context.Context, nodeID string, newSize int) (*CloudDatabaseMessageResponse, error) {
	var iar = CloudDatabaseAction{Action: resizeVolume, NewSize: newSize}

	req, err := no.client.NewRequest(ctx, http.MethodPost, databaseServiceName, no.resourceActionPath(nodeID), &iar)
	if err != nil {
		return nil, err
	}

	resp, err := no.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()
	var dmr *CloudDatabaseMessageResponse

	if err := json.NewDecoder(resp.Body).Decode(&dmr); err != nil {
		return nil, err
	}

	return dmr, nil
}

// ResizeVolumeSuggestion get suggestion when resize an instances.
func (no *cloudDatabaseNodes) ResizeVolumeSuggestion(ctx context.Context, nodeID string, newSize int) (*CloudDatabaseSuggestion, error) {
	var iar = CloudDatabaseAction{Action: resizeVolume, NewSize: newSize, Suggestion: true}

	req, err := no.client.NewRequest(ctx, http.MethodPost, databaseServiceName, no.resourceActionPath(nodeID), &iar)
	if err != nil {
		return nil, err
	}

	resp, err := no.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()
	var suggestions *CloudDatabaseSuggestion

	if err := json.NewDecoder(resp.Body).Decode(&suggestions); err != nil {
		return nil, err
	}

	return suggestions, nil
}

// Restart with an instances.
func (no *cloudDatabaseNodes) Restart(ctx context.Context, nodeID string) (*CloudDatabaseMessageResponse, error) {
	var iar = CloudDatabaseConfigAction{Action: restart}

	req, err := no.client.NewRequest(ctx, http.MethodPost, databaseServiceName, no.resourceActionPath(nodeID), &iar)
	if err != nil {
		return nil, err
	}

	resp, err := no.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()
	var dmr *CloudDatabaseMessageResponse

	if err := json.NewDecoder(resp.Body).Decode(&dmr); err != nil {
		return nil, err
	}

	return dmr, nil
}

// DetachReplica with an instances.
func (no *cloudDatabaseNodes) DetachReplica(ctx context.Context, nodeID string) (*CloudDatabaseMessageResponse, error) {
	var iar = CloudDatabaseAction{Action: detachReplica}

	req, err := no.client.NewRequest(ctx, http.MethodPost, databaseServiceName, no.resourceActionPath(nodeID), &iar)
	if err != nil {
		return nil, err
	}

	resp, err := no.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()
	var dmr *CloudDatabaseMessageResponse

	if err := json.NewDecoder(resp.Body).Decode(&dmr); err != nil {
		return nil, err
	}

	return dmr, nil
}

// EnableRoot with an instances.
func (no *cloudDatabaseNodes) EnableRoot(ctx context.Context, nodeID string) (*CloudDatabaseMessageResponse, error) {
	var iar = CloudDatabaseAction{Action: enableRoot}

	req, err := no.client.NewRequest(ctx, http.MethodPost, databaseServiceName, no.resourceActionPath(nodeID), &iar)
	if err != nil {
		return nil, err
	}

	resp, err := no.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()
	var dmr *CloudDatabaseMessageResponse

	if err := json.NewDecoder(resp.Body).Decode(&dmr); err != nil {
		return nil, err
	}

	return dmr, nil
}

// Delete a node.
func (no *cloudDatabaseNodes) Delete(ctx context.Context, nodeID string, idr *CloudDatabaseDelete) (*CloudDatabaseMessageResponse, error) {
	req, err := no.client.NewRequest(ctx, http.MethodDelete, databaseServiceName, no.resourcePath(nodeID), &idr)
	if err != nil {
		return nil, err
	}

	resp, err := no.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()
	var dmr *CloudDatabaseMessageResponse

	if err := json.NewDecoder(resp.Body).Decode(&dmr); err != nil {
		return nil, err
	}

	return dmr, nil
}
