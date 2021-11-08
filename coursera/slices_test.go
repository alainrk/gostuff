package main

import (
	"testing"
)

func TestSlices1(t *testing.T) {
	// sli := []string{"a", "b", "c", "d", "e", "f"} // Slice (it hasn't length specified)
	arr := [...]string{"a", "b", "c", "d", "e", "f"} // Array (it has length specified => implicit with "...")

	s1 := arr[2:4]
	s2 := arr[3:5]

	s1[1] = "x"
	if s2[0] != "x" && arr[3] != "x" {
		t.Error("should be 'x'")
	}
}

func TestSlices2(t *testing.T) {
	s1 := make([]int, 5)     // length (of the slice) = capacity (length of the underlying array)
	s2 := make([]int, 5, 10) // length (of the slice) < capacity (length of the underlying array)

	if len(s1) != 5 || cap(s1) != 5 {
		t.Error("should be 5 and 5")
	}

	if len(s2) != 5 || cap(s2) != 10 {
		t.Error("should be 5 and 10")
	}
}
