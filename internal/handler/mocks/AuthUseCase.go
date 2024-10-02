// Code generated by mockery v2.45.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/patyukin/mbs-api-gateway/internal/model"
)

// AuthUseCase is an autogenerated mock type for the AuthUseCase type
type AuthUseCase struct {
	mock.Mock
}

// GetJWTToken provides a mock function with given fields:
func (_m *AuthUseCase) GetJWTToken() []byte {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetJWTToken")
	}

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// SignUpV1 provides a mock function with given fields: ctx, in
func (_m *AuthUseCase) SignUpV1(ctx context.Context, in model.SignUpV1Request) error {
	ret := _m.Called(ctx, in)

	if len(ret) == 0 {
		panic("no return value specified for SignUpV1")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.SignUpV1Request) error); ok {
		r0 = rf(ctx, in)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewAuthUseCase creates a new instance of AuthUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthUseCase {
	mock := &AuthUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
