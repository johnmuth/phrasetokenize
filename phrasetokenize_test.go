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
		`"Hello world," she said`,
		[]string{`"Hello world,"`, "she", "said"},
		"a phrase and two words",
	},
}

func TestPhrasetokenize(t *testing.T) {
	for _, tc := range testCases {
		assert.Equal(t, tc.expectedOutput, Phrasetokenize(tc.input), tc.message)
	}
}
