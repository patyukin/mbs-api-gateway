package auth

import (
	"context"
	authpb "github.com/patyukin/mbs-pkg/pkg/proto/auth_v1"
	"google.golang.org/grpc"
)

type ProtoClient interface {
	SignUp(ctx context.Context, in *authpb.SignUpRequest, opts ...grpc.CallOption) (*authpb.SignUpResponse, error)
	SignIn(ctx context.Context, in *authpb.SignInRequest, opts ...grpc.CallOption) (*authpb.SignInResponse, error)
	SignInConfirmation(ctx context.Context, in *authpb.SignInConfirmationRequest, opts ...grpc.CallOption) (*authpb.SignInConfirmationResponse, error)
	AuthorizeUser(ctx context.Context, in *authpb.AuthorizeUserRequest, opts ...grpc.CallOption) (*authpb.AuthorizeUserResponse, error)
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
