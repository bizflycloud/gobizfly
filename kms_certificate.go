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
	Create(ctx context.Context, req *KMSCertificateContainerCreateRequest) (*KMSCertificateCreateResponse, error)
	Delete(ctx context.Context, id string) error
}

type kmsCertificateService struct {
	client *Client
}

func (k *kmsService) Certificates() *kmsCertificateService {
	return &kmsCertificateService{
		client: k.client,
	}
}

type KMSCertificate struct {
	ContainerID string `json:"container_id"`
	Name        string `json:"name"`
}

type KMSCertificateContainerCreateRequest struct {
	CertContainer KMSCertContainer `json:"cert_container"`
}

type KMSCertContainer struct {
	Name                 string                              `json:"name"`
	Certificate          KMSCertificateCreateReqest          `json:"certificate"`
	PrivateKey           KMSPrivateKeyCreateReqest           `json:"private_key"`
	PrivateKeyPassphrase KMSPrivateKeyPassphraseCreateReqest `json:"private_key_passphrase"`
	Intermediates        *KMSIntermediatesCreateReqest       `json:"intermediates,omitempty"`
}

type KMSCertificateCreateReqest struct {
	Name    string `json:"name"`
	Payload string `json:"payload"`
}

type KMSPrivateKeyCreateReqest struct {
	Name    string `json:"name"`
	Payload string `json:"payload"`
}

type KMSPrivateKeyPassphraseCreateReqest struct {
	Name    string `json:"name"`
	Payload string `json:"payload"`
}

type KMSIntermediatesCreateReqest struct {
	Name    string `json:"name,omitempty"`
	Payload string `json:"payload,omitempty"`
}

type KMSCertificateCreateResponse struct {
	CertificateHref string `json:"certificate_href"`
}

type KMSCertificateListResponse struct {
	CertificateContainer []*KMSCertificate `json:"certificate_container"`
	Total                int               `json:"total"`
}

type KMSCertificateGetResponse struct {
	ContainerID string `json:"container_id"`
	Name        string `json:"name"`
	Certificate string `json:"certificate"`
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
	defer func() {
		_ = resp.Body.Close()
	}()

	var respDecode KMSCertificateListResponse
	if err := json.NewDecoder(resp.Body).Decode(&respDecode); err != nil {
		return nil, err
	}

	return respDecode.CertificateContainer, nil
}

func (c *kmsCertificateService) Get(ctx context.Context, id string) (*KMSCertificateGetResponse, error) {
	path := certificateServicePath + "/" + id
	req, err := c.client.NewRequest(ctx, http.MethodGet, kmsServiceName, path, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	var data *KMSCertificateGetResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func (c *kmsCertificateService) Create(ctx context.Context, payload *KMSCertificateContainerCreateRequest) (*KMSCertificateCreateResponse, error) {
	path := certificateServicePath

	req, err := c.client.NewRequest(ctx, http.MethodPost, kmsServiceName, path, payload)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

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
