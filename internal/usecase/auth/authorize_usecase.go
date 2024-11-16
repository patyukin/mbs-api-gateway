package auth

import (
	"context"
	"fmt"
	"github.com/patyukin/mbs-api-gateway/internal/model"
)

func (uc *UseCase) Authorize(ctx context.Context, in model.AuthorizeRequest) error {
	pbm := model.ToProtoAuthorizeFromRequest(in)
	_, err := uc.authClient.Authorize(ctx, &pbm)
	if err != nil {
		return fmt.Errorf("failed to uc.authClient.SignInVerify: %w", err)
	}

	return nil
}
