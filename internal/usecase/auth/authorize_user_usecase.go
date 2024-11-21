package auth

import (
	"context"
	"fmt"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/rs/zerolog/log"
)

func (uc *UseCase) Authorize(ctx context.Context, in model.AuthorizeRequest) error {
	pbm := model.ToProtoAuthorizeFromRequest(in)

	log.Debug().Msgf("uc.authClient: %v", uc.authClient)

	response, err := uc.authClient.AuthorizeUser(ctx, &pbm)
	if err != nil || response.Error != nil {
		return fmt.Errorf("failed to uc.authClient.Authorize: %w", err)
	}

	return nil
}
