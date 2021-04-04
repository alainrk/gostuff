// Thanks to https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
package golangz

import (
	"fmt"
)

// v is not of any type; it is of interface{} type, Go runtime will perform a
// type conversion (if necessary), and convert the value to an interface{} value
func testEmptyInterface(val interface{}) int {
	assertedTypeVal := val.(int)
	// valueAssertedInt, ok := v.(int) -- no panic
	return assertedTypeVal + 2
}

func TestInterfaces () {
	res := testEmptyInterface(5)
	fmt.Println(res)
}

func Test () {
	TestInterfaces()
}