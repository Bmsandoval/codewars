package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountNSum(t *testing.T) {
	unitTestData := []struct {
		Description string
		Input       []int
		Expect      []int
	}{
		{
			Description: "happy path",
			Input:       []int{10, -10},
			Expect:      []int{1, -10},
		},
		{
			Description: "order shouldn't matter",
			Input:       []int{-10, 10},
			Expect:      []int{1, -10},
		},
		{
			Description: "can handle empty array",
			Input:       []int{},
			Expect:      []int{},
		},
		{
			Description: "can handle missing negatives",
			Input:       []int{10},
			Expect:      []int{1, 0},
		},
		{
			Description: "can handle missing positives",
			Input:       []int{-10},
			Expect:      []int{0, -10},
		},
		{
			Description: "longer example",
			Input:       []int{20, -2, 47, -11, 33, -17, 100, -1},
			Expect:      []int{4, -31},
		},
	}

	for _, testData := range unitTestData {
		t.Run(testData.Description, func(t *testing.T) {
			got := CountNSum(testData.Input)
			assert.Equal(t, testData.Expect, got)
		})
	}
}
