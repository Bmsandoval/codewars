package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleAssembler(t *testing.T) {
	unitTestData := []struct {
		Description string
		Input []string
		Expect map[string]int
	}{
		{
			Description: "no movement",
			Input: []string{"mov a 5", "inc a", "dec a", "dec a", "jnz a -1", "inc a"},
			Expect: map[string]int{"a": 1},
		},
	}

	for _, testData := range unitTestData {
		t.Run(testData.Description, func(t *testing.T) {
			got := SimpleAssembler(testData.Input)
			assert.Equal(t, testData.Expect, got)
		})
	}
}

var result map[string]int
func benchmarkSimpleAssembler(i int, b *testing.B) {
	var r map[string]int
	for n := 0; n < b.N; n++ {
		r = SimpleAssembler([]string{"mov a 5", "inc a", "dec a", "dec a", "jnz a -1", "inc a"})
	}
	result = r
}