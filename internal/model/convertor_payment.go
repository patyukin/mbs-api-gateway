package model

import (
	"fmt"

	"github.com/patyukin/mbs-pkg/pkg/mapping/paymentmapper"
	paymentpb "github.com/patyukin/mbs-pkg/pkg/proto/payment_v1"
)

func ToProtoCreateAccountFromRequest(in CreateAccountV1Request, userID string) paymentpb.CreateAccountRequest {
	return paymentpb.CreateAccountRequest{
		UserId:   userID,
		Currency: in.Currency,
		Balance:  in.Balance,
	}
}

func ToProtoCreatePaymentFromRequest(in *CreatePaymentV1Request, userID string) *paymentpb.CreatePaymentRequest {
	return &paymentpb.CreatePaymentRequest{
		SenderAccountId:   in.SenderAccountID,
		ReceiverAccountId: in.ReceiverAccountID,
		Amount:            in.Amount,
		Currency:          in.Currency,
		Description:       in.Description,
		UserId:            userID,
	}
}

func ToProtoVerifyPaymentFromRequest(in ConfirmationPaymentV1Request, userID string) paymentpb.ConfirmationPaymentRequest {
	return paymentpb.ConfirmationPaymentRequest{
		UserId: userID,
		Code:   in.Code,
	}
}

func ToProtoGetPaymentFromRequest(in GetPaymentV1Request, userID string) paymentpb.GetPaymentRequest {
	return paymentpb.GetPaymentRequest{
		PaymentId: in.PaymentID,
		UserId:    userID,
	}
}

func ToModelPayment(in *paymentpb.Payment) (Payment, error) {
	status, err := paymentmapper.EnumToStringPaymentStatus(in.GetStatus())
	if err != nil {
		return Payment{}, fmt.Errorf("failed to map payment status: %w", err)
	}

	return Payment{
		ID:                in.GetId(),
		SenderAccountID:   in.GetSenderAccountId(),
		ReceiverAccountID: in.GetReceiverAccountId(),
		Amount:            in.GetAmount(),
		Currency:          in.GetCurrency(),
		Description:       in.GetDescription(),
		Status:            status,
		CreatedAt:         in.GetCreatedAt(),
	}, nil
}

func ToModelGetPaymentV1Response(in *paymentpb.GetPaymentResponse) (GetPaymentV1Response, error) {
	payment, err := ToModelPayment(in.GetPayment())
	if err != nil {
		return GetPaymentV1Response{}, fmt.Errorf("failed to map payment: %w", err)
	}

	return GetPaymentV1Response{Payment: payment}, nil
}
