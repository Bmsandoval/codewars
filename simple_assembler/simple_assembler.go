package main

import (
	"strconv"
	"strings"
)


func SimpleAssembler(program []string) map[string]int {
	var Registers = map[string]int{}

	for i:=0; i<len(program);i++ {
		instruction := program[i]

		// tokenize the instruction
		tokens := strings.Fields(instruction)

		// THE BIG SWITCH
		command := tokens[0]
		switch command {
		case "mov":
			registerLeft := tokens[1]
			if val, err := strconv.Atoi(tokens[2]); err == nil {
				// converted to int, use the resulting value
				Registers[registerLeft] = val
			} else {
				// use the value contained in the register
				registerRight := tokens[2]
				Registers[registerLeft] = Registers[registerRight]
			}
			break
		case "inc":
			register := tokens[1]
			// increment the value in the register
			Registers[register]++
			break
		case "dec":
			register := tokens[1]
			// decrement the value in the register
			Registers[register]--
			break
		case "jnz":
			lhs := 0
			if val, err := strconv.Atoi(tokens[1]); err == nil {
				// constant
				lhs = val
			} else {
				// register
				lhs = Registers[tokens[1]]
			}

			if lhs != 0 {
				rhs := 0
				if val, err := strconv.Atoi(tokens[2]); err == nil {
					// constant
					rhs = val
				} else {
					// register
					rhs = Registers[tokens[2]]
				}

				i+=rhs-1
			}
			break
		}

	}

	return Registers
}
