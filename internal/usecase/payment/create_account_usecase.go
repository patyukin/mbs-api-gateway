package payment

import (
	"context"
	"fmt"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
	"github.com/rs/zerolog/log"
)

func (u *UseCase) CreateAccountUseCase(ctx context.Context, in model.CreateAccountV1Request) (model.CreateAccountV1Response, *error_v1.ErrorResponse) {
	pbm := model.ToProtoCreateAccountFromRequest(in)
	result, err := u.paymentClient.CreateAccount(ctx, &pbm)

	log.Debug().Msgf("result: %v", result)

	if err != nil {
		return model.CreateAccountV1Response{}, &error_v1.ErrorResponse{
			Code:        500,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to GetLogReport: %v", err),
		}
	}

	if result.Error != nil {
		return model.CreateAccountV1Response{}, result.Error
	}

	return model.CreateAccountV1Response{}, nil
}
