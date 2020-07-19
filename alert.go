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
	"fmt"
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
	alertBasePath         = "/api/alert"
	alarmsResourcePath    = "alarms"
	receiversResourcePath = "receivers"
	historiesResourcePath = "histories"
	getVerificationPath   = "resend"
)

// Comparison is ...
type Comparison struct {
	CompareType string  `json:"compare_type"`
	Measurement string  `json:"measurement"`
	RangeTime   int     `json:"range_time"`
	Value       float64 `json:"value"`
}

// Instances is ...
type Instances struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Volumes is ...
type Volumes struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Device string `json:"device"`
}

// HTTPHeaders is ...
type HTTPHeaders struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// LoadBalancers is ...
type LoadBalancers struct {
	LoadBalancerID   string `json:"load_balancer_id"`
	LoadBalancerName string `json:"load_balancer_name"`
	TargetID         string `json:"target_id"`
	TargetName       string `json:"target_name"`
	TargetType       string `json:"target_type"`
}

// AlarmCreateReceivers is ...
type AlarmCreateReceivers struct {
	AutoscaleClusterName string `json:"autoscale_cluster_name,omitempty"`
	EmailAddress         string `json:"email_address,omitempty"`
	Name                 string `json:"name"`
	ReceiverID           string `json:"receiver_id"`
	SlackChannelName     string `json:"slack_channel_name,omitempty"`
	SMSInterval          string `json:"sms_interval,omitempty"`
	SMSNumber            string `json:"sms_number,omitempty"`
	TelegramChatID       string `json:"telegram_chat_id,omitempty"`
	WebhookURL           string `json:"webhook_url,omitempty"`
}

// AlarmGetReceivers is ...
type AlarmGetReceivers struct {
	AutoscaleClusterName string `json:"autoscale_cluster_name,omitempty"`
	EmailAddress         string `json:"email_address,omitempty"`
	Name                 string `json:"name"`
	ReceiverID           string `json:"_id"`
	SlackChannelName     string `json:"slack_channel_name,omitempty"`
	SMSInterval          string `json:"sms_interval,omitempty"`
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
	AlertInterval    int                    `json:"alert_interval"`
	ClusterID        string                 `json:"cluster_id,omitempty"`
	ClusterName      string                 `json:"cluster_name,omitempty"`
	Comparison       *Comparison            `json:"comparison,omitempty"`
	Hostname         string                 `json:"hostname,omitempty"`
	HTTPExpectedCode int                    `json:"http_expected_code,omitempty"`
	HTTPHeaders      *[]HTTPHeaders         `json:"http_headers,omitempty"`
	HTTPURL          string                 `json:"http_url,omitempty"`
	Instances        *[]Instances           `json:"instances,omitempty"`
	LoadBalancers    *[]LoadBalancers       `json:"load_balancers,omitempty"`
	Name             string                 `json:"name"`
	Receivers        []AlarmCreateReceivers `json:"receivers"`
	ResourceType     string                 `json:"resource_type"`
	Volumes          *[]Volumes             `json:"volumes,omitempty"`
}

// ReceiverCreateRequest contains receiver information.
type ReceiverCreateRequest struct {
	AutoScale      *AutoScale `json:"autoscale,omitempty"`
	EmailAddress   string     `json:"email_address,omitempty"`
	Name           string     `json:"name"`
	Slack          *Slack     `json:"slack,omitempty"`
	SMSInterval    string     `json:"sms_interval,omitempty"`
	SMSNumber      string     `json:"sms_number,omitempty"`
	TelegramChatID string     `json:"telegram_chat_id,omitempty"`
	WebhookURL     string     `json:"webhook_url,omitempty"`
}

// AlarmUpdateRequest represents update alarm request payload.
type AlarmUpdateRequest struct {
	AlertInterval    int                     `json:"alert_interval"`
	ClusterID        string                  `json:"cluster_id,omitempty"`
	ClusterName      string                  `json:"cluster_name,omitempty"`
	Comparison       *Comparison             `json:"comparison,omitempty"`
	Enable           bool                    `json:"enable"`
	Hostname         string                  `json:"hostname,omitempty"`
	HTTPExpectedCode int                     `json:"http_expected_code,omitempty"`
	HTTPHeaders      *[]HTTPHeaders          `json:"http_headers,omitempty"`
	HTTPURL          string                  `json:"http_url,omitempty"`
	Instances        *[]Instances            `json:"instances,omitempty"`
	LoadBalancers    *[]LoadBalancers        `json:"load_balancers,omitempty"`
	Name             string                  `json:"name"`
	Receivers        *[]AlarmCreateReceivers `json:"receivers,omitempty"`
	ResourceType     *string                 `json:"resource_type,omitempty"`
	Volumes          *[]Volumes              `json:"volumes,omitempty"`
}

