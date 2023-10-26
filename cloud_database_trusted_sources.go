// This file is part of gobizfly

package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
)

const (
	cloudDatabaseTrustedSourcesResourcePath = "/trusted_sources"
)

type cloudDatabaseTrustedSources struct {
	client *Client
}

// CloudDatabaseTrustedSources contains TrustedSource information of node.
type CloudDatabaseTrustedSources struct {
	TrustedSources []string `json:"trusted_sources"`
}

func (db *cloudDatabaseService) TrustedSources() *cloudDatabaseTrustedSources {
	return &cloudDatabaseTrustedSources{client: db.client}
}

// CloudDatabase Trusted Source Resource Path
func (ts *cloudDatabaseTrustedSources) resourcePath(nodeID string) string {
	return cloudDatabaseNodesResourcePath + "/" + nodeID + cloudDatabaseTrustedSourcesResourcePath
}

// Get TrustedSource of a node.
func (ts *cloudDatabaseTrustedSources) Get(ctx context.Context, nodeID string) (*CloudDatabaseTrustedSources, error) {
	req, err := ts.client.NewRequest(ctx, http.MethodGet, databaseServiceName, ts.resourcePath(nodeID), nil)
	if err != nil {
		return nil, err
	}

	resp, err := ts.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var trustedSources *CloudDatabaseTrustedSources

	if err := json.NewDecoder(resp.Body).Decode(&trustedSources); err != nil {
		return nil, err
	}

	return trustedSources, nil
}

// Update a TrustedSource for nodeID
func (ts *cloudDatabaseTrustedSources) Update(ctx context.Context, nodeID string, tsc *CloudDatabaseTrustedSources) (*CloudDatabaseTrustedSources, error) {
	req, err := ts.client.NewRequest(ctx, http.MethodPost, databaseServiceName, ts.resourcePath(nodeID), &tsc)
	if err != nil {
		return nil, err
	}

	resp, err := ts.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var trustedSources *CloudDatabaseTrustedSources

	if err := json.NewDecoder(resp.Body).Decode(&trustedSources); err != nil {
		return nil, err
	}

	return trustedSources, nil
}
