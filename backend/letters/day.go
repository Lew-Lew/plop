package letters

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/go-co-op/gocron"
)

var word [10]string
var wordMux sync.RWMutex

func init() {
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
	word = GenerateLetters()
}

func WordOfTheDay() ([10]string, error) {
	if len(word) == 0 {
		return word, fmt.Errorf("word is not initialized")
	}
	return word, nil
}
