package gobizfly

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

// KafkaClusterListOptions represents the options for listing Kafka clusters.
// Name is the filter of cluster name.
type KafkaClusterListOptions struct {
	Name string `url:"name,omitempty"`
}

type MetaData struct {
	HasNext      bool `json:"has_next" example:"false"`
	CurrentPage  int  `json:"current_page" example:"1"`
	HasPrevious  int  `json:"has_previous" example:"0"`
	Total        int  `json:"total" example:"1"`
	PreviousPage int  `json:"previous_page" example:"0"`
	NextPage     int  `json:"next_page" example:"0"`
	Pages        int  `json:"page" example:"1"`
}

type OBS struct {
	DashboardURL string `json:"dashboard_url" bson:"dashboard_url,omitempty" validate:"required"`
	OBSID        string `json:"obs_id" bson:"obs_id,omitempty" validate:"required"`
	QueryLog     bool   `json:"query_log" bson:"query_log,omitempty" validate:"required"`
}

// KafkaCluster contains Kafka cluster information.
type KafkaCluster struct {
	ID               string `json:"id" bson:"id,omitempty"`
	Name             string `json:"name" bson:"name,omitempty"`
	KafkaVersion     string `json:"version_id" bson:"version_id,omitempty"`
	Nodes            int    `json:"nodes" bson:"nodes,omitempty"`
	Flavor           string `json:"flavor" bson:"flavor,omitempty"`
	VolumeSize       int    `json:"volume_size" bson:"volume_size,omitempty"`
	Status           string `json:"status" bson:"status,omitempty"`
	CreatedAt        string `json:"created_at" bson:"created_at,omitempty"`
	AvailabilityZone string `json:"availability_zone" bson:"availability_zone,omitempty"`
	VPCNetworkID     string `json:"vpc_network_id" bson:"vpc_network_id"`
	PublicAccess     bool   `json:"public_access" bson:"public_access,omitempty"`
	TaskID           string `json:"task_id,omitempty" bson:"task_id,omitempty"`
	OBS              OBS    `json:"obs,omitempty" bson:"obs,omitempty"`
	ProjectID        string `json:"project_id" bson:"project_id,omitempty"`
}
type NodeResponse struct {
	ID               string          `json:"id" bson:"id,omitempty"`
	ClusterID        string          `json:"cluster_id" bson:"cluster_id,omitempty"`
	Name             string          `json:"name" bson:"name,omitempty"`
	Flavor           string          `json:"flavor" bson:"flavor,omitempty"`
	Address          AddressResponse `json:"address" bson:"addresses,omitempty"`
	VolumeSize       int             `json:"volume_size" bson:"volume_size,omitempty"`
	Used             float64         `json:"used" bson:"used,omitempty"`
	Status           string          `json:"status" bson:"status,omitempty"`
	CreatedAt        string          `json:"created_at" bson:"created_at,omitempty"`
	AvailabilityZone string          `json:"availability_zone" bson:"availability_zone,omitempty"`
}

type AddressResponse struct {
	LAN    []KIP `json:"lan" bson:"lan,omitempty"`
	WAN_V4 []KIP `json:"wan_v4" bson:"wan_v4,omitempty"`
	WAN_V6 []KIP `json:"wan_v6" bson:"wan_v6,omitempty"`
}

type KIP struct {
	IPAddress   string `json:"ip_address" bson:"ip_address,omitempty"`
	NetworkName string `json:"network_name" bson:"network_name,omitempty"`
	Port        int    `json:"port" bson:"port,omitempty"`
}

type ClusterResponse struct {
	ID               string         `json:"id" bson:"id,omitempty"`
	Name             string         `json:"name" bson:"name,omitempty"`
	KafkaVersion     string         `json:"version_id" bson:"version_id,omitempty"`
	Nodes            []NodeResponse `json:"nodes" bson:"nodes,omitempty"`
	Flavor           string         `json:"flavor" bson:"flavor,omitempty"`
	VolumeSize       int            `json:"volume_size" bson:"volume_size,omitempty"`
	Status           string         `json:"status" bson:"status,omitempty"`
	CreatedAt        string         `json:"created_at" bson:"created_at,omitempty"`
	AvailabilityZone string         `json:"availability_zone" bson:"availability_zone,omitempty"`
	PublicAccess     bool           `json:"public_access" bson:"public_access,omitempty"`
	NumTopics        int            `json:"num_topics" bson:"num_topics,omitempty"`
	NumPartitions    int            `json:"num_partitions" bson:"num_partitions,omitempty"`
	TaskID           string         `json:"task_id,omitempty" bson:"task_id,omitempty"`
	OBS              OBS            `json:"obs,omitempty" bson:"obs,omitempty"`
	ProjectID        string         `json:"project_id" bson:"project_id,omitempty"`
}

type ListClusterResponse struct {
	Success   bool            `json:"success" example:"true"`
	ErrorCode int             `json:"error_code" example:"0"`
	Message   string          `json:"message" example:"success"`
	Data      []*KafkaCluster `json:"data"`
	MetaData  MetaData        `json:"metadata"`
}

