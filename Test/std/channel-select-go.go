package main

import (
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

/*
go并发素数筛
在fliter开启的goroutine没有退出，会一直存在，
在fliterWithCancle中开启的goroutine通过cancleContext退出
*/
func generate() chan int {
	in := make(chan int)
	go func() {
		for i := 2; ; i++ {
			in <- i
		}
	}()
	return in
}

func fliter(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func fliterWithCancle(ctx context.Context, in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case i := <-in:
				if i%prime != 0 {
					out <- i
				}
			}
		}
	}()
	return out
}

func primefliter() {
	go func() {
		http.ListenAndServe(":4321", nil)
	}()
	out := generate()
	for i := 1; i <= 1000; i++ {
		prime := <-out
		out = fliter(out, prime)
	}
	fmt.Println("done")
	select {}
}

func primefliterWithCancle() {
	go func() {
		http.ListenAndServe(":4321", nil)
	}()

	ctx, cancle := context.WithCancel(context.Background())
	out := generate()
	for i := 0; i < 1000; i++ {
		prime := <-out
		out = fliterWithCancle(ctx, out, prime)
	}
	cancle()
	fmt.Println("done")
	select {}
}
