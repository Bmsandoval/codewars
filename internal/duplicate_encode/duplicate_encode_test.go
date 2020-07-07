package main

import (
	"codewars/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDuplicateEncode(t *testing.T) {
	unitTestData := []struct {
		Description string
		Input string
		Expect string
	}{
		{
			Description: "happy path",
			Input: "happy path",
			Expect: "))))(())()",
		},
	}

	for _, testData := range unitTestData {
		t.Run(testData.Description, func(t *testing.T) {
			got := DuplicateEncode(testData.Input)
			assert.Equal(t, testData.Expect, got)
		})
	}
}

var result string
func benchmarkDuplicateEncode(i int, b *testing.B) {
	var r string
	randString := pkg.RandomString(i)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = DuplicateEncode(randString)
	}
	result = r
}

func BenchmarkDuplicateEncode1(b *testing.B)  { benchmarkDuplicateEncode(1, b) }
func BenchmarkDuplicateEncode2(b *testing.B)  { benchmarkDuplicateEncode(2, b) }
func BenchmarkDuplicateEncode3(b *testing.B)  { benchmarkDuplicateEncode(3, b) }
func BenchmarkDuplicateEncode10(b *testing.B)  { benchmarkDuplicateEncode(10, b) }
func BenchmarkDuplicateEncode20(b *testing.B)  { benchmarkDuplicateEncode(20, b) }
func BenchmarkDuplicateEncode40(b *testing.B)  { benchmarkDuplicateEncode(40, b) }
