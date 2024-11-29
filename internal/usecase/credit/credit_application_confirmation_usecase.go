package credit

import (
	"context"
	"fmt"
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

	if response == nil {
		return model.CreditApplicationConfirmationV1Response{}, &error_v1.ErrorResponse{
			Code:        http.StatusInternalServerError,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to CreateCreditApplication: %v", err),
		}
	}

	return model.CreditApplicationConfirmationV1Response{
		ApplicationID: response.GetApplicationId(),
		Message:       response.GetMessage(),
		Status:        response.GetStatus().String(),
	}, nil
}
