package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRun(t *testing.T) {
	unitTestData := []struct {
		Description string
		Input string
		Expect bool
	}{
		{
			Description: "happy",
			Input: "(){}[]",
			Expect: true,
		},
	}

	for _, testData := range unitTestData {
		t.Run(testData.Description, func(t *testing.T) {
			got := Run(testData.Input)
			assert.Equal(t, testData.Expect, got)
		})
	}
}

var result bool
func BenchmarkValidBraces(b *testing.B) {
	var r bool
	for n := 0; n < b.N; n++ {
		r = Run("[{([](){})}{}()[]]")
	}
	result = r
}