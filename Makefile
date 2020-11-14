GO_CMD=go run
DOCKER_CMD=docker

MINIO_ACCESS_DEFAULT_KEY=Cg2g6f63KGvzm2a623UEGdiPKYTe66Nb
MINIO_SECRET_DEFAULT_KEY=Mf67pN3LsJRabd8j97pk7nxGLq7B3mD8

PUB_KEY=$(MINIO_ACCESS_DEFAULT_KEY)
SEC_KEY=$(MINIO_SECRET_DEFAULT_KEY)

all:
	$(GO_CMD) scripts/keyGen.go
minio-run:
	sh scripts/launch_minio.sh $(PUB_KEY) $(SEC_KEY)
minio-down:
	docker stop minio_s3