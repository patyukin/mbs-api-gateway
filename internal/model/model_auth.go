package model

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
	"github.com/patyukin/mbs-pkg/pkg/validator"
	"net/http"
	"regexp"
)

const minPasswordLength = 8

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
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if re.MatchString(req.Email) == false {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "email: Invalid",
			Description: "email is invalid",
		}
	}

	if len(req.Password) < minPasswordLength {
		msg := "password: Invalid, password must be at least 8 characters long"
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

	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if re.MatchString(req.Login) == false {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "login: Invalid",
			Description: "login is invalid",
		}
	}

	if len(req.Password) < minPasswordLength {
		msg := "password: Invalid, password must be at least 8 characters long"
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

func (r *GetUserByIDV1Request) Validate(userIDFromHeader, roleFromHeader string) *error_v1.ErrorResponse {
	if r.UserID == "" {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "user_id: required",
			Description: "user_id is required",
		}
	}

	_, err := uuid.Parse(r.UserID)
	if err != nil {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "user_id: Invalid",
			Description: "user_id is invalid",
		}
	}

	if roleFromHeader == "sys-admin" {
		return nil
	}

	if r.UserID != userIDFromHeader {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "user_id: Invalid",
			Description: "user_id is invalid",
		}
	}

	return nil
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
	UserInfoV1 UserInfoV1 `json:"user"`
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

type GetUsersV1Request struct {
	Page  int32 `json:"page"`
	Limit int32 `json:"limit"`
}

func (req *GetUsersV1Request) Validate() *error_v1.ErrorResponse {
	if req.Page < 0 {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "page: Invalid",
			Description: "page is invalid",
		}
	}

	if req.Limit < 0 {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "limit: Invalid",
			Description: "limit is invalid",
		}
	}

	return nil
}

type GetUsersV1Response struct {
	UsersInfoV1 []UserInfoV1 `json:"users"`
	Total       int32        `json:"total"`
}

type SignInConfirmationV1Request struct {
	Login string `json:"login"`
	Code  string `json:"code"`
}

func (req *SignInConfirmationV1Request) Validate() *error_v1.ErrorResponse {
	if req.Login == "" {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "login: required",
			Description: "login is required",
		}
	}

	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if re.MatchString(req.Login) == false {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "login: Invalid",
			Description: "login is invalid",
		}
	}

	if req.Code == "" {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "code: required",
			Description: "code is required"}
	}

	return nil
}

type SignInConfirmationV1Response struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AddUserRoleV1Request struct {
	UserID string `json:"user_id"`
	RoleID string `json:"role_id"`
}

func (r *AddUserRoleV1Request) Validate() *error_v1.ErrorResponse {
	if r.UserID == "" {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "user_id: required",
			Description: "user_id is required",
		}
	}

	_, err := uuid.Parse(r.UserID)
	if err != nil {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "user_id: Invalid",
			Description: "user_id is invalid",
		}
	}

	if r.RoleID == "" {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "role: required",
			Description: "role is required",
		}
	}

	_, err = uuid.Parse(r.RoleID)
	if err != nil {
		return &error_v1.ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "role: Invalid",
			Description: "role is invalid",
		}
	}

	return nil
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
