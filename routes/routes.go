package routes

import (
	"fmt"
	"log"
	"net/http"

	"gobootcamp/controllers"
	//csv "gobootcamp/common" //  Question: why is csv required? otherwise it doesn't work

	"github.com/gorilla/mux"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: mainHandler")
}

func HandleRequests() {
	r := mux.NewRouter()

	r.HandleFunc("/", mainHandler)
	r.HandleFunc("/csv", controllers.ReadCsv).Methods("POST")
	r.HandleFunc("/pokemon", controllers.GetOnePokemon).Methods("GET")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
