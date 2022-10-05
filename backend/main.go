package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/Lew-Lew/plop/backend/bst"
	"github.com/Lew-Lew/plop/backend/generator"
)

var nn = bst.NewNode()

func lettersOfTheDay(w http.ResponseWriter, _ *http.Request) {
	log.Println("lettersOfTheDay")
	l, err := generator.LettersOfTheDay()
	if err != nil {
		http.Error(w, "WORD_UNINITIALIZED", http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(l); err != nil {
		http.Error(w, "INTERNAL_ERROR", http.StatusInternalServerError)
		return
	}
}

// verifyWordRequest contains all data of the JSON request body
type verifyWordRequest struct {
	Word string `json:"word"`
}

// verifyWordResponse contains all data of the JSON response
type verifyWordResponse struct {
	Exist bool `json:"exist"`
}

func verifyWord(w http.ResponseWriter, r *http.Request) {
	var reqBody verifyWordRequest

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "INTERNAL_ERROR", http.StatusInternalServerError)
		return
	}

	resBody := verifyWordResponse{
		Exist: nn.WordExist(reqBody.Word),
	}

	if err := json.NewEncoder(w).Encode(resBody); err != nil {
		http.Error(w, "INTERNAL_ERROR", http.StatusInternalServerError)
		return
	}
}

func main() {
	file, err := os.Open("words.txt")
	if err != nil {
		log.Fatal("open words.txt:", err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nn.InsertWord(scanner.Text())
	}
	file.Close()

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// create REST API server
	addr := ":8081"
	router := mux.NewRouter()
	router.Use(middlewareMain)
	router.HandleFunc("/word", lettersOfTheDay).Methods("GET")
	router.HandleFunc("/word", verifyWord).Methods("POST")

	log.Println("Server listening on", addr)
	err = http.ListenAndServe(addr, router)
	log.Fatal(err)
}

func middlewareMain(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("%v %v\n", req.Method, req.URL)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		next.ServeHTTP(w, req)
	})
}
