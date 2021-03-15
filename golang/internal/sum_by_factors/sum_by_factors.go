package main

import (
	"fmt"
)

func SumOfDivided(lst []int) string {
	var Primes = []int{}
	var Sums = []*int{}
	var LastTested = 1

	// for each list item
	for _,l := range lst {
		// test each prime against that item
		for i,p := range Primes {
			// sum each l which is divisible by prime p
			if l%p==0 {
				if Sums[i] == nil {
					tmp := l
					Sums[i] = &tmp
				} else {
					*Sums[i] += l
				}
			}
		}

		for p:=LastTested+1; p<=l; p++ {
			LastTested++
			if ! IsPrime(p) {
				continue
			}

			Primes=append(Primes, p)
			Sums=append(Sums, nil)

			// sum each l which is divisible by prime p
			if l%p==0 {
				tmp := l
				Sums[len(Sums)-1]=&tmp
			}
		}
	}

	result := ""

	for i, p := range Primes {
		if Sums[i] != nil {
			result += fmt.Sprintf("(%d %d)", p, *Sums[i])
		}
	}

	return result
}

func IsPrime(candidate int) bool {
	for i:=2; i<=candidate/2; i++ {
		if candidate%i == 0 {
			return false
		}
	}
	return true
}