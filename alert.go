// This file is part of gobizfly
//
// Copyright (C) 2020  BizFly Cloud
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>

package gobizfly

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

var _ AlertService = (*alertService)(nil)

type alertService struct {
	client *Client
}

// AlertService is interface wrap others resource's interfaces
type AlertService interface {
	Alarms() *alarms
	Receivers() *receivers
	Histories() *histories
}

func (as *alertService) Alarms() *alarms {
	return &alarms{client: as.client}
}

func (as *alertService) Receivers() *receivers {
	return &receivers{client: as.client}
}

func (as *alertService) Histories() *histories {
	return &histories{client: as.client}
}

const (
	alarmsResourcePath    = "alarms"
	receiversResourcePath = "receivers"
	historiesResourcePath = "histories"
	getVerificationPath   = "resend"
)

// Comparison is represents comparison payload in alarms
type Comparison struct {
	CompareType string  `json:"compare_type"`
	Measurement string  `json:"measurement"`
	RangeTime   int     `json:"range_time"`
	Value       float64 `json:"value"`
}

// AlarmInstancesMonitors is represents instances payload - which servers will be monitored
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

// HTTPHeaders is is represents http headers - which using call to http_url
type HTTPHeaders struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// AlarmLoadBalancersMonitor is represents load balancer payload - which load balancer will be monitored
type AlarmLoadBalancersMonitor struct {
	LoadBalancerID   string `json:"load_balancer_id"`
	LoadBalancerName string `json:"load_balancer_name"`
	TargetID         string `json:"target_id"`
	TargetName       string `json:"target_name,omitempty"`
	TargetType       string `json:"target_type"`
}

// AlarmReceiversUse is represents receiver's payload will be use create alarms
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

type alarms struct {
	client *Client
}

type receivers struct {
	client *Client
}

