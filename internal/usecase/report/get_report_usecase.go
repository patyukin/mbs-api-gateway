package report

import (
	"context"
	"fmt"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
	"github.com/rs/zerolog/log"
)

func (u *UseCase) GetUserReportV1UseCase(ctx context.Context, in model.GetUserReportV1Request) (model.GetUserReportV1Response, *error_v1.ErrorResponse) {
	pbm := model.ToProtoGetUserReport(in)
	result, err := u.reportClient.GetUserReport(ctx, &pbm)
	if err != nil {
		return model.GetUserReportV1Response{}, &error_v1.ErrorResponse{
			Code:        500,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to GetLogReportV1UseCase: %v", err),
		}
	}

	if result.Error != nil {
		return model.GetUserReportV1Response{}, result.Error
	}

	log.Debug().Msgf("result.FileUrl: %v", result.Message)

	return model.GetUserReportV1Response{Message: result.Message}, nil

}
