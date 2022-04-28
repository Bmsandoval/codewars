package main

var Mapping = map[rune]rune{ '{': '}', '[': ']', '(': ')', }

func Run(str string) bool {
	s := NewStack()

	for _, r := range str {
		if match, exists := Mapping[r]; exists {
			// push the matching closing brace
			s.Push(match)
		} else if s.Peek() != nil && r == s.Peek().(rune) {
			// pop closing brace if matching brace found
			s.Pop()
		} else {
			// mismatch or invalid character
			return false
		}
	}

	// return false if still expecting a closing brace
	return s.Len() == 0
}