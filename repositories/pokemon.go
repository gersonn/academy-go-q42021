package repositories

import "gobootcamp/models"

var pokemons []models.Pokemon

func GetOnePokemon(id int) {

}

func SaveManyPokemon(data []models.Pokemon) {
	pokemons = data
	// question: why is there a warning in pokemons
	// is there something similar like this, tu reference the local variable
}
