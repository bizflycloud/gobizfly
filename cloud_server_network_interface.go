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

var _ NetworkInterfaceService = (*cloudServerNetworkInterfaceResource)(nil)

type cloudServerNetworkInterfaceResource struct {
	client *Client
}

func (cs *cloudServerService) NetworkInterfaces() *cloudServerNetworkInterfaceResource {
	return &cloudServerNetworkInterfaceResource{client: cs.client}
}

type NetworkInterfaceService interface {
	Create(ctx context.Context, networkInterfaceID string, payload *CreateNetworkInterfacePayload) (*NetworkInterface, error)
	Update(ctx context.Context, networkInterfaceID string, payload *UpdateNetworkInterfacePayload) (*NetworkInterface, error)
	Delete(ctx context.Context, networkInterfaceID string) error
	Get(ctx context.Context, networkInterfaceID string) (*NetworkInterface, error)
	Action(ctx context.Context, networkInterfaceID string, payload *ActionNetworkInterfacePayload) (*NetworkInterface, error)
	List(ctx context.Context, opts *ListNetworkInterfaceOptions) ([]*NetworkInterface, error)
}

// ListNetworkInterfaceOptions represents the options for listing network interfaces.
type ListNetworkInterfaceOptions struct {
	VPCNetworkID string `json:"vpc-network-id,omitempty"`
	Status       string `json:"status,omitempty"`
	Detailed     string `json:"detailed,omitempty"`
	Type         string `json:"type,omitempty"`
}

// NetworkInterface represents a network interface's information.
type NetworkInterface struct {
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
	ExtraDhcpOpts       []ExtraDHCPOpt       `json:"extra_dhcp_opts"`
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
	Type                string               `json:"type"`
	BillingType         string               `json:"billing_type"`
	AttachedServer      struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"attached_server"`
	IPVersion int    `json:"ip_version"`
	IPAddress string `json:"ip_address"`
}

// FixedIp represents a fixed IP's information.
type FixedIp struct {
	SubnetID  string `json:"subnet_id"`
	IPAddress string `json:"ip_address"`
	IPVersion int    `json:"ip_version"`
}

type ExtraDHCPOpt struct {
	OptName   string      `json:"opt_name"`
	OptValue  interface{} `json:"opt_value"`
	IPVersion int         `json:"ip_version,omitempty"`
}

type AllowedAddressPair struct {
	IPAddress  string `json:"ip_address"`
	MacAddress string `json:"mac_address"`
}

// UpdateNetworkInterfacePayload represents the payload for updating a network interface.
type UpdateNetworkInterfacePayload struct {
	Name string `json:"name"`
}

// CreateNetworkInterfacePayload represents the payload for creating a network interface.
type CreateNetworkInterfacePayload struct {
	AttachedServer string `json:"attached_server,omitempty"`
	FixedIP        string `json:"fixed_ip,omitempty"`
	Name           string `json:"name"`
}

// ActionNetworkInterfacePayload represents the payload for action on a network interface
type ActionNetworkInterfacePayload struct {
	Action         string   `json:"action,omitempty"`
	ServerID       string   `json:"server_id,omitempty"`
	SecurityGroups []string `json:"security_groups,omitempty"`
}

func (n cloudServerNetworkInterfaceResource) createPath(networkID string) string {
	return strings.Join([]string{vpcPath, networkID, "network-interfaces"}, "/")
}

func (n cloudServerNetworkInterfaceResource) actionPath(id string) string {
	return strings.Join([]string{networkInterfacePath, id, "action"}, "/")
}

func (n cloudServerNetworkInterfaceResource) resourcePath() string {
	return networkInterfacePath
}

func (n cloudServerNetworkInterfaceResource) itemPath(id string) string {
	return strings.Join([]string{networkInterfacePath, id}, "/")
}

// Create - Create a network interface
func (n cloudServerNetworkInterfaceResource) Create(ctx context.Context, networkID string, payload *CreateNetworkInterfacePayload) (*NetworkInterface, error) {
	req, err := n.client.NewRequest(ctx, http.MethodPost, serverServiceName, n.createPath(networkID), payload)
	if err != nil {
		return nil, err
	}
	var data *NetworkInterface
	resp, err := n.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

// Get - Get a network interface
func (n cloudServerNetworkInterfaceResource) Get(ctx context.Context, networkInterfaceID string) (*NetworkInterface, error) {
	req, err := n.client.NewRequest(ctx, http.MethodGet, serverServiceName, n.itemPath(networkInterfaceID), nil)
	if err != nil {
		return nil, err
	}
	var data *NetworkInterface
	resp, err := n.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

// Update - Update the network interface information
func (n cloudServerNetworkInterfaceResource) Update(ctx context.Context, networkInterfaceID string, payload *UpdateNetworkInterfacePayload) (*NetworkInterface, error) {
	req, err := n.client.NewRequest(ctx, http.MethodPut, serverServiceName, n.itemPath(networkInterfaceID), payload)
	if err != nil {
		return nil, err
	}
	var data *NetworkInterface
	resp, err := n.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

// Delete - Delete the network interface
func (n cloudServerNetworkInterfaceResource) Delete(ctx context.Context, networkInterfaceID string) error {
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

// List - List network interfaces with options
func (n cloudServerNetworkInterfaceResource) List(ctx context.Context, opts *ListNetworkInterfaceOptions) ([]*NetworkInterface, error) {
	req, err := n.client.NewRequest(ctx, http.MethodGet, serverServiceName, n.resourcePath(), nil)
	if err != nil {
		return nil, err
	}

	params := req.URL.Query()
	if opts != nil {
		if opts.VPCNetworkID != "" {
			params.Add("vpc-network-id", opts.VPCNetworkID)
		}
		if opts.Status != "" {
			params.Add("status", opts.Status)
		}
		if opts.Detailed != "" {
			params.Add("detailed", opts.Detailed)
		}
		if opts.Type != "" {
			params.Add("type", opts.Type)
		}
	}
	req.URL.RawQuery = params.Encode()

	var data []*NetworkInterface
	resp, err := n.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

// Action - Execute action on a network interface
func (n cloudServerNetworkInterfaceResource) Action(ctx context.Context, networkInterfaceID string, payload *ActionNetworkInterfacePayload) (*NetworkInterface, error) {
	req, err := n.client.NewRequest(ctx, http.MethodPost, serverServiceName, n.actionPath(networkInterfaceID), payload)
	if err != nil {
		return nil, err
	}
	var data *NetworkInterface
	resp, err := n.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}
