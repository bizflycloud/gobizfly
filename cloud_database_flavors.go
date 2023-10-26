// This file is part of gobizfly

package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
)

const (
	cloudDatabaseFlavorsResourcePath = "/flavors"
)

type cloudDatabaseFlavors struct {
	client *Client
}

// CloudDatabaseFlavor contains detail about a datastore.
type CloudDatabaseFlavor struct {
	Category           string                 `json:"category"`
	Datastore          string                 `json:"datastore"`
	DatastoreID        string                 `json:"datastore_id"`
	DatastoreVersionID string                 `json:"datastore_version_id"`
	Version            string                 `json:"version"`
	Flavors            []string               `json:"flavors"`
	LimitVolumesSize   map[string]interface{} `json:"limit_volumes_size"`
}

func (flv *cloudDatabaseService) Flavors() *cloudDatabaseFlavors {
	return &cloudDatabaseFlavors{client: flv.client}
}

// CloudDatabase Flavor Resource Path
func (flv *cloudDatabaseFlavors) resourcePath(datastore string, datastoreVersion string) string {
	return cloudDatabaseFlavorsResourcePath + "/" + datastore + "/versions/" + datastoreVersion
}

// List all database flavor.
func (flv *cloudDatabaseFlavors) List(ctx context.Context) ([]*CloudDatabaseFlavor, error) {
	req, err := flv.client.NewRequest(ctx, http.MethodGet, databaseServiceName, cloudDatabaseFlavorsResourcePath, nil)
	if err != nil {
		return nil, err
	}

	resp, err := flv.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var flavors []*CloudDatabaseFlavor

	if err := json.NewDecoder(resp.Body).Decode(&flavors); err != nil {
		return nil, err
	}

	return flavors, nil
}

// Get a database flavor parameters.
func (flv *cloudDatabaseFlavors) Get(ctx context.Context, datastore string, datastoreVersion string) ([]*CloudDatabaseFlavor, error) {
	req, err := flv.client.NewRequest(ctx, http.MethodGet, databaseServiceName, flv.resourcePath(datastore, datastoreVersion), nil)
	if err != nil {
		return nil, err
	}

	resp, err := flv.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var flavors []*CloudDatabaseFlavor

	if err := json.NewDecoder(resp.Body).Decode(&flavors); err != nil {
		return nil, err
	}

	return flavors, nil
}
