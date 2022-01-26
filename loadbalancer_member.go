package gobizfly

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

var _ MemberService = (*member)(nil)

// MemberService is an interface to interact with Bizfly API Members endpoint.
type MemberService interface {
	List(ctx context.Context, poolID string, opts *ListOptions) ([]*Member, error)
	Get(ctx context.Context, poolID, id string) (*Member, error)
	Update(ctx context.Context, poolID, id string, req *MemberUpdateRequest) (*Member, error)
	Delete(ctx context.Context, poolID, id string) error
	Create(ctx context.Context, poolID string, req *MemberCreateRequest) (*Member, error)
	BatchUpdate(ctx context.Context, poolID string, members *BatchMemberUpdateRequest) error
}

// MemberUpdateRequest represents update member request payload.
type MemberUpdateRequest struct {
	Name           string `json:"name"`
	Weight         int    `json:"weight,omitempty"`
	MonitorAddress string `json:"monitor_address,omitempty"`
	MonitorPort    int    `json:"monitor_port,omitempty"`
	Backup         bool   `json:"backup,omitempty"`
}

// ExtendMemberUpdateRequest represents update member request payload.
type ExtendMemberUpdateRequest struct {
	MemberUpdateRequest
	Address      string `json:"address"`
	ProtocolPort int    `json:"protocol_port"`
}

// BatchMemberUpdateRequest represents batch update member request payload.
type BatchMemberUpdateRequest struct {
	Members []ExtendMemberUpdateRequest `json:"members"`
}

// MemberCreateRequest represents create member request payload
type MemberCreateRequest struct {
	Name           string `json:"name"`
	Weight         int    `json:"weight,omitempty"`
	Address        string `json:"address"`
	ProtocolPort   int    `json:"protocol_port"`
	MonitorAddress string `json:"monitor_address,omitempty"`
	MonitorPort    int    `json:"monitor_port,omitempty"`
	Backup         bool   `json:"backup,omitempty"`
}

// Member contains member information.
type Member struct {
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

type member struct {
	client *Client
}

func (m *member) resourcePath(poolID string) string {
	return strings.Join([]string{poolPath, poolID, "member"}, "/")
}

func (m *member) itemPath(poolID string, id string) string {
	return strings.Join([]string{poolPath, poolID, "member", id}, "/")
}

// List - list members' information
func (m *member) List(ctx context.Context, poolID string, opts *ListOptions) ([]*Member, error) {
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
		Members []*Member `json:"members"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Members, nil
}

// Get - Get member's information
func (m *member) Get(ctx context.Context, poolID, id string) (*Member, error) {
	req, err := m.client.NewRequest(ctx, http.MethodGet, loadBalancerServiceName, m.itemPath(poolID, id), nil)
	if err != nil {
		return nil, err
	}
	resp, err := m.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	mb := &Member{}
	if err := json.NewDecoder(resp.Body).Decode(mb); err != nil {
		return nil, err
	}
	return mb, nil
}

// Update - Update member's information
func (m *member) Update(ctx context.Context, poolID, id string, mur *MemberUpdateRequest) (*Member, error) {
	var data struct {
		Member *MemberUpdateRequest `json:"member"`
	}
	data.Member = mur
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
		Member *Member `json:"member"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, err
	}
	return respData.Member, nil
}

// BatchUpdate - Batch update members
func (m *member) BatchUpdate(ctx context.Context, poolID string, members *BatchMemberUpdateRequest) error {
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
func (m *member) Delete(ctx context.Context, poolID, id string) error {
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
func (m *member) Create(ctx context.Context, poolID string, mcr *MemberCreateRequest) (*Member, error) {
	var data struct {
		Member *MemberCreateRequest `json:"member"`
	}
	data.Member = mcr
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
		Member *Member `json:"member"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response.Member, nil
}
