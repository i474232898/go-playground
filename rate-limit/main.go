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
	time.Sleep(500 * time.Millisecond)
	p("sending request", request.Payload)
	return nil
}

// limit camount of connections
// limit amount of gorutines
// limit rps
func (c client) WithLimiter(ctx context.Context, reqs []Request) {
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
