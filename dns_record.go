package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

// WrappedRecordPayload represents a wrapped interface record payload.
type WrappedRecordPayload struct {
	Record interface{} `json:"record"`
}

// MXData - contains the data for an MX record.
type MXData struct {
	Value    string `json:"value"`
	Priority int    `json:"priority"`
}

// SRVData - contains the data for an SRV record.
type SRVData struct {
	Port     int    `json:"port"`
	Priority int    `json:"priority"`
	Protocol string `json:"protocol"`
	Service  string `json:"service"`
	Target   string `json:"target"`
	Weight   int    `json:"weight"`
}

// BaseCreateRecordPayload - contains the base payload for creating a record.
type BaseCreateRecordPayload struct {
	Name string `json:"name"`
	Type string `json:"type"`
	TTL  int    `json:"ttl"`
}

// CreateNormalRecordPayload - contains the payload for creating a normal record.
type CreateNormalRecordPayload struct {
	BaseCreateRecordPayload
	Data []string `json:"data"`
}

// CreateMXRecordPayload - contains the payload for creating an MX record.
type CreateMXRecordPayload struct {
	BaseCreateRecordPayload
	Data []MXData `json:"data"`
}

// CreateSRVRecordPayload - contains the payload for creating an SRV record.
type CreateSRVRecordPayload struct {
	BaseCreateRecordPayload
	Data []SRVData `json:"data"`
}

// UpdateRecordPayload - contains the payload for updating a record.
type BaseUpdateRecordPayload struct {
	Name string   `json:"name,omitempty"`
	Type string   `json:"type,omitempty"`
	TTL  int      `json:"ttl,omitempty"`
	Data []MXData `json:"data,omitempty"`
}
type UpdateNormalRecordPayload struct {
	BaseUpdateRecordPayload
	Data []string `json:"data"`
}

// CreateMXRecordPayload - contains the payload for creating an MX record.
type UpdateMXRecordPayload struct {
	BaseUpdateRecordPayload
	Data []MXData `json:"data"`
}

// UpdateSRVRecordPayload - contains the payload for creating an SRV record.
type UpdateSRVRecordPayload struct {
	BaseUpdateRecordPayload
	Data []SRVData `json:"data"`
}

// Record - contains the information of a record.
type Record struct {
	ID        string        `json:"id"`
	Name      string        `json:"name"`
	Delete    int           `json:"deleted"`
	CreatedAt string        `json:"created_at"`
	UpdatedAt string        `json:"updated_at"`
	TenantID  string        `json:"tenant_id"`
	ZoneID    string        `json:"zone_id"`
	Type      string        `json:"type"`
	TTL       int           `json:"ttl"`
	Data      []interface{} `json:"data"`
}

// Records - contains the list of records.
type Records struct {
	Records []Record `json:"records"`
}

// CreateRecord - Creates a DNS record.
func (d *dnsService) CreateRecord(ctx context.Context, zoneID string, crpl interface{}) (*Record, error) {
	payload := WrappedRecordPayload{
		Record: crpl,
	}
	req, err := d.client.NewRequest(ctx, http.MethodPost, dnsName, strings.Join([]string{d.zoneItemPath(zoneID), "record"}, "/"), &payload)
	if err != nil {
		return nil, err
	}
	resp, err := d.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var data *struct {
		Record *Record `json:"record"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data.Record, nil
}

// GetRecord - Get a DNS record
func (d *dnsService) GetRecord(ctx context.Context, recordID string) (*Record, error) {
	req, err := d.client.NewRequest(ctx, http.MethodGet, dnsName, d.recordItemPath(recordID), nil)
	if err != nil {
		return nil, err
	}
	resp, err := d.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var data *struct {
		Record *Record `json:"record"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data.Record, nil
}

// UpdateRecord - Update a DNS record
func (d *dnsService) UpdateRecord(ctx context.Context, recordID string, urpl interface{}) (*Record, error) {
	req, err := d.client.NewRequest(ctx, http.MethodPut, dnsName, d.recordItemPath(recordID), urpl)
	if err != nil {
		return nil, err
	}
	resp, err := d.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var data *Record
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

// DeleteRecord - Delete a DNS record
func (d *dnsService) DeleteRecord(ctx context.Context, recordID string) error {
	req, err := d.client.NewRequest(ctx, http.MethodDelete, dnsName, d.recordItemPath(recordID), nil)
	if err != nil {
		return err
	}
	resp, err := d.client.Do(ctx, req)
	if err != nil {
		return err
	}
	return resp.Body.Close()
}
