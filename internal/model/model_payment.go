package model

type CreateAccountV1Request struct {
	Currency string `json:"currency"`
	Balance  int64  `json:"balance"`
	UserID   string `json:"user_id"`
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
