package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(timeDiff("Asia/Taipei"))
}

func timeDiff(timezone string) (string, string) {
	current := time.Now()
	remoteZone, _ := time.LoadLocation(timezone)
	remoteTime := current.In(remoteZone)
	fmt.Println("the current time is:", current.Format(time.ANSIC))
	fmt.Println("The timezone:", timezone, "time is:", remoteTime)
	return current.Format(time.ANSIC), remoteTime.Format(time.ANSIC)
}
