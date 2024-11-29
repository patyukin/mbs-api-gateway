package handler

import (
	"bytes"
	"encoding/json"
	"github.com/patyukin/mbs-api-gateway/internal/handler/mocks"
	"github.com/patyukin/mbs-api-gateway/internal/metrics"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type SignUpV1TestSuite struct {
	suite.Suite
	handler                     *Handler
	mockAUC                     *mocks.AuthUseCase
	mockLUC                     *mocks.LoggerUseCase
	mockPUC                     *mocks.PaymentUseCase
	mockCUC                     *mocks.CreditUseCase
	mockRUC                     *mocks.ReportUseCase
	mockTotalRegistrations      *mocks.Counter
	mockFailedRegistrations     *mocks.Counter
	mockSuccessfulRegistrations *mocks.Counter
}

func (suite *SignUpV1TestSuite) SetupTest() {
	suite.mockAUC = &mocks.AuthUseCase{}
	suite.mockLUC = &mocks.LoggerUseCase{}
	suite.mockPUC = &mocks.PaymentUseCase{}
	suite.mockCUC = &mocks.CreditUseCase{}
	suite.mockRUC = &mocks.ReportUseCase{}
	suite.handler = New(suite.mockAUC, suite.mockLUC, suite.mockPUC, suite.mockCUC, suite.mockRUC)

	suite.mockTotalRegistrations = new(mocks.Counter)
	suite.mockFailedRegistrations = new(mocks.Counter)
	suite.mockSuccessfulRegistrations = new(mocks.Counter)

	metrics.TotalRegistrations = suite.mockTotalRegistrations
	metrics.FailedRegistrations = suite.mockFailedRegistrations
	metrics.SuccessfulRegistrations = suite.mockSuccessfulRegistrations

	suite.mockTotalRegistrations.On("Inc").Return(nil)
	suite.mockFailedRegistrations.On("Inc").Return(nil)
	suite.mockSuccessfulRegistrations.On("Inc").Return(nil)
}

func (suite *SignUpV1TestSuite) TearDownTest() {
	metrics.TotalRegistrations = nil
	metrics.FailedRegistrations = nil
	metrics.SuccessfulRegistrations = nil
}

func (suite *SignUpV1TestSuite) TestSignUpV1_Success() {
	requestData := model.SignUpV1Request{
		Email:         "john.doe@example.com",
		Password:      "securepassword123",
		TelegramLogin: "johndoe_telegram",
		FirstName:     "John",
		LastName:      "Doe",
		Patronymic:    "Jonathan",
		DateOfBirth:   "1990-01-01",
		Phone:         "79000000000",
		Address:       "Moscow, Russia, 1, 1, 1",
	}

	responseData := model.SignUpV1Response{Message: "success"}

	suite.mockAUC.On("SignUpV1UseCase", mock.Anything, requestData).Return(responseData, nil)

	body, err := json.Marshal(requestData)
	suite.NoError(err)

	req, err := http.NewRequest(http.MethodPost, "/sign-up", bytes.NewReader(body))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	suite.handler.SignUpV1Handler(rr, req)
	suite.Equal(http.StatusCreated, rr.Code)

	suite.Equal("application/json; charset=UTF-8", rr.Header().Get("Content-Type"))

	var resp model.SignUpV1Response
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	suite.NoError(err)
	suite.Equal(responseData, resp)

	suite.mockAUC.AssertExpectations(suite.T())
	suite.mockTotalRegistrations.AssertCalled(suite.T(), "Inc")
	suite.mockSuccessfulRegistrations.AssertCalled(suite.T(), "Inc")
	suite.mockFailedRegistrations.AssertNotCalled(suite.T(), "Inc")
}

func (suite *SignUpV1TestSuite) TestSignUpV1_DecodeError() {
	// Создание запроса с некорректным JSON
	req, err := http.NewRequest(http.MethodPost, "/sign-up", bytes.NewReader([]byte("invalid json")))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	// Создание рекордера ответа
	rr := httptest.NewRecorder()

	// Вызов обработчика
	suite.handler.SignUpV1Handler(rr, req)

	// Проверка кода статуса
	suite.Equal(http.StatusBadRequest, rr.Code)

	// Проверка тела ответа
	expectedBody := `{"error":"invalid data"}`
	suite.JSONEq(expectedBody, rr.Body.String())

	// Проверка, что UseCase не был вызван
	suite.mockAUC.AssertNotCalled(suite.T(), "SignUpV1UseCase", mock.Anything, mock.Anything)
	suite.mockFailedRegistrations.AssertCalled(suite.T(), "Inc")
	suite.mockTotalRegistrations.AssertCalled(suite.T(), "Inc")
	suite.mockSuccessfulRegistrations.AssertNotCalled(suite.T(), "Inc")
}

func (suite *SignUpV1TestSuite) TestSignUpV1_ValidationError() {
	requestData := model.SignUpV1Request{
		Email:         "john.doeexample.com",
		Password:      "short",
		TelegramLogin: "johndoe_telegram",
		FirstName:     "John",
		LastName:      "Doe",
		Patronymic:    "Jonathan",
		DateOfBirth:   "1990-01-01",
		Phone:         "79000000000",
		Address:       "Moscow, Russia, 1, 1, 1",
	}

	body, err := json.Marshal(requestData)
	suite.NoError(err)

	req, err := http.NewRequest(http.MethodPost, "/sign-up", bytes.NewReader(body))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	suite.handler.SignUpV1Handler(rr, req)

	suite.Equal(http.StatusBadRequest, rr.Code)

	expectedBody := `{"error":"email: Invalid"}`
	suite.JSONEq(expectedBody, rr.Body.String())

	suite.mockAUC.AssertNotCalled(suite.T(), "SignUpV1UseCase", mock.Anything, mock.Anything)
	suite.mockFailedRegistrations.AssertCalled(suite.T(), "Inc")
	suite.mockTotalRegistrations.AssertCalled(suite.T(), "Inc")
	suite.mockSuccessfulRegistrations.AssertNotCalled(suite.T(), "Inc")
}

func (suite *SignUpV1TestSuite) TestSignUpV1_UseCaseError() {
	requestData := model.SignUpV1Request{
		Email:         "john.doe@example.com",
		Password:      "securepassword123",
		TelegramLogin: "johndoe_telegram",
		FirstName:     "John",
		LastName:      "Doe",
		Patronymic:    "Jonathan",
		DateOfBirth:   "1990-01-01",
		Phone:         "79000000000",
		Address:       "Moscow, Russia, 1, 1, 1",
	}

	useCaseError := &error_v1.ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: "use case error",
	}

	suite.mockAUC.On("SignUpV1UseCase", mock.Anything, requestData).Return(model.SignUpV1Response{}, useCaseError)

	body, err := json.Marshal(requestData)
	suite.NoError(err)

	req, err := http.NewRequest(http.MethodPost, "/sign-up", bytes.NewReader(body))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	suite.handler.SignUpV1Handler(rr, req)

	suite.Equal(useCaseError.Code, int32(rr.Code))

	expectedBody := `{"error":"use case error"}`
	suite.JSONEq(expectedBody, rr.Body.String())

	// Проверка вызовов моков
	suite.mockAUC.AssertExpectations(suite.T())
	suite.mockFailedRegistrations.AssertCalled(suite.T(), "Inc")
	suite.mockTotalRegistrations.AssertCalled(suite.T(), "Inc")
	suite.mockSuccessfulRegistrations.AssertNotCalled(suite.T(), "Inc")
}

func TestSignUpV1TestSuite(t *testing.T) {
	suite.Run(t, new(SignUpV1TestSuite))
}
