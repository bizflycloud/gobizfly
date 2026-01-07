package gobizfly

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

var _ CloudLoadBalancerPoolService = (*cloudLoadBalancerPoolResource)(nil)

// CloudLoadBalancerPoolService is an interface to interact with Bizfly API Pools endpoint.
type CloudLoadBalancerPoolService interface {
	List(ctx context.Context, loadBalancerID string, opts *ListOptions) ([]*CloudLoadBalancerPool, error)
	Create(ctx context.Context, loadBalancerID string, req *CloudLoadBalancerPoolCreateRequest) (*CloudLoadBalancerPool, error)
	Get(ctx context.Context, id string) (*CloudLoadBalancerPool, error)
	Update(ctx context.Context, id string, req *CloudLoadBalancerPoolUpdateRequest) (*CloudLoadBalancerPool, error)
	Delete(ctx context.Context, id string) error
}

// SessionPersistence object controls how LoadBalancer sends request to backend.
// See https://support.bizflycloud.vn/api/loadbalancer/#post-loadbalancer-load_balancer_id-pools
type SessionPersistence struct {
	Type                   string  `json:"type"`
	CookieName             *string `json:"cookie_name,omitempty"`
	PersistenceTimeout     *string `json:"persistence_timeout,omitempty"`
	PersistenceGranularity *string `json:"persistence_granularity,omitempty"`
}

type CloudLoadBalancerPoolHealthMonitorRequest struct {
	ID             string `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Delay          int    `json:"delay"`
	ExpectedCodes  string `json:"expected_codes,omitempty"`
	HttpMethod     string `json:"http_method,omitempty"`
	MaxRetries     int    `json:"max_retries"`
	MaxRetriesDown int    `json:"max_retries_down"`
	Timeout        int    `json:"timeout"`
	Type           string `json:"type"`
	URLPath        string `json:"url_path,omitempty"`
}

type CloudLoadBalancerPoolMemberRequest struct {
	ID          int    `json:"id"`
	Address     string `json:"address"`
	Name        string `json:"name"`
	Weight      int    `json:"weight"`
	Port        int    `json:"protocol_port"`
	NetworkName string `json:"network_name,omitempty"`
}

// CloudLoadBalancerPoolCreateRequest represents create new pool request payload.
type CloudLoadBalancerPoolCreateRequest struct {
	LBAlgorithm        string                                     `json:"lb_algorithm"`
	ListenerID         string                                     `json:"listener_id,omitempty"`
	Name               *string                                    `json:"name,omitempty"`
	Protocol           string                                     `json:"protocol"`
	SessionPersistence *SessionPersistence                        `json:"session_persistence,omitempty"`
	HealthMonitor      *CloudLoadBalancerPoolHealthMonitorRequest `json:"healthmonitor,omitempty"`
	Members            []CloudLoadBalancerPoolMemberRequest       `json:"members,omitempty"`
}

// CloudLoadBalancerPoolUpdateRequest represents update pool request payload.
type CloudLoadBalancerPoolUpdateRequest struct {
	AdminStateUp       *bool                                      `json:"admin_state_up,omitempty"`
	Description        *string                                    `json:"description,omitempty"`
	LBAlgorithm        *string                                    `json:"lb_algorithm,omitempty"`
	Name               *string                                    `json:"name,omitempty"`
	SessionPersistence *SessionPersistence                        `json:"session_persistence"`
	Members            []CloudLoadBalancerPoolMemberRequest       `json:"members,omitempty"`
	HealthMonitor      *CloudLoadBalancerPoolHealthMonitorRequest `json:"healthmonitor,omitempty"`
}

// CloudLoadBalancerPool contains pool information.
type CloudLoadBalancerPool struct {
	ID                 string                          `json:"id"`
	TenandID           string                          `json:"tenant_id"`
	Description        string                          `json:"description"`
	LBAlgorithm        string                          `json:"lb_algorithm"`
	Name               string                          `json:"name"`
	HealthMonitor      *CloudLoadBalancerHealthMonitor `json:"healthmonitor"`
	UpdatedAt          string                          `json:"updated_at"`
	OperatingStatus    string                          `json:"operating_status"`
	Listeners          []resourceID                    `json:"listeners"`
	SessionPersistence *SessionPersistence             `json:"session_persistence"`
	ProvisoningStatus  string                          `json:"provisioning_status"`
	ProjectID          string                          `json:"project_id"`
	LoadBalancers      []resourceID                    `json:"loadbalancers"`
	AdminStateUp       bool                            `json:"admin_state_up"`
	Protocol           string                          `json:"protocol"`
	CreatedAt          string                          `json:"created_at"`
	HealthMonitorID    string                          `json:"healthmonitor_id"`
	Members            []CloudLoadBalancerMember       `json:"members"`
}

type cloudLoadBalancerPoolResource struct {
	client *Client
}

func (lbs *cloudLoadBalancerService) Pools() *cloudLoadBalancerPoolResource {
	return &cloudLoadBalancerPoolResource{client: lbs.client}
}

func (p *cloudLoadBalancerPoolResource) resourcePath(lbID string) string {
	return strings.Join([]string{loadBalancerResourcePath, lbID, "pools"}, "/")
}

func (p *cloudLoadBalancerPoolResource) itemPath(id string) string {
	return strings.Join([]string{poolPath, id}, "/")
}

// List - retrieves a list of pools' information.
func (p *cloudLoadBalancerPoolResource) List(ctx context.Context, lbID string, opts *ListOptions) ([]*CloudLoadBalancerPool, error) {
	req, err := p.client.NewRequest(ctx, http.MethodGet, loadBalancerServiceName, p.resourcePath(lbID), nil)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	var data struct {
		Pools []*CloudLoadBalancerPool `json:"pools"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Pools, nil
}

