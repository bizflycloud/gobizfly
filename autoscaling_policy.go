package gobizfly

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

// ScalePolicy - information about a scale policy
type ScalePolicy struct {
	BestEffort bool    `json:"best_effort"`
	CoolDown   int     `json:"cooldown"`
	Event      string  `json:"event,omitempty"`
	ID         *string `json:"id,omitempty"`
	MetricType string  `json:"metric"`
	RangeTime  int     `json:"range_time"`
	ScaleSize  int     `json:"number"`
	Threshold  float64 `json:"threshold"`
	Type       *string `json:"type,omitempty"`
}

// DeletePolicyHooks represents a hooks was triggered when delete node in autoscaling group
type DeletePolicyHooks struct {
	URL     string                  `json:"url"`
	Method  *string                 `json:"method,omitempty"`
	Params  *map[string]interface{} `json:"params,omitempty"`
	Headers *map[string]interface{} `json:"headers,omitempty"`
	Verify  *bool                   `json:"verify,omitempty"`
}

// DeletionPolicyInformation - represents a deletion policy
type DeletionPolicyInformation struct {
	ID                    string            `json:"id,omitempty"`
	Criteria              string            `json:"criteria"`
	DestroyAfterDeletion  bool              `json:"destroy_after_deletion,omitempty"`
	GracePeriod           int               `json:"grace_period,omitempty"`
	Hooks                 DeletePolicyHooks `json:"hooks,omitempty"`
	ReduceDesiredCapacity bool              `json:"reduce_desired_capacity,omitempty"`
}

// LoadBalancerPolicy - information of load balancer will be use for auto scaling group
type LoadBalancerPolicy struct {
	LoadBalancerID   string `json:"id"`
	LoadBalancerName string `json:"name,omitempty"`
	ServerGroupID    string `json:"server_group_id"`
	ServerGroupName  string `json:"server_group_name,omitempty"`
	ServerGroupPort  int    `json:"server_group_port"`
}

// LoadBalancerScalingPolicy represents payload load balancers in LoadBalancersPolicyCreateRequest/Update
type LoadBalancerScalingPolicy struct {
	ID         string `json:"load_balancer_id"`
	Name       string `json:"load_balancer_name,omitempty"`
	TargetID   string `json:"target_id"`
	TargetName string `json:"target_name,omitempty"`
	TargetType string `json:"target_type"`
}

// LoadBalancersPolicyCreateRequest represents payload in create a load balancer policy
type LoadBalancersPolicyCreateRequest struct {
	CoolDown      int                       `json:"cooldown,omitempty"`
	Event         string                    `json:"event,omitempty"`
	LoadBalancers LoadBalancerScalingPolicy `json:"load_balancers,omitempty"`
	MetricType    string                    `json:"metric,omitempty"`
	RangeTime     int                       `json:"range_time,omitempty"`
	ScaleSize     int                       `json:"number,omitempty"`
	Threshold     int                       `json:"threshold,omitempty"`
}

// LoadBalancersPolicyUpdateRequest represents payload in create a load balancer policy
type LoadBalancersPolicyUpdateRequest struct {
	CoolDown      int                       `json:"cooldown,omitempty"`
	Event         string                    `json:"event,omitempty"`
	LoadBalancers LoadBalancerScalingPolicy `json:"load_balancers,omitempty"`
	MetricType    string                    `json:"metric,omitempty"`
	RangeTime     int                       `json:"range_time,omitempty"`
	ScaleSize     int                       `json:"number,omitempty"`
	Threshold     int                       `json:"threshold,omitempty"`
}

// PolicyAutoScalingCreateRequest represents payload use create a  policy
type PolicyAutoScalingCreateRequest struct {
	CoolDown   int    `json:"cooldown,omitempty"`
	Event      string `json:"event,omitempty"`
	MetricType string `json:"metric,omitempty"`
	RangeTime  int    `json:"range_time,omitempty"`
	ScaleSize  int    `json:"number,omitempty"`
	Threshold  int    `json:"threshold,omitempty"`
	Type       string `json:"policy_type,omitempty"`
}

// PolicyDeletionCreateRequest represents payload use create a  policy
type PolicyDeletionCreateRequest struct {
	Criteria              string             `json:"criteria"`
	DestroyAfterDeletion  bool               `json:"destroy_after_deletion,omitempty"`
	GracePeriod           int                `json:"grace_period,omitempty"`
	Hooks                 *DeletePolicyHooks `json:"hooks,omitempty"`
	ReduceDesiredCapacity bool               `json:"reduce_desired_capacity,omitempty"`
	Type                  string             `json:"policy_type,omitempty"`
}

// PolicyAutoScalingUpdateRequest represents payload use update a  policy
type PolicyAutoScalingUpdateRequest struct {
	CoolDown   int    `json:"cooldown,omitempty"`
	Event      string `json:"event,omitempty"`
	MetricType string `json:"metric,omitempty"`
	RangeTime  int    `json:"range_time,omitempty"`
	ScaleSize  int    `json:"number,omitempty"`
	Threshold  int    `json:"threshold,omitempty"`
}

