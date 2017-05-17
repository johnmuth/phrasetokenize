package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testCases = []struct {
	input           string
	expectedTerms   []string
	expectedPhrases []string
	message         string
}{
	{
		"",
		nil,
		nil,
		"empty string",
	},
	{
		"hello",
		[]string{"hello"},
		nil,
		"single one-word string",
	},
	{
		"hello world",
		[]string{"hello", "world"},
		nil,
		"two one-word strings",
	},
	{
		`"Hello world," she said`,
		[]string{"she", "said"},
		[]string{"Hello world"},
		"a phrase and two words",
	},
	{
		`"Hello," she said`,
		[]string{"Hello", "she", "said"},
		nil,
		"single word in quotes is a term",
	},
	{
		`she said "Hello how`,
		[]string{"she", "said", "Hello", "how"},
		nil,
		"unclosed quotes, is a term",
	},
}

func TestPhrasetokenize(t *testing.T) {
	for _, tc := range testCases {
		terms, phrases := phrasetokenize(tc.input)
		assert.Equal(t, tc.expectedTerms, terms, tc.message)
		assert.Equal(t, tc.expectedPhrases, phrases, tc.message)
	}
}
