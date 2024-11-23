package credit

import (
	"context"
	"fmt"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
)

func (u *UseCase) UpdateCreditApplicationStatusV1UseCase(ctx context.Context, in model.UpdateCreditApplicationStatusV1Request) (model.UpdateCreditApplicationStatusV1Response, *error_v1.ErrorResponse) {
	mpb, err := model.ToProtoV1UpdateCreditApplicationStatusRequest(in)
	if err != nil {
		return model.UpdateCreditApplicationStatusV1Response{}, &error_v1.ErrorResponse{
			Code:        500,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed model.ToProtoV1UpdateCreditApplicationStatusRequest: %v", err),
		}
	}

	response, err := u.creditClient.UpdateCreditApplicationStatus(ctx, &mpb)
	if err != nil {
		return model.UpdateCreditApplicationStatusV1Response{}, &error_v1.ErrorResponse{
			Code:        500,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to CreateCreditApplication: %v", err),
		}
	}

	if response == nil {
		return model.UpdateCreditApplicationStatusV1Response{}, &error_v1.ErrorResponse{
			Code:        500,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to CreateCreditApplication: %v", err),
		}
	}

	return model.UpdateCreditApplicationStatusV1Response{Message: response.Message}, nil
}
