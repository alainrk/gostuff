package example

import (
	"fmt"
	"crypto/sha256"
)

// nil slice
func TestSha(str string) {
	sum := sha256.Sum256([]byte(str))
	fmt.Printf("%x", sum)
}

// func Test() {
// 	TestSha("diocan")
// }