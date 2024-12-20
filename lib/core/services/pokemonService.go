package services

import (
	"context"
	"pokedex/lib/core/domain"
)

func (s *Service) GetPokemon(ctx context.Context, pokemon string) (*domain.Pokemon, error) {
	result, err := s.repo.GetPokemon(ctx, pokemon)
	if err != nil {
		return &domain.Pokemon{}, err
	}

	return result, nil
}

func (s *Service) GetAllPokemon(ctx context.Context) ([]*domain.Pokemon, error) {
	result, err := s.repo.GetAllPokemon(ctx)

	if err != nil {
		return []*domain.Pokemon{}, err
	}

	return result, nil
}
