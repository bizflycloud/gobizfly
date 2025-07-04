// This file is part of gobizfly

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

// TokenService is an interface to interact with Bizfly API token endpoint.
type TokenService interface {
	Create(ctx context.Context, request *TokenCreateRequest) (*Token, error)
	Refresh(ctx context.Context) (*Token, error)
	Init(ctx context.Context, request *TokenCreateRequest) (*Token, error)
}

// TokenCreateRequest represents create new token request payload.
type TokenCreateRequest struct {
	AppCredID     string `json:"credential_id,omitempty"`
	AppCredSecret string `json:"credential_secret,omitempty"`
	AuthMethod    string `json:"auth_method"`
	Password      string `json:"password,omitempty"`
	ProjectID     string `json:"project_id,omitempty"`
	Username      string `json:"username,omitempty"`
	AuthType      string `json:"auth_type,omitempty"`
	Token         string `json:"token,omitempty"`
}

// Token contains token information.
type Token struct {
	ExpiresAt     string `json:"expire_at"`
	KeystoneToken string `json:"token"`
	ProjectID     string `json:"project_id"`
	ProjectName   string `json:"project_name"`
}

type token struct {
	client *Client
}

// Create creates new token base on the information in TokenCreateRequest.
func (t *token) Create(ctx context.Context, tcr *TokenCreateRequest) (*Token, error) {
	return t.create(ctx, tcr)
}

func (t *token) Init(ctx context.Context, tcr *TokenCreateRequest) (*Token, error) {
	return t.init(ctx, tcr)
}

// Refresh retrieves new token base on underlying client information.
func (t *token) Refresh(ctx context.Context) (*Token, error) {
	tcr := &TokenCreateRequest{
		AuthMethod:    t.client.authMethod,
		Username:      t.client.username,
		Password:      t.client.password,
		AppCredID:     t.client.appCredID,
		AppCredSecret: t.client.appCredSecret,
		ProjectID:     t.client.projectID,
	}
	return t.create(ctx, tcr)
}

func (t *token) create(ctx context.Context, tcr *TokenCreateRequest) (*Token, error) {
	var tok Token
	if tcr.AuthType != appCredentialAuthType {
		if tcr.Token != "" {
			tokDto, err := t.getUserInfo(ctx, tcr.Token)
			if err != nil {
				return nil, err
			}
			tok = *tokDto
		} else {
			req, err := t.client.NewRequest(ctx, http.MethodPost, authServiceName, tokenPath, tcr)
			if err != nil {
				return nil, err
			}
			resp, err := t.client.Do(ctx, req)
			if err != nil {
				return nil, err
			}
			defer resp.Body.Close()

			if err := json.NewDecoder(resp.Body).Decode(&tok); err != nil {
				return nil, err
			}
		}
	}

	if len(t.client.services) == 0 {
		// Get new services catalog after create token
		services, err := t.client.Service.List(ctx)
		if err != nil {
			return nil, err
		}
		t.client.services = services
	}

	t.client.authMethod = tcr.AuthMethod
	t.client.username = tcr.Username
	t.client.password = tcr.Password
	t.client.projectID = tcr.ProjectID
	t.client.appCredID = tcr.AppCredID
	t.client.appCredSecret = tcr.AppCredSecret
	t.client.authType = tcr.AuthType

	return &tok, nil
}

func (t *token) init(ctx context.Context, tcr *TokenCreateRequest) (*Token, error) {
	var tok Token
	if tcr.AuthType != appCredentialAuthType {
		if tcr.Token != "" {
			tokDto, err := t.getUserInfo(ctx, tcr.Token)
			if err != nil {
				return nil, err
			}
			tok = *tokDto
		} else {
			req, err := t.client.NewRequest(ctx, http.MethodPost, authServiceName, tokenPath, tcr)
			if err != nil {
				return nil, err
			}
			resp, err := t.client.DoInit(ctx, req)
			if err != nil {
				return nil, err
			}
			defer resp.Body.Close()
			if err := json.NewDecoder(resp.Body).Decode(&tok); err != nil {
				return nil, err
			}
		}
	}

	if len(t.client.services) == 0 {
		// Get new services catalog after create token
		services, err := t.client.Service.List(ctx)
		if err != nil {
			return nil, err
		}
		t.client.services = services
	}

	t.client.authMethod = tcr.AuthMethod
	t.client.username = tcr.Username
	t.client.password = tcr.Password
	t.client.projectID = tcr.ProjectID
	t.client.appCredID = tcr.AppCredID
	t.client.appCredSecret = tcr.AppCredSecret
	t.client.authType = tcr.AuthType

	return &tok, nil
}

func (t *token) getUserInfo(ctx context.Context, token string) (*Token, error) {
	var tok Token
	t.client.keystoneToken = token
	services, err := t.client.Service.List(ctx)
	if err != nil {
		return nil, err
	}
	t.client.services = services
	user, err := t.client.Account.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}
	tok.ExpiresAt = user.Expires
	tok.KeystoneToken = token
	tok.ProjectID = user.TenantID
	tok.ProjectName = user.TenantName
	return &tok, nil
}
