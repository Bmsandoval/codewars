package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIpValidation(t *testing.T) {
	unitTestData := []struct {
		Description string
		Input       string
		Expect      bool
	}{
		{
			Description: "happy path (1)",
			Input:       "1.2.3.4",
			Expect:      true,
		},
		{
			Description: "happy path (2)",
			Input:       "123.45.67.89",
			Expect:      true,
		},
		{
			Description: "can detect not enough octets",
			Input:       "1.2.3",
			Expect:      false,
		},
		{
			Description: "can detect too many octets",
			Input:       "1.2.3.4.5",
			Expect:      false,
		},
		{
			Description: "can detect large octet values",
			Input:       "123.456.78.90",
			Expect:      false,
		},
		{
			Description: "can detect leading zeros in octet values",
			Input:       "123.045.067.089",
			Expect:      false,
		},
	}

	for _, testData := range unitTestData {
		t.Run(testData.Description, func(t *testing.T) {
			got := IpValidation(testData.Input)
			assert.Equal(t, testData.Expect, got)
		})
	}
}

var result bool

func BenchmarkValidBraces(b *testing.B) {
	var r bool
	for n := 0; n < b.N; n++ {
		r = IpValidation("1.2.3.4")
	}
	result = r
}
