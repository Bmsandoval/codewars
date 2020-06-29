package main

var Mapping = map[rune]rune{ '{': '}', '[': ']', '(': ')', }

func ValidBraces(str string) bool {
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

// ##########################################################################################
// codewars doesn't allow the stack package. I'll have to make this one myself
type Stack struct {
	stack []interface{}
}

func NewStack() Stack {
	return Stack{
		stack: []interface{}{},
	}
}

func (s *Stack) Push(i interface{}) {
	s.stack = append(s.stack, i)
}

func (s *Stack) Peek() interface{} {
	if s.Len() > 0 {
		return s.stack[s.Len()-1]
	}
	return nil
}

func (s *Stack) Pop() {
	s.stack = s.stack[:s.Len()-1]
}

func (s *Stack) Len() int {
	return len(s.stack)
}

