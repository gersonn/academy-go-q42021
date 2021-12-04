package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"gobootcamp/common"
	"gobootcamp/repositories"
)

type PokemonController struct {
	PokemonRepo *repositories.PokemonRepository
}

func (p *PokemonController) ReadCsv(c *gin.Context) {

	fileHeader, _ := c.FormFile("file")
	file, _ := fileHeader.Open()

	pokemons, err := common.CsvToPokemon(file)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "csv not well formated"})
	}

	p.PokemonRepo.SaveManyPokemons(pokemons)

	c.JSON(http.StatusCreated, pokemons)
}

func (p *PokemonController) GetPokemonById(c *gin.Context) {
	// question: is there a simpliest way to parse the param?
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	pokemon, err := p.PokemonRepo.GetPokemonById(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "pokemon not found"})
	}

	c.JSON(http.StatusOK, pokemon)
}

func (p *PokemonController) GetPokemonsFromPokeApi(c *gin.Context) {

	resp, err := p.PokemonRepo.GetPokemonsFromPokeAPI()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "something went wrong"})
	}

	c.JSON(http.StatusOK, resp)
}

func (p *PokemonController) GetPokemonsWithWorkerPool(c *gin.Context) {
	itemsQuery := c.Query("items")
	items, _ := strconv.Atoi(itemsQuery)

	itemsPerWorkerQuery := c.Query("items_per_workers")
	itemsPerWorker, _ := strconv.Atoi(itemsPerWorkerQuery)

	typePokemon := c.Query("type")

	fileHeader, _ := c.FormFile("file")
	file, _ := fileHeader.Open()

	pokemons, err := common.WorkerPoolReadCSV(file, items, itemsPerWorker, typePokemon)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "csv not well formated"})
	}

	c.JSON(http.StatusOK, pokemons)
}
