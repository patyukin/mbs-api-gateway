package auth

import (
	"context"
	"fmt"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
)

func (uc *UseCase) SignInV1UseCase(ctx context.Context, in model.SignInV1Request) (model.SignInV1Response, *error_v1.ErrorResponse) {
	dto := model.ToProtoV1SignInRequest(in)

	response, err := uc.authClient.SignIn(ctx, &dto)
	if err != nil {
		return model.SignInV1Response{}, &error_v1.ErrorResponse{
			Code:        500,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to uc.authClient.SignIn: %v", err),
		}
	}

	if response.Error != nil {
		response.Error.Description = fmt.Sprintf("failed to uc.authClient.SignIn: %v", response.Error.Description)
		return model.SignInV1Response{}, response.Error
	}

	return model.SignInV1Response{Message: "Код отправлен в telegram bot"}, nil
}
