package model

import (
	authpb "github.com/patyukin/mbs-api-gateway/proto/auth"
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

func ToProtoSignInFromRequest(in SignInV1Request) authpb.SignInRequest {
	return authpb.SignInRequest{
		Email:    in.Login,
		Password: in.Password,
	}
}

func FromProtoSignInToResponse(in *authpb.SignInResponse) SignInV1Response {
	return SignInV1Response{
		AccessToken:  in.AccessToken,
		RefreshToken: in.RefreshToken,
	}
}
