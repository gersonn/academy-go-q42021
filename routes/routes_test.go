package routes

import (
	"encoding/json"
	"gobootcamp/controllers"
	"gobootcamp/mocks"
	"gobootcamp/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type HealthResponse struct {
	Message string `json:"message"`
}

func TestHealthEndpoint(t *testing.T) {
	mockedRepo := new(mocks.MockedPokemonRepository)
	pokemonController := controllers.PokemonController{PokemonRepo: mockedRepo}
	router := HandleRequests(pokemonController)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/_health", nil)
	router.ServeHTTP(w, req)

	var response HealthResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "app is working!", response.Message)
}

func TestGetPokemonById(t *testing.T) {
	mockedRepo := new(mocks.MockedPokemonRepository)

	mockPokemonResp := models.Pokemon{
		Id:   1,
		Name: "Pikachu",
	}

	mockedRepo.On("GetPokemonById", mockPokemonResp.Id).Return(mockPokemonResp, nil)

	pokemonController := controllers.PokemonController{PokemonRepo: mockedRepo}
	router := HandleRequests(pokemonController)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/pokemon/1", nil)
	router.ServeHTTP(w, req)

	var response models.Pokemon
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, response, mockPokemonResp)
}

func TestGetPokemonByIdWithWrongParam(t *testing.T) {
	mockedRepo := new(mocks.MockedPokemonRepository)
	pokemonController := controllers.PokemonController{PokemonRepo: mockedRepo}
	router := HandleRequests(pokemonController)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/pokemon/1a", nil)
	router.ServeHTTP(w, req)

	var response models.Pokemon
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, 400, w.Code)
}
