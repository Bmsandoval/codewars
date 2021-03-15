package main

import (
	"math"
	"sort"
)

func NextBigger(n int) int {
	rhs := []int{}
	lhs := n
	for i := 0; lhs != 0; i++ {
		// move lhs's rightmost digit to rhs
		rhs=append(rhs,lhs % 10)
		lhs /= 10

		// get left digit
		leftDigit := lhs % 10

		// find first one we can swap
		if !(lhs == 0 || leftDigit>=rhs[i]) {
			nextLargest:=10
			digit:=0
			for d,n := range rhs {
				if n > leftDigit && n <= nextLargest{
					nextLargest=n
					digit=d
				}
			}
			rhs[digit]=leftDigit
			//// we found one to swap
			//// throw away rightmost digit
			lhs /= 10
			// add back all the zeroes we threw away
			lhs *= int(math.Pow10( i+2 ))
			// add new digit
			lhs += nextLargest * int(math.Pow10( i+1 ))

			sort.Sort(sort.Reverse(sort.IntSlice(rhs)))

			for d,n := range rhs {
				lhs += n * int(math.Pow10(d))

			}

			// return swapped result
			return lhs
		}
	}
	// if no swap necessary, return -1
	return -1
}
