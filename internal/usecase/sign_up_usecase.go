package usecase

import (
	"context"
	dtoRequest "github.com/patyukin/mbs-api-gateway/internal/dto/request"
)

func (uc *UseCase) SignUpV1(ctx context.Context, in dtoRequest.SignUpV1Request) error {
	return nil
}
