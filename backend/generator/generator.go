package generator

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/go-co-op/gocron"
)

var word [10]string      // 10 letters of the day
var wordMux sync.RWMutex // mutex for word

func init() {
	rand.Seed(time.Now().UnixNano())

	s := gocron.NewScheduler(time.UTC)

	setWord()
	s.Every(1).Day().At("00:00").Do(func() {
		log.Println("New day: new letters have been generated")
		setWord()
	})

	s.StartAsync()
}

func setWord() {
	wordMux.Lock()
	defer wordMux.Unlock()
	word = generateLetters()
}

// LettersOfTheDay return the 10 letters of the current day
func LettersOfTheDay() ([10]string, error) {
	wordMux.RLock()
	defer wordMux.RUnlock()
	if len(word) == 0 {
		return word, fmt.Errorf("word is not initialized")
	}
	return word, nil
}

var voyels = []string{"a", "e", "i", "o", "u", "y"}
var consonants = []string{"z", "r", "t", "p", "m", "l", "k", "j", "h", "g", "f", "d", "s", "q", "w", "x", "c", "v", "b", "n"}

// generateLetters generate a list of 10 letters
func generateLetters() [10]string {
	var arrayLetters [10]string

	for i := 0; i < 5; i++ {
		arrayLetters[i] = voyels[rand.Intn(len(voyels))]
	}

	for i := 5; i < 10; i++ {
		arrayLetters[i] = consonants[rand.Intn(len(consonants))]
	}

	rand.Shuffle(
		len(arrayLetters),
		func(i, j int) {
			arrayLetters[i], arrayLetters[j] = arrayLetters[j], arrayLetters[i]
		},
	)

	return arrayLetters
}
