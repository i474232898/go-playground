package main

import (
	"context"
	"fmt"
	"time"
)

var p = fmt.Println

type User struct{
	name string
}

func main() {
	names := []User{
		{name: "Ann"},
		{name: "Cindy"},
		{name: "Bob"},
		{name: "Ann"},
	}
	
	ctx := context.Background()
	
	start := time.Now()
	
	res, err := process(ctx, names, realFetcher{})
	if err != nil {
		p("an error occured:", err.Error())
	}
	p("time:", time.Since(start))
	p(res)
}

func fetch(_ context.Context, user User) (string, error) {
	time.Sleep(1 * time.Second)
	return user.name, nil
}

type Fetcher interface {
	fetch(ctx context.Context, user User) (string, error)
}

type realFetcher struct{}

func (realFetcher) fetch(ctx context.Context, user User) (string, error) {
	return fetch(ctx, user)
}

func process(ctx context.Context, users []User, f Fetcher) (map[string]int64, error) {
	names := make(map[string]int64, 0)
	
	for _, u := range users {
		name, err := f.fetch(ctx, u)
		if err != nil {}
		
		names[name] = names[name] + 1
	}
	
	return names, nil
}

// tasks
// 1. make process running in parallel


