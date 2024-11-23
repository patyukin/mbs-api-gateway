package credit

import (
	"context"
	"fmt"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
)

func (u *UseCase) CreateCreditApplicationV1UseCase(ctx context.Context, in model.CreateCreditApplicationV1Request, userID string) (model.CreateCreditApplicationV1Response, *error_v1.ErrorResponse) {
	mpb := model.ToProtoV1CreateCreditApplicationRequest(in, userID)
	response, err := u.creditClient.CreateCreditApplication(ctx, &mpb)
	if err != nil {
		return model.CreateCreditApplicationV1Response{}, &error_v1.ErrorResponse{
			Code:        500,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to CreateCreditApplication: %v", err),
		}
	}

	if response == nil {
		return model.CreateCreditApplicationV1Response{}, &error_v1.ErrorResponse{
			Code:        500,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to CreateCreditApplication: %v", err),
		}
	}

	if response.Error != nil {
		return model.CreateCreditApplicationV1Response{}, &error_v1.ErrorResponse{
			Code:        response.Error.Code,
			Message:     response.Error.Message,
			Description: response.Error.Description,
		}
	}

	return model.CreateCreditApplicationV1Response{}, nil
}
