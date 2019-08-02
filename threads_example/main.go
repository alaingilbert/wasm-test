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
