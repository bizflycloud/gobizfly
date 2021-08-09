package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

const (
	networkInterfacePath = "/network-interfaces"
)

var _ NetworkInterfaceService = (*networkInterfaceService)(nil)

type ListNetworkInterfaceOptions struct {
	VPCNetworkID string `json:"vpc-network-id,omitempty"`
	Status       string `json:"status,omitempty"`
	Detailed     string `json:"detailed,omitempty"`
	Type         string `json:"type,omitempty"`
}

type networkInterfaceService struct {
	client *Client
}

type NetworkInterfaceService interface {
	Create(ctx context.Context, networkInterfaceID string, payload *CreateNetworkInterfacePayload) (*NetworkInterface, error)
	Update(ctx context.Context, networkInterfaceID string, payload *UpdateNetworkInterfacePayload) (*NetworkInterface, error)
	Delete(ctx context.Context, networkInterfaceID string) error
	Get(ctx context.Context, networkInterfaceID string) (*NetworkInterface, error)
	Action(ctx context.Context, networkInterfaceID string, payload *ActionNetworkInterfacePayload) (*NetworkInterface, error)
	List(ctx context.Context, opts *ListNetworkInterfaceOptions) ([]*NetworkInterface, error)
}

type NetworkInterface struct {
	ID                  string    `json:"id"`
	Name                string    `json:"name"`
	NetworkID           string    `json:"network_id"`
	TenantID            string    `json:"tenant_id"`
	MacAddress          string    `json:"mac_address"`
	AdminStateUp        bool      `json:"admin_state_up"`
	Status              string    `json:"status"`
	DeviceID            string    `json:"device_id"`
	DeviceOwner         string    `json:"device_owner"`
	FixedIps            []FixedIp `json:"fixed_ips"`
	AllowedAddressPairs []string  `json:"allowed_address_pairs"`
	ExtraDhcpOpts       []string  `json:"extra_dhcp_opts"`
	SecurityGroups      []string  `json:"security_groups"`
	Description         string    `json:"description"`
	BindingVnicType     string    `json:"binding:vnic_type"`
	PortSecurityEnabled bool      `json:"port_security_enabled"`
	QosPolicyID         string    `json:"qos_policy_id"`
	Tags                []string  `json:"tags"`
	CreatedAt           string    `json:"created_at"`
	UpdatedAt           string    `json:"updated_at"`
	RevisionNumber      int       `json:"revision_number"`
	ProjectID           string    `json:"project_id"`
	AttachedServer      struct {
	} `json:"attached_server"`
}

type FixedIp struct {
	SubnetID  string `json:"subnet_id"`
	IPAddress string `json:"ip_address"`
}

type UpdateNetworkInterfacePayload struct {
	Name string `json:"name"`
}

type CreateNetworkInterfacePayload struct {
	AttachedServer string `json:"attached_server,omitempty"`
	FixedIP        string `json:"fixed_ip,omitempty"`
	Name           string `json:"name"`
}

type ActionNetworkInterfacePayload struct {
	Action         string   `json:"action,omitempty"`
	ServerID       string   `json:"server_id,omitempty"`
	SecurityGroups []string `json:"security_groups,omitempty"`
}

func (n networkInterfaceService) createPath(networkID string) string {
	return strings.Join([]string{vpcPath, networkID, "network-interfaces"}, "/")
}

func (n networkInterfaceService) actionPath(id string) string {
	return strings.Join([]string{networkInterfacePath, id, "action"}, "/")
}

func (n networkInterfaceService) resourcePath() string {
	return networkInterfacePath
}

func (n networkInterfaceService) itemPath(id string) string {
	return strings.Join([]string{networkInterfacePath, id}, "/")
}

func (n networkInterfaceService) Create(ctx context.Context, networkID string, payload *CreateNetworkInterfacePayload) (*NetworkInterface, error) {
	req, err := n.client.NewRequest(ctx, http.MethodPost, serverServiceName, n.createPath(networkID), payload)
	if err != nil {
		return nil, err
	}
	var data *NetworkInterface
	resp, err := n.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func (n networkInterfaceService) Get(ctx context.Context, networkInterfaceID string) (*NetworkInterface, error) {
	req, err := n.client.NewRequest(ctx, http.MethodGet, serverServiceName, n.itemPath(networkInterfaceID), nil)
	if err != nil {
		return nil, err
	}
	var data *NetworkInterface
	resp, err := n.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func (n networkInterfaceService) Update(ctx context.Context, networkInterfaceID string, payload *UpdateNetworkInterfacePayload) (*NetworkInterface, error) {
	req, err := n.client.NewRequest(ctx, http.MethodPut, serverServiceName, n.itemPath(networkInterfaceID), payload)
	if err != nil {
		return nil, err
	}
	var data *NetworkInterface
	resp, err := n.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func (n networkInterfaceService) Delete(ctx context.Context, networkInterfaceID string) error {
	req, err := n.client.NewRequest(ctx, http.MethodDelete, serverServiceName, n.itemPath(networkInterfaceID), nil)
	if err != nil {
		return err
	}
	resp, err := n.client.Do(ctx, req)
	if err != nil {
		return err
	}
	return resp.Body.Close()
}

func (n networkInterfaceService) List(ctx context.Context, opts *ListNetworkInterfaceOptions) ([]*NetworkInterface, error) {
	req, err := n.client.NewRequest(ctx, http.MethodGet, serverServiceName, n.resourcePath(), nil)
	if err != nil {
		return nil, err
	}

	params := req.URL.Query()
	params.Add("vpc-network-id", opts.VPCNetworkID)
	params.Add("status", opts.Status)
	params.Add("detailed", opts.Detailed)
	params.Add("type", opts.Type)
	req.URL.RawQuery = params.Encode()

	var data []*NetworkInterface
	resp, err := n.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func (n networkInterfaceService) Action(ctx context.Context, networkInterfaceID string, payload *ActionNetworkInterfacePayload) (*NetworkInterface, error) {
	req, err := n.client.NewRequest(ctx, http.MethodPost, serverServiceName, n.actionPath(networkInterfaceID), payload)
	if err != nil {
		return nil, err
	}
	var data *NetworkInterface
	resp, err := n.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}
