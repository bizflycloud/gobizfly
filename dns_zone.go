package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
)

// Meta - Metadata of list zone response
type Meta struct {
	MaxResults int `json:"max_results"`
	Total      int `json:"total"`
	Page       int `json:"page"`
}

// Zone - contains the information of a zone
type Zone struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Deleted    int      `json:"deleted"`
	CreatedAt  string   `json:"created_at"`
	UpdatedAt  string   `json:"updated_at"`
	TenantId   string   `json:"tenant_id"`
	NameServer []string `json:"nameserver"`
	TTL        int      `json:"ttl"`
	Active     bool     `json:"active"`
}

// WrappedZonePayload - wrapped struct for creating zone payload
type WrappedZonePayload struct {
	Zones *CreateZonePayload `json:"zones"`
}

// ExtendedZone - contains the information of a zone and embedded record sets
type ExtendedZone struct {
	Zone
	RecordsSet []Record `json:"record_set"`
}

// ListZoneResp - contains the response of list zones and metadata
type ListZoneResp struct {
	Zones []Zone `json:"zones"`
	Meta  Meta   `json:"_meta"`
}

// CreateZonePayload - the payload of creating a zone request
type CreateZonePayload struct {
	Name        string `json:"name"`
	Required    bool   `json:"required,omitempty"`
	Description string `json:"description,omitempty"`
}

// ListZones - List DNS zones
func (d *dnsService) ListZones(ctx context.Context, opts *ListOptions) (*ListZoneResp, error) {
	req, err := d.client.NewRequest(ctx, http.MethodGet, dnsName, d.resourcePath(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := d.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var data *ListZoneResp
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

// CreateZone - Create a DNS zone
func (d *dnsService) CreateZone(ctx context.Context, czpl *CreateZonePayload) (*ExtendedZone, error) {
	payload := WrappedZonePayload{
		Zones: czpl,
	}
	req, err := d.client.NewRequest(ctx, http.MethodPost, dnsName, d.resourcePath(), payload)
	if err != nil {
		return nil, err
	}
	resp, err := d.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var data *ExtendedZone
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

// GetZone - Get a DNS zone
func (d *dnsService) GetZone(ctx context.Context, zoneID string) (*ExtendedZone, error) {
	req, err := d.client.NewRequest(ctx, http.MethodGet, dnsName, d.zoneItemPath(zoneID), nil)
	if err != nil {
		return nil, err
	}
	resp, err := d.client.Do(ctx, req)
	var data *ExtendedZone
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

// DeleteZone - Delete a DNS zone
func (d *dnsService) DeleteZone(ctx context.Context, zoneID string) error {
	req, err := d.client.NewRequest(ctx, http.MethodDelete, dnsName, d.zoneItemPath(zoneID), nil)
	if err != nil {
		return err
	}
	resp, err := d.client.Do(ctx, req)
	if err != nil {
		return err
	}
	return resp.Body.Close()
}
