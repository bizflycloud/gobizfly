package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

const (
	wanIpPath = "/wanips"
)

var _ WanIPService = (*wanIPService)(nil)

type wanIPService struct {
	client *Client
}

type WanIPService interface {
	Create(ctx context.Context, payload *CreateWanIpPayload) (*WanIP, error)
	List(ctx context.Context) ([]*WanIP, error)
	Get(ctx context.Context, wanIPId string) (*WanIP, error)
	Delete(ctx context.Context, wanIpId string) error
	Action(ctx context.Context, wanIpId string, payload *ActionWanIpPayload) error
}

type WanIP struct {
	ID                  string               `json:"id"`
	Name                string               `json:"name"`
	NetworkID           string               `json:"network_id"`
	TenantID            string               `json:"tenant_id"`
	MacAddress          string               `json:"mac_address"`
	AdminStateUp        bool                 `json:"admin_state_up"`
	Status              string               `json:"status"`
	DeviceID            string               `json:"device_id"`
	DeviceOwner         string               `json:"device_owner"`
	FixedIps            []FixedIp            `json:"fixed_ips"`
	AllowedAddressPairs []AllowedAddressPair `json:"allowed_address_pairs"`
	ExtraDhcpOpts       []string             `json:"extra_dhcp_opts"`
	SecurityGroups      []string             `json:"security_groups"`
	Description         string               `json:"description"`
	BindingVnicType     string               `json:"binding:vnic_type"`
	PortSecurityEnabled bool                 `json:"port_security_enabled"`
	QosPolicyID         string               `json:"qos_policy_id"`
	Tags                []string             `json:"tags"`
	CreatedAt           string               `json:"created_at"`
	UpdatedAt           string               `json:"updated_at"`
	RevisionNumber      int                  `json:"revision_number"`
	ProjectID           string               `json:"project_id"`
	BillingType         string               `json:"billing_type"`
	Bandwidth           int                  `json:"bandwidth"`
	AvailabilityZone    string               `json:"availability_zone"`
	IsMain              bool                 `json:"is_main"`
	AttachedServer      Server               `json:"attached_server"`
	IpAddress           string               `json:"ip_address"`
	IpVersion           int                  `json:"ip_version"`
}

type CreateWanIpPayload struct {
	Name             string `json:"name"`
	AttachedServer   string `json:"attached_server"`
	AvailabilityZone string `json:"availability_zone"`
}

type ActionWanIpPayload struct {
	Action   string `json:"action"`
	ServerId string `json:"server_id,omitempty"`
}

func (w wanIPService) resourcePath() string {
	return wanIpPath
}

func (w wanIPService) itemPath(id string) string {
	return strings.Join([]string{wanIpPath, id}, "/")
}

func (w wanIPService) actionPath(id string) string {
	return strings.Join([]string{wanIpPath, id, "action"}, "/")
}

func (w wanIPService) Create(ctx context.Context, payload *CreateWanIpPayload) (*WanIP, error) {
	req, err := w.client.NewRequest(ctx, http.MethodPost, serverServiceName, w.resourcePath(), payload)
	if err != nil {
		return nil, err
	}
	var wanIP *WanIP
	resp, err := w.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&wanIP); err != nil {
		return nil, err
	}
	return wanIP, nil
}

func (w wanIPService) List(ctx context.Context) ([]*WanIP, error) {
	req, err := w.client.NewRequest(ctx, http.MethodGet, serverServiceName, w.resourcePath(), nil)
	if err != nil {
		return nil, err
	}
	var wanIps []*WanIP
	resp, err := w.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&wanIps); err != nil {
		return nil, err
	}
	return wanIps, nil
}

func (w wanIPService) Get(ctx context.Context, id string) (*WanIP, error) {
	req, err := w.client.NewRequest(ctx, http.MethodGet, serverServiceName, w.itemPath(id), nil)
	if err != nil {
		return nil, err
	}
	var wanIp *WanIP
	resp, err := w.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&wanIp); err != nil {
		return nil, err
	}
	return wanIp, nil
}

func (w wanIPService) Delete(ctx context.Context, id string) error {
	req, err := w.client.NewRequest(ctx, http.MethodDelete, serverServiceName, w.itemPath(id), nil)
	if err != nil {
		return err
	}
	resp, err := w.client.Do(ctx, req)
	if err != nil {
		return err
	}
	return resp.Body.Close()
}

func (w wanIPService) Action(ctx context.Context, id string, payload *ActionWanIpPayload) error {
	req, err := w.client.NewRequest(ctx, http.MethodPost, serverServiceName, w.actionPath(id), payload)
	if err != nil {
		return err
	}
	resp, err := w.client.Do(ctx, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return err
}
