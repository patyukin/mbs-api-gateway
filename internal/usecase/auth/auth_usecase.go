package auth

import (
	"context"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	authpb "github.com/patyukin/mbs-api-gateway/pkg/auth_v1"
	"google.golang.org/grpc"
)

type ProtoClient interface {
	SignUp(ctx context.Context, in *authpb.SignUpRequest, opts ...grpc.CallOption) (*authpb.SignUpResponse, error)
	SignIn(ctx context.Context, in *authpb.SignInRequest, opts ...grpc.CallOption) (*authpb.SignInResponse, error)
	SignInVerify(ctx context.Context, in *authpb.SignInVerifyRequest, opts ...grpc.CallOption) (*authpb.SignInVerifyResponse, error)
}

type UseCase struct {
	jwtSecret  []byte
	authClient ProtoClient
}

func (uc *UseCase) SignInVerifyV1(_ context.Context, _ model.SignInVerifyV1Request) (model.SignInVerifyV1Response, error) {
	//TODO implement me
	panic("implement me")
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
