package model

type ErrorResponse struct {
	Error string `json:"error"`
}

type TokensResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
