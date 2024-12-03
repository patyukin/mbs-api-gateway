package credit

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"

	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
)

func (u *UseCase) CreditApplicationConfirmationV1UseCase(ctx context.Context, in model.CreditApplicationConfirmationV1Request, userID string) (model.CreditApplicationConfirmationV1Response, *error_v1.ErrorResponse) {
	mpb := model.ToProtoV1CreditApplicationConfirmationRequest(in, userID)
	response, err := u.creditClient.CreditApplicationConfirmation(ctx, &mpb)
	if err != nil {
		return model.CreditApplicationConfirmationV1Response{}, &error_v1.ErrorResponse{
			Code:        http.StatusInternalServerError,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to CreateCreditApplication: %v", err),
		}
	}

	if response.Error != nil {
		return model.CreditApplicationConfirmationV1Response{}, &error_v1.ErrorResponse{
			Code:        http.StatusInternalServerError,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to CreateCreditApplication: %v", err),
		}
	}

	log.Debug().Msgf("response: %v", response)

	return model.CreditApplicationConfirmationV1Response{Message: response.GetMessage()}, nil
}
