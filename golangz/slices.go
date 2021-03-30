package golangz

import (
	"fmt"
)

func Test () {
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