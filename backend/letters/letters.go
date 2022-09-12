package letters // signifier à quel dossier le fichier appartient

import (
	"math/rand"
)

var voyels = []string{"a", "e", "i", "o", "u", "y"}
var consonants = []string{"z", "r", "t", "p", "m", "l", "k", "j", "h", "g", "f", "d", "s", "q", "w", "x", "c", "v", "b", "n"}

func GenerateLetters() [10]string {
	var arrayLetters [10]string

	for i := 0; i < 5; i++ {
		arrayLetters[i] = voyels[rand.Intn(len(voyels))]
	}

	for i := 5; i < 10; i++ {
		arrayLetters[i] = consonants[rand.Intn(len(consonants))]
	}
	return arrayLetters
}
