// This file is part of gobizfly

package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

const (
	firewallBasePath = "/firewalls"
)

var _ FirewallService = (*cloudServerFirewallResource)(nil)

type cloudServerFirewallResource struct {
	client *Client
}

func (cs *cloudServerService) Firewalls() *cloudServerFirewallResource {
	return &cloudServerFirewallResource{client: cs.client}
}

type FirewallService interface {
	List(ctx context.Context, opts *ListOptions) ([]*Firewall, error)
	Create(ctx context.Context, fcr *FirewallRequestPayload) (*FirewallDetail, error)
	Get(ctx context.Context, id string) (*FirewallDetail, error)
	Delete(ctx context.Context, id string) (*FirewallDeleteResponse, error)
	RemoveServer(ctx context.Context, id string, rsfr *FirewallRemoveServerRequest) (*Firewall, error)
	Update(ctx context.Context, id string, ufr *FirewallRequestPayload) (*FirewallDetail, error)
	DeleteRule(ctx context.Context, id string) (*FirewallDeleteResponse, error)
}

// BaseFirewall - contains base information fields of a firewall
type BaseFirewall struct {
	ID                    string         `json:"id"`
	Name                  string         `json:"name"`
	Description           string         `json:"description"`
	Tags                  []string       `json:"tags"`
	CreatedAt             string         `json:"created_at"`
	UpdatedAt             string         `json:"updated_at"`
	RevisionNumber        int            `json:"revision_number"`
	ProjectID             string         `json:"project_id"`
	ServersCount          int            `json:"servers_count"`
	RulesCount            int            `json:"rules_count"`
	NetworkInterfaceCount int            `json:"network_interface_count"`
	InBound               []FirewallRule `json:"inbound"`
	OutBound              []FirewallRule `json:"outbound"`
}

// Firewall - contains information fields of a firewall and the applied servers
type Firewall struct {
	BaseFirewall
	Servers []string `json:"servers"`
}

// FirewallDetail - contains information fields of a firewall and the applied servers and network interfaces
type FirewallDetail struct {
	BaseFirewall
	Servers          []*Server           `json:"servers"`
	NetworkInterface []*NetworkInterface `json:"network_interface"`
}

// FirewallRule - contains information fields of a firewall rule
type FirewallRule struct {
	ID             string   `json:"id"`
	FirewallID     string   `json:"security_group_id"`
	EtherType      string   `json:"ethertype"`
	Direction      string   `json:"direction"`
	Protocol       string   `json:"protocol"`
	PortRangeMin   int      `json:"port_range_min"`
	PortRangeMax   int      `json:"port_range_max"`
	RemoteIPPrefix string   `json:"remote_ip_prefix"`
	Description    string   `json:"description"`
	Tags           []string `json:"tags"`
	CreatedAt      string   `json:"created_at"`
	UpdatedAt      string   `json:"updated_at"`
	RevisionNumber int      `json:"revision_number"`
	ProjectID      string   `json:"project_id"`
	Type           string   `json:"type"`
	CIDR           string   `json:"cidr"`
	PortRange      string   `json:"port_range"`
}

// FirewallRuleCreateRequest - payload for creating a firewall rule
type FirewallRuleCreateRequest struct {
	Type      interface{} `json:"type"`
	Protocol  interface{} `json:"protocol"`
	PortRange interface{} `json:"port_range"`
	CIDR      string      `json:"cidr"`
}

// FirewallRequestPayload - payload for creating a firewall
type FirewallRequestPayload struct {
	Name              string                      `json:"name"`
	InBound           []FirewallRuleCreateRequest `json:"inbound,omitempty"`
	OutBound          []FirewallRuleCreateRequest `json:"outbound,omitempty"`
	Targets           []string                    `json:"targets,omitempty"` // Deprecated: This field will be removed in the near future
	NetworkInterfaces []string                    `json:"network_interfaces,omitempty"`
}

// FirewallDeleteResponse represents the response body when deleting a firewall
type FirewallDeleteResponse struct {
	Message string `json:"message"`
}

// FirewallRemoveServerRequest represents the request body when removing a server from a firewall
type FirewallRemoveServerRequest struct {
	Servers []string `json:"servers"`
}

// FirewallRuleCreateResponse represents the response body when creating a firewall rule
type FirewallRuleCreateResponse struct {
	Rule FirewallRule `json:"security_group_rule"`
}

