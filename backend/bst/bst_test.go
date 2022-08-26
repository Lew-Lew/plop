package bst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBst(t *testing.T) {
	nn := NewNode()
	w := nn.WordExist("banane")
	assert.Equal(t, w, false)

	nn.InsertWord("maison")
	w = nn.WordExist("maison")
	assert.Equal(t, w, true)

	w = nn.WordExist("mais")
	assert.Equal(t, w, false)
}
