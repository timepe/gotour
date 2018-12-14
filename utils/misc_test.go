package utils

import "testing"

func TestWordCount(t *testing.T) {
	t.Log(WordCount("the other day i went home, and I found the wallet gone disappear."))
}

func TestFibonacci(t *testing.T) {
	f := Fibonacci()
	for i := 0; i < 10; i++ {
		t.Log(f())
	}
}
