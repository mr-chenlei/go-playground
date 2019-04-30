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
	ctx, _ := context.WithCancel(context.Background())
	ctxA, cancelA := context.WithCancel(ctx)
	ctxB, cancelB := context.WithCancel(ctx)

	go printA(ctxA)
	go printB(ctxB)

	time.Sleep(5 * time.Second)
	cancelA()
	time.Sleep(5 * time.Second)
	cancelB()

	select {}
}
