package main

import (
	"codewars/golang/pkg"
	"strings"
)

var MORSE_CODE = pkg.MORSE_DECODER

func DecodeMorse(morseCode string) string {
	result := ""
	for _,word := range Tokenize(morseCode) {
		if result != "" {
			result += " "
		}
		for _,char := range strings.Split(word," ") {
			result += MORSE_CODE[char]
		}
	}
	return result
}

func Tokenize(str string) []string {
	return strings.Split(str,"   ")
}