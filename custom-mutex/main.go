package main

import (
	"fmt"
	"sync"
	// "sync/atomic"
	// "time"
)

var p = fmt.Println

type CustomMu struct {
	locked chan struct{}
}

func (cm *CustomMu) Lock() {
	cm.locked <- struct{}{}
}
func (cm *CustomMu) Unlock() {
	<- cm.locked
}

func main() {
	var wg sync.WaitGroup
	mu := CustomMu{ locked: make(chan struct{}, 1)}
	c := 0

	wg.Add(1000)
	for range 1000 {
		go func() {
			defer wg.Done()
			mu.Lock()
			c++
			mu.Unlock()
		}()
	}
	wg.Wait()
	p(c)
}
