package common

import (
	"encoding/csv"
	"fmt"
	"io"
	"mime/multipart"
	"strconv"
	"sync"

	"gobootcamp/models"
)

//	CsvToPokemon ~Receives a multipart csv file with format id,name and returns a pokemon list
func CsvToPokemon(f multipart.File) (models.Pokemons, error) {
	var pokemons models.Pokemons

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return models.Pokemons{}, err
	}

	for _, item := range lines {
		pokemon := parsePokemon(item)

		pokemons = append(pokemons, pokemon)
	}

	fmt.Println(pokemons)
	return pokemons, nil
}

func worker(t string, ipw int, jobs <-chan []string, results chan<- models.Pokemon) {
	counter := 0

	for j := range jobs {

		//	this validates a worker only reads the number of items_per_worker established
		counter++
		if counter > ipw {
			break
		}

		p := parsePokemon(j)
		results <- p

	}
}

//	WorkerPoolReadCSV ~Receives a multipart csv file with format id,name and returns a pokemon list
// f: csv file
//	items: number of items from the csv file to be returned
// itemsPerWorker: number of jobs each worker is going to execute
// t: type of items that will be returned, valid values are odd and even
func WorkerPoolReadCSV(f multipart.File, items int, itemsPerWorker int, t string) (models.Pokemons, error) {
	reader := csv.NewReader(f)
	var pokemons models.Pokemons

	numWorkers := 5
	jobs := make(chan []string, items)
	res := make(chan models.Pokemon, items)

	var wg sync.WaitGroup
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)

		go func() {
			worker(t, itemsPerWorker, jobs, res)
			defer wg.Done()
		}()
	}

	for j := 1; j <= items*2; j++ {
		rStr, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("ERROR: ", err.Error())
			return models.Pokemons{}, err
		}

		if t == "odd" && j%2 != 0 {
			continue
		} else if t == "even" && j%2 == 0 {
			continue
		}

		jobs <- rStr
	}

	close(jobs)
	wg.Wait()
	close(res)

	for r := range res {
		pokemons = append(pokemons, r)
	}

	return pokemons, nil
}

func parsePokemon(data []string) models.Pokemon {
	id, _ := strconv.Atoi(data[0])
	pokemon := models.Pokemon{
		Id:   id,
		Name: data[1],
	}

	return pokemon
}
