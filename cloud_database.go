// This file is part of gobizfly

package gobizfly

import (
	"errors"
	"net/http"
	"strconv"
)

const (
	attach        = "attach"
	backups       = "backups"
	detach        = "detach"
	detachReplica = "detach_replica"
	enableRoot    = "enable_root"
	mariaDB       = "MariaDB"
	mongoDB       = "MongoDB"
	nodes         = "nodes"
	resizeFlavor  = "resize"
	resizeVolume  = "resize_volume"
	restart       = "restart"
	schedules     = "schedules"
)

var (
	// ErrRequireFlavorName for resource resize flavor
	ErrRequireFlavorName = errors.New("resize flavor require flavor_name")
	// ErrRequireNewSize for resource resize volume
	ErrRequireNewSize = errors.New("resize volume require new_size")
	// ErrListOption for get resource with list option
	ErrListOption = errors.New("can't determine resource list option")
)

var _ CloudDatabaseService = (*cloudDatabaseService)(nil)

type cloudDatabaseService struct {
	client *Client
}

// CloudDatabaseService is an interface to interact with service
type CloudDatabaseService interface {
	AutoScalings() *cloudDatabaseAutoScalings
	Backups() *cloudDatabaseBackups
	BackupSchedules() *cloudDatabaseBackupSchedules
	Configurations() *cloudDatabaseConfigurations
	EngineParameters() *cloudDatabaseEngineParameters
	Engines() *cloudDatabaseEngines
	Flavors() *cloudDatabaseFlavors
	Instances() *cloudDatabaseInstances
	Nodes() *cloudDatabaseNodes
	Tasks() *cloudDatabaseTasks
	TrustedSources() *cloudDatabaseTrustedSources
}

// CloudDatabaseDNS contains DNS information.
type CloudDatabaseDNS struct {
	Private string `json:"private"`
	Public  string `json:"public"`
	SRV     string `json:"srv"`
}

// CloudDatabaseAddressesDetail contains detail addresses information of database.
type CloudDatabaseAddressesDetail struct {
	IPAddress string `json:"ip_address"`
	Network   string `json:"network_name"`
	Port      int    `json:"port"`
}

// CloudDatabaseAddresses contains addresses information of database.
type CloudDatabaseAddresses struct {
	Private []CloudDatabaseAddressesDetail `json:"private"`
	Public  []CloudDatabaseAddressesDetail `json:"public"`
}

// CloudDatabaseDatastore contains datastore information.
type CloudDatabaseDatastore struct {
	ID          string `json:"id,omitempty"`
	Type        string `json:"type,omitempty"`
	VersionID   string `json:"version_id,omitempty"`
	VersionName string `json:"name,omitempty"`
}

// CloudDatabaseReplicasConfiguration is template to create secondary/ replica nodes
type CloudDatabaseReplicasConfiguration struct {
	AvailabilityZone string `json:"availability_zone"`
	Region           string `json:"region,omitempty"`
}

// CloudDatabaseReplicaNodeCreate is payload used to declare about secondary/ replica nodes
type CloudDatabaseReplicaNodeCreate struct {
	Configurations CloudDatabaseReplicasConfiguration `json:"configurations"`
	Quantity       int                                `json:"quantity"`
}

// CloudDatabaseVolume contains info of volume.
type CloudDatabaseVolume struct {
	Size int     `json:"size"`
	Used float32 `json:"used"`
}

// CloudDatabaseSuggestion contains structure of suggestion.
type CloudDatabaseSuggestion map[string]interface{}

// CloudDatabaseMessageResponse contains message response from Database Service API.
type CloudDatabaseMessageResponse struct {
	Message string `json:"message"`
	TaskID  string `json:"task_id"`
}

// CloudDatabaseDB - define for a database.
type CloudDatabaseDB struct {
	Name string `json:"name"`
}

// CloudDatabaseUser - define for an user.
type CloudDatabaseUser struct {
	Databases []CloudDatabaseDB `json:"databases"`
	Host      string            `json:"host"`
	Name      string            `json:"name"`
	Password  string            `json:"password"`
}

// Utils Function

// AddParamsListOption update params option when list resource.
func AddParamsListOption(req *http.Request, opts *CloudDatabaseListOption) {
	params := req.URL.Query()
	if opts.Page != 0 {
		params.Add("page", strconv.Itoa(opts.Page))
	}
	if opts.ResultsPerPage != 0 {
		params.Add("results_per_page", strconv.Itoa(opts.ResultsPerPage))
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
