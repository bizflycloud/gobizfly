// This file is part of goBizFly

package gobizfly

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

const (
	cloudDatabaseInstancesResourcePath      = "/instances"
	cloudDatabaseNodesResourcePath          = "/nodes"
	cloudDatabaseTasksResourcePath          = "/tasks"
	cloudDatabaseEnginesResourcePath        = "/engines"
	cloudDatabaseConfigurationsResourcePath = "/configurations"
	cloudDatabaseBackupsResourcePath        = "/backups"
	cloudDatabaseSchedulesResourcePath      = "/schedules"
	cloudDatabaseAutoScalingsResourcePath   = "/autoscaling"
	cloudDatabaseTrustedSourcesResourcePath = "/trusted_sources"
	attach                                  = "attach"
	backups                                 = "backups"
	detach                                  = "detach"
	detachReplica                           = "detach_replica"
	enableRoot                              = "enable_root"
	mariaDB                                 = "MariaDB"
	mongoDB                                 = "MongoDB"
	nodes                                   = "nodes"
	resizeFlavor                            = "resize"
	resizeVolume                            = "resize_volume"
	restart                                 = "restart"
	schedules                               = "schedules"
)

var (
	// ErrRequireFlavorName for resource resize flavor
	ErrRequireFlavorName = errors.New("resize flavor require flavor_name")
	// ErrRequireNewSize for resource resize volume
	ErrRequireNewSize = errors.New("resize volume require new_size")
	// ErrMongoDBReplicas for create resource datastore type mongodb
	ErrMongoDBReplicas = errors.New("MongoDB can't have replica node")
	// ErrMariaDBSecondariesQuantity for create resource datastore type mariadb
	ErrMariaDBSecondariesQuantity = errors.New("MariaDB can't have more than one secondary node")
	// ErrListOption for get resource with list option
	ErrListOption = errors.New("can't determine resource list option")
)

var _ CloudDatabaseService = (*cloudDatabaseService)(nil)

type cloudDatabaseService struct {
	client *Client
}

type CloudDatabaseService interface {
	Instances() *cloudDatabaseInstances
	Nodes() *cloudDatabaseNodes
	Tasks() *cloudDatabaseTasks
	Configurations() *cloudDatabaseConfigurations
	Backups() *cloudDatabaseBackups
	Schedules() *cloudDatabaseSchedules
	AutoScalings() *cloudDatabaseAutoScalings
	TrustedSources() *cloudDatabaseTrustedSources
	Engines() *cloudDatabaseEngines
	EngineParameters() *cloudDatabaseEngineParameters
}

func (db *cloudDatabaseService) Instances() *cloudDatabaseInstances {
	return &cloudDatabaseInstances{client: db.client}
}

func (db *cloudDatabaseService) Nodes() *cloudDatabaseNodes {
	return &cloudDatabaseNodes{client: db.client}
}

func (db *cloudDatabaseService) Tasks() *cloudDatabaseTasks {
	return &cloudDatabaseTasks{client: db.client}
}

func (db *cloudDatabaseService) Configurations() *cloudDatabaseConfigurations {
	return &cloudDatabaseConfigurations{client: db.client}
}

func (db *cloudDatabaseService) Backups() *cloudDatabaseBackups {
	return &cloudDatabaseBackups{client: db.client}
}

func (db *cloudDatabaseService) Schedules() *cloudDatabaseSchedules {
	return &cloudDatabaseSchedules{client: db.client}
}

func (db *cloudDatabaseService) AutoScalings() *cloudDatabaseAutoScalings {
	return &cloudDatabaseAutoScalings{client: db.client}
}

func (db *cloudDatabaseService) TrustedSources() *cloudDatabaseTrustedSources {
	return &cloudDatabaseTrustedSources{client: db.client}
}

func (db *cloudDatabaseService) Engines() *cloudDatabaseEngines {
	return &cloudDatabaseEngines{client: db.client}
}

func (db *cloudDatabaseService) EngineParameters() *cloudDatabaseEngineParameters {
	return &cloudDatabaseEngineParameters{client: db.client}
}

type cloudDatabaseInstances struct {
	client *Client
}

type cloudDatabaseNodes struct {
	client *Client
}

type cloudDatabaseTasks struct {
	client *Client
}

type cloudDatabaseConfigurations struct {
	client *Client
}

type cloudDatabaseBackups struct {
	client *Client
}

type cloudDatabaseSchedules struct {
	client *Client
}

type cloudDatabaseAutoScalings struct {
	client *Client
}

type cloudDatabaseTrustedSources struct {
	client *Client
}

type cloudDatabaseEngines struct {
	client *Client
}

