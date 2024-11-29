package model

import (
	"fmt"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
	"github.com/patyukin/mbs-pkg/pkg/validator"
	"net/http"
	"regexp"
)

type Secret string

func (s Secret) String() string {
	return "***"
}

type SignUpV1Request struct {
	Email         string `json:"email"`
	Password      Secret `json:"password"`
	TelegramLogin string `json:"telegram_login"`
	LastName      string `json:"last_name"`
	FirstName     string `json:"first_name"`
	Patronymic    string `json:"patronymic"`
	Phone         string `json:"phone"`
	Address       string `json:"address"`
	DateOfBirth   string `json:"date_of_birth"`
}

func (req *SignUpV1Request) Validate() *error_v1.ErrorResponse {
	re, err := regexp.Compile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if err != nil {
		return &error_v1.ErrorResponse{
			Code:        http.StatusInternalServerError,
			Message:     "Internal Server Error",
			Description: fmt.Sprintf("failed to compile regex: %v", err),
		}
	}

	if re.MatchString(req.Email) == false {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "email: Invalid",
			Description: "email is invalid",
		}
	}

	if len(req.Password) < 6 {
		msg := "password: Invalid, password must be at least 6 characters long"
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     msg,
			Description: msg,
		}
	}

	if req.FirstName == "" {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "first_name: required",
			Description: "first_name is required",
		}
	}

	if req.TelegramLogin == "" {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "telegram_login: required",
			Description: "telegram_login is required",
		}
	}

	if req.LastName == "" {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "last_name: required",
			Description: "last_name is required",
		}
	}

	if req.Patronymic == "" {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "patronymic: required",
			Description: "patronymic is required",
		}
	}

	isValid, err := validator.ValidateDate(req.DateOfBirth)
	if err != nil || !isValid {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "date_of_birth: Invalid",
			Description: fmt.Sprintf("failed to parse date_of_birth: %v", err),
		}
	}

	if req.Phone == "" {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "phone: required",
			Description: "phone is required",
		}
	}

	if req.Address == "" {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "address: required",
			Description: "address is required",
		}
	}

	return nil
}

type SignUpV1Response struct {
	Message string `json:"message"`
}

type SignInV1Request struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (req *SignInV1Request) Validate() *error_v1.ErrorResponse {
	if req.Login == "" {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "login: required",
			Description: "login is required",
		}
	}

	if len(req.Password) < 6 {
		msg := "password: Invalid, password must be at least 6 characters long"
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     msg,
			Description: msg,
		}
	}

	return nil
}

type SignInV1Response struct {
	Message string `json:"message,omitempty"`
}

type GetUserByIDV1Request struct {
	UserID string `json:"user_id"`
}

type UserV1 struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Patronymic    string `json:"patronymic"`
	DateOfBirth   string `json:"date_of_birth"`
	Phone         string `json:"phone"`
	Address       string `json:"address"`
	TelegramLogin string `json:"telegram_login"`
	TelegramID    string `json:"telegram_id"`
	ChatID        string `json:"chat_id"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

type GetUserByIDV1Response struct {
	UserV1 UserV1 `json:"user"`
}

type GetUsersV1Request struct {
	Page  int32 `json:"page"`
	Limit int32 `json:"limit"`
}

type ProfileV1 struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Patronymic  string `json:"patronymic"`
	DateOfBirth string `json:"date_of_birth"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
}

type UserInfoV1 struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	ProfileV1 ProfileV1 `json:"profile"`
}

type GetUsersV1Response struct {
	UsersInfoV1 []UserInfoV1 `json:"users"`
	Total       int32        `json:"total"`
}

type SignInConfirmationV1Request struct {
	Code string `json:"code"`
}

type SignInConfirmationV1Response struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AddUserRoleV1Request struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
}

type AddUserRoleV1Response struct {
	Message string `json:"message"`
}

type AuthorizeUserV1Request struct {
	UserID    string `json:"user_id"`
	RoutePath string `json:"route_path"`
	Method    string `json:"method"`
}

type AuthorizeUserV1Response struct {
	Message string `json:"message"`
}

type RefreshTokenV1Request struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenV1Response struct {
	AccessToken string `json:"access_token"`
}
