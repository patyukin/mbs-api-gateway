package model

import (
	authpb "github.com/patyukin/mbs-pkg/pkg/proto/auth_v1"
)

func ToProtoV1SignUpRequest(in *SignUpV1Request) *authpb.SignUpRequest {
	return &authpb.SignUpRequest{
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
		Code:  in.Code,
		Login: in.Login,
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

func ToProtoAddUserRoleV1Request(in AddUserRoleV1Request) authpb.AddUserRoleRequest {
	return authpb.AddUserRoleRequest{
		UserId: in.UserID,
		RoleId: in.RoleID,
	}
}

func ToModelAddUserRoleV1Response(in *authpb.AddUserRoleResponse) AddUserRoleV1Response {
	return AddUserRoleV1Response{
		Message: in.GetMessage(),
	}
}

func ToProtoGetUserByIDV1Request(in GetUserByIDV1Request) authpb.GetUserByIDRequest {
	return authpb.GetUserByIDRequest{
		UserId: in.UserID,
	}
}

func ToModelUserInfoV1(in *authpb.UserInfo) UserInfoV1 {
	return UserInfoV1{
		ID:    in.GetId(),
		Email: in.GetEmail(),
		ProfileV1: ProfileV1{
			FirstName:   in.GetProfile().GetFirstName(),
			LastName:    in.GetProfile().GetLastName(),
			Patronymic:  in.GetProfile().GetPatronymic(),
			DateOfBirth: in.GetProfile().GetDateOfBirth(),
			Phone:       in.GetProfile().GetPhone(),
			Address:     in.GetProfile().GetAddress(),
		},
	}
}

func ToModelGetUserByIDV1Response(in *authpb.GetUserByIDResponse) GetUserByIDV1Response {
	return GetUserByIDV1Response{
		UserInfoV1: ToModelUserInfoV1(in.GetUser()),
	}
}

func ToProtoGetUsersV1Request(in GetUsersV1Request) authpb.GetUsersRequest {
	return authpb.GetUsersRequest{
		Page:  in.Page,
		Limit: in.Limit,
	}
}

func ToModelUsersV1(in []*authpb.UserInfo) []UserInfoV1 {
	users := make([]UserInfoV1, 0, len(in))
	for _, user := range in {
		users = append(users, ToModelUserInfoV1(user))
	}
	return users
}

func ToModelGetUsersV1Response(in *authpb.GetUsersResponse) GetUsersV1Response {
	return GetUsersV1Response{
		UsersInfoV1: ToModelUsersV1(in.GetUsers()),
		Total:       in.GetTotal(),
	}
}
