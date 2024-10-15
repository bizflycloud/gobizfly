package gobizfly

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

var _ CloudLoadBalancerMemberService = (*cloudLoadBalancerMemberResource)(nil)

// CloudLoadBalancerMemberService is an interface to interact with Bizfly API Members endpoint.
type CloudLoadBalancerMemberService interface {
	List(ctx context.Context, poolID string, opts *ListOptions) ([]*CloudLoadBalancerMember, error)
	Get(ctx context.Context, poolID, id string) (*CloudLoadBalancerMember, error)
	Update(ctx context.Context, poolID, id string, req *CloudLoadBalancerMemberUpdateRequest) (*CloudLoadBalancerMember, error)
	Delete(ctx context.Context, poolID, id string) error
	Create(ctx context.Context, poolID string, req *CloudLoadBalancerMemberCreateRequest) (*CloudLoadBalancerMember, error)
	BatchUpdate(ctx context.Context, poolID string, members *CloudLoadBalancerBatchMemberUpdateRequest) error
}

// CloudLoadBalancerMemberUpdateRequest represents update member request payload.
type CloudLoadBalancerMemberUpdateRequest struct {
	Name           string `json:"name"`
	Weight         int    `json:"weight,omitempty"`
	MonitorAddress string `json:"monitor_address,omitempty"`
	MonitorPort    int    `json:"monitor_port,omitempty"`
	Backup         bool   `json:"backup,omitempty"`
}

// CloudLoadBalancerExtendMemberUpdateRequest represents update member request payload.
type CloudLoadBalancerExtendMemberUpdateRequest struct {
	CloudLoadBalancerMemberUpdateRequest
	Address      string `json:"address"`
	ProtocolPort int    `json:"protocol_port"`
	SubnetID     string `json:"subnet_id,omitempty"`
}

// CloudLoadBalancerBatchMemberUpdateRequest represents batch update member request payload.
type CloudLoadBalancerBatchMemberUpdateRequest struct {
	Members []CloudLoadBalancerExtendMemberUpdateRequest `json:"members"`
}

// CloudLoadBalancerMemberCreateRequest represents create member request payload
type CloudLoadBalancerMemberCreateRequest struct {
	Name           string `json:"name"`
	Weight         int    `json:"weight,omitempty"`
	Address        string `json:"address"`
	ProtocolPort   int    `json:"protocol_port"`
	MonitorAddress string `json:"monitor_address,omitempty"`
	MonitorPort    int    `json:"monitor_port,omitempty"`
	Backup         bool   `json:"backup,omitempty"`
}

// CloudLoadBalancerMember contains member information.
type CloudLoadBalancerMember struct {
	ID                string  `json:"id"`
	TenandID          string  `json:"tenant_id"`
	AdminStateUp      bool    `json:"admin_state_up"`
	Name              string  `json:"name"`
	UpdatedAt         string  `json:"updated_at"`
	OperatingStatus   string  `json:"operating_status"`
	MonitorAddress    *string `json:"monitor_address"`
	ProvisoningStatus string  `json:"provisioning_status"`
	ProjectID         string  `json:"project_id"`
	ProtocolPort      int     `json:"protocol_port"`
	SubnetID          string  `json:"subnet_id"`
	MonitorPort       *int    `json:"monitor_port"`
	Address           string  `json:"address"`
	Weight            int     `json:"weight"`
	CreatedAt         string  `json:"created_at"`
	Backup            bool    `json:"backup"`
}

type cloudLoadBalancerMemberResource struct {
	client *Client
}

func (lbs *cloudLoadBalancerService) Members() *cloudLoadBalancerMemberResource {
	return &cloudLoadBalancerMemberResource{client: lbs.client}
}

func (m *cloudLoadBalancerMemberResource) resourcePath(poolID string) string {
	return strings.Join([]string{poolPath, poolID, "member"}, "/")
}

func (m *cloudLoadBalancerMemberResource) itemPath(poolID string, id string) string {
	return strings.Join([]string{poolPath, poolID, "member", id}, "/")
}

// List - list members' information
func (m *cloudLoadBalancerMemberResource) List(ctx context.Context, poolID string, opts *ListOptions) ([]*CloudLoadBalancerMember, error) {
	req, err := m.client.NewRequest(ctx, http.MethodGet, loadBalancerServiceName, m.resourcePath(poolID), nil)
	if err != nil {
		return nil, err
	}
	resp, err := m.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		Members []*CloudLoadBalancerMember `json:"members"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Members, nil
}

// Get - Get member's information
func (m *cloudLoadBalancerMemberResource) Get(ctx context.Context, poolID, id string) (*CloudLoadBalancerMember, error) {
	req, err := m.client.NewRequest(ctx, http.MethodGet, loadBalancerServiceName, m.itemPath(poolID, id), nil)
	if err != nil {
		return nil, err
	}
	resp, err := m.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	mb := &CloudLoadBalancerMember{}
	if err := json.NewDecoder(resp.Body).Decode(mb); err != nil {
		return nil, err
	}
	return mb, nil
}

// Update - Update member's information
func (m *cloudLoadBalancerMemberResource) Update(ctx context.Context, poolID, id string, mur *CloudLoadBalancerMemberUpdateRequest) (*CloudLoadBalancerMember, error) {
	var data struct {
		CloudLoadBalancerMember *CloudLoadBalancerMemberUpdateRequest `json:"member"`
	}
	data.CloudLoadBalancerMember = mur
	req, err := m.client.NewRequest(ctx, http.MethodPut, loadBalancerServiceName, m.itemPath(poolID, id), &data)
	if err != nil {

		return nil, err
	}
	resp, err := m.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var respData struct {
		CloudLoadBalancerMember *CloudLoadBalancerMember `json:"member"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, err
	}
	return respData.CloudLoadBalancerMember, nil
}

// BatchUpdate - Batch update members
func (m *cloudLoadBalancerMemberResource) BatchUpdate(ctx context.Context, poolID string, members *CloudLoadBalancerBatchMemberUpdateRequest) error {
	req, err := m.client.NewRequest(ctx, http.MethodPut, loadBalancerServiceName, m.resourcePath(poolID), &members)
	if err != nil {
		return err
	}
	resp, err := m.client.Do(ctx, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

// Delete - Delete a member
func (m *cloudLoadBalancerMemberResource) Delete(ctx context.Context, poolID, id string) error {
	req, err := m.client.NewRequest(ctx, http.MethodDelete, loadBalancerServiceName, m.itemPath(poolID, id), nil)
	if err != nil {
		return err
	}
	resp, err := m.client.Do(ctx, req)
	if err != nil {
		return err
	}
	_, _ = io.Copy(ioutil.Discard, resp.Body)

	return resp.Body.Close()
}

// Create - Create a new member
func (m *cloudLoadBalancerMemberResource) Create(ctx context.Context, poolID string, mcr *CloudLoadBalancerMemberCreateRequest) (*CloudLoadBalancerMember, error) {
	var data struct {
		CloudLoadBalancerMember *CloudLoadBalancerMemberCreateRequest `json:"member"`
	}
	data.CloudLoadBalancerMember = mcr
	req, err := m.client.NewRequest(ctx, http.MethodPost, loadBalancerServiceName, m.resourcePath(poolID), &data)
	if err != nil {
		return nil, err
	}
	resp, err := m.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var response struct {
		CloudLoadBalancerMember *CloudLoadBalancerMember `json:"member"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response.CloudLoadBalancerMember, nil
}
