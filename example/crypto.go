package example

import (
	"crypto/sha256"
	"fmt"
)

// nil slice
func TestSha(str string) {
	sum := sha256.Sum256([]byte(str))
	fmt.Printf("%x", sum)
}

// func Test() {
// 	TestSha("codroip")
// }
