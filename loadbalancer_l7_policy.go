package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

// L7PolicyService is an interface to interact with Bizfly API L7Policy endpoint.
type L7PolicyService interface {
	Create(ctx context.Context, listenerId string, paylaod *CreateL7PolicyRequest) (*DetailL7Policy, error)
	Get(ctx context.Context, policyId string) (*DetailL7Policy, error)
	Update(ctx context.Context, policyId string, payload *UpdateL7PolicyRequest) (*DetailL7Policy, error)
	Delete(ctx context.Context, policyId string) error
	ListL7PolicyRules(ctx context.Context, policyId string) ([]DetailL7PolicyRule, error)
}

// L7PolicyRuleRequest is rule of l7 policy payload
type L7PolicyRuleRequest struct {
	Invert      bool   `json:"invert"`
	Type        string `json:"type"`
	CompareType string `json:"compare_type"`
	Key         string `json:"key"`
	Value       string `json:"value"`
}

// CreateL7PolicyRequest is create l7 policy payload
type CreateL7PolicyRequest struct {
	Action         string                `json:"action"`
	Description    string                `json:"description"`
	Name           string                `json:"name"`
	Position       string                `json:"position"`
	RedirectPoolId string                `json:"redirect_pool_id"`
	RedirectPrefix *string               `json:"redirect_prefix"`
	RedirectUrl    string                `json:"redirect_url"`
	Rules          []L7PolicyRuleRequest `json:"rules"`
}

// L7PolicyRule is l7 policy rule id response
type L7PolicyRule struct {
	Id string `json:"id"`
}

// DetailL7PolicyRule is detail l7 policy rule response
type DetailL7PolicyRule struct {
	Id                 string  `json:"id"`
	AdminStateUp       bool    `json:"admin_state_up"`
	CompareType        string  `json:"compare_type"`
	Invert             bool    `json:"invert"`
	Key                *string `json:"key"`
	Value              string  `json:"value"`
	OperatingStatus    string  `json:"operating_status"`
	ProjectId          string  `json:"project_id"`
	TenantId           string  `json:"tenant_id"`
	ProvisioningStatus string  `json:"provisioning_status"`
	Type               string  `json:"type"`
	CreatedAt          string  `json:"created_at"`
	UpdatedAt          string  `json:"updated_at"`
}

// DetailL7Policy is detail l7 policy response
type DetailL7Policy struct {
	Action             string         `json:"action"`
	AdminStateUp       bool           `json:"admin_state_up"`
	CreatedAt          string         `json:"created_at"`
	Description        string         `json:"description"`
	Id                 string         `json:"id"`
	ListenerId         string         `json:"listener_id"`
	Name               string         `json:"name"`
	OperatingStatus    string         `json:"operating_status"`
	Position           int            `json:"position"`
	ProjectId          string         `json:"project_id"`
	TenantId           string         `json:"tenant_id"`
	ProvisioningStatus string         `json:"provisioning_status"`
	RedirectHttpCode   *int           `json:"redirect_http_code"`
	RedirectPoolId     *string        `json:"redirect_pool_id"`
	RedirectPrefix     *string        `json:"redirect_prefix"`
	RedirectUrl        *string        `json:"redirect_url"`
	Rules              []L7PolicyRule `json:"rules"`
}

// UpdateL7PolicyRuleRequest is update l7 policy rule payload
type UpdateL7PolicyRuleRequest struct {
	L7PolicyRuleRequest `json:",inline"`
	ID                  string `json:"id"`
}

// UpdateL7PolicyRequest is update l7 policy payload
type UpdateL7PolicyRequest struct {
	Action         string                      `json:"action"`
	Description    string                      `json:"description"`
	Name           string                      `json:"name"`
	Position       int                         `json:"position"`
	RedirectPoolId *string                     `json:"redirect_pool_id"`
	RedirectPrefix *string                     `json:"redirect_prefix"`
	RedirectUrl    *string                     `json:"redirect_url"`
	Rules          []UpdateL7PolicyRuleRequest `json:"rules"`
}

type l7Policy struct {
	client *Client
}

func (p *l7Policy) itemPath(policyId string) string {
	return strings.Join([]string{"l7policy", policyId}, "/")
}

// Create - create policy for listener
func (p *l7Policy) Create(ctx context.Context, listenerId string, payload *CreateL7PolicyRequest) (*DetailL7Policy, error) {
	createL7PolicyPath := strings.Join([]string{"listen", listenerId, "l7policy"}, "/")
	clpr := struct {
		L7Policy CreateL7PolicyRequest `json:"l7policy"`
	}{L7Policy: *payload}
	req, err := p.client.NewRequest(ctx, http.MethodPost, loadBalancerServiceName, createL7PolicyPath, clpr)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var data struct {
		L7Policy DetailL7Policy `json:"l7policy"`
	}
	if err := json.NewDecoder(req.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data.L7Policy, nil
}

// Get - get detail l7 policy
func (p *l7Policy) Get(ctx context.Context, policyId string) (*DetailL7Policy, error) {
	req, err := p.client.NewRequest(ctx, http.MethodGet, loadBalancerServiceName, p.itemPath(policyId), nil)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var data DetailL7Policy
	if err := json.NewDecoder(req.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, err
}

// Update - update l7 policy
func (p *l7Policy) Update(ctx context.Context, policyId string, payload *UpdateL7PolicyRequest) (*DetailL7Policy, error) {
	ulpr := struct {
		L7Plicy UpdateL7PolicyRequest `json:"l7policy"`
	}{L7Plicy: *payload}
	req, err := p.client.NewRequest(ctx, http.MethodPut, loadBalancerServiceName, p.itemPath(policyId), ulpr)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var data struct {
		L7Policy DetailL7Policy `json:"l7policy"`
	}
	if err := json.NewDecoder(req.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data.L7Policy, nil
}

// Delete - delete l7 policy
func (p *l7Policy) Delete(ctx context.Context, policyId string) error {
	req, err := p.client.NewRequest(ctx, http.MethodDelete, loadBalancerServiceName, p.itemPath(policyId), nil)
	if err != nil {
		return err
	}
	resp, err := p.client.Do(ctx, req)
	if err != nil {
		return err
	}
	return resp.Body.Close()
}

// ListL7PolicyRules - list l7 policy rules
func (p *l7Policy) ListL7PolicyRules(ctx context.Context, policyId string) ([]DetailL7PolicyRule, error) {
	path := strings.Join([]string{p.itemPath(policyId), "rules"}, "/")
	req, err := p.client.NewRequest(ctx, http.MethodGet, loadBalancerServiceName, path, nil)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var data struct {
		Rules []DetailL7PolicyRule `json:"rules"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data.Rules, nil
}
