package letters // signifier à quel dossier le fichier appartient

import (
	"math/rand"
)

var letters = []string{"a", "z", "e", "r", "t", "y", "u", "i", "o", "p", "m", "l", "k", "j", "h", "g", "f", "d", "s", "q", "w", "x", "c", "v", "b", "n"}

// nommage: si commence par une majuscule c'est public sinon privé
func GenerateLetters() [9]string {
	// déclaration de variable:
	// var letters []string
	// initialisation de variable:
	// letters = []string{"a","z","e","r","t","y","u","i","o","p","m","l","k","j","h","g","f","d","s","q","w","x","c","v","b","n"}
	//  =
	var arrayLetters [9]string

	for i := 0; i < 9; i++ {
		rand.Shuffle(len(letters), func(i, j int) { letters[i], letters[j] = letters[j], letters[i] })
		arrayLetters[i] = letters[0]
	}
	return arrayLetters
}

func GenerateLetters2() [9]string {
	var arrayLetters [9]string

	for i := 0; i < 9; i++ {
		arrayLetters[i] = letters[rand.Intn(len(letters))]
	}
	return arrayLetters
}
