package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
)

func (uc *UseCase) SignUpV1UseCase(ctx context.Context, in model.SignUpV1Request) (model.SignUpV1Response, *error_v1.ErrorResponse) {
	dto := model.ToProtoV1SignUpRequest(in)
	result, err := uc.authClient.SignUp(ctx, &dto)
	if err != nil {
		return model.SignUpV1Response{}, &error_v1.ErrorResponse{
			Code:        http.StatusInternalServerError,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to uc.authClient.SignUp: %v", err),
		}
	}

	return model.SignUpV1Response{Message: result.GetMessage()}, nil
}
