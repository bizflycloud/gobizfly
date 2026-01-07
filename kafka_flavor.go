package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
)

var (
	kafkaFlavorPath = "/flavors"
)

type KafkaFlavorListOptions struct {
	Name string `url:"name,omitempty"`
}

type KafkaFlavorResponse struct {
	BillingPlan string `json:"billing_plan" example:"string"`
	Code        string `json:"code" example:"string"`
	Description string `json:"description" example:"string"`
	Disk        int    `json:"disk" example:"0"`
	FlavorType  string `json:"flavor_type" example:"string"`
	ID          string `json:"id" example:"string"`
	IsDefault   bool   `json:"is_default" example:"true"`
	Name        string `json:"name" example:"string"`
	PlanName    string `json:"plan_name" example:"string"`
	RAM         int    `json:"ram" example:"0"`
	VCPUs       int    `json:"vcpus" example:"0"`
}

type ListKafkaFlavorResponse struct {
	Success   bool                   `json:"success" example:"true"`
	ErrorCode int                    `json:"error_code" example:"0"`
	Message   string                 `json:"message" example:"success"`
	Data      []*KafkaFlavorResponse `json:"data"`
}

// List all Kafka flavors.
func (s *kafkaService) ListFlavor(ctx context.Context, opts *KafkaFlavorListOptions) ([]*KafkaFlavorResponse, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, kafkaServiceName, kafkaFlavorPath, nil)
	if err != nil {
		return nil, err
	}
	params := req.URL.Query()
	if opts != nil {
		if opts.Name != "" {
			params.Add("name", opts.Name)
		}
	}
	req.URL.RawQuery = params.Encode()

	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var flavorsRes *ListKafkaFlavorResponse

	if err := json.NewDecoder(resp.Body).Decode(&flavorsRes); err != nil {
		return nil, err
	}

	return flavorsRes.Data, nil
}
