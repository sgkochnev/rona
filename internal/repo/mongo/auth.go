package mongo

import (
	"context"
	"errors"
	"github.com/sgkochnev/rona/internal/entity"
	e "github.com/sgkochnev/rona/internal/error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepo struct {
	db *MongoDB
}

func NewAuthRepo(db *MongoDB) *AuthRepo {
	return &AuthRepo{db: db}
}

func (repo *AuthRepo) Add(ctx context.Context, user *entity.User) error {
	_, err := repo.db.InsertOne(ctx, *user)
	if err != nil {
		return e.ErrDuplicateEntry
	}

	return nil
}

func (repo *AuthRepo) Get(ctx context.Context, username string) (*entity.User, error) {
	user := &entity.User{}

	err := repo.db.FindOne(ctx, bson.M{"_id": username}).Decode(user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, e.ErrUserDoesNotExist
		}
		return nil, err
	}

	return user, nil
}

func (repo *AuthRepo) UpdateRefreshToken(ctx context.Context, user *entity.User) error {
	_, err := repo.db.UpdateOne(ctx,
		bson.M{"_id": user.Username},
		bson.M{"$set": bson.M{"refresh_token": user.RefreshToken}},
	)

	return err
}
