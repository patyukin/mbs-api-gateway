package model

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
)

type CreateAccountV1Request struct {
	Currency string `json:"currency"`
	Balance  int64  `json:"balance"`
}

func (r *CreateAccountV1Request) Validate(userID string) *error_v1.ErrorResponse {
	if r.Currency != "RUB" {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "currency: Invalid",
			Description: "currency is invalid",
		}
	}

	if r.Balance < 0 {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "balance: Invalid",
			Description: "balance is invalid",
		}
	}

	if userID == "" {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "user_id: required",
			Description: "user_id is required",
		}
	}

	_, err := uuid.Parse(userID)
	if err != nil {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "user_id: Invalid",
			Description: "user_id is invalid",
		}
	}

	return nil
}

type CreateAccountV1Response struct {
	Message string `json:"message"`
}

type CreatePaymentV1Request struct {
	SenderAccountID   string `json:"sender_account_id"`
	ReceiverAccountID string `json:"receiver_account_id"`
	Amount            int64  `json:"amount"`
	Currency          string `json:"currency"`
	Description       string `json:"description"`
	UserID            string `json:"user_id"`
}

func (r *CreatePaymentV1Request) Validate(userID string) *error_v1.ErrorResponse {
	if r.ReceiverAccountID == "" {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "receiver_account_id: required",
			Description: "receiver_account_id is required",
		}
	}

	_, err := uuid.Parse(r.ReceiverAccountID)
	if err != nil {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "receiver_account_id: Invalid",
			Description: "receiver_account_id is invalid",
		}
	}

	if r.SenderAccountID == "" {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "sender_account_id: required",
			Description: "sender_account_id is required",
		}
	}

	_, err = uuid.Parse(r.SenderAccountID)
	if err != nil {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "sender_account_id: Invalid",
			Description: "sender_account_id is invalid",
		}
	}

	if r.Amount < 0 {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "amount: Invalid",
			Description: "amount is invalid",
		}
	}

	if r.Currency != "RUB" {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "currency: Invalid",
			Description: "currency is invalid",
		}
	}

	if userID == "" {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "user_id: required",
			Description: "user_id is required",
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

	return nil
}

type CreatePaymentV1Response struct {
	Message string `json:"message"`
}

type ConfirmationPaymentV1Request struct {
	Code string `json:"code"`
}

func (r *ConfirmationPaymentV1Request) Validate(userID string) *error_v1.ErrorResponse {
	if r.Code == "" {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "code: required",
			Description: "code is required",
		}
	}

	_, err := uuid.Parse(r.Code)
	if err != nil {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "code: Invalid",
			Description: "code is invalid",
		}
	}

	if userID == "" {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "user_id: required",
			Description: "user_id is required",
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

	return nil
}

type VerifyPaymentV1Response struct {
	Message string `json:"message"`
}

type GetPaymentV1Request struct {
	PaymentID string `json:"payment_id"`
}

func (r *GetPaymentV1Request) Validate(userID string) *error_v1.ErrorResponse {
	if r.PaymentID == "" {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "payment_id: required",
			Description: "payment_id is required",
		}
	}

	_, err := uuid.Parse(r.PaymentID)
	if err != nil {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "payment_id: Invalid",
			Description: "payment_id is invalid",
		}
	}

	if userID == "" {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "user_id: required",
			Description: "user_id is required",
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

	return nil
}

type Payment struct {
	ID                string `json:"id"`
	SenderAccountID   string `json:"sender_account_id"`
	ReceiverAccountID string `json:"receiver_account_id"`
	Amount            int64  `json:"amount"`
	Currency          string `json:"currency"`
	Description       string `json:"description"`
	Status            string `json:"status"`
	CreatedAt         string `json:"created_at"`
}

type GetPaymentV1Response struct {
	Payment Payment `json:"payment"`
}
