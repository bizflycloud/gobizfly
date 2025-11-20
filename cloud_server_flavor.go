package gobizfly

import (
    "context"
    "encoding/json"
    "net/http"
)

// ServerFlavorResponse represents a server flavor from the API
type ServerFlavorResponse struct {
    ID           string           `json:"id"`
    Name         string           `json:"name"`
    RAM          int              `json:"ram"`           // RAM in MB
    Disk         int              `json:"disk"`          // Disk in GB
    Swap         string           `json:"swap"`
    VCPUs        int              `json:"vcpus"`
    GenerationID string           `json:"generation_id"`
    Generation   FlavorGeneration `json:"generation"`
    Category     string           `json:"category"`      // premium, basic, enterprise, etc.
    GPU          *FlavorGPU       `json:"gpu"`           // Nullable
    BillingPlans []string         `json:"billing_plans"`
    IsNew        bool             `json:"is_new"`
}

// FlavorGeneration represents hardware generation information
type FlavorGeneration struct {
    ID                string   `json:"id"`
    Name              string   `json:"name"`
    Code              string   `json:"code"`
    Category          string   `json:"category"`
    Vendor            string   `json:"vendor"`
    Model             string   `json:"model"`
    Group             string   `json:"group"`
    Description       string   `json:"description"`
    AvailabilityZones []string `json:"availability_zones"`
    Icon              string   `json:"icon"`
}

// FlavorGPU represents GPU information
type FlavorGPU struct {
    Name  string `json:"name"`
    Count int    `json:"count"`
}

type cloudServerFlavorResource struct {
    client *Client
}

func (cs *cloudServerService) Flavors() *cloudServerFlavorResource {
    return &cloudServerFlavorResource{client: cs.client}
}

// List lists server flavors
func (s *cloudServerFlavorResource) List(ctx context.Context) ([]*ServerFlavorResponse, error) {
    req, err := s.client.NewRequest(ctx, http.MethodGet, serverServiceName, flavorPath, nil)
    if err != nil {
        return nil, err
    }
    
    q := req.URL.Query()
    q.Set("new", "true")
    req.URL.RawQuery = q.Encode()
    
    resp, err := s.client.Do(ctx, req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    
    var flavors []*ServerFlavorResponse
    if err := json.NewDecoder(resp.Body).Decode(&flavors); err != nil {
        return nil, err
    }
    
    return flavors, nil
}
