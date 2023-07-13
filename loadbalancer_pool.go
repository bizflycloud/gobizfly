package gobizfly

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

var _ PoolService = (*pool)(nil)

// PoolService is an interface to interact with BizFly API Pools endpoint.
type PoolService interface {
	List(ctx context.Context, loadBalancerID string, opts *ListOptions) ([]*Pool, error)
	Create(ctx context.Context, loadBalancerID string, req *PoolCreateRequest) (*Pool, error)
	Get(ctx context.Context, id string) (*Pool, error)
	Update(ctx context.Context, id string, req *PoolUpdateRequest) (*Pool, error)
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

// PoolCreateRequest represents create new pool request payload.
type PoolCreateRequest struct {
	LBAlgorithm        string              `json:"lb_algorithm"`
	ListenerID         string              `json:"listener_id,omitempty"`
	Name               *string             `json:"name,omitempty"`
	Protocol           string              `json:"protocol"`
	SessionPersistence *SessionPersistence `json:"session_persistence,omitempty"`
}

// PoolUpdateRequest represents update pool request payload.
type PoolUpdateRequest struct {
	AdminStateUp       *bool               `json:"admin_state_up,omitempty"`
	Description        *string             `json:"description,omitempty"`
	LBAlgorithm        *string             `json:"lb_algorithm,omitempty"`
	Name               *string             `json:"name,omitempty"`
	SessionPersistence *SessionPersistence `json:"session_persistence"`
}

// Pool contains pool information.
type Pool struct {
	ID                 string              `json:"id"`
	TenandID           string              `json:"tenant_id"`
	Description        string              `json:"description"`
	LBAlgorithm        string              `json:"lb_algorithm"`
	Name               string              `json:"name"`
	HealthMonitor      *HealthMonitor      `json:"healthmonitor"`
	UpdatedAt          string              `json:"updated_at"`
	OperatingStatus    string              `json:"operating_status"`
	Listeners          []resourceID        `json:"listeners"`
	SessionPersistence *SessionPersistence `json:"session_persistence"`
	ProvisoningStatus  string              `json:"provisioning_status"`
	ProjectID          string              `json:"project_id"`
	LoadBalancers      []resourceID        `json:"loadbalancers"`
	Members            []string            `json:"memebers"`
	AdminStateUp       bool                `json:"admin_state_up"`
	Protocol           string              `json:"protocol"`
	CreatedAt          string              `json:"created_at"`
	HealthMonitorID    string              `json:"healthmonitor_id"`
}

type pool struct {
	client *Client
}

func (p *pool) resourcePath(lbID string) string {
	return strings.Join([]string{loadBalancerResourcePath, lbID, "pools"}, "/")
}

func (p *pool) itemPath(id string) string {
	return strings.Join([]string{poolPath, id}, "/")
}

// List - retrieves a list of pools' information.
func (p *pool) List(ctx context.Context, lbID string, opts *ListOptions) ([]*Pool, error) {
	req, err := p.client.NewRequest(ctx, http.MethodGet, loadBalancerServiceName, p.resourcePath(lbID), nil)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		Pools []*Pool `json:"pools"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Pools, nil
}

// Create - Create a new pool
func (p *pool) Create(ctx context.Context, lbID string, pcr *PoolCreateRequest) (*Pool, error) {
	var data struct {
		Pool *PoolCreateRequest `json:"pool"`
	}
	data.Pool = pcr
	req, err := p.client.NewRequest(ctx, http.MethodPost, loadBalancerServiceName, p.resourcePath(lbID), &data)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var respData struct {
		Pool *Pool `json:"pool"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, err
	}
	return respData.Pool, nil
}

// Get - Get a pool's information
func (p *pool) Get(ctx context.Context, id string) (*Pool, error) {
	req, err := p.client.NewRequest(ctx, http.MethodGet, loadBalancerServiceName, p.itemPath(id), nil)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	pool := &Pool{}
	if err := json.NewDecoder(resp.Body).Decode(pool); err != nil {
		return nil, err
	}
	return pool, nil
}

// Update - Update a pool's information
func (p *pool) Update(ctx context.Context, id string, pur *PoolUpdateRequest) (*Pool, error) {
	var data struct {
		Pool *PoolUpdateRequest `json:"pool"`
	}
	data.Pool = pur
	req, err := p.client.NewRequest(ctx, http.MethodPut, loadBalancerServiceName, p.itemPath(id), data)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var respData struct {
		Pool *Pool `json:"pool"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, err
	}
	return respData.Pool, nil
}

// Delete - Delete a pool
func (p *pool) Delete(ctx context.Context, id string) error {
	req, err := p.client.NewRequest(ctx, http.MethodDelete, loadBalancerServiceName, p.itemPath(id), nil)
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
