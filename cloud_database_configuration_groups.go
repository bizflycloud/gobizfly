// This file is part of gobizfly

package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
)

const (
	cloudDatabaseConfigurationsResourcePath = "/configurations"
)

type cloudDatabaseConfigurations struct {
	client *Client
}

// CloudDatabaseNodeUseConfiguration contains info of node use configuration.
type CloudDatabaseNodeUseConfiguration struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// CloudDatabaseConfiguration contains detail of a configuration.
type CloudDatabaseConfiguration struct {
	CreatedAt string                              `json:"created_at"`
	Datastore CloudDatabaseDatastore              `json:"datastore"`
	ID        string                              `json:"id"`
	Message   string                              `json:"message"`
	Name      string                              `json:"name"`
	NodeCount int                                 `json:"node_count"`
	Nodes     []CloudDatabaseNodeUseConfiguration `json:"nodes"`
	Values    map[string]interface{}              `json:"values"`
}

// CloudDatabaseConfigurationCreate contains payload create configuration.
type CloudDatabaseConfigurationCreate struct {
	Datastore  CloudDatabaseDatastore `json:"datastore,omitempty" validate:"required"`
	Name       string                 `json:"configuration_name,omitempty" validate:"required"`
	Parameters map[string]interface{} `json:"configuration_parameters,omitempty" validate:"required"`
}

// CloudDatabaseConfigurationUpdate contains payload to update a configuration.
type CloudDatabaseConfigurationUpdate struct {
	Datastore  CloudDatabaseDatastore `json:"datastore,omitempty" validate:"required"`
	Parameters map[string]interface{} `json:"configuration_parameters" validate:"required"`
}

func (db *cloudDatabaseService) Configurations() *cloudDatabaseConfigurations {
	return &cloudDatabaseConfigurations{client: db.client}
}

// CloudDatabase Configuration Resource Path
func (cfg *cloudDatabaseConfigurations) resourcePath(cfgID string) string {
	return cloudDatabaseConfigurationsResourcePath + "/" + cfgID
}

// CloudDatabase Configuration Resource Action Path
func (cfg *cloudDatabaseConfigurations) resourceActionPath(nodeID string, cfgID string) string {
	return cloudDatabaseNodesResourcePath + "/" + nodeID + "/configuration/" + cfgID
}

// List all configurations.
func (cfg *cloudDatabaseConfigurations) List(ctx context.Context, opts *CloudDatabaseListOption) ([]*CloudDatabaseConfiguration, error) {
	req, err := cfg.client.NewRequest(ctx, http.MethodGet, databaseServiceName, cloudDatabaseConfigurationsResourcePath, nil)
	if err != nil {
		return nil, err
	}

	AddParamsListOption(req, opts)

	resp, err := cfg.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var data struct {
		Configurations []*CloudDatabaseConfiguration `json:"configurations"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Configurations, nil
}

// Create a new configurations.
func (cfg *cloudDatabaseConfigurations) Create(ctx context.Context, cr *CloudDatabaseConfigurationCreate) (*CloudDatabaseConfiguration, error) {
	req, err := cfg.client.NewRequest(ctx, http.MethodPost, databaseServiceName, cloudDatabaseConfigurationsResourcePath, &cr)
	if err != nil {
		return nil, err
	}

	resp, err := cfg.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var configuration *CloudDatabaseConfiguration
	if err := json.NewDecoder(resp.Body).Decode(&configuration); err != nil {
		return nil, err
	}
	return configuration, nil
}

// Get a configurations.
func (cfg *cloudDatabaseConfigurations) Get(ctx context.Context, cfgID string) (*CloudDatabaseConfiguration, error) {
	req, err := cfg.client.NewRequest(ctx, http.MethodGet, databaseServiceName, cfg.resourcePath(cfgID), nil)
	if err != nil {
		return nil, err
	}

	resp, err := cfg.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var configuration *CloudDatabaseConfiguration

	if err := json.NewDecoder(resp.Body).Decode(&configuration); err != nil {
		return nil, err
	}

	return configuration, nil
}

// Action with a configurations to node.
func (cfg *cloudDatabaseConfigurations) Action(ctx context.Context, nodeID string, cfgID string, iar *CloudDatabaseAction) (*CloudDatabaseMessageResponse, error) {
	req, err := cfg.client.NewRequest(ctx, http.MethodPost, databaseServiceName, cfg.resourceActionPath(nodeID, cfgID), &iar)
	if err != nil {
		return nil, err
	}

	resp, err := cfg.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var dmr *CloudDatabaseMessageResponse

	if err := json.NewDecoder(resp.Body).Decode(&dmr); err != nil {
		return nil, err
	}

	return dmr, nil
}

// Attach with a configurations to node.
func (cfg *cloudDatabaseConfigurations) Attach(ctx context.Context, nodeID string, cfgID string, all bool) (*CloudDatabaseMessageResponse, error) {
	var iar = CloudDatabaseConfigAction{Action: attach}

	req, err := cfg.client.NewRequest(ctx, http.MethodPost, databaseServiceName, cfg.resourceActionPath(nodeID, cfgID), &iar)
	if err != nil {
		return nil, err
	}

	resp, err := cfg.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var dmr *CloudDatabaseMessageResponse

	if err := json.NewDecoder(resp.Body).Decode(&dmr); err != nil {
		return nil, err
	}

	return dmr, nil
}

// Detach with a configurations to node.
func (cfg *cloudDatabaseConfigurations) Detach(ctx context.Context, nodeID string, cfgID string, all bool) (*CloudDatabaseMessageResponse, error) {
	var iar = CloudDatabaseAction{Action: detach, ActionAll: all}

	req, err := cfg.client.NewRequest(ctx, http.MethodPost, databaseServiceName, cfg.resourceActionPath(nodeID, cfgID), &iar)
	if err != nil {
		return nil, err
	}

	resp, err := cfg.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var dmr *CloudDatabaseMessageResponse

	if err := json.NewDecoder(resp.Body).Decode(&dmr); err != nil {
		return nil, err
	}

	return dmr, nil
}

// Update a configurations.
func (cfg *cloudDatabaseConfigurations) Update(ctx context.Context, cfgID string, cu *CloudDatabaseConfigurationUpdate) (*CloudDatabaseMessageResponse, error) {
	req, err := cfg.client.NewRequest(ctx, http.MethodPut, databaseServiceName, cfg.resourcePath(cfgID), &cu)
	if err != nil {
		return nil, err
	}

	resp, err := cfg.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var dmr *CloudDatabaseMessageResponse

	if err := json.NewDecoder(resp.Body).Decode(&dmr); err != nil {
		return nil, err
	}

	return dmr, nil
}

// Delete a configurations.
func (cfg *cloudDatabaseConfigurations) Delete(ctx context.Context, cfgID string) (*CloudDatabaseMessageResponse, error) {
	req, err := cfg.client.NewRequest(ctx, http.MethodDelete, databaseServiceName, cfg.resourcePath(cfgID), nil)
	if err != nil {
		return nil, err
	}

	resp, err := cfg.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var dmr *CloudDatabaseMessageResponse

	if err := json.NewDecoder(resp.Body).Decode(&dmr); err != nil {
		return nil, err
	}

	return dmr, nil
}
