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
	poolPath = "/pools"
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
	LoadBalancerID     *string             `json:"load_balancer_id"`
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

func (l *pool) List(ctx context.Context, lbID string, opts *ListOptions) ([]*Pool, error) {
	path := strings.Join([]string{loadBalancerPath, lbID, "pools"}, "/")
	req, err := l.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	resp, err := l.client.Do(ctx, req)
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

func (l *pool) Create(ctx context.Context, lbID string, pcr *PoolCreateRequest) (*Pool, error) {
	var data struct {
		Pool *PoolCreateRequest `json:"pool"`
	}
	data.Pool = pcr
	path := strings.Join([]string{loadBalancerPath, lbID, "pools"}, "/")
	req, err := l.client.NewRequest(ctx, http.MethodPost, path, &data)
	if err != nil {
		return nil, err
	}
	resp, err := l.client.Do(ctx, req)
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
	return respData.Pool, err
}

func (l *pool) Get(ctx context.Context, id string) (*Pool, error) {
	req, err := l.client.NewRequest(ctx, http.MethodGet, poolPath+"/"+id, nil)
	if err != nil {
		return nil, err
	}
	resp, err := l.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	lb := &Pool{}
	if err := json.NewDecoder(resp.Body).Decode(lb); err != nil {
		return nil, err
	}
	return lb, nil
}

func (l *pool) Update(ctx context.Context, id string, lbur *PoolUpdateRequest) (*Pool, error) {
	var data struct {
		Pool *PoolUpdateRequest
	}
	data.Pool = lbur
	req, err := l.client.NewRequest(ctx, http.MethodPut, poolPath+"/"+id, data)
	if err != nil {
		return nil, err
	}
	resp, err := l.client.Do(ctx, req)
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
	return respData.Pool, err
}

func (l *pool) Delete(ctx context.Context, id string) error {
	req, err := l.client.NewRequest(ctx, http.MethodDelete, poolPath+"/"+id, nil)
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
