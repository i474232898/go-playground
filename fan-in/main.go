package main

import (
	"fmt"
	"sync"
)

var p = fmt.Println

func generator(n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := range n {
			out <- i
		}
	}()

	return out
}

func main() {
	one := generator(5)
	two := generator(5)
	
	out := fain(one, two)
	
	for data := range out {
		p(data)
	}
}

func fain(chans ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(chans))

	for _, ch := range chans {
		go func() {
			defer wg.Done()
			for d := range ch {
				out <- d
			}
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
