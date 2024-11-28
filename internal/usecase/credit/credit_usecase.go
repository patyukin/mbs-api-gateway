package credit

import (
	"context"
	creditpb "github.com/patyukin/mbs-pkg/pkg/proto/credit_v1"
	"google.golang.org/grpc"
)

type ProtoClient interface {
	CreateCreditApplication(ctx context.Context, in *creditpb.CreateCreditApplicationRequest, opts ...grpc.CallOption) (*creditpb.CreateCreditApplicationResponse, error)
	CreditApplicationConfirmation(ctx context.Context, in *creditpb.CreditApplicationConfirmationRequest, opts ...grpc.CallOption) (*creditpb.CreditApplicationConfirmationResponse, error)
	CreateCredit(ctx context.Context, in *creditpb.CreateCreditRequest, opts ...grpc.CallOption) (*creditpb.CreateCreditResponse, error)
	GetCreditApplication(ctx context.Context, in *creditpb.GetCreditApplicationRequest, opts ...grpc.CallOption) (*creditpb.GetCreditApplicationResponse, error)
	UpdateCreditApplicationSolution(ctx context.Context, in *creditpb.UpdateCreditApplicationSolutionRequest, opts ...grpc.CallOption) (*creditpb.UpdateCreditApplicationSolutionResponse, error)
	GetCredit(ctx context.Context, in *creditpb.GetCreditRequest, opts ...grpc.CallOption) (*creditpb.GetCreditResponse, error)
	GetListUserCredits(ctx context.Context, in *creditpb.GetListUserCreditsRequest, opts ...grpc.CallOption) (*creditpb.GetListUserCreditsResponse, error)
	GetPaymentSchedule(ctx context.Context, in *creditpb.GetPaymentScheduleRequest, opts ...grpc.CallOption) (*creditpb.GetPaymentScheduleResponse, error)
}

type UseCase struct {
	creditClient ProtoClient
}

func New(creditClient ProtoClient) *UseCase {
	return &UseCase{
		creditClient: creditClient,
	}
}
