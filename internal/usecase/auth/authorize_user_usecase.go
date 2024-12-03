package auth

import (
	"context"
	"fmt"

	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/rs/zerolog/log"
)

func (uc *UseCase) AuthorizeUserV1UseCase(ctx context.Context, in model.AuthorizeUserV1Request) error {
	pbm := model.ToProtoV1AuthorizeUserRequest(in)
	response, err := uc.authClient.AuthorizeUser(ctx, &pbm)

	log.Debug().Msgf("response: %v", response)
	log.Debug().Msgf("err: %v", err)

	if err != nil || response.GetError() != nil {
		return fmt.Errorf("failed to uc.authClient.AuthorizeUser: %w", err)
	}

	return nil
}
