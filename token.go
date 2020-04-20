// Copyright 2019 The BizFly Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
)

const (
	tokenPath = "/iaas-cloud/api/token"
)

var _ TokenService = (*token)(nil)

// TokenService is an interface to interact with BizFly API token endpoint.
type TokenService interface {
	Create(ctx context.Context, request *TokenCreateRequest) (*Token, error)
}

// TokenCreateRequest represents create new token request payload.
type TokenCreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Token contains token information.
type Token struct {
	KeystoneToken string `json:"token"`
	ExpiresAt     string `json:"expires_at"`
}

type token struct {
	client *Client
}

func (t *token) Create(ctx context.Context, tcr *TokenCreateRequest) (*Token, error) {
	req, err := t.client.NewRequest(ctx, http.MethodPost, tokenPath, tcr)
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

	t.client.username = tcr.Username
	t.client.password = tcr.Password

	return tok, nil
}
