package auth

import (
	"context"
	"fmt"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/rs/zerolog/log"
)

func (uc *UseCase) AuthorizeUser(ctx context.Context, in model.AuthorizeUserV1Request) error {
	pbm := model.ToProtoV1AuthorizeUserRequest(in)

	log.Debug().Msgf("uc.authClient: %v", uc.authClient)

	response, err := uc.authClient.AuthorizeUser(ctx, &pbm)
	if err != nil || response.Error != nil {
		return fmt.Errorf("failed to uc.authClient.Authorize: %w", err)
	}

	return nil
}
