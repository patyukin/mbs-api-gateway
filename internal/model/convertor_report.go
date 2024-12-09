package model

import (
	reportpb "github.com/patyukin/mbs-pkg/pkg/proto/report_v1"
)

func ToProtoGetUserReport(in GetUserReportV1Request) reportpb.GetUserReportRequest {
	return reportpb.GetUserReportRequest{
		UserId:    in.UserID,
		StartDate: in.StartDate,
		EndDate:   in.EndDate,
	}
}
