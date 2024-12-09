package main

import (
	"context"
	"fmt"
	"time"
)

// Context
// - Controlling timeouts
// - Cancelling go routines
// - Passing Metadata

func main() {
	// ctx := context.Background()
	// ExampleTimeout(ctx)

	ExampleWithValues()
}

func ExampleTimeout(ctx context.Context) {

	ctxWithTimout, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	done := make(chan struct{})

	go func() {
		time.Sleep(3 * time.Second)
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("Called the api")
	case <-ctxWithTimout.Done():
		fmt.Printf("Api expired: %v\n", ctxWithTimout.Err())
	}
}

func ExampleWithValues() {
	type key int
	const UserKey key = 0

	ctx := context.Background()

	ctxWithValue := context.WithValue(ctx, UserKey, "shres")

	if userID, ok := ctxWithValue.Value(UserKey).(string); ok {
		fmt.Printf("This is the user: %s\n", userID)
	} else {
		fmt.Println("This is a protected route: No userId found")
	}
}
