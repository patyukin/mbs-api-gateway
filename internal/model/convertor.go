package model

import (
	authpb "github.com/patyukin/mbs-pkg/pkg/proto/auth_v1"
	loggerpb "github.com/patyukin/mbs-pkg/pkg/proto/logger_v1"
	paymentpb "github.com/patyukin/mbs-pkg/pkg/proto/payment_v1"
)

func ToProtoSignUpFromRequest(in SignUpV1Request) authpb.SignUpRequest {
	return authpb.SignUpRequest{
		Email:         string(in.Email),
		Password:      string(in.Password),
		TelegramLogin: in.TelegramLogin,
		FirstName:     in.FirstName,
		LastName:      in.LastName,
		Patronymic:    in.Patronymic,
		DateOfBirth:   in.DateOfBirth,
		Phone:         in.Phone,
		Address:       in.Address,
	}
}

func ToProtoSignInVerifyFromRequest(in SignInVerifyV1Request) authpb.SignInConfirmationRequest {
	return authpb.SignInConfirmationRequest{
		Code: in.Code,
	}
}

func ToProtoSignInFromRequest(in SignInV1Request) authpb.SignInRequest {
	return authpb.SignInRequest{
		Email:    in.Login,
		Password: in.Password,
	}
}

func FromProtoSignInVerifyToResponse(in *authpb.SignInConfirmationResponse) SignInVerifyV1Response {
	return SignInVerifyV1Response{
		AccessToken:  in.AccessToken,
		RefreshToken: in.RefreshToken,
	}
}

func ToProtoAuthorizeFromRequest(in AuthorizeRequest) authpb.AuthorizeUserRequest {
	return authpb.AuthorizeUserRequest{
		UserId:    in.UserID,
		RoutePath: in.RoutePath,
		Method:    in.Method,
	}
}

func ToProtoRefreshTokenFromRequest(in RefreshTokenV1Request) authpb.RefreshTokenRequest {
	return authpb.RefreshTokenRequest{
		RefreshToken: in.RefreshToken,
	}
}

func FromProtoRefreshTokenToResponse(in *authpb.RefreshTokenResponse) RefreshTokenV1Response {
	return RefreshTokenV1Response{
		AccessToken: in.AccessToken,
	}
}

func ToProtoLogReportFromRequest(in GetLogReportV1Request) (loggerpb.LogReportRequest, error) {
	return loggerpb.LogReportRequest{
		StartTime:   in.StartDate,
		EndTime:     in.EndDate,
		ServiceName: in.ServiceName,
	}, nil
}

func ToProtoCreateAccountFromRequest(in CreateAccountV1Request) paymentpb.CreateAccountRequest {
	return paymentpb.CreateAccountRequest{
		UserId:   in.UserID,
		Currency: in.Currency,
		Balance:  in.Balance,
	}
}

func ToProtoCreatePaymentFromRequest(in CreatePaymentV1Request) paymentpb.CreatePaymentRequest {
	return paymentpb.CreatePaymentRequest{
		SenderAccountId:   in.SenderAccountID,
		ReceiverAccountId: in.ReceiverAccountID,
		Amount:            in.Amount,
		Currency:          in.Currency,
		Description:       in.Description,
		UserId:            in.UserID,
	}
}

func ToProtoVerifyPaymentFromRequest(in VerifyPaymentV1Request) paymentpb.ConfirmationPaymentRequest {
	return paymentpb.ConfirmationPaymentRequest{
		UserId: in.UserID,
		Code:   in.Code,
	}
}
