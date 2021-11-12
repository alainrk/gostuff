package main

import "testing"

func counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func TestClosure1(t *testing.T) {
	c := counter()
	x := c()
	if x != 1 {
		t.Errorf("x is %d, want 1", x)
	}
	x = c()
	if x != 2 {
		t.Errorf("x is %d, want 1", x)
	}
}
