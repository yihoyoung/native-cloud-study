package main

import (
	"context"
	"fmt"
	_ "lesson2/lihaorong"
	"time"
)

// func Consumer(queue *chan int, timeOut int) error {
// 	return error.New("kkkk")
// }

// func producer(queue *chan int)

func producer(ch chan int) {
	n := 0
	for {
		fmt.Println("create data: ", n)
		ch <- n
		n++
		time.Sleep(time.Second)
	}
}

func consumer(ch chan int) {
	for n := range ch {
		fmt.Println("received: ", n)
		time.Sleep(time.Second)
	}
}

func process(ctx context.Context, duration time.Duration) {
	fmt.Println("duration: ", duration)
	select {
	case <-ctx.Done():
		fmt.Println("process is done.")
	}
}
func main() {
	ch := make(chan int, 10)
	// selectCh := make(chan int)
	// closeCh := make(chan int)
	go producer(ch)
	go consumer(ch)

	timer := time.NewTimer(time.Second * 10)

	ch2 := make(chan int)

	select {
	case <-timer.C:
		fmt.Println("timeout form channel")
	case <-ch2:
		fmt.Println("chan2 closed")
	}

	defer close(ch)
	// lihaorong.GetApp()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()
	go process(ctx, 5000*time.Millisecond)
	<-ctx.Done()

	fmt.Println("main: ", ctx.Err())
	// <-selectCh

}
