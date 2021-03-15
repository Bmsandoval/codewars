package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateSpiral(t *testing.T) {
	unitTestData := []struct {
		Description string
		Input int
		Expect [][]int
	}{
		{
			Description: "input too small",
			Input: -1,
			Expect: [][]int{},
		},
		{
			Description: "1",
			Input: 1,
			Expect: [][]int{
				{1}},
		},
		{
			Description: "2",
			Input: 2,
			Expect: [][]int{
				{1,2},
				{4,3}},
		},
		{
			Description: "3",
			Input: 3,
			Expect: [][]int{
				{1,2,3},
				{8,9,4},
				{7,6,5}},
		},
		{
			Description: "4",
			Input: 4,
			Expect: [][]int{
				{ 1, 2, 3, 4},
				{12,13,14, 5},
				{11,16,15, 6},
				{10, 9, 8, 7}},
		},
		{
			Description: "5",
			Input: 5,
			Expect: [][]int{
				{ 1, 2, 3, 4, 5},
				{16,17,18,19, 6},
				{15,24,25,20, 7},
				{14,23,22,21, 8},
				{13,12,11,10, 9}},
		},
	}

	for _, testData := range unitTestData {
		t.Run(testData.Description, func(t *testing.T) {
			got := CreateSpiral(testData.Input)
			assert.Equal(t, testData.Expect, got)
		})
	}
}

var result [][]int
func benchmarkCreateSpiral(i int, b *testing.B) {
	var r [][]int
	for n := 0; n < b.N; n++ {
		r = CreateSpiral(i)
	}
	result = r
}

func BenchmarkCreateSpiral1(b *testing.B)  { benchmarkCreateSpiral(1, b) }
func BenchmarkCreateSpiral2(b *testing.B)  { benchmarkCreateSpiral(2, b) }
func BenchmarkCreateSpiral3(b *testing.B)  { benchmarkCreateSpiral(3, b) }
func BenchmarkCreateSpiral10(b *testing.B)  { benchmarkCreateSpiral(10, b) }
func BenchmarkCreateSpiral20(b *testing.B)  { benchmarkCreateSpiral(20, b) }
func BenchmarkCreateSpiral40(b *testing.B)  { benchmarkCreateSpiral(40, b) }
