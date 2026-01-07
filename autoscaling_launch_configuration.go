package gobizfly

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// LaunchConfiguration - is represents a launch configurations
type LaunchConfiguration struct {
	AvailabilityZone string                     `json:"availability_zone"`
	Created          string                     `json:"created_at,omitempty"`
	DataDisks        []*AutoScalingDataDisk     `json:"datadisks,omitempty"`
	Flavor           string                     `json:"flavor"`
	ID               string                     `json:"id,omitempty"`
	Metadata         map[string]interface{}     `json:"metadata"`
	Name             string                     `json:"name"`
	NetworkPlan      string                     `json:"network_plan,omitempty"`
	Networks         []*AutoScalingNetworks     `json:"networks,omitempty"`
	OperatingSystem  AutoScalingOperatingSystem `json:"os"`
	ProfileType      string                     `json:"profile_type,omitempty"`
	RootDisk         *AutoScalingDataDisk       `json:"rootdisk"`
	SecurityGroups   []*string                  `json:"security_groups,omitempty"`
	SSHKey           string                     `json:"key_name,omitempty"`
	Status           string                     `json:"status,omitempty"`
	Type             string                     `json:"type,omitempty"`
	UserData         string                     `json:"user_data,omitempty"`
}

// List launch configurations
func (lc *launchConfiguration) List(ctx context.Context, all bool) ([]*LaunchConfiguration, error) {
	req, err := lc.client.NewRequest(ctx, http.MethodGet, autoScalingServiceName, lc.resourcePath(), nil)
	if err != nil {
		return nil, err
	}

	if all {
		q := req.URL.Query()
		q.Add("all", "true")
		req.URL.RawQuery = q.Encode()
	}

	resp, err := lc.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	var data struct {
		LaunchConfigurations []*LaunchConfiguration `json:"profiles"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	for _, LaunchConfiguration := range data.LaunchConfigurations {
		// Force ProfileType = Type
		LaunchConfiguration.ProfileType = LaunchConfiguration.Type

		if LaunchConfiguration.OperatingSystem.Error != "" {
			LaunchConfiguration.Status = statusError
		} else {
			LaunchConfiguration.Status = statusActive
		}
	}
	return data.LaunchConfigurations, nil
}

// Get a launch configuration
func (lc *launchConfiguration) Get(ctx context.Context, profileID string) (*LaunchConfiguration, error) {
	req, err := lc.client.NewRequest(ctx, http.MethodGet, autoScalingServiceName, lc.itemPath(profileID), nil)
	if err != nil {
		return nil, err
	}
	resp, err := lc.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	data := &LaunchConfiguration{}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	// Force ProfileType = Type
	data.ProfileType = data.Type

	if data.OperatingSystem.Error != "" {
		data.Status = statusError
	} else {
		data.Status = statusActive
	}
	return data, nil
}

// Delete a launch configuration
func (lc *launchConfiguration) Delete(ctx context.Context, profileID string) error {
	req, err := lc.client.NewRequest(ctx, http.MethodDelete, autoScalingServiceName, lc.itemPath(profileID), nil)
	if err != nil {
		return err
	}
	resp, err := lc.client.Do(ctx, req)
	if err != nil {
		return err
	}
	_, _ = io.Copy(io.Discard, resp.Body)

	return resp.Body.Close()
}

// Create a launch configuration
func (lc *launchConfiguration) Create(ctx context.Context, lcr *LaunchConfiguration) (*LaunchConfiguration, error) {
	if _, ok := SliceContains(networkPlan, lcr.NetworkPlan); !ok {
		return nil, errors.New("UNSUPPORTED network plan")
	}

	req, err := lc.client.NewRequest(ctx, http.MethodPost, autoScalingServiceName, lc.resourcePath(), &lcr)
	if err != nil {
		return nil, err
	}

	resp, err := lc.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	data := &LaunchConfiguration{}

	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return nil, err
	}
	return data, nil
}
