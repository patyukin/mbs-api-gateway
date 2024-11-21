package model

import (
	"errors"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
	"regexp"
	"time"
)

type Email string

func (e Email) Validate() bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(string(e))
}

type Secret string

func (s Secret) String() string {
	return "***"
}

type SignUpV1Request struct {
	Email         Email  `json:"email"`
	Password      Secret `json:"password"`
	TelegramLogin string `json:"telegram_login"`
	LastName      string `json:"last_name"`
	FirstName     string `json:"first_name"`
	Patronymic    string `json:"patronymic"`
	Phone         string `json:"phone"`
	Address       string `json:"address"`
	DateOfBirth   string `json:"date_of_birth"`
}

type SignUpV1Response struct {
	Message string `json:"message"`
}

func (req *SignUpV1Request) Validate() error {
	if req.Email == "" || !req.Email.Validate() {
		return errors.New("не валидный email")
	}

	if len(req.Password) < 6 {
		return errors.New("пароль должен содержать не менее 6 символов")
	}

	if req.FirstName == "" {
		return errors.New("имя не может быть пустым")
	}

	if req.TelegramLogin == "" {
		return errors.New("telegram логин не может быть пустым")
	}

	if req.LastName == "" {
		return errors.New("фамилия не может быть пустой")
	}

	if req.Patronymic == "" {
		return errors.New("отчество не может быть пустым")
	}

	if _, err := time.Parse("2006-01-02", req.DateOfBirth); err != nil {
		return errors.New("день рождения должно быть в формате 2006-01-02")
	}

	if req.Phone == "" {
		return errors.New("мобильный телефон не может быть пустым")
	}

	if req.Address == "" {
		return errors.New("адрес не может быть пустым")
	}

	return nil
}

type UpdateUserProfileV1Request struct {
	Address     string `json:"address"`
	DateOfBirth string `json:"date_of_birth"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Patronymic  string `json:"patronymic"`
}

type SignInV1Response struct {
	Error   *error_v1.ErrorResponse `json:"error,omitempty"`
	Message string                  `json:"message,omitempty"`
}

type SignInV1Request struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (req *SignInV1Request) Validate() error {
	if req.Login == "" {
		return errors.New("логин не может быть пустым")
	}

	if len(req.Password) < 6 {
		return errors.New("пароль должен содержать не менее 6 символов")
	}

	return nil
}

type SignInVerifyV1Request struct {
	Code string `json:"code"`
}

type AuthorizeRequest struct {
	UserID    string `json:"user_id"`
	RoutePath string `json:"route_path"`
	Method    string `json:"method"`
}

type AuthorizeResponse struct {
}

type RefreshTokenV1Request struct {
	RefreshToken string `json:"token"`
}

type RefreshTokenV1Response struct {
	AccessToken string `json:"access_token"`
}

type GetLogReportV1Request struct {
	ServiceName string `json:"service_name"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}

type GetLogReportV1Response struct {
	FileUrl string `json:"file_url"`
}

func (req *GetLogReportV1Request) Validate() error {
	if req.ServiceName == "" {
		return errors.New("название сервиса не может быть пустым")
	}

	if req.StartDate == "" {
		return errors.New("дата начала не может быть пустой")
	}

	if req.EndDate == "" {
		return errors.New("дата окончания не может быть пустой")
	}

	return nil
}

type CreateAccountV1Request struct {
	Currency string `json:"currency"`
	Balance  int64  `json:"balance"`
	UserID   string `json:"user_id"`
}

type CreateAccountV1Response struct {
	Error *string `json:"error"`
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
	Error *string `json:"error"`
}

type VerifyPaymentV1Request struct {
	Code   string `json:"code"`
	UserID string `json:"user_id"`
}

type VerifyPaymentV1Response struct {
	Error *string `json:"error"`
}
