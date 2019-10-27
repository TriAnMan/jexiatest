package translit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNextChar(t *testing.T) {
	a := assert.New(t)

	var resChar []int
	var resOffset int

	a.Panics(func() {
		nextChar([]byte{}, 0)
	}, "empty slice")

	a.Panics(func() {
		nextChar([]byte{}, 1)
	}, "offset for an empty slice")

	a.Panics(func() {
		nextChar([]byte(" "), 2)
	}, "too big offset for non empty slice")

	for i := range missed {
		resChar, resOffset = nextChar(missed, i)
		a.Nil(resChar, "non translatable character")
		a.Equal(i, resOffset, "non translatable character")
	}

	for _, trans := range transitions {
		resChar, resOffset = nextChar(trans.from, 0)
		a.Equal(trans.to, resChar, "translatable character")
		a.Equal(len(trans.from), resOffset, "translatable character")
	}

	resChar, resOffset = nextChar([]byte("angh"), 0)
	a.Equal([]int{0xF8D0}, resChar, "common case")
	a.Equal(1, resOffset, "common case")

	resChar, resOffset = nextChar([]byte("angh"), 1)
	a.Equal([]int{0xF8DB, 0xF8DC}, resChar, "corner case for n,gh")
	a.Equal(4, resOffset, "corner case for n,gh")

	resChar, resOffset = nextChar([]byte("qQ"), 0)
	a.Equal([]int{0xF8DF}, resChar, "corner case q")
	a.Equal(1, resOffset, "corner case q")

	resChar, resOffset = nextChar([]byte("qQ"), 1)
	a.Equal([]int{0xF8E0}, resChar, "corner case Q")
	a.Equal(2, resOffset, "corner case Q")

	resChar, resOffset = nextChar([]byte("A"), 0)
	a.Equal([]int{0xF8D0}, resChar, "uppercase")
	a.Equal(1, resOffset, "uppercase")

	for char := byte('a'); char < byte('z'); char++ {
		if char == byte('f') || char == byte('k') || char == byte('x') || char == byte('z') {
			continue
		}
		resChar, _ = nextChar([]byte{char}, 0)
		a.NotNil(resChar, "possible chars")
	}
}
