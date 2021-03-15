package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPrime(t *testing.T) {
	unitTestData := []struct {
		Description string
		Input int
		Expect bool
	}{
		{
			Description: "2",
			Input: 2,
			Expect: true,
		},
		{
			Description: "11",
			Input: 11,
			Expect: true,
		},
		{
			Description: "101",
			Input: 101,
			Expect: true,
		},
		{
			Description: "199",
			Input: 199,
			Expect: true,
		},
		{
			Description: "200",
			Input: 200,
			Expect: false,
		},
	}

	for _, testData := range unitTestData {
		t.Run(testData.Description, func(t *testing.T) {
			got := IsPrime(testData.Input)
			assert.Equal(t, testData.Expect, got)
		})
	}
}

var primeResult bool
func benchmarkIsPrime(i int, b *testing.B) {
	var r bool
	for n := 0; n < b.N; n++ {
		r = IsPrime(i)
	}
	primeResult = r
}

func BenchmarkIsPrime10pow3(b *testing.B)  { benchmarkIsPrime(1000, b) }
func BenchmarkIsPrime10pow6(b *testing.B)  { benchmarkIsPrime(1000000, b) }
func BenchmarkIsPrime10pow9(b *testing.B)  { benchmarkIsPrime(1000000000, b) }
func BenchmarkIsPrime10pow12(b *testing.B)  { benchmarkIsPrime(1000000000000, b) }
func BenchmarkIsPrime10pow15(b *testing.B)  { benchmarkIsPrime(1000000000000000, b) }
func BenchmarkIsPrime10pow18(b *testing.B)  { benchmarkIsPrime(1000000000000000000, b) }
