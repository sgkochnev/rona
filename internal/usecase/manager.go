package usecase

var _ Usecase = (*manager)(nil)

type manager struct {
	*auth
}

func NewManager(authRepo AuthRepository, signedKey []byte) *manager {
	return &manager{
		auth: NewAuth(authRepo, signedKey),
	}
}
