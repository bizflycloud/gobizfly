package gobizfly

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

var _ CloudLoadBalancerListenerService = (*cloudLoadBalancerListenerResource)(nil)

// CloudLoadBalancerListenerService is an interface to interact with Bizfly API Listeners endpoint.
type CloudLoadBalancerListenerService interface {
	List(ctx context.Context, loadBalancerID string, opts *ListOptions) ([]*CloudLoadBalancerListener, error)
	Create(ctx context.Context, loadBalancerID string, req *CloudLoadBalancerListenerCreateRequest) (*CloudLoadBalancerListener, error)
	Get(ctx context.Context, id string) (*CloudLoadBalancerListener, error)
	Update(ctx context.Context, id string, req *CloudLoadBalancerListenerUpdateRequest) (*CloudLoadBalancerListener, error)
	Delete(ctx context.Context, id string) error
}

// CloudLoadBalancerListenerCreateRequest represents create new listener request payload.
type CloudLoadBalancerListenerCreateRequest struct {
	TimeoutTCPInspect      *int               `json:"timeout_tcp_inspect,omitempty"`
	TimeoutMemberData      *int               `json:"timeout_member_data,omitempty"`
	TimeoutMemberConnect   *int               `json:"timeout_member_connect,omitempty"`
	TimeoutClientData      *int               `json:"timeout_client_data,omitempty"`
	SNIContainerRefs       *[]string          `json:"sni_container_refs,omitempty"`
	ProtocolPort           int                `json:"protocol_port"`
	Protocol               string             `json:"protocol"`
	Name                   *string            `json:"name,omitempty"`
	L7Policies             *[]resourceID      `json:"l7policies,omitempty"`
	InsertHeaders          *map[string]string `json:"insert_headers,omitempty"`
	Description            *string            `json:"description,omitempty"`
	DefaultTLSContainerRef *string            `json:"default_tls_container_ref,omitempty"`
	DefaultPoolID          *string            `json:"default_pool_id,omitempty"`
}

// CloudLoadBalancerListenerUpdateRequest represents update listener request payload.
type CloudLoadBalancerListenerUpdateRequest struct {
	TimeoutTCPInspect      *int               `json:"timeout_tcp_inspect,omitempty"`
	TimeoutMemberData      *int               `json:"timeout_member_data,omitempty"`
	TimeoutMemberConnect   *int               `json:"timeout_member_connect,omitempty"`
	TimeoutClientData      *int               `json:"timeout_client_data,omitempty"`
	SNIContainerRefs       *[]string          `json:"sni_container_refs,omitempty"`
	Name                   *string            `json:"name,omitempty"`
	L7Policies             *[]resourceID      `json:"l7policies,omitempty"`
	InsertHeaders          *map[string]string `json:"insert_headers,omitempty"`
	Description            *string            `json:"description,omitempty"`
	DefaultTLSContainerRef *string            `json:"default_tls_container_ref,omitempty"`
	DefaultPoolID          *string            `json:"default_pool_id,omitempty"`
	AdminStateUp           *bool              `json:"admin_state_up,omitempty"`
}

// CloudLoadBalancerListener contains listener information.
type CloudLoadBalancerListener struct {
	ID                     string            `json:"id"`
	TimeoutClientData      int               `json:"timeout_client_data"`
	Description            string            `json:"description"`
	SNIContainerRefs       []string          `json:"sni_container_refs"`
	Name                   string            `json:"name"`
	ConnectionLimit        int               `json:"connection_limit"`
	UpdatedAt              string            `json:"updated_at"`
	ProjectID              string            `json:"project_id"`
	TimeoutMemberData      int               `json:"timeout_member_data"`
	TimeoutMemberConnect   int               `json:"timeout_member_connect"`
	L7Policies             []resourceID      `json:"l7policies"`
	TenandID               string            `json:"tenant_id"`
	DefaultTLSContainerRef *string           `json:"default_tls_container_ref"`
	AdminStateUp           bool              `json:"admin_state_up"`
	CreatedAt              string            `json:"created_at"`
	OperatingStatus        string            `json:"operating_status"`
	ProtocolPort           int               `json:"protocol_port"`
	LoadBalancers          []resourceID      `json:"loadbalancers"`
	ProvisoningStatus      string            `json:"provisioning_status"`
	DefaultPoolID          string            `json:"default_pool_id"`
	Protocol               string            `json:"protocol"`
	InsertHeaders          map[string]string `json:"insert_headers"`
	TimeoutTCPInspect      int               `json:"timeout_tcp_inspect"`
}

