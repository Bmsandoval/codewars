package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidBraces(t *testing.T) {
	unitTestData := []struct {
		Description string
		Input string
		Expect bool
	}{
		{
			Description: "immediate match",
			Input: "(){}[]",
			Expect: true,
		},
		{
			Description: "wrapping match",
			Input: "{([])}",
			Expect: true,
		},
		{
			Description: "unmatched",
			Input: "(}",
			Expect: false,
		},
		{
			Description: "mismatched",
			Input: "[(])",
			Expect: false,
		},
		{
			Description: "extras",
			Input: "[({})](]",
			Expect: false,
		},
	}

	for _, testData := range unitTestData {
		t.Run(testData.Description, func(t *testing.T) {
			got := ValidBraces(testData.Input)
			assert.Equal(t, testData.Expect, got)
		})
	}
}

var result bool
func benchmarkValidBraces(i int, b *testing.B) {
	var r bool
	for n := 0; n < b.N; n++ {
		r = ValidBraces("[{([](){})}{}()[]]")
	}
	result = r
}