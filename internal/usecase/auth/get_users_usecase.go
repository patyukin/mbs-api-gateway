package auth

import (
	"context"
	"fmt"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
	"net/http"
)

func (uc *UseCase) GetUsersV1UseCase(ctx context.Context, in model.GetUsersV1Request) (model.GetUsersV1Response, *error_v1.ErrorResponse) {
	pbm := model.ToProtoGetUsersV1Request(in)

	result, err := uc.authClient.GetUsers(ctx, &pbm)
	if err != nil {
		return model.GetUsersV1Response{}, &error_v1.ErrorResponse{
			Code:        http.StatusInternalServerError,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to uc.authClient.GetUsers: %v", err),
		}
	}

	if result.GetError() != nil {
		return model.GetUsersV1Response{}, result.GetError()
	}

	return model.ToModelGetUsersV1Response(result), nil
}
