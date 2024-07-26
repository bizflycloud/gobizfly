package gobizfly

type kmsService struct {
	client *Client
}

type KMSService interface {
	Certificates() *kmsCertificateService
	Secrets() *kmsSecretService
}
