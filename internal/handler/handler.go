package handler

import (
	"net/http"
)

type UseCase interface {
}

type Handler struct {
	uc UseCase
}

func (h *Handler) RefreshTokenV1(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) SignInV1(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

// SignUpV1 docs
// @Summary Register a new user
// @Description Register a new user in the system
// @Tags Auth
// @Accept json
// @Produce json
// @Param SignUpRequest body request.SignUpV1Request true "Запрос пользователя на регистрацию"
// @Success 201 "Пользователь успешно зарегистрирован"
// @Failure 400 {object} response.ErrorResponse "Invalid request"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /v1/sign-up [post]
func (h *Handler) SignUpV1(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) GetUserProfileV1(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) UpdateUserProfileV1(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func New(uc UseCase) *Handler {
	return &Handler{uc: uc}
}
