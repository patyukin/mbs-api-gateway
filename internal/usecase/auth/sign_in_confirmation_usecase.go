package auth

import (
	"context"
	"fmt"

	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
)

func (uc *UseCase) SignInVerifyV1(ctx context.Context, in model.SignInConfirmationV1Request) (model.SignInConfirmationV1Response, *error_v1.ErrorResponse) {
	pbm := model.ToProtoV1SignInConfirmationRequest(in)
	result, err := uc.authClient.SignInConfirmation(ctx, &pbm)
	if err != nil {
		return model.SignInConfirmationV1Response{}, &error_v1.ErrorResponse{
			Code:        500,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to uc.authClient.SignInVerify: %v", err),
		}
	}

	if result.GetError() != nil {
		return model.SignInConfirmationV1Response{}, result.GetError()
	}

	return model.ToModelSignInConfirmationV1Response(result), nil
}
