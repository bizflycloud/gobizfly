package gobizfly

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

var _ LoadBalancerService = (*cloudLoadBalancerService)(nil)

// LoadBalancerService is an interface to interact with Bizfly API Load Balancers endpoint.
type LoadBalancerService interface {
	Create(ctx context.Context, req *LoadBalancerCreateRequest) (*LoadBalancer, error)
	Delete(ctx context.Context, req *LoadBalancerDeleteRequest) error
	Get(ctx context.Context, id string) (*LoadBalancer, error)
	List(ctx context.Context, opts *ListOptions) ([]*LoadBalancer, error)
	Resize(ctx context.Context, id string, newType string) error
	Update(ctx context.Context, id string, req *LoadBalancerUpdateRequest) (*LoadBalancer, error)

	Listeners() *cloudLoadBalancerListenerResource
	Pools() *cloudLoadBalancerPoolResource
	HealthMonitors() *cloudLoadBalancerHealthMonitorResource
	L7Policies() *cloudLoadBalancerL7PolicyResource
	Members() *cloudLoadBalancerMemberResource
}

type ListenerHealthMonitor struct {
	Type           string `json:"type"`
	URLPath        string `json:"url_path"`
	HTTPMethod     string `json:"http_method"`
	ExpectedCodes  string `json:"expected_codes"`
	MaxRetries     int    `json:"max_retries"`
	MaxRetriesDown int    `json:"max_retries_down"`
	Delay          int    `json:"delay"`
	Timeout        int    `json:"timeout"`
}

type LoadBalancerResizeRequest struct {
	Action  string `json:"action"`
	NewType string `json:"new_type"`
}

type ListenerPool struct {
	LbAlgorithm                    string                `json:"lb_algorithm"`
	Name                           string                `json:"name"`
	Protocol                       string                `json:"protocol"`
	Members                        []string              `json:"members"`
	CloudLoadBalancerHealthMonitor ListenerHealthMonitor `json:"healthmonitor"`
}

type LoadBalancerListener struct {
	Name          string       `json:"name"`
	Protocol      string       `json:"protocol"`
	ProtocolPort  int          `json:"protocol_port"`
	DefaultTLSRef string       `json:"default_tls_container_ref,omitempty"`
	DefaultPool   ListenerPool `json:"default_pool"`
}

// LoadBalancerCreateRequest represents create new load balancer request payload.
type LoadBalancerCreateRequest struct {
	Name         string                 `json:"name"`
	Description  string                 `json:"description,omitempty"`
	NetworkType  string                 `json:"network_type"`
	VPCNetworkID string                 `json:"vip_network_id"`
	Listeners    []LoadBalancerListener `json:"listeners,omitempty"`
	Type         string                 `json:"type"`
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

type cloudLoadBalancerService struct {
	client *Client
}

func (l *cloudLoadBalancerService) resourcePath() string {
	return loadBalancersPath
}

func (l *cloudLoadBalancerService) itemPath(id string) string {
	return strings.Join([]string{loadBalancerResourcePath, id}, "/")
}

// List returns a list of load balancers' information.
func (l *cloudLoadBalancerService) List(ctx context.Context, opts *ListOptions) ([]*LoadBalancer, error) {
	req, err := l.client.NewRequest(ctx, http.MethodGet, loadBalancerServiceName, l.resourcePath(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := l.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	var data struct {
		LoadBalancers []*LoadBalancer `json:"loadbalancers"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.LoadBalancers, nil
}

// Create - creates a new load balancer.
func (l *cloudLoadBalancerService) Create(ctx context.Context, lbcr *LoadBalancerCreateRequest) (*LoadBalancer, error) {
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
	defer func() {
		_ = resp.Body.Close()
	}()

	var respData struct {
		LoadBalancer *LoadBalancer `json:"loadbalancer"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, err
	}
	return respData.LoadBalancer, nil
}

// Get - retrieves a load balancer by its ID.
func (l *cloudLoadBalancerService) Get(ctx context.Context, id string) (*LoadBalancer, error) {
	req, err := l.client.NewRequest(ctx, http.MethodGet, loadBalancerServiceName, l.itemPath(id), nil)
	if err != nil {
		return nil, err
	}
	resp, err := l.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	lb := &LoadBalancer{}
	if err := json.NewDecoder(resp.Body).Decode(lb); err != nil {
		return nil, err
	}
	return lb, nil
}

// Update - update the load balancer's information.
func (l *cloudLoadBalancerService) Update(ctx context.Context, id string, lbur *LoadBalancerUpdateRequest) (*LoadBalancer, error) {
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
	defer func() {
		_ = resp.Body.Close()
	}()

	var respData struct {
		LoadBalancer *LoadBalancer `json:"loadbalancer"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, err
	}
	return respData.LoadBalancer, nil
}

// Delete - deletes a load balancer by its ID.
func (l *cloudLoadBalancerService) Delete(ctx context.Context, lbdr *LoadBalancerDeleteRequest) error {
	req, err := l.client.NewRequest(ctx, http.MethodDelete, loadBalancerServiceName, l.itemPath(lbdr.ID), lbdr)
	if err != nil {
		return err
	}
	resp, err := l.client.Do(ctx, req)
	if err != nil {
		return err
	}
	_, _ = io.Copy(io.Discard, resp.Body)

	return resp.Body.Close()
}

func (l *cloudLoadBalancerService) Resize(ctx context.Context, id string, newType string) error {
	lrq := &LoadBalancerResizeRequest{
		NewType: newType,
		Action:  "resize",
	}
	req, err := l.client.NewRequest(ctx, http.MethodPost, loadBalancerServiceName, l.itemPath(id)+"/action", lrq)
	if err != nil {
		return err
	}
	resp, err := l.client.Do(ctx, req)
	if err != nil {
		return err
	}
	_, _ = io.Copy(io.Discard, resp.Body)
	return resp.Body.Close()
}
