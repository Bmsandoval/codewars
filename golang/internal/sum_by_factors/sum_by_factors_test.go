package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumOfDivided(t *testing.T) {
	unitTestData := []struct {
		Description string
		Input []int
		Expect string
	}{
		{
			Description: "Input:513",
			Input: []int{12,15},
			Expect: "(2 12)(3 27)(5 15)",
		},
	}

	for _, testData := range unitTestData {
		t.Run(testData.Description, func(t *testing.T) {
			got := SumOfDivided(testData.Input)
			assert.Equal(t, testData.Expect, got)
		})
	}
}

var result string
func BenchmarkSumOfDivided(b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		r = SumOfDivided([]int{12,15})
	}
	result = r
}
