package model

import (
	"fmt"
	"net/http"

	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
	"github.com/patyukin/mbs-pkg/pkg/validator"
)

type GetUserReportV1Request struct {
	UserID    string `json:"user_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

func (r *GetUserReportV1Request) Validate() *error_v1.ErrorResponse {
	if r.StartDate == "" || r.EndDate == "" {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "start_date or end_date is empty",
			Description: "start_date or end_date is empty",
		}
	}

	isValid, err := validator.ValidateDate(r.StartDate)
	if err != nil || !isValid {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "start_date: Invalid",
			Description: fmt.Sprintf("failed to parse start_date: %v", err),
		}
	}

	isValid, err = validator.ValidateDate(r.EndDate)
	if err != nil || !isValid {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "end_date: Invalid",
			Description: fmt.Sprintf("failed to parse end_date: %v", err),
		}
	}

	return nil
}

type GetUserReportV1Response struct {
	Message string `json:"message"`
}
