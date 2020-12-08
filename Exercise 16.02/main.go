package main

import (
	"log"
	"sync"
)

func main() {
	s1 := 0
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go sum(1, 100, wg, &s1)
	wg.Wait()
	log.Println(s1)
}

func sum(from, to int, wg *sync.WaitGroup, res *int) {
	*res = 0
	for i := from; i <= to; i++ {
		*res += i
	}

	wg.Done()
	return
}
