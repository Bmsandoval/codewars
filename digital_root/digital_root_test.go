package main

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestDigitalRoot(t *testing.T) {
	unitTestData := []struct {
		Description string
		Input int
		Expect int
	}{
		{
			Description: "happy path",
			Input: 123,
			Expect: 6,
		},
		{
			Description: "zero",
			Input: 0,
			Expect: 0,
		},
		{
			Description: "int max",
			Input: math.MaxInt32,
			Expect: 1,
		},
	}

	for _, testData := range unitTestData {
		t.Run(testData.Description, func(t *testing.T) {
			got := DigitalRoot(testData.Input)
			assert.Equal(t, testData.Expect, got)
		})
	}
}

var result int
func benchmarkDigitalRoot(i int, b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = DigitalRoot(i)
	}
	result = r
}