FROM golang:1.13-alpine AS build-env
WORKDIR /go/src/github.com/miraikeitai2020/backend-file-proxy
COPY ./ ./
RUN go build -o server cmd/main.go

FROM alpine:latest
WORKDIR /usr/local/bin
RUN apk add --no-cache --update ca-certificates
COPY --from=build-env /go/src/github.com/miraikeitai2020/backend-file-proxy/server /usr/local/bin/server
COPY --from=build-env /go/src/github.com/miraikeitai2020/backend-file-proxy/config/bucket.json /usr/local/bin/config/bucket.json

EXPOSE 8080
CMD ["/usr/local/bin/server"]