// Create - Create a new pool
func (p *cloudLoadBalancerPoolResource) Create(ctx context.Context, lbID string, pcr *CloudLoadBalancerPoolCreateRequest) (*CloudLoadBalancerPool, error) {
	var data struct {
		CloudLoadBalancerPool *CloudLoadBalancerPoolCreateRequest `json:"pool"`
	}
	data.CloudLoadBalancerPool = pcr
	req, err := p.client.NewRequest(ctx, http.MethodPost, loadBalancerServiceName, p.resourcePath(lbID), &data)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	var respData struct {
		CloudLoadBalancerPool *CloudLoadBalancerPool `json:"pool"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, err
	}
	return respData.CloudLoadBalancerPool, nil
}

// Get - Get a pool's information
func (p *cloudLoadBalancerPoolResource) Get(ctx context.Context, id string) (*CloudLoadBalancerPool, error) {
	req, err := p.client.NewRequest(ctx, http.MethodGet, loadBalancerServiceName, p.itemPath(id), nil)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	pool := &CloudLoadBalancerPool{}
	if err := json.NewDecoder(resp.Body).Decode(pool); err != nil {
		return nil, err
	}
	return pool, nil
}

// Update - Update a pool's information
func (p *cloudLoadBalancerPoolResource) Update(ctx context.Context, id string, pur *CloudLoadBalancerPoolUpdateRequest) (*CloudLoadBalancerPool, error) {
	var data struct {
		CloudLoadBalancerPool *CloudLoadBalancerPoolUpdateRequest `json:"pool"`
	}
	data.CloudLoadBalancerPool = pur
	req, err := p.client.NewRequest(ctx, http.MethodPut, loadBalancerServiceName, p.itemPath(id), data)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	var respData struct {
		CloudLoadBalancerPool *CloudLoadBalancerPool `json:"pool"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, err
	}
	return respData.CloudLoadBalancerPool, nil
}

// Delete - Delete a pool
func (p *cloudLoadBalancerPoolResource) Delete(ctx context.Context, id string) error {
	req, err := p.client.NewRequest(ctx, http.MethodDelete, loadBalancerServiceName, p.itemPath(id), nil)
	if err != nil {
		return err
	}
	resp, err := p.client.Do(ctx, req)
	if err != nil {
		return err
	}
	_, _ = io.Copy(io.Discard, resp.Body)

	return resp.Body.Close()
}
