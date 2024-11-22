package model

import (
	paymentpb "github.com/patyukin/mbs-pkg/pkg/proto/payment_v1"
)

func ToProtoCreateAccountFromRequest(in CreateAccountV1Request) paymentpb.CreateAccountRequest {
	return paymentpb.CreateAccountRequest{
		UserId:   in.UserID,
		Currency: in.Currency,
		Balance:  in.Balance,
	}
}

func ToProtoCreatePaymentFromRequest(in CreatePaymentV1Request) paymentpb.CreatePaymentRequest {
	return paymentpb.CreatePaymentRequest{
		SenderAccountId:   in.SenderAccountID,
		ReceiverAccountId: in.ReceiverAccountID,
		Amount:            in.Amount,
		Currency:          in.Currency,
		Description:       in.Description,
		UserId:            in.UserID,
	}
}

func ToProtoVerifyPaymentFromRequest(in ConfirmationPaymentV1Request) paymentpb.ConfirmationPaymentRequest {
	return paymentpb.ConfirmationPaymentRequest{
		UserId: in.UserID,
		Code:   in.Code,
	}
}
