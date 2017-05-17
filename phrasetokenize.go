package main

import (
	"bufio"
	"strings"
	"unicode"
)

func phrasetokenize(input string) (terms []string, phrases []string) {
	bs := bufio.NewScanner(strings.NewReader(input))
	bs.Split(bufio.ScanWords)
	var currentPhrase []string
	for bs.Scan() {
		if bs.Text()[:1] == `"` {
			if bs.Text()[len(bs.Text())-1:] == `"` {
				terms = append(terms, removePunctuation(bs.Text()))
			} else {
				currentPhrase = append(currentPhrase, removePunctuation(bs.Text()))
			}
		} else if bs.Text()[len(bs.Text())-1:] == `"` && len(currentPhrase) > 0 {
			currentPhrase = append(currentPhrase, removePunctuation(bs.Text()))
			phrases = append(phrases, strings.Join(currentPhrase, " "))
			currentPhrase = []string{}
		} else if len(currentPhrase) > 0 {
			currentPhrase = append(currentPhrase, removePunctuation(bs.Text()))
		} else {
			terms = append(terms, removePunctuation(bs.Text()))
		}
	}
	if len(currentPhrase) > 0 {
		for _, term := range currentPhrase {
			terms = append(terms, term)
		}
	}
	return
}

func removePunctuation(input string) (output string) {
	var listOfPunct []string
	for c := range RunesFromRange(unicode.Punct) {
		// the Latin runes category < 0x600
		if c < 0x600 {
			listOfPunct = append(listOfPunct, string(c))
		}
	}
	for _, punct := range listOfPunct {
		if strings.Contains(input, punct) {
			input = strings.Replace(input, punct, "", -1)
		}
	}
	return input
}

func RunesFromRange(tab *unicode.RangeTable) <-chan rune {
	res := make(chan rune)
	go func() {
		defer close(res)
		for _, r16 := range tab.R16 {
			for c := r16.Lo; c <= r16.Hi; c += r16.Stride {
				res <- rune(c)
			}
		}
		for _, r32 := range tab.R32 {
			for c := r32.Lo; c <= r32.Hi; c += r32.Stride {
				res <- rune(c)
			}
		}
	}()
	return res
}
