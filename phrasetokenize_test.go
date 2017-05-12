package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	input          string
	expectedOutput []string
	message        string
}{
	{
		"",
		[]string{},
		"empty string",
	},
	{
		"hello",
		[]string{"hello"},
		"single one-word string",
	},
	{
		"hello world",
		[]string{"hello","world"},
		"two one-word strings",
	},
	{
		`hello "end" world`,
		[]string{"hello", `"end"`, "world"},
		"three one-word strings, one in quotes",
	},
	{
		`"hello world"`,
		[]string{`"hello world"`},
		"single two-word phrase",
	},
	{
		`"hello world" they say`,
		[]string{`"hello world"`, "they", "say"},
		"two-word phrase and two single words",
	},
}

func TestPhrasetokenize(t *testing.T) {
	for _, tc := range testCases {
		assert.Equal(t, tc.expectedOutput, Phrasetokenize(tc.input), tc.message)
	}
}
