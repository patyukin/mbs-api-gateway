package credit

import (
	"context"
	"fmt"
	"net/http"

	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
)

func (u *UseCase) CreateCreditV1UseCase(ctx context.Context, in model.CreateCreditV1Request, userID string) (model.CreateCreditV1Response, *error_v1.ErrorResponse) {
	mpb := model.ToProtoV1CreateCreditRequest(in, userID)
	response, err := u.creditClient.CreateCredit(ctx, &mpb)
	if err != nil {
		return model.CreateCreditV1Response{}, &error_v1.ErrorResponse{
			Code:        http.StatusInternalServerError,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to CreateCreditApplication: %v", err),
		}
	}

	if response.GetError() != nil {
		return model.CreateCreditV1Response{}, &error_v1.ErrorResponse{
			Code:        response.GetError().GetCode(),
			Message:     response.GetError().GetMessage(),
			Description: fmt.Sprintf("failed to CreateCreditApplication: %v", response.GetError().GetDescription()),
		}
	}

	return model.CreateCreditV1Response{Message: response.GetMessage()}, nil
}
