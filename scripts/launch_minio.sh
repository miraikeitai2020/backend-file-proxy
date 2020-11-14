#!/bin/sh

# Env Var #
CONTAINER_IMAGE=minio/minio
CONTAINER_NAME=minio_s3
MINIO_ACCESS_KEY=$1
MINIO_SECRET_KEY=$2

# launch minio container
echo "launch minio container.\tname = $CONTAINER_NAME"
docker run -p 9000:9000 -d \
    --name $CONTAINER_NAME \
    --rm \
	-e "MINIO_ACCESS_KEY=$MINIO_ACCESS_KEY" \
	-e "MINIO_SECRET_KEY=$MINIO_SECRET_KEY" \
	$CONTAINER_IMAGE server /data
