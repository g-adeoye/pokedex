package ports

import (
	"context"
	model "pokedex/lib/core/domain"
)

type PokePort interface {
	GetPokemon(ctx context.Context, pokemon string) (*model.Pokemon, error)
	GetAllPokemon(ctx context.Context) ([]*model.Pokemon, error)
}