type cloudDatabaseEngineParameters struct {
	client *Client
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

// CloudDatabase Configuration Resource Path
func (cfg *cloudDatabaseConfigurations) resourcePath(cfgID string) string {
	return cloudDatabaseConfigurationsResourcePath + "/" + cfgID
}

// CloudDatabase Configuration Resource Action Path
func (cfg *cloudDatabaseConfigurations) resourceActionPath(nodeID string, cfgID string) string {
	return cloudDatabaseNodesResourcePath + "/" + nodeID + "/configuration/" + cfgID
}

// CloudDatabase Backup Resource Path
func (bk *cloudDatabaseBackups) resourcePath(backupID string) string {
	return cloudDatabaseBackupsResourcePath + "/" + backupID
}

// CloudDatabase Backup resourceType Create Path
func (bk *cloudDatabaseBackups) resourceCreatePath(resourceType string, resourceID string) string {
	return "/" + resourceType + "/" + resourceID + cloudDatabaseBackupsResourcePath
}

// CloudDatabase Schedule Resource List Path
func (sc *cloudDatabaseSchedules) resourcePath(scheduleID string) string {
	return cloudDatabaseSchedulesResourcePath + "/" + scheduleID
}

// CloudDatabase Schedule List In resourceType Path
func (sc *cloudDatabaseSchedules) resourceCreatePath(resourceType string, resourceID string) string {
	return "/" + resourceType + "/" + resourceID + cloudDatabaseSchedulesResourcePath
}

// CloudDatabase Schedule Resource Backup Path
func (sc *cloudDatabaseSchedules) resourceBackupPath(scheduleID string) string {
	return cloudDatabaseSchedulesResourcePath + "/" + scheduleID + cloudDatabaseBackupsResourcePath
}

// CloudDatabase AutoScaling Resource Path
func (au *cloudDatabaseAutoScalings) resourcePath(resourceID string) string {
	return cloudDatabaseInstancesResourcePath + "/" + resourceID + cloudDatabaseAutoScalingsResourcePath
}

// CloudDatabase Trusted Source Resource Path
func (ts *cloudDatabaseTrustedSources) resourcePath(nodeID string) string {
	return cloudDatabaseNodesResourcePath + "/" + nodeID + cloudDatabaseTrustedSourcesResourcePath
}

// CloudDatabase Task Resource Path
func (ta *cloudDatabaseTasks) resourcePath(taskID string) string {
	return cloudDatabaseTasksResourcePath + "/" + taskID + "/status"
}

// CloudDatabase Engine Resource Path
func (en *cloudDatabaseEngineParameters) resourcePath(datastore string, datastoreVersion string) string {
	return cloudDatabaseEnginesResourcePath + "/" + datastore + "/versions/" + datastoreVersion + "/parameters"
}

// CloudDatabaseListOption contains option when list resource
type CloudDatabaseListOption struct {
	DatabaseEngine  string `json:"database_engine,omitempty"`
	DatabaseVersion string `json:"database_version,omitempty"`
	Detailed        bool   `json:"detailed,omitempty"`
	EndTime         string `json:"end_time,omitempty"`
	Name            string `json:"name,omitempty"`
	Page            int    `json:"page,omitempty"`
	ResultsPerPages int    `json:"results_per_pages,omitempty"`
	StartTime       string `json:"start_time,omitempty"`
}

// CloudDatabaseAutoScalingVolume contains volume information of autoscaling.
type CloudDatabaseAutoScalingVolume struct {
	Limited   int `json:"limited"`
	Threshold int `json:"threshold"`
}

// CloudDatabaseAutoScalingReceivers contains information about receivers of autoscaling
type CloudDatabaseAutoScalingReceivers struct {
	Action string `json:"action"`
	ID     string `json:"id"`
	Name   string `json:"name"`
}

// CloudDatabaseAutoScalingAlarms contains information about alarms of autoscaling
type CloudDatabaseAutoScalingAlarms struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	ReceiverID string `json:"receiver_id"`
}

// CloudDatabaseAutoScaling contains autoscaling information of a database instance.
type CloudDatabaseAutoScaling struct {
	Alarms    []CloudDatabaseAutoScalingAlarms    `json:"alarms,omitempty"`
	Enable    bool                                `json:"enable,omitempty"`
	Receivers []CloudDatabaseAutoScalingReceivers `json:"receivers,omitempty"`
	Volume    CloudDatabaseAutoScalingVolume      `json:"volume,omitempty"`
}

// CloudDatabaseDNS contains DNS information.
type CloudDatabaseDNS struct {
	Private string `json:"private"`
	Public  string `json:"public"`
}

// CloudDatabaseAddressesDetail contains detail addresses information of database.
type CloudDatabaseAddressesDetail struct {
	IPAddress   string `json:"ip_address"`
	NetworkName string `json:"network_name"`
}

// CloudDatabaseAddresses contains addresses information of database.
type CloudDatabaseAddresses struct {
	Private []CloudDatabaseAddressesDetail `json:"private"`
	Public  []CloudDatabaseAddressesDetail `json:"public"`
}

// CloudDatabaseDatastore contains datastore information.
type CloudDatabaseDatastore struct {
	Type      string `json:"type,omitempty"`
	Name      string `json:"name,omitempty"`
	ID        string `json:"id,omitempty"`
	VersionID string `json:"version_id,omitempty"`
}

// CloudDatabaseAction contains database instance action information.
type CloudDatabaseAction struct {
	Action     string   `json:"action,omitempty" validate:"required"`
	NewSize    int      `json:"new_size,omitempty"`
	FlavorName string   `json:"flavor_name,omitempty"`
	ResizeAll  bool     `json:"resize_all,omitempty"`
	ActionAll  bool     `json:"action_all,omitempty"`
	NodeType   string   `json:"node_type,omitempty"`
	NodeIDs    []string `json:"node_ids,omitempty"`
	Suggestion bool     `json:"suggestion,omitempty"`
}

// CloudDatabaseDelete contains database delete information.
type CloudDatabaseDelete struct {
	PurgeBackup     bool `json:"purge_backup"`
	PurgeAutobackup bool `json:"purge_autobackup"`
}

type CloudDatabaseReplicasConfiguration struct {
	AvailabilityZone string `json:"availability_zone"`
	Region           string `json:"region"`
}

type CloudDatabaseSecondariesConfiguration struct {
	AvailabilityZone string `json:"availability_zone"`
	Region           string `json:"region"`
}

type CloudDatabaseReplicas struct {
	Configurations CloudDatabaseReplicasConfiguration `json:"configurations"`
	Quantity       int                                `json:"quantity"`
}

