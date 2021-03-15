package main

import (
	"log"
	"strconv"
)

func main() {
	num := 124888
	log.Println(DigitalRoot(num))
}

func SumOfDigits(num int, digits int) int {
	total := 0
	for i := 1; i<=digits; i++ {
		total += NthDigit(num, i)
	}
	return total
}

func NthDigit(num int, digit int) int {
	val := num
	// throw away trailing values
	for n := 1; n < digit; n++ {
		val /= 10
	}
	// throw away preceding values
	return val % 10
}

func DigitalRoot(n int) int {
	val := n
	for ;; {
		digits := len(strconv.Itoa(val))
		if digits == 1 {
			return val
		}
		val = SumOfDigits(val, digits)
	}
}
