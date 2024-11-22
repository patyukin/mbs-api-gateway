package model

import (
	"fmt"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
	"github.com/patyukin/mbs-pkg/pkg/validator"
	"net/http"
)

type GetLogReportV1Request struct {
	ServiceName string `json:"service_name"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}

type GetLogReportV1Response struct {
	FileUrl string `json:"file_url"`
}

func (req *GetLogReportV1Request) Validate() *error_v1.ErrorResponse {
	if req.ServiceName == "" {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "service_name: required",
			Description: "service_name is required",
		}
	}

	if req.StartDate == "" {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "start_date: required",
			Description: "start_date is required",
		}
	}

	isValid, err := validator.ValidateDate(req.StartDate)
	if err != nil || !isValid {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "start_date: Invalid",
			Description: fmt.Sprintf("failed to parse start_date: %v", err),
		}
	}

	if req.EndDate == "" {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "end_date: required",
			Description: "end_date is required",
		}
	}

	isValid, err = validator.ValidateDate(req.EndDate)
	if err != nil || !isValid {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "end_date: Invalid",
			Description: fmt.Sprintf("failed to parse end_date: %v", err),
		}
	}

	return nil
}
