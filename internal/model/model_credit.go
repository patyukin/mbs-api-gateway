package model

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/patyukin/mbs-pkg/pkg/mapping/creditmapper"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
	"net/http"
	"time"
)

type CreateCreditApplicationV1Request struct {
	RequestedAmount int64  `json:"requested_amount" validate:"required,gt=0"`
	InterestRate    int32  `json:"interest_rate" validate:"required,gt=0"`
	StartDate       string `json:"start_date" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
	EndDate         string `json:"end_date" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
	Description     string `json:"description" validate:"omitempty,max=500"`
}

func (req *CreateCreditApplicationV1Request) Validate() *error_v1.ErrorResponse {
	validate := validator.New()

	endDate, err := time.Parse(time.RFC3339, req.EndDate)
	if err != nil {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "end_date: Invalid",
			Description: fmt.Sprintf("failed to parse end_date: %v", err),
		}
	}

	startDate, err := time.Parse(time.RFC3339, req.StartDate)
	if err != nil {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "start_date: Invalid",
			Description: fmt.Sprintf("failed to parse start_date: %w", err),
		}
	}

	if endDate.After(startDate) {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "end_date: end_date must be after start_date",
			Description: "end_date must be after start_date",
		}
	}

	err = validate.Struct(req)
	if err != nil {
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			return &error_v1.ErrorResponse{
				Code:        http.StatusBadRequest,
				Message:     "Invalid Data",
				Description: fmt.Sprintf("failed to validate request: %v", err),
			}
		}

		var validationErrors validator.ValidationErrors
		errors.As(err, &validationErrors)
		errorMessages := ""
		for _, fieldErr := range validationErrors {
			switch fieldErr.Tag() {
			case "required":
				errorMessages += fmt.Sprintf("Field '%s' is required.\n", fieldErr.Field())
			case "gt":
				errorMessages += fmt.Sprintf("Field '%s' must be greater than zero.\n", fieldErr.Field())
			case "datetime":
				errorMessages += fmt.Sprintf("Field '%s' must be in RFC3339 format.\n", fieldErr.Field())
			case "max":
				errorMessages += fmt.Sprintf("Field '%s' must be less than or equal to 500 characters.\n", fieldErr.Field())
			default:
				errorMessages += fmt.Sprintf("Field '%s' is invalid.\n", fieldErr.Field())
			}
		}

		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "Invalid Data",
			Description: fmt.Sprintf("failed to validate request: %s", errorMessages),
		}
	}

	return nil
}

type CreateCreditApplicationV1Response struct {
	Message string `json:"message"`
}

type CreateCreditV1Request struct {
	ApplicationID string `json:"application_id"`
	UserID        string `json:"user_id"`
}

type CreateCreditV1Response struct {
	Message string `json:"message"`
}

type CreditApplicationConfirmationV1Request struct {
	Code string `json:"code"`
}

func (req *CreditApplicationConfirmationV1Request) Validate() *error_v1.ErrorResponse {
	if req.Code == "" {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "code: required",
			Description: "code is required",
		}
	}

	return nil
}

type CreditApplicationConfirmationV1Response struct {
	ApplicationID string `json:"application_id"`
	Message       string `json:"message"`
	Status        string `json:"status"`
}

type GetCreditApplicationV1Request struct {
	ApplicationID string `json:"application_id"`
}

type GetCreditApplicationV1Response struct {
	ApplicationID  string `json:"application_id"`
	Status         string `json:"status"`
	ApprovedAmount int64  `json:"approved_amount"`
	DecisionDate   string `json:"decision_date"`
	Description    string `json:"description"`
}

type UpdateCreditApplicationStatusV1Request struct {
	ApplicationID string `json:"application_id"`
	Status        string `json:"status"`
	DecisionNotes string `json:"decision_notes"`
}

func (r *UpdateCreditApplicationStatusV1Request) Validate() *error_v1.ErrorResponse {
	_, err := uuid.Parse(r.ApplicationID)
	if err != nil {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "application_id: Invalid",
			Description: fmt.Sprintf("failed to parse application_id: %v", err),
		}
	}

	err = creditmapper.ValidateStringCreditApplicationStatus(r.Status)
	if err != nil {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "status: Invalid",
			Description: fmt.Sprintf("failed to validate status: %v", err),
		}
	}

	return nil
}

type UpdateCreditApplicationStatusV1Response struct {
	Message string `json:"message"`
}

type GetListUserCreditsV1Request struct {
	UserID string `json:"user_id"`
	Page   int32  `json:"page"`
	Limit  int32  `json:"limit"`
}

type CreditV1 struct {
	CreditID        string `json:"credit_id"`
	UserID          string `json:"user_id"`
	Amount          int64  `json:"amount"`
	InterestRate    int32  `json:"interest_rate"`
	RemainingAmount int64  `json:"remaining_amount"`
	Status          string `json:"status"`
	StartDate       string `json:"start_date"`
	EndDate         string `json:"end_date"`
	Description     string `json:"description"`
}

type GetCreditV1Response struct {
	CreditV1 CreditV1 `json:"credit"`
}

type GetListUserCreditsV1Response struct {
	Credits []CreditV1 `json:"credits"`
	Total   int32      `json:"total"`
}

type PaymentSchedule struct {
	ID      string `json:"id"`
	Amount  int64  `json:"amount"`
	DueDate string `json:"due_date"`
	Status  string `json:"status"`
}

type GetPaymentScheduleV1Response struct {
	Payments []PaymentSchedule `json:"payments"`
}
