package credit

import (
	"context"
	"fmt"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	creditpb "github.com/patyukin/mbs-pkg/pkg/proto/credit_v1"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
)

func (u *UseCase) GetListUserCreditsV1UseCase(ctx context.Context, in model.GetListUserCreditsV1Request) (model.GetListUserCreditsV1Response, *error_v1.ErrorResponse) {
	response, err := u.creditClient.GetListUserCredits(ctx, &creditpb.GetListUserCreditsRequest{
		UserId: in.UserID,
		Limit:  in.Limit,
		Page:   in.Page,
	})
	if err != nil {
		return model.GetListUserCreditsV1Response{}, &error_v1.ErrorResponse{
			Code:        500,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to CreateCreditApplication: %v", err),
		}
	}

	if response == nil {
		return model.GetListUserCreditsV1Response{}, &error_v1.ErrorResponse{
			Code:        500,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to CreateCreditApplication: %v", err),
		}
	}

	return model.ToModelGetListUserCreditsResponse(response), nil
}
