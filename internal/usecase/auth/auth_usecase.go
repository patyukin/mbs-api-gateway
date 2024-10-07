package auth

import (
	"context"
	authpb "github.com/patyukin/mbs-api-gateway/proto/auth"
)

type ProtoClient interface {
	SignUp(ctx context.Context, in *authpb.SignUpRequest) (*authpb.SignUpResponse, error)
	SignIn(ctx context.Context, in *authpb.SignInRequest) (*authpb.SignInResponse, error)
	SignInVerify(ctx context.Context, in *authpb.SignInVerifyRequest) (*authpb.SignInVerifyResponse, error)
}

type UseCase struct {
	jwtSecret  []byte
	authClient ProtoClient
}

func New(jwtSecret []byte, authClient ProtoClient) *UseCase {
	return &UseCase{
		jwtSecret:  jwtSecret,
		authClient: authClient,
	}
}

func (uc *UseCase) GetJWTToken() []byte {
	return uc.jwtSecret
}
