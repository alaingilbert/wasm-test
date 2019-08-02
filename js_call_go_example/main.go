package main

import (
	"syscall/js"
)

func add(a, b int) int {
	return a + b
}

func main() {
	cb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 2 || args[0].Type() != js.TypeNumber || args[1].Type() != js.TypeNumber {
			return nil
		}
		return add(args[0].Int(), args[1].Int())
	})
	js.Global().Set("goFn", cb)
	js.Global().Call("startCb")
}
