package usecase

type AuthUseCase interface {
}

type UseCase struct {
	jwtSecret []byte
}

func New(auth AuthUseCase, jwtSecret []byte) *UseCase {
	return &UseCase{
		jwtSecret: jwtSecret,
	}
}

func (uc *UseCase) GetJWTToken() []byte {
	return uc.jwtSecret
}
