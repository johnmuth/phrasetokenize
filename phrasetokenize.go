package main

import (
	"strings"
	"fmt"
	"text/scanner"
)

func Phrasetokenize(input string) (tokens []string) {
	tokens = []string{}
	var s scanner.Scanner
	s.Init(strings.NewReader(input))
	var tok rune
	for tok != scanner.EOF {
		tok = s.Scan()
		fmt.Println("At position", s.Pos(), ":", s.TokenText())
		if len(s.TokenText()) > 0 {
			tokens = append(tokens, s.TokenText())
		}
	}
	return
}
