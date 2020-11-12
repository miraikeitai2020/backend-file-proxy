package config

import "github.com/kelseyhightower/envconfig"

const nullStr = ""

// MinioConnConfig packs connection config
type minioConnConfig struct {
	URL       string `envconfig:"MINIO_PATH" default:"localhost:9000"`
	PublicKey string `envconfig:"MINIO_PUBLIC_KEY" default:"Cg2g6f63KGvzm2a623UEGdiPKYTe66Nb"`
	SecretKey string `envconfig:"MINIO_SECRET_KEY" default:"Mf67pN3LsJRabd8j97pk7nxGLq7B3mD8"`
}

// MinioConnParams provides minio connection config
func MinioConnParams() (string, string, string, error) {
	var c minioConnConfig
	if err := envconfig.Process("", &c); err != nil {
		return nullStr, nullStr, nullStr, err
	}
	return c.URL, c.PublicKey, c.SecretKey, nil
}
