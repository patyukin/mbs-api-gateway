package auth

import (
	"context"
	"fmt"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
)

func (uc *UseCase) SignInConfirmation(ctx context.Context, in model.SignInConfirmationV1Request) (model.SignInConfirmationV1Response, *error_v1.ErrorResponse) {
	pbm := model.ToProtoV1SignInConfirmationRequest(in)
	tokens, err := uc.authClient.SignInConfirmation(ctx, &pbm)
	if err != nil {
		return model.SignInConfirmationV1Response{}, &error_v1.ErrorResponse{
			Code:        500,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to uc.authClient.SignInConfirmation: %v", err),
		}
	}

	if tokens == nil {
		return model.SignInConfirmationV1Response{}, &error_v1.ErrorResponse{
			Code:        500,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to uc.authClient.SignInConfirmation: %v", err),
		}
	}

	if tokens.Error != nil {
		return model.SignInConfirmationV1Response{}, tokens.Error
	}

	return model.ToModelSignInConfirmationV1Response(tokens), nil
}
