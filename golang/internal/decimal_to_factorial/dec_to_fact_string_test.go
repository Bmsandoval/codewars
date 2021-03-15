package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDec2FactString(t *testing.T) {
	unitTestData := []struct {
		Description string
		Input int
		Expect int
	}{
		{
			Description: "Input:513",
			Input: 513,
			Expect: 531,
		},
		//{
		//	Description: "Input:7811",
		//	Input: 7811,
		//	Expect: 8117,
		//},
		//{
		//	Description: "Input:97531",
		//	Input: 97531,
		//	Expect: -1,
		//},
		//{
		//	Description: "Input:144",
		//	Input: 144,
		//	Expect: 414,
		//},
		//{
		//	Description: "Input:1234567890",
		//	Input: 1234567890,
		//	Expect: 1234567908,
		//},
	}

	for _, testData := range unitTestData {
		t.Run(testData.Description, func(t *testing.T) {
			got := Dec2FactString(testData.Input)
			assert.Equal(t, testData.Expect, got)
		})
	}
}

var result string
func benchmarkDec2FactString(i int, b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		r = Dec2FactString(i)
	}
	result = r
}

func BenchmarkDec2FactString10x3(b *testing.B)  { benchmarkDec2FactString(1000, b) }
func BenchmarkDec2FactString10x6(b *testing.B)  { benchmarkDec2FactString(1000000, b) }
func BenchmarkDec2FactString10x9(b *testing.B)  { benchmarkDec2FactString(1000000000, b) }
func BenchmarkDec2FactString10x12(b *testing.B)  { benchmarkDec2FactString(1000000000000, b) }
func BenchmarkDec2FactString10x15(b *testing.B)  { benchmarkDec2FactString(1000000000000000, b) }
func BenchmarkDec2FactString10x16(b *testing.B)  { benchmarkDec2FactString(1000000000000000000, b) }
