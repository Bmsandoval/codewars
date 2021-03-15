package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLCS(t *testing.T) {
	unitTestData := []struct {
		Description string
		Input1 string
		Input2 string
		Expect string
	}{
		{
			Description: "abcde/abc",
			Input1: "abcdef",
			Input2: "abc",
			Expect: "abc",
		},
		//{
		//	Description: "abcdef/abf",
		//	Input1: "abcdef",
		//	Input2: "acf",
		//	Expect: "acf",
		//},
		//{
		//	Description: "132535365/123456789",
		//	Input1: "132535365",
		//	Input2: "123456789",
		//	Expect: "12356",
		//},
		//{
		//	Description: "123512345123456/123456",
		//	Input1: "123512345123456",
		//	Input2: "123456",
		//	Expect: "123456",
		//},
	}

	for _, testData := range unitTestData {
		t.Run(testData.Description, func(t *testing.T) {
			got := LCS(testData.Input1, testData.Input2)
			assert.Equal(t, testData.Expect, got)
		})
	}
}

var result string
func BenchmarkLCS(b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		r = LCS("132535365", "123456789")
	}
	result = r
}