type GetClusterResponse struct {
	Success   bool             `json:"success" example:"true"`
	ErrorCode int              `json:"error_code" example:"0"`
	Message   string           `json:"message" example:"success"`
	Data      *ClusterResponse `json:"data"`
}

type KafkaTaskResponse struct {
	TaskID string `json:"data"`
}

type KafkaInitClusterRequest struct {
	ClusterName      string `json:"name"`
	VersionID        string `json:"version_id"`
	Nodes            int    `json:"nodes"`
	Flavor           string `json:"flavor"`
	PublicAccess     bool   `json:"public_access"`
	VolumeSize       int    `json:"volume_size"`
	VPCNetworkID     string `json:"vpc_network_id"`
	AvailabilityZone string `json:"availability_zone"`
}

// Resize cluster
type KafkaResizeClusterRequest struct {
	Flavor string `json:"flavor,omitempty"`

	// VolumeSize is the new volume size in GB
	VolumeSize int    `json:"volume_size,omitempty"`

	// Type can be "flavor" or "volume"
	// If Type is "flavor", Flavor field must be set
	// If Type is "volume", VolumeSize field must be set
	Type       string `json:"type" validate:"required"`
}

type KafkaAddNodeRequest struct {
	// NumNodes is the number of nodes to add
	Nodes int `json:"num_nodes"`

	// Type is the type of resize operation, e.g., "increase" or "decrease".
	// Now only "increase" is supported.
	Type string `json:"type" validate:"required"`
}

// List all Kafka clusters.
func (s *kafkaService) List(ctx context.Context, opts *KafkaClusterListOptions) ([]*KafkaCluster, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, kafkaServiceName, kafkaClusterPath, nil)
	if err != nil {
		return nil, err
	}
	params := req.URL.Query()
	if opts != nil {
		if opts.Name != "" {
			params.Add("name", opts.Name)
		}
	}
	req.URL.RawQuery = params.Encode()

	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var clustersRes *ListClusterResponse

	if err := json.NewDecoder(resp.Body).Decode(&clustersRes); err != nil {
		return nil, err
	}

	return clustersRes.Data, nil
}

// Get a Kafka cluster by ID.
func (s *kafkaService) Get(ctx context.Context, id string) (*ClusterResponse, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, kafkaServiceName, kafkaClusterPath+"/"+id, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var clusterRes GetClusterResponse

	if err := json.NewDecoder(resp.Body).Decode(&clusterRes); err != nil {
		return nil, err
	}

	return clusterRes.Data, nil
}

// Delete a Kafka cluster by ID.
func (s *kafkaService) Delete(ctx context.Context, id string) (*KafkaTaskResponse, error) {
	req, err := s.client.NewRequest(ctx, http.MethodDelete, kafkaServiceName, kafkaClusterPath+"/"+id, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var deleteRes KafkaTaskResponse

	if err := json.NewDecoder(resp.Body).Decode(&deleteRes); err != nil {
		return nil, err
	}

	return &deleteRes, nil
}

// Create a new Kafka cluster.
func (s *kafkaService) Create(ctx context.Context, reqBody *KafkaInitClusterRequest) (*KafkaTaskResponse, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, kafkaServiceName, kafkaClusterPath, reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var createRes KafkaTaskResponse

	if err := json.NewDecoder(resp.Body).Decode(&createRes); err != nil {
		return nil, err
	}

	return &createRes, nil
}

// Resize a Kafka cluster.
func (s *kafkaService) Resize(ctx context.Context, id string, reqBody *KafkaResizeClusterRequest) (*KafkaTaskResponse, error) {
	// validate the request body
	switch reqBody.Type {
	case "flavor":
		if reqBody.Flavor == "" {
			return nil, errors.New("Missing required field: Flavor for type 'flavor'")
		}
	case "volume":
		if reqBody.VolumeSize < 10 {
			return nil, errors.New("VolumeSize must be at least 10 GB")
		}
	default:
		return nil, errors.New("Invalid Type field: must be 'flavor' or 'volume'")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPut, kafkaServiceName, strings.Join([]string{kafkaClusterPath, id}, "/"), reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var resizeResp KafkaTaskResponse

	if err := json.NewDecoder(resp.Body).Decode(&resizeResp); err != nil {
		return nil, err
	}

	return &resizeResp, nil
}

// Add node to Kafka cluster
func (s *kafkaService) AddNode(ctx context.Context, id string, reqBody *KafkaAddNodeRequest) (*KafkaTaskResponse, error) {
	// validate the request body
	if reqBody.Nodes <= 0 {
		return nil, errors.New("Number of nodes to add must be greater than 0")
	}
	if reqBody.Type != "increase" {
		return nil, errors.New("Invalid Type field: only 'increase' is supported")
	}
	req, err := s.client.NewRequest(ctx, http.MethodPost, kafkaServiceName, strings.Join([]string{kafkaClusterPath, id, "resize"}, "/"), reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var addNodeResp KafkaTaskResponse

	if err := json.NewDecoder(resp.Body).Decode(&addNodeResp); err != nil {
		return nil, err
	}

	return &addNodeResp, nil
}
