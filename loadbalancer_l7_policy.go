package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

// L7PolicyService is an interface to interact with Bizfly API L7Policy endpoint.
type L7PolicyService interface {
	Create(ctx context.Context, listenerID string, payload *CreateL7PolicyRequest) (*DetailL7Policy, error)
	Get(ctx context.Context, policyID string) (*DetailL7Policy, error)
	Update(ctx context.Context, policyID string, payload *UpdateL7PolicyRequest) (*DetailL7Policy, error)
	Delete(ctx context.Context, policyID string) error
	ListL7PolicyRules(ctx context.Context, policyID string) ([]DetailL7PolicyRule, error)
	CreateL7PolicyRule(ctx context.Context, policyID string, payload L7PolicyRuleRequest) (*DetailL7PolicyRule, error)
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
	RedirectPoolID string                `json:"redirect_pool_id"`
	RedirectPrefix *string               `json:"redirect_prefix"`
	RedirectURL    string                `json:"redirect_url"`
	Rules          []L7PolicyRuleRequest `json:"rules"`
}

// L7PolicyRule is l7 policy rule id response
type L7PolicyRule struct {
	ID string `json:"id"`
}

// DetailL7PolicyRule is detail l7 policy rule response
type DetailL7PolicyRule struct {
	ID                 string  `json:"id"`
	AdminStateUp       bool    `json:"admin_state_up"`
	CompareType        string  `json:"compare_type"`
	Invert             bool    `json:"invert"`
	Key                *string `json:"key"`
	Value              string  `json:"value"`
	OperatingStatus    string  `json:"operating_status"`
	ProjectID          string  `json:"project_id"`
	TenantID           string  `json:"tenant_id"`
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
	ID                 string         `json:"id"`
	ListenerID         string         `json:"listener_id"`
	Name               string         `json:"name"`
	OperatingStatus    string         `json:"operating_status"`
	Position           int            `json:"position"`
	ProjectID          string         `json:"project_id"`
	TenantID           string         `json:"tenant_id"`
	ProvisioningStatus string         `json:"provisioning_status"`
	RedirectHttpCode   *int           `json:"redirect_http_code"`
	RedirectPoolID     *string        `json:"redirect_pool_id"`
	RedirectPrefix     *string        `json:"redirect_prefix"`
	RedirectURL        *string        `json:"redirect_url"`
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
	RedirectPoolID *string                     `json:"redirect_pool_id"`
	RedirectPrefix *string                     `json:"redirect_prefix"`
	RedirectURL    *string                     `json:"redirect_url"`
	Rules          []UpdateL7PolicyRuleRequest `json:"rules"`
}

type cloudLoadBalancerL7PolicyResource struct {
	client *Client
}

func (lbs *cloudLoadBalancerService) L7Policies() *cloudLoadBalancerL7PolicyResource {
	return &cloudLoadBalancerL7PolicyResource{client: lbs.client}
}

func (p *cloudLoadBalancerL7PolicyResource) itemPath(policyID string) string {
	return strings.Join([]string{l7PolicyPath, policyID}, "/")
}

// Create - create policy for listener
func (p *cloudLoadBalancerL7PolicyResource) Create(ctx context.Context, listenerID string, payload *CreateL7PolicyRequest) (*DetailL7Policy, error) {
	createL7PolicyPath := strings.Join([]string{listenerPath, listenerID, "l7policy"}, "/")
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
	defer func() {
		_ = resp.Body.Close()
	}()
	var data struct {
		L7Policy DetailL7Policy `json:"l7policy"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data.L7Policy, nil
}

// Get - get detail l7 policy
func (p *cloudLoadBalancerL7PolicyResource) Get(ctx context.Context, policyID string) (*DetailL7Policy, error) {
	req, err := p.client.NewRequest(ctx, http.MethodGet, loadBalancerServiceName, p.itemPath(policyID), nil)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	var data DetailL7Policy
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, err
}

// Update - update l7 policy
func (p *cloudLoadBalancerL7PolicyResource) Update(ctx context.Context, policyID string, payload *UpdateL7PolicyRequest) (*DetailL7Policy, error) {
	ulpr := struct {
		L7Plicy UpdateL7PolicyRequest `json:"l7policy"`
	}{L7Plicy: *payload}
	req, err := p.client.NewRequest(ctx, http.MethodPut, loadBalancerServiceName, p.itemPath(policyID), ulpr)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	var data struct {
		L7Policy DetailL7Policy `json:"l7policy"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data.L7Policy, nil
}

// Delete - delete l7 policy
func (p *cloudLoadBalancerL7PolicyResource) Delete(ctx context.Context, policyID string) error {
	req, err := p.client.NewRequest(ctx, http.MethodDelete, loadBalancerServiceName, p.itemPath(policyID), nil)
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
func (p *cloudLoadBalancerL7PolicyResource) ListL7PolicyRules(ctx context.Context, policyID string) ([]DetailL7PolicyRule, error) {
	path := strings.Join([]string{p.itemPath(policyID), "rules"}, "/")
	req, err := p.client.NewRequest(ctx, http.MethodGet, loadBalancerServiceName, path, nil)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	var data struct {
		Rules []DetailL7PolicyRule `json:"rules"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data.Rules, nil
}

func (p *cloudLoadBalancerL7PolicyResource) CreateL7PolicyRule(ctx context.Context, policyID string, payload L7PolicyRuleRequest) (*DetailL7PolicyRule, error) {
	path := strings.Join([]string{p.itemPath(policyID), "rules"}, "/")
	clpr := struct {
		Rule L7PolicyRuleRequest `json:"rule"`
	}{Rule: payload}
	req, err := p.client.NewRequest(ctx, http.MethodPost, loadBalancerServiceName, path, clpr)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	var data struct {
		Rule DetailL7PolicyRule `json:"rule"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data.Rule, nil
}
