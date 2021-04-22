// Thanks to https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
package example

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

func printAll(vals []interface{}) {
	for _, val := range vals {
			fmt.Println(val)
	}
}

func TestInterfaces () {
	res := testEmptyInterface(5)
	fmt.Println(res)

	names := []string{"this", "is", "as", "ugly", "as", "duck"}
	vals := make([]interface{}, len(names))
	for i, v := range names {
			vals[i] = v
	}
	printAll(vals)
}

// func Test () {
// 	TestInterfaces()
// }