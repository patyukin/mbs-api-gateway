package grpc_client

import (
	"fmt"
	grpcopentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGRPCClientServiceConn(host string, port int) (*grpc.ClientConn, error) {
	url := fmt.Sprintf("%s:%d", host, port)
	log.Info().Msgf("url: %s", url)
	clientGRPCConn, err := grpc.NewClient(
		url,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(grpcopentracing.UnaryClientInterceptor()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to auth service: %w", err)
	}

	return clientGRPCConn, nil
}
