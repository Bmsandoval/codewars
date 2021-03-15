package main

import (
	"codewars/golang/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleMorse(t *testing.T) {
	unitTestData := []struct {
		Description string
		Input string
		Expect string
	}{
		{
			Description: "single letter",
			Input: ".-",
			Expect: "A",
		},
		{
			Description: "single word",
			Input: ".- -... -.-.",
			Expect: "ABC",
		},
		{
			Description: "multiple words",
			Input: ".- -... -.-.   -.. . ..-.",
			Expect: "ABC DEF",
		},
	}

	for _, testData := range unitTestData {
		t.Run(testData.Description, func(t *testing.T) {
			got := DecodeMorse(testData.Input)
			assert.Equal(t, testData.Expect, got)
		})
	}
}

var result string
func benchmarkSimpleMorse(i int, b *testing.B) {
	var r string
	randString := pkg.RandomString(i)
	randMorse := ""
	for i,char := range randString {
		randMorse += pkg.MORSE_ENCODER[string(char)]
		if i != len(randString) {
			if pkg.RandomBool(6) { // make one in 6 characters a space
				randMorse += " "
			}
		}
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = DecodeMorse(randString)
	}
	result = r
}

func BenchmarkSimpleMorse1(b *testing.B)  { benchmarkSimpleMorse(1, b) }
func BenchmarkSimpleMorse2(b *testing.B)  { benchmarkSimpleMorse(2, b) }
func BenchmarkSimpleMorse3(b *testing.B)  { benchmarkSimpleMorse(3, b) }
func BenchmarkSimpleMorse10(b *testing.B)  { benchmarkSimpleMorse(10, b) }
func BenchmarkSimpleMorse20(b *testing.B)  { benchmarkSimpleMorse(20, b) }
func BenchmarkSimpleMorse40(b *testing.B)  { benchmarkSimpleMorse(40, b) }
