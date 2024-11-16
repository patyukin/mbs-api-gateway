package auth

import (
	"context"
	authpb "github.com/patyukin/mbs-pkg/pkg/proto/auth_v1"
	"google.golang.org/grpc"
)

type ProtoClient interface {
	SignUp(ctx context.Context, in *authpb.SignUpRequest, opts ...grpc.CallOption) (*authpb.SignUpResponse, error)
	SignIn(ctx context.Context, in *authpb.SignInRequest, opts ...grpc.CallOption) (*authpb.SignInResponse, error)
	SignInVerify(ctx context.Context, in *authpb.SignInVerifyRequest, opts ...grpc.CallOption) (*authpb.SignInVerifyResponse, error)
	Authorize(ctx context.Context, in *authpb.AuthorizeRequest, opts ...grpc.CallOption) (*authpb.AuthorizeResponse, error)
	RefreshToken(ctx context.Context, in *authpb.RefreshTokenRequest, opts ...grpc.CallOption) (*authpb.RefreshTokenResponse, error)
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
