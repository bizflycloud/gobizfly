package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

const (
	serverTypeBasePath = "/server-types"
)

type server struct {
	client *Client
}

// ServerSecurityGroup contains information of security group of a server.
type ServerSecurityGroup struct {
	Name string `json:"name"`
}

// ServerTaskResult represents the response of getting server task result
type ServerTaskResult struct {
	Action   string `json:"action"`
	Progress int    `json:"progress"`
	Success  bool   `json:"success"`
	Server
}

// ServerTaskResponse represents the response of getting server task
type ServerTaskResponse struct {
	Ready  bool             `json:"ready"`
	Result ServerTaskResult `json:"result"`
}

// AttachedVolume contains attached volumes of a server.
type AttachedVolume struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Size         int    `json:"size"`
	AttachedType string `json:"attached_type"`
	Type         string `json:"type"`
	Category     string `json:"category"`
}

// IP represents the IP address, version and mac address of a port
type IP struct {
	Version    int    `json:"version"`
	Address    string `json:"addr"`
	Type       string `json:"OS-EXT-IPS:type"`
	MacAddress string `json:"0S-EXT-IPS-MAC:mac_addr"`
}

// IPAddress contains LAN & WAN Ip address of a Cloud Server
type IPAddress struct {
	LanAddresses   []IP `json:"LAN"`
	WanV4Addresses []IP `json:"WAN_V4"`
	WanV6Addresses []IP `json:"WAN_V6"`
}

// Flavor contains flavor information.
type Flavor struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Ram  int    `json:"ram"`
	VCPU int    `json:"vcpu"`
}

// DeletedVolumes represent payload when delete server
type DeletedVolumes struct {
	Ids []string `json:"delete_volume"`
}

// Server contains server information.
type Server struct {
	ID               string                 `json:"id"`
	Name             string                 `json:"name"`
	KeyName          string                 `json:"key_name"`
	UserID           string                 `json:"user_id"`
	ProjectID        string                 `json:"tenant_id"`
	CreatedAt        string                 `json:"created"`
	UpdatedAt        string                 `json:"updated"`
	Status           string                 `json:"status"`
	IPv6             bool                   `json:"ipv6"`
	SecurityGroup    []ServerSecurityGroup  `json:"security_group"`
	Addresses        map[string]interface{} `json:"addresses"` // Deprecated: This field will be removed in the near future
	Metadata         map[string]string      `json:"metadata"`
	Flavor           Flavor                 `json:"flavor"`
	FlavorName       string                 `json:"flavor_name"`
	Progress         int                    `json:"progress"`
	AttachedVolumes  []AttachedVolume       `json:"os-extended-volumes:volumes_attached"`
	AvailabilityZone string                 `json:"OS-EXT-AZ:availability_zone"`
	Category         string                 `json:"category"`
	IPAddresses      IPAddress              `json:"ip_addresses"`
	RegionName       string                 `json:"region_name"`
	NetworkPlan      string                 `json:"network_plan"`
	Locked           bool                   `json:"locked"`
	IsCreatedWan     bool                   `json:"is_created_wan"`
	ZoneName         string                 `json:"zone_name"`
	BillingPlan      string                 `json:"billing_plan"`
	IsAvailable      bool                   `json:"is_available"`
}

// CreateCustomImagePayload represents payload when create custom image
type CreateCustomImagePayload struct {
	Name        string `json:"name"`
	DiskFormat  string `json:"disk_format"`
	Description string `json:"description,omitempty"`
	ImageURL    string `json:"image_url,omitempty"`
}

// Location represents a location of a server.
type Location struct {
	URL      string            `json:"url"`
	Metadata map[string]string `json:"metadata,"`
}

// ServerListOptions represents the options for listing servers.
type ServerListOptions struct {
	detailed bool
	fields   []string
}

// ServerConsoleResponse contains information of server console url.
type ServerConsoleResponse struct {
	URL  string `json:"url"`
	Type string `json:"type"`
}

