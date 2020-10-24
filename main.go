package main

import (
	"fmt"
	"time"
)

func main() {
	t1, _ := time.Parse(time.RFC3339, "2019-09-27T22:18:11+00:00")
	t2, _ := time.Parse(time.UnixDate, "2019-09-27T22:18:11+00:00")
	t3, _ := time.Parse(time.ANSIC, "2019-09-27T22:18:11+00:00")
	fmt.Println("RFC3339:", t1)
	fmt.Println("UnixDate", t2)
	fmt.Println("ANSIC", t3)

	date := time.Date(2019, 10, 1, 9, 0, 0, 0, time.UTC)
	fmt.Println(date)
	nextDate := date.AddDate(1, 2, 3)
	fmt.Println(nextDate)

	Current := time.Now()
	Berlin, _ := time.LoadLocation("America/Los_Angeles")
	fmt.Println("The local current time is:", Current.Format(time.ANSIC))
	fmt.Println("The time in Berlin is: ", Current.In(Berlin).Format(time.ANSIC))
}
