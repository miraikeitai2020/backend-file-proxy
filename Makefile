s3:
	docker run -p 9000:9000 -d \
		--name minio_mirai \
		--rm \
		-e "MINIO_ACCESS_KEY=Cg2g6f63KGvzm2a623UEGdiPKYTe66Nb" \
		-e "MINIO_SECRET_KEY=Mf67pN3LsJRabd8j97pk7nxGLq7B3mD8" \
		minio/minio server /data
s3-clean:
	docker stop minio_mirai