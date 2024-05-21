// This file is part of gobizfly

package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

const (
	vpcPath = "/vpc-networks"
)

var _ VPCNetworkService = (*cloudServerVPCNetworkResource)(nil)

type cloudServerVPCNetworkResource struct {
	client *Client
}

func (cs *cloudServerService) VPCNetworks() *cloudServerVPCNetworkResource {
	return &cloudServerVPCNetworkResource{client: cs.client}
}

type VPCNetworkService interface {
	List(ctx context.Context) ([]*VPCNetwork, error)
	Get(ctx context.Context, vpcID string) (*VPCNetwork, error)
	Update(ctx context.Context, vpcID string, uvpl *UpdateVPCPayload) (*VPCNetwork, error)
	Create(ctx context.Context, cvpl *CreateVPCPayload) (*VPCNetwork, error)
	Delete(ctx context.Context, vpcID string) error
}

type VPCNetwork struct {
	ID                    string   `json:"id"`
	Name                  string   `json:"name"`
	TenantID              string   `json:"tenant_id"`
	AdminStateUp          bool     `json:"admin_state_up"`
	MTU                   int      `json:"mtu"`
	Status                string   `json:"status"`
	Subnets               []Subnet `json:"subnets"`
	Shared                bool     `json:"shared"`
	AvailabilityZoneHints []string `json:"availability_zone_hints"`
	AvailabilityZones     []string `json:"availability_zones"`
	IPv4AddressScope      string   `json:"ipv4_address_scope"`
	IPv6AddressScope      string   `json:"ipv6_address_scope"`
	RouterExternal        bool     `json:"router:external"`
	Description           string   `json:"description"`
	PortSecurityEnabled   bool     `json:"port_security_enabled"`
	QosPolicyID           string   `json:"qos_policy_id"`
	Tags                  []string `json:"tags"`
	CreatedAt             string   `json:"created_at"`
	UpdatedAt             string   `json:"updated_at"`
	RevisionNumber        int      `json:"revision_number"`
	IsDefault             bool     `json:"is_default"`
}

type Subnet struct {
	ID              string           `json:"id"`
	Name            string           `json:"name"`
	TenantID        string           `json:"tenant_id"`
	NetworkID       string           `json:"network_id"`
	IPVersion       int              `json:"ip_version"`
	SubnetPoolID    string           `json:"subnet_pool_id"`
	EnableDHCP      bool             `json:"enable_dhcp"`
	IPv6RaMode      string           `json:"ipv6_ra_mode"`
	IPv6AddressMode string           `json:"ipv6_address_mode"`
	GatewayIP       string           `json:"gateway_ip"`
	CIDR            string           `json:"cidr"`
	AllocationPools []AllocationPool `json:"allocation_pools"`
	HostRoutes      []HostRoute      `json:"host_routes"`
	DNSNameServers  []string         `json:"dns_nameservers"`
	Description     string           `json:"description"`
	ServiceTypes    []string         `json:"service_types"`
	Tags            []string         `json:"tags"`
	CreatedAt       string           `json:"created_at"`
	UpdatedAt       string           `json:"updated_at"`
	RevisionNumber  int              `json:"revision_number"`
	ProjectID       string           `json:"project_id"`
}

type HostRoute struct {
	Destination string `json:"destination"`
	NextHop     string `json:"nexthop"`
}

type AllocationPool struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type CreateVPCPayload struct {
	Name        string `json:"name"`
	CIDR        string `json:"cidr,omitempty"`
	Description string `json:"description,omitempty"`
	IsDefault   bool   `json:"is_default,omitempty"`
}

type UpdateVPCPayload struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	CIDR        string `json:"cidr,omitempty"`
	IsDefault   bool   `json:"is_default,omitempty"`
}

func (v cloudServerVPCNetworkResource) resourcePath() string {
	return vpcPath
}

func (v cloudServerVPCNetworkResource) itemPath(id string) string {
	return strings.Join([]string{vpcPath, id}, "/")
}

func (v cloudServerVPCNetworkResource) List(ctx context.Context) ([]*VPCNetwork, error) {
	req, err := v.client.NewRequest(ctx, http.MethodGet, serverServiceName, v.resourcePath(), nil)
	if err != nil {
		return nil, err
	}
	var data []*VPCNetwork
	resp, err := v.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func (v cloudServerVPCNetworkResource) Get(ctx context.Context, vpcID string) (*VPCNetwork, error) {
	req, err := v.client.NewRequest(ctx, http.MethodGet, serverServiceName, v.itemPath(vpcID), nil)
	if err != nil {
		return nil, err
	}
	var data *VPCNetwork
	resp, err := v.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func (v cloudServerVPCNetworkResource) Update(ctx context.Context, vpcID string, uvpl *UpdateVPCPayload) (*VPCNetwork, error) {
	req, err := v.client.NewRequest(ctx, http.MethodPut, serverServiceName, v.itemPath(vpcID), uvpl)
	if err != nil {
		return nil, err
	}
	var data *VPCNetwork
	resp, err := v.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func (v cloudServerVPCNetworkResource) Create(ctx context.Context, cvpl *CreateVPCPayload) (*VPCNetwork, error) {
	req, err := v.client.NewRequest(ctx, http.MethodPost, serverServiceName, v.resourcePath(), cvpl)
	if err != nil {
		return nil, err
	}
	var data *VPCNetwork
	resp, err := v.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func (v cloudServerVPCNetworkResource) Delete(ctx context.Context, vpcID string) error {
	req, err := v.client.NewRequest(ctx, http.MethodDelete, serverServiceName, v.itemPath(vpcID), nil)
	if err != nil {
		return err
	}
	resp, err := v.client.Do(ctx, req)
	if err != nil {
		return err
	}
	return resp.Body.Close()
}
