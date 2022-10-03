package repo

import (
	"github.com/sgkochnev/rona/config"
	"github.com/sgkochnev/rona/internal/repo/mongo"
)

var _ Repository = (*store)(nil)

type store struct {
	*mongo.AuthRepo
}

func NewStore(cfg *config.Config) (*store, error) {
	mDB, err := mongo.Dial(cfg)
	if err != nil {
		return nil, err
	}

	return &store{
		AuthRepo: mongo.NewAuthRepo(mDB),
	}, nil
}
