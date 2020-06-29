package main

import (
	"codewars/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDigitalRoot(t *testing.T) {
	unitTestData := []struct {
		Description string
		Input string
		Expect string
	}{
		{
			Description: "happypath",
			Input: "happypath",
			Expect: "htapyppah",
		},
		{
			Description: "happy path",
			Input: "happy path",
			Expect: "yppah path",
		},
		{
			Description: "another random testypoo",
			Input: "another random testypoo",
			Expect: "rehtona modnar oopytset",
		},
	}

	for _, testData := range unitTestData {
		t.Run(testData.Description, func(t *testing.T) {
			got := SpinWords(testData.Input)
			assert.Equal(t, testData.Expect, got)
		})
	}
}

var result string
func benchmarkDigitalRoot(i int, b *testing.B) {
	var r string
	randString := pkg.RandomString(i)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = SpinWords(randString)
	}
	result = r
}