type CloudDatabaseSecondaries struct {
	Configurations CloudDatabaseSecondariesConfiguration `json:"configurations"`
	Quantity       int                                   `json:"quantity"`
}

// CloudDatabaseNodeCreate contains detail to create database node payload.
type CloudDatabaseNodeCreate struct {
	ReplicaOf   string                 `json:"replica_of,omitempty" validate:"required"`
	Replicas    *CloudDatabaseReplicas `json:"replicas,omitempty"`
	Role        string                 `json:"role,omitempty"`
	Secondaries *CloudDatabaseReplicas `json:"secondaries,omitempty"`
	Suggestion  bool                   `json:"suggestion,omitempty"`
}

// CloudDatabaseVolume contains info of volume.
type CloudDatabaseVolume struct {
	Size int     `json:"size"`
	Used float32 `json:"used"`
}

// CloudDatabaseNode contains detail of a database node.
type CloudDatabaseNode struct {
	Addresses        CloudDatabaseAddresses `json:"addresses"`
	AvailabilityZone string                 `json:"availability_zone"`
	CreatedAt        string                 `json:"created_at"`
	Datastore        CloudDatabaseDatastore `json:"datastore"`
	Description      string                 `json:"description"`
	DNS              CloudDatabaseDNS       `json:"dns"`
	Flavor           string                 `json:"flavor.id"`
	ID               string                 `json:"id"`
	InstanceID       string                 `json:"instance_id"`
	Message          string                 `json:"message"`
	Name             string                 `json:"name"`
	RegionName       string                 `json:"region_name"`
	Replicas         []CloudDatabaseNode    `json:"replicas"`
	Role             string                 `json:"role"`
	Status           string                 `json:"status"`
	TaskID           string                 `json:"task_id"`
	Volume           CloudDatabaseVolume    `json:"volume"`
}

// CloudDatabase Instance Struct

// CloudDatabaseNetworks contains network information to create database instance.
type CloudDatabaseNetworks struct {
	NetworkID string `json:"network_id"`
}

// CloudDatabaseInstanceCreate contains payload to create database instance.
type CloudDatabaseInstanceCreate struct {
	AutoScaling      *CloudDatabaseAutoScaling `json:"autoscaling,omitempty"`
	AvailabilityZone string                    `json:"availability_zone,omitempty" validate:"required"`
	BackupID         string                    `json:"backup_id,omitempty"`
	Datastore        CloudDatabaseDatastore    `json:"datastore,omitempty" validate:"required"`
	EnableFailover   bool                      `json:"enable_failover,omitempty"`
	FlavorName       string                    `json:"flavor_name,omitempty" validate:"required"`
	InstanceType     string                    `json:"instance_type,omitempty"`
	Name             string                    `json:"name,omitempty" validate:"required"`
	Networks         []CloudDatabaseNetworks   `json:"networks,omitempty" validate:"required"`
	PublicAccess     bool                      `json:"public_access,omitempty"`
	Replicas         *CloudDatabaseReplicas    `json:"replicas,omitempty"`
	Secondaries      *CloudDatabaseSecondaries `json:"secondaries,omitempty"`
	Suggestion       bool                      `json:"suggestion,omitempty"`
	VolumeSize       int                       `json:"volume_size,omitempty" validate:"required"`
}

type CloudDatabaseLog struct {
	Enable string `json:"enable"`
	Name   string `json:"name"`
}

// CloudDatabaseInstance contains database instance information.
type CloudDatabaseInstance struct {
	AutoScaling  CloudDatabaseAutoScaling `json:"autoscaling"`
	CreatedAt    string                   `json:"created"`
	Datastore    CloudDatabaseDatastore   `json:"datastore"`
	Description  string                   `json:"description"`
	ID           string                   `json:"id"`
	Logs         CloudDatabaseLog         `json:"logs"`
	Message      string                   `json:"message"`
	Name         string                   `json:"name"`
	Networks     []CloudDatabaseNetworks  `json:"networks"`
	Nodes        []CloudDatabaseNode      `json:"nodes"`
	ProjectID    string                   `json:"project_id"`
	PublicAccess bool                     `json:"public_access"`
	TaskID       string                   `json:"task_id"`
}

// CloudDatabase Configuration Struct

// CloudDatabaseConfigurationCreate contains payload create configuration.
type CloudDatabaseConfigurationCreate struct {
	ConfigurationName       string                 `json:"configuration_name,omitempty" validate:"required"`
	ConfigurationParameters map[string]interface{} `json:"configuration_parameters,omitempty" validate:"required"`
	Datastore               CloudDatabaseDatastore `json:"datastore,omitempty" validate:"required"`
}

// CloudDatabaseNodeUseConfiguration contains info of node use configuration.
type CloudDatabaseNodeUseConfiguration struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// CloudDatabaseConfiguration contains detail of a configuration.
type CloudDatabaseConfiguration struct {
	CreatedAt string                              `json:"created_at"`
	Datastore CloudDatabaseDatastore              `json:"datastore"`
	ID        string                              `json:"id"`
	Message   string                              `json:"message"`
	Name      string                              `json:"name"`
	NodeCount int                                 `json:"node_count"`
	Nodes     []CloudDatabaseNodeUseConfiguration `json:"nodes"`
	Values    map[string]interface{}              `json:"values"`
}

// CloudDatabaseConfigurationUpdate contains payload to update a configuration.
type CloudDatabaseConfigurationUpdate struct {
	ConfigurationParameters map[string]interface{} `json:"configuration_parameters" validate:"required"`
}

// CloudDatabase Backup Struct

// CloudDatabaseDatastoreBackup contains datastore info of a backup.
type CloudDatabaseDatastoreBackup struct {
	Type      string `json:"type"`
	Version   string `json:"version"`
	VersionID string `json:"version_id"`
}

// CloudDatabaseBackup contains detail of a backup.
type CloudDatabaseBackup struct {
	Created     string                       `json:"created"`
	Datastore   CloudDatabaseDatastoreBackup `json:"datastore"`
	Description string                       `json:"description"`
	ID          string                       `json:"id"`
	Message     string                       `json:"message"`
	Name        string                       `json:"name"`
	NodeID      string                       `json:"node_id"`
	ParentID    string                       `json:"parent_id"`
	ProjectID   string                       `json:"project_id"`
	Size        float32                      `json:"size"`
	Status      string                       `json:"status"`
	Type        string                       `json:"type"`
	Updated     string                       `json:"updated"`
}

// CloudDatabaseBackupCreate contains payload require create backup.
type CloudDatabaseBackupCreate struct {
	BackupName string `json:"backup_name,omitempty" validate:"required"`
	NodeID     string `json:"node_id,omitempty"`
	ParentID   string `json:"parent_id,omitempty"`
	Suggestion bool   `json:"suggestion"`
}

// CloudDatabaseBackupResource contains option list backup.
type CloudDatabaseBackupResource struct {
	ResourceID   string `json:"resource_id"`
	ResourceType string `json:"resource_type"`
}

// CloudDatabase Schedule Backup Struct

// CloudDatabaseSchedule contains detail of a schedule.
type CloudDatabaseSchedule struct {
	FirstExecutionTime  string `json:"first_execution_time"`
	ID                  string `json:"id"`
	InstanceID          string `json:"instance_id"`
	LimitBackup         int    `json:"limit_backup"`
	Message             string `json:"message"`
	Name                string `json:"name"`
	NextExecutionTime   string `json:"next_execution_time"`
	NodeID              string `json:"node_id"`
	NodeName            string `json:"node_name"`
	Pattern             string `json:"pattern"`
	ProjectID           string `json:"project_id"`
	RemainingExecutions int    `json:"remaining_executions"`
	WorkflowID          string `json:"workflow_id"`
	WorkflowName        string `json:"workflow_name"`
}

// CloudDatabaseScheduleCreate contains schedule create payload info.
type CloudDatabaseScheduleCreate struct {
	DayOfMonth   []int  `json:"day_of_month,omitempty"`
	DayOfWeek    []int  `json:"day_of_week,omitempty"`
	Hour         []int  `json:"hour,omitempty"`
	LimitBackup  int    `json:"limit_backup,omitempty" validate:"required"`
	Minute       []int  `json:"minute,omitempty" validate:"required"`
	ScheduleName string `json:"schedule_name,omitempty" validate:"required"`
	ScheduleType string `json:"schedule_type,omitempty" validate:"required"`
}

// CloudDatabaseScheduleDelete contains option when delete a schedule.
type CloudDatabaseScheduleDelete struct {
	PurgeBackup bool `json:"purge_backup"`
}

// CloudDatabaseScheduleListResourceOption contains option when list a database schedule.
type CloudDatabaseScheduleListResourceOption struct {
	All          bool   `json:"all,omitempty"`
	ListBackup   bool   `json:"list_backup,omitempty"`
	ResourceID   string `json:"resource_id,omitempty"`
	ResourceType string `json:"resource_type,omitempty"`
}

// CloudDatabase TrustedSource Struct

// CloudDatabaseTrustedSources contains TrustedSource information of node.
type CloudDatabaseTrustedSources struct {
	TrustedSources []string `json:"trusted_sources"`
}

// CloudDatabase Task Struct

// CloudDatabaseTaskResult contains task result detail.
type CloudDatabaseTaskResult struct {
	Action   string                 `json:"action"`
	Data     map[string]interface{} `json:"data"`
	Message  string                 `json:"message"`
	Progress int                    `json:"progress"`
}

// CloudDatabaseTask contains task result.
type CloudDatabaseTask struct {
	Ready  bool                    `json:"ready"`
	Result CloudDatabaseTaskResult `json:"result"`
}

// CloudDatabase Engine Struct

// CloudDatabaseEngineVersion contains datastore version detail.
type CloudDatabaseEngineVersion struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

// CloudDatabaseEngine contains detail about a datastore.
type CloudDatabaseEngine struct {
	ID       string                       `json:"id"`
	Name     string                       `json:"name"`
	Versions []CloudDatabaseEngineVersion `json:"versions"`
}

// CloudDatabaseEngineParameters contains datastore parameters info.
type CloudDatabaseEngineParameters struct {
	ConfigurationParameters []map[string]interface{} `json:"configuration_parameters"`
}

// CloudDatabase Engine Message Struct

// CloudDatabaseMessageResponse contains message response from Database Service API.
type CloudDatabaseMessageResponse struct {
	Message string `json:"message"`
	TaskID  string `json:"task_id"`
}

// CloudDatabase Instance

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

	defer resp.Body.Close()
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

	defer resp.Body.Close()
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

	defer resp.Body.Close()
	var data struct {
		Backups []*CloudDatabaseBackup `json:"backups"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Backups, nil
}

// ListSchedules List all schedule in instances.
func (ins *cloudDatabaseInstances) ListSchedules(ctx context.Context, instanceID string, opts *CloudDatabaseListOption) ([]*CloudDatabaseSchedule, error) {
	req, err := ins.client.NewRequest(ctx, http.MethodGet, databaseServiceName, ins.resourceTypePath(instanceID, schedules), nil)
	if err != nil {
		return nil, err
	}

	AddParamsListOption(req, opts)

	resp, err := ins.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var data struct {
		Schedules []*CloudDatabaseSchedule `json:"schedules"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Schedules, nil
}

// Create a new instances
func (ins *cloudDatabaseInstances) Create(ctx context.Context, icr *CloudDatabaseInstanceCreate) (*CloudDatabaseInstance, error) {
	if icr.Datastore.Type == mongoDB && icr.Replicas != nil {
		return nil, ErrMongoDBReplicas
	}

	if icr.Datastore.Type == mariaDB && icr.Secondaries != nil && icr.Secondaries.Quantity > 1 {
		return nil, ErrMariaDBSecondariesQuantity
	}

	req, err := ins.client.NewRequest(ctx, http.MethodPost, databaseServiceName, cloudDatabaseInstancesResourcePath, &icr)
	if err != nil {
		return nil, err
	}

	resp, err := ins.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var instances *CloudDatabaseInstance
	if err := json.NewDecoder(resp.Body).Decode(&instances); err != nil {
		return nil, err
	}
	return instances, nil
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

	defer resp.Body.Close()
	var instances *CloudDatabaseInstance

	if err := json.NewDecoder(resp.Body).Decode(&instances); err != nil {
		return nil, err
	}

	return instances, nil
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

	defer resp.Body.Close()
	var dmr *CloudDatabaseMessageResponse

	if err := json.NewDecoder(resp.Body).Decode(&dmr); err != nil {
		return nil, err
	}

	return dmr, nil
}

// ResizeFlavor with an instances.
func (ins *cloudDatabaseInstances) ResizeFlavor(ctx context.Context, instanceID string, flavorName string) (*CloudDatabaseMessageResponse, error) {
	var iar = CloudDatabaseAction{Action: resizeFlavor, FlavorName: flavorName}

	req, err := ins.client.NewRequest(ctx, http.MethodPost, databaseServiceName, ins.resourceActionPath(instanceID), &iar)
	if err != nil {
		return nil, err
	}

	resp, err := ins.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var dmr *CloudDatabaseMessageResponse

	if err := json.NewDecoder(resp.Body).Decode(&dmr); err != nil {
		return nil, err
	}

	return dmr, nil
}

// ResizeVolume with an instances.
func (ins *cloudDatabaseInstances) ResizeVolume(ctx context.Context, instanceID string, newSize int) (*CloudDatabaseMessageResponse, error) {
	var iar = CloudDatabaseAction{Action: resizeVolume, NewSize: newSize}

	req, err := ins.client.NewRequest(ctx, http.MethodPost, databaseServiceName, ins.resourceActionPath(instanceID), &iar)
	if err != nil {
		return nil, err
	}

	resp, err := ins.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var dmr *CloudDatabaseMessageResponse

	if err := json.NewDecoder(resp.Body).Decode(&dmr); err != nil {
		return nil, err
	}

	return dmr, nil
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

	defer resp.Body.Close()
	var dmr *CloudDatabaseMessageResponse

	if err := json.NewDecoder(resp.Body).Decode(&dmr); err != nil {
		return nil, err
	}

	return dmr, nil
}

// CloudDatabase Node

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

	defer resp.Body.Close()
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

	defer resp.Body.Close()
	var data struct {
		Backups []*CloudDatabaseBackup `json:"backups"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Backups, nil
}

// ListSchedules List all schedule in nodes.
func (no *cloudDatabaseNodes) ListSchedules(ctx context.Context, nodeID string, opts *CloudDatabaseListOption) ([]*CloudDatabaseSchedule, error) {
	req, err := no.client.NewRequest(ctx, http.MethodGet, databaseServiceName, no.resourceTypePath(nodeID, schedules), nil)
	if err != nil {
		return nil, err
	}

	AddParamsListOption(req, opts)

	resp, err := no.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var data struct {
		Schedules []*CloudDatabaseSchedule `json:"schedules"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Schedules, nil
}

// Create a new replica or secondary nodes
func (no *cloudDatabaseNodes) Create(ctx context.Context, icr *CloudDatabaseNodeCreate) (*CloudDatabaseNode, error) {
	req, err := no.client.NewRequest(ctx, http.MethodPost, databaseServiceName, cloudDatabaseNodesResourcePath, &icr)
	if err != nil {
		return nil, err
	}

	resp, err := no.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var nodes *CloudDatabaseNode

	if err := json.NewDecoder(resp.Body).Decode(&nodes); err != nil {
		return nil, err
	}

	return nodes, nil
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

	defer resp.Body.Close()
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

	defer resp.Body.Close()
	var dmr *CloudDatabaseMessageResponse

	if err := json.NewDecoder(resp.Body).Decode(&dmr); err != nil {
		return nil, err
	}

	return dmr, nil
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

	defer resp.Body.Close()
	var dmr *CloudDatabaseMessageResponse

	if err := json.NewDecoder(resp.Body).Decode(&dmr); err != nil {
		return nil, err
	}

	return dmr, nil
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

	defer resp.Body.Close()
	var dmr *CloudDatabaseMessageResponse

	if err := json.NewDecoder(resp.Body).Decode(&dmr); err != nil {
		return nil, err
	}

	return dmr, nil
}

// Restart with an instances.
func (no *cloudDatabaseNodes) Restart(ctx context.Context, nodeID string) (*CloudDatabaseMessageResponse, error) {
	var iar = CloudDatabaseAction{Action: restart}

	req, err := no.client.NewRequest(ctx, http.MethodPost, databaseServiceName, no.resourceActionPath(nodeID), &iar)
	if err != nil {
		return nil, err
	}

	resp, err := no.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
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

	defer resp.Body.Close()
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

	defer resp.Body.Close()
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

	defer resp.Body.Close()
	var dmr *CloudDatabaseMessageResponse

	if err := json.NewDecoder(resp.Body).Decode(&dmr); err != nil {
		return nil, err
	}

	return dmr, nil
}

// CloudDatabase Configuration

// List all configurations.
func (cfg *cloudDatabaseConfigurations) List(ctx context.Context, opts *CloudDatabaseListOption) ([]*CloudDatabaseConfiguration, error) {
	req, err := cfg.client.NewRequest(ctx, http.MethodGet, databaseServiceName, cloudDatabaseConfigurationsResourcePath, nil)
	if err != nil {
		return nil, err
	}

	AddParamsListOption(req, opts)

	resp, err := cfg.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var data struct {
		Configurations []*CloudDatabaseConfiguration `json:"configurations"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Configurations, nil
}

// Create a new configurations.
func (cfg *cloudDatabaseConfigurations) Create(ctx context.Context, cr *CloudDatabaseConfigurationCreate) (*CloudDatabaseConfiguration, error) {
	req, err := cfg.client.NewRequest(ctx, http.MethodPost, databaseServiceName, cloudDatabaseConfigurationsResourcePath, &cr)
	if err != nil {
		return nil, err
	}

	resp, err := cfg.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var configuration *CloudDatabaseConfiguration
	if err := json.NewDecoder(resp.Body).Decode(&configuration); err != nil {
		return nil, err
	}
	return configuration, nil
}

// Get a configurations.
func (cfg *cloudDatabaseConfigurations) Get(ctx context.Context, cfgID string) (*CloudDatabaseConfiguration, error) {
	req, err := cfg.client.NewRequest(ctx, http.MethodGet, databaseServiceName, cfg.resourcePath(cfgID), nil)
	if err != nil {
		return nil, err
	}

	resp, err := cfg.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var configuration *CloudDatabaseConfiguration

	if err := json.NewDecoder(resp.Body).Decode(&configuration); err != nil {
		return nil, err
	}

	return configuration, nil
}

// Action with a configurations to node.
func (cfg *cloudDatabaseConfigurations) Action(ctx context.Context, nodeID string, cfgID string, iar *CloudDatabaseAction) (*CloudDatabaseMessageResponse, error) {
	req, err := cfg.client.NewRequest(ctx, http.MethodPost, databaseServiceName, cfg.resourceActionPath(nodeID, cfgID), &iar)
	if err != nil {
		return nil, err
	}

	resp, err := cfg.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var dmr *CloudDatabaseMessageResponse

	if err := json.NewDecoder(resp.Body).Decode(&dmr); err != nil {
		return nil, err
	}

	return dmr, nil
}

// Attach with a configurations to node.
func (cfg *cloudDatabaseConfigurations) Attach(ctx context.Context, nodeID string, cfgID string, all bool) (*CloudDatabaseMessageResponse, error) {
	var iar = CloudDatabaseAction{Action: attach, ActionAll: all}

	req, err := cfg.client.NewRequest(ctx, http.MethodPost, databaseServiceName, cfg.resourceActionPath(nodeID, cfgID), &iar)
	if err != nil {
		return nil, err
	}

	resp, err := cfg.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var dmr *CloudDatabaseMessageResponse

	if err := json.NewDecoder(resp.Body).Decode(&dmr); err != nil {
		return nil, err
	}

	return dmr, nil
}

// Detach with a configurations to node.
func (cfg *cloudDatabaseConfigurations) Detach(ctx context.Context, nodeID string, cfgID string, all bool) (*CloudDatabaseMessageResponse, error) {
	var iar = CloudDatabaseAction{Action: detach, ActionAll: all}

	req, err := cfg.client.NewRequest(ctx, http.MethodPost, databaseServiceName, cfg.resourceActionPath(nodeID, cfgID), &iar)
	if err != nil {
		return nil, err
	}

	resp, err := cfg.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var dmr *CloudDatabaseMessageResponse

	if err := json.NewDecoder(resp.Body).Decode(&dmr); err != nil {
		return nil, err
	}

	return dmr, nil
}

// Update a configurations.
func (cfg *cloudDatabaseConfigurations) Update(ctx context.Context, cfgID string, cu *CloudDatabaseConfigurationUpdate) (*CloudDatabaseMessageResponse, error) {
	req, err := cfg.client.NewRequest(ctx, http.MethodPut, databaseServiceName, cfg.resourcePath(cfgID), &cu)
	if err != nil {
		return nil, err
	}

	resp, err := cfg.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var dmr *CloudDatabaseMessageResponse

	if err := json.NewDecoder(resp.Body).Decode(&dmr); err != nil {
		return nil, err
	}

	return dmr, nil
}

// Delete a configurations.
func (cfg *cloudDatabaseConfigurations) Delete(ctx context.Context, cfgID string) (*CloudDatabaseMessageResponse, error) {
	req, err := cfg.client.NewRequest(ctx, http.MethodDelete, databaseServiceName, cfg.resourcePath(cfgID), nil)
	if err != nil {
		return nil, err
	}

	resp, err := cfg.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var dmr *CloudDatabaseMessageResponse

	if err := json.NewDecoder(resp.Body).Decode(&dmr); err != nil {
		return nil, err
	}

	return dmr, nil
}

// CloudDatabase Backup

// List all backups.
func (bk *cloudDatabaseBackups) List(ctx context.Context, resource *CloudDatabaseBackupResource, opts *CloudDatabaseListOption) ([]*CloudDatabaseBackup, error) {
	resourcePath := cloudDatabaseBackupsResourcePath
	if resource != nil {
		resourcePath = bk.resourceCreatePath(resource.ResourceType, resource.ResourceID)
	}

	req, err := bk.client.NewRequest(ctx, http.MethodGet, databaseServiceName, resourcePath, nil)
	if err != nil {
		return nil, err
	}

	AddParamsListOption(req, opts)

	resp, err := bk.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var data struct {
		Backups []*CloudDatabaseBackup `json:"backups"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Backups, nil
}

// Create a backups.
func (bk *cloudDatabaseBackups) Create(ctx context.Context, resourceType string, resourceID string, cr *CloudDatabaseBackupCreate) (*CloudDatabaseBackup, error) {
	req, err := bk.client.NewRequest(ctx, http.MethodPost, databaseServiceName, bk.resourceCreatePath(resourceType, resourceID), &cr)
	if err != nil {
		return nil, err
	}

	resp, err := bk.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var backups *CloudDatabaseBackup

	if err := json.NewDecoder(resp.Body).Decode(&backups); err != nil {
		return nil, err
	}

	return backups, nil
}

// Get a backups.
func (bk *cloudDatabaseBackups) Get(ctx context.Context, backupID string) (*CloudDatabaseBackup, error) {
	req, err := bk.client.NewRequest(ctx, http.MethodGet, databaseServiceName, bk.resourcePath(backupID), nil)
	if err != nil {
		return nil, err
	}

	resp, err := bk.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var backups *CloudDatabaseBackup

	if err := json.NewDecoder(resp.Body).Decode(&backups); err != nil {
		return nil, err
	}

	return backups, nil
}

// Delete a backups.
func (bk *cloudDatabaseBackups) Delete(ctx context.Context, backupID string) (*CloudDatabaseMessageResponse, error) {
	req, err := bk.client.NewRequest(ctx, http.MethodDelete, databaseServiceName, bk.resourcePath(backupID), nil)
	if err != nil {
		return nil, err
	}

	resp, err := bk.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var dmr *CloudDatabaseMessageResponse

	if err := json.NewDecoder(resp.Body).Decode(&dmr); err != nil {
		return nil, err
	}

	return dmr, nil
}

// CloudDatabase schedule

// List all schedule.
func (sc *cloudDatabaseSchedules) List(ctx context.Context, resource *CloudDatabaseScheduleListResourceOption, opts *CloudDatabaseListOption) ([]*CloudDatabaseSchedule, error) {
	var resourcePath string

	switch {
	case resource.All:
		resourcePath = cloudDatabaseSchedulesResourcePath
	case resource.ListBackup && resource.ResourceID != "":
		resourcePath = sc.resourceBackupPath(resource.ResourceID)
	case resource.ResourceType != "" && resource.ResourceID != "":
		resourcePath = sc.resourceCreatePath(resource.ResourceType, resource.ResourceID)
	default:
		return nil, ErrListOption
	}

	req, err := sc.client.NewRequest(ctx, http.MethodGet, databaseServiceName, resourcePath, nil)
	if err != nil {
		return nil, err
	}

	AddParamsListOption(req, opts)

	resp, err := sc.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var data struct {
		Schedules []*CloudDatabaseSchedule `json:"schedules"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Schedules, nil
}

// ListBackups list all backup in schedule.
func (sc *cloudDatabaseSchedules) ListBackups(ctx context.Context, scheduleID string, opts *CloudDatabaseListOption) ([]*CloudDatabaseBackup, error) {
	req, err := sc.client.NewRequest(ctx, http.MethodGet, databaseServiceName, sc.resourceBackupPath(scheduleID), nil)
	if err != nil {
		return nil, err
	}

	AddParamsListOption(req, opts)

	resp, err := sc.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var data struct {
		Backups []*CloudDatabaseBackup `json:"backups"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Backups, nil
}

// Create new schedule.
func (sc *cloudDatabaseSchedules) Create(ctx context.Context, nodeID string, scc *CloudDatabaseScheduleCreate) (*CloudDatabaseSchedule, error) {
	req, err := sc.client.NewRequest(ctx, http.MethodPost, databaseServiceName, sc.resourceCreatePath("nodes", nodeID), &scc)
	if err != nil {
		return nil, err
	}

	resp, err := sc.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var schedules *CloudDatabaseSchedule

	if err := json.NewDecoder(resp.Body).Decode(&schedules); err != nil {
		return nil, err
	}

	return schedules, nil
}

// Get a schedule.
func (sc *cloudDatabaseSchedules) Get(ctx context.Context, scheduleID string) (*CloudDatabaseSchedule, error) {
	req, err := sc.client.NewRequest(ctx, http.MethodGet, databaseServiceName, sc.resourcePath(scheduleID), nil)
	if err != nil {
		return nil, err
	}

	resp, err := sc.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var schedules *CloudDatabaseSchedule

	if err := json.NewDecoder(resp.Body).Decode(&schedules); err != nil {
		return nil, err
	}

	return schedules, nil
}

// Delete a schedule.
func (sc *cloudDatabaseSchedules) Delete(ctx context.Context, scheduleID string, option *CloudDatabaseScheduleDelete) (*CloudDatabaseMessageResponse, error) {
	req, err := sc.client.NewRequest(ctx, http.MethodDelete, databaseServiceName, sc.resourcePath(scheduleID), option)
	if err != nil {
		return nil, err
	}

	resp, err := sc.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var dmr *CloudDatabaseMessageResponse

	if err := json.NewDecoder(resp.Body).Decode(&dmr); err != nil {
		return nil, err
	}

	return dmr, nil
}

// CloudDatabase autoscaling

// Create a new autoscaling.
func (au *cloudDatabaseAutoScalings) Create(ctx context.Context, instanceID string, option *CloudDatabaseAutoScaling) (*CloudDatabaseMessageResponse, error) {
	req, err := au.client.NewRequest(ctx, http.MethodPost, databaseServiceName, au.resourcePath(instanceID), option)
	if err != nil {
		return nil, err
	}

	resp, err := au.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var dmr *CloudDatabaseMessageResponse
	if err := json.NewDecoder(resp.Body).Decode(&dmr); err != nil {
		return nil, err
	}
	return dmr, nil
}

// Update a autoscaling.
func (au *cloudDatabaseAutoScalings) Update(ctx context.Context, instanceID string, option *CloudDatabaseAutoScaling) (*CloudDatabaseMessageResponse, error) {
	req, err := au.client.NewRequest(ctx, http.MethodPut, databaseServiceName, au.resourcePath(instanceID), option)
	if err != nil {
		return nil, err
	}

	resp, err := au.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var dmr *CloudDatabaseMessageResponse

	if err := json.NewDecoder(resp.Body).Decode(&dmr); err != nil {
		return nil, err
	}

	return dmr, nil
}

// Delete a autoscaling.
func (au *cloudDatabaseAutoScalings) Delete(ctx context.Context, instanceID string) (*CloudDatabaseMessageResponse, error) {
	req, err := au.client.NewRequest(ctx, http.MethodDelete, databaseServiceName, au.resourcePath(instanceID), nil)
	if err != nil {
		return nil, err
	}

	resp, err := au.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var dmr *CloudDatabaseMessageResponse

	if err := json.NewDecoder(resp.Body).Decode(&dmr); err != nil {
		return nil, err
	}

	return dmr, nil
}

// CloudDatabase TrustedSource

// Get TrustedSource of a node.
func (ts *cloudDatabaseTrustedSources) Get(ctx context.Context, nodeID string) (*CloudDatabaseTrustedSources, error) {
	req, err := ts.client.NewRequest(ctx, http.MethodGet, databaseServiceName, ts.resourcePath(nodeID), nil)
	if err != nil {
		return nil, err
	}

	resp, err := ts.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var trustedSources *CloudDatabaseTrustedSources

	if err := json.NewDecoder(resp.Body).Decode(&trustedSources); err != nil {
		return nil, err
	}

	return trustedSources, nil
}

// Update a TrustedSource.
func (ts *cloudDatabaseTrustedSources) Update(ctx context.Context, nodeID string, tsc *CloudDatabaseTrustedSources) (*CloudDatabaseTrustedSources, error) {
	req, err := ts.client.NewRequest(ctx, http.MethodPost, databaseServiceName, ts.resourcePath(nodeID), &tsc)
	if err != nil {
		return nil, err
	}

	resp, err := ts.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var trustedSources *CloudDatabaseTrustedSources

	if err := json.NewDecoder(resp.Body).Decode(&trustedSources); err != nil {
		return nil, err
	}

	return trustedSources, nil
}

// CloudDatabase Task status

// Get a task status.
func (ta *cloudDatabaseTasks) Get(ctx context.Context, taskID string) (*CloudDatabaseTask, error) {
	req, err := ta.client.NewRequest(ctx, http.MethodGet, databaseServiceName, ta.resourcePath(taskID), nil)
	if err != nil {
		return nil, err
	}

	resp, err := ta.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var task *CloudDatabaseTask

	if err := json.NewDecoder(resp.Body).Decode(&task); err != nil {
		return nil, err
	}

	return task, nil
}

// CloudDatabase Engine

// List all database engine.
func (en *cloudDatabaseEngines) List(ctx context.Context) ([]*CloudDatabaseEngine, error) {
	req, err := en.client.NewRequest(ctx, http.MethodGet, databaseServiceName, cloudDatabaseEnginesResourcePath, nil)
	if err != nil {
		return nil, err
	}

	resp, err := en.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var data struct {
		Engines []*CloudDatabaseEngine `json:"engines"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Engines, nil
}

// Get a database engine parameters.
func (en *cloudDatabaseEngineParameters) Get(ctx context.Context, datastore string, datastoreVersion string) (*CloudDatabaseEngineParameters, error) {
	req, err := en.client.NewRequest(ctx, http.MethodGet, databaseServiceName, en.resourcePath(datastore, datastoreVersion), nil)
	if err != nil {
		return nil, err
	}

	resp, err := en.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var engineParameters *CloudDatabaseEngineParameters

	if err := json.NewDecoder(resp.Body).Decode(&engineParameters); err != nil {
		return nil, err
	}

	return engineParameters, nil
}

// Utils Function

// AddParamsListOption update params option when list resource.
func AddParamsListOption(req *http.Request, opts *CloudDatabaseListOption) {
	params := req.URL.Query()
	if opts.Page != 0 {
		params.Add("page", strconv.Itoa(opts.Page))
	}
	if opts.ResultsPerPages != 0 {
		params.Add("results_per_page", strconv.Itoa(opts.ResultsPerPages))
	}
	if opts.Name != "" {
		params.Add("name", opts.Name)
	}
	if opts.StartTime != "" {
		params.Add("start_time", opts.Name)
	}
	if opts.EndTime != "" {
		params.Add("end_time", opts.Name)
	}
	if opts.DatabaseEngine != "" {
		params.Add("database_engine", opts.Name)
	}
	if opts.DatabaseVersion != "" {
		params.Add("database_version", opts.Name)
	}
	if strconv.FormatBool(opts.Detailed) != "" {
		params.Add("detailed", strconv.FormatBool(opts.Detailed))
	}
	req.URL.RawQuery = params.Encode()
}
