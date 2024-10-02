package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/patyukin/mbs-api-gateway/internal/handler/mocks"
	"github.com/patyukin/mbs-api-gateway/internal/metrics"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type SignUpV1TestSuite struct {
	suite.Suite
	handler                     *Handler
	mockUC                      *mocks.AuthUseCase
	mockTotalRegistrations      *mocks.Counter
	mockFailedRegistrations     *mocks.Counter
	mockSuccessfulRegistrations *mocks.Counter
}

// SetupTest выполняется перед каждым тестом
func (suite *SignUpV1TestSuite) SetupTest() {
	suite.mockUC = &mocks.AuthUseCase{}
	suite.handler = New(suite.mockUC)

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
	suite.mockUC.On("SignUpV1", mock.Anything, mock.Anything).Return(nil)

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

	body, _ := json.Marshal(requestData)
	req, err := http.NewRequest(http.MethodPost, "/sign-up", bytes.NewReader(body))
	suite.NoError(err)

	rr := httptest.NewRecorder()

	suite.handler.SignUpV1(rr, req)

	suite.Equal(http.StatusCreated, rr.Code)
	suite.Equal("application/json; charset=UTF-8", rr.Header().Get("Content-Type"))
	suite.mockUC.AssertExpectations(suite.T())
}

func (suite *SignUpV1TestSuite) TestSignUpV1_DecodeError() {
	req, err := http.NewRequest(http.MethodPost, "/sign-up", bytes.NewReader([]byte("invalid json")))
	suite.NoError(err)

	rr := httptest.NewRecorder()

	suite.handler.SignUpV1(rr, req)

	suite.Equal(http.StatusBadRequest, rr.Code)

	suite.mockUC.AssertNotCalled(suite.T(), "SignUpV1", mock.Anything, mock.Anything)
}

func (suite *SignUpV1TestSuite) TestSignUpV1_ValidationError() {
	requestData := model.SignUpV1Request{
		Email:         "john.doeexample.com",
		Password:      "securepassword123",
		TelegramLogin: "johndoe_telegram",
		FirstName:     "John",
		LastName:      "Doe",
		Patronymic:    "Jonathan",
		DateOfBirth:   "1990-01-01",
		Phone:         "79000000000",
		Address:       "Moscow, Russia, 1, 1, 1",
	}

	body, _ := json.Marshal(requestData)
	req, err := http.NewRequest(http.MethodPost, "/sign-up", bytes.NewReader(body))
	suite.NoError(err)

	rr := httptest.NewRecorder()

	suite.handler.SignUpV1(rr, req)
	suite.Equal(http.StatusBadRequest, rr.Code)
	suite.mockUC.AssertNotCalled(suite.T(), "SignUpV1", mock.Anything, mock.Anything)
}

func (suite *SignUpV1TestSuite) TestSignUpV1_UseCaseError() {
	suite.mockUC.On("SignUpV1", mock.Anything, mock.Anything).Return(errors.New("use case error"))

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

	body, _ := json.Marshal(requestData)
	req, err := http.NewRequest(http.MethodPost, "/sign-up", bytes.NewReader(body))
	suite.NoError(err)

	rr := httptest.NewRecorder()

	suite.handler.SignUpV1(rr, req)
	suite.Equal(http.StatusInternalServerError, rr.Code)

	suite.mockUC.AssertExpectations(suite.T())
}

func TestSignUpV1TestSuite(t *testing.T) {
	suite.Run(t, new(SignUpV1TestSuite))
}
