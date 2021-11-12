package common

import (
	"encoding/csv"
	"fmt"
	types "gobootcamp/models/person"
	"mime/multipart"
	"strconv"
)

func CsvToPersons(f multipart.File) ([]types.Person, error) {
	var persons []types.Person

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return []types.Person{}, err
	}

	for _, item := range lines {
		id, err := strconv.Atoi(item[0])

		if err != nil {
			fmt.Println(err)
		}

		person := types.Person{
			Id:   id,
			Name: item[1],
		}

		persons = append(persons, person)
	}

	fmt.Println(lines)
	return persons, nil
}
