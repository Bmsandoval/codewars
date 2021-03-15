package main

func Interpreter(code string) string {
	var b byte = 0
	var bs []byte
	for _, r := range code {
		switch r {
		case '+':
			b += 1
		case '.':
			bs = append(bs, b)
		}
	}
	return string(bs)
}

