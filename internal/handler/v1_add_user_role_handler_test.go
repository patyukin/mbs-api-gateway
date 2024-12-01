package handler

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/patyukin/mbs-api-gateway/internal/handler/mocks"
	"github.com/patyukin/mbs-api-gateway/internal/metrics"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type AddUserRoleV1TestSuite struct {
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

func TestAddUserRoleV1TestSuiteTestSuite(t *testing.T) {
	suite.Run(t, new(AddUserRoleV1TestSuite))
}

func (suite *AddUserRoleV1TestSuite) SetupTest() {
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

func (suite *AddUserRoleV1TestSuite) TearDownTest() {
	metrics.TotalRegistrations = nil
	metrics.FailedRegistrations = nil
	metrics.SuccessfulRegistrations = nil
}

func (suite *AddUserRoleV1TestSuite) TestAddUserRoleV1_Success() {
	userUUID, err := uuid.NewUUID()
	suite.Require().NoError(err)

	roleUUID, err := uuid.NewUUID()
	suite.Require().NoError(err)

	reqBody := model.AddUserRoleV1Request{UserID: userUUID.String(), RoleID: roleUUID.String()}
	body, err := json.Marshal(reqBody)
	suite.Require().NoError(err)

	response := model.AddUserRoleV1Response{
		Message: "Success",
	}
	suite.mockAUC.On("AddUserRoleV1UseCase", mock.Anything, reqBody).Return(response, nil)

	req, err := http.NewRequest("POST", "/add-user-role", bytes.NewBuffer(body))
	suite.Require().NoError(err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	suite.handler.AddUserRoleV1Handler(rr, req)

	suite.Equal(http.StatusOK, rr.Code)

	var resp model.AddUserRoleV1Response
	err = json.NewDecoder(rr.Body).Decode(&resp)
	suite.Require().NoError(err)
	suite.Equal(response, resp)

	suite.mockAUC.AssertExpectations(suite.T())
}
