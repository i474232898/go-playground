package main

import (
	"context"
	"testing"
)

type mockFetcher struct {}
func (*mockFetcher) fetch(_ context.Context, u User) (string, error) {
	return u.name, nil
}

func TestProcess(t *testing.T) {
	users := []User{{name: "Ann"}, {name: "Ann"}, {name: "Bob"}}
	got, err := process(context.Background(), users, &mockFetcher{})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got["Ann"] != 2 {
		t.Errorf("Ann: got %d, want 2", got["Ann"])
	}
	if got["Bob"] != 1 {
		t.Errorf("Bob: got %d, want 1", got["Bob"])
	}
}
