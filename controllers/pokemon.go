package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"gobootcamp/common"
	"gobootcamp/models"
)

type PokemonController struct {
	PokemonRepo models.PokemonRepository
}

func (p *PokemonController) ReadCsv(c *gin.Context) {

	fileHeader, _ := c.FormFile("file")
	file, _ := fileHeader.Open()

	pokemons, err := common.CsvToPokemon(file)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "csv not well formated"})
		return
	}

	p.PokemonRepo.SaveManyPokemons(pokemons)

	c.JSON(http.StatusCreated, pokemons)
}

func (p *PokemonController) GetPokemonById(c *gin.Context) {
	// question: is there a simpliest way to parse the param?
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a number"})
		return
	}

	pokemon, err := p.PokemonRepo.GetPokemonById(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "pokemon not found"})
		return
	}

	c.JSON(http.StatusOK, pokemon)
}

func (p *PokemonController) GetPokemonsFromPokeApi(c *gin.Context) {

	resp, err := p.PokemonRepo.GetPokemonsFromPokeAPI()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "something went wrong"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (p *PokemonController) GetPokemonsWithWorkerPool(c *gin.Context) {
	itemsQuery := c.Query("items")
	items, err := strconv.Atoi(itemsQuery)

	if err != nil || items == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "items must be a valid number and bigger than 0"})
		return
	}

	itemsPerWorkerQuery := c.Query("items_per_workers")
	itemsPerWorker, err := strconv.Atoi(itemsPerWorkerQuery)

	if err != nil || itemsPerWorker == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "items_per_workers must be a valid number and bigger than 0"})
		return
	}

	typePokemon := c.Query("type")
	if !isValidType(typePokemon) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "type valid values are: odd, even"})
		return
	}

	fileHeader, _ := c.FormFile("file")
	file, _ := fileHeader.Open()

	pokemons, err := common.WorkerPoolReadCSV(file, items, itemsPerWorker, typePokemon)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "csv not well formated"})
		return
	}

	c.JSON(http.StatusOK, pokemons)
}

func isValidType(t string) bool {
	switch t {
	case
		"odd",
		"even":
		return true
	}
	return false
}
