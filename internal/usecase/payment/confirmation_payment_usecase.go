package payment

import (
	"context"
	"fmt"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
	"github.com/rs/zerolog/log"
)

func (u *UseCase) ConfirmationPaymentV1UseCase(ctx context.Context, in model.ConfirmationPaymentV1Request) (model.VerifyPaymentV1Response, *error_v1.ErrorResponse) {
	pbm := model.ToProtoVerifyPaymentFromRequest(in)
	log.Debug().Msgf("pbm: %v", &pbm)

	result, err := u.paymentClient.ConfirmationPayment(ctx, &pbm)
	if err != nil {
		return model.VerifyPaymentV1Response{}, &error_v1.ErrorResponse{
			Code:        500,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to GetLogReportV1UseCase: %v", err),
		}
	}

	if result != nil {
		return model.VerifyPaymentV1Response{}, result.Error
	}

	return model.VerifyPaymentV1Response{}, nil
}
