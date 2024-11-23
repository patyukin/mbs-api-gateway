package logger

import (
	"context"
	"fmt"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
	"github.com/rs/zerolog/log"
)

func (u *UseCase) GetLogReportV1UseCase(ctx context.Context, in model.GetLogReportV1Request) (model.GetLogReportV1Response, *error_v1.ErrorResponse) {
	pbm, err := model.ToProtoLogReportFromRequest(in)
	if err != nil {
		return model.GetLogReportV1Response{}, &error_v1.ErrorResponse{
			Code:        400,
			Message:     err.Error(),
			Description: fmt.Sprintf("failed to ToProtoLogReportFromRequest: %v", err),
		}
	}

	result, err := u.loggerClient.GetLogReport(ctx, &pbm)
	if err != nil {
		return model.GetLogReportV1Response{}, &error_v1.ErrorResponse{
			Code:        500,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to GetLogReportV1UseCase: %v", err),
		}
	}

	if result.Error != nil {
		return model.GetLogReportV1Response{}, result.Error
	}

	log.Debug().Msgf("result.FileUrl: %v", result.Message)

	return model.GetLogReportV1Response{FileUrl: result.Message}, nil
}
