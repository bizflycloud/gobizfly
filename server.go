package gobizfly

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	serverBasePath = "/iaas-cloud/api/servers"
)

var _ ServerService = (*server)(nil)

// ServerSecurityGroup contains information of security group of a server.
type ServerSecurityGroup struct {
	Name string `json:"name"`
}

// AttachedVolume contains attached volumes of a server.
type AttachedVolume struct {
	ID string `json:"id"`
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
	Addresses        map[string]interface{} `json:"addresses"`
	Metadata         map[string]string      `json:"metadata"`
	Flavor           map[string]interface{} `json:"flavor"`
	Progress         int                    `json:"progress"`
	AttachedVolumes  []AttachedVolume       `json:"os-extended-volumes:volumes_attached"`
	AvailabilityZone string                 `json:"OS-EXT-AZ:availability_zone"`
}

type server struct {
	client *Client
}

// ServerService is an interface to interact with BizFly Cloud API.
type ServerService interface {
	List(ctx context.Context, opts *ListOptions) ([]*Server, error)
	Create(ctx context.Context, scr *ServerCreateRequest) (*ServerTask, error)
	Get(ctx context.Context, id string) (*Server, error)
	Delete(ctx context.Context, id string) error
	Resize(ctx context.Context, id string, newFlavor string) (*ServerTask, error)
	Start(ctx context.Context, id string) (*Server, error)
	Stop(ctx context.Context, id string) (*Server, error)
	SoftReboot(ctx context.Context, id string) (*ServerMessageResponse, error)
	HardReboot(ctx context.Context, id string) (*ServerMessageResponse, error)
	Rebuild(ctx context.Context, id string, imageID string) (*ServerTask, error)
	GetVNC(ctx context.Context, id string) (*ServerConsoleResponse, error)
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

// ServerAction represents server action request payload.
type ServerAction struct {
	Action      string   `json:"action"`
	ImageID     string   `json:"image,omitempty"`
	FlavorName  string   `json:"flavor_name,omitempty"`
	ConsoleType string   `json:"type,omitempty"`
	FirewallIDs []string `json:"firewall_ids,omitempty"`
	NewType     string   `json:"new_type,omitempty"`
}

// ServerTask contains task information.
type ServerTask struct {
	TaskID string `json:"task_id"`
}

// ServerDisk contains server's disk information.
type ServerDisk struct {
	Size int    `json:"size"`
	Type string `json:"type"`
}

// ServerOS contains OS information of server.
type ServerOS struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// ServerCreateRequest represents create a new server payload.
type ServerCreateRequest struct {
	Name             string        `json:"name"`
	FlavorName       string        `json:"flavor"`
	SSHKey           string        `json:"sshkey,omitempty"`
	Password         bool          `json:"password"`
	RootDisk         *ServerDisk   `json:"rootdisk"`
	DataDisks        []*ServerDisk `json:"datadisks,omitempty"`
	Type             string        `json:"type"`
	AvailabilityZone string        `json:"availability_zone"`
	OS               *ServerOS     `json:"os"`
}

// itemActionPath return http path of server action
func (s *server) itemActionPath(id string) string {
	return strings.Join([]string{serverBasePath, id, "action"}, "/")
}

// List lists all servers.
func (s *server) List(ctx context.Context, opts *ListOptions) ([]*Server, error) {

	req, err := s.client.NewRequest(ctx, http.MethodGet, serverBasePath, nil)
	if err != nil {
		return nil, err
	}

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
func (s *server) Create(ctx context.Context, scr *ServerCreateRequest) (*ServerTask, error) {
	payload := []*ServerCreateRequest{scr}
	req, err := s.client.NewRequest(ctx, http.MethodPost, serverBasePath, payload)
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

// Get gets a server.
func (s *server) Get(ctx context.Context, id string) (*Server, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, serverBasePath+"/"+id, nil)
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
func (s *server) Delete(ctx context.Context, id string) error {
	req, err := s.client.NewRequest(ctx, http.MethodDelete, serverBasePath+"/"+id, nil)
	if err != nil {
		return err
	}
	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return err
	}
	_, _ = io.Copy(ioutil.Discard, resp.Body)
	return resp.Body.Close()
}

// Resize resizes a server.
func (s *server) Resize(ctx context.Context, id string, newFlavor string) (*ServerTask, error) {
	var payload = &ServerAction{
		Action:     "resize",
		FlavorName: newFlavor}
	req, err := s.client.NewRequest(ctx, http.MethodPost, s.itemActionPath(id), payload)
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
	req, err := s.client.NewRequest(ctx, http.MethodPost, s.itemActionPath(id), payload)
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
	req, err := s.client.NewRequest(ctx, http.MethodPost, s.itemActionPath(id), payload)
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
	req, err := s.client.NewRequest(ctx, http.MethodPost, s.itemActionPath(id), payload)
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
	req, err := s.client.NewRequest(ctx, http.MethodPost, s.itemActionPath(id), payload)
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
	req, err := s.client.NewRequest(ctx, http.MethodPost, s.itemActionPath(id), payload)
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
	req, err := s.client.NewRequest(ctx, http.MethodPost, s.itemActionPath(id), payload)
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
