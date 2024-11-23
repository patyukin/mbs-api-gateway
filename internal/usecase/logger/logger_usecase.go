package logger

import (
	"context"
	loggerpb "github.com/patyukin/mbs-pkg/pkg/proto/logger_v1"
	"google.golang.org/grpc"
)

type ProtoClient interface {
	GetLogReport(ctx context.Context, in *loggerpb.LogReportRequest, opts ...grpc.CallOption) (*loggerpb.LogReportResponse, error)
}

type UseCase struct {
	loggerClient ProtoClient
}

func New(loggerClient ProtoClient) *UseCase {
	return &UseCase{
		loggerClient: loggerClient,
	}
}
