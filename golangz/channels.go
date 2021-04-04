package golangz

import (
	"fmt"
	"strconv"
)

func TestChannels() {
	messages := make(chan string, 2)
	messages <- "buffered"
	// UnBuffered
	toast := make(chan string)
	loops := 5
	for i := 0; i < loops; i++ {
		go func(x int) {
			var msg string = "unbuffered_" + strconv.Itoa(x)
			toast <- msg
		}(i)
	}

	for i := 0; i < loops; i++ {
		fmt.Println(<-toast)
	}

	// Buffered

	messages <- "channel"
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

func Test() {
	TestChannels()
}
