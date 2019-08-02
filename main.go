package main

import (
	"syscall/js"
)

var done chan struct{}

func main() {
	cb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return args[0].Int() + args[1].Int()
	})
	js.Global().Set("goFn", cb)
	js.Global().Call("startCb")
	<-done
}
