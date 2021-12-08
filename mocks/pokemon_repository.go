package mocks

import (
	"gobootcamp/models"

	"github.com/stretchr/testify/mock"
)

type MockedPokemonRepository struct {
	mock.Mock
	pokemons []models.Pokemon
}

func (p *MockedPokemonRepository) SaveManyPokemons(pokemons []models.Pokemon) {
	p.Called()
}

func (p *MockedPokemonRepository) GetPokemonById(id int) (models.Pokemon, error) {
	ret := p.Called(id)
	var r0 models.Pokemon
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(models.Pokemon)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

func (p *MockedPokemonRepository) GetPokemonsFromPokeAPI() (models.Pokemons, error) {
	p.Called()
	return p.pokemons, nil
}
