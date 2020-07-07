package main

import (
	"codewars/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpinWords(t *testing.T) {
	unitTestData := []struct {
		Description string
		Input string
		Expect string
	}{
		{
			Description: "happypath",
			Input: "happypath",
			Expect: "htapyppah",
		},
		{
			Description: "happy path",
			Input: "happy path",
			Expect: "yppah path",
		},
		{
			Description: "another random testypoo",
			Input: "another random testypoo",
			Expect: "rehtona modnar oopytset",
		},
	}

	for _, testData := range unitTestData {
		t.Run(testData.Description, func(t *testing.T) {
			got := SpinWords(testData.Input)
			assert.Equal(t, testData.Expect, got)
		})
	}
}

var result string
func benchmarkSpinWords(i int, b *testing.B) {
	var r string
	randString := pkg.RandomString(i)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = SpinWords(randString)
	}
	result = r
}

func BenchmarkSpinWords1(b *testing.B)  { benchmarkSpinWords(1, b) }
func BenchmarkSpinWords2(b *testing.B)  { benchmarkSpinWords(2, b) }
func BenchmarkSpinWords3(b *testing.B)  { benchmarkSpinWords(3, b) }
func BenchmarkSpinWords10(b *testing.B)  { benchmarkSpinWords(10, b) }
func BenchmarkSpinWords20(b *testing.B)  { benchmarkSpinWords(20, b) }
func BenchmarkSpinWords40(b *testing.B)  { benchmarkSpinWords(40, b) }
