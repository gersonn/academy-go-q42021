package repositories

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"gobootcamp/models"
)

type PokemonRepository struct {
	pokemons []models.Pokemon
}

func (p *PokemonRepository) SaveManyPokemons(pokemons []models.Pokemon) {
	p.pokemons = pokemons
}

func (p *PokemonRepository) GetPokemonById(id int) (models.Pokemon, error) {
	for _, val := range p.pokemons {
		if val.Id == id {
			return val, nil
		}

	}

	// Question: is this right? passing an empty model?
	return models.Pokemon{}, errors.New("pokemon not found")

}

func (p *PokemonRepository) GetPokemonsFromPokeAPI() (models.Pokemons, error) {
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon?limit=10")

	if err != nil {
		return models.Pokemons{}, errors.New("pokemon not found")
	}

	defer resp.Body.Close()
	var pokemonsResp models.PokemonsResponse

	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &pokemonsResp)

	return pokemonsResp.Results, nil
}
