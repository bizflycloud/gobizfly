// This file is part of gobizfly

package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
)

const (
	cloudDatabaseInstancesResourcePath = "/instances"
)

type cloudDatabaseInstances struct {
	client *Client
}

// CloudDatabaseListOption contains option when list resource
type CloudDatabaseListOption struct {
	DatabaseEngine  string `json:"database_engine,omitempty"`
	DatabaseVersion string `json:"database_version,omitempty"`
	Detailed        bool   `json:"detailed,omitempty"`
	EndTime         string `json:"end_time,omitempty"`
	Name            string `json:"name,omitempty"`
	Page            int    `json:"page,omitempty"`
	ResultsPerPage  int    `json:"results_per_page,omitempty"`
	StartTime       string `json:"start_time,omitempty"`
}

// CloudDatabaseNetworks contains network information to create database instance.
type CloudDatabaseNetworks struct {
	NetworkID string `json:"network_id"`
}

// CloudDatabaseLog include about logs of nodes
type CloudDatabaseLog struct {
	Enable string `json:"enable"`
	Name   string `json:"name"`
}

// CloudDatabaseInstance contains database instance information.
type CloudDatabaseInstance struct {
	AutoScaling    CloudDatabaseAutoScaling `json:"autoscaling"`
	EnableFailover bool                     `json:"enable_failover"`
	CreatedAt      string                   `json:"created"`
	Datastore      CloudDatabaseDatastore   `json:"datastore"`
	Description    string                   `json:"description"`
	DNS            CloudDatabaseDNS         `json:"dns"`
	ID             string                   `json:"id"`
	InstanceType   string                   `json:"instance_type"`
	Logs           CloudDatabaseLog         `json:"logs"`
	Message        string                   `json:"message"`
	Name           string                   `json:"name"`
	Nodes          []CloudDatabaseNode      `json:"nodes"`
	ProjectID      string                   `json:"project_id"`
	PublicAccess   bool                     `json:"public_access"`
	Status         string                   `json:"status"`
	TaskID         string                   `json:"task_id"`
	Volume         CloudDatabaseVolume      `json:"volume"`
}

// CloudDatabaseInstanceCreate contains payload to create database instance.
type CloudDatabaseInstanceCreate struct {
	AutoScaling      *CloudDatabaseAutoScaling       `json:"autoscaling,omitempty"`
	AvailabilityZone string                          `json:"availability_zone" validate:"required"`
	BackupID         string                          `json:"backup_id,omitempty"`
	Datastore        CloudDatabaseDatastore          `json:"datastore" validate:"required"`
	EnableFailover   bool                            `json:"enable_failover,omitempty"`
	FlavorName       string                          `json:"flavor_name" validate:"required"`
	InstanceType     string                          `json:"instance_type,omitempty"`
	Name             string                          `json:"name" validate:"required"`
	Networks         []CloudDatabaseNetworks         `json:"networks" validate:"required"`
	PublicAccess     bool                            `json:"public_access,omitempty"`
	Replicas         *CloudDatabaseReplicaNodeCreate `json:"replicas,omitempty"`
	Secondaries      *CloudDatabaseReplicaNodeCreate `json:"secondaries,omitempty"`
	Suggestion       bool                            `json:"suggestion,omitempty"`
	VolumeSize       int                             `json:"volume_size" validate:"required"`
}

// CloudDatabaseDelete contains database delete information.
type CloudDatabaseDelete struct {
	PurgeBackup     bool `json:"purge_backup"`
	PurgeAutobackup bool `json:"purge_autobackup"`
}

// CloudDatabaseAction contains database instance action information.
type CloudDatabaseAction struct {
	Action       string                 `json:"action,omitempty" validate:"required"`
	ActionAll    bool                   `json:"action_all,omitempty"`
	Datastore    CloudDatabaseDatastore `json:"datastore" validate:"required"`
	FlavorName   string                 `json:"flavor_name,omitempty"`
	InstanceType string                 `json:"instance_type" validate:"required"`
	NewSize      int                    `json:"new_size,omitempty"`
	NodeIDs      []string               `json:"node_ids,omitempty"`
	NodeType     string                 `json:"node_type,omitempty"`
	ResizeAll    bool                   `json:"resize_all,omitempty"`
	Suggestion   bool                   `json:"suggestion,omitempty"`
}

