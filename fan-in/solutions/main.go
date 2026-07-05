
package main

import (
	"fmt"
	"sync"
)

var p = fmt.Println

func fanIn(chans ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(chans))

	for _, ch := range chans {
		go func(c <-chan int) {
			defer wg.Done()
			for data := range c {
				out <- data
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	gen1, gen2 := generator(15), generator(15)

	for data := range fanIn(gen1, gen2) {
		p(data)
	}
}

func generator(n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := range n {
			out <- i + 1
		}
	}()
	return out
}
