package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"

	// "runtime"
	// "sync"
	// "sync/atomic"
	"time"
)

var p = fmt.Println

func unpredictable() int {
	n := rand.Intn(40)
	time.Sleep(time.Duration(n) * time.Second)
	return n
}

var timeout = 5 * time.Second
func predictable(ctx context.Context) (int, error) {
	ch := make(chan struct{})
	var result int
	
	var cancel context.CancelFunc
	if _, ok := ctx.Deadline(); !ok {
		ctx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()
	}
	
	go func(){
		result = unpredictable()
		close(ch)
	}()
	
	select {
		case <- ch:
			return result, nil
		case <-ctx.Done():
			return 0, errors.New("timed out")
	}
}

func main() {
	_, _ = predictable(context.Background())
}
