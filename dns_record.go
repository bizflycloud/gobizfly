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

// Addrs - contains the list of addresses in regions.
type Addrs struct {
	HN  []string `json:"HN"`
	HCM []string `json:"HCM"`
	SG  []string `json:"SG"`
	USA []string `json:"USA"`
}

// RoutingData - contains the routing data for version 4 and 6 addresses.
type RoutingData struct {
	AddrsV4 Addrs `json:"addrs_v4,omitempty"`
	AddrsV6 Addrs `json:"addrs_v6,omitempty"`
}

// RoutingPolicyData - contains the routing policy data for version 4 and 6 addresses and the health check information.
type RoutingPolicyData struct {
	RoutingData RoutingData `json:"routing_data,omitempty"`
	HealthCheck HealthCheck `json:"healthcheck,omitempty"`
}

// HealthCheck - contains the health check information.
type HealthCheck struct {
	TCPConnect TCPHealthCheck  `json:"tcp_connect,omitempty"`
	HTTPStatus HTTPHealthCheck `json:"http_status,omitempty"`
}

// TCPHealthCheck - contains the TCP health check information.
type TCPHealthCheck struct {
	TCPPort int `json:"tcp_port,omitempty"`
}

// HTTPHealthCheck - contains the HTTP health check information.
type HTTPHealthCheck struct {
	HTTPPort int    `json:"http_port,omitempty"`
	URLPath  string `json:"url_path,omitempty"`
	VHost    string `json:"vhost,omitempty"`
	OkCodes  []int  `json:"ok_codes,omitempty"`
	Interval int    `json:"interval,omitempty"`
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

// CreatePolicyRecordPayload - contains the payload for creating a policy record.
type CreatePolicyRecordPayload struct {
	BaseCreateRecordPayload
	RoutingPolicyData RoutingPolicyData `json:"routing_policy_data"`
}

// CreateMXRecordPayload - contains the payload for creating an MX record.
type CreateMXRecordPayload struct {
	BaseCreateRecordPayload
	Data []MXData `json:"data"`
}

// MXData - contains the data for an MX record.
type MXData struct {
	Value    string `json:"value"`
	Priority int    `json:"priority"`
}

// RecordData - contains the data for a record.
type RecordData struct {
	Value    string `json:"value"`
	Priority int    `json:"priority"`
}

// UpdateRecordPayload - contains the payload for updating a record.
type UpdateRecordPayload struct {
	Name              string            `json:"name,omitempty"`
	Type              string            `json:"type,omitempty"`
	TTL               int               `json:"ttl,omitempty"`
	Data              []MXData          `json:"data,omitempty"`
	RoutingPolicyData RoutingPolicyData `json:"routing_policy_data,omitempty"`
}

// Record - contains the information of a record.
type Record struct {
	ID                string            `json:"id"`
	Name              string            `json:"name"`
	Delete            int               `json:"deleted"`
	CreatedAt         string            `json:"created_at"`
	UpdatedAt         string            `json:"updated_at"`
	TenantID          string            `json:"tenant_id"`
	ZoneID            string            `json:"zone_id"`
	Type              string            `json:"type"`
	TTL               int               `json:"ttl"`
	Data              []interface{}     `json:"data"`
	RoutingPolicyData RoutingPolicyData `json:"routing_policy_data"`
}

// ExtendedRecord - contains the extended information of a record.
type ExtendedRecord struct {
	Record
	Data []RecordData `json:"data"`
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
func (d *dnsService) UpdateRecord(ctx context.Context, recordID string, urpl *UpdateRecordPayload) (*ExtendedRecord, error) {
	req, err := d.client.NewRequest(ctx, http.MethodPut, dnsName, d.recordItemPath(recordID), urpl)
	if err != nil {
		return nil, err
	}
	resp, err := d.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var data *ExtendedRecord
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
