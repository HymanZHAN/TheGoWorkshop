package main

import (
	"fmt"
	"time"
)

func main() {
	localTime := time.Now()
	laTimezone, _ := time.LoadLocation("America/Los_Angeles")
	laTime := localTime.In(laTimezone)

	nyTimezone, _ := time.LoadLocation("America/New_York")
	nyTime := localTime.In(nyTimezone)

	fmt.Println("The local current time is:", localTime.Format(time.ANSIC))
	fmt.Println("The time in New York is:", nyTime.Format(time.ANSIC))
	fmt.Println("The time in Los Angeles is:", laTime.Format(time.ANSIC))

}
