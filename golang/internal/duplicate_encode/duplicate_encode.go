package main

import "strings"

func DuplicateEncode(word string) string {
	sByte := []byte(strings.ToLower(word))
	characters := map[uint8]int{}
	for i:=0;i<len(sByte);i++{
		character := sByte[i]

		val,exists := characters[character]
		if ! exists {
			// record first occurrence
			sByte[i]='('
			characters[character]=i
		} else {
			if val >= 0 {
				// update first occurrence
				sByte[val] = ')'
				characters[character] = -1
			}
			// set current occurrence
			sByte[i]=')'
		}
	}

	return string(sByte)
}