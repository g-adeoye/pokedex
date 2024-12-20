package utils

import (
	"fmt"
	"testing"
)

func TestFetchPokemon(t *testing.T) {
	t.Run("test request", func(t *testing.T) {
		want, _ := FetchPokemon("ditto")

		fmt.Print(want)
	})
}