// ServerMessageResponse contains message response from Cloud Server API.
type ServerMessageResponse struct {
	Message string `json:"message"`
}

// ServerType represents a server type.
type ServerType struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Enabled      bool     `json:"enabled"`
	ComputeClass []string `json:"compute_class"`
	Priority     int      `json:"priority"`
}

// ServerAction represents server action request payload.
type ServerAction struct {
	Action         string   `json:"action"`
	ImageID        string   `json:"image,omitempty"`
	FlavorName     string   `json:"flavor_name,omitempty"`
	ConsoleType    string   `json:"type,omitempty"`
	FirewallIDs    []string `json:"firewall_ids,omitempty"`
	NewType        string   `json:"new_type,omitempty"`
	VPCNetworkIDs  []string `json:"vpc_network_ids,omitempty"`
	AttachWanIPs   []string `json:"wan_ips,omitempty"`
	NewNetworkPlan string   `json:"new_network_plan,omitempty"`
	NewBillingPlan string   `json:"billing_plan,omitempty"`
}

// ServerTask contains task information.
type ServerTask struct {
	TaskID string `json:"task_id"`
}

// ServerCreateResponse contains response tasks when create server
type ServerCreateResponse struct {
	Task []string `json:"task_id"`
}

// ServerDisk contains server's disk information.
type ServerDisk struct {
	Size       int     `json:"size"`
	Type       *string `json:"type,omitempty"` // Deprecated: This field will be removed in the near future
	VolumeType *string `json:"volume_type,omitempty"`
}

// ServerOS contains OS information of server.
type ServerOS struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// ServerCreateRequest represents create a new server payload.
type ServerCreateRequest struct {
	Name                 string        `json:"name"`
	FlavorName           string        `json:"flavor"`
	SSHKey               string        `json:"sshkey,omitempty"`
	Password             bool          `json:"password"`
	RootDisk             *ServerDisk   `json:"rootdisk"`
	DataDisks            []*ServerDisk `json:"datadisks,omitempty"`
	Type                 string        `json:"type"`
	AvailabilityZone     string        `json:"availability_zone"`
	OS                   *ServerOS     `json:"os"`
	Quantity             int           `json:"quantity,omitempty"`
	NetworkInterface     []string      `json:"network_interfaces,omitempty"`
	WanNetworkInterfaces []string      `json:"wan_network_interfaces,omitempty"`
	Firewalls            []string      `json:"firewalls,omitempty"`
	NetworkPlan          string        `json:"network_plan,omitempty"`
	VPCNetworkIds        []string      `json:"vpc_network_ids,omitempty"`
	BillingPlan          string        `json:"billing_plan,omitempty"`
	IPv6                 bool          `json:"ipv6,omitempty"`
	IsCreatedWan         bool          `json:"is_created_wan,omitempty"`
	UserData             string        `json:"user_data,omitempty"`
}

// itemActionPath return http path of server action
func (s *server) itemActionPath(id string) string {
	return strings.Join([]string{serverBasePath, id, "action"}, "/")
}

// List lists all servers.
func (s *server) List(ctx context.Context, opts *ServerListOptions) ([]*Server, error) {

	req, err := s.client.NewRequest(ctx, http.MethodGet, serverServiceName, serverBasePath, nil)
	if err != nil {
		return nil, err
	}
	params := req.URL.Query()
	if opts.detailed {
		params.Add("detailed", "True")
	}
	if len(opts.fields) != 0 {
		params.Add("fields", strings.Join(opts.fields, ","))
	}
	req.URL.RawQuery = params.Encode()

	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var servers []*Server

	if err := json.NewDecoder(resp.Body).Decode(&servers); err != nil {
		return nil, err
	}

	return servers, nil
}

// Create creates a new server.
func (s *server) Create(ctx context.Context, scr *ServerCreateRequest) (*ServerCreateResponse, error) {
	payload := []*ServerCreateRequest{scr}
	req, err := s.client.NewRequest(ctx, http.MethodPost, serverServiceName, serverBasePath, payload)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var task *ServerCreateResponse
	if err := json.NewDecoder(resp.Body).Decode(&task); err != nil {
		return nil, err
	}
	return task, nil
}

