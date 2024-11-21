package model

import (
	"github.com/patyukin/mbs-pkg/pkg/mapping/creditmapper"
	creditpb "github.com/patyukin/mbs-pkg/pkg/proto/credit_v1"
)

func ToProtoCreateCreditApplicationRequest(in CreateCreditApplicationV1Request, userID string) creditpb.CreateCreditApplicationRequest {
	return creditpb.CreateCreditApplicationRequest{
		UserId:          userID,
		RequestedAmount: in.RequestedAmount,
		InterestRate:    in.InterestRate,
		StartDate:       in.StartDate,
		EndDate:         in.EndDate,
		Description:     in.Description,
	}
}

func ToProtoCreditApplicationConfirmationRequest(in CreditApplicationConfirmationV1Request, userID string) creditpb.CreditApplicationConfirmationRequest {
	return creditpb.CreditApplicationConfirmationRequest{
		Code:   in.Code,
		UserId: userID,
	}
}

func ToModelGetCreditApplicationResponse(in *creditpb.GetCreditApplicationResponse) GetCreditApplicationV1Response {
	return GetCreditApplicationV1Response{
		ApplicationID:  in.ApplicationId,
		ApprovedAmount: in.ApprovedAmount,
		Status:         in.Status.String(),
		DecisionDate:   in.DecisionDate,
		Message:        in.Message,
	}
}

func ToProtoUpdateCreditApplicationStatusRequest(in UpdateCreditApplicationStatusV1Request) (creditpb.UpdateCreditApplicationStatusRequest, error) {
	status, err := creditmapper.StringToEnumCreditApplicationStatus(in.Status)
	if err != nil {
		return creditpb.UpdateCreditApplicationStatusRequest{}, err
	}

	return creditpb.UpdateCreditApplicationStatusRequest{
		ApplicationId: in.ApplicationID,
		Status:        status,
		DecisionNotes: in.DecisionNotes,
	}, nil
}

func ToModelCredit(in *creditpb.Credit) Credit {
	return Credit{
		CreditID:        in.CreditId,
		UserID:          in.UserId,
		Amount:          in.Amount,
		InterestRate:    in.InterestRate,
		RemainingAmount: in.RemainingAmount,
		Status:          in.Status.String(),
		StartDate:       in.StartDate,
		EndDate:         in.EndDate,
		Description:     in.Description,
	}
}

func ToModelGetCreditResponse(in *creditpb.Credit) GetCreditV1Response {
	return GetCreditV1Response{
		Credit: ToModelCredit(in),
	}
}

func ToModelsGetListUserCreditsResponse(in []*creditpb.Credit) []Credit {
	var credits []Credit
	for _, credit := range in {
		credits = append(credits, ToModelCredit(credit))
	}

	return credits
}

func ToModelGetListUserCreditsResponse(in *creditpb.GetListUserCreditsResponse) GetListUserCreditsV1Response {
	return GetListUserCreditsV1Response{
		Credits:     ToModelsGetListUserCreditsResponse(in.Credits),
		CurrentPage: in.CurrentPage,
		TotalPages:  in.TotalPages,
	}
}

func ToModelPaymentSchedule(in []*creditpb.PaymentSchedule) []PaymentSchedule {
	var payments []PaymentSchedule
	for _, payment := range in {
		payments = append(payments, PaymentSchedule{
			ID:      payment.PaymentId,
			Amount:  payment.Amount,
			DueDate: payment.DueDate,
			Status:  payment.Status.String(),
		})
	}

	return payments
}

func ToModelGetPaymentScheduleResponse(in *creditpb.GetPaymentScheduleResponse) GetPaymentScheduleV1Response {
	return GetPaymentScheduleV1Response{
		Payments: ToModelPaymentSchedule(in.Payments),
	}
}
