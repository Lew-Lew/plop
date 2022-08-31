package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/Lew-Lew/plop/backend/bst"
	"github.com/Lew-Lew/plop/backend/letters"
)

var nn = bst.NewNode()

func lettersOfTheDay(w http.ResponseWriter, r *http.Request) {
	log.Println("lettersOfTheDay")
	l := letters.GenerateLetters2()
	json.NewEncoder(w).Encode(l)
}

// la structure de données qui va contenir le JSON pour la requête
type verifyWordRequest struct {
	Word string `json:"word"`
}

// la structure de donnée pour la réponse qu'on renvoit
type verifyWordResponse struct {
	Exist bool `json:"exist"`
}

func verifyWord(w http.ResponseWriter, r *http.Request) {
	log.Println("verifyWord")
	// variable qui permet de stocker résultat décodé du JSON récupéré de la requête POST ie la props du user
	var rBody verifyWordRequest
	json.NewDecoder(r.Body).Decode(&rBody)

	response := verifyWordResponse{
		Exist: nn.WordExist(rBody.Word),
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	file, err := os.Open("words.txt")
	if err != nil {
		log.Fatal("open words.txt:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nn.InsertWord(scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// create REST API server
	router := mux.NewRouter()
	router.HandleFunc("/word", lettersOfTheDay).Methods("GET")
	router.HandleFunc("/word", verifyWord).Methods("POST")
	log.Println("Server listening on :8081")
	err = http.ListenAndServe(":8081", router)
	log.Fatal(err)
}
