package handler

import (
	"bytes"
	"encoding/json"
	"github.com/patyukin/mbs-api-gateway/internal/handler/mocks"
	"github.com/patyukin/mbs-api-gateway/internal/metrics"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type SignInV1TestSuite struct {
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

func (suite *SignInV1TestSuite) SetupTest() {
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

func (suite *SignInV1TestSuite) TearDownTest() {
	metrics.TotalRegistrations = nil
	metrics.FailedRegistrations = nil
	metrics.SuccessfulRegistrations = nil
}

func (suite *SignInV1TestSuite) TestSignInV1_Success() {
	requestData := model.SignInV1Request{
		Login:    "john.doe@example.com",
		Password: "password",
	}

	responseData := model.SignInV1Response{
		Message: "code is sent to telegram",
	}

	suite.mockAUC.On("SignInV1UseCase", mock.Anything, requestData).Return(responseData, nil)

	body, err := json.Marshal(requestData)
	suite.NoError(err)

	req, err := http.NewRequest(http.MethodPost, "/sign-in", bytes.NewReader(body))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	suite.handler.SignInV1Handler(rr, req)

	suite.Equal(http.StatusOK, rr.Code)
	suite.Equal("application/json; charset=UTF-8", rr.Header().Get("Content-Type"))

	var resp model.SignInV1Response
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	suite.NoError(err)
	suite.Equal(responseData, resp)

	suite.mockAUC.AssertExpectations(suite.T())
}

func (suite *SignInV1TestSuite) TestSignInV1_DecodeError() {
	invalidJSON := `{"email": "john.doe@example.com", "password": "123"`

	req, err := http.NewRequest(http.MethodPost, "/sign-in", bytes.NewReader([]byte(invalidJSON)))
	suite.NoError(err)

	rr := httptest.NewRecorder()

	suite.handler.SignInV1Handler(rr, req)

	suite.Equal(http.StatusBadRequest, rr.Code)

	var resp model.ErrorResponse
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	suite.NoError(err)
	suite.Equal("invalid data", resp.Error)

	suite.mockAUC.AssertExpectations(suite.T())
	suite.TotalLogin.AssertCalled(suite.T(), "Inc")
	suite.FailedLogin.AssertCalled(suite.T(), "Inc")
	suite.SuccessfulLogin.AssertNotCalled(suite.T(), "Inc")
}

func (suite *SignInV1TestSuite) TestSignInV1_ValidationError() {
	requestData := model.SignInV1Request{
		Login:    "invalid-email",
		Password: "",
	}

	validationErr := &error_v1.ErrorResponse{
		Code:        http.StatusBadRequest,
		Message:     "login: Invalid",
		Description: "invalid email format and password cannot be empty",
	}

	suite.mockAUC.On("SignInV1UseCase", mock.Anything, requestData).
		Return(model.SignInV1Response{}, validationErr)

	body, err := json.Marshal(requestData)
	suite.NoError(err)

	req, err := http.NewRequest(http.MethodPost, "/sign-in", bytes.NewReader(body))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	suite.handler.SignInV1Handler(rr, req)

	var resp model.ErrorResponse
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	suite.NoError(err)
	suite.Equal(validationErr.Message, resp.Error)

	suite.TotalLogin.AssertCalled(suite.T(), "Inc")
	suite.FailedLogin.AssertCalled(suite.T(), "Inc")
	suite.SuccessfulLogin.AssertNotCalled(suite.T(), "Inc")
}

func (suite *SignInV1TestSuite) TestSignInV1_UseCaseError() {
	requestData := model.SignInV1Request{
		Login:    "john.doe@example.com",
		Password: "wrongpassword",
	}

	useCaseErr := &error_v1.ErrorResponse{
		Code:        401,
		Message:     "authentication failed",
		Description: "invalid credentials",
	}

	suite.mockAUC.
		On("SignInV1UseCase", mock.Anything, requestData).
		Return(model.SignInV1Response{}, useCaseErr)

	body, err := json.Marshal(requestData)
	suite.NoError(err)

	req, err := http.NewRequest(http.MethodPost, "/sign-in", bytes.NewReader(body))
	suite.NoError(err)

	rr := httptest.NewRecorder()

	suite.handler.SignInV1Handler(rr, req)

	suite.Equal(int(useCaseErr.Code), rr.Code)

	var resp model.ErrorResponse
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	suite.NoError(err)

	suite.mockAUC.AssertExpectations(suite.T())
	suite.TotalLogin.AssertCalled(suite.T(), "Inc")
	suite.FailedLogin.AssertCalled(suite.T(), "Inc")
	suite.SuccessfulLogin.AssertNotCalled(suite.T(), "Inc")
}

func TestSignInV1TestSuite(t *testing.T) {
	suite.Run(t, new(SignInV1TestSuite))
}
