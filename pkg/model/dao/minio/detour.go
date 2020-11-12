package minio

import (
	"context"

	"github.com/minio/minio-go"
)

const (
	detourBucketName = "detour"
)

type detourRepository struct {
	Client *minio.Client
	Bucket string
}

func newDetourClient(client *minio.Client) repository {
	return &detourRepository{Client: client, Bucket: detourBucketName}
}

func (r *detourRepository) Get(fileName string) (*minio.Object, int64, error) {
	object, err := r.Client.GetObject(context.Background(), r.Bucket, fileName, minio.GetObjectOptions{})
	if err != nil {
		return nil, -1, err
	}

	stat, err := object.Stat()
	if err != nil {
		return nil, -1, err
	}

	return object, stat.Size, nil
}

func (r *detourRepository) Add(fileName string) error {
	return nil
}
