package controller

import (
	"github.com/miraikeitai2020/backend-file-proxy/pkg/model/dao/minio"
)

// Controllers manages multiple controllers
type Controllers struct {
	minio *minio.Minio
}

// New controllers client
func New(m *minio.Minio) Controllers {
	return Controllers{m}
}
