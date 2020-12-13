package main

import (
	"fmt"
	"log"
)

func main() {
	ch := make(chan string)
	ch <- "Hello John"
	go greet(ch)
	log.Println(<-ch)
	log.Println(<-ch)
}

func greet(ch chan string) {
	msg := <-ch
	ch <- fmt.Sprintf("Thanks for %s", msg)
	ch <- "Hello David"
}
