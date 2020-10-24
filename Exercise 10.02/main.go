package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	Start := time.Now()
	time.Sleep(2 * time.Second)
	End := time.Now()
	fmt.Println(elapsedTime(Start, End))
}

func whatstheclock() string {
	return time.Now().Format(time.ANSIC)
}

func elapsedTime(start time.Time, end time.Time) string {
	elapsed := end.Sub(start)
	hours := strconv.Itoa(int(elapsed.Hours()))
	mins := strconv.Itoa(int(elapsed.Minutes()))
	seconds := strconv.Itoa(int(elapsed.Seconds()))

	return fmt.Sprintf(
		"The total execution time elapsed is: %v hour(s) and %v minute(s) and %v second(s)!",
		hours, mins, seconds)
}
