package v1

import (
	"context"

	"github.com/sgkochnev/rona/internal/entity"
)

type TokenParser interface {
	ParseToken(string) (*entity.Claims, error)
}

type Auth interface {
	SignUp(context.Context, *entity.User) error
	SignIn(context.Context, *entity.User) (*entity.Token, error)
}

type Usecases interface {
	Auth
	TokenParser
}
