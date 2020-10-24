package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(2 * time.Second)
	end := time.Now()

	elapsed := end.Sub(start).Seconds()

	fmt.Printf("The execution took exactly %v seconds!", elapsed)
}
