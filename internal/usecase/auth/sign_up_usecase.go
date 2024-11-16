package auth

import (
	"context"
	"fmt"
	"github.com/patyukin/mbs-api-gateway/internal/model"
)

func (uc *UseCase) SignUpV1(ctx context.Context, in model.SignUpV1Request) (model.SignUpV1Response, error) {
	dto := model.ToProtoSignUpFromRequest(in)

	result, err := uc.authClient.SignUp(ctx, &dto)
	if err != nil {
		return model.SignUpV1Response{}, fmt.Errorf("failed to uc.authClient.SignUp: %w", err)
	}

	return model.SignUpV1Response{Message: result.Message}, nil
}
