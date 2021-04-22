package example

import (
	"flag"
	"fmt"
)

// $> go build . && ./gostuff -n 43 --flagname=100 --par xxxx
func TestFlags() {
	var nFlag = flag.Int("n", -1, "help message for flag n")

	var flagvar int
	flag.IntVar(&flagvar, "flagname", -1, "help message for flagname")

	var strpar string
	flag.StringVar(&strpar, "par", "default", "help message for string param")

	flag.Parse()

	fmt.Println(*nFlag)
	fmt.Println(flagvar)
	fmt.Println(strpar)
}

// func Test() {
// 	TestFlags()
// }
