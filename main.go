package main

import (
	"fmt"
	"github.com/alainrk/gostuff/datastruct"
)

func main() {
	hm := datastruct.HashMap{}
	fmt.Println(hm.Get("xxx"))
}