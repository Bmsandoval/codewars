package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDirReduc(t *testing.T) {
	unitTestData := []struct {
		Description string
		Input []string
		Expect []string
	}{
		{
			Description: "no movement",
			Input: []string{"NORTH", "SOUTH", "EAST", "WEST"},
			Expect: []string{},
		},
		{
			Description: "no movement wrapped",
			Input: []string{"NORTH", "EAST", "WEST", "SOUTH"},
			Expect: []string{},
		},
		{
			Description: "wrapped movement",
			Input: []string{"NORTH", "NORTH", "EAST", "WEST", "SOUTH", "EAST"},
			Expect: []string{"NORTH", "EAST"},
		},
	}

	for _, testData := range unitTestData {
		t.Run(testData.Description, func(t *testing.T) {
			got := DirReduc(testData.Input)
			assert.Equal(t, testData.Expect, got)
		})
	}
}

var result []string
func BenchmarkDirReduc(b *testing.B) {
	var r []string
	for n := 0; n < b.N; n++ {
		r = DirReduc([]string{"NORTH", "EAST", "WEST", "SOUTH"})
	}
	result = r
}