type histories struct {
	client *Client
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

// ReceiverCreateRequest contains receiver information.
type ReceiverCreateRequest struct {
	AutoScale      *AutoScalingWebhook `json:"autoscale,omitempty"`
	EmailAddress   string              `json:"email_address,omitempty"`
	Name           string              `json:"name"`
	Slack          *Slack              `json:"slack,omitempty"`
	SMSNumber      string              `json:"sms_number,omitempty"`
	TelegramChatID string              `json:"telegram_chat_id,omitempty"`
	WebhookURL     string              `json:"webhook_url,omitempty"`
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

// Slack is represents slack payload - which will be use create a receiver
type Slack struct {
	SlackChannelName string `json:"channel_name"`
	WebhookURL       string `json:"webhook_url"`
}

// Receivers contains receiver information.
type Receivers struct {
	AutoScale              *AutoScalingWebhook `json:"autoscale,omitempty"`
	Created                string              `json:"_created"`
	Creator                string              `json:"creator"`
	EmailAddress           string              `json:"email_address,omitempty"`
	Name                   string              `json:"name"`
	ProjectID              string              `json:"project_id,omitempty"`
	ReceiverID             string              `json:"_id"`
	Slack                  *Slack              `json:"slack,omitempty"`
	SMSNumber              string              `json:"sms_number,omitempty"`
	TelegramChatID         string              `json:"telegram_chat_id,omitempty"`
	UserID                 string              `json:"user_id,omitempty"`
	VerifiedEmailDddress   bool                `json:"verified_email_address,omitempty"`
	VerifiedSMSNumber      bool                `json:"verified_sms_number,omitempty"`
	VerifiedTelegramChatID bool                `json:"verified_telegram_chat_id,omitempty"`
	VerifiedWebhookURL     bool                `json:"verified_webhook_url,omitempty"`
	WebhookURL             string              `json:"webhook_url,omitempty"`
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

func (a *alarms) resourcePath() string {
	return strings.Join([]string{alarmsResourcePath}, "/")
}

func (a *alarms) itemPath(id string) string {
	return strings.Join([]string{alarmsResourcePath, id}, "/")
}

func (r *receivers) resourcePath() string {
	return strings.Join([]string{receiversResourcePath}, "/")
}

func (r *receivers) itemPath(id string) string {
	return strings.Join([]string{receiversResourcePath, id}, "/")
}

func (r *receivers) verificationPath() string {
	return strings.Join([]string{getVerificationPath}, "/")
}

func (h *histories) resourcePath() string {
	return strings.Join([]string{historiesResourcePath}, "/")
}

// List is function using list alarms
func (a *alarms) List(ctx context.Context, filters *string) ([]*Alarms, error) {
	req, err := a.client.NewRequest(ctx, http.MethodGet, alertServiceName, a.resourcePath(), nil)
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

// List is function using list receivers
func (r *receivers) List(ctx context.Context, filters *string) ([]*Receivers, error) {
	req, err := r.client.NewRequest(ctx, http.MethodGet, alertServiceName, r.resourcePath(), nil)
	if err != nil {
		return nil, err
	}

	if filters != nil {
		q := req.URL.Query()
		q.Add("where", *filters)
		req.URL.RawQuery = q.Encode()
	}

	resp, err := r.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		Receivers []*Receivers `json:"_items"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Receivers, nil
}

// List is function using list histories
func (h *histories) List(ctx context.Context, filters *string) ([]*Histories, error) {
	req, err := h.client.NewRequest(ctx, http.MethodGet, alertServiceName, h.resourcePath(), nil)
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
	defer resp.Body.Close()

	var data struct {
		Histories []*Histories `json:"_items"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Histories, nil
}

func (a *alarms) Create(ctx context.Context, acr *AlarmCreateRequest) (*ResponseRequest, error) {
	req, err := a.client.NewRequest(ctx, http.MethodPost, alertServiceName, a.resourcePath(), &acr)
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

func (r *receivers) Create(ctx context.Context, rcr *ReceiverCreateRequest) (*ResponseRequest, error) {
	req, err := r.client.NewRequest(ctx, http.MethodPost, alertServiceName, r.resourcePath(), &rcr)
	if err != nil {
		return nil, err
	}
	resp, err := r.client.Do(ctx, req)
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

func (a *alarms) Get(ctx context.Context, id string) (*Alarms, error) {
	req, err := a.client.NewRequest(ctx, http.MethodGet, alertServiceName, a.itemPath(id), nil)
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

func (r *receivers) Get(ctx context.Context, id string) (*Receivers, error) {
	req, err := r.client.NewRequest(ctx, http.MethodGet, alertServiceName, r.itemPath(id), nil)
	if err != nil {
		return nil, err
	}
	resp, err := r.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	receiver := &Receivers{}
	if err := json.NewDecoder(resp.Body).Decode(receiver); err != nil {
		return nil, err
	}
	return receiver, nil
}

func (a *alarms) Update(ctx context.Context, id string, aur *AlarmUpdateRequest) (*ResponseRequest, error) {
	req, err := a.client.NewRequest(ctx, http.MethodPatch, alertServiceName, a.itemPath(id), &aur)
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

func (r *receivers) Update(ctx context.Context, id string, rur *ReceiverCreateRequest) (*ResponseRequest, error) {
	req, err := r.client.NewRequest(ctx, http.MethodPut, alertServiceName, r.itemPath(id), &rur)
	if err != nil {
		return nil, err
	}
	resp, err := r.client.Do(ctx, req)
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

func (a *alarms) Delete(ctx context.Context, id string) error {
	req, err := a.client.NewRequest(ctx, http.MethodDelete, alertServiceName, a.itemPath(id), nil)
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

func (r *receivers) Delete(ctx context.Context, id string) error {
	req, err := r.client.NewRequest(ctx, http.MethodDelete, alertServiceName, r.itemPath(id), nil)
	if err != nil {
		return err
	}
	resp, err := r.client.Do(ctx, req)
	if err != nil {
		return err
	}
	_, _ = io.Copy(ioutil.Discard, resp.Body)

	return resp.Body.Close()
}

// ResendVerificationLink is use get a link verification
func (r *receivers) ResendVerificationLink(ctx context.Context, id string, rType string) error {
	req, err := r.client.NewRequest(ctx, http.MethodGet, alertServiceName, r.verificationPath(), nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("rec_id", id)
	q.Add("rec_type", rType)
	req.URL.RawQuery = q.Encode()

	resp, err := r.client.Do(ctx, req)
	if err != nil {
		return err
	}
	_, _ = io.Copy(ioutil.Discard, resp.Body)

	return resp.Body.Close()
}
