package letters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateLetters(t *testing.T) {
	runGenerateFirst := GenerateLetters()
	runGenerateSecond := GenerateLetters()

	assert.NotEqual(t, runGenerateFirst, runGenerateSecond)
}

func TestGenerateLetters2(t *testing.T) {
	assert.NotEqual(t, GenerateLetters2(), GenerateLetters2())
}

func BenchmarkGenerateLetters2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateLetters2()
	}
}

func BenchmarkGenerateLetters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateLetters()
	}
}
