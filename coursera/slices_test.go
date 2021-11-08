package main

import (
	"testing"
)

func TestSlices1(t *testing.T) {
	arr := [...]string{"a", "b", "c", "d", "e", "f"}

	s1 := arr[2:4]
	s2 := arr[3:5]

	s1[1] = "x"
	if s2[0] != "x" && arr[3] != "x" {
		t.Error("should be 'x'")
	}
}
