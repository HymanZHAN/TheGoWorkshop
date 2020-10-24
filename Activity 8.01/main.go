package main

import (
	"activity/payroll"
	"fmt"
)

var eric payroll.Developer
var boss payroll.Manager

func init() {
	fmt.Println("Initializing variables")
	eric = payroll.Developer{
		Individual:        payroll.Employee{Id: 1, FirstName: "Eric", LastName: "Davis"},
		HourlyRate:        40,
		HoursWorkedInYear: 2100,
		Review: map[string]interface{}{
			"Work Quality":  5,
			"Work Quantity": 2,
			"Teamwork":      2,
			"Communication": 3,
			"Commitment":    2,
		},
	}

	boss = payroll.Manager{
		Individual:     payroll.Employee{Id: 2, FirstName: "Mr.", LastName: "Boss"},
		CommissionRate: 0.5,
		Salary:         107000,
	}
}

func main() {
	payroll.ReviewDeveloper(eric)
	payroll.PayDetails(eric)
	payroll.PayDetails(boss)
}
