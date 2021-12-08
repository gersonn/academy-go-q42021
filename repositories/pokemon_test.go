package repositories

import (
	"gobootcamp/models"
	"testing"
)

func TestGetPokemonById(t *testing.T) {
	r := PokemonRepository{}
	pokemons := models.Pokemons{
		models.Pokemon{Id: 1, Name: "Pikachu"},
		models.Pokemon{Id: 2, Name: "Charmander"},
		models.Pokemon{Id: 3, Name: "Bulbasaur"},
	}

	r.SaveManyPokemons(pokemons)
	result, _ := r.GetPokemonById(1)

	if result.Name != "Pikachu" {
		t.Error("GetPokemonById not getting the right pokemon")
	}
}
