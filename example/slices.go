package example

import (
	"fmt"
)

func TestSliceLenAndCap () {
	// creates a slice of ints with len 10 and cap 20
	myslice := make([]int, 10, 20)
	fmt.Println(len(myslice), cap(myslice), myslice)

	// throws an error
	// myslice[15] = 42

	// expands the slice to its max cap
	myslice = myslice[:cap(myslice)]

	// now it's ok
	myslice[15] = 42
	fmt.Println(len(myslice), cap(myslice), myslice)
}

func TestSliceAndUnderlyingArr () {
	// creates a slice of len and cap 6
	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(len(s), cap(s), s)

	// slice is rewrote, but the underlying arr is still intact
	s = s[:3]
	fmt.Println(s)

	// indeed here I can get all the other elements excluded by the previous slice
	s = s[3:cap(s)]
	fmt.Println(s)
}

// nil slice
func TestNilSliceAndAppend() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("Slice is nil")
	}

	// Throws an err
	// s[0] = 123

	// len increases with the appending, cap doubles when the half on len is passed
	for i := 0; i < 16; i++ {
		s = append(s, i)
		fmt.Println(s, "len:", len(s), "cap:", cap(s))
	}
}

// func Test() {
// 	TestNilSlice()
// }