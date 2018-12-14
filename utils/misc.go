package utils

import "strings"

// WordCount : count the words in `string`
func WordCount(s string) map[string]int {
	wc := make(map[string]int)
	tokens := strings.Fields(s)

	for i := 0; i < len(tokens); i++ {
		v := wc[tokens[i]]
		v++
		wc[tokens[i]] = v
	}
	return wc
}

// Fibonacci :generate fibonacci sequence.
func Fibonacci() func() int {
	pos := 0
	prev := 1
	next := 1
	return func() int {
		if pos == 0 || pos == 1 {
			pos++
			return next
		}
		/*
			curr := next
			next += prev
			prev = curr
		*/
		prev, next = next, prev+next
		pos++

		return next
	}
}
