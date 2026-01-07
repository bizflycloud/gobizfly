package gobizfly

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// Comparison represents comparison payload in alarms
type Comparison struct {
	CompareType string  `json:"compare_type"`
	Measurement string  `json:"measurement"`
	RangeTime   int     `json:"range_time"`
	Value       float64 `json:"value"`
}

// AlarmInstancesMonitors represents instances payload - which servers will be monitored
type AlarmInstancesMonitors struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// AlarmVolumesMonitor is represents volumes payload - which volumes will be monitored
type AlarmVolumesMonitor struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Device string `json:"device,omitempty"`
}

// HTTPHeaders represents http headers - which using call to http_url
type HTTPHeaders struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// AlarmLoadBalancersMonitor represents load balancer payload - which load balancer will be monitored
type AlarmLoadBalancersMonitor struct {
	LoadBalancerID   string `json:"load_balancer_id"`
	LoadBalancerName string `json:"load_balancer_name"`
	TargetID         string `json:"target_id"`
	TargetName       string `json:"target_name,omitempty"`
	TargetType       string `json:"target_type"`
}

// AlarmReceiversUse represents receiver's payload will be use create alarms
type AlarmReceiversUse struct {
	AutoscaleClusterName string `json:"autoscale_cluster_name,omitempty"`
	EmailAddress         string `json:"email_address,omitempty"`
	Name                 string `json:"name"`
	ReceiverID           string `json:"receiver_id"`
	SlackChannelName     string `json:"slack_channel_name,omitempty"`
	SMSInterval          int    `json:"sms_interval,omitempty"`
	SMSNumber            string `json:"sms_number,omitempty"`
	TelegramChatID       string `json:"telegram_chat_id,omitempty"`
	WebhookURL           string `json:"webhook_url,omitempty"`
}

// AlarmCreateRequest represents create new alarm request payload.
type AlarmCreateRequest struct {
	AlertInterval    int                          `json:"alert_interval"`
	ClusterID        string                       `json:"cluster_id,omitempty"`
	ClusterName      string                       `json:"cluster_name,omitempty"`
	Comparison       *Comparison                  `json:"comparison,omitempty"`
	Hostname         string                       `json:"hostname,omitempty"`
	HTTPExpectedCode int                          `json:"http_expected_code,omitempty"`
	HTTPHeaders      *[]HTTPHeaders               `json:"http_headers,omitempty"`
	HTTPURL          string                       `json:"http_url,omitempty"`
	Instances        *[]AlarmInstancesMonitors    `json:"instances,omitempty"`
	LoadBalancers    []*AlarmLoadBalancersMonitor `json:"load_balancers,omitempty"`
	Name             string                       `json:"name"`
	Receivers        []AlarmReceiversUse          `json:"receivers"`
	ResourceType     string                       `json:"resource_type"`
	Volumes          *[]AlarmVolumesMonitor       `json:"volumes,omitempty"`
}

// AlarmUpdateRequest represents update alarm request payload.
type AlarmUpdateRequest struct {
	AlertInterval    int                          `json:"alert_interval,omitempty"`
	ClusterID        string                       `json:"cluster_id,omitempty"`
	ClusterName      string                       `json:"cluster_name,omitempty"`
	Comparison       *Comparison                  `json:"comparison,omitempty"`
	Enable           bool                         `json:"enable"`
	Hostname         string                       `json:"hostname,omitempty"`
	HTTPExpectedCode int                          `json:"http_expected_code,omitempty"`
	HTTPHeaders      *[]HTTPHeaders               `json:"http_headers,omitempty"`
	HTTPURL          string                       `json:"http_url,omitempty"`
	Instances        *[]AlarmInstancesMonitors    `json:"instances,omitempty"`
	LoadBalancers    []*AlarmLoadBalancersMonitor `json:"load_balancers,omitempty"`
	Name             string                       `json:"name,omitempty"`
	Receivers        *[]AlarmReceiversUse         `json:"receivers,omitempty"`
	ResourceType     string                       `json:"resource_type,omitempty"`
	Volumes          *[]AlarmVolumesMonitor       `json:"volumes,omitempty"`
}

// ResponseRequest represents api's response.
type ResponseRequest struct {
	Created string `json:"_created,omitempty"`
	Deleted bool   `json:"_deleted,omitempty"`
	ID      string `json:"_id,omitempty"`
	Status  string `json:"_status,omitempty"`
}

