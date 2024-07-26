package gobizfly

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type KMSCertificateService interface {
	List(ctx context.Context) ([]*KMSCertificate, error)
	Get(ctx context.Context, id string) (*KMSCertificate, error)
	Create(ctx context.Context, req *KMSCertificateCreateRequest) (*KMSCertificateCreateResponse, error)
	Delete(ctx context.Context, id string) error
}

type kmsCertificateService struct {
	client *Client
}

func (k *kmsService) Certificate() *kmsCertificateService {
	return &kmsCertificateService{
		client: k.client,
	}
}

type KMSCertificate struct {
	ContainerId string `json:"container_id"`
	Name        string `json:"name"`
}

type KMSCertificateCreateRequest struct {
	CertContainer KMSCertContainer `json:"cert_container"`
}

type KMSCertContainer struct {
	Name                 string `json:"name"`
	Certificate          string `json:"certificate"`
	PrivateKey           string `json:"private_key"`
	PrivateKeyPassphrase string `json:"private_key_passphrase"`
	Intermediates        string `json:"intermediates"`
}

type KMSCertificateCreateResponse struct {
	CertificateHref string `json:"certificate_href"`
}

type KMSCertificateListResponse struct {
	CertificateContrainer []*KMSCertificate `json:"certificate_container"`
	Total                 int               `json:"total"`
}

const (
	certificateServicePath = "/certificate_container"
)

func (c *kmsCertificateService) List(ctx context.Context) ([]*KMSCertificate, error) {
	path := certificateServicePath
	req, err := c.client.NewRequest(ctx, http.MethodGet, kmsServiceName, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var respDecode KMSCertificateListResponse
	if err := json.NewDecoder(resp.Body).Decode(&respDecode); err != nil {
		return nil, err
	}

	return respDecode.CertificateContrainer, nil
}

func (c *kmsCertificateService) Get(ctx context.Context, id string) (*KMSCertificate, error) {
	path := certificateServicePath + "/" + id
	req, err := c.client.NewRequest(ctx, http.MethodGet, kmsServiceName, path, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data *KMSCertificate
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func (c *kmsCertificateService) Create(ctx context.Context, payload *KMSCertificateCreateRequest) (*KMSCertificateCreateResponse, error) {
	path := certificateServicePath

	req, err := c.client.NewRequest(ctx, http.MethodPost, kmsServiceName, path, payload)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data *KMSCertificateCreateResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func (c *kmsCertificateService) Delete(ctx context.Context, id string) error {
	path := certificateServicePath + "/" + id
	req, err := c.client.NewRequest(ctx, http.MethodDelete, kmsServiceName, path, nil)
	if err != nil {
		return err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return err
	}
	_, _ = io.Copy(io.Discard, resp.Body)

	return resp.Body.Close()
}
