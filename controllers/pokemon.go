package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetOnePokemon(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Hello"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened. Err: %s", err)
	}
	w.Write(jsonResp)
}