// Alarms contains alarm information.
type Alarms struct {
	AlertInterval    int                          `json:"alert_interval"`
	ClusterID        string                       `json:"cluster_id,omitempty"`
	ClusterName      string                       `json:"cluster_name,omitempty"`
	Comparison       *Comparison                  `json:"comparison,omitempty"`
	Created          string                       `json:"_created"`
	Creator          string                       `json:"creator"`
	Enable           bool                         `json:"enable"`
	Hostname         string                       `json:"hostname,omitempty"`
	HTTPExpectedCode int                          `json:"http_expected_code,omitempty"`
	HTTPHeaders      []HTTPHeaders                `json:"http_headers,omitempty"`
	HTTPURL          string                       `json:"http_url,omitempty"`
	ID               string                       `json:"_id"`
	Instances        []AlarmInstancesMonitors     `json:"instances,omitempty"`
	LoadBalancers    []*AlarmLoadBalancersMonitor `json:"load_balancers,omitempty"`
	Name             string                       `json:"name"`
	ProjectID        string                       `json:"project_id"`
	Receivers        []AlarmReceiversUse          `json:"receivers"`
	ResourceType     string                       `json:"resource_type"`
	UserID           string                       `json:"user_id"`
	Volumes          []AlarmVolumesMonitor        `json:"volumes,omitempty"`
}

func (a *alarms) resourcePath() string {
	return strings.Join([]string{alarmsResourcePath}, "/")
}

func (a *alarms) itemPath(id string) string {
	return strings.Join([]string{alarmsResourcePath, id}, "/")
}

// List alarms
func (a *alarms) List(ctx context.Context, filters *string) ([]*Alarms, error) {
	req, err := a.client.NewRequest(ctx, http.MethodGet, cloudwatcherServiceName, a.resourcePath(), nil)
	if err != nil {
		return nil, err
	}

	if filters != nil {
		q := req.URL.Query()
		q.Add("where", *filters)
		req.URL.RawQuery = q.Encode()
	}

	resp, err := a.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	var data struct {
		Alarms []*Alarms `json:"_items"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Alarms, nil
}

// Create an alarm
func (a *alarms) Create(ctx context.Context, acr *AlarmCreateRequest) (*ResponseRequest, error) {
	req, err := a.client.NewRequest(ctx, http.MethodPost, cloudwatcherServiceName, a.resourcePath(), &acr)
	if err != nil {
		return nil, err
	}
	resp, err := a.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	var respData = &ResponseRequest{}
	if err := json.NewDecoder(resp.Body).Decode(respData); err != nil {
		return nil, err
	}
	return respData, nil
}

// Get an alarm
func (a *alarms) Get(ctx context.Context, id string) (*Alarms, error) {
	req, err := a.client.NewRequest(ctx, http.MethodGet, cloudwatcherServiceName, a.itemPath(id), nil)
	if err != nil {
		return nil, err
	}
	resp, err := a.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	alarm := &Alarms{}
	if err := json.NewDecoder(resp.Body).Decode(alarm); err != nil {
		return nil, err
	}
	// hardcode in here
	for _, loadbalancer := range alarm.LoadBalancers {
		if loadbalancer.TargetType == "frontend" {
			frontend, err := a.client.CloudLoadBalancer.Listeners().Get(ctx, loadbalancer.TargetID)
			if err != nil {
				loadbalancer.TargetName = ""
			}
			loadbalancer.TargetName = frontend.Name
		} else {
			backend, err := a.client.CloudLoadBalancer.Pools().Get(ctx, loadbalancer.TargetID)
			if err != nil {
				loadbalancer.TargetName = ""
			}
			loadbalancer.TargetName = backend.Name
		}
	}
	return alarm, nil
}

// Update an alarm
func (a *alarms) Update(ctx context.Context, id string, aur *AlarmUpdateRequest) (*ResponseRequest, error) {
	req, err := a.client.NewRequest(ctx, http.MethodPatch, cloudwatcherServiceName, a.itemPath(id), &aur)
	if err != nil {
		return nil, err
	}
	resp, err := a.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	respData := &ResponseRequest{}
	if err := json.NewDecoder(resp.Body).Decode(respData); err != nil {
		return nil, err
	}

	return respData, nil
}

// Delete an alarm
func (a *alarms) Delete(ctx context.Context, id string) error {
	req, err := a.client.NewRequest(ctx, http.MethodDelete, cloudwatcherServiceName, a.itemPath(id), nil)
	if err != nil {
		return err
	}
	resp, err := a.client.Do(ctx, req)
	if err != nil {
		return err
	}
	_, _ = io.Copy(io.Discard, resp.Body)

	return resp.Body.Close()
}