// PolicyDeletionUpdateRequest represents payload use update a  policy
type PolicyDeletionUpdateRequest struct {
	Criteria              string             `json:"criteria"`
	DestroyAfterDeletion  bool               `json:"destroy_after_deletion,omitempty"`
	GracePeriod           int                `json:"grace_period,omitempty"`
	Hooks                 *DeletePolicyHooks `json:"hooks,omitempty"`
	ReduceDesiredCapacity bool               `json:"reduce_desired_capacity,omitempty"`
	Type                  string             `json:"policy_type,omitempty"`
}

// TaskResponses is responses
type TaskResponses struct {
	Message string `json:"message"`
	TaskID  string `json:"task_id,omitempty"`
}

// AutoScalingPolicies - information of policies using for auto scaling group
type AutoScalingPolicies struct {
	ScaleInPolicies      []ScalePolicy             `json:"scale_in_policy,omitempty"`
	ScaleOutPolicies     []ScalePolicy             `json:"scale_out_policy,omitempty"`
	LoadBalancerPolicies LoadBalancerPolicy        `json:"load_balancer_policy,omitempty"`
	DeletionPolicy       DeletionPolicyInformation `json:"deletion_policy,omitempty"`
	DoingTasks           []string                  `json:"doing_task,omitempty"`
}

// List autoscaling policies of the given cluster
func (p *policy) List(ctx context.Context, clusterID string) (*AutoScalingPolicies, error) {
	if clusterID == "" {
		return nil, errors.New("Auto Scaling Group ID is required")
	}

	req, err := p.client.NewRequest(ctx, http.MethodGet, autoScalingServiceName, p.resourcePath(clusterID), nil)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data = &AutoScalingPolicies{}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}

// Delete an autoscaling policies of the given cluster
func (p *policy) Delete(ctx context.Context, clusterID, PolicyID string) error {
	req, err := p.client.NewRequest(ctx, http.MethodDelete, autoScalingServiceName, p.itemPath(clusterID, PolicyID), nil)
	if err != nil {
		return err
	}

	resp, err := p.client.Do(ctx, req)
	if err != nil {
		return err
	}
	_, _ = io.Copy(ioutil.Discard, resp.Body)

	return resp.Body.Close()
}

// CreateAutoScaling - Create an autoscaling for the given cluster
func (p *policy) CreateAutoScaling(ctx context.Context, clusterID string, pcr *PolicyAutoScalingCreateRequest) (*TaskResponses, error) {
	pcr.Type = policyTypeAutoScaling
	req, err := p.client.NewRequest(ctx, http.MethodPost, autoScalingServiceName, p.resourcePath(clusterID), &pcr)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	data := &TaskResponses{}

	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return nil, err
	}
	return data, nil
}

// CreateDeletion - Create a deletion policy for the given cluster
func (p *policy) CreateDeletion(ctx context.Context, clusterID string, pcr *PolicyDeletionCreateRequest) (*TaskResponses, error) {
	pcr.Type = policyTypeDeletion
	req, err := p.client.NewRequest(ctx, http.MethodPost, autoScalingServiceName, p.resourcePath(clusterID), &pcr)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	data := &TaskResponses{}

	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return nil, err
	}
	return data, nil
}

// CreateLoadBalancers - Create a load balancers for the given cluster
func (p *policy) CreateLoadBalancers(ctx context.Context, clusterID string, lbpcr *LoadBalancersPolicyCreateRequest) (*TaskResponses, error) {
	req, err := p.client.NewRequest(ctx, http.MethodPost, autoScalingServiceName, p.resourcePath(clusterID), &lbpcr)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	data := &TaskResponses{}

	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return nil, err
	}
	return data, nil
}

// UpdateLoadBalancers - Update load balancers policy for the given cluster
func (p *policy) UpdateLoadBalancers(ctx context.Context, clusterID, PolicyID string, lbpur *LoadBalancersPolicyUpdateRequest) (*TaskResponses, error) {
	req, err := p.client.NewRequest(ctx, http.MethodPut, autoScalingServiceName, p.itemPath(clusterID, PolicyID), &lbpur)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	data := &TaskResponses{}

	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return nil, err
	}
	return data, nil
}

// UpdateAutoScaling - Update autoscaling policy for the given cluster
func (p *policy) UpdateAutoScaling(ctx context.Context, clusterID, PolicyID string, pur *PolicyAutoScalingUpdateRequest) (*TaskResponses, error) {
	req, err := p.client.NewRequest(ctx, http.MethodPut, autoScalingServiceName, p.itemPath(clusterID, PolicyID), &pur)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	data := &TaskResponses{}

	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return nil, err
	}
	return data, nil
}

// Get - Get a policy of the given cluster
func (p *policy) Get(ctx context.Context, clusterID, PolicyID string) (*ScalePolicy, error) {
	req, err := p.client.NewRequest(ctx, http.MethodGet, autoScalingServiceName, p.itemPath(clusterID, PolicyID), nil)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	data := &ScalePolicy{}

	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return nil, err
	}
	return data, nil
}

// UpdateDeletion - Update deletion policy of the given cluster
func (p *policy) UpdateDeletion(ctx context.Context, clusterID, PolicyID string, pur *PolicyDeletionUpdateRequest) (*TaskResponses, error) {
	pur.Type = policyTypeDeletion
	req, err := p.client.NewRequest(ctx, http.MethodPut, autoScalingServiceName, p.itemPath(clusterID, PolicyID), &pur)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	data := &TaskResponses{}

	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return nil, err
	}
	return data, nil
}
