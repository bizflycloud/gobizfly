package gobizfly

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

var _ LoadBalancerService = (*loadbalancer)(nil)

// LoadBalancerService is an interface to interact with BizFly API Load Balancers endpoint.
type LoadBalancerService interface {
	List(ctx context.Context, opts *ListOptions) ([]*LoadBalancer, error)
	Create(ctx context.Context, req *LoadBalancerCreateRequest) (*LoadBalancer, error)
	Get(ctx context.Context, id string) (*LoadBalancer, error)
	Update(ctx context.Context, id string, req *LoadBalancerUpdateRequest) (*LoadBalancer, error)
	Delete(ctx context.Context, req *LoadBalancerDeleteRequest) error
}

// LoadBalancerCreateRequest represents create new load balancer request payload.
type LoadBalancerCreateRequest struct {
	Description  string   `json:"description,omitempty"`
	Type         string   `json:"type"`
	Listeners    []string `json:"listeners,omitempty"`
	Name         string   `json:"name"`
	NetworkType  string   `json:"network_type"`
	VPCNetworkID string   `json:"vip_network_id"`
}

// LoadBalancerUpdateRequest represents update load balancer request payload.
type LoadBalancerUpdateRequest struct {
	Name         *string `json:"name,omitempty"`
	Description  *string `json:"description,omitempty"`
	AdminStateUp *bool   `json:"admin_state_up,omitempty"`
}

// LoadBalancerDeleteRequest represents delete load balancer request payload.
type LoadBalancerDeleteRequest struct {
	Cascade bool   `json:"cascade"`
	ID      string `json:"loadbalancer_id"`
}

// LoadBalancer contains load balancer information.
type LoadBalancer struct {
	ID                 string       `json:"id"`
	FlavorID           string       `json:"flavor_id"`
	Description        string       `json:"description"`
	Provider           string       `json:"provider"`
	UpdatedAt          string       `json:"updated_at"`
	Listeners          []resourceID `json:"listeners"`
	VipSubnetID        string       `json:"vip_subnet_id"`
	ProjectID          string       `json:"project_id"`
	VipQosPolicyID     string       `json:"vip_qos_policy_id"`
	VipNetworkID       string       `json:"vip_network_id"`
	NetworkType        string       `json:"network_type"`
	VipAddress         string       `json:"vip_address"`
	VipPortID          string       `json:"vip_port_id"`
	AdminStateUp       bool         `json:"admin_state_up"`
	Name               string       `json:"name"`
	OperatingStatus    string       `json:"operating_status"`
	ProvisioningStatus string       `json:"provisioning_status"`
	Pools              []resourceID `json:"pools"`
	Type               string       `json:"type"`
	TenantID           string       `json:"tenant_id"`
	CreatedAt          string       `json:"created_at"`
}

type loadbalancer struct {
	client *Client
}

func (l *loadbalancer) resourcePath() string {
	return loadBalancersPath
}

func (l *loadbalancer) itemPath(id string) string {
	return strings.Join([]string{loadBalancerResourcePath, id}, "/")
}

// List returns a list of load balancers' information.
func (l *loadbalancer) List(ctx context.Context, opts *ListOptions) ([]*LoadBalancer, error) {
	req, err := l.client.NewRequest(ctx, http.MethodGet, loadBalancerServiceName, l.resourcePath(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := l.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		LoadBalancers []*LoadBalancer `json:"loadbalancers"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.LoadBalancers, nil
}

// Create - creates a new load balancer.
func (l *loadbalancer) Create(ctx context.Context, lbcr *LoadBalancerCreateRequest) (*LoadBalancer, error) {
	var data struct {
		LoadBalancer *LoadBalancerCreateRequest `json:"loadbalancer"`
	}
	data.LoadBalancer = lbcr
	req, err := l.client.NewRequest(ctx, http.MethodPost, loadBalancerServiceName, l.resourcePath(), &data)
	if err != nil {
		return nil, err
	}
	resp, err := l.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var respData struct {
		LoadBalancer *LoadBalancer `json:"loadbalancer"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, err
	}
	return respData.LoadBalancer, nil
}

// Get - retrieves a load balancer by its ID.
func (l *loadbalancer) Get(ctx context.Context, id string) (*LoadBalancer, error) {
	req, err := l.client.NewRequest(ctx, http.MethodGet, loadBalancerServiceName, l.itemPath(id), nil)
	if err != nil {
		return nil, err
	}
	resp, err := l.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	lb := &LoadBalancer{}
	if err := json.NewDecoder(resp.Body).Decode(lb); err != nil {
		return nil, err
	}
	return lb, nil
}

// Update - update the load balancer's information.
func (l *loadbalancer) Update(ctx context.Context, id string, lbur *LoadBalancerUpdateRequest) (*LoadBalancer, error) {
	var data struct {
		LoadBalancer *LoadBalancerUpdateRequest `json:"loadbalancer"`
	}
	data.LoadBalancer = lbur
	req, err := l.client.NewRequest(ctx, http.MethodPut, loadBalancerServiceName, l.itemPath(id), &data)
	if err != nil {
		return nil, err
	}
	resp, err := l.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var respData struct {
		LoadBalancer *LoadBalancer `json:"loadbalancer"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, err
	}
	return respData.LoadBalancer, nil
}

// Delete - deletes a load balancer by its ID.
func (l *loadbalancer) Delete(ctx context.Context, lbdr *LoadBalancerDeleteRequest) error {
	req, err := l.client.NewRequest(ctx, http.MethodDelete, loadBalancerServiceName, l.itemPath(lbdr.ID), lbdr)
	if err != nil {
		return err
	}
	resp, err := l.client.Do(ctx, req)
	if err != nil {
		return err
	}
	_, _ = io.Copy(ioutil.Discard, resp.Body)

	return resp.Body.Close()
}
