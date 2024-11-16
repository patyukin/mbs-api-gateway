FROM golang:1.23.2-alpine3.20 AS build-env

RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY . /app
WORKDIR /app

RUN go mod download
RUN go mod tidy
RUN go build -gcflags="all=-N -l" -o bin/api_gateway cmd/api_gateway/main.go

FROM alpine:3.20

EXPOSE 5002 40001

WORKDIR /root

COPY --from=build-env /go/bin/dlv .
COPY --from=build-env /app/bin/api_gateway ./api_gateway
ENV YAML_CONFIG_FILE_PATH=config.yaml

CMD ["/root/dlv", "--listen=:40001", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/root/api_gateway"]
