package auth

type ProtoClient interface {
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
