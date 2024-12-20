// Code generated by mockery v2.45.1. DO NOT EDIT.

package mocks

import (
	context "context"

	error_v1 "github.com/patyukin/mbs-pkg/pkg/proto/error_v1"

	mock "github.com/stretchr/testify/mock"

	model "github.com/patyukin/mbs-api-gateway/internal/model"
)

// PaymentUseCase is an autogenerated mock type for the PaymentUseCase type
type PaymentUseCase struct {
	mock.Mock
}

// ConfirmationPaymentV1UseCase provides a mock function with given fields: ctx, in, userID
func (_m *PaymentUseCase) ConfirmationPaymentV1UseCase(ctx context.Context, in model.ConfirmationPaymentV1Request, userID string) (model.VerifyPaymentV1Response, *error_v1.ErrorResponse) {
	ret := _m.Called(ctx, in, userID)

	if len(ret) == 0 {
		panic("no return value specified for ConfirmationPaymentV1UseCase")
	}

	var r0 model.VerifyPaymentV1Response
	var r1 *error_v1.ErrorResponse
	if rf, ok := ret.Get(0).(func(context.Context, model.ConfirmationPaymentV1Request, string) (model.VerifyPaymentV1Response, *error_v1.ErrorResponse)); ok {
		return rf(ctx, in, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.ConfirmationPaymentV1Request, string) model.VerifyPaymentV1Response); ok {
		r0 = rf(ctx, in, userID)
	} else {
		r0 = ret.Get(0).(model.VerifyPaymentV1Response)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.ConfirmationPaymentV1Request, string) *error_v1.ErrorResponse); ok {
		r1 = rf(ctx, in, userID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*error_v1.ErrorResponse)
		}
	}

	return r0, r1
}

// CreateAccountV1UseCase provides a mock function with given fields: ctx, in, userID
func (_m *PaymentUseCase) CreateAccountV1UseCase(ctx context.Context, in model.CreateAccountV1Request, userID string) (model.CreateAccountV1Response, *error_v1.ErrorResponse) {
	ret := _m.Called(ctx, in, userID)

	if len(ret) == 0 {
		panic("no return value specified for CreateAccountV1UseCase")
	}

	var r0 model.CreateAccountV1Response
	var r1 *error_v1.ErrorResponse
	if rf, ok := ret.Get(0).(func(context.Context, model.CreateAccountV1Request, string) (model.CreateAccountV1Response, *error_v1.ErrorResponse)); ok {
		return rf(ctx, in, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.CreateAccountV1Request, string) model.CreateAccountV1Response); ok {
		r0 = rf(ctx, in, userID)
	} else {
		r0 = ret.Get(0).(model.CreateAccountV1Response)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.CreateAccountV1Request, string) *error_v1.ErrorResponse); ok {
		r1 = rf(ctx, in, userID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*error_v1.ErrorResponse)
		}
	}

	return r0, r1
}

// CreatePaymentV1UseCase provides a mock function with given fields: ctx, in, userID
func (_m *PaymentUseCase) CreatePaymentV1UseCase(ctx context.Context, in *model.CreatePaymentV1Request, userID string) (model.CreatePaymentV1Response, *error_v1.ErrorResponse) {
	ret := _m.Called(ctx, in, userID)

	if len(ret) == 0 {
		panic("no return value specified for CreatePaymentV1UseCase")
	}

	var r0 model.CreatePaymentV1Response
	var r1 *error_v1.ErrorResponse
	if rf, ok := ret.Get(0).(func(context.Context, *model.CreatePaymentV1Request, string) (model.CreatePaymentV1Response, *error_v1.ErrorResponse)); ok {
		return rf(ctx, in, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.CreatePaymentV1Request, string) model.CreatePaymentV1Response); ok {
		r0 = rf(ctx, in, userID)
	} else {
		r0 = ret.Get(0).(model.CreatePaymentV1Response)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.CreatePaymentV1Request, string) *error_v1.ErrorResponse); ok {
		r1 = rf(ctx, in, userID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*error_v1.ErrorResponse)
		}
	}

	return r0, r1
}

// GetPaymentV1UseCase provides a mock function with given fields: ctx, in, userID
func (_m *PaymentUseCase) GetPaymentV1UseCase(ctx context.Context, in model.GetPaymentV1Request, userID string) (model.GetPaymentV1Response, *error_v1.ErrorResponse) {
	ret := _m.Called(ctx, in, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetPaymentV1UseCase")
	}

	var r0 model.GetPaymentV1Response
	var r1 *error_v1.ErrorResponse
	if rf, ok := ret.Get(0).(func(context.Context, model.GetPaymentV1Request, string) (model.GetPaymentV1Response, *error_v1.ErrorResponse)); ok {
		return rf(ctx, in, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.GetPaymentV1Request, string) model.GetPaymentV1Response); ok {
		r0 = rf(ctx, in, userID)
	} else {
		r0 = ret.Get(0).(model.GetPaymentV1Response)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.GetPaymentV1Request, string) *error_v1.ErrorResponse); ok {
		r1 = rf(ctx, in, userID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*error_v1.ErrorResponse)
		}
	}

	return r0, r1
}

// NewPaymentUseCase creates a new instance of PaymentUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPaymentUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *PaymentUseCase {
	mock := &PaymentUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
