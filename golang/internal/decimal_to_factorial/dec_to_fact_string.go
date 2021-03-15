package main

import "strconv"

func Dec2FactString(nb int) string {
	factorials := map[int]int{}

	total:=0
	factorials[0]=0
	for i:=1;;i++{
		total*=i
		factorials[i]=total
		if nb%total==0{
			break
		}
	}

	num := nb
	result:=""
	for j:=len(factorials)-1; j>=0; j-- {
		var val int
		if factorials[j] > 0 {
			count := num / factorials[j]
			val = count * factorials[j]
			num -= val
		} else {
			val = 0
		}

		var char rune
		if j < 10 {
			char = rune(strconv.Itoa(val)[0])
		} else {
			val -= 10
			char = rune('A' + val)
		}
		result+=string(char)
	}

	return result
}

func FactString2Dec(str string) int {
	return 0
}
