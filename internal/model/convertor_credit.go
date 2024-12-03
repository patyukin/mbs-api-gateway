package model

import (
	"fmt"
	"github.com/patyukin/mbs-pkg/pkg/mapping/creditmapper"
	creditpb "github.com/patyukin/mbs-pkg/pkg/proto/credit_v1"
)

func ToProtoV1CreateCreditApplicationRequest(in CreateCreditApplicationV1Request, userID string) creditpb.CreateCreditApplicationRequest {
	return creditpb.CreateCreditApplicationRequest{
		UserId:          userID,
		RequestedAmount: in.RequestedAmount,
		InterestRate:    in.InterestRate,
		Description:     in.Description,
	}
}

func ToProtoV1CreditApplicationConfirmationRequest(in CreditApplicationConfirmationV1Request, userID string) creditpb.CreditApplicationConfirmationRequest {
	return creditpb.CreditApplicationConfirmationRequest{
		Code:   in.Code,
		UserId: userID,
	}
}

func ToModelGetCreditApplicationV1Response(in *creditpb.GetCreditApplicationResponse) GetCreditApplicationV1Response {
	return GetCreditApplicationV1Response{
		ApplicationID:  in.GetApplicationId(),
		ApprovedAmount: in.GetApprovedAmount(),
		Status:         in.GetStatus().String(),
		DecisionDate:   in.GetDecisionDate(),
		Description:    in.GetDescription(),
	}
}

func ToProtoV1UpdateCreditApplicationStatusRequest(in UpdateCreditApplicationStatusV1Request, applicationID string) (creditpb.UpdateCreditApplicationSolutionRequest, error) {
	status, err := creditmapper.StringToEnumCreditApplicationStatus(in.Status)
	if err != nil {
		return creditpb.UpdateCreditApplicationSolutionRequest{}, fmt.Errorf("failed creditmapper.StringToEnumCreditApplicationStatus: %w", err)
	}

	return creditpb.UpdateCreditApplicationSolutionRequest{
		ApplicationId:  applicationID,
		Status:         status,
		DecisionNotes:  in.DecisionNotes,
		ApprovedAmount: in.ApprovedAmount,
	}, nil
}

func ToModelCredit(in *creditpb.Credit) CreditV1 {
	return CreditV1{
		CreditID:        in.GetCreditId(),
		UserID:          in.GetUserId(),
		Amount:          in.GetAmount(),
		InterestRate:    in.GetInterestRate(),
		RemainingAmount: in.GetRemainingAmount(),
		Status:          in.GetStatus().String(),
		StartDate:       in.GetStartDate(),
		EndDate:         in.GetEndDate(),
		Description:     in.GetDescription(),
	}
}

func ToModelGetCreditResponse(in *creditpb.Credit) GetCreditV1Response {
	return GetCreditV1Response{
		CreditV1: ToModelCredit(in),
	}
}

func ToModelsGetListUserCreditsResponse(in []*creditpb.Credit) []CreditV1 {
	credits := make([]CreditV1, 0, len(in))
	for _, credit := range in {
		credits = append(credits, ToModelCredit(credit))
	}

	return credits
}

func ToModelGetListUserCreditsResponse(in *creditpb.GetListUserCreditsResponse) GetListUserCreditsV1Response {
	return GetListUserCreditsV1Response{
		Credits: ToModelsGetListUserCreditsResponse(in.GetCredits()),
		Total:   in.GetTotal(),
	}
}

func ToModelPaymentSchedule(in []*creditpb.PaymentSchedule) []PaymentSchedule {
	payments := make([]PaymentSchedule, 0, len(in))
	for _, payment := range in {
		payments = append(
			payments, PaymentSchedule{
				ID:      payment.GetPaymentId(),
				Amount:  payment.GetAmount(),
				DueDate: payment.GetDueDate(),
				Status:  payment.GetStatus().String(),
			},
		)
	}

	return payments
}

func ToModelGetPaymentScheduleResponse(in *creditpb.GetPaymentScheduleResponse) GetPaymentScheduleV1Response {
	return GetPaymentScheduleV1Response{
		Payments: ToModelPaymentSchedule(in.GetPayments()),
	}
}

func ToProtoV1CreateCreditRequest(in CreateCreditV1Request, userID string) creditpb.CreateCreditRequest {
	return creditpb.CreateCreditRequest{
		UserId:           userID,
		ApplicationId:    in.ApplicationID,
		CreditTermMonths: in.CreditTermMonths,
		AccountId:        in.AccountID,
	}
}
