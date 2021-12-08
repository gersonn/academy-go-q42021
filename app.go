package main

import (
	"gobootcamp/controllers"
	"gobootcamp/repositories"
	"gobootcamp/routes"
)

func main() {
	pokemonController := controllers.PokemonController{PokemonRepo: &repositories.PokemonRepository{}}
	r := routes.HandleRequests(pokemonController)
	r.Run()
}
