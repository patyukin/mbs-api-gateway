package auth

import (
	"context"
	"fmt"
	"github.com/patyukin/mbs-api-gateway/internal/model"
)

func (uc *UseCase) SignInV1(ctx context.Context, in model.SignInV1Request) (model.SignInV1Response, error) {
	dto := model.ToProtoSignInFromRequest(in)

	_, err := uc.authClient.SignIn(ctx, &dto)
	if err != nil {
		return model.SignInV1Response{}, fmt.Errorf("failed to uc.authClient.SignUp: %w", err)
	}

	return model.SignInV1Response{}, nil
}
