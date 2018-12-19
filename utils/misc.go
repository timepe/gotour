package utils

import (
	"io"
	"strings"
)

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

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(p []byte) (int, error) {
	q := make([]byte, len(p))
	n, err := r.Read(q)
	if err != nil {
		return -1, err
	}

	for i := 0; i < n; i++ {
		p[i] = q[i] + (q[i]-'a'+13)%26
	}

	return n, nil
}

type tree struct {
	Left  *tree
	Value int
	Right *tree
}

func walk(t *tree, c chan int) {
	if t != nil {
		walk(t.Left, c)
		c <- t.Value
		walk(t.Right, c)
	}
}

func Same(t1, t2 *tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	for x := range c1 {
		y, ok := <-c2
		if !ok || x != y {
			return false
		}
	}
	y, ok := <-c2
	if !ok {
		return false
	}

	return true
}
