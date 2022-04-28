package main

import (
	"codewars/golang/pkg"
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

var result []int

// benchmarkCountNSum can be run via `$ go test -bench . ./...`
func benchmarkCountNSum(i int, b *testing.B) {
	var r []int
	randArr := pkg.RandomIntArray(i)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = CountNSum(randArr)
	}
	result = r
}

func BenchmarkCountNSum1(b *testing.B)    { benchmarkCountNSum(1, b) }
func BenchmarkCountNSum2(b *testing.B)    { benchmarkCountNSum(2, b) }
func BenchmarkCountNSum3(b *testing.B)    { benchmarkCountNSum(3, b) }
func BenchmarkCountNSum10(b *testing.B)   { benchmarkCountNSum(10, b) }
func BenchmarkCountNSum20(b *testing.B)   { benchmarkCountNSum(20, b) }
func BenchmarkCountNSum40(b *testing.B)   { benchmarkCountNSum(40, b) }
func BenchmarkCountNSum100(b *testing.B)  { benchmarkCountNSum(100, b) }
func BenchmarkCountNSum1000(b *testing.B) { benchmarkCountNSum(1000, b) }
