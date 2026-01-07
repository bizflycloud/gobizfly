package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
)

var (
	kafkaVersionPath = "/versions"
)

type KafkaVersionListOptions struct {
	Name string `url:"name,omitempty"`
}

type KafkaVersionResponse struct {
	ID        string `json:"id" example:"string"`
	Code      string `json:"code" example:"string"`
	Name      string `json:"name" example:"string"`
	IsDefault bool   `json:"is_default" example:"true"`
}

type ListKafkaVersionResponse struct {
	Success   bool                    `json:"success" example:"true"`
	ErrorCode int                     `json:"error_code" example:"0"`
	Message   string                  `json:"message" example:"success"`
	Data      []*KafkaVersionResponse `json:"data"`
}

// List all Kafka flavors.
func (s *kafkaService) ListVersion(ctx context.Context, opts *KafkaVersionListOptions) ([]*KafkaVersionResponse, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, kafkaServiceName, kafkaVersionPath, nil)
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
	defer func() {
		_ = resp.Body.Close()
	}()
	var versionsRes *ListKafkaVersionResponse

	if err := json.NewDecoder(resp.Body).Decode(&versionsRes); err != nil {
		return nil, err
	}

	return versionsRes.Data, nil
}
