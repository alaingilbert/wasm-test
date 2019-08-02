package main

import (
	"fmt"
	"syscall/js"
)

var done chan struct{}

func main() {
	val := js.Global().Call("jsFn", 1, 2)
	fmt.Println("res:", val)
}
