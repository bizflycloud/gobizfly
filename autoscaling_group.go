package gobizfly

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// Webhook - informaion of cluster's receiver
type Webhook struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ASWebhookIDs - information about cluster's receivers
type ASWebhookIDs struct {
	ScaleIn  Webhook `json:"scale_in"`
	ScaleOut Webhook `json:"scale_out"`
}

// ASAlarm - alarm to triggers scale
type ASAlarm struct {
	AutoScaling  []string `json:"auto_scaling"`
	LoadBalancer []string `json:"load_balancer"`
}

// ASAlarms - alarms to do trigger scale
type ASAlarms struct {
	ScaleIn  ASAlarm `json:"scale_in"`
	ScaleOut ASAlarm `json:"scale_out"`
}

// ASMetadata - Medata of an auto scaling group
type ASMetadata struct {
	DeletionPolicy   string       `json:"deletion_policy"`
	ScaleInReceiver  string       `json:"scale_in_receiver"`
	ScaleOutReceiver string       `json:"scale_out_receiver"`
	WebhookIDs       ASWebhookIDs `json:"webhook_ids"`
}

// AutoScalingDataDisk is represents for a data disk being created with servers
type AutoScalingDataDisk struct {
	DeleteOnTermination bool   `json:"delete_on_termination"`
	Size                int    `json:"size"`
	Type                string `json:"type"`
}

// AutoScalingOperatingSystem is represents for operating system being use to create servers
type AutoScalingOperatingSystem struct {
	CreateFrom string `json:"type,omitempty"`
	Error      string `json:"error,omitempty"`
	ID         string `json:"id,omitempty"`
	OSName     string `json:"os_name,omitempty"`
}

// AutoScalingNetworks - is represents for relationship between network and firewalls
type AutoScalingNetworks struct {
	ID             string    `json:"id"`
	SecurityGroups []*string `json:"security_groups,omitempty"`
}

// AutoScalingGroupCreateRequest - payload use to create auto scaling group
type AutoScalingGroupCreateRequest struct {
	DeletionPolicy       *PolicyDeletionCreateRequest `json:"deletion_policy,omitempty"`
	DesiredCapacity      int                          `json:"desired_capacity"`
	LoadBalancerPolicies *LoadBalancerPolicy          `json:"load_balancer_policy,omitempty"`
	MaxSize              int                          `json:"max_size"`
	MinSize              int                          `json:"min_size"`
	Name                 string                       `json:"name"`
	ProfileID            string                       `json:"profile_id"`
	ScaleInPolicies      *[]ScalePolicy               `json:"scale_in_policy,omitempty"`
	ScaleOutPolicies     *[]ScalePolicy               `json:"scale_out_policy,omitempty"`
}

// AutoScalingGroupUpdateRequest - payload use to update auto scaling group
type AutoScalingGroupUpdateRequest struct {
	DesiredCapacity      int                 `json:"desired_capacity"`
	LoadBalancerPolicies *LoadBalancerPolicy `json:"load_balancer_policy,omitempty"`
	MaxSize              int                 `json:"max_size"`
	MinSize              int                 `json:"min_size"`
	Name                 string              `json:"name"`
	ProfileID            string              `json:"profile_id"`
	ProfileOnly          bool                `json:"profile_only"`
}

// AutoScalingGroup - is represents an auto scaling group
type AutoScalingGroup struct {
	Created              string                    `json:"created_at"`
	Data                 map[string]interface{}    `json:"data"`
	DeletionPolicy       DeletionPolicyInformation `json:"deletion_policy,omitempty"`
	DesiredCapacity      int                       `json:"desired_capacity"`
	ID                   string                    `json:"id"`
	LoadBalancerPolicies LoadBalancerPolicy        `json:"load_balancer_policy,omitempty"`
	MaxSize              int                       `json:"max_size"`
	Metadata             ASMetadata                `json:"metadata"`
	MinSize              int                       `json:"min_size"`
	Name                 string                    `json:"name"`
	NodeIDs              []string                  `json:"node_ids"`
	ProfileID            string                    `json:"profile_id"`
	ProfileName          string                    `json:"profile_name"`
	ScaleInPolicies      []ScalePolicy             `json:"scale_in_policy,omitempty"`
	ScaleOutPolicies     []ScalePolicy             `json:"scale_out_policy,omitempty"`
	Status               string                    `json:"status"`
	TaskID               string                    `json:"task_id,omitempty"`
	Timeout              int                       `json:"timeout"`
	Updated              string                    `json:"updated_at"`
}

// List autoscaling groups
func (asg *autoScalingGroup) List(ctx context.Context, all bool) ([]*AutoScalingGroup, error) {
	req, err := asg.client.NewRequest(ctx, http.MethodGet, autoScalingServiceName, asg.resourcePath(), nil)
	if err != nil {
		return nil, err
	}

	if all {
		q := req.URL.Query()
		q.Add("all", "true")
		req.URL.RawQuery = q.Encode()
	}

	resp, err := asg.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	var data struct {
		AutoScalingGroups []*AutoScalingGroup `json:"clusters"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.AutoScalingGroups, nil
}

// Get an autoscaling group
func (asg *autoScalingGroup) Get(ctx context.Context, clusterID string) (*AutoScalingGroup, error) {
	req, err := asg.client.NewRequest(ctx, http.MethodGet, autoScalingServiceName, asg.itemPath(clusterID), nil)
	if err != nil {
		return nil, err
	}
	resp, err := asg.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	data := &AutoScalingGroup{}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}

// Delete an autoscaling group
func (asg *autoScalingGroup) Delete(ctx context.Context, clusterID string) error {
	req, err := asg.client.NewRequest(ctx, http.MethodDelete, autoScalingServiceName, asg.itemPath(clusterID), nil)
	if err != nil {
		return err
	}
	resp, err := asg.client.Do(ctx, req)
	if err != nil {
		return err
	}
	_, _ = io.Copy(io.Discard, resp.Body)

	return resp.Body.Close()
}

// Create an autoscaling group
func (asg *autoScalingGroup) Create(ctx context.Context, ascr *AutoScalingGroupCreateRequest) (*AutoScalingGroup, error) {
	if valid, _ := isValidQuotas(ctx, asg.client, "", ascr.ProfileID, ascr.DesiredCapacity, ascr.MaxSize); !valid {
		return nil, errors.New("not enough quotas to create new auto scaling group")
	}
	req, err := asg.client.NewRequest(ctx, http.MethodPost, autoScalingServiceName, asg.resourcePath(), &ascr)
	if err != nil {
		return nil, err
	}

	resp, err := asg.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	data := &AutoScalingGroup{}

	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return nil, err
	}
	return data, nil
}

// Update an autoscaling group
func (asg *autoScalingGroup) Update(ctx context.Context, clusterID string, asur *AutoScalingGroupUpdateRequest) (*AutoScalingGroup, error) {
	if valid, _ := isValidQuotas(ctx, asg.client, clusterID, asur.ProfileID, asur.DesiredCapacity, asur.MaxSize); !valid {
		return nil, errors.New("not enough quotas to update new auto scaling group")
	}

	req, err := asg.client.NewRequest(ctx, http.MethodPut, autoScalingServiceName, asg.itemPath(clusterID), &asur)
	if err != nil {
		return nil, err
	}

	resp, err := asg.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	data := &AutoScalingGroup{}

	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return nil, err
	}
	return data, nil
}
