package credit

import (
	"context"
	"fmt"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
)

func (u *UseCase) CreditApplicationConfirmationUseCase(ctx context.Context, in model.CreditApplicationConfirmationV1Request, userID string) (model.CreditApplicationConfirmationV1Response, *error_v1.ErrorResponse) {
	mpb := model.ToProtoV1CreditApplicationConfirmationRequest(in, userID)
	response, err := u.creditClient.CreditApplicationConfirmation(ctx, &mpb)
	if err != nil {
		return model.CreditApplicationConfirmationV1Response{}, &error_v1.ErrorResponse{
			Code:        500,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to CreateCreditApplication: %v", err),
		}
	}

	if response == nil {
		return model.CreditApplicationConfirmationV1Response{}, &error_v1.ErrorResponse{
			Code:        500,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to CreateCreditApplication: %v", err),
		}
	}

	return model.CreditApplicationConfirmationV1Response{
		ApplicationID: response.ApplicationId,
		Message:       response.Message,
		Status:        response.Status.String(),
	}, nil
}
