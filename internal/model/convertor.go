package model

import (
	"fmt"
	authpb "github.com/patyukin/mbs-pkg/pkg/proto/auth_v1"
	loggerpb "github.com/patyukin/mbs-pkg/pkg/proto/logger_v1"
	"github.com/patyukin/mbs-pkg/pkg/utils"
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

func ToProtoSignInVerifyFromRequest(in SignInVerifyV1Request) authpb.SignInVerifyRequest {
	return authpb.SignInVerifyRequest{
		Code: in.Code,
	}
}

func ToProtoSignInFromRequest(in SignInV1Request) authpb.SignInRequest {
	return authpb.SignInRequest{
		Email:    in.Login,
		Password: in.Password,
	}
}

func FromProtoSignInVerifyToResponse(in *authpb.SignInVerifyResponse) SignInVerifyV1Response {
	return SignInVerifyV1Response{
		AccessToken:  in.AccessToken,
		RefreshToken: in.RefreshToken,
	}
}

func ToProtoAuthorizeFromRequest(in AuthorizeRequest) authpb.AuthorizeRequest {
	return authpb.AuthorizeRequest{
		UserId:    in.UserID,
		RoutePath: in.RoutePath,
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
	isValid, err := utils.ValidateDate(in.StartDate)
	if err != nil {
		return loggerpb.LogReportRequest{}, fmt.Errorf("некорректная дата начала")
	}

	if !isValid {
		return loggerpb.LogReportRequest{}, fmt.Errorf("некорректная дата начала")
	}

	isValid, err = utils.ValidateDate(in.EndDate)
	if err != nil {
		return loggerpb.LogReportRequest{}, fmt.Errorf("некорректная дата окончания")
	}

	if !isValid {
		return loggerpb.LogReportRequest{}, fmt.Errorf("некорректная дата окончания")
	}

	return loggerpb.LogReportRequest{
		StartTime:   in.StartDate,
		EndTime:     in.EndDate,
		ServiceName: in.ServiceName,
	}, nil
}
