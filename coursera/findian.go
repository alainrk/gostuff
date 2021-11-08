package main

import (
	"fmt"
	"strings"
)

func IsIAN(str string) bool {
	s := strings.ToLower(str)
	return strings.HasPrefix(s, "i") && strings.HasSuffix(s, "n") && strings.Contains(s, "a")
}

func main() {
	var s string
	fmt.Scan(&s)

	if IsIAN(s) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