// List lists all firewall.
func (f *cloudServerFirewallResource) List(ctx context.Context, opts *ListOptions) ([]*Firewall, error) {

	req, err := f.client.NewRequest(ctx, http.MethodGet, serverServiceName, firewallBasePath, nil)
	if err != nil {
		return nil, err
	}

	resp, err := f.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var firewalls []*Firewall

	if err := json.NewDecoder(resp.Body).Decode(&firewalls); err != nil {
		return nil, err
	}

	for _, firewall := range firewalls {
		for index, rule := range firewall.InBound {
			if rule.CIDR != "" {
				continue
			}

			if rule.EtherType == "IPv4" {
				firewall.InBound[index].CIDR = "0.0.0.0/0"
			}
			if rule.EtherType == "IPv6" {
				firewall.InBound[index].CIDR = "::/0"
			}
		}

		for index, rule := range firewall.OutBound {
			if rule.CIDR != "" {
				continue
			}

			if rule.EtherType == "IPv4" {
				firewall.OutBound[index].CIDR = "0.0.0.0/0"
			}
			if rule.EtherType == "IPv6" {
				firewall.OutBound[index].CIDR = "::/0"
			}
		}
	}
	return firewalls, nil
}

// Create a firewall.
func (f *cloudServerFirewallResource) Create(ctx context.Context, fcr *FirewallRequestPayload) (*FirewallDetail, error) {

	req, err := f.client.NewRequest(ctx, http.MethodPost, serverServiceName, firewallBasePath, fcr)
	if err != nil {
		return nil, err
	}

	resp, err := f.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var firewall *FirewallDetail

	if err := json.NewDecoder(resp.Body).Decode(&firewall); err != nil {
		return nil, err
	}

	return firewall, nil
}

// Get detail a firewall.
func (f *cloudServerFirewallResource) Get(ctx context.Context, id string) (*FirewallDetail, error) {

	req, err := f.client.NewRequest(ctx, http.MethodGet, serverServiceName, firewallBasePath+"/"+id, nil)
	if err != nil {
		return nil, err
	}

	resp, err := f.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var firewall *FirewallDetail

	if err := json.NewDecoder(resp.Body).Decode(&firewall); err != nil {
		return nil, err
	}

	for index, rule := range firewall.InBound {
		if rule.CIDR != "" {
			continue
		}

		if rule.EtherType == "IPv4" {
			firewall.InBound[index].CIDR = "0.0.0.0/0"
		}
		if rule.EtherType == "IPv6" {
			firewall.InBound[index].CIDR = "::/0"
		}
	}

	for index, rule := range firewall.OutBound {
		if rule.CIDR != "" {
			continue
		}

		if rule.EtherType == "IPv4" {
			firewall.OutBound[index].CIDR = "0.0.0.0/0"
		}
		if rule.EtherType == "IPv6" {
			firewall.OutBound[index].CIDR = "::/0"
		}
	}

	return firewall, nil
}

// RemoveServer - Remove applied servers from a firewall.
func (f *cloudServerFirewallResource) RemoveServer(ctx context.Context, id string, rsfr *FirewallRemoveServerRequest) (*Firewall, error) {

	req, err := f.client.NewRequest(ctx, http.MethodDelete, serverServiceName, strings.Join([]string{firewallBasePath, id, "servers"}, "/"), rsfr)

	if err != nil {
		return nil, err
	}

	resp, err := f.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var firewall *Firewall

	if err := json.NewDecoder(resp.Body).Decode(&firewall); err != nil {
		return nil, err
	}

	return firewall, nil
}

// Update Firewall
func (f *cloudServerFirewallResource) Update(ctx context.Context, id string, ufr *FirewallRequestPayload) (*FirewallDetail, error) {

	req, err := f.client.NewRequest(ctx, http.MethodPatch, serverServiceName, firewallBasePath+"/"+id, ufr)

	if err != nil {
		return nil, err
	}

	resp, err := f.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var firewall *FirewallDetail

	if err := json.NewDecoder(resp.Body).Decode(&firewall); err != nil {
		return nil, err
	}

	return firewall, nil
}

// Delete a Firewall
func (f *cloudServerFirewallResource) Delete(ctx context.Context, id string) (*FirewallDeleteResponse, error) {

	req, err := f.client.NewRequest(ctx, http.MethodDelete, serverServiceName, firewallBasePath+"/"+id, nil)

	if err != nil {
		return nil, err
	}

	resp, err := f.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var dwr *FirewallDeleteResponse

	if err := json.NewDecoder(resp.Body).Decode(&dwr); err != nil {
		return nil, err
	}

	return dwr, nil
}

// DeleteRule - delete a rule in a firewall
func (f *cloudServerFirewallResource) DeleteRule(ctx context.Context, id string) (*FirewallDeleteResponse, error) {
	req, err := f.client.NewRequest(ctx, http.MethodDelete, serverServiceName, firewallBasePath+"/"+id, nil)

	if err != nil {
		return nil, err
	}

	resp, err := f.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var dwr *FirewallDeleteResponse

	if err := json.NewDecoder(resp.Body).Decode(&dwr); err != nil {
		return nil, err
	}

	return dwr, nil
}
