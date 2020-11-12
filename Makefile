s3:
	docker run -p 9000:9000 -d \
		--name minio_mirai \
		--rm \
		-e "MINIO_ACCESS_KEY=AKIAIOSFODNN7EXAMPLE" \
		-e "MINIO_SECRET_KEY=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY" \
		minio/minio server /data
s3-clean:
	docker stop minio_mirai