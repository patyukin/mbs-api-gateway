FROM golang:1.23.1-alpine3.20 AS builder

COPY . /app
WORKDIR /app

RUN go mod download
RUN go mod tidy
RUN go build -o ./bin/api_gateway cmd/api_gateway/main.go

FROM alpine3.20

WORKDIR /app
COPY --from=builder /app/bin/api_gateway .
ENV YAML_CONFIG_FILE_PATH=config.yaml
COPY migrations migrations

CMD ["./api_gateway"]
