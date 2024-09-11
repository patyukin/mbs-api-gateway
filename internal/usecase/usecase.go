package usecase

type Usecase struct {
	jwtSecret []byte
}

func New(jwtSecret []byte) *Usecase {
	return &Usecase{
		jwtSecret: jwtSecret,
	}
}
