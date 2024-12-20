package repository

import (
	"context"
	"encoding/json"
	"pokedex/lib/core/domain"
)

func (r *RedisClient) GetPokemon(ctx context.Context, pokemon string) (*domain.Pokemon, error) {
	result, err := r.client.HGet(ctx, "pokemon", pokemon).Result()
	if err != nil {
		return &domain.Pokemon{}, nil
	}
	data := &domain.Pokemon{}

	err = json.Unmarshal([]byte(result), data)
	if err != nil {
		return &domain.Pokemon{}, err
	}
	return data, nil
}

func (r *RedisClient) GetAllPokemon(ctx context.Context) ([]*domain.Pokemon, error) {
	var results []*domain.Pokemon
	values, err := r.client.HGetAll(ctx, "pokemon").Result()
	if err != nil {
		return nil, err
	}

	for _, val := range values {
		pokemon := &domain.Pokemon{}
		err := json.Unmarshal([]byte(val), pokemon)
		if err != nil {
			return nil, err
		}
		results = append(results, pokemon)
	}

	return results, nil
}
