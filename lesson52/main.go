package main

import (
	"context"
	"fmt"
)

func main() {
	// ctx := context.WithValue(context.Background(), "ctxKey1", "ctxVal")
	// ctx = context.WithValue(ctx, "ctxKey2", "ctxVal2")
	// ctx = context.WithValue(ctx, "ctxKey3", "ctxVal3")
	// go func(ctx context.Context) {
	// 	// 读取 ctx 的 value
	// 	data, ok := ctx.Value("ctxKey1").(string)
	// 	if ok {
	// 		fmt.Printf("sub goroutine get value from parent goroutine, val=%s\n", data)
	// 	}
	// 	data, ok = ctx.Value("ctxKey2").(string)
	// 	if ok {
	// 		fmt.Printf("sub goroutine get value from parent goroutine, val=%s\n", data)
	// 	}
	// 	data, ok = ctx.Value("ctxKey3").(string)
	// 	if ok {
	// 		fmt.Printf("sub goroutine get value from parent goroutine, val=%s\n", data)
	// 	}
	// }(ctx)

	// ctxVal := make(map[string]string)
	// ctxVal["k1"] = "v1"
	// ctxVal["k2"] = "v2"
	// ctx := context.WithValue(context.Background(), "ctxKey1", ctxVal)
	// go func(ctx context.Context) {
	// 	// 读取 ctx 的 value
	// 	data, ok := ctx.Value("ctxKey1").(map[string]string)
	// 	if ok {
	// 		fmt.Printf("sub goroutine get value from parent goroutine, val=%+v\n", data)
	// 	}
	// }(ctx)

	// ctxVal := make(map[string]string)
	// ctxVal["k1"] = "v1"
	// ctxVal["k2"] = "v2"
	// ctx := context.WithValue(context.Background(), "ctxKey1", ctxVal)
	// go func(ctx context.Context) {
	// 	// 读取 ctx 的 value
	// 	data, ok := ctx.Value("ctxKey1").(map[string]string)
	// 	if ok {
	// 		ctxVal := make(map[string]string)
	// 		for k, v := range data {
	// 			ctxVal[k] = v
	// 		}
	// 		ctxVal["k3"] = "v3"
	// 		ctx = context.WithValue(ctx, "ctxKey1", ctxVal)
	// 		data, ok := ctx.Value("ctxKey1").(map[string]string)
	// 		if !ok {
	// 			fmt.Printf("sub goroutine get value from parent goroutine, val=%+v\n", nil)
	// 		}
	// 		fmt.Printf("sub goroutine get value from parent goroutine, val=%+v\n", data)
	// 	}
	// }(ctx)

	// time.Sleep(1 * time.Second)

	// ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	// defer cancel()
	// select {
	// case <-time.After(1 * time.Second):
	// 	fmt.Println("overslept")
	// case <-ctx.Done():
	// 	fmt.Println(ctx.Err())
	// }

	// gen := func() <-chan int {
	// 	dst := make(chan int)
	// 	go func() {
	// 		var n int
	// 		for {
	// 			dst <- n
	// 			n++
	// 		}
	// 	}()
	// 	return dst
	// }
	// for n := range gen() {
	// 	fmt.Println(n)
	// 	if n == 5 {
	// 		break
	// 	}
	// }

	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		go func() {
			var n int
			for {
				select {
				case <-ctx.Done():
					return
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			cancel()
			break
		}
	}
}
