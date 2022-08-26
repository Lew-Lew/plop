package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Lew-Lew/plop/backend/letters"
)

func lettersOfTheDay(w http.ResponseWriter, r *http.Request) {
	log.Println("lettersOfTheDay")
	l := letters.GenerateLetters2()
	json.NewEncoder(w).Encode(l)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/word", lettersOfTheDay).Methods("GET")
	log.Println("Server listening on :8080")
	err := http.ListenAndServe(":8080", router)
	log.Fatal(err)
}
