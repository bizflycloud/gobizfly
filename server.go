package gobizfly

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	// "strings"
	"encoding/json"
)

const (
	serverBasePath = "/iaas-cloud/api/servers"
)

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
	ID              string                 `json:"id"`
	Name            string                 `json:"name"`
	KeyName         string                 `json:"key_name"`
	UserID          string                 `json:"user_id"`
	ProjectID       string                 `json:"tenant_id"`
	CreatedAt       string                 `json:"created"`
	UpdatedAt       string                 `json:"updated"`
	Status          string                 `json:"status"`
	IPv6            bool                   `json:"ipv6"`
	SecurityGroup   []ServerSecurityGroup  `json:"security_group"`
	Addresses       map[string]interface{} `json:"addresses"`
	Metadata        map[string]string      `json:"metadata"`
	Flavor          map[string]interface{} `json:"flavor"`
	Progress        int                    `json:"progress"`
	AttachedVolumes []AttachedVolume       `json:"os-extended-volumes:volumes_attached"`
}

// ServerCreateRequest represents create a new server payload.
type ServerCreateRequest struct {
}

type server struct {
	client *Client
}

// ServerService is an interface to interact with BizFly Cloud API.
type ServerService interface {
	List(ctx context.Context, opts *ListOptions) ([]*Server, error)
	Create(ctx context.Context, scr *ServerCreateRequest) (*Server, error)
	Get(ctx context.Context, id string) (*Server, error)
	Delete(ctx context.Context, id string) error
	Resize(ctx context.Context, id string, newFlavor string) (*Server, error)
	Start(ctx context.Context, id string) (*Server, error)
	Stop(ctx context.Context, id string) (*Server, error)
	SoftReboot(ctx context.Context, id string) (*ServerMessageResponse, error)
	HardReboot(ctx context.Context, id string) (*ServerMessageResponse, error)
}

// ServerConsoleResponse contains information of server console url.
type ServerConsoleResponse struct {
	Url  string `json:"url"`
	Type string `json:"type"`
}

// ServerMessageResponse contains message response from Cloud Server API.
type ServerMessageResponse struct {
	Message string `json:"message"`
}

// ServerAction represents server action request payload.
type ServerAction struct {
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
func (s *server) Create(ctx context.Context, scr *ServerCreateRequest) (*Server, error) {
	return nil, nil
}

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

	var server *Server
	if err := json.NewDecoder(resp.Body).Decode(&server); err != nil {
		return nil, err
	}
	return server, nil
}

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

func (s *server) Resize(ctx context.Context, id string, newFlavor string) (*Server, error) {
	return nil, nil
}

func (s *server) Start(ctx context.Context, id string) (*Server, error) {
	return nil, nil
}

func (s *server) Stop(ctx context.Context, id string) (*Server, error) {
	return nil, nil
}

func (s *server) SoftReboot(ctx context.Context, id string) (*ServerMessageResponse, error) {
	return nil, nil
}

func (s *server) HardReboot(ctx context.Context, id string) (*ServerMessageResponse, error) {
	return nil, nil
}

func (s *server) Rebuild(ctx context.Context, id string, imageID string) (*Server, error) {
	return nil, nil
}

func (s *server) GetVNC(ctx context.Context, id string) (*ServerConsoleResponse, error) {
	return nil, nil
}
