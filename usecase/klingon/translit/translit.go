// Package translit provides routines for transliteration of latin characters into Klingon ones.
// Refer to https://en.wikipedia.org/wiki/Klingon_scripts for additional info.

package translit

import (
	"bytes"
	"strings"
)

type transition struct {
	from []byte
	to   []int
}

var missed = []byte("fkxz")

var q = transition{[]byte("Q"), []int{0xF8E0}}

var transitions = []transition{
	{[]byte("ngh"), []int{0xF8DB, 0xF8DC}},
	{[]byte("tlh"), []int{0xF8E4}},
	{[]byte("ch"), []int{0xF8D2}},
	{[]byte("gh"), []int{0xF8D5}},
	{[]byte("ng"), []int{0xF8DC}},
	{[]byte("a"), []int{0xF8D0}},
	{[]byte("b"), []int{0xF8D1}},
	{[]byte("c"), []int{0xF8D2}},
	{[]byte("d"), []int{0xF8D3}},
	{[]byte("e"), []int{0xF8D4}},
	{[]byte("g"), []int{0xF8D5}},
	{[]byte("h"), []int{0xF8D6}},
	{[]byte("i"), []int{0xF8D7}},
	{[]byte("j"), []int{0xF8D8}},
	{[]byte("l"), []int{0xF8D9}},
	{[]byte("m"), []int{0xF8DA}},
	{[]byte("n"), []int{0xF8DB}},
	{[]byte("o"), []int{0xF8DD}},
	{[]byte("p"), []int{0xF8DE}},
	{[]byte("q"), []int{0xF8DF}},
	{[]byte("r"), []int{0xF8E1}},
	{[]byte("s"), []int{0xF8E2}},
	{[]byte("t"), []int{0xF8E3}},
	{[]byte("u"), []int{0xF8E5}},
	{[]byte("v"), []int{0xF8E6}},
	{[]byte("w"), []int{0xF8E7}},
	{[]byte("y"), []int{0xF8E8}},
	{[]byte("'"), []int{0xF8E9}},
	{[]byte(" "), []int{0x0020}},
}

// nextChar returns a corresponding Klingon substitution for a latin characters found in a specified offset.
// Returns nil if input can't be translated
func nextChar(latin []byte, offset int) (klingon []int, nextOffset int) {
	if latin[offset] == q.from[0] {
		return q.to, offset + len(q.from)
	}

	for _, trans := range transitions {
		if len(trans.from) > len(latin)-offset {
			continue
		}
		stripe := latin[offset : offset+len(trans.from)]
		stripe = []byte(strings.ToLower(string(stripe)))
		if 0 == bytes.Compare(trans.from, stripe) {
			return trans.to, offset + len(trans.from)
		}
	}

	return nil, offset
}