// CloudDatabaseConfigAction contains database instance action information.
type CloudDatabaseConfigAction struct {
	Action string `json:"action" validate:"required"`
}

// CloudDatabaseDBReq contains databases information to send API
type CloudDatabaseDBReq struct {
	Databases []*CloudDatabaseDB `json:"databases,omitempty" validate:"required"`
}

// CloudDatabaseUserReq contains users information to send API
type CloudDatabaseUserReq struct {
	Users []*CloudDatabaseUser `json:"users,omitempty" validate:"required"`
}

func (db *cloudDatabaseService) Instances() *cloudDatabaseInstances {
	return &cloudDatabaseInstances{client: db.client}
}

// CloudDatabase Instance Resource Path
func (ins *cloudDatabaseInstances) resourcePath(instanceID string) string {
	return cloudDatabaseInstancesResourcePath + "/" + instanceID
}

// CloudDatabase Instance Action Path
func (ins *cloudDatabaseInstances) resourceActionPath(instanceID string) string {
	return cloudDatabaseInstancesResourcePath + "/" + instanceID + "/action"
}

// CloudDatabase Instance resourceType Path
func (ins *cloudDatabaseInstances) resourceTypePath(instanceID string, resourceType string) string {
	return cloudDatabaseInstancesResourcePath + "/" + instanceID + "/" + resourceType
}

