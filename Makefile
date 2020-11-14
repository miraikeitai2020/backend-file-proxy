GO_CMD=go
GO_RUN=$(GO_CMD) run
GO_BUILD=$(GO_BUILD) build
GO_VET=$(GO_CMD) vet
DOCKER_CMD=docker

MINIO_ACCESS_DEFAULT_KEY=Cg2g6f63KGvzm2a623UEGdiPKYTe66Nb
MINIO_SECRET_DEFAULT_KEY=Mf67pN3LsJRabd8j97pk7nxGLq7B3mD8

PUB_KEY=$(MINIO_ACCESS_DEFAULT_KEY)
SEC_KEY=$(MINIO_SECRET_DEFAULT_KEY)

all:
	$(GO_RUN) scripts/keys/keyGen.go
build:
	$(GO_BUILD)
docker:
	$(DOCKER_CMD) ./ -t miraikeitai2020/file-proxy:1.0.0
test:
	$(GO_RUN) scripts/postObject.go
minio-run:
	sh scripts/launch_minio.sh $(PUB_KEY) $(SEC_KEY)
minio-down:
	docker stop minio_s3