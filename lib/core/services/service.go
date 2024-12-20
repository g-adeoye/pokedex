package services

import "pokedex/lib/core/ports"

type Service struct {
	repo ports.PokePort
}

func NewPokemonService(store ports.PokePort) *Service {
	return &Service{
		repo: store,
	}
}

