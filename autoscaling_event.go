package gobizfly

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

// AutoScalingEvent is represents for a event of auto scaling group
type AutoScalingEvent struct {
	ActionName string `json:"action"`
	ActionType string `json:"type"`
	Metadata   struct {
		Action struct {
			Outputs map[string]interface{} `json:"outputs"`
		} `json:"action"`
	} `json:"meta_data"`
	ClusterID    string `json:"cluster_id"`
	ID           string `json:"id"`
	Level        string `json:"level"`
	ObjectType   string `json:"otype"`
	StatusReason string `json:"status_reason"`
	Timestamp    string `json:"timestamp"`
}

// List autoscaling events of a cluster
func (e *event) List(ctx context.Context, clusterID string, page, total int) ([]*AutoScalingEvent, error) {
	if clusterID == "" {
		return nil, errors.New("auto scaling group ID is required")
	}

	req, err := e.client.NewRequest(ctx, http.MethodGet, autoScalingServiceName, e.resourcePath(clusterID, page, total), nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("cluster_id", clusterID)
	q.Add("page", strconv.Itoa(page))
	q.Add("total", strconv.Itoa(total))
	req.URL.RawQuery = q.Encode()

	resp, err := e.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	var data struct {
		AutoScalingEvents []*AutoScalingEvent `json:"events"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.AutoScalingEvents, nil
}
