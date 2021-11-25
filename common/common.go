package common

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"

	"gobootcamp/models"
)

func CsvToPokemon(f multipart.File) ([]models.Pokemon, error) {
	var pokemons []models.Pokemon

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return []models.Pokemon{}, err
	}

	for _, item := range lines {
		id, err := strconv.Atoi(item[0])

		if err != nil {
			fmt.Println(err)
			continue
		}

		person := models.Pokemon{
			Id:   id,
			Name: item[1],
		}

		pokemons = append(pokemons, person)
	}

	fmt.Println(lines)
	return pokemons, nil
}

func HandleInternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Some Error Occurred"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened. Err: %s", err)
	}
	w.Write(jsonResp)
}
