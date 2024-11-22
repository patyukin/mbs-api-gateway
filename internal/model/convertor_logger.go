package model

import (
	loggerpb "github.com/patyukin/mbs-pkg/pkg/proto/logger_v1"
)

func ToProtoLogReportFromRequest(in GetLogReportV1Request) (loggerpb.LogReportRequest, error) {
	return loggerpb.LogReportRequest{
		StartTime:   in.StartDate,
		EndTime:     in.EndDate,
		ServiceName: in.ServiceName,
	}, nil
}
