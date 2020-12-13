package main

import (
	"context"
	"log"
	"time"
)

func main() {
	r := make(chan int)
	c := context.TODO()
	cl, stop := context.WithCancel(c)
	go counterNumbers(cl, r)

	go func() {
		time.Sleep(time.Millisecond * 5 * 100)
		stop()
	}()

	v := <-r
	log.Println(v)
}

func counterNumbers(c context.Context, r chan int) {
	v := 0
	for {
		select {
		case <-c.Done():
			r <- v
		default:
			time.Sleep(time.Millisecond * 100)
			v++
		}
	}
}
