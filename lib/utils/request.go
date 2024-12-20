package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func FetchPokemon(pokemon string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemon))
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
