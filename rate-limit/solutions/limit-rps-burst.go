package main

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"
)

var p = fmt.Println

type Request struct {
	Payload string
}
type Client interface {
	SendRequest(ctx context.Context, request Request) error
	WithLimiter(ctx context.Context, requests []Request)
}
type client struct{}

func (c client) SendRequest(ctx context.Context, request Request) error {
	p("sending request", request.Payload)
	time.Sleep(500 * time.Millisecond)
	return nil
}

// limit rps
var rps = 10
var burst = 5
func (c client) WithLimiter(ctx context.Context, reqs []Request) {
	ticker := time.NewTicker(time.Second / time.Duration(rps))
	tickets := make(chan struct{}, burst)
	wg := sync.WaitGroup{}

	go func(){
		for range burst{
			tickets <- struct{}{}
		}
	}()
	go func() {
		for {
			select {
				case <- ticker.C:
				tickets <- struct{}{}
			}
		}
	}()

	wg.Add(len(reqs))
	for _, r := range reqs {
		<-tickets
		go func(){
			defer wg.Done()
			c.SendRequest(ctx, r)
		}()
	}
	wg.Wait()
}

func main() {
	ctx := context.Background()
	c := client{}
	rq := 100
	requests := make([]Request, rq)
	for i := 0; i < rq; i++ {
		requests[i] = Request{Payload: strconv.Itoa(i)}
	}
	c.WithLimiter(ctx, requests)
}
