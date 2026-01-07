package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

// AlarmsInHistories contains alarm information in a history
type AlarmsInHistories struct {
	AlertInterval    int                          `json:"alert_interval"`
	ClusterID        string                       `json:"cluster_id,omitempty"`
	ClusterName      string                       `json:"cluster_name,omitempty"`
	Comparison       *Comparison                  `json:"comparison,omitempty"`
	Enable           bool                         `json:"enable"`
	Hostname         string                       `json:"hostname,omitempty"`
	HTTPExpectedCode int                          `json:"http_expected_code,omitempty"`
	HTTPHeaders      *[]HTTPHeaders               `json:"http_headers,omitempty"`
	HTTPURL          string                       `json:"http_url,omitempty"`
	ID               string                       `json:"_id"`
	Instances        *[]AlarmInstancesMonitors    `json:"instances,omitempty"`
	LoadBalancers    *[]AlarmLoadBalancersMonitor `json:"load_balancers,omitempty"`
	Name             string                       `json:"name"`
	ProjectID        string                       `json:"project_id"`
	Receivers        *[]struct {
		ReceiverID string   `json:"receiver_id"`
		Methods    []string `json:"methods"`
	} `json:"receivers"`
	ResourceType string                 `json:"resource_type"`
	UserID       string                 `json:"user_id"`
	Volumes      *[]AlarmVolumesMonitor `json:"volumes,omitempty"`
}

// Histories contains history information.
type Histories struct {
	HistoryID   string            `json:"_id"`
	Name        string            `json:"name"`
	ProjectID   string            `json:"project_id,omitempty"`
	UserID      string            `json:"user_id,omitempty"`
	Resource    interface{}       `json:"resource,omitempty"`
	State       string            `json:"state,omitempty"`
	Measurement string            `json:"measurement,omitempty"`
	AlarmID     string            `json:"alarm_id"`
	Alarm       AlarmsInHistories `json:"alarm,omitempty"`
	Created     string            `json:"_created,omitempty"`
}

func (h *histories) resourcePath() string {
	return strings.Join([]string{historiesResourcePath}, "/")
}

// List histories
func (h *histories) List(ctx context.Context, filters *string) ([]*Histories, error) {
	req, err := h.client.NewRequest(ctx, http.MethodGet, cloudwatcherServiceName, h.resourcePath(), nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	if filters != nil {
		q.Add("where", *filters)
	}
	q.Add("sort", "-_created")
	req.URL.RawQuery = q.Encode()

	resp, err := h.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	var data struct {
		Histories []*Histories `json:"_items"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Histories, nil
}
