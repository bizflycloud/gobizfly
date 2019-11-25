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
	listenerPath = "/listeners"
)

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
	Listeners              *Listener              `json:"listeners"`
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

func (l *listener) List(ctx context.Context, lbID string, opts *ListOptions) ([]*Listener, error) {
	path := strings.Join([]string{loadBalancerPath, lbID, "listeners"}, "/")
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
		Listeners []*Listener `json:"listeners"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Listeners, nil
}

func (l *listener) Create(ctx context.Context, lbID string, lcr *ListenerCreateRequest) (*Listener, error) {
	var data struct {
		Listener *ListenerCreateRequest
	}
	data.Listener = lcr
	path := strings.Join([]string{loadBalancerPath, lbID, "listeners"}, "/")
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
		Listener *Listener `json:"listener"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, err
	}
	return respData.Listener, err
}

func (l *listener) Get(ctx context.Context, id string) (*Listener, error) {
	req, err := l.client.NewRequest(ctx, http.MethodGet, listenerPath+"/"+id, nil)
	if err != nil {
		return nil, err
	}
	resp, err := l.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	lb := &Listener{}
	if err := json.NewDecoder(resp.Body).Decode(lb); err != nil {
		return nil, err
	}
	return lb, nil
}

func (l *listener) Update(ctx context.Context, id string, lur *ListenerUpdateRequest) (*Listener, error) {
	var data struct {
		Listener *ListenerUpdateRequest
	}
	data.Listener = lur
	req, err := l.client.NewRequest(ctx, http.MethodPut, listenerPath+"/"+id, &data)
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

func (l *listener) Delete(ctx context.Context, id string) error {
	req, err := l.client.NewRequest(ctx, http.MethodDelete, listenerPath+"/"+id, nil)
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
