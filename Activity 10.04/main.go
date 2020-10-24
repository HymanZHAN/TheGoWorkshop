package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("The current time:", start.Format(time.ANSIC))

	duration := time.Duration(6*time.Hour + 6*time.Minute + 6*time.Second)
	future := start.Add(duration)
	fmt.Println("6 hours, 6 minutes and 6 seconds from now will be:", future.Format(time.ANSIC))
}
