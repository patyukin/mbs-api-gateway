package credit

import (
	"context"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
)

func (u *UseCase) CreateCreditUseCase(ctx context.Context, in model.CreateCreditV1Request, userID string) (model.CreateCreditV1Response, *error_v1.ErrorResponse) {
	panic("")
}
