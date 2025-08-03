package easy

import "fmt"

var mapper = map[rune]rune{
	')': '(',
	'}': '{',
	']': '[',
}

func isValid(s string) bool {
	var stack []rune = []rune{}
	if len(s) < 2 {
		return false
	}

	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
			continue
		}
		if len(stack) == 0 || stack[len(stack)-1] != v {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) > 0 {
		return false
	}
	return true
}
