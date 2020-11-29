package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool)

	signal.Notify(sigs, syscall.SIGINT)

	go func() {
		for {
			s := <-sigs
			switch s {
			case syscall.SIGINT:
				fmt.Println()
				fmt.Println("My process has been interrupted.")
				fmt.Println("Some clean up is occurring.")
				done <- true
			}
		}
	}()

	fmt.Println("Waiting for SIGINT...")
	<-done

	fmt.Println("Out of here!")
}
