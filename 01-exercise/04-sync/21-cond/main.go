package main

import (
	"fmt"
	"sync"
)

var sharedRsc = make(map[string]interface{})

func main() {
	var wg sync.WaitGroup
	m := sync.Mutex{}
	n := sync.NewCond(&m)

	wg.Add(1)
	go func() {
		defer wg.Done()

		//TODO: suspend goroutine until sharedRsc is populated.
		n.L.Lock()
		for len(sharedRsc) == 0 {
			n.Wait()
		}

		fmt.Println(sharedRsc["rsc1"])
		n.L.Unlock()
	}()

	// TODO: writes changes to sharedRsc
	n.L.Lock()
	sharedRsc["rsc1"] = "foo"
	n.Signal()
	n.L.Unlock()
	wg.Wait()
}
