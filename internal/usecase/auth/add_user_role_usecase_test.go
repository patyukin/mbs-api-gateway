package auth

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/patyukin/mbs-api-gateway/internal/usecase/auth/mocks"
	authpb "github.com/patyukin/mbs-pkg/pkg/proto/auth_v1"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
	"github.com/stretchr/testify/assert"
)

func TestAddUserRoleV1UseCase(t *testing.T) {
	ctx := context.Background()

	t.Run(
		"Success case", func(t *testing.T) {
			mockAuthClient := &mocks.ProtoClient{}
			uc := UseCase{authClient: mockAuthClient, jwtSecret: []byte("secret")}

			input := model.AddUserRoleV1Request{
				UserID: "12345",
				RoleID: "54321",
			}
			mockProtoReq := &authpb.AddUserRoleRequest{
				UserId: "12345",
				RoleId: "54321",
			}
			mockProtoResp := &authpb.AddUserRoleResponse{
				Message: "success",
			}
			expectedOutput := model.AddUserRoleV1Response{
				Message: "success",
			}

			mockAuthClient.On("AddUserRole", ctx, mockProtoReq).Return(mockProtoResp, nil)

			// Act
			output, errResp := uc.AddUserRoleV1UseCase(ctx, input)

			// Assert
			assert.Equal(t, expectedOutput, output)
			assert.Nil(t, errResp)
			mockAuthClient.AssertExpectations(t)
		},
	)

	t.Run(
		"Error from authClient", func(t *testing.T) {
			mockAuthClient := &mocks.ProtoClient{}
			uc := UseCase{authClient: mockAuthClient, jwtSecret: []byte("secret")}

			input := model.AddUserRoleV1Request{
				UserID: "12345",
				RoleID: "54321",
			}
			mockProtoReq := &authpb.AddUserRoleRequest{
				UserId: "12345",
				RoleId: "54321",
			}
			expectedError := &error_v1.ErrorResponse{
				Code:        http.StatusInternalServerError,
				Message:     "Internal Server Error",
				Description: "failed to uc.authClient.AddUserRole: failed to uc.authClient.AddUserRole",
			}

			mockAuthClient.On("AddUserRole", ctx, mockProtoReq).Return(nil, errors.New("failed to uc.authClient.AddUserRole"))

			// Act
			output, errResp := uc.AddUserRoleV1UseCase(ctx, input)

			// Assert
			assert.Equal(t, model.AddUserRoleV1Response{}, output)
			assert.NotNil(t, errResp)
			assert.Equal(t, expectedError.Code, errResp.Code)
			assert.Contains(t, errResp.Description, expectedError.Description)
			mockAuthClient.AssertExpectations(t)
		},
	)

	t.Run(
		"Error in response", func(t *testing.T) {
			mockAuthClient := &mocks.ProtoClient{}
			uc := UseCase{authClient: mockAuthClient, jwtSecret: []byte("secret")}

			input := model.AddUserRoleV1Request{
				UserID: "12345",
				RoleID: "admin",
			}
			mockProtoReq := &authpb.AddUserRoleRequest{
				UserId: "12345",
				RoleId: "admin",
			}
			mockProtoResp := &authpb.AddUserRoleResponse{
				Error: &error_v1.ErrorResponse{
					Code:    int32(http.StatusBadRequest),
					Message: "Invalid role",
				},
			}
			expectedError := &error_v1.ErrorResponse{
				Code:    int32(http.StatusBadRequest),
				Message: "Invalid role",
			}

			mockAuthClient.On("AddUserRole", ctx, mockProtoReq).Return(mockProtoResp, nil)

			// Act
			output, errResp := uc.AddUserRoleV1UseCase(ctx, input)

			// Assert
			assert.Equal(t, model.AddUserRoleV1Response{}, output)
			assert.NotNil(t, errResp)
			assert.Equal(t, expectedError.Code, errResp.Code)
			assert.Contains(t, errResp.Message, expectedError.Message)
			mockAuthClient.AssertExpectations(t)
		},
	)
}
