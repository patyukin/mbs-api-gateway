package auth

import (
	"context"
	"fmt"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
	"net/http"
)

func (uc *UseCase) GetUserByIDV1UseCase(ctx context.Context, in model.GetUserByIDV1Request) (model.GetUserByIDV1Response, *error_v1.ErrorResponse) {
	pbm := model.ToProtoGetUserByIDV1Request(in)
	response, err := uc.authClient.GetUserByID(ctx, &pbm)
	if err != nil {
		return model.GetUserByIDV1Response{}, &error_v1.ErrorResponse{
			Code:        http.StatusInternalServerError,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to uc.authClient.GetUserByIDV1UseCase: %v", err),
		}
	}

	if response.Error != nil {
		return model.GetUserByIDV1Response{}, response.Error
	}

	return model.ToModelGetUserByIDV1Response(response), nil
}
