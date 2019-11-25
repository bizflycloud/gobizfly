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
	memberPath = "/members"
)

var _ MemberService = (*member)(nil)

// MemberService is an interface to interact with BizFly API Members endpoint.
type MemberService interface {
	List(ctx context.Context, poolID string, opts *ListOptions) ([]*Member, error)
	Get(ctx context.Context, poolID, id string) (*Member, error)
	Update(ctx context.Context, id string, req *MemberUpdateRequest) (*Member, error)
	Delete(ctx context.Context, id string) error
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

func (m *member) List(ctx context.Context, poolID string, opts *ListOptions) ([]*Member, error) {
	path := strings.Join([]string{poolPath, poolID, "members"}, "/")
	req, err := m.client.NewRequest(ctx, http.MethodGet, path, nil)
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
	path := strings.Join([]string{poolPath, poolID, "members", id}, "/")
	req, err := m.client.NewRequest(ctx, http.MethodGet, path, nil)
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

func (m *member) Update(ctx context.Context, id string, mur *MemberUpdateRequest) (*Member, error) {
	var data struct {
		Member *MemberUpdateRequest `json:"member"`
	}
	data.Member = mur
	req, err := m.client.NewRequest(ctx, http.MethodPut, memberPath+"/"+id, &data)
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
	return respData.Member, err
}

func (m *member) Delete(ctx context.Context, id string) error {
	req, err := m.client.NewRequest(ctx, http.MethodDelete, memberPath+"/"+id, nil)
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
