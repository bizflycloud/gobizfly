// This file is part of gobizfly

package gobizfly

import (
	"context"
)

const (
	kafkaClusterPath = "/clusters"
)

var _ KafkaService = (*kafkaService)(nil)

type kafkaService struct {
	client *Client
}

// KafkaService is an interface to interact with Bizfly Cloud API.
type KafkaService interface {
	List(ctx context.Context, opts *KafkaClusterListOptions) ([]*KafkaCluster, error)
	Get(ctx context.Context, id string) (*ClusterResponse, error)
	Create(ctx context.Context, reqBody *KafkaInitClusterRequest) (*KafkaTaskResponse, error)
	Delete(ctx context.Context, id string) (*KafkaTaskResponse, error)
	ListFlavor(ctx context.Context, opts *KafkaFlavorListOptions) ([]*KafkaFlavorResponse, error)
	ListVersion(ctx context.Context, opts *KafkaVersionListOptions) ([]*KafkaVersionResponse, error)
	Resize(ctx context.Context, id string, reqBody *KafkaResizeClusterRequest) (*KafkaTaskResponse, error)
	AddNode(ctx context.Context, id string, reqBody *KafkaAddNodeRequest) (*KafkaTaskResponse, error)
}


