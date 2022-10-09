package repo

import (
	"context"
	"github.com/sgkochnev/rona/internal/entity"
)

type AuthRepository interface {
	Add(context.Context, *entity.User) error
	Get(context.Context, string) (*entity.User, error)
	UpdateRefreshToken(context.Context, *entity.User) error
}

type Repository interface {
	AuthRepository
}
