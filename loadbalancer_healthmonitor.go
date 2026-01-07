package gobizfly

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// CloudLoadBalancerHealthMonitor represent the load balancer health monitor information
type CloudLoadBalancerHealthMonitor struct {
	Name                  string       `json:"name"`
	Type                  string       `json:"type"`
	Delay                 int          `json:"delay"`
	MaxRetries            int          `json:"max_retries"`
	MaxRetriesDown        int          `json:"max_retries_down"`
	TimeOut               int          `json:"timeout"`
	HTTPMethod            string       `json:"http_method"`
	URLPath               string       `json:"url_path"`
	ExpectedCodes         string       `json:"expected_codes"`
	HTTPVersion           float32      `json:"http_version"`
	OperatingStatus       string       `json:"operating_status"`
	ProvisioningStatus    string       `json:"provisioning_status"`
	DomainName            string       `json:"domain_name"`
	ID                    string       `json:"id"`
	CreatedAt             string       `json:"created_at"`
	UpdatedAt             string       `json:"updated_at"`
	TenantID              string       `json:"tenant_id"`
	CloudLoadBalancerPool []resourceID `json:"pool"`
}

// CloudLoadBalancerHealthMonitorCreateRequest represent the request bodfor creating a health monitor
type CloudLoadBalancerHealthMonitorCreateRequest struct {
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

// CloudLoadBalancerHealthMonitorUpdateRequest represent the request bodfor updating a health monitor
type CloudLoadBalancerHealthMonitorUpdateRequest struct {
	Name           string   `json:"name"`
	TimeOut        *int     `json:"timeout,omitempty"`
	Delay          *int     `json:"delay,omitempty"`
	MaxRetries     *int     `json:"max_retries,omitempty"`
	MaxRetriesDown *int     `json:"max_retries_down,omitempty"`
	HTTPMethod     *string  `json:"http_method,omitempty"`
	HTTPVersion    *float32 `json:"http_version,omitempty"`
	URLPath        *string  `json:"url_path,omitempty"`
	ExpectedCodes  *string  `json:"expected_codes,omitempty"`
	DomainName     *string  `json:"domain_name,omitempty"`
}

type cloudLoadBalancerHealthMonitorResource struct {
	client *Client
}

var _ HealthMonitorService = (*cloudLoadBalancerHealthMonitorResource)(nil)

func (lbs *cloudLoadBalancerService) HealthMonitors() *cloudLoadBalancerHealthMonitorResource {
	return &cloudLoadBalancerHealthMonitorResource{client: lbs.client}
}

// HealthMonitorService is an interface to interact with Bizfly API Health Monitor endpoint.
type HealthMonitorService interface {
	Get(ctx context.Context, healthMonitorID string) (*CloudLoadBalancerHealthMonitor, error)
	Delete(ctx context.Context, healthMonitorID string) error
	Create(ctx context.Context, poolID string, hmcr *CloudLoadBalancerHealthMonitorCreateRequest) (*CloudLoadBalancerHealthMonitor, error)
	Update(Ctx context.Context, healthMonitorID string, hmur *CloudLoadBalancerHealthMonitorUpdateRequest) (*CloudLoadBalancerHealthMonitor, error)
}

func (h *cloudLoadBalancerHealthMonitorResource) itemPath(hmID string) string {
	return strings.Join([]string{healthMonitorPath, hmID}, "/")
}

// Get gets detail a health monitor
func (h *cloudLoadBalancerHealthMonitorResource) Get(ctx context.Context, hmID string) (*CloudLoadBalancerHealthMonitor, error) {
	req, err := h.client.NewRequest(ctx, http.MethodGet, loadBalancerServiceName, h.itemPath(hmID), nil)
	if err != nil {
		return nil, err
	}
	resp, err := h.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	hm := &CloudLoadBalancerHealthMonitor{}
	if err := json.NewDecoder(resp.Body).Decode(hm); err != nil {
		return nil, err
	}
	return hm, nil
}

// Create creates a health monitor for a pool
func (h *cloudLoadBalancerHealthMonitorResource) Create(ctx context.Context, poolID string, hmcr *CloudLoadBalancerHealthMonitorCreateRequest) (*CloudLoadBalancerHealthMonitor, error) {
	var data struct {
		CloudLoadBalancerHealthMonitor *CloudLoadBalancerHealthMonitorCreateRequest `json:"healthmonitor"`
	}
	hmcr.PoolID = poolID
	data.CloudLoadBalancerHealthMonitor = hmcr
	req, err := h.client.NewRequest(ctx, http.MethodPost, loadBalancerServiceName, "/"+strings.Join([]string{"pool", poolID, "healthmonitor"}, "/"), &data)
	if err != nil {
		return nil, err
	}
	resp, err := h.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	var respData struct {
		CloudLoadBalancerHealthMonitor *CloudLoadBalancerHealthMonitor `json:"healthmonitor"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, err
	}
	return respData.CloudLoadBalancerHealthMonitor, nil
}

// Delete deletes a health monitor
func (h *cloudLoadBalancerHealthMonitorResource) Delete(ctx context.Context, hmID string) error {
	req, err := h.client.NewRequest(ctx, http.MethodDelete, loadBalancerServiceName, h.itemPath(hmID), nil)
	if err != nil {
		return err
	}
	resp, err := h.client.Do(ctx, req)
	if err != nil {
		return err
	}
	_, _ = io.Copy(io.Discard, resp.Body)

	return resp.Body.Close()
}

// Update - updates a health monitor
func (h *cloudLoadBalancerHealthMonitorResource) Update(ctx context.Context, hmID string, hmur *CloudLoadBalancerHealthMonitorUpdateRequest) (*CloudLoadBalancerHealthMonitor, error) {
	var data struct {
		CloudLoadBalancerHealthMonitor *CloudLoadBalancerHealthMonitorUpdateRequest `json:"healthmonitor"`
	}
	data.CloudLoadBalancerHealthMonitor = hmur
	req, err := h.client.NewRequest(ctx, http.MethodPut, loadBalancerServiceName, h.itemPath(hmID), data)
	if err != nil {
		return nil, err
	}
	resp, err := h.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	var respData struct {
		CloudLoadBalancerHealthMonitor *CloudLoadBalancerHealthMonitor `json:"healthmonitor"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, err
	}
	return respData.CloudLoadBalancerHealthMonitor, nil
}
