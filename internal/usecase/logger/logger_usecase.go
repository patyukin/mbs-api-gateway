package logger

import (
	"context"
	"fmt"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
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

func (u *UseCase) GetLogReport(ctx context.Context, in model.GetLogReportV1Request) *error_v1.ErrorResponse {
	pbm, err := model.ToProtoLogReportFromRequest(in)
	if err != nil {
		return &error_v1.ErrorResponse{
			Code:        400,
			Message:     err.Error(),
			Description: fmt.Sprintf("failed to ToProtoLogReportFromRequest: %v", err),
		}
	}

	result, err := u.loggerClient.GetLogReport(ctx, &pbm)
	if err != nil {
		return &error_v1.ErrorResponse{
			Code:        500,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to GetLogReport: %v", err),
		}
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}