// List all instances.
func (ins *cloudDatabaseInstances) List(ctx context.Context, opts *CloudDatabaseListOption) ([]*CloudDatabaseInstance, error) {
	req, err := ins.client.NewRequest(ctx, http.MethodGet, databaseServiceName, cloudDatabaseInstancesResourcePath, nil)
	if err != nil {
		return nil, err
	}

	AddParamsListOption(req, opts)

	resp, err := ins.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()
	var data struct {
		Instances []*CloudDatabaseInstance `json:"instances"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Instances, nil
}

// ListNodes List all nodes in instances.
func (ins *cloudDatabaseInstances) ListNodes(ctx context.Context, instanceID string, opts *CloudDatabaseListOption) ([]*CloudDatabaseNode, error) {
	req, err := ins.client.NewRequest(ctx, http.MethodGet, databaseServiceName, ins.resourceTypePath(instanceID, nodes), nil)
	if err != nil {
		return nil, err
	}

	AddParamsListOption(req, opts)

	resp, err := ins.client.Do(ctx, req)
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

// ListBackups List all backup in instances.
func (ins *cloudDatabaseInstances) ListBackups(ctx context.Context, instanceID string, opts *CloudDatabaseListOption) ([]*CloudDatabaseBackup, error) {
	req, err := ins.client.NewRequest(ctx, http.MethodGet, databaseServiceName, ins.resourceTypePath(instanceID, backups), nil)
	if err != nil {
		return nil, err
	}

	AddParamsListOption(req, opts)

	resp, err := ins.client.Do(ctx, req)
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

// ListBackupSchedules List all schedule in instances.
func (ins *cloudDatabaseInstances) ListBackupSchedules(ctx context.Context, instanceID string, opts *CloudDatabaseListOption) ([]*CloudDatabaseBackupSchedule, error) {
	req, err := ins.client.NewRequest(ctx, http.MethodGet, databaseServiceName, ins.resourceTypePath(instanceID, schedules), nil)
	if err != nil {
		return nil, err
	}

	AddParamsListOption(req, opts)

	resp, err := ins.client.Do(ctx, req)
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

// Create a new instances
func (ins *cloudDatabaseInstances) Create(ctx context.Context, icr *CloudDatabaseInstanceCreate) (*CloudDatabaseInstance, error) {
	req, err := ins.client.NewRequest(ctx, http.MethodPost, databaseServiceName, cloudDatabaseInstancesResourcePath, &icr)
	if err != nil {
		return nil, err
	}

	resp, err := ins.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()
	var instances *CloudDatabaseInstance
	if err := json.NewDecoder(resp.Body).Decode(&instances); err != nil {
		return nil, err
	}
	return instances, nil
}

// CreateSuggestion get suggestion when create a new instance
func (ins *cloudDatabaseInstances) CreateSuggestion(ctx context.Context, icr *CloudDatabaseInstanceCreate) (*CloudDatabaseSuggestion, error) {
	// set true for get suggestion
	icr.Suggestion = true

	req, err := ins.client.NewRequest(ctx, http.MethodPost, databaseServiceName, cloudDatabaseInstancesResourcePath, &icr)
	if err != nil {
		return nil, err
	}

	resp, err := ins.client.Do(ctx, req)
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

// Get an instances.
func (ins *cloudDatabaseInstances) Get(ctx context.Context, instanceID string) (*CloudDatabaseInstance, error) {
	req, err := ins.client.NewRequest(ctx, http.MethodGet, databaseServiceName, ins.resourcePath(instanceID), nil)
	if err != nil {
		return nil, err
	}

	resp, err := ins.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()
	var instance *CloudDatabaseInstance

	if err := json.NewDecoder(resp.Body).Decode(&instance); err != nil {
		return nil, err
	}

	return instance, nil
}

// Action with an instances.
func (ins *cloudDatabaseInstances) Action(ctx context.Context, instanceID string, iar *CloudDatabaseAction) (*CloudDatabaseMessageResponse, error) {
	if iar.Action == resizeFlavor && iar.FlavorName == "" {
		return nil, ErrRequireFlavorName
	}

	if iar.Action == resizeVolume && iar.NewSize == 0 {
		return nil, ErrRequireNewSize
	}

	req, err := ins.client.NewRequest(ctx, http.MethodPost, databaseServiceName, ins.resourceActionPath(instanceID), &iar)
	if err != nil {
		return nil, err
	}

	resp, err := ins.client.Do(ctx, req)
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

// ActionSuggestion get suggestion when action with instances.
func (ins *cloudDatabaseInstances) ActionSuggestion(ctx context.Context, instanceID string, iar *CloudDatabaseAction) (*CloudDatabaseSuggestion, error) {
	if iar.Action == resizeFlavor && iar.FlavorName == "" {
		return nil, ErrRequireFlavorName
	}

	if iar.Action == resizeVolume && iar.NewSize == 0 {
		return nil, ErrRequireNewSize
	}

	// set true for get suggestion
	iar.Suggestion = true

	req, err := ins.client.NewRequest(ctx, http.MethodPost, databaseServiceName, ins.resourceActionPath(instanceID), &iar)
	if err != nil {
		return nil, err
	}

	resp, err := ins.client.Do(ctx, req)
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
func (ins *cloudDatabaseInstances) ResizeFlavor(ctx context.Context, instanceID string, ds CloudDatabaseDatastore, instanceType, flavorName string) (*CloudDatabaseMessageResponse, error) {
	var iar = CloudDatabaseAction{
		Action:       resizeFlavor,
		Datastore:    ds,
		FlavorName:   flavorName,
		InstanceType: instanceType,
	}

	req, err := ins.client.NewRequest(ctx, http.MethodPost, databaseServiceName, ins.resourceActionPath(instanceID), &iar)
	if err != nil {
		return nil, err
	}

	resp, err := ins.client.Do(ctx, req)
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

// ResizeFlavorSuggestion get suggestion when resize instances.
func (ins *cloudDatabaseInstances) ResizeFlavorSuggestion(ctx context.Context, instanceID string, flavorName string) (*CloudDatabaseSuggestion, error) {
	var iar = CloudDatabaseAction{Action: resizeFlavor, FlavorName: flavorName, Suggestion: true}

	req, err := ins.client.NewRequest(ctx, http.MethodPost, databaseServiceName, ins.resourceActionPath(instanceID), &iar)
	if err != nil {
		return nil, err
	}

	resp, err := ins.client.Do(ctx, req)
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
func (ins *cloudDatabaseInstances) ResizeVolume(ctx context.Context, instanceID string, ds CloudDatabaseDatastore, instanceType string, newSize int) (*CloudDatabaseMessageResponse, error) {
	var iar = CloudDatabaseAction{
		Action:       resizeVolume,
		Datastore:    ds,
		InstanceType: instanceType,
		NewSize:      newSize,
	}

	req, err := ins.client.NewRequest(ctx, http.MethodPost, databaseServiceName, ins.resourceActionPath(instanceID), &iar)
	if err != nil {
		return nil, err
	}

	resp, err := ins.client.Do(ctx, req)
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

// ResizeVolumeSuggestion get suggestion when resize instances.
func (ins *cloudDatabaseInstances) ResizeVolumeSuggestion(ctx context.Context, instanceID string, newSize int) (*CloudDatabaseSuggestion, error) {
	var iar = CloudDatabaseAction{Action: resizeVolume, NewSize: newSize, Suggestion: true}

	req, err := ins.client.NewRequest(ctx, http.MethodPost, databaseServiceName, ins.resourceActionPath(instanceID), &iar)
	if err != nil {
		return nil, err
	}

	resp, err := ins.client.Do(ctx, req)
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

// Delete an instances.
func (ins *cloudDatabaseInstances) Delete(ctx context.Context, instanceID string, idr *CloudDatabaseDelete) (*CloudDatabaseMessageResponse, error) {
	req, err := ins.client.NewRequest(ctx, http.MethodDelete, databaseServiceName, ins.resourcePath(instanceID), &idr)
	if err != nil {
		return nil, err
	}

	resp, err := ins.client.Do(ctx, req)
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

// ListDatabases - List all databases in instances.
func (ins *cloudDatabaseInstances) ListDatabases(ctx context.Context, instanceID string) ([]*CloudDatabaseDB, error) {
	req, err := ins.client.NewRequest(ctx, http.MethodGet, databaseServiceName, ins.resourceTypePath(instanceID, "databases"), nil)
	if err != nil {
		return nil, err
	}

	resp, err := ins.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()
	var data struct {
		Databases []*CloudDatabaseDB `json:"databases"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Databases, nil
}

// CreateDatabases - Create databases in instances.
func (ins *cloudDatabaseInstances) CreateDatabases(ctx context.Context, instanceID string, databases []*CloudDatabaseDB) error {
	req, err := ins.client.NewRequest(ctx, http.MethodPost, databaseServiceName, ins.resourceTypePath(instanceID, "databases"), CloudDatabaseDBReq{Databases: databases})
	if err != nil {
		return err
	}

	resp, err := ins.client.Do(ctx, req)

	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	return nil
}

// DeleteDatabases - Delete databases in instances.
func (ins *cloudDatabaseInstances) DeleteDatabases(ctx context.Context, instanceID string, databases []*CloudDatabaseDB) error {
	req, err := ins.client.NewRequest(ctx, http.MethodDelete, databaseServiceName, ins.resourceTypePath(instanceID, "databases"), CloudDatabaseDBReq{Databases: databases})
	if err != nil {
		return err
	}

	resp, err := ins.client.Do(ctx, req)

	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	return nil
}

// ListUsers - List all users in instances.
func (ins *cloudDatabaseInstances) ListUsers(ctx context.Context, instanceID string) ([]*CloudDatabaseUser, error) {
	req, err := ins.client.NewRequest(ctx, http.MethodGet, databaseServiceName, ins.resourceTypePath(instanceID, "users"), nil)
	if err != nil {
		return nil, err
	}

	resp, err := ins.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()
	var data struct {
		Users []*CloudDatabaseUser `json:"users"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Users, nil
}

// CreateUsers - Create users in instances.
func (ins *cloudDatabaseInstances) CreateUsers(ctx context.Context, instanceID string, users []*CloudDatabaseUser) error {
	req, err := ins.client.NewRequest(ctx, http.MethodPost, databaseServiceName, ins.resourceTypePath(instanceID, "users"), CloudDatabaseUserReq{Users: users})
	if err != nil {
		return err
	}

	resp, err := ins.client.Do(ctx, req)

	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	return nil
}

// ChangePasswordUsers - change password users in instances.
func (ins *cloudDatabaseInstances) ChangePasswordUsers(ctx context.Context, instanceID string, users []*CloudDatabaseUser) error {
	req, err := ins.client.NewRequest(ctx, http.MethodPut, databaseServiceName, ins.resourceTypePath(instanceID, "users"), CloudDatabaseUserReq{Users: users})
	if err != nil {
		return err
	}

	resp, err := ins.client.Do(ctx, req)

	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	return nil
}

// DeleteUsers - Delete users in instances.
func (ins *cloudDatabaseInstances) DeleteUsers(ctx context.Context, instanceID string, users []*CloudDatabaseUser) error {
	req, err := ins.client.NewRequest(ctx, http.MethodDelete, databaseServiceName, ins.resourceTypePath(instanceID, "users"), CloudDatabaseUserReq{Users: users})
	if err != nil {
		return err
	}

	resp, err := ins.client.Do(ctx, req)

	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	return nil
}
