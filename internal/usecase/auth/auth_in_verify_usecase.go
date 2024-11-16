package auth

import (
	"context"
	"fmt"
	"github.com/patyukin/mbs-api-gateway/internal/model"
)

func (uc *UseCase) SignInVerify(ctx context.Context, in model.SignInVerifyV1Request) (model.SignInVerifyV1Response, error) {
	siv := model.ToProtoSignInVerifyFromRequest(in)
	tokens, err := uc.authClient.SignInVerify(ctx, &siv)
	if err != nil {
		return model.SignInVerifyV1Response{}, fmt.Errorf("failed to uc.authClient.SignInVerify: %w", err)
	}

	if tokens == nil {
		return model.SignInVerifyV1Response{}, fmt.Errorf("failed to get tokens")
	}

	return model.FromProtoSignInVerifyToResponse(tokens), nil
}
