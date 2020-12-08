package main

import "fmt"

func main() {
	a := 0
	go next(&a)
	go next(&a)
	go next(&a)
	fmt.Println(a)
}

func next(v *int) {
	fmt.Printf("Inside next: %d", *v)
	c := *v
	*v = c + 1
}
