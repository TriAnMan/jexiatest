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

func TestString(t *testing.T) {
	a := assert.New(t)

	var result []int
	var err error
	result, err = String("")
	a.Nil(result, "empty name")
	a.EqualError(err, "invalid Klingon name", "empty name")

	for _, char := range missed {
		result, err = String(string(char))
		a.Nil(result, "missed characters")
		a.EqualError(err, "invalid Klingon name", "missed characters")
	}

	for char := byte(0); char < 255; char++ {
		if char >= byte('A') && char <= byte('Z') {
			continue
		}
		if char >= byte('a') && char <= byte('z') {
			continue
		}
		if char == byte(' ') || char == byte('\'') {
			continue
		}
		result, err = String(string(char))
		a.Nil(result, "invalid characters")
		a.EqualError(err, "invalid Klingon name", "invalid characters")
	}

	for _, trans := range transitions {
		result, err = String(string(trans.from))
		a.Equal(trans.to, result, "valid characters")
		a.Nil(err, "valid characters")
	}

	result, err = String("James Tiberius Kirk")
	a.Nil(result, "name with invalid characters")
	a.EqualError(err, "invalid Klingon name", "invalid characters")

	result, err = String("Uhura")
	a.Equal([]int{0xF8E5, 0xF8D6, 0xF8E5, 0xF8E1, 0xF8D0}, result, "valid name")
	a.Nil(err, "valid name")

	result, err = String("Leonard McCoy")
	a.Equal(
		[]int{0xF8D9, 0xF8D4, 0xF8DD, 0xF8DB, 0xF8D0, 0xF8E1, 0xF8D3, 0x0020, 0xF8DA, 0xF8D2, 0xF8D2, 0xF8DD, 0xF8E8},
		result,
		"valid name",
	)
	a.Nil(err, "valid name")
}
