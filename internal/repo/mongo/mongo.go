package mongo

import (
	"context"
	"fmt"
	"github.com/sgkochnev/rona/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	timeout = 5 * time.Second
)

type MongoDB struct {
	*mongo.Collection
}

func Dial(cfg *config.Config) (*MongoDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	client, err := mongo.Connect(ctx,
		options.Client().ApplyURI(cfg.Mongo.URI),
		options.Client().SetServerSelectionTimeout(timeout),
	)
	if err != nil {
		return nil, fmt.Errorf("mongo connection failed: %w", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("mongo ping failed: %w", err)
	}

	return &MongoDB{
		Collection: client.
			Database(cfg.Mongo.Name).
			Collection(cfg.Mongo.Collection),
	}, nil
}
