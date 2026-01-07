// This file is part of gobizfly

package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
)

const (
	cloudDatabaseEnginesResourcePath = "/engines"
)

type cloudDatabaseEngines struct {
	client *Client
}

type cloudDatabaseEngineParameters struct {
	client *Client
}

// CloudDatabaseEngineVersion contains datastore version detail.
type CloudDatabaseEngineVersion struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

// CloudDatabaseEngine contains detail about a datastore.
type CloudDatabaseEngine struct {
	ID       string                       `json:"id"`
	Name     string                       `json:"name"`
	Versions []CloudDatabaseEngineVersion `json:"versions"`
}

// CloudDatabaseEngineParameters contains datastore parameters info.
type CloudDatabaseEngineParameters struct {
	Parameters []map[string]interface{} `json:"configuration_parameters"`
}

func (db *cloudDatabaseService) Engines() *cloudDatabaseEngines {
	return &cloudDatabaseEngines{client: db.client}
}

func (db *cloudDatabaseService) EngineParameters() *cloudDatabaseEngineParameters {
	return &cloudDatabaseEngineParameters{client: db.client}
}

// CloudDatabase Engine Resource Path
func (en *cloudDatabaseEngineParameters) resourcePath(datastore string, datastoreVersion string) string {
	return cloudDatabaseEnginesResourcePath + "/" + datastore + "/versions/" + datastoreVersion + "/parameters"
}

// List all database engine.
func (en *cloudDatabaseEngines) List(ctx context.Context) ([]*CloudDatabaseEngine, error) {
	req, err := en.client.NewRequest(ctx, http.MethodGet, databaseServiceName, cloudDatabaseEnginesResourcePath, nil)
	if err != nil {
		return nil, err
	}

	resp, err := en.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()
	var data struct {
		Engines []*CloudDatabaseEngine `json:"engines"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Engines, nil
}

// Get a database engine parameters.
func (en *cloudDatabaseEngineParameters) Get(ctx context.Context, datastore string, datastoreVersion string) (*CloudDatabaseEngineParameters, error) {
	req, err := en.client.NewRequest(ctx, http.MethodGet, databaseServiceName, en.resourcePath(datastore, datastoreVersion), nil)
	if err != nil {
		return nil, err
	}

	resp, err := en.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()
	var engineParameters *CloudDatabaseEngineParameters

	if err := json.NewDecoder(resp.Body).Decode(&engineParameters); err != nil {
		return nil, err
	}

	return engineParameters, nil
}
