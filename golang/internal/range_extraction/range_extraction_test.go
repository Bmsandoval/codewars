package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRangeExtraction(t *testing.T) {
	unitTestData := []struct {
		Description string
		Input []int
		Expect string
	}{
		{
			Description: "single digit",
			Input: []int{-1},
			Expect: "-1",
		},
		{
			Description: "subsequent but no range",
			Input: []int{1,2},
			Expect: "1,2",
		},
		{
			Description: "simple range",
			Input: []int{1,2,3},
			Expect: "1-3",
		},
		{
			Description: "simple range of negatives",
			Input: []int{-3, -2, -1},
			Expect: "-3--1",
		},
		{
			Description: "multiple ranges",
			Input: []int{1,2,3,5,6,7},
			Expect: "1-3,5-7",
		},
		{
			Description: "compound",
			Input: []int{1,2,3,5,7,8,9},
			Expect: "1-3,5,7-9",
		},
		{
			Description: "failed sample from codewars",
			Input: []int{-6,-3,-2,-1,0,1,3,4,5,7,8,9,10,11,14,15,17,18,19,20},
			Expect: "-6,-3-1,3-5,7-11,14,15,17-20",
		},
	}

	for _, testData := range unitTestData {
		t.Run(testData.Description, func(t *testing.T) {
			got := Solution(testData.Input)
			assert.Equal(t, testData.Expect, got)
		})
	}
}
