package minio

import (
	"github.com/minio/minio-go"
	"github.com/minio/minio-go/pkg/credentials"
	"github.com/miraikeitai2020/backend-file-proxy/config"
)

// Minio packs repositories
type Minio struct {
	Client  *minio.Client
	Buckets Repositories
}

// Repositories is...
type Repositories struct {
	Detour repository
}

type repository interface {
	Get(string) (*minio.Object, int64, error)
	Add(string) error
}

type connectionConfig struct {
	URL       string
	PublicKey string
	SecretKey string
	TLS       bool
}

// New returns Clients structer
func New() (*Minio, error) {
	// Get connection config
	url, pk, sk, err := config.MinioConnParams()
	if err != nil {
		return nil, err
	}

	// Connect minio service
	client, err := minio.New(url, &minio.Options{
		Creds:  credentials.NewStaticV4(pk, sk, ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}

	// init repositories
	detour := newDetourClient(client)

	return &Minio{
		Client:  client,
		Buckets: Repositories{detour},
	}, nil
}

// UpdateConfig can change connection infomation
func (m *Minio) UpdateConfig(url, pk, sk string) error {
	client, err := minio.New(url, &minio.Options{
		Creds:  credentials.NewStaticV4(pk, sk, ""),
		Secure: false,
	})
	if err != nil {
		return err
	}
	m.Client = client
	return nil
}
