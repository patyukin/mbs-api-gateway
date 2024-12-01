package auth

import (
	"context"
	"fmt"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
	"net/http"
)

func (uc *UseCase) AddUserRoleV1UseCase(ctx context.Context, in model.AddUserRoleV1Request) (model.AddUserRoleV1Response, *error_v1.ErrorResponse) {
	pbm := model.ToProtoAddUserRoleV1Request(in)

	result, err := uc.authClient.AddUserRole(ctx, &pbm)
	if err != nil {
		return model.AddUserRoleV1Response{}, &error_v1.ErrorResponse{
			Code:        http.StatusInternalServerError,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to uc.authClient.AddUserRole: %v", err),
		}
	}

	if result.GetError() != nil {
		return model.AddUserRoleV1Response{}, result.GetError()
	}

	return model.ToModelAddUserRoleV1Response(result), nil
}
