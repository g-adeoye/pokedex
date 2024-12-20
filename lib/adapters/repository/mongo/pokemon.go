package repository

import (
	"context"
	"pokedex/lib/core/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *MongoRepo) GetPokemon(ctx context.Context, pokemon string) (*domain.Pokemon, error) {
	var result domain.Pokemon

	err := m.Coll.FindOne(context.Background(), bson.D{{Key: "name", Value: pokemon}}).Decode(&result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (m *MongoRepo) UpdatePokemon(ctx context.Context, pokemon domain.Pokemon, upsert bool) (*mongo.UpdateResult, error) {
	result, err := m.Coll.UpdateOne(
		context.Background(),
		bson.D{{Key: "name", Value: pokemon.Name}},
		bson.D{{Key: "$set", Value: bson.D{{Key: "type", Value: pokemon.Types}}}},
	)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (m *MongoRepo) GetAllPokemon(ctx context.Context) ([]*domain.Pokemon, error) {
	var results []*domain.Pokemon
	cursor, err := m.Coll.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		result := domain.Pokemon{}
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}

		results = append(results, &result)
	}
	return results, nil

}
