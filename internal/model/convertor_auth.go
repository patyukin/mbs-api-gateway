package model

import (
	authpb "github.com/patyukin/mbs-pkg/pkg/proto/auth_v1"
)

func ToProtoV1SignUpRequest(in SignUpV1Request) authpb.SignUpRequest {
	return authpb.SignUpRequest{
		Email:         in.Email,
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

func ToProtoV1SignInConfirmationRequest(in SignInConfirmationV1Request) authpb.SignInConfirmationRequest {
	return authpb.SignInConfirmationRequest{
		Code: in.Code,
	}
}

func ToProtoV1SignInRequest(in SignInV1Request) authpb.SignInRequest {
	return authpb.SignInRequest{
		Email:    in.Login,
		Password: in.Password,
	}
}

func ToModelSignInConfirmationV1Response(in *authpb.SignInConfirmationResponse) SignInConfirmationV1Response {
	return SignInConfirmationV1Response{
		AccessToken:  in.GetAccessToken(),
		RefreshToken: in.GetRefreshToken(),
	}
}

func ToProtoV1AuthorizeUserRequest(in AuthorizeUserV1Request) authpb.AuthorizeUserRequest {
	return authpb.AuthorizeUserRequest{
		UserId:    in.UserID,
		RoutePath: in.RoutePath,
		Method:    in.Method,
	}
}

func ToProtoV1RefreshTokenRequest(in RefreshTokenV1Request) authpb.RefreshTokenRequest {
	return authpb.RefreshTokenRequest{
		RefreshToken: in.RefreshToken,
	}
}

func ToModelRefreshTokenV1Response(in *authpb.RefreshTokenResponse) RefreshTokenV1Response {
	return RefreshTokenV1Response{
		AccessToken: in.GetAccessToken(),
	}
}
