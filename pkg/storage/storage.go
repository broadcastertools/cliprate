package storage

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

type (
	MongoStorageClient struct {
		db *mongo.Database
	}
)

var (
	ErrNotFound = errors.New("not found")
)

func NewMongoStorageClient(ctx context.Context, uri string) (*MongoStorageClient, error) {
	c, err := connstring.Parse(uri)
	if err != nil {
		return nil, err
	}

	db, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	dbName := c.Database
	if dbName == "" || dbName == "admin" {
		dbName = "cliprate"
	}

	return &MongoStorageClient{db: db.Database(dbName)}, nil
}

func (c *MongoStorageClient) Subscribers() *mongoSubscribersStore {
	return newMongoStorageClient(c.db.Collection("subscribers"))
}
