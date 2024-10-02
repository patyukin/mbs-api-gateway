package grpc_client

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/credentials/insecure"
	"time"

	grpcRetry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	traceutils "github.com/opentracing-contrib/go-grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

const (
	backoffLinear = 100 * time.Millisecond
)

func NewGRPCClientServiceConn(port int) (*grpc.ClientConn, error) {
	opts := []grpcRetry.CallOption{
		grpcRetry.WithBackoff(grpcRetry.BackoffLinear(backoffLinear)),
		grpcRetry.WithCodes(codes.NotFound, codes.Aborted),
	}

	clientGRPCConn, err := grpc.NewClient(
		fmt.Sprintf("http://0.0.0.0:%d", port),
		grpc.WithUnaryInterceptor(traceutils.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(grpcRetry.UnaryClientInterceptor(opts...)),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to auth service: %w", err)
	}

	return clientGRPCConn, nil
}