type cloudLoadBalancerListenerResource struct {
	client *Client
}

func (lbs *cloudLoadBalancerService) Listeners() *cloudLoadBalancerListenerResource {
	return &cloudLoadBalancerListenerResource{client: lbs.client}
}

func (l *cloudLoadBalancerListenerResource) resourcePath(lbID string) string {
	return strings.Join([]string{loadBalancerResourcePath, lbID, "listeners"}, "/")
}

func (l *cloudLoadBalancerListenerResource) itemPath(id string) string {
	return strings.Join([]string{listenerPath, id}, "/")
}

// List returns a list of listeners' information.
func (l *cloudLoadBalancerListenerResource) List(ctx context.Context, lbID string, opts *ListOptions) ([]*CloudLoadBalancerListener, error) {
	req, err := l.client.NewRequest(ctx, http.MethodGet, loadBalancerServiceName, l.resourcePath(lbID), nil)
	if err != nil {
		return nil, err
	}
	resp, err := l.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	var data struct {
		Listeners []*CloudLoadBalancerListener `json:"listeners"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Listeners, nil
}

// Create - Create a new listener.
func (l *cloudLoadBalancerListenerResource) Create(ctx context.Context, lbID string, lcr *CloudLoadBalancerListenerCreateRequest) (*CloudLoadBalancerListener, error) {
	var data struct {
		CloudLoadBalancerListener *CloudLoadBalancerListenerCreateRequest `json:"listener"`
	}
	data.CloudLoadBalancerListener = lcr
	req, err := l.client.NewRequest(ctx, http.MethodPost, loadBalancerServiceName, l.resourcePath(lbID), &data)
	if err != nil {
		return nil, err
	}
	resp, err := l.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	var respData struct {
		CloudLoadBalancerListener *CloudLoadBalancerListener `json:"listener"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, err
	}
	return respData.CloudLoadBalancerListener, err
}

// Get - Get a listener's information
func (l *cloudLoadBalancerListenerResource) Get(ctx context.Context, id string) (*CloudLoadBalancerListener, error) {
	req, err := l.client.NewRequest(ctx, http.MethodGet, loadBalancerServiceName, l.itemPath(id), nil)
	if err != nil {
		return nil, err
	}
	resp, err := l.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	listener := &CloudLoadBalancerListener{}
	if err := json.NewDecoder(resp.Body).Decode(listener); err != nil {
		return nil, err
	}
	return listener, nil
}

// Update - Update a listener's information.
func (l *cloudLoadBalancerListenerResource) Update(ctx context.Context, id string, lur *CloudLoadBalancerListenerUpdateRequest) (*CloudLoadBalancerListener, error) {
	var data struct {
		CloudLoadBalancerListener *CloudLoadBalancerListenerUpdateRequest `json:"listener"`
	}
	data.CloudLoadBalancerListener = lur
	req, err := l.client.NewRequest(ctx, http.MethodPut, loadBalancerServiceName, l.itemPath(id), &data)
	if err != nil {
		return nil, err
	}
	resp, err := l.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	var respData struct {
		CloudLoadBalancerListener *CloudLoadBalancerListener `json:"listener"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, err
	}
	return respData.CloudLoadBalancerListener, nil
}

// Delete - Delete a listener
func (l *cloudLoadBalancerListenerResource) Delete(ctx context.Context, id string) error {
	req, err := l.client.NewRequest(ctx, http.MethodDelete, loadBalancerServiceName, l.itemPath(id), nil)
	if err != nil {
		return err
	}
	resp, err := l.client.Do(ctx, req)
	if err != nil {
		return err
	}
	_, _ = io.Copy(io.Discard, resp.Body)

	return resp.Body.Close()
}
