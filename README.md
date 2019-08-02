# WASM demo

## Setup

Get JS wasm lib

```
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
```

Compile for wasm

```
GOOS=js GOARCH=wasm go build -o main.wasm
```

Caddyfile

```go
127.0.0.1:8080
mime .wasm application/wasm
```

---

## Hello world

[http://127.0.0.1:8080/hello_world_example/](http://127.0.0.1:8080/hello_world_example/)

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, WebAssembly!")
}
```

---

## DOM example

[http://127.0.0.1:8080/dom_example/](http://127.0.0.1:8080/dom_example/)

```go
package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	document := js.Global().Get("document")
	p := document.Call("createElement", "p")
	p.Set("innerHTML", "Hello WASM from Go!")
	document.Get("body").Call("appendChild", p)
}
```

---

## Html button to call go code

[http://127.0.0.1:8080/button_example/](http://127.0.0.1:8080/button_example/)

```go
package main

import (
	"fmt"
	"syscall/js"
)

var done chan struct{}

func main() {
	var cb js.Func
	cb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		fmt.Println("button clicked")
		return nil
	})
	js.Global().Get("document").Call("getElementById", "myButton").Call("addEventListener", "click", cb)
	<-done
}
```

```html
<script src="../wasm_exec.js"></script>
<script>
const go = new Go();
WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
	go.run(result.instance);
});
</script>
<button id="myButton">Test button</button>
```

---

## Threads example

[http://127.0.0.1:8080/threads_example/](http://127.0.0.1:8080/threads_example/)

```go
package main

import (
	"math/rand"
	"strconv"
	"sync"
	"syscall/js"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	nb := 20
	wg := sync.WaitGroup{}
	wg.Add(nb)
	for i := 0; i < nb; i++ {
		go func(i int) {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			document := js.Global().Get("document")
			p := document.Call("createElement", "p")
			p.Set("innerHTML", "Div "+strconv.Itoa(i))
			document.Get("body").Call("appendChild", p)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
```

---

## Go call Js

[http://127.0.0.1:8080/go_call_js_ecxample/](http://127.0.0.1:8080/go_call_js_ecxample/)

```html
<script src="../wasm_exec.js"></script>
<script>
function jsFn(a, b) {
    return a + b;
}

const go = new Go();
WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
    go.run(result.instance);
});
</script>
```

```go
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
```

---

## Js call Go

[http://127.0.0.1:8080/js_call_go_example/](http://127.0.0.1:8080/js_call_go_example/)

```go
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
```

```html
<script src="../wasm_exec.js"></script>
<script>
function startCb() {
    console.log(window.goFn(1, 2));
    console.log(window.goFn(1, "2"));
    console.log(window.goFn(3, 4.2));
}

const go = new Go();
WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
    go.run(result.instance);
});
</script>
```

---

## TinyGo

[https://tinygo.org/webassembly/webassembly/](https://tinygo.org/webassembly/webassembly/)

[http://127.0.0.1:8080/tinygo_example/](http://127.0.0.1:8080/tinygo_example/)

```
tinygo build -o main.wasm -target wasm ./main.go
```

wasm_exec.js for tinygo:
[https://github.com/tinygo-org/tinygo/blob/master/targets/wasm_exec.js](https://github.com/tinygo-org/tinygo/blob/master/targets/wasm_exec.js)


NormalGo:

```
-rwxr-xr-x  1 agilbert  staff   1.4M Aug  2 16:01 main.wasm
```

TinyGo:

```
-rwxr-xr-x  1 agilbert  staff    31K Aug  2 16:09 main.wasm
```