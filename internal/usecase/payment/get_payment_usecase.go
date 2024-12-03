package payment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
)

func (u *UseCase) GetPaymentV1UseCase(ctx context.Context, in model.GetPaymentV1Request, userID string) (model.GetPaymentV1Response, *error_v1.ErrorResponse) {
	pbm := model.ToProtoGetPaymentFromRequest(in, userID)
	result, err := u.paymentClient.GetPayment(ctx, &pbm)
	if err != nil {
		return model.GetPaymentV1Response{}, &error_v1.ErrorResponse{
			Code:        http.StatusInternalServerError,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to GetLogReportV1UseCase: %v", err),
		}
	}

	if result.GetError() != nil {
		return model.GetPaymentV1Response{}, result.GetError()
	}

	payment, err := model.ToModelGetPaymentV1Response(result)
	if err != nil {
		return model.GetPaymentV1Response{}, &error_v1.ErrorResponse{
			Code:        http.StatusInternalServerError,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to GetLogReportV1UseCase: %v", err),
		}
	}

	return payment, nil
}
