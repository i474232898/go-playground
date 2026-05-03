package main

import (
	"context"
	"fmt"
	"strconv"
	// "sync"
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
	time.Sleep(500 * time.Millisecond)
	p("sending request", request.Payload)
	return nil
}

// limit amount of gorutines
// limit rps
// Using the semaphore channel for both limiting and waiting
var maxGorutines = 10
func (c client) WithLimiter(ctx context.Context, reqs []Request) {
	tokens := make(chan struct{}, maxGorutines)
	for range maxGorutines{
		tokens<-struct{}{}
	}
	for _, r := range reqs{
		<-tokens
		go func(){
			defer func(){
				tokens<-struct{}{}
			}()
			c.SendRequest(ctx, r)
		}()
	}
	for range maxGorutines{
		<-tokens
	}
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
