package generator

import (
	"testing"
	"unicode"

	"github.com/stretchr/testify/require"
)

func TestGenerateLetters(t *testing.T) {
	runGenerateFirst := generateLetters()
	runGenerateSecond := generateLetters()

	require.NotEqual(t, runGenerateFirst, runGenerateSecond)
}

func TestGenerateLettersLength(t *testing.T) {
	g := generateLetters()
	require.Equal(t, 10, len(g))

	for _, v := range g {
		require.True(t, unicode.IsLetter(rune(v[0])))
	}
}

func BenchmarkGenerateLetters2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateLetters()
	}
}
