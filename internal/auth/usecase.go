package auth

import (
	"context"
	dtoRequest "github.com/patyukin/mbs-api-gateway/internal/dto/request"
)

type UseCase struct {
}

func New() *UseCase {
	return &UseCase{}
}

func (uc *UseCase) SignUpV1(ctx context.Context, in dtoRequest.SignUpV1Request) error {
	return nil
}
