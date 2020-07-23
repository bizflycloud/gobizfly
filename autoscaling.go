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
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const (
	autoscalingServiceBasePath       = "/auto-scaling/api"
	launchConfigurationsResourcePath = "launch_configs"
	autoscalingResourcePath          = "groups"
	webhooksResourcePath             = "webhooks"
)

var (
	actionTypeSupportedWebhooks = []string{
		"CLUSTER SCALE IN",
		"CLUSTER SCALE OUT",
	}
)

var _ AutoScalingService = (*autoscalingService)(nil)

type autoscalingService struct {
	client *Client
}

// AutoScalingService is interface wrap others resource's interfaces
type AutoScalingService interface {
	AutoScalingGroups() *autoScalingGroup
	LaunchConfigurations() *launchConfiguration
	// Nodes() *histories
	// Policies() *histories
	// Events() *histories
	Webhooks() *webhook
	// Schedules() *histories
}

func (as *autoscalingService) AutoScalingGroups() *autoScalingGroup {
	return &autoScalingGroup{client: as.client}
}

func (as *autoscalingService) LaunchConfigurations() *launchConfiguration {
	return &launchConfiguration{client: as.client}
}

func (as *autoscalingService) Webhooks() *webhook {
	return &webhook{client: as.client}
}

type autoScalingGroup struct {
	client *Client
}

type launchConfiguration struct {
	client *Client
}

type webhook struct {
	client *Client
}

// AutoScalingGroup - is represents an auto scaling group
type AutoScalingGroup struct {
}

// AutoScalingDataDisk is represents for a data disk being created with servers
type AutoScalingDataDisk struct {
	Type                string `json:"type"`
	Size                int    `json:"size"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
}

// AutoScalingOperatingSystem is represents for operating system being use to create servers
type AutoScalingOperatingSystem struct {
	CreateFrom string `json:"type,omitempty"`
	ID         string `json:"id,omitempty"`
	OSName     string `json:"os_name,omitempty"`
	Error      string `json:"error,omitempty"`
}

// LaunchConfiguration - is represents a launch configurations
type LaunchConfiguration struct {
	AvailabilityZone string                     `json:"availability_zone"`
	Created          string                     `json:"created_at"`
	DataDisks        []AutoScalingDataDisk      `json:"datadisks"`
	Flavor           string                     `json:"flavor"`
	ID               string                     `json:"id"`
	Name             string                     `json:"name"`
	OperatingSystem  AutoScalingOperatingSystem `json:"os"`
	RootDisk         AutoScalingDataDisk        `json:"rootdisk"`
	SecurityGroups   []string                   `json:"security_groups"`
	SSHKey           string                     `json:"key_name"`
	Status           string
	Type             string `json:"type"`
	UserData         string `json:"user_data"`
}

// AutoScalingWebhook is represents for a Webhook to trigger scale
type AutoScalingWebhook struct {
	ActionID    string `json:"action_id"`
	ActionType  string `json:"action_type"`
	ClusterID   string `json:"cluster_id"`
	ClusterName string `json:"cluster_name"`
}

// Launch Configurations path
func (lc *launchConfiguration) resourcePath(all bool) string {
	if all {
		return fmt.Sprintf("%s?all", strings.Join([]string{autoscalingServiceBasePath, launchConfigurationsResourcePath}, "/"))
	}
	return strings.Join([]string{autoscalingServiceBasePath, launchConfigurationsResourcePath}, "/")
}

func (lc *launchConfiguration) itemPath(id string) string {
	return strings.Join([]string{autoscalingServiceBasePath, launchConfigurationsResourcePath, id}, "/")
}

// Webhook path
func (wh *webhook) resourcePath(clusterID string) string {
	return strings.Join([]string{autoscalingServiceBasePath, autoscalingResourcePath, clusterID, webhooksResourcePath}, "/")
}

// List
func (lc *launchConfiguration) List(ctx context.Context, all bool) ([]*LaunchConfiguration, error) {
	req, err := lc.client.NewRequest(ctx, http.MethodGet, lc.resourcePath(all), nil)
	if err != nil {
		return nil, err
	}
	resp, err := lc.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		LaunchConfigurations []*LaunchConfiguration `json:"profiles"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	for _, LaunchConfiguration := range data.LaunchConfigurations {
		if LaunchConfiguration.OperatingSystem.Error != "" {
			LaunchConfiguration.Status = "Error"
		} else {
			LaunchConfiguration.Status = "Active"
		}
	}
	return data.LaunchConfigurations, nil
}

func (wh *webhook) List(ctx context.Context, clusterID string) ([]*AutoScalingWebhook, error) {
	if clusterID == "" {
		return nil, errors.New("Auto Scaling Group ID is required")
	}

	req, err := wh.client.NewRequest(ctx, http.MethodGet, wh.resourcePath(clusterID), nil)
	if err != nil {
		return nil, err
	}
	resp, err := wh.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data []*AutoScalingWebhook

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}

// Get
func (lc *launchConfiguration) Get(ctx context.Context, profileID string) (*LaunchConfiguration, error) {
	req, err := lc.client.NewRequest(ctx, http.MethodGet, lc.itemPath(profileID), nil)
	if err != nil {
		return nil, err
	}
	resp, err := lc.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data := &LaunchConfiguration{}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	if data.OperatingSystem.Error != "" {
		data.Status = "Error"
	} else {
		data.Status = "Active"
	}
	return data, nil
}

func (wh *webhook) Get(ctx context.Context, clusterID string, ActionType string) (*AutoScalingWebhook, error) {
	if clusterID == "" {
		return nil, errors.New("Auto Scaling Group ID is required")
	}
	if _, ok := SliceContains(actionTypeSupportedWebhooks, ActionType); !ok {
		return nil, errors.New("UNSUPPORTED action type")
	}
	req, err := wh.client.NewRequest(ctx, http.MethodGet, wh.resourcePath(clusterID), nil)
	if err != nil {
		return nil, err
	}
	resp, err := wh.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

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
