package credit

import (
	"context"
	"fmt"
	"net/http"

	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
)

func (u *UseCase) CreateCreditApplicationV1UseCase(ctx context.Context, in model.CreateCreditApplicationV1Request, userID string) (model.CreateCreditApplicationV1Response, *error_v1.ErrorResponse) {
	mpb := model.ToProtoV1CreateCreditApplicationRequest(in, userID)
	response, err := u.creditClient.CreateCreditApplication(ctx, &mpb)
	if err != nil {
		return model.CreateCreditApplicationV1Response{}, &error_v1.ErrorResponse{
			Code:        http.StatusInternalServerError,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to CreateCreditApplication: %v", err),
		}
	}

	if response == nil {
		return model.CreateCreditApplicationV1Response{}, &error_v1.ErrorResponse{
			Code:        http.StatusInternalServerError,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to CreateCreditApplication: %v", err),
		}
	}

	if response.GetError() != nil {
		return model.CreateCreditApplicationV1Response{}, &error_v1.ErrorResponse{
			Code:        response.GetError().GetCode(),
			Message:     response.GetError().GetMessage(),
			Description: response.GetError().GetDescription(),
		}
	}

	return model.CreateCreditApplicationV1Response{Message: response.GetMessage()}, nil
}
