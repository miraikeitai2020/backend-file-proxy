FROM golang:1.13-alpine AS build-env
WORKDIR /go/src/github.com/miraikeitai2020/backend-file-proxy
COPY ./ ./
RUN go build -o server cmd/main.go

FROM alpine:latest
RUN apk add --no-cache --update ca-certificates
COPY --from=build-env /go/src/github.com/miraikeitai2020/backend-file-proxy/server /usr/local/bin/server

EXPOSE 9000
CMD ["/usr/local/bin/server"]