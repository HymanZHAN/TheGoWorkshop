package main

import "fmt"

type employee struct {
	id        int
	firstName string
	lastName  string
}

type developer struct {
	Individual employee
	HourlyRate int
	WorkWeek   [7]int
}

type weekday int

const (
	monday weekday = iota
	tuesday
	wednesday
	thursday
	friday
	saturday
)

func main() {
	dev := developer{
		Individual: employee{id: 1, firstName: "Xucong", lastName: "Zhan"},
		HourlyRate: 22,
		WorkWeek:   [7]int{},
	}

	dev.LogHours(monday, 8)
	dev.LogHours(tuesday, 10)

	fmt.Println(dev.HoursWorked())
}

func (d developer) HoursWorked() string {
	var result string
	var totalHours int
	week := []string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
	}
	for i, hour := range d.WorkWeek {
		if hour > 0 {
			result += fmt.Sprintf("Hours worked on %s: %d\n", week[i], hour)
			totalHours += hour
		}
	}
	result += fmt.Sprintf("Hours worked this week: %d", totalHours)
	return result
}

func (d *developer) LogHours(day weekday, hours int) {
	d.WorkWeek[day] = hours
}
