package main

import (
	"fmt"
	"log"
)

func main() {
	log.SetFlags(0)
	in, out := make(chan string), make(chan string)
	go readthem(in, out)
	strs := []string{"a", "b", "c", "d", "e", "f"}
	for _, s := range strs {
		in <- s
	}
	close(in)
	fmt.Println(<-out)
}

func readthem(in, out chan string) {
	for i := range in {
		log.Println(i)
	}
	// out <- "done"
}
