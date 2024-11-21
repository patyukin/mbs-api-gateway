package credit

import (
	"context"
	"fmt"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	creditpb "github.com/patyukin/mbs-pkg/pkg/proto/credit_v1"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
)

func (u *UseCase) GetPaymentScheduleUseCase(ctx context.Context, userID, creditID string) (model.GetPaymentScheduleV1Response, *error_v1.ErrorResponse) {
	response, err := u.creditClient.GetPaymentSchedule(ctx, &creditpb.GetPaymentScheduleRequest{
		UserId:   userID,
		CreditId: creditID,
	})
	if err != nil {
		return model.GetPaymentScheduleV1Response{}, &error_v1.ErrorResponse{
			Code:        500,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to CreateCreditApplication: %v", err),
		}
	}

	if response == nil {
		return model.GetPaymentScheduleV1Response{}, &error_v1.ErrorResponse{
			Code:        500,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to CreateCreditApplication: %v", err),
		}
	}

	return model.ToModelGetPaymentScheduleResponse(response), nil

}
