package models

type PokemonRepository interface {
	SaveManyPokemons(pokemons []Pokemon)
	GetPokemonById(id int) (Pokemon, error)
	GetPokemonsFromPokeAPI() (Pokemons, error)
}
