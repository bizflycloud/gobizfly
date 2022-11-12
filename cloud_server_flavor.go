package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
)

// serverFlavorResponse is the response from the server flavor API.
type serverFlavorResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"` // Deprecated: Will be removed in the future.
	VCPUs    int    `json:"vcpus"`
	RAM      int    `json:"ram"`
	Disk     int    `json:"disk"`
	Category string `json:"category"`
}

// ListFlavors lists server flavors
func (s *server) ListFlavors(ctx context.Context) ([]*serverFlavorResponse, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, serverServiceName, flavorPath, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	var flavors []*serverFlavorResponse

	if err := json.NewDecoder(resp.Body).Decode(&flavors); err != nil {
		return nil, err
	}
	return flavors, nil
}
