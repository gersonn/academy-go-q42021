package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gobootcamp/common"
)

func ReadCsv(w http.ResponseWriter, r *http.Request) {
	// ParseMultipartForm parses a request body as multipart/form-data
	r.ParseMultipartForm(32 << 20)

	file, handler, err := r.FormFile("file") // Retrieve the file from form data

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	defer file.Close()

	persons, err := common.CsvToPokemon(file)

	if err != nil {
		common.HandleInternalServerError(w)
	}

	w.WriteHeader(http.StatusOK)
	jsonResp, err := json.Marshal(persons)

	if err != nil {
		common.HandleInternalServerError(w)
	}

	w.Write(jsonResp)
	//fmt.Fprintf(w, "csv", persons, handler.Filename)
	fmt.Println("Endpoint Hit: read csv!", handler.Filename) // Question: How to ignore an unused var?

}