// ReceiverUpdateRequest contains receiver information.
type ReceiverUpdateRequest struct {
	AutoScale      *AutoScale `json:"autoscale,omitempty"`
	EmailAddress   string     `json:"email_address,omitempty"`
	Name           string     `json:"name"`
	Slack          *Slack     `json:"slack,omitempty"`
	SMSInterval    string     `json:"sms_interval,omitempty"`
	SMSNumber      string     `json:"sms_number,omitempty"`
	TelegramChatID string     `json:"telegram_chat_id,omitempty"`
	WebhookURL     string     `json:"webhook_url,omitempty"`
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
	AlertInterval    int                 `json:"alert_interval"`
	ClusterID        string              `json:"cluster_id,omitempty"`
	ClusterName      string              `json:"cluster_name,omitempty"`
	Comparison       Comparison          `json:"comparison,omitempty"`
	Enable           bool                `json:"enable"`
	Hostname         string              `json:"hostname,omitempty"`
	HTTPExpectedCode int                 `json:"http_expected_code,omitempty"`
	HTTPHeaders      []HTTPHeaders       `json:"http_headers,omitempty"`
	HTTPURL          string              `json:"http_url,omitempty"`
	ID               string              `json:"_id"`
	Instances        []Instances         `json:"instances,omitempty"`
	LoadBalancers    []LoadBalancers     `json:"load_balancers,omitempty"`
	Name             string              `json:"name"`
	ProjectID        string              `json:"project_id"`
	Receivers        []AlarmGetReceivers `json:"receivers"`
	ResourceType     string              `json:"resource_type"`
	UserID           string              `json:"user_id"`
	Volumes          []Volumes           `json:"volumes,omitempty"`
}

// AlarmsInHistories contains alarm information.
type AlarmsInHistories struct {
	AlertInterval    int             `json:"alert_interval"`
	ClusterID        string          `json:"cluster_id,omitempty"`
	ClusterName      string          `json:"cluster_name,omitempty"`
	Comparison       Comparison      `json:"comparison,omitempty"`
	Enable           bool            `json:"enable"`
	Hostname         string          `json:"hostname,omitempty"`
	HTTPExpectedCode int             `json:"http_expected_code,omitempty"`
	HTTPHeaders      []HTTPHeaders   `json:"http_headers,omitempty"`
	HTTPURL          string          `json:"http_url,omitempty"`
	ID               string          `json:"_id"`
	Instances        []Instances     `json:"instances,omitempty"`
	LoadBalancers    []LoadBalancers `json:"load_balancers,omitempty"`
	Name             string          `json:"name"`
	ProjectID        string          `json:"project_id"`
	Receivers        []struct {
		ReceiverID string   `json:"receiver_id"`
		Methods    []string `json:"methods"`
	} `json:"receivers"`
	ResourceType string    `json:"resource_type"`
	UserID       string    `json:"user_id"`
	Volumes      []Volumes `json:"volumes,omitempty"`
}

// AutoScale is ...
type AutoScale struct {
	ClusterName string `json:"cluster_name"`
	ClusterID   string `json:"cluster_id"`
	ActionID    string `json:"action_id"`
	ActionType  string `json:"action_type"`
}

// Slack is ...
type Slack struct {
	SlackChannelName string `json:"channel_name"`
	WebhookURL       string `json:"webhook_url"`
}

