package handlers

import (
	"fmt"
	"log"
	"net/http"
	"pokedex/lib/core/services"
	"pokedex/lib/utils"

	"github.com/m7shapan/njson"

	"github.com/TwiN/go-color"
	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	svc services.Service
}

func NewHTTPHandler(service services.Service) *HTTPHandler {
	return &HTTPHandler{
		svc: service,
	}
}

func (h *HTTPHandler) GetPokemon(ctx *gin.Context) {
	name := ctx.Param("name")

	pokemon, err := h.svc.GetPokemon(ctx, name)

	if err != nil {
		log.Printf(color.Red + fmt.Sprintf(
			"\n%s does not exist locally\ncalling poké API...",
			pokemon) + color.Reset)
	}

	pokeData, err := utils.FetchPokemon(name)

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			"message: pokemon does not exist",
		)
	}

	err = njson.Unmarshal(pokeData, pokemon)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Err": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, pokemon)
}

func (h *HTTPHandler) GetAllPokemon(ctx *gin.Context) {

	all_pokemon, err := h.svc.GetAllPokemon(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Err": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, all_pokemon)

}
