package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	min := now.Minute()
	sec := now.Second()

	timeString := fmt.Sprintf(
		"%02d:%02d:%02d %d/%2d/%02d",
		hour, min, sec, year, month, day,
	)

	fmt.Println(timeString)
}
