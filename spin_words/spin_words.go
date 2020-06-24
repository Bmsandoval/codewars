package main

import (
	"log"
	"strings"
)

func main() {
	sample := "another random testypoo"
	log.Println(SpinWords(sample))
}

func SpinWords(sample string) string {
	fields := strings.Fields(sample)
	for i, field := range fields {
		if len(field) >= 5 {
			fields[i] = ReverseField(field)
		}
	}
	return strings.Join(fields, " ")
}

func ReverseField(field string) string {
	runeField := []rune(field)
	for i:=0; i < len(field)/2; i++ {
		foreIndex := i
		aftIndex := len(field)-1 - i
		runeField[foreIndex], runeField[aftIndex] = runeField[aftIndex], runeField[foreIndex]
	}
	return string(runeField)
}
