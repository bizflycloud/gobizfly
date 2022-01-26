package gobizfly

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// HealthMonitor represent the load balancer health monitor information
type HealthMonitor struct {
	Name           string       `json:"name"`
	Type           string       `json:"type"`
	Delay          int          `json:"delay"`
	MaxRetries     int          `json:"max_retries"`
	MaxRetriesDown int          `json:"max_retries_down"`
	TimeOut        int          `json:"timeout"`
	HTTPMethod     string       `json:"http_method"`
	UrlPath        string       `json:"url_path"`
	ExpectedCodes  string       `json:"expected_codes"`
	HTTPVersion    float32      `json:"http_version"`
	OpratingStatus string       `json:"oprating_status"`
	DomainName     string       `json:"domain_name"`
	ID             string       `json:"id"`
	CreatedAt      string       `json:"created_at"`
	UpdatedAt      string       `json:"updated_at"`
	TenantID       string       `json:"tenant_id"`
	Pool           []resourceID `json:"pool"`
}

// HealthMonitorCreateRequest represent the request bodfor creating a health monitor
type HealthMonitorCreateRequest struct {
	Name           string  `json:"name"`
	Type           string  `json:"type"`
	TimeOut        int     `json:"timeout,omitempty"`
	PoolID         string  `json:"pool_id"`
	Delay          int     `json:"delay,omitempty"`
	MaxRetries     int     `json:"max_retries,omitempty"`
	MaxRetriesDown int     `json:"max_retries_down,omitempty"`
	HTTPMethod     string  `json:"http_method,omitempty"`
	HTTPVersion    float32 `json:"http_version,omitempty"`
	URLPath        string  `json:"url_path,omitempty"`
	ExpectedCodes  string  `json:"expected_codes,omitempty"`
	DomainName     string  `json:"domain_name,omitempty"`
}

// HealthMonitorUpdateRequest represent the request bodfor updating a health monitor
type HealthMonitorUpdateRequest struct {
	Name           string  `json:"name"`
	TimeOut        int     `json:"timeout,omitempty"`
	Delay          int     `json:"delay,omitempty"`
	MaxRetries     int     `json:"max_retries,omitempty"`
	MaxRetriesDown int     `json:"max_retries_down,omitempty"`
	HTTPMethod     string  `json:"http_method,omitempty"`
	HTTPVersion    float32 `json:"http_version,omitempty"`
	URLPath        string  `json:"url_path,omitempty"`
	ExpectedCodes  string  `json:"expected_codes,omitempty"`
	DomainName     string  `json:"domain_name,omitempty"`
}

type healthmonitor struct {
	client *Client
}

var _ HealthMonitorService = (*healthmonitor)(nil)

// HealthMonitorService is an interface to interact with BizFly API Health Monitor endpoint.
type HealthMonitorService interface {
	Get(ctx context.Context, healthMonitorID string) (*HealthMonitor, error)
	Delete(ctx context.Context, healthMonitorID string) error
	Create(ctx context.Context, poolID string, hmcr *HealthMonitorCreateRequest) (*HealthMonitor, error)
	Update(Ctx context.Context, healthMonitorID string, hmur *HealthMonitorUpdateRequest) (*HealthMonitor, error)
}

func (h *healthmonitor) itemPath(hmID string) string {
	return strings.Join([]string{healthMonitorPath, hmID}, "/")
}

// Get gets detail a health monitor
func (h *healthmonitor) Get(ctx context.Context, hmID string) (*HealthMonitor, error) {
	req, err := h.client.NewRequest(ctx, http.MethodGet, loadBalancerServiceName, h.itemPath(hmID), nil)
	if err != nil {
		return nil, err
	}
	resp, err := h.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	hm := &HealthMonitor{}
	if err := json.NewDecoder(resp.Body).Decode(hm); err != nil {
		return nil, err
	}
	return hm, nil
}

// Create creates a health monitor for a pool
func (h *healthmonitor) Create(ctx context.Context, poolID string, hmcr *HealthMonitorCreateRequest) (*HealthMonitor, error) {
	var data struct {
		HealthMonitor *HealthMonitorCreateRequest `json:"healthmonitor"`
	}
	hmcr.PoolID = poolID
	data.HealthMonitor = hmcr
	req, err := h.client.NewRequest(ctx, http.MethodPost, loadBalancerServiceName, "/"+strings.Join([]string{"pool", poolID, "healthmonitor"}, "/"), &data)
	if err != nil {
		return nil, err
	}
	resp, err := h.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var respData struct {
		HealthMonitor *HealthMonitor `json:"healthmonitor"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, err
	}
	return respData.HealthMonitor, nil
}

// Delete deletes a health monitor
func (h *healthmonitor) Delete(ctx context.Context, hmID string) error {
	req, err := h.client.NewRequest(ctx, http.MethodDelete, loadBalancerServiceName, h.itemPath(hmID), nil)
	if err != nil {
		return err
	}
	resp, err := h.client.Do(ctx, req)
	if err != nil {
		return err
	}
	_, _ = io.Copy(ioutil.Discard, resp.Body)

	return resp.Body.Close()
}

// Update - updates a health monitor
func (h *healthmonitor) Update(ctx context.Context, hmID string, hmur *HealthMonitorUpdateRequest) (*HealthMonitor, error) {
	var data struct {
		HealthMonitor *HealthMonitorUpdateRequest `json:"healthmonitor"`
	}
	data.HealthMonitor = hmur
	req, err := h.client.NewRequest(ctx, http.MethodPut, loadBalancerServiceName, h.itemPath(hmID), data)
	if err != nil {
		return nil, err
	}
	resp, err := h.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var respData struct {
		HealthMonitor *HealthMonitor `json:"healthmonitor"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, err
	}
	return respData.HealthMonitor, nil
}
