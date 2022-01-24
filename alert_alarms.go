package gobizfly

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

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
	defer resp.Body.Close()

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
	defer resp.Body.Close()

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
	defer resp.Body.Close()

	alarm := &Alarms{}
	if err := json.NewDecoder(resp.Body).Decode(alarm); err != nil {
		return nil, err
	}
	// hardcode in here
	for _, loadbalancer := range alarm.LoadBalancers {
		if loadbalancer.TargetType == "frontend" {
			frontend, err := a.client.Listener.Get(ctx, loadbalancer.TargetID)
			if err != nil {
				loadbalancer.TargetName = ""
			}
			loadbalancer.TargetName = frontend.Name
		} else {
			backend, err := a.client.Pool.Get(ctx, loadbalancer.TargetID)
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
	defer resp.Body.Close()

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
	_, _ = io.Copy(ioutil.Discard, resp.Body)

	return resp.Body.Close()
}
