package model

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/patyukin/mbs-pkg/pkg/mapping/creditmapper"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
)

type CreateCreditApplicationV1Request struct {
	RequestedAmount int64  `json:"requested_amount"`
	InterestRate    int32  `json:"interest_rate"`
	Description     string `json:"description"`
}

func (req *CreateCreditApplicationV1Request) Validate(userID string) *error_v1.ErrorResponse {
	_, err := uuid.Parse(userID)
	if err != nil {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "user_id: Invalid",
			Description: fmt.Sprintf("failed to parse user_id: %v", err),
		}
	}

	if req.RequestedAmount <= 0 {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "requested_amount: Invalid",
			Description: "requested_amount is invalid",
		}
	}

	if req.InterestRate <= 0 || req.InterestRate >= 100 {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "interest_rate: Invalid",
			Description: "interest_rate is invalid",
		}
	}

	if req.Description == "" {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "description: required",
			Description: "description is required",
		}
	}

	return nil
}

type CreateCreditApplicationV1Response struct {
	Message string `json:"message"`
}

type CreateCreditV1Request struct {
	ApplicationID    string `json:"application_id"`
	AccountID        string `json:"account_id"`
	CreditTermMonths int32  `json:"credit_term_months"`
}

func (in *CreateCreditV1Request) Validate(userID string) *error_v1.ErrorResponse {
	if in.ApplicationID == "" {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "application_id: required",
			Description: "application_id is required",
		}
	}

	_, err := uuid.Parse(in.ApplicationID)
	if err != nil {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "application_id: Invalid",
			Description: "application_id is invalid",
		}
	}

	_, err = uuid.Parse(userID)
	if err != nil {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "user_id: Invalid",
			Description: "user_id is invalid",
		}
	}

	if in.AccountID == "" {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "account_id: required",
			Description: "account_id is required",
		}
	}

	_, err = uuid.Parse(in.AccountID)
	if err != nil {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "account_id: Invalid",
			Description: "account_id is invalid",
		}
	}

	if in.CreditTermMonths <= 0 {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "credit_term_months: required",
			Description: "credit_term_months is required",
		}
	}

	return nil
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
	Message string `json:"message"`
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
	Status         string `json:"status"`
	DecisionNotes  string `json:"decision_notes"`
	ApprovedAmount int64  `json:"approved_amount"`
}

func (r *UpdateCreditApplicationStatusV1Request) Validate(userID, applicationID string) *error_v1.ErrorResponse {
	_, err := uuid.Parse(applicationID)
	if err != nil {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "application_id: Invalid",
			Description: fmt.Sprintf("failed to parse application_id: %v", err),
		}
	}

	_, err = uuid.Parse(userID)
	if err != nil {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "user_id: Invalid",
			Description: fmt.Sprintf("failed to parse user_id: %v", err),
		}
	}

	if r.DecisionNotes == "" {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "decision_notes: required",
			Description: "decision_notes is required",
		}
	}

	if r.Status == "" {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "status: required",
			Description: "status is required",
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

	if r.ApprovedAmount < 0 {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "approved_amount: Invalid",
			Description: "approved_amount must be greater than zero",
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
