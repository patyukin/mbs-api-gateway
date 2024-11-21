package payment

import (
	"context"
	"fmt"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
	"github.com/rs/zerolog/log"
)

func (u *UseCase) CreatePaymentUseCase(ctx context.Context, in model.CreatePaymentV1Request) (model.CreatePaymentV1Response, *error_v1.ErrorResponse) {
	pbm := model.ToProtoCreatePaymentFromRequest(in)

	log.Debug().Msgf("pbm: %v", &pbm)

	result, err := u.paymentClient.CreatePayment(ctx, &pbm)

	log.Debug().Msgf("result: %v", result)

	if err != nil {
		return model.CreatePaymentV1Response{}, &error_v1.ErrorResponse{
			Code:        500,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to GetLogReport: %v", err),
		}
	}

	if result.Error != nil {
		return model.CreatePaymentV1Response{}, result.Error
	}

	return model.CreatePaymentV1Response{}, nil
}
