// This file is part of gobizfly

package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

const (
	statusActive          = "Active"
	statusError           = "Error"
	policyTypeDeletion    = "deletion"
	policyTypeAutoScaling = "autoscaling"
)

const (
	autoscalingGroupResourcePath     = "/groups"
	eventsResourcePath               = "/events"
	launchConfigurationsResourcePath = "/launch_configs"
	nodesResourcePath                = "nodes"
	policiesResourcePath             = "policies"
	quotasResourcePath               = "/quotas"
	schedulesResourcePath            = "cron_triggers"
	suggestionResourcePath           = "/suggestion"
	tasksResourcePath                = "/task"
	usingResourcePath                = "/using_resource"
	webhooksResourcePath             = "webhooks"
)

var (
	actionTypeSupportedWebhooks = []string{
		"CLUSTER SCALE IN",
		"CLUSTER SCALE OUT",
	}
	networkPlan = []string{
		"free_datatransfer",
		"free_bandwidth",
	}
)

var _ AutoScalingService = (*autoscalingService)(nil)

type autoscalingService struct {
	client *Client
}

/*
AutoScalingService is interface wrap others resource's interfaces. Includes:
1. AutoScalingGroups: Provides function interact with an autoscaling group such as:
    Create, Update, Delete
2. Events: Provides function to list events of an autoscaling group
3. LaunchConfigurations: Provides function to interact with launch configurations
4. Nodes: Provides function to interact with members of autoscaling group
5. Policies: Provides function to interact with autoscaling policies of autoscaling group
6. Schedules: Provides function to interact with schedule for auto scaling
7. Tasks: Provides function to get information of task
8. Webhooks: Provides fucntion to list webhook triggers scale of autoscaling group
*/
type AutoScalingService interface {
	AutoScalingGroups() *autoScalingGroup
	Common() *common
	Events() *event
	LaunchConfigurations() *launchConfiguration
	Nodes() *node
	Policies() *policy
	Schedules() *schedule
	Tasks() *task
	Webhooks() *webhook
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

func (as *autoscalingService) Events() *event {
	return &event{client: as.client}
}

func (as *autoscalingService) Nodes() *node {
	return &node{client: as.client}
}

func (as *autoscalingService) Policies() *policy {
	return &policy{client: as.client}
}

func (as *autoscalingService) Schedules() *schedule {
	return &schedule{client: as.client}
}

func (as *autoscalingService) Tasks() *task {
	return &task{client: as.client}
}

func (as *autoscalingService) Common() *common {
	return &common{client: as.client}
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

type event struct {
	client *Client
}

type policy struct {
	client *Client
}

type node struct {
	client *Client
}

type schedule struct {
	client *Client
}

type task struct {
	client *Client
}

type common struct {
	client *Client
}

// taskResult - Struct of data was returned by workers
type taskResult struct {
	Action   string      `json:"action,omitempty"`
	Data     interface{} `json:"data,omitempty"`
	Message  string      `json:"message,omitempty"`
	Progress int         `json:"progress,omitempty"`
	Success  bool        `json:"success,omitempty"`
}

// ASTask is information of doing task
type ASTask struct {
	Ready  bool       `json:"ready"`
	Result taskResult `json:"result"`
}

// usingResource - list snapshot, ssh key using to create launch configurations
type usingResource struct {
	SSHKeys   []string `json:"ssh_keys"`
	Snapshots []string `json:"snapshots"`
}

// usingResource - list snapshot, ssh key using to create launch configurations
type autoscalingQuotas struct {
	Availability map[string]int `json:"can_create,omitempty"`
	Limited      map[string]int `json:"limited,omitempty"`
	Valid        bool           `json:"valid"`
}

// Auto Scaling Group path
func (asg *autoScalingGroup) resourcePath() string {
	return autoscalingGroupResourcePath
}

func (asg *autoScalingGroup) itemPath(id string) string {
	return strings.Join([]string{autoscalingGroupResourcePath, id}, "/")
}

// Launch Configurations path
func (lc *launchConfiguration) resourcePath() string {
	return launchConfigurationsResourcePath
}

func (lc *launchConfiguration) itemPath(id string) string {
	return strings.Join([]string{launchConfigurationsResourcePath, id}, "/")
}

// Webhook path
func (wh *webhook) resourcePath(clusterID string) string {
	return strings.Join([]string{autoscalingGroupResourcePath, clusterID, webhooksResourcePath}, "/")
}

// Events path
func (e *event) resourcePath(clusterID string, page, total int) string {
	return eventsResourcePath
}

// Policy path
func (p *policy) resourcePath(clusterID string) string {
	return strings.Join([]string{autoscalingGroupResourcePath, clusterID, policiesResourcePath}, "/")
}

func (p *policy) itemPath(clusterID, policyID string) string {
	return strings.Join([]string{autoscalingGroupResourcePath, clusterID, policiesResourcePath, policyID}, "/")
}

// Node path
func (n *node) resourcePath(clusterID string) string {
	return strings.Join([]string{autoscalingGroupResourcePath, clusterID, nodesResourcePath}, "/")
}

// Schedule path
func (s *schedule) resourcePath(clusterID string) string {
	return strings.Join([]string{autoscalingGroupResourcePath, clusterID, schedulesResourcePath}, "/")
}

func (s *schedule) itemPath(clusterID, scheduleID string) string {
	return strings.Join([]string{autoscalingGroupResourcePath, clusterID, schedulesResourcePath, scheduleID}, "/")
}

// Task path
func (t *task) resourcePath(taskID string) string {
	return strings.Join([]string{tasksResourcePath, taskID, "status"}, "/")
}

// Using Resource path
func (c *common) usingResourcePath() string {
	return usingResourcePath
}

func getQuotasResourcePath() string {
	return quotasResourcePath
}

func getSuggestionResourcePath() string {
	return suggestionResourcePath
}

func (t *task) Get(ctx context.Context, taskID string) (*ASTask, error) {
	req, err := t.client.NewRequest(ctx, http.MethodGet, autoScalingServiceName, t.resourcePath(taskID), nil)
	if err != nil {
		return nil, err
	}
	resp, err := t.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data = &ASTask{}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}

// Common
func (c *common) AutoScalingUsingResource(ctx context.Context) (*usingResource, error) {
	req, err := c.client.NewRequest(ctx, http.MethodGet, autoScalingServiceName, c.usingResourcePath(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	data := &usingResource{}

	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (c *common) AutoScalingIsValidQuotas(ctx context.Context, clusterID, ProfileID string, desiredCapacity, maxSize int) (bool, error) {
	return isValidQuotas(ctx, c.client, clusterID, ProfileID, desiredCapacity, maxSize)
}

func isValidQuotas(ctx context.Context, client *Client, clusterID, ProfileID string, desiredCapacity, maxSize int) (bool, error) {
	payload := map[string]interface{}{
		"desired_capacity": desiredCapacity,
		"max_size":         maxSize,
		"profile_id":       ProfileID,
	}

	if clusterID != "" {
		payload["cluster_id"] = clusterID
	}

	req, err := client.NewRequest(ctx, http.MethodPost, autoScalingServiceName, getQuotasResourcePath(), &payload)
	if err != nil {
		return false, err
	}

	resp, err := client.Do(ctx, req)
	if err != nil {
		return false, err
	}

	var data struct {
		Quotas autoscalingQuotas `json:"message"`
	}

	// data := &map[string]interface{}{}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return false, err
	}
	return data.Quotas.Valid, nil
}

func (c *common) AutoScalingGetSuggestion(ctx context.Context, ProfileID string, desiredCapacity, maxSize int) (interface{}, error) {
	return getSuggestion(ctx, c.client, ProfileID, desiredCapacity, maxSize)
}

// getSuggestion do get suggestion
func getSuggestion(ctx context.Context, client *Client, ProfileID string, desiredCapacity, maxSize int) (interface{}, error) {
	payload := map[string]interface{}{
		"desired_capacity": desiredCapacity,
		"max_size":         maxSize,
		"profile_id":       ProfileID,
	}

	req, err := client.NewRequest(ctx, http.MethodPost, autoScalingServiceName, getSuggestionResourcePath(), &payload)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	data := map[string]interface{}{}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}
