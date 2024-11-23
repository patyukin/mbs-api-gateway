package auth

import (
	"context"
	"fmt"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
)

func (uc *UseCase) RefreshTokenV1UseCase(ctx context.Context, in model.RefreshTokenV1Request) (model.RefreshTokenV1Response, *error_v1.ErrorResponse) {
	pbm := model.ToProtoV1RefreshTokenRequest(in)
	response, err := uc.authClient.RefreshToken(ctx, &pbm)
	if err != nil {
		return model.RefreshTokenV1Response{}, &error_v1.ErrorResponse{
			Code:        500,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to uc.authClient.Authorize: %v", err),
		}
	}

	if response.Error != nil {
		return model.RefreshTokenV1Response{}, response.Error
	}

	if response.AccessToken == "" {
		return model.RefreshTokenV1Response{}, &error_v1.ErrorResponse{
			Code:        500,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to uc.authClient.Authorize: %v", err),
		}
	}

	return model.ToModelRefreshTokenV1Response(response), nil
}
