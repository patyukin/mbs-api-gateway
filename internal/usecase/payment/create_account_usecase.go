package payment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
	"github.com/rs/zerolog/log"
)

func (u *UseCase) CreateAccountV1UseCase(ctx context.Context, in model.CreateAccountV1Request, userID string) (
	model.CreateAccountV1Response, *error_v1.ErrorResponse,
) {
	pbm := model.ToProtoCreateAccountFromRequest(in, userID)
	result, err := u.paymentClient.CreateAccount(ctx, &pbm)

	log.Debug().Msgf("result: %v", result)

	if err != nil {
		return model.CreateAccountV1Response{}, &error_v1.ErrorResponse{
			Code:        http.StatusInternalServerError,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to GetLogReportV1UseCase: %v", err),
		}
	}

	if result.GetError() != nil {
		return model.CreateAccountV1Response{}, result.GetError()
	}

	return model.CreateAccountV1Response{Message: result.GetMessage()}, nil
}
