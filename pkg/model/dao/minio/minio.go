package minio

import (
	"context"
	"fmt"
	"os"

	"github.com/minio/minio-go"
	"github.com/minio/minio-go/pkg/credentials"
	"github.com/miraikeitai2020/backend-file-proxy/config"
	"github.com/miraikeitai2020/backend-file-proxy/pkg/model/service/log"
)

var BucketList []string

func init() {
	var err error
	BucketList, err = config.MinioBucketList()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

// Minio packs repositories
type Minio struct {
	Config  minioConfig
	Client  *minio.Client
	Buckets Repositories
}

// Repositories is...
type Repositories struct {
	Detour repository
}

type repository interface {
	Get(string) (*minio.Object, int64, error)
	Add(string) (int64, error)
}

type minioConfig struct {
	URL       string `json:"endpoint"`
	PublicKey string `json:"publicKey"`
	SecretKey string `json:"secretKey"`
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
		Config:  minioConfig{url, pk, sk},
		Client:  client,
		Buckets: Repositories{detour},
	}, nil
}

// UpdateConfig can change connection infomation
func (m *Minio) UpdateConfig(url, pk, sk string) *Minio {
	client, err := minio.New(url, &minio.Options{
		Creds:  credentials.NewStaticV4(pk, sk, ""),
		Secure: false,
	})
	if err != nil {
		return nil
	}

	detour := newDetourClient(client)
	m.Buckets = Repositories{detour}
	m.Config = minioConfig{url, pk, sk}

	return m
}

// CreateBucket create bucket in minio
func (m *Minio) CreateBucket(bucket string) error {
	cxt := context.Background()
	if err := m.Client.MakeBucket(cxt, bucket, minio.MakeBucketOptions{}); err != nil {
		exists, errBucketExists := m.Client.BucketExists(cxt, bucket)
		if errBucketExists != nil {
			return errBucketExists
		}
		if exists {
			msg := fmt.Sprintf("%s is already exists", bucket)
			log.Info(msg)
			return nil
		}
		return err
	}

	msg := fmt.Sprintf("Successfully created %s", bucket)
	log.Info(msg)
	return nil
}
