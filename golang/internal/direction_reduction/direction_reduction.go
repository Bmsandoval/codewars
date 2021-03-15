package main

var Mapping = map[string]string{
	"NORTH": "SOUTH",
	"SOUTH": "NORTH",
	"EAST": "WEST",
	"WEST": "EAST",
}

func DirReduc(arr []string) []string {
	// new stack
	stack := []string{}

	for _, dir := range arr {
		if len(stack) > 0 && Mapping[dir] == stack[len(stack)-1] {
			// pop complementary direction
			stack = stack[:len(stack)-1]
		} else {
			// push direction
			stack = append(stack, dir)
		}
	}

	// return final directions
	return stack
}
