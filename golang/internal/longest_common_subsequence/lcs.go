package main

import "log"

// LCS returns the longest subsequence common to s1 and s2
func LCS(s1, s2 string) string {
	prevIdx := -1
	sequence := ""
	candidateSequence := ""

	subSq := s1
	for ;len(subSq)>len(sequence); {
		for _,r := range subSq {
			for i:=prevIdx+1; i<len(s2); i++ {
				log.Printf("r1: %s, r2: %s", string(r), string(s2[i]))
				if r == int32(s2[i]) {
					prevIdx = i
					candidateSequence += string(r)
					break
				}
			}
		}

		if len(candidateSequence) > len(sequence) {
			sequence, candidateSequence = candidateSequence, ""
		}
		subSq = subSq[1:]
		prevIdx=-1
	}

	sequence += "1"
	return sequence
}
