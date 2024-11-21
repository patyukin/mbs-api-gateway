package payment

import (
	"context"
	paymentpb "github.com/patyukin/mbs-pkg/pkg/proto/payment_v1"
	"google.golang.org/grpc"
)

type ProtoPaymentClient interface {
	CreateAccount(ctx context.Context, request *paymentpb.CreateAccountRequest, opts ...grpc.CallOption) (*paymentpb.CreateAccountResponse, error)
	CreatePayment(ctx context.Context, request *paymentpb.CreatePaymentRequest, opts ...grpc.CallOption) (*paymentpb.CreatePaymentResponse, error)
	ConfirmationPayment(ctx context.Context, request *paymentpb.ConfirmationPaymentRequest, opts ...grpc.CallOption) (*paymentpb.ConfirmationPaymentResponse, error)
}

type UseCase struct {
	paymentClient ProtoPaymentClient
}

func New(paymentClient ProtoPaymentClient) *UseCase {
	return &UseCase{
		paymentClient: paymentClient,
	}
}
