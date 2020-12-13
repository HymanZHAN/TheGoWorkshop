package main

import (
	"fmt"
	"sync"
)

func main() {
	lock := sync.Mutex{}
	result := ""
	wg := sync.WaitGroup{}
	wg.Add(4)
	go sumAsStr(1, 25, &result, &lock, &wg)
	go sumAsStr(26, 50, &result, &lock, &wg)
	go sumAsStr(52, 75, &result, &lock, &wg)
	go sumAsStr(76, 100, &result, &lock, &wg)
	wg.Wait()

	fmt.Println(result)
}

func sumAsStr(from, to int, s *string, m *sync.Mutex, wg *sync.WaitGroup) {
	for i := from; i <= to; i++ {
		m.Lock()
		newStr := fmt.Sprintf("|%d|", i)
		*s = fmt.Sprintf("%s%s", *s, newStr)
		m.Unlock()
	}
	wg.Done()
}
