package random

import (
	"testing"
)

func TestShuffle(t *testing.T) {
	b := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	Shuffle(b)
	lastElement := b[0]
	shuffled := false
	for _, element := range b[1:] {
		if element < lastElement {
			shuffled = true
		}
	}
	if !shuffled {
		t.Fatalf("shuffle did not shuffle slice out of order")
	}
}
