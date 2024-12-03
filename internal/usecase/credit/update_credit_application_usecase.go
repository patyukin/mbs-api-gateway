package credit

import (
	"context"
	"fmt"
	"net/http"

	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
)

func (u *UseCase) UpdateCreditApplicationStatusV1UseCase(ctx context.Context, in model.UpdateCreditApplicationStatusV1Request, applicationID string) (model.UpdateCreditApplicationStatusV1Response, *error_v1.ErrorResponse) {
	mpb, err := model.ToProtoV1UpdateCreditApplicationStatusRequest(in, applicationID)
	if err != nil {
		return model.UpdateCreditApplicationStatusV1Response{}, &error_v1.ErrorResponse{
			Code:        http.StatusInternalServerError,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed model.ToProtoV1UpdateCreditApplicationStatusRequest: %v", err),
		}
	}

	response, err := u.creditClient.UpdateCreditApplicationSolution(ctx, &mpb)
	if err != nil {
		return model.UpdateCreditApplicationStatusV1Response{}, &error_v1.ErrorResponse{
			Code:        http.StatusInternalServerError,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to CreateCreditApplication: %v", err),
		}
	}

	if response.Error != nil {
		return model.UpdateCreditApplicationStatusV1Response{}, response.GetError()
	}

	return model.UpdateCreditApplicationStatusV1Response{Message: response.GetMessage()}, nil
}
