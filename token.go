// This file is part of gobizfly
//
// Copyright (C) 2020  BizFly Cloud
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>

package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
)

const (
	tokenPath = "/token"
)

var _ TokenService = (*token)(nil)

// TokenService is an interface to interact with BizFly API token endpoint.
type TokenService interface {
	Create(ctx context.Context, request *TokenCreateRequest) (*Token, error)
	Refresh(ctx context.Context) (*Token, error)
}

// TokenCreateRequest represents create new token request payload.
type TokenCreateRequest struct {
	AuthMethod    string `json:"auth_method"`
	Username      string `json:"username,omitempty"`
	Password      string `json:"password,omitempty"`
	AppCredID     string `json:"credential_id,omitempty"`
	AppCredSecret string `json:"credential_secret,omitempty"`
}

// Token contains token information.
type Token struct {
	KeystoneToken string `json:"token"`
	ExpiresAt     string `json:"expires_at"`
}

type token struct {
	client *Client
}

// Create creates new token base on the information in TokenCreateRequest.
func (t *token) Create(ctx context.Context, tcr *TokenCreateRequest) (*Token, error) {
	return t.create(ctx, tcr)
}

// Refresh retrieves new token base on underlying client information.
func (t *token) Refresh(ctx context.Context) (*Token, error) {
	tcr := &TokenCreateRequest{
		AuthMethod:    t.client.authMethod,
		Username:      t.client.username,
		Password:      t.client.password,
		AppCredID:     t.client.appCredID,
		AppCredSecret: t.client.appCredSecret,
	}
	return t.create(ctx, tcr)
}

func (t *token) create(ctx context.Context, tcr *TokenCreateRequest) (*Token, error) {

	req, err := t.client.NewRequest(ctx, http.MethodPost, authServiceName, tokenPath, tcr)
	if err != nil {
		return nil, err
	}
	resp, err := t.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	tok := &Token{}
	if err := json.NewDecoder(resp.Body).Decode(tok); err != nil {
		return nil, err
	}
	// Get new services catalog after create token
	services, err := t.client.Service.List(ctx)
	if err != nil {
		return nil, err
	}

	t.client.authMethod = tcr.AuthMethod
	t.client.username = tcr.Username
	t.client.password = tcr.Password
	t.client.appCredID = tcr.AppCredID
	t.client.appCredSecret = tcr.AppCredSecret
	t.client.services = services

	return tok, nil
}
