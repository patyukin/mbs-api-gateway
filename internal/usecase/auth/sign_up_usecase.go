package auth

import (
	"context"
	"fmt"
	"github.com/patyukin/mbs-api-gateway/internal/model"
)

func (uc *UseCase) SignUpV1(ctx context.Context, in model.SignUpV1Request) error {
	dto := model.ToProtoSignUpFromRequest(in)

	_, err := uc.authClient.SignUp(ctx, &dto)
	if err != nil {
		return fmt.Errorf("failed to uc.authClient.SignUp: %w", err)
	}

	return nil
}
