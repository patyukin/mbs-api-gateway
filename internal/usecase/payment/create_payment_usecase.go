package payment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
	"github.com/rs/zerolog/log"
)

func (u *UseCase) CreatePaymentV1UseCase(ctx context.Context, in model.CreatePaymentV1Request, userID string) (
	model.CreatePaymentV1Response, *error_v1.ErrorResponse,
) {
	pbm := model.ToProtoCreatePaymentFromRequest(in, userID)

	log.Debug().Msgf("pbm: %v", &pbm)

	result, err := u.paymentClient.CreatePayment(ctx, &pbm)

	log.Debug().Msgf("result: %v", result)

	if err != nil {
		return model.CreatePaymentV1Response{}, &error_v1.ErrorResponse{
			Code:        http.StatusInternalServerError,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to GetLogReportV1UseCase: %v", err),
		}
	}

	if result.GetError() != nil {
		return model.CreatePaymentV1Response{}, result.GetError()
	}

	return model.CreatePaymentV1Response{Message: result.GetMessage()}, nil
}