// Get gets a server.
func (s *server) Get(ctx context.Context, id string) (*Server, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, serverServiceName, serverBasePath+"/"+id, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var svr *Server
	if err := json.NewDecoder(resp.Body).Decode(&svr); err != nil {
		return nil, err
	}
	return svr, nil
}

// Delete deletes a server.
func (s *server) Delete(ctx context.Context, id string, deletedRootDisk []string) (*ServerTask, error) {
	deletedVolumes := &DeletedVolumes{
		Ids: deletedRootDisk,
	}
	req, err := s.client.NewRequest(ctx, http.MethodDelete, serverServiceName, serverBasePath+"/"+id, deletedVolumes)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var task *ServerTask
	if err := json.NewDecoder(resp.Body).Decode(&task); err != nil {
		return nil, err
	}
	return task, nil
}

// Resize resizes a server.
func (s *server) Resize(ctx context.Context, id string, newFlavor string) (*ServerTask, error) {
	var payload = &ServerAction{
		Action:     "resize",
		FlavorName: newFlavor}
	req, err := s.client.NewRequest(ctx, http.MethodPost, serverServiceName, s.itemActionPath(id), payload)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(ctx, req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var task *ServerTask
	if err := json.NewDecoder(resp.Body).Decode(&task); err != nil {
		return nil, err
	}
	return task, nil
}

// Start starts a server.
func (s *server) Start(ctx context.Context, id string) (*Server, error) {
	payload := &ServerAction{Action: "start"}
	req, err := s.client.NewRequest(ctx, http.MethodPost, serverServiceName, s.itemActionPath(id), payload)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var svr *Server
	if err := json.NewDecoder(resp.Body).Decode(&svr); err != nil {
		return nil, err
	}
	return svr, nil
}

// Stop stops a server
func (s *server) Stop(ctx context.Context, id string) (*Server, error) {
	payload := &ServerAction{Action: "stop"}
	req, err := s.client.NewRequest(ctx, http.MethodPost, serverServiceName, s.itemActionPath(id), payload)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var svr *Server
	if err := json.NewDecoder(resp.Body).Decode(&svr); err != nil {
		return nil, err
	}
	return svr, nil
}

// SoftReboot soft reboots a server.
func (s *server) SoftReboot(ctx context.Context, id string) (*ServerMessageResponse, error) {
	payload := &ServerAction{Action: "soft_reboot"}
	req, err := s.client.NewRequest(ctx, http.MethodPost, serverServiceName, s.itemActionPath(id), payload)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var srm *ServerMessageResponse
	if err := json.NewDecoder(resp.Body).Decode(&srm); err != nil {
		return nil, err
	}
	return srm, nil
}

// HardReboot hard reboots a server.
func (s *server) HardReboot(ctx context.Context, id string) (*ServerMessageResponse, error) {
	payload := &ServerAction{Action: "hard_reboot"}
	req, err := s.client.NewRequest(ctx, http.MethodPost, serverServiceName, s.itemActionPath(id), payload)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var smr *ServerMessageResponse
	if err := json.NewDecoder(resp.Body).Decode(&smr); err != nil {
		return nil, err
	}
	return smr, nil
}

// Rebuild rebuilds a server.
func (s *server) Rebuild(ctx context.Context, id string, imageID string) (*ServerTask, error) {
	var payload = &ServerAction{
		Action:  "rebuild",
		ImageID: imageID}
	req, err := s.client.NewRequest(ctx, http.MethodPost, serverServiceName, s.itemActionPath(id), payload)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(ctx, req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var task *ServerTask
	if err := json.NewDecoder(resp.Body).Decode(&task); err != nil {
		return nil, err
	}
	return task, nil
}

// GetVNC gets vnc console of a server.
func (s *server) GetVNC(ctx context.Context, id string) (*ServerConsoleResponse, error) {
	payload := &ServerAction{
		Action:      "get_vnc",
		ConsoleType: "novnc"}
	req, err := s.client.NewRequest(ctx, http.MethodPost, serverServiceName, s.itemActionPath(id), payload)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var respPayload struct {
		Console *ServerConsoleResponse `json:"console"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&respPayload); err != nil {
		return nil, err
	}
	return respPayload.Console, nil
}

// GetTask get tasks result from Server API
func (s *server) GetTask(ctx context.Context, id string) (*ServerTaskResponse, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, serverServiceName, taskPath+"/"+id, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	var str *ServerTaskResponse
	if err := json.NewDecoder(resp.Body).Decode(&str); err != nil {
		return nil, err
	}
	return str, nil
}

// ChangeCategory changes category of the server
func (s server) ChangeCategory(ctx context.Context, id string, newCategory string) (*ServerTask, error) {
	payload := &ServerAction{
		Action:  "change_type",
		NewType: newCategory}
	req, err := s.client.NewRequest(ctx, http.MethodPost, serverServiceName, s.itemActionPath(id), payload)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var svt *ServerTask
	if err := json.NewDecoder(resp.Body).Decode(&svt); err != nil {
		return nil, err
	}
	return svt, nil
}

// AddVPC add the VPC to the server
func (s server) AddVPC(ctx context.Context, id string, vpcs []string) (*Server, error) {
	payload := &ServerAction{
		Action:        "add_vpc",
		VPCNetworkIDs: vpcs,
	}
	req, err := s.client.NewRequest(ctx, http.MethodPost, serverServiceName, s.itemActionPath(id), payload)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var server *Server
	if err := json.NewDecoder(resp.Body).Decode(&server); err != nil {
		return nil, err
	}
	return server, nil
}

// RemoveVPC remove the VPC from the server
func (s server) RemoveVPC(ctx context.Context, id string, vpcs []string) (*Server, error) {
	payload := &ServerAction{
		Action:        "remove_vpc",
		VPCNetworkIDs: vpcs,
	}
	req, err := s.client.NewRequest(ctx, http.MethodPost, serverServiceName, s.itemActionPath(id), payload)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var server *Server
	if err := json.NewDecoder(resp.Body).Decode(&server); err != nil {
		return nil, err
	}
	return server, nil
}

// AttachWanIps attach batch WAN IPs to the server
func (s server) AttachWanIps(ctx context.Context, id string, wanIpIds []string) error {
	payload := &ServerAction{
		Action:       "attach_wan_ips",
		AttachWanIPs: wanIpIds,
	}
	req, err := s.client.NewRequest(ctx, http.MethodPost, serverServiceName, s.itemActionPath(id), payload)
	if err != nil {
		return err
	}
	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return err
	}
	return resp.Body.Close()
}

func (s server) ListServerTypes(ctx context.Context) ([]*ServerType, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, serverServiceName, serverTypeBasePath, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var serverTypes struct {
		ServerTypes []*ServerType `json:"server_types"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&serverTypes); err != nil {
		return nil, err
	}
	return serverTypes.ServerTypes, nil
}

func (s server) ChangeNetworkPlan(ctx context.Context, id string, newNetworkPlan string) error {
	payload := &ServerAction{
		Action:         "change_network_plan",
		NewNetworkPlan: newNetworkPlan,
	}
	req, err := s.client.NewRequest(ctx, http.MethodPost, serverServiceName, s.itemActionPath(id), payload)
	if err != nil {
		return err
	}
	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return err
	}
	return resp.Body.Close()
}

func (s server) SwitchBillingPlan(ctx context.Context, id string, newBillingPlan string) error {
	payload := &ServerAction{
		Action:         "switch_billing_plan",
		NewBillingPlan: newBillingPlan,
	}
	req, err := s.client.NewRequest(ctx, http.MethodPost, serverServiceName, s.itemActionPath(id), payload)
	if err != nil {
		return err
	}
	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return err
	}
	return resp.Body.Close()
}
