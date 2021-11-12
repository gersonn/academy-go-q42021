package routes

import (
	"fmt"
	"log"
	"net/http"

	csv "gobootcamp/common"

	"github.com/gorilla/mux"
)

type Person struct {
	Id   int
	Name string
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: mainHandler")
}

func readCsv(w http.ResponseWriter, r *http.Request) {
	// ParseMultipartForm parses a request body as multipart/form-data
	r.ParseMultipartForm(32 << 20)

	file, handler, err := r.FormFile("file") // Retrieve the file from form data

	if err != nil {
		// todo: improve error handling
	}

	defer file.Close()

	persons, err := csv.CsvToPersons(file)

	if err != nil {
		// todo: improve error handling
	}

	fmt.Fprintf(w, "csv", persons, handler.Filename)
	fmt.Println("Endpoint Hit: read csv!")

}

func HandleRequests() {
	r := mux.NewRouter()

	r.HandleFunc("/", mainHandler)
	r.HandleFunc("/csv", readCsv).Methods("POST")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
