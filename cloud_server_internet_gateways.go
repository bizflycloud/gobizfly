package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

const igwPath = "/internet-gateways"

type CloudServerInternetGatewayInterface interface {
	Create(ctx context.Context, payload CreateInternetGatewayPayload) (*ExtendedInternetGateway, error)
	List(ctx context.Context, opts ListInternetGatewayOpts) (*ListInternetGatewaysResult, error)
	Get(ctx context.Context, internetGatewayID string) (*ExtendedInternetGateway, error)
	Update(ctx context.Context, internetGatewayID string, payload UpdateInternetGatewayPayload) (*ExtendedInternetGateway, error)
	Delete(ctx context.Context, internetGatewayID string) error
}

type cloudServerInternetGatewayResource struct {
	client *Client
}

func (cs *cloudServerService) InternetGateways() CloudServerInternetGatewayInterface {
	return &cloudServerInternetGatewayResource{client: cs.client}
}

// result
type ExternalFixedIP struct {
	SubnetID  string `json:"subnet_id"`
	IpAddress string `json:"ip_address"`
}

type ExternalGatewayInfo struct {
	NetworkID        string            `json:"network_id"`
	EnableSnat       bool              `json:"enable_snat"`
	ExternalFixedIPs []ExternalFixedIP `json:"external_fixed_ips"`
}

type InternetGateway struct {
	ID                    string               `json:"id"`
	Name                  string               `json:"name"`
	Description           string               `json:"description"`
	TenantID              string               `json:"tenant_id"`
	ProjectID             string               `json:"project_id"`
	Status                string               `json:"status"`
	AvailabilityZones     []string             `json:"availability_zones"`
	AvailabilityZoneHints []string             `json:"availability_zone_hints"`
	FlavorID              *string              `json:"flavor_id"`
	Tags                  []string             `json:"tags"`
	CreatedAt             string               `json:"created_at"`
	UpdatedAt             string               `json:"updated_at"`
	ExTernalGatewayInfo   *ExternalGatewayInfo `json:"external_gateway_info"`
}

type RouterInterfaceNetwork struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type InterfaceInfo struct {
	SubnetID    string                  `json:"subnet_id"`
	PortID      string                  `json:"port_id"`
	IPAddress   string                  `json:"ip_address"`
	NetworkID   string                  `json:"network_id"`
	NetworkInfo *RouterInterfaceNetwork `json:"network_info"`
}

type ExtendedInternetGateway struct {
	InternetGateway `json:",inline"`
	InterfacesInfo  []InterfaceInfo `json:"interfaces_info"`
}

func (igw *cloudServerInternetGatewayResource) parsePath(igwID string) string {
	return strings.Join([]string{igwPath, igwID}, "/")
}

// create internet gateway
type CreateInternetGatewayPayload struct {
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	NetworkIDs  *[]string `json:"network_ids,omitempty"`
}

func (igw *cloudServerInternetGatewayResource) Create(ctx context.Context, payload CreateInternetGatewayPayload) (*ExtendedInternetGateway, error) {
	req, err := igw.client.NewRequest(ctx, http.MethodPost, serverServiceName, igwPath, &payload)
	if err != nil {
		return nil, err
	}
	resp, err := igw.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var igwData *ExtendedInternetGateway
	if err := json.NewDecoder(resp.Body).Decode(&igwData); err != nil {
		return nil, err
	}
	return igwData, nil
}

// list internet gateway
type ListInternetGatewayOpts struct {
	Name       *string `json:"name,omitempty"`
	Detailed   *bool   `json:"detailed,omitempty"`
	Limit      *int    `json:"limit,omitempty"`
	NextCursor *string `json:"next_cursor,omitempty"`
	PrevCursor *string `json:"prev_cursor,omitempty"`
}

type ListInternetGatewaysResult struct {
	InternetGateways []*ExtendedInternetGateway `json:"internet_gateways"`
	Meta             struct {
		NextPage *string `json:"next_page"`
		PrevPage *string `json:"prev_page"`
	} `json:"meta"`
}

func (igw *cloudServerInternetGatewayResource) List(ctx context.Context, opts ListInternetGatewayOpts) (*ListInternetGatewaysResult, error) {
	req, err := igw.client.NewRequest(ctx, http.MethodGet, serverServiceName, igwPath, nil)
	if err != nil {
		return nil, err
	}
	params := req.URL.Query()
	if opts.Name != nil {
		params.Add("name", *opts.Name)
	}
	if opts.Detailed != nil {
		detailedStr := strconv.FormatBool(*opts.Detailed)
		params.Add("detailed", detailedStr)
	}
	if opts.Limit != nil {
		limitStr := strconv.Itoa(*opts.Limit)
		params.Add("limit", limitStr)
	}
	if opts.NextCursor != nil {
		params.Add("next_cursor", *opts.NextCursor)
	}
	if opts.PrevCursor != nil {
		params.Add("prev_cursor", *opts.PrevCursor)
	}
	req.URL.RawQuery = params.Encode()
	resp, err := igw.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var data *ListInternetGatewaysResult
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

// get internet gateway
func (igw *cloudServerInternetGatewayResource) Get(ctx context.Context, internetGatewayID string) (*ExtendedInternetGateway, error) {
	getPath := igw.parsePath(internetGatewayID)
	req, err := igw.client.NewRequest(ctx, http.MethodGet, serverServiceName, getPath, nil)
	if err != nil {
		return nil, err
	}
	resp, err := igw.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var data *ExtendedInternetGateway
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

// update internet gateway
type UpdateInternetGatewayPayload struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	NetworkIDs  []string `json:"network_ids"`
}

func (igw *cloudServerInternetGatewayResource) Update(ctx context.Context, internetGatewayID string, payload UpdateInternetGatewayPayload) (*ExtendedInternetGateway, error) {
	getPath := igw.parsePath(internetGatewayID)
	req, err := igw.client.NewRequest(ctx, http.MethodPut, serverServiceName, getPath, payload)
	if err != nil {
		return nil, err
	}
	resp, err := igw.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var data *ExtendedInternetGateway
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

// delete internet gateway
func (igw *cloudServerInternetGatewayResource) Delete(ctx context.Context, internetGatewayID string) error {
	getPath := igw.parsePath(internetGatewayID)
	req, err := igw.client.NewRequest(ctx, http.MethodDelete, serverServiceName, getPath, nil)
	if err != nil {
		return err
	}
	resp, err := igw.client.Do(ctx, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
