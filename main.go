package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	Background()
	// Todo()
	// WithCancel()
	// WithDeadline()
	// WithTimeout()
	// WithValue()

}

// Background
func Background() {
	ctx := context.Background()

	fmt.Println(ctx)
}

// TODO
func Todo() {
	ctx := context.TODO()

	fmt.Println(ctx)
}

// WithCancel
func WithCancel() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println(ctx.Err())

	cancel()

	fmt.Println(ctx.Err())
}

// WithDeadline
func WithDeadline() {
	tmp := time.Now().Add(time.Second * 2)
	ctx, _ := context.WithDeadline(context.Background(), tmp)

	fmt.Println(ctx.Err())

	time.Sleep(time.Second * 4)

	fmt.Println(ctx.Err())
}

// WithTimeout
func WithTimeout() {
	tmp := time.Second * 2
	ctx, _ := context.WithTimeout(context.Background(), tmp)

	fmt.Println(ctx.Err())

	time.Sleep(time.Second * 4)

	fmt.Println(ctx.Err())
}

// WithValue
func WithValue() {
	ctx := context.WithValue(context.Background(), "chave", "valor")

	fmt.Println(ctx.Value("chave"))
}
