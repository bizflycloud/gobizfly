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

// TokenService is an interface to interact with the token endpoint of BizFly API.
type TokenService interface {
	Create(ctx context.Context, request *TokenCreateRequest) (*Token, error)
}

// TokenCreateRequest represents create new token request payload.
type TokenCreateRequest struct {
	username string
	password string
}

// Token contains token information.
type Token struct {
	Token     string
	ExpiresAt string
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

	return tok, err
}
