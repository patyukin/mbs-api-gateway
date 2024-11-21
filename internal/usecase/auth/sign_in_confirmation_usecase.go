package auth

import (
	"context"
	"fmt"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
)

func (uc *UseCase) SignInVerifyV1(ctx context.Context, in model.SignInVerifyV1Request) (model.SignInVerifyV1Response, *error_v1.ErrorResponse) {
	pbm := model.ToProtoSignInVerifyFromRequest(in)
	result, err := uc.authClient.SignInConfirmation(ctx, &pbm)
	if err != nil {
		return model.SignInVerifyV1Response{}, &error_v1.ErrorResponse{
			Code:        500,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to uc.authClient.SignInVerify: %v", err),
		}
	}

	if result.Error != nil {
		return model.SignInVerifyV1Response{}, result.Error
	}

	return model.FromProtoSignInVerifyToResponse(result), nil
}
