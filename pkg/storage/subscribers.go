package storage

import (
	"context"

	"github.com/broadcastertools/cliprate/core/api"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	SubscribersStore interface {
		FindByIdentifier(ctx context.Context, subscriberID string) (*api.Subscriber, error)
		FindByTwitchIdentifier(ctx context.Context, twitchID string) (*api.Subscriber, error)
	}

	mongoSubscribersStore struct {
		c *mongo.Collection
	}
)

func newMongoStorageClient(c *mongo.Collection) *mongoSubscribersStore {
	return &mongoSubscribersStore{c: c}
}

func (s *mongoSubscribersStore) FindByIdentifier(ctx context.Context, subscriberID string) (*api.Subscriber, error) {
	return nil, ErrNotFound
}

func (s *mongoSubscribersStore) FindByTwitchIdentifier(ctx context.Context, twitchID string) (*api.Subscriber, error) {
	return nil, ErrNotFound
}
