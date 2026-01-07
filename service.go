// This file is part of gobizfly

package gobizfly

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

var _ ServiceInterface = (*service)(nil)

const serviceURL = "/api/auth/service"

type Service struct {
	Name          string `json:"name"`
	Code          string `json:"code"`
	CanonicalName string `json:"canonical_name"`
	ID            int    `json:"id"`
	Region        string `json:"region"`
	Icon          string `json:"icon"`
	Description   string `json:"description"`
	Enabled       bool   `json:"enabled"`
	ServiceURL    string `json:"service_url"`
}

type ServiceList struct {
	Services []*Service `json:"services"`
}

type service struct {
	client *Client
}

type ServiceInterface interface {
	List(ctx context.Context) ([]*Service, error)
	//GetEndpoint(ctx context.Context, name string, region string) (string, error)
}

func (s *service) List(ctx context.Context) ([]*Service, error) {
	u, err := s.client.apiURL.Parse(serviceURL)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)

	req, err := http.NewRequest("GET", u.String(), buf)

	if s.client.basicAuth != "" {
		req.Header.Set("Authorization", "Basic "+s.client.basicAuth)
	}
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	var services ServiceList

	if err := json.NewDecoder(resp.Body).Decode(&services); err != nil {
		return nil, err
	}
	return services.Services, nil
}
