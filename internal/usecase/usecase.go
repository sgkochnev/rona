package usecase

import (
	"context"
	"time"

	"github.com/sgkochnev/rona/internal/entity"
)

const expiresAtForAccess = 15 * time.Minute
const expiresAtForRefresh = 14 * 24 * time.Hour

type AuthRepository interface {
	Add(context.Context, *entity.User) error
	Get(ctx context.Context, username string) (*entity.User, error)
	UpdateRefreshToken(context.Context, *entity.User) error
}

type TokenParser interface {
	ParseToken(string) (*entity.Claims, error)
}

type AuthUsecase interface {
	SignUp(context.Context, *entity.User) error
	SignIn(context.Context, *entity.User) (*entity.Token, error)
}

type Usecase interface {
	AuthUsecase
	TokenParser
}
