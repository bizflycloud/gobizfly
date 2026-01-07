package gobizfly

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// AutoScalingWebhook is represents for a Webhook to trigger scale
type AutoScalingWebhook struct {
	ActionID    string `json:"action_id"`
	ActionType  string `json:"action_type"`
	ClusterID   string `json:"cluster_id"`
	ClusterName string `json:"cluster_name"`
}

// List of all available webhooks
func (wh *webhook) List(ctx context.Context, clusterID string) ([]*AutoScalingWebhook, error) {
	if clusterID == "" {
		return nil, errors.New("auto scaling group ID is required")
	}

	req, err := wh.client.NewRequest(ctx, http.MethodGet, autoScalingServiceName, wh.resourcePath(clusterID), nil)
	if err != nil {
		return nil, err
	}
	resp, err := wh.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	var data []*AutoScalingWebhook

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}

// Get information of a webhook
func (wh *webhook) Get(ctx context.Context, clusterID string, ActionType string) (*AutoScalingWebhook, error) {
	if clusterID == "" {
		return nil, errors.New("auto scaling group ID is required")
	}
	if _, ok := SliceContains(actionTypeSupportedWebhooks, ActionType); !ok {
		return nil, errors.New("UNSUPPORTED action type")
	}
	req, err := wh.client.NewRequest(ctx, http.MethodGet, autoScalingServiceName, wh.resourcePath(clusterID), nil)
	if err != nil {
		return nil, err
	}
	resp, err := wh.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	var data []*AutoScalingWebhook

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	for _, webhook := range data {
		if webhook.ActionType == ActionType {
			return webhook, nil
		}
	}
	return nil, nil
}
