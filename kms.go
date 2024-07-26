package gobizfly

type kmsService struct {
	client *Client
}

type KMSService interface {
	Certificate() *kmsCertificateService
	Secret() *kmsSecretService
}
