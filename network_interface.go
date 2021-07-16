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
	NetworkID string `json:"network_id"`
	Status    string `json:"status"`
}

type networkInterfaceService struct {
	client *Client
}

type NetworkInterfaceService interface {
	CreateNetworkInterface(ctx context.Context, networkID string, payload *NetworkInterfaceRequestPayload) (*NetworkInterface, error)
	UpdateNetworkInterface(ctx context.Context, networkID string, netInterfaceID string, payload *NetworkInterfaceRequestPayload) (*NetworkInterface, error)
	DeleteNetworkInterface(ctx context.Context, networkID string, netInterfaceID string) error
	GetNetworkInterface(ctx context.Context, networkID string, netInterfaceID string) (*NetworkInterface, error)
	ListNetworkInterface(ctx context.Context, opts *ListNetworkInterfaceOptions) ([]*NetworkInterface, error)
	ActionVPCNetworkInterface(ctx context.Context, networkID string, netInterfaceID string, payload *NetworkInterfaceRequestPayload) (*NetworkInterface, error)
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
	Firewalls []string `json:"firewalls"`
}

type FixedIp struct {
	SubnetID  string `json:"subnet_id"`
	IPAddress string `json:"ip_address"`
}

type NetworkInterfaceRequestPayload struct {
	AttachedServer string `json:"attached_server,omitempty"`
	FixedIP        string `json:"fixed_ip,omitempty"`
	ServerID       string `json:"server_id,omitempty"`
	Name           string `json:"name"`
	Action         string `json:"action"`
}

func (n networkInterfaceService) resourceCreateNetworkInterfacePath(networkID string) string {
	return strings.Join([]string{vpcPath, networkID, "network-interfaces"}, "/")
}

func (n networkInterfaceService) resourceNetworkInterfacePath(networkID string, netInterfaceID string) string {
	return strings.Join([]string{vpcPath, networkID, "network-interfaces", netInterfaceID}, "/")
}

func (n networkInterfaceService) CreateNetworkInterface(ctx context.Context, networkID string, payload *NetworkInterfaceRequestPayload) (*NetworkInterface, error) {
	req, err := n.client.NewRequest(ctx, http.MethodPost, serverServiceName, n.resourceCreateNetworkInterfacePath(networkID), payload)
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

func (n networkInterfaceService) GetNetworkInterface(ctx context.Context, networkID string, netInterfaceID string) (*NetworkInterface, error) {
	req, err := n.client.NewRequest(ctx, http.MethodGet, serverServiceName, n.resourceNetworkInterfacePath(networkID, netInterfaceID), nil)
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

func (n networkInterfaceService) UpdateNetworkInterface(ctx context.Context, networkID string, netInterfaceID string, payload *NetworkInterfaceRequestPayload) (*NetworkInterface, error) {
	req, err := n.client.NewRequest(ctx, http.MethodPut, serverServiceName, n.resourceNetworkInterfacePath(networkID, netInterfaceID), payload)
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

func (n networkInterfaceService) DeleteNetworkInterface(ctx context.Context, networkID string, netInterfaceID string) error {
	req, err := n.client.NewRequest(ctx, http.MethodDelete, serverServiceName, n.resourceNetworkInterfacePath(networkID, netInterfaceID), nil)
	if err != nil {
		return err
	}
	resp, err := n.client.Do(ctx, req)
	if err != nil {
		return err
	}
	return resp.Body.Close()
}

func (n networkInterfaceService) ListNetworkInterface(ctx context.Context, opts *ListNetworkInterfaceOptions) ([]*NetworkInterface, error) {
	req, err := n.client.NewRequest(ctx, http.MethodGet, serverServiceName, networkInterfacePath, nil)
	if err != nil {
		return nil, err
	}

	params := req.URL.Query()
	params.Add("network_id", opts.NetworkID)
	params.Add("status", opts.Status)
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

func (n networkInterfaceService) ActionVPCNetworkInterface(ctx context.Context, networkID string, netInterfaceID string, payload *NetworkInterfaceRequestPayload) (*NetworkInterface, error) {
	req, err := n.client.NewRequest(ctx, http.MethodPut, serverServiceName, n.resourceNetworkInterfacePath(networkID, netInterfaceID), payload)
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
