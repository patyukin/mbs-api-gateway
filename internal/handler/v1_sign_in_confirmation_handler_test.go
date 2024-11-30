package handler

import (
	"bytes"
	"encoding/json"
	"github.com/patyukin/mbs-api-gateway/internal/handler/mocks"
	"github.com/patyukin/mbs-api-gateway/internal/metrics"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type SignInConfirmationV1TestSuite struct {
	suite.Suite
	handler         *Handler
	mockAUC         *mocks.AuthUseCase
	mockLUC         *mocks.LoggerUseCase
	mockPUC         *mocks.PaymentUseCase
	mockCUC         *mocks.CreditUseCase
	mockRUC         *mocks.ReportUseCase
	TotalLogin      *mocks.Counter
	FailedLogin     *mocks.Counter
	SuccessfulLogin *mocks.Counter
}

func (suite *SignInConfirmationV1TestSuite) SetupTest() {
	suite.mockAUC = &mocks.AuthUseCase{}
	suite.mockLUC = &mocks.LoggerUseCase{}
	suite.mockPUC = &mocks.PaymentUseCase{}
	suite.mockCUC = &mocks.CreditUseCase{}
	suite.mockRUC = &mocks.ReportUseCase{}
	suite.handler = New(suite.mockAUC, suite.mockLUC, suite.mockPUC, suite.mockCUC, suite.mockRUC)

	suite.TotalLogin = new(mocks.Counter)
	suite.FailedLogin = new(mocks.Counter)
	suite.SuccessfulLogin = new(mocks.Counter)

	metrics.TotalLogin = suite.TotalLogin
	metrics.FailedLogin = suite.FailedLogin
	metrics.SuccessfulLogin = suite.SuccessfulLogin

	suite.TotalLogin.On("Inc").Return(nil)
	suite.FailedLogin.On("Inc").Return(nil)
	suite.SuccessfulLogin.On("Inc").Return(nil)
}

func (suite *SignInConfirmationV1TestSuite) TearDownTest() {
	metrics.TotalRegistrations = nil
	metrics.FailedRegistrations = nil
	metrics.SuccessfulRegistrations = nil
}

func TestSignInConfirmationV1TestSuite(t *testing.T) {
	suite.Run(t, new(SignInConfirmationV1TestSuite))
}
func (suite *SignInConfirmationV1TestSuite) TestSignInConfirmationV1_Success() {
	requestData := model.SignInConfirmationV1Request{
		Code: "12345",
	}

	responseData := model.SignInConfirmationV1Response{
		AccessToken:  "new.jwt.token",
		RefreshToken: "new.refresh.token",
	}

	suite.mockAUC.
		On("SignInConfirmationV1UseCase", mock.Anything, requestData).
		Return(responseData, nil)

	body, err := json.Marshal(requestData)
	suite.NoError(err)

	req, err := http.NewRequest(http.MethodPost, "/sign-in-confirmation", bytes.NewReader(body))
	suite.NoError(err)

	rr := httptest.NewRecorder()

	suite.handler.SignInConfirmationHandler(rr, req)

	// Проверка кода статуса
	suite.Equal(http.StatusOK, rr.Code)

	var resp model.SignInConfirmationV1Response
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	suite.NoError(err)
	suite.Equal(responseData, resp)

	suite.mockAUC.AssertExpectations(suite.T())
	suite.TotalLogin.AssertCalled(suite.T(), "Inc")
	suite.FailedLogin.AssertCalled(suite.T(), "Inc")
	suite.SuccessfulLogin.AssertNotCalled(suite.T(), "Inc")
}
