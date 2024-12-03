package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/patyukin/mbs-api-gateway/internal/metrics"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
	"github.com/rs/zerolog/log"
)

const (
	HeaderAuthorization string = "Authorization"
	HeaderUserID        string = "X-User-Id"
	HeaderUserRole      string = "X-User-Role"
	RequestID           string = "X-Request-Id"
	TraceID             string = "X-Trace-Id"
	minLimit                   = 10
	SysAdmin                   = "sys-admin"
)

//go:generate go run github.com/vektra/mockery/v2@v2.45.1 --name=AuthUseCase
type AuthUseCase interface {
	GetJWTToken() []byte
	SignUpV1UseCase(ctx context.Context, in model.SignUpV1Request) (model.SignUpV1Response, *error_v1.ErrorResponse)
	SignInV1UseCase(ctx context.Context, in model.SignInV1Request) (model.SignInV1Response, *error_v1.ErrorResponse)
	SignInConfirmationV1UseCase(ctx context.Context, in model.SignInConfirmationV1Request) (model.SignInConfirmationV1Response, *error_v1.ErrorResponse)
	GetUserByIDV1UseCase(ctx context.Context, in model.GetUserByIDV1Request) (model.GetUserByIDV1Response, *error_v1.ErrorResponse)
	RefreshTokenV1UseCase(ctx context.Context, in model.RefreshTokenV1Request) (model.RefreshTokenV1Response, *error_v1.ErrorResponse)
	AuthorizeUserV1UseCase(ctx context.Context, in model.AuthorizeUserV1Request) error
	AddUserRoleV1UseCase(ctx context.Context, in model.AddUserRoleV1Request) (model.AddUserRoleV1Response, *error_v1.ErrorResponse)
	GetUsersV1UseCase(ctx context.Context, in model.GetUsersV1Request) (model.GetUsersV1Response, *error_v1.ErrorResponse)
}

//go:generate go run github.com/vektra/mockery/v2@v2.45.1 --name=PaymentUseCase
type PaymentUseCase interface {
	CreateAccountV1UseCase(ctx context.Context, in model.CreateAccountV1Request, userID string) (model.CreateAccountV1Response, *error_v1.ErrorResponse)
	CreatePaymentV1UseCase(ctx context.Context, in model.CreatePaymentV1Request, userID string) (model.CreatePaymentV1Response, *error_v1.ErrorResponse)
	ConfirmationPaymentV1UseCase(ctx context.Context, in model.ConfirmationPaymentV1Request, userID string) (model.VerifyPaymentV1Response, *error_v1.ErrorResponse)
	GetPaymentV1UseCase(ctx context.Context, in model.GetPaymentV1Request, userID string) (model.GetPaymentV1Response, *error_v1.ErrorResponse)
}

//go:generate go run github.com/vektra/mockery/v2@v2.45.1 --name=CreditUseCase
type CreditUseCase interface {
	CreateCreditV1UseCase(ctx context.Context, in model.CreateCreditV1Request, userID string) (model.CreateCreditV1Response, *error_v1.ErrorResponse)
	CreateCreditApplicationV1UseCase(ctx context.Context, in model.CreateCreditApplicationV1Request, userID string) (model.CreateCreditApplicationV1Response, *error_v1.ErrorResponse)
	CreditApplicationConfirmationV1UseCase(ctx context.Context, in model.CreditApplicationConfirmationV1Request, userID string) (model.CreditApplicationConfirmationV1Response, *error_v1.ErrorResponse)
	GetCreditApplicationV1UseCase(ctx context.Context, applicationID, userID string) (model.GetCreditApplicationV1Response, *error_v1.ErrorResponse)
	UpdateCreditApplicationStatusV1UseCase(ctx context.Context, in model.UpdateCreditApplicationStatusV1Request, applicationID string) (model.UpdateCreditApplicationStatusV1Response, *error_v1.ErrorResponse)
	GetCreditV1UseCase(ctx context.Context, creditID, userID string) (model.GetCreditV1Response, *error_v1.ErrorResponse)
	GetListUserCreditsV1UseCase(ctx context.Context, in model.GetListUserCreditsV1Request) (model.GetListUserCreditsV1Response, *error_v1.ErrorResponse)
	GetPaymentScheduleV1UseCase(ctx context.Context, userID, creditID string) (model.GetPaymentScheduleV1Response, *error_v1.ErrorResponse)
}

//go:generate go run github.com/vektra/mockery/v2@v2.45.1 --name=ReportUseCase
type ReportUseCase interface {
	GetUserReportV1UseCase(ctx context.Context, in model.GetUserReportV1Request) (model.GetUserReportV1Response, *error_v1.ErrorResponse)
}

//go:generate go run github.com/vektra/mockery/v2@v2.45.1 --name=LoggerUseCase
type LoggerUseCase interface {
	GetLogReportV1UseCase(ctx context.Context, in model.GetLogReportV1Request) (model.GetLogReportV1Response, *error_v1.ErrorResponse)
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
