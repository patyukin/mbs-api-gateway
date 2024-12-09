package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
)

func (uc *UseCase) SignInV1UseCase(ctx context.Context, in model.SignInV1Request) (model.SignInV1Response, *error_v1.ErrorResponse) {
	dto := model.ToProtoV1SignInRequest(in)

	response, err := uc.authClient.SignIn(ctx, &dto)
	if err != nil {
		return model.SignInV1Response{}, &error_v1.ErrorResponse{
			Code:        http.StatusInternalServerError,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to uc.authClient.SignIn: %v", err),
		}
	}

	if response.GetError() != nil {
		response.Error.Description = fmt.Sprintf("failed to uc.authClient.SignIn: %v", response.GetError().GetDescription())
		return model.SignInV1Response{}, response.GetError()
	}

	return model.SignInV1Response{Message: "Код отправлен в telegram bot"}, nil
}
