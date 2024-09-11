package usecase

type UseCase struct {
	jwtSecret []byte
}

func New(jwtSecret []byte) *UseCase {
	return &UseCase{
		jwtSecret: jwtSecret,
	}
}

func (uc *UseCase) GetJWTToken() []byte {
	return uc.jwtSecret
}
