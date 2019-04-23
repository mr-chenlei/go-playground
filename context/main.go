package main

import (
	"context"
	"fmt"
	"time"
)

func printA(ctx context.Context) {
	t := time.NewTimer(0)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("printA done")
			return
		case <-t.C:
			fmt.Println("A")
			t.Reset(1 * time.Second)
		}
	}
}

func printB(ctx context.Context) {
	t := time.NewTimer(0)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("printB done")
			return
		case <-t.C:
			fmt.Println("B")
			t.Reset(1 * time.Second)
		}
	}
}

func printC(ctx context.Context) {
	t := time.NewTimer(0)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("printC done")
			return
		case <-t.C:
			fmt.Println("C")
			t.Reset(1 * time.Second)
		}
	}
}

func main() {
	ctxA := context.WithValue(context.Background(), "fmt", "A")
	ctxB := context.WithValue(context.Background(), "fmt", "B")

	go printA(ctxA)
	go printB(ctxB)

	time.Sleep(5 * time.Second)

	ctxA.Done()

	select {}
}
