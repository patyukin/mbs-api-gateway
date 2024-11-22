package handler

import (
	"context"
	"encoding/json"
	"github.com/patyukin/mbs-api-gateway/internal/metrics"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
	"github.com/rs/zerolog/log"
	"net/http"
)

const (
	HeaderAuthorization = "Authorization"
	HeaderUserID        = "X-User-ID"
	HeaderRequestUUID   = "X-Request-UUID"
)

//go:generate go run github.com/vektra/mockery/v2@v2.45.1 --name=AuthUseCase
type AuthUseCase interface {
	GetJWTToken() []byte
	SignUpV1(ctx context.Context, in model.SignUpV1Request) (model.SignUpV1Response, *error_v1.ErrorResponse)
	SignInV1(ctx context.Context, in model.SignInV1Request) (model.SignInV1Response, *error_v1.ErrorResponse)
	SignInConfirmation(ctx context.Context, in model.SignInConfirmationV1Request) (model.SignInConfirmationV1Response, *error_v1.ErrorResponse)
	RefreshTokenV1(ctx context.Context, in model.RefreshTokenV1Request) (model.RefreshTokenV1Response, *error_v1.ErrorResponse)
	AuthorizeUser(ctx context.Context, in model.AuthorizeUserV1Request) error
}

//go:generate go run github.com/vektra/mockery/v2@v2.45.1 --name=PaymentUseCase
type PaymentUseCase interface {
	CreateAccountUseCase(ctx context.Context, in model.CreateAccountV1Request) (model.CreateAccountV1Response, *error_v1.ErrorResponse)
	CreatePaymentUseCase(ctx context.Context, in model.CreatePaymentV1Request) (model.CreatePaymentV1Response, *error_v1.ErrorResponse)
	ConfirmationPaymentUseCase(ctx context.Context, in model.ConfirmationPaymentV1Request) (model.VerifyPaymentV1Response, *error_v1.ErrorResponse)
}

//go:generate go run github.com/vektra/mockery/v2@v2.45.1 --name=CreditUseCase
type CreditUseCase interface {
	CreateCreditUseCase(ctx context.Context, in model.CreateCreditV1Request, userID string) (model.CreateCreditV1Response, *error_v1.ErrorResponse)
	CreateCreditApplicationUseCase(ctx context.Context, in model.CreateCreditApplicationV1Request, userID string) (model.CreateCreditApplicationV1Response, *error_v1.ErrorResponse)
	CreditApplicationConfirmationUseCase(ctx context.Context, in model.CreditApplicationConfirmationV1Request, userID string) (model.CreditApplicationConfirmationV1Response, *error_v1.ErrorResponse)
	GetCreditApplicationUseCase(ctx context.Context, applicationID, userID string) (model.GetCreditApplicationV1Response, *error_v1.ErrorResponse)
	UpdateCreditApplicationStatusUseCase(ctx context.Context, in model.UpdateCreditApplicationStatusV1Request) (model.UpdateCreditApplicationStatusV1Response, *error_v1.ErrorResponse)
	GetCreditUseCase(ctx context.Context, creditID, userID string) (model.GetCreditV1Response, *error_v1.ErrorResponse)
	GetListUserCreditsUseCase(ctx context.Context, in model.GetListUserCreditsV1Request) (model.GetListUserCreditsV1Response, *error_v1.ErrorResponse)
	GetPaymentScheduleUseCase(ctx context.Context, userID, creditID string) (model.GetPaymentScheduleV1Response, *error_v1.ErrorResponse)
}

//go:generate go run github.com/vektra/mockery/v2@v2.45.1 --name=ReportUseCase
type ReportUseCase interface {
	GetUserReport(ctx context.Context, in model.GetUserReportV1Request) (model.GetUserReportV1Response, *error_v1.ErrorResponse)
}

//go:generate go run github.com/vektra/mockery/v2@v2.45.1 --name=LoggerUseCase
type LoggerUseCase interface {
	GetLogReport(ctx context.Context, in model.GetLogReportV1Request) (model.GetLogReportV1Response, *error_v1.ErrorResponse)
}

type Handler struct {
	auc AuthUseCase
	luc LoggerUseCase
	puc PaymentUseCase
	cuc CreditUseCase
	ruc ReportUseCase
}

func New(auc AuthUseCase, luc LoggerUseCase, puc PaymentUseCase, cus CreditUseCase, ruc ReportUseCase) *Handler {
	return &Handler{
		auc: auc,
		luc: luc,
		puc: puc,
		cuc: cus,
		ruc: ruc,
	}
}

func (h *Handler) HandleError(w http.ResponseWriter, code int, message string) {
	log.Error().Msgf("Error: %s", message)

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(model.ErrorResponse{Error: message}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) HealthCheck(w http.ResponseWriter, _ *http.Request) {
	metrics.IncomingTraffic.Inc()
	w.WriteHeader(http.StatusOK)
}
