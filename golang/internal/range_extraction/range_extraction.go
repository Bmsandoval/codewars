package main

import (
	"strconv"
)

func Solution(list []int) string {
	result := ""
	count := 0
	for i,_ := range list {
		count++
		if i == 0 {
			result+=strconv.Itoa(list[i])
			continue
		}
		if list[i-1] == list[i]-1 { // previous one is one less than current one, start of a range
			if i == len(list) - 1 { // capture last element of list
				if count == 2 {     // ranges of 2 don't get a dash
					result += ","
				} else {			// ranges of 3+ get a dash
					result += "-"
				}
				result += strconv.Itoa(list[i])
			} else {				// not the last element of a list
				if list[i]+1 != list[i+1] { // the next element does not match the current range, close out this range
					if count == 2 { // ranges of 2 don't get a dash
						result += ","
					} else {		// ranges of 3+ get a dash
						result += "-"
					}
					result += strconv.Itoa(list[i])
					count = 0 		// current range is over, reset the counts
				}
			}
		} else {
			result += "," + strconv.Itoa(list[i])
			count = 1 				// not currently a range based on previous input, but might be with the next val
		}
	}
	return result
}
