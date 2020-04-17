// Copyright 2019 The BizFly Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gobizfly

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	loadBalancerBasePath     = "/loadbalancers/api"
	loadBalancersPath        = "loadbalancers"
	loadBalancerResourcePath = "loadbalancer"
	listenerPath             = "listener"
	poolPath                 = "pool"
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
	Description  string        `json:"description,omitempty"`
	Type         string        `json:"type"`
	Listeners    []string      `json:"listeners,omitempty"`
	Name         string        `json:"name"`
	NetworkType  string        `json:"network_type"`
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
	ID                 string                `json:"id"`
	FlavorID           string                `json:"flavor_id"`
	Description        string                `json:"description"`
	Provider           string                `json:"provider"`
	UpdatedAt          string                `json:"updated_at"`
	Listeners          []struct{ ID string } `json:"listeners"`
	VipSubnetID        string                `json:"vip_subnet_id"`
	ProjectID          string                `json:"project_id"`
	VipQosPolicyID     string                `json:"vip_qos_policy_id"`
	VipNetworkID       string                `json:"vip_network_id"`
	NetworkType        string                `json:"network_type"`
	VipAddress         string                `json:"vip_address"`
	VipPortID          string                `json:"vip_port_id"`
	AdminStateUp       bool                  `json:"admin_state_up"`
	Name               string                `json:"name"`
	OperatingStatus    string                `json:"operating_status"`
	ProvisioningStatus string                `json:"provisioning_status"`
	Pools              []struct{ ID string } `json:"pools"`
	Type               string                `json:"type"`
	TenantID           string                `json:"tenant_id"`
	CreatedAt          string                `json:"created_at"`
}

type loadbalancer struct {
	client *Client
}

func (l *loadbalancer) resourcePath() string {
	return strings.Join([]string{loadBalancerBasePath, loadBalancersPath}, "/")
}

func (l *loadbalancer) itemPath(id string) string {
	return strings.Join([]string{loadBalancerBasePath, loadBalancerResourcePath, id}, "/")
}

