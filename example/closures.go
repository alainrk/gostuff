package example

import (
	"fmt"
)

func iterator (step int) func () int {
	count := 0
	return func () int {
		count += step
		return count
	}
}

func TestClosures() {
	doubler := iterator(2)
	tripler := iterator(3)

	fmt.Println(doubler())
	fmt.Println(doubler())
	fmt.Println(doubler())

	fmt.Println(tripler())
	fmt.Println(tripler())
	fmt.Println(tripler())
}

// func Test() {
// 	TestClosures()
// }
