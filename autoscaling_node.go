package gobizfly

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

// AutoScalingNode is represents a cloud server in auto scaling group
type AutoScalingNode struct {
	Status       string                 `json:"status"`
	Name         string                 `json:"name"`
	ProfileID    string                 `json:"profile_id"`
	ProfileName  string                 `json:"profile_name"`
	PhysicalID   string                 `json:"physical_id"`
	StatusReason string                 `json:"status_reason"`
	ID           string                 `json:"id"`
	Addresses    map[string]interface{} `json:"addresses,omitempty"`
}

// AutoScalingNodesDelete is represents a list cloud server being deleted
type AutoScalingNodesDelete struct {
	Nodes []string `json:"nodes"`
}

// List nodes of the given cluster
func (n *node) List(ctx context.Context, clusterID string, all bool) ([]*AutoScalingNode, error) {
	if clusterID == "" {
		return nil, errors.New("Auto Scaling Group ID is required")
	}

	req, err := n.client.NewRequest(ctx, http.MethodGet, autoScalingServiceName, n.resourcePath(clusterID), nil)
	if err != nil {
		return nil, err
	}

	if all {
		q := req.URL.Query()
		q.Add("all", "true")
		req.URL.RawQuery = q.Encode()
	}

	resp, err := n.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		AutoScalingNodes []*AutoScalingNode `json:"nodes"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.AutoScalingNodes, nil
}

// Delete nodes of the given cluster
func (n *node) Delete(ctx context.Context, clusterID string, asnd *AutoScalingNodesDelete) error {
	req, err := n.client.NewRequest(ctx, http.MethodDelete, autoScalingServiceName, n.resourcePath(clusterID), &asnd)
	if err != nil {
		return err
	}

	resp, err := n.client.Do(ctx, req)
	if err != nil {
		return err
	}
	_, _ = io.Copy(ioutil.Discard, resp.Body)

	return resp.Body.Close()
}