// Receivers contains receiver information.
type Receivers struct {
	AutoScale              AutoScale `json:"autoscale,omitempty"`
	EmailAddress           string    `json:"email_address,omitempty"`
	Name                   string    `json:"name"`
	ReceiverID             string    `json:"_id"`
	Slack                  Slack     `json:"slack,omitempty"`
	SMSInterval            string    `json:"sms_interval,omitempty"`
	SMSNumber              string    `json:"sms_number,omitempty"`
	TelegramChatID         string    `json:"telegram_chat_id,omitempty"`
	WebhookURL             string    `json:"webhook_url,omitempty"`
	UserID                 string    `json:"user_id,omitempty"`
	ProjectID              string    `json:"project_id,omitempty"`
	VerifiedSMSNumber      bool      `json:"verified_sms_number,omitempty"`
	VerifiedEmailDddress   bool      `json:"verified_email_address,omitempty"`
	VerifiedWebhookURL     bool      `json:"verified_webhook_url,omitempty"`
	VerifiedTelegramChatID bool      `json:"verified_telegram_chat_id,omitempty"`
}

// Histories contains receiver information.
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
}

func (a *alarms) resourcePath(filters *string) string {
	if filters != nil {
		return fmt.Sprintf("%s?%s", strings.Join([]string{alertBasePath, alarmsResourcePath}, "/"), *filters)
	}
	return strings.Join([]string{alertBasePath, alarmsResourcePath}, "/")
}

func (a *alarms) itemPath(id string) string {
	return strings.Join([]string{alertBasePath, alarmsResourcePath, id}, "/")
}

func (r *receivers) resourcePath(filters *string) string {
	if filters != nil {
		return fmt.Sprintf("%s?%s", strings.Join([]string{alertBasePath, receiversResourcePath}, "/"), *filters)
	}
	return strings.Join([]string{alertBasePath, receiversResourcePath}, "/")
}

func (r *receivers) itemPath(id string) string {
	return strings.Join([]string{alertBasePath, receiversResourcePath, id}, "/")
}

func (r *receivers) verificationPath() string {
	return strings.Join([]string{alertBasePath, getVerificationPath}, "/")
}

func (h *histories) resourcePath(filters *string) string {
	if filters != nil {
		return fmt.Sprintf("%s?%s", strings.Join([]string{alertBasePath, historiesResourcePath}, "/"), *filters)
	}
	return strings.Join([]string{alertBasePath, historiesResourcePath}, "/")
}

// List is function using list alarms
func (a *alarms) List(ctx context.Context, filters *string) ([]*Alarms, error) {
	req, err := a.client.NewRequest(ctx, http.MethodGet, a.resourcePath(filters), nil)
	if err != nil {
		return nil, err
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
	req, err := r.client.NewRequest(ctx, http.MethodGet, r.resourcePath(filters), nil)
	if err != nil {
		return nil, err
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

// List is function using list histories'
func (h *histories) List(ctx context.Context, filters *string) ([]*Histories, error) {
	req, err := h.client.NewRequest(ctx, http.MethodGet, h.resourcePath(filters), nil)
	if err != nil {
		return nil, err
	}
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
	req, err := a.client.NewRequest(ctx, http.MethodPost, a.resourcePath(nil), &acr)
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
	req, err := r.client.NewRequest(ctx, http.MethodPost, r.resourcePath(nil), &rcr)
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
	req, err := a.client.NewRequest(ctx, http.MethodGet, a.itemPath(id), nil)
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
	return alarm, nil
}

func (r *receivers) Get(ctx context.Context, id string) (*Receivers, error) {
	req, err := r.client.NewRequest(ctx, http.MethodGet, r.itemPath(id), nil)
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

func (a *alarms) Update(ctx context.Context, id string, adr *AlarmUpdateRequest) (*ResponseRequest, error) {
	req, err := a.client.NewRequest(ctx, http.MethodPatch, a.itemPath(id), &adr)
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

func (r *receivers) Update(ctx context.Context, id string, rdr *ReceiverUpdateRequest) (*ResponseRequest, error) {
	req, err := r.client.NewRequest(ctx, http.MethodPut, r.itemPath(id), &rdr)
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
	req, err := a.client.NewRequest(ctx, http.MethodDelete, a.itemPath(id), nil)
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
	req, err := r.client.NewRequest(ctx, http.MethodDelete, r.itemPath(id), nil)
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

func (r *receivers) ResendVerificationLink(ctx context.Context, id string, rType string) error {
	req, err := r.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s?rec_id=%s&rec_type=%s", r.verificationPath(), id, rType), nil)
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
