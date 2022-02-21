package gobizfly

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// SecretsCreateRequest contains receiver information.
type SecretsCreateRequest struct {
	Name string `json:"name,omitempty"`
}

// SecretCreateRequest represents create new secret request payload.
type SecretCreateRequest struct {
	Name string `json:"name,omitempty"`
}

// Secrets contains secrets information.
type Secrets struct {
	Created   string `json:"_created,omitempty"`
	ID        string `json:"_id"`
	Name      string `json:"name"`
	ProjectID string `json:"project_id,omitempty"`
	Secret    string `json:"secret,omitempty"`
	UserID    string `json:"user_id,omitempty"`
}

func (s *secrets) resourcePath() string {
	return strings.Join([]string{secretsResourcePath}, "/")
}

func (s *secrets) itemPath(id string) string {
	return strings.Join([]string{secretsResourcePath, id}, "/")
}

// List secrets
func (s *secrets) List(ctx context.Context, filters *string) ([]*Secrets, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, cloudwatcherServiceName, s.resourcePath(), nil)
	if err != nil {
		return nil, err
	}

	if filters != nil {
		q := req.URL.Query()
		q.Add("where", *filters)
		req.URL.RawQuery = q.Encode()
	}

	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		Secrets []*Secrets `json:"_items"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Secrets, nil
}

// Create a secret
func (s *secrets) Create(ctx context.Context, scr *SecretsCreateRequest) (*ResponseRequest, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, cloudwatcherServiceName, s.resourcePath(), &scr)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var respData = &ResponseRequest{}
	if err := json.NewDecoder(resp.Body).Decode(respData); err != nil {
		return nil, err
	}
	return respData, nil
}

// Get a secret
func (s *secrets) Get(ctx context.Context, id string) (*Secrets, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, cloudwatcherServiceName, s.itemPath(id), nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	secret := &Secrets{}
	if err := json.NewDecoder(resp.Body).Decode(secret); err != nil {
		return nil, err
	}
	return secret, nil
}

// Delete secret
func (s *secrets) Delete(ctx context.Context, id string) error {
	req, err := s.client.NewRequest(ctx, http.MethodDelete, cloudwatcherServiceName, s.itemPath(id), nil)
	if err != nil {
		return err
	}
	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return err
	}
	_, _ = io.Copy(ioutil.Discard, resp.Body)

	return resp.Body.Close()
}
