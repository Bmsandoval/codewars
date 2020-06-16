package main

import (
	"codewars/shared"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDuplicateEncode(t *testing.T) {
	unitTestData := []struct {
		Description string
		Input string
		Expect string
	}{
		{
			Description: "happy path",
			Input: "happy path",
			Expect: "))))(())()",
		},
	}

	for _, testData := range unitTestData {
		t.Run(testData.Description, func(t *testing.T) {
			got := DuplicateEncode(testData.Input)
			assert.Equal(t, testData.Expect, got)
		})
	}
}

var result string
func benchmarkDuplicateEncode(i int, b *testing.B) {
	var r string
	randString := shared.RandomString(i)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = DuplicateEncode(randString)
	}
	result = r
}