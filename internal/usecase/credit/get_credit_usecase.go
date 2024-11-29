package credit

import (
	"context"
	"fmt"

	"github.com/patyukin/mbs-api-gateway/internal/model"
	creditpb "github.com/patyukin/mbs-pkg/pkg/proto/credit_v1"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
)

func (u *UseCase) GetCreditV1UseCase(ctx context.Context, creditID, userID string) (model.GetCreditV1Response, *error_v1.ErrorResponse) {
	response, err := u.creditClient.GetCredit(ctx, &creditpb.GetCreditRequest{
		CreditId: creditID,
		UserId:   userID,
	})
	if err != nil {
		return model.GetCreditV1Response{}, &error_v1.ErrorResponse{
			Code:        500,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to CreateCreditApplication: %v", err),
		}
	}

	if response == nil {
		return model.GetCreditV1Response{}, &error_v1.ErrorResponse{
			Code:        500,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to CreateCreditApplication: %v", err),
		}
	}

	return model.ToModelGetCreditResponse(response.GetCredit()), nil
}
