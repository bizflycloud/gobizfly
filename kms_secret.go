package gobizfly

import "context"

type KMSSecretService interface {
	List(ctx context.Context, page, total int) ([]*KMSKey, error)
	Get(ctx context.Context, id string) (*KMSKey, error)
	Create(ctx context.Context, key *KMSKey) (*KMSKey, error)
	Delete(ctx context.Context, id string) error
}

type kmsSecretService struct {
	client *Client
}

func (k *kmsService) Secret() *kmsSecretService {
	return &kmsSecretService{
		client: k.client,
	}
}

type KMSKey struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
