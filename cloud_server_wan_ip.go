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

var _ CloudServerPublicNetworkInterfaceService = (*cloudServerPublicNetworkInterfaceResource)(nil)

type cloudServerPublicNetworkInterfaceResource struct {
	client *Client
}

func (cs *cloudServerService) PublicNetworkInterfaces() *cloudServerPublicNetworkInterfaceResource {
	return &cloudServerPublicNetworkInterfaceResource{client: cs.client}
}

type CloudServerPublicNetworkInterfaceService interface {
	Create(ctx context.Context, payload *CreatePublicNetworkInterfacePayload) (*CloudServerPublicNetworkInterface, error)
	List(ctx context.Context) ([]*CloudServerPublicNetworkInterface, error)
	Get(ctx context.Context, wanIPID string) (*CloudServerPublicNetworkInterface, error)
	Delete(ctx context.Context, publicNetworkInterfaceID string) error
	Action(ctx context.Context, publicNetworkInterfaceID string, payload *ActionPublicNetworkInterfacePayload) error
}

type CloudServerPublicNetworkInterface struct {
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
	IPAddress           string               `json:"ip_address"`
	IpVersion           int                  `json:"ip_version"`
}

type CreatePublicNetworkInterfacePayload struct {
	Name             string `json:"name"`
	AttachedServer   string `json:"attached_server"`
	AvailabilityZone string `json:"availability_zone"`
}

type ActionPublicNetworkInterfacePayload struct {
	Action   string `json:"action"`
	ServerID string `json:"server_id,omitempty"`
}

func (w cloudServerPublicNetworkInterfaceResource) resourcePath() string {
	return wanIpPath
}

func (w cloudServerPublicNetworkInterfaceResource) itemPath(id string) string {
	return strings.Join([]string{wanIpPath, id}, "/")
}

func (w cloudServerPublicNetworkInterfaceResource) actionPath(id string) string {
	return strings.Join([]string{wanIpPath, id, "action"}, "/")
}

func (w cloudServerPublicNetworkInterfaceResource) Create(ctx context.Context, payload *CreatePublicNetworkInterfacePayload) (*CloudServerPublicNetworkInterface, error) {
	req, err := w.client.NewRequest(ctx, http.MethodPost, serverServiceName, w.resourcePath(), payload)
	if err != nil {
		return nil, err
	}
	var wanIP *CloudServerPublicNetworkInterface
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

func (w cloudServerPublicNetworkInterfaceResource) List(ctx context.Context) ([]*CloudServerPublicNetworkInterface, error) {
	req, err := w.client.NewRequest(ctx, http.MethodGet, serverServiceName, w.resourcePath(), nil)
	if err != nil {
		return nil, err
	}

	params := req.URL.Query()
	params.Add("detailed", "true")
	req.URL.RawQuery = params.Encode()

	var wanIps []*CloudServerPublicNetworkInterface
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

func (w cloudServerPublicNetworkInterfaceResource) Get(ctx context.Context, id string) (*CloudServerPublicNetworkInterface, error) {
	req, err := w.client.NewRequest(ctx, http.MethodGet, serverServiceName, w.itemPath(id), nil)
	if err != nil {
		return nil, err
	}
	var wanIp *CloudServerPublicNetworkInterface
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

func (w cloudServerPublicNetworkInterfaceResource) Delete(ctx context.Context, id string) error {
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

func (w cloudServerPublicNetworkInterfaceResource) Action(ctx context.Context, id string, payload *ActionPublicNetworkInterfacePayload) error {
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
