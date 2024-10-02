package model

import (
	"errors"
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
	LastName      string `json:"last_name"`
	FirstName     string `json:"first_name"`
	Patronymic    string `json:"patronymic"`
	Phone         string `json:"phone"`
	Address       string `json:"address"`
	DateOfBirth   string `json:"date_of_birth"`
	TelegramLogin string `json:"telegram_login"`
}

func (req *SignUpV1Request) Validate() error {
	if req.Email == "" || !req.Email.Validate() {
		return errors.New("не валидный email")
	}

	if len(req.Password) < 6 {
		return errors.New("пароль должен содержать не менее 6 символов")
	}

	if req.TelegramLogin == "" {
		return errors.New("логин телеграм не может быть пустым")
	}

	if req.FirstName == "" {
		return errors.New("имя не может быть пустым")
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
		return errors.New("аддресс не может быть пустым")
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

type SignInV1Request struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
