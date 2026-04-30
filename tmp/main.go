package main

import (
	"context"
	"errors"
	"fmt"
	"sync"

	// "sync"
	"time"
)

var p = fmt.Println

type User struct {
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
	res, err := process(ctx, names)
	if err != nil {
		p("an error occured:", err.Error())
	}
	p("time:", time.Since(start))
	p(res)
}

func fetch(ctx context.Context, user User) (string, error) {
	if user.name == "Ann" {
		return "", errors.New("invalid name")
	}
	ch := make(chan struct{})
	
	go func(){
		time.Sleep(1 * time.Second)
		close(ch)
	}()
	
	select{
		case <-ch:
			return user.name, nil
		case <-ctx.Done():
			return "", errors.New("context cancelled")
	}
}

func process(ctx context.Context, users []User) (map[string]int64, error) {
	names := make(map[string]int64, 0)
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(len(users))
	
	ctx, cancel := context.WithCancel(ctx)
	
	var commonError error
	
	defer cancel()
	for _, u := range users {
		go func() {
			defer wg.Done()
			name, err := fetch(ctx, u)
			if err != nil {
				sync.OnceFunc(func(){
					cancel()
					commonError = err
				})()
			}
			
			mu.Lock()
			defer mu.Unlock()
			names[name] = names[name] + 1
		}()
	}
	wg.Wait()
	
	if commonError != nil {
		return nil, commonError
	}

	return names, nil
}
