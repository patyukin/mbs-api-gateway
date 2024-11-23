package credit

import (
	"context"
	"fmt"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
)

func (u *UseCase) CreateCreditV1UseCase(ctx context.Context, in model.CreateCreditV1Request) (model.CreateCreditV1Response, *error_v1.ErrorResponse) {
	mpb := model.ToProtoV1CreateCreditRequest(in)
	response, err := u.creditClient.CreateCredit(ctx, &mpb)
	if err != nil {
		return model.CreateCreditV1Response{}, &error_v1.ErrorResponse{
			Code:        500,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to CreateCreditApplication: %v", err),
		}
	}

	if response == nil {
		return model.CreateCreditV1Response{}, &error_v1.ErrorResponse{
			Code:        500,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to CreateCreditApplication: %v", err),
		}
	}

	if response.Error != nil {
		return model.CreateCreditV1Response{}, &error_v1.ErrorResponse{
			Code:        response.Error.Code,
			Message:     response.Error.Message,
			Description: response.Error.Description,
		}
	}

	return model.CreateCreditV1Response{Message: response.Message}, nil
}
