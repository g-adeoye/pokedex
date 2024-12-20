package main

import (
	"flag"
	"log"
	"net/http"
	"pokedex/lib/adapters/handlers"
	mongoClient "pokedex/lib/adapters/repository/mongo"
	redisClient "pokedex/lib/adapters/repository/redis"
	"pokedex/lib/core/services"

	"github.com/gin-gonic/gin"
)

var (
	uri         = "mongodb://localhost:27017"
	db_name     = "test"
	coll        = "streets"
	redisHost   = "localhost:6379"
	repo        = flag.String("db", "mongo", "Database for stroing message")
	httpHandler *handlers.HTTPHandler
	svc         *services.Service
)

func main() {
	flag.Parse()

	switch *repo {
	case "mongo":
		store, err := mongoClient.NewMongoClient(uri, db_name, coll)
		if err != nil {
			log.Fatalf("%v", err)
		}
		defer store.CloseDB()

		svc = services.NewPokemonService(store)
	case "redis":
		store := redisClient.NewRedisClient(redisHost)
		svc = services.NewPokemonService(store)

	}

	InitRoutes()

}

func helloFunc(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK, "message: pong",
	)
}

func InitRoutes() {
	router := gin.Default()
	httpHandler = handlers.NewHTTPHandler(*svc)
	router.GET("/", helloFunc)
	router.GET("/pokemon/:name", httpHandler.GetPokemon)
	router.GET("/pokemon", httpHandler.GetAllPokemon)

	router.Run(":3000")
}
