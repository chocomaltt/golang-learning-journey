package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(200*time.Millisecond):
		fmt.Println("Job is done.")
	case <-ctx.Done():
		fmt.Println("Cancelled: ", ctx.Err())
	}
}