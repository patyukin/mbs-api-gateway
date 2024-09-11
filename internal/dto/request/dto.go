package request

type SignUpV1Request struct {
	Email         string `json:"email"`
	Password      string `json:"password"`
	LastName      string `json:"last_name"`
	FirstName     string `json:"first_name"`
	Patronymic    string `json:"patronymic"`
	Phone         string `json:"phone"`
	Address       string `json:"address"`
	DateOfBirth   string `json:"date_of_birth"`
	TelegramLogin string `json:"telegram_login"`
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
