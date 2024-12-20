package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepo struct {
	Client *mongo.Client
	Coll   *mongo.Collection
}

func NewMongoClient(uri, db_name, default_collection string) (*MongoRepo, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))

	if err != nil {
		return nil, err
	}

	coll := client.Database(db_name).Collection(default_collection)

	return &MongoRepo{
		Client: client,
		Coll:   coll,
	}, nil
}

func (m *MongoRepo) CloseDB() error {
	if err := m.Client.Disconnect(context.Background()); err != nil {
		return err
	}
	return nil
}
