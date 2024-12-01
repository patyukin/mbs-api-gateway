package model

import (
	"github.com/google/uuid"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
	"net/http"
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

type CreatePaymentV1Response struct {
	Message string `json:"message"`
}

type ConfirmationPaymentV1Request struct {
	Code   string `json:"code"`
	UserID string `json:"user_id"`
}

type VerifyPaymentV1Response struct {
	Message string `json:"message"`
}