func (l *loadbalancer) List(ctx context.Context, opts *ListOptions) ([]*LoadBalancer, error) {
	req, err := l.client.NewRequest(ctx, http.MethodGet, l.resourcePath(), nil)
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

func (l *loadbalancer) Create(ctx context.Context, lbcr *LoadBalancerCreateRequest) (*LoadBalancer, error) {
	var data struct {
		LoadBalancer *LoadBalancerCreateRequest `json:"loadbalancer"`
	}
	data.LoadBalancer = lbcr
	req, err := l.client.NewRequest(ctx, http.MethodPost, l.resourcePath(), &data)
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

func (l *loadbalancer) Get(ctx context.Context, id string) (*LoadBalancer, error) {
	req, err := l.client.NewRequest(ctx, http.MethodGet, l.itemPath(id), nil)
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

func (l *loadbalancer) Update(ctx context.Context, id string, lbur *LoadBalancerUpdateRequest) (*LoadBalancer, error) {
	var data struct {
		LoadBalancer *LoadBalancerUpdateRequest `json:"loadbalancer"`
	}
	data.LoadBalancer = lbur
	req, err := l.client.NewRequest(ctx, http.MethodPut, l.itemPath(id), &data)
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

func (l *loadbalancer) Delete(ctx context.Context, lbdr *LoadBalancerDeleteRequest) error {
	req, err := l.client.NewRequest(ctx, http.MethodDelete, l.itemPath(lbdr.ID), lbdr)
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

var _ ListenerService = (*listener)(nil)

// ListenerService is an interface to interact with BizFly API Listeners endpoint.
type ListenerService interface {
	List(ctx context.Context, loadBalancerID string, opts *ListOptions) ([]*Listener, error)
	Create(ctx context.Context, loadBalancerID string, req *ListenerCreateRequest) (*Listener, error)
	Get(ctx context.Context, id string) (*Listener, error)
	Update(ctx context.Context, id string, req *ListenerUpdateRequest) (*Listener, error)
	Delete(ctx context.Context, id string) error
}

// ListenerCreateRequest represents create new listener request payload.
type ListenerCreateRequest struct {
	TimeoutTCPInspect      *int                   `json:"timeout_tcp_inspect,omitempty"`
	TimeoutMemberData      *int                   `json:"timeout_member_data,omitempty"`
	TimeoutMemberConnect   *int                   `json:"timeout_member_connect,omitempty"`
	TimeoutClientData      *int                   `json:"timeout_client_data,omitempty"`
	SNIContainerRefs       *[]string              `json:"sni_container_refs,omitempty"`
	ProtocolPort           int                    `json:"protocol_port"`
	Protocol               string                 `json:"protocol"`
	Name                   *string                `json:"name,omitempty"`
	L7Policies             *[]struct{ ID string } `json:"l7policies,omitempty"`
	InsertHeaders          *map[string]string     `json:"insert_headers,omitempty"`
	Description            *string                `json:"description,omitempty"`
	DefaultTLSContainerRef *string                `json:"default_tls_container_ref,omitempty"`
	DefaultPoolID          *string                `json:"default_pool_id,omitempty"`
}

// ListenerUpdateRequest represents update listener request payload.
type ListenerUpdateRequest struct {
	TimeoutTCPInspect      *int                   `json:"timeout_tcp_inspect,omitempty"`
	TimeoutMemberData      *int                   `json:"timeout_member_data,omitempty"`
	TimeoutMemberConnect   *int                   `json:"timeout_member_connect,omitempty"`
	TimeoutClientData      *int                   `json:"timeout_client_data,omitempty"`
	SNIContainerRefs       *[]string              `json:"sni_container_refs,omitempty"`
	Name                   *string                `json:"name,omitempty"`
	L7Policies             *[]struct{ ID string } `json:"l7policies,omitempty"`
	InsertHeaders          *map[string]string     `json:"insert_headers,omitempty"`
	Description            *string                `json:"description,omitempty"`
	DefaultTLSContainerRef *string                `json:"default_tls_container_ref,omitempty"`
	DefaultPoolID          *string                `json:"default_pool_id,omitempty"`
	AdminStateUp           *bool                  `json:"admin_state_up,omitempty"`
}

// Listener contains listener information.
type Listener struct {
	ID                     string                `json:"id"`
	TimeoutClientData      int                   `json:"timeout_client_data"`
	Description            string                `json:"description"`
	SNIContainerRefs       []string              `json:"sni_container_refs"`
	Name                   string                `json:"name"`
	ConnectionLimit        int                   `json:"connection_limit"`
	UpdatedAt              string                `json:"updated_at"`
	ProjectID              string                `json:"project_id"`
	TimeoutMemberData      int                   `json:"timeout_member_data"`
	TimeoutMemberConnect   int                   `json:"timeout_member_connect"`
	L7Policies             []struct{ ID string } `json:"l7policies"`
	TenandID               string                `json:"tenant_id"`
	DefaultTLSContainerRef *string               `json:"default_tls_container_ref"`
	AdminStateUp           bool                  `json:"admin_state_up"`
	CreatedAt              string                `json:"created_at"`
	OperatingStatus        string                `json:"operating_status"`
	ProtocolPort           int                   `json:"protocol_port"`
	LoadBalancers          []struct{ ID string } `json:"loadbalancers"`
	ProvisoningStatus      string                `json:"provisioning_status"`
	DefaultPoolID          string                `json:"default_pool_id"`
	Protocol               string                `json:"protocol"`
	InsertHeaders          map[string]string     `json:"insert_headers"`
	TimeoutTCPInspect      int                   `json:"timeout_tcp_inspect"`
}

type listener struct {
	client *Client
}

func (l *listener) resourcePath(lbID string) string {
	return strings.Join([]string{loadBalancerBasePath, loadBalancerResourcePath, lbID, "listeners"}, "/")
}

func (l *listener) itemPath(id string) string {
	return strings.Join([]string{loadBalancerBasePath, listenerPath, id}, "/")
}

func (l *listener) List(ctx context.Context, lbID string, opts *ListOptions) ([]*Listener, error) {
	req, err := l.client.NewRequest(ctx, http.MethodGet, l.resourcePath(lbID), nil)
	if err != nil {
		return nil, err
	}
	resp, err := l.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		Listeners []*Listener `json:"listeners"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Listeners, nil
}

func (l *listener) Create(ctx context.Context, lbID string, lcr *ListenerCreateRequest) (*Listener, error) {
	var data struct {
		Listener *ListenerCreateRequest `json:"listener"`
	}
	data.Listener = lcr
	req, err := l.client.NewRequest(ctx, http.MethodPost, l.resourcePath(lbID), &data)
	if err != nil {
		return nil, err
	}
	resp, err := l.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var respData struct {
		Listener *Listener `json:"listener"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, err
	}
	return respData.Listener, err
}

func (l *listener) Get(ctx context.Context, id string) (*Listener, error) {
	req, err := l.client.NewRequest(ctx, http.MethodGet, l.itemPath(id), nil)
	if err != nil {
		return nil, err
	}
	resp, err := l.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	listener := &Listener{}
	if err := json.NewDecoder(resp.Body).Decode(listener); err != nil {
		return nil, err
	}
	return listener, nil
}

func (l *listener) Update(ctx context.Context, id string, lur *ListenerUpdateRequest) (*Listener, error) {
	var data struct {
		Listener *ListenerUpdateRequest
	}
	data.Listener = lur
	req, err := l.client.NewRequest(ctx, http.MethodPut, l.itemPath(id), &data)
	if err != nil {
		return nil, err
	}
	resp, err := l.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var respData struct {
		Listener *Listener `json:"listener"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, err
	}
	return respData.Listener, nil
}

func (l *listener) Delete(ctx context.Context, id string) error {
	req, err := l.client.NewRequest(ctx, http.MethodDelete, l.itemPath(id), nil)
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

var _ MemberService = (*member)(nil)

// MemberService is an interface to interact with BizFly API Members endpoint.
type MemberService interface {
	List(ctx context.Context, poolID string, opts *ListOptions) ([]*Member, error)
	Get(ctx context.Context, poolID, id string) (*Member, error)
	Update(ctx context.Context, poolID, id string, req *MemberUpdateRequest) (*Member, error)
	Delete(ctx context.Context, poolID, id string) error
	Create(ctx context.Context, poolID string, req *MemberCreateRequest) (*Member, error)
}

// MemberUpdateRequest represents update member request payload.
type MemberUpdateRequest struct {
	Name           string  `json:"name"`
	Weight         int     `json:"weight"`
	AdminStateUp   bool    `json:"admin_state_up"`
	MonitorAddress *string `json:"monitor_address"`
	MonitorPort    *int    `json:"monitor_port"`
	Backup         bool    `json:"backup"`
}

// MemberCreateRequest represents create member request payload
type MemberCreateRequest struct {
	Name           string `json:"name"`
	Weight         int    `json:"weight,omitempty"`
	Address        string `json:"address"`
	ProtocolPort   int    `json:"protocol_port"`
	MonitorAddress string `json:"monitor_address,omitempty"`
	MonitorPort    int    `json:"monitor_port,omitempty"`
	Backup         bool   `json:"backup,omitempty"`
}

// Member contains member information.
type Member struct {
	ID                string  `json:"id"`
	TenandID          string  `json:"tenant_id"`
	AdminStateUp      bool    `json:"admin_state_up"`
	Name              string  `json:"name"`
	UpdatedAt         string  `json:"updated_at"`
	OperatingStatus   string  `json:"operating_status"`
	MonitorAddress    *string `json:"monitor_address"`
	ProvisoningStatus string  `json:"provisioning_status"`
	ProjectID         string  `json:"project_id"`
	ProtocolPort      int     `json:"protocol_port"`
	SubnetID          string  `json:"subnet_id"`
	MonitorPort       *int    `json:"monitor_port"`
	Address           string  `json:"address"`
	Weight            int     `json:"weight"`
	CreatedAt         string  `json:"created_at"`
	Backup            bool    `json:"backup"`
}

type member struct {
	client *Client
}

func (m *member) resourcePath(poolID string) string {
	return strings.Join([]string{loadBalancerBasePath, poolPath, poolID, "member"}, "/")
}

func (m *member) itemPath(poolID string, id string) string {
	return strings.Join([]string{loadBalancerBasePath, poolPath, poolID, "member", id}, "/")
}

func (m *member) List(ctx context.Context, poolID string, opts *ListOptions) ([]*Member, error) {
	req, err := m.client.NewRequest(ctx, http.MethodGet, m.resourcePath(poolID), nil)
	if err != nil {
		return nil, err
	}
	resp, err := m.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		Members []*Member `json:"members"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Members, nil
}

func (m *member) Get(ctx context.Context, poolID, id string) (*Member, error) {
	req, err := m.client.NewRequest(ctx, http.MethodGet, m.itemPath(poolID, id), nil)
	if err != nil {
		return nil, err
	}
	resp, err := m.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	mb := &Member{}
	if err := json.NewDecoder(resp.Body).Decode(mb); err != nil {
		return nil, err
	}
	return mb, nil
}

func (m *member) Update(ctx context.Context, poolID, id string, mur *MemberUpdateRequest) (*Member, error) {
	var data struct {
		Member *MemberUpdateRequest `json:"member"`
	}
	data.Member = mur
	req, err := m.client.NewRequest(ctx, http.MethodPut, m.itemPath(poolID, id), &data)
	if err != nil {

		return nil, err
	}
	resp, err := m.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var respData struct {
		Member *Member `json:"member"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, err
	}
	return respData.Member, nil
}

func (m *member) Delete(ctx context.Context, poolID, id string) error {
	req, err := m.client.NewRequest(ctx, http.MethodDelete, m.itemPath(poolID, id), nil)
	if err != nil {
		return err
	}
	resp, err := m.client.Do(ctx, req)
	if err != nil {
		return err
	}
	_, _ = io.Copy(ioutil.Discard, resp.Body)

	return resp.Body.Close()
}

func (m *member) Create(ctx context.Context, poolID string, mcr *MemberCreateRequest) (*Member, error) {
	var data struct {
		Member *MemberCreateRequest `json:"member"`
	}
	data.Member = mcr
	req, err := m.client.NewRequest(ctx, http.MethodPost, m.resourcePath(poolID), &data)
	if err != nil {
		return nil, err
	}
	resp, err := m.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var response struct {
		Member *Member `json:"member"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response.Member, nil
}

var _ PoolService = (*pool)(nil)

// PoolService is an interface to interact with BizFly API Pools endpoint.
type PoolService interface {
	List(ctx context.Context, loadBalancerID string, opts *ListOptions) ([]*Pool, error)
	Create(ctx context.Context, loadBalancerID string, req *PoolCreateRequest) (*Pool, error)
	Get(ctx context.Context, id string) (*Pool, error)
	Update(ctx context.Context, id string, req *PoolUpdateRequest) (*Pool, error)
	Delete(ctx context.Context, id string) error
}

// SessionPersistence object controls how LoadBalacner sends request to backend.
// See https://support.bizflycloud.vn/api/loadbalancer/#post-loadbalancer-load_balancer_id-pools
type SessionPersistence struct {
	Type                   string  `json:"type"`
	CookieName             *string `json:"cookie_name,omitempty"`
	PersistenceTimeout     *string `json:"persistence_timeout,omitempty"`
	PersistenceGranularity *string `json:"persistence_granularity,omitempty"`
}

// PoolCreateRequest represents create new pool request payload.
type PoolCreateRequest struct {
	Description        *string             `json:"description,omitempty"`
	LBAlgorithm        string              `json:"lb_algorithm"`
	ListenerID         *string             `json:"listener_id"`
	Name               *string             `json:"name,omitempty"`
	Protocol           string              `json:"protocol"`
	SessionPersistence *SessionPersistence `json:"session_persistence"`
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
	ID          string `json:"id"`
	TenandID    string `json:"tenant_id"`
	Description string `json:"description"`
	LBAlgorithm string `json:"lb_algorithm"`
	Name        string `json:"name"`
	// TODO: change later when HealthMonitor entity added
	HealthMonitor      interface{}           `json:"healthmonitor"`
	UpdatedAt          string                `json:"updated_at"`
	OperatingStatus    string                `json:"operating_status"`
	Listeners          []struct{ ID string } `json:"listeners"`
	SessionPersistence *SessionPersistence   `json:"session_persistence"`
	ProvisoningStatus  string                `json:"provisioning_status"`
	ProjectID          string                `json:"project_id"`
	LoadBalancers      []struct{ ID string } `json:"loadbalancers"`
	Members            []string              `json:"memebers"`
	AdminStateUp       bool                  `json:"admin_state_up"`
	Protocol           string                `json:"protocol"`
	CreatedAt          string                `json:"created_at"`
	HealthMonitorID    string                `json:"healthmonitor_id"`
}

type pool struct {
	client *Client
}

func (p *pool) resourcePath(lbID string) string {
	return strings.Join([]string{loadBalancerBasePath, loadBalancerResourcePath, lbID, "pools"}, "/")
}

func (p *pool) itemPath(id string) string {
	return strings.Join([]string{loadBalancerBasePath, poolPath, id}, "/")
}

func (p *pool) List(ctx context.Context, lbID string, opts *ListOptions) ([]*Pool, error) {
	req, err := p.client.NewRequest(ctx, http.MethodGet, p.resourcePath(lbID), nil)
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

func (p *pool) Create(ctx context.Context, lbID string, pcr *PoolCreateRequest) (*Pool, error) {
	var data struct {
		Pool *PoolCreateRequest `json:"pool"`
	}
	data.Pool = pcr
	req, err := p.client.NewRequest(ctx, http.MethodPost, p.resourcePath(lbID), &data)
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

func (p *pool) Get(ctx context.Context, id string) (*Pool, error) {
	req, err := p.client.NewRequest(ctx, http.MethodGet, p.itemPath(id), nil)
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

func (p *pool) Update(ctx context.Context, id string, pur *PoolUpdateRequest) (*Pool, error) {
	var data struct {
		Pool *PoolUpdateRequest `json:"pool"`
	}
	data.Pool = pur
	req, err := p.client.NewRequest(ctx, http.MethodPut, p.itemPath(id), data)
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

func (p *pool) Delete(ctx context.Context, id string) error {
	req, err := p.client.NewRequest(ctx, http.MethodDelete, p.itemPath(id), nil)
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
