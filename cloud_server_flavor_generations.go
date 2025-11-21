package gobizfly

import (
    "context"
    "encoding/json"
    "fmt"
    "net/http"
    "net/url"
)

// Define the resource path for flavor generations
const (
    flavorGenerationsResourcePath = "/flavor-generations"
)

// cloudFlavorGenerations handles flavor generation requests.
type cloudFlavorGenerations struct {
    client *Client
}

// FlavorGeneration represents the flavor generation info structure.
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
}

// flavorGenerationsResponse is used to decode the top-level response structure.
type flavorGenerationsResponse struct {
    Data []FlavorGeneration `json:"data"`
}

// ListOption is a function type for optional parameters
type ListOption func(*listOptions)

// listOptions contains the optional filter parameters for List
type listOptions struct {
    az       string
    category string
}

// WithAZ sets the availability zone filter
func WithAZ(az string) ListOption {
    return func(opts *listOptions) {
        opts.az = az
    }
}

// WithCategory sets the category filter
func WithCategory(category string) ListOption {
    return func(opts *listOptions) {
        opts.category = category
    }
}

// FlavorGenerations returns a cloudFlavorGenerations client.
func (c *Client) FlavorGenerations() *cloudFlavorGenerations {
    return &cloudFlavorGenerations{client: c}
}

// List fetches flavor generations with optional filters.
func (fg *cloudFlavorGenerations) List(ctx context.Context, opts ...ListOption) ([]FlavorGeneration, error) {
    // Apply options
    options := &listOptions{}
    for _, opt := range opts {
        opt(options)
    }

    // Build query parameters
    queryParams := url.Values{}
    if options.az != "" {
        queryParams.Add("az", options.az)
    }
    if options.category != "" {
        queryParams.Add("category", options.category)
    }

    path := flavorGenerationsResourcePath
    if len(queryParams) > 0 {
        path = fmt.Sprintf("%s?%s", flavorGenerationsResourcePath, queryParams.Encode())
    }

    // Create the request with the path and query
    req, err := fg.client.NewRequest(ctx, http.MethodGet, "", path, nil)
    if err != nil {
        return nil, err
    }

    resp, err := fg.client.Do(ctx, req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var r flavorGenerationsResponse
    if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
        return nil, err
    }

    return r.Data, nil
}
