package auth

import (
	"context"
	"fmt"
	"github.com/patyukin/mbs-api-gateway/internal/model"
)

func (uc *UseCase) SignInV1(ctx context.Context, in model.SignInV1Request) (model.SignInV1Response, error) {
	dto := model.ToProtoSignInFromRequest(in)

	response, err := uc.authClient.SignIn(ctx, &dto)
	if err != nil {
		return model.SignInV1Response{}, fmt.Errorf("failed to uc.authClient.SignIn: %w", err)
	}

	if response.Error != nil {
		response.Error.Description = fmt.Sprintf("failed to uc.authClient.SignIn: %v", response.Error.Description)
		return model.SignInV1Response{Error: response.Error}, nil
	}

	return model.SignInV1Response{Message: "Код отправлен в telegram bot"}, nil
}
