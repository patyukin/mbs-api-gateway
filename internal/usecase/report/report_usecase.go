package report

import (
	"context"

	"github.com/patyukin/mbs-pkg/pkg/proto/report_v1"
	"google.golang.org/grpc"
)

type ProtoReportClient interface {
	GetUserReport(ctx context.Context, in *report_v1.GetUserReportRequest, opts ...grpc.CallOption) (*report_v1.GetUserReportResponse, error)
}

type UseCase struct {
	reportClient ProtoReportClient
}

func New(reportClient ProtoReportClient) *UseCase {
	return &UseCase{
		reportClient: reportClient,
	}
}
