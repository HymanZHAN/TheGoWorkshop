package main

import (
	"errors"
	"fmt"
)

type Payer interface {
	Pay() (string, float64)
}

type Employee struct {
	Id        int
	FirstName string
	LastName  string
}

type Developer struct {
	Individual        Employee
	HourlyRate        float64
	HoursWorkedInYear float64
	Review            map[string]interface{}
}

func (d Developer) Pay() (string, float64) {
	name := d.Individual.FirstName + " " + d.Individual.LastName
	yearlyPay := d.HourlyRate * d.HoursWorkedInYear
	return name, yearlyPay
}

type Manager struct {
	Individual     Employee
	Salary         float64
	CommissionRate float64
}

func (m Manager) Pay() (string, float64) {
	name := m.Individual.FirstName + " " + m.Individual.LastName
	yearlyPay := m.CommissionRate*m.Salary + m.Salary
	return name, yearlyPay
}

func main() {
	eric := Developer{
		Individual:        Employee{Id: 1, FirstName: "Eric", LastName: "Davis"},
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
	reviewDeveloper(eric)
	payDetails(eric)

	boss := Manager{
		Individual:     Employee{Id: 2, FirstName: "Mr.", LastName: "Boss"},
		CommissionRate: 0.5,
		Salary:         107000,
	}

	payDetails(boss)

	// Eric Davis got a review rating of 2.80
	// Eric Davis got paid 84000.00 for the year
	// Mr. Boss got paid 160500.00 for the year

}

func payDetails(p Payer) {
	fullName, yearlyPay := p.Pay()
	fmt.Printf("%s got pad %.2f for the year\n", fullName, yearlyPay)
}

func reviewDeveloper(d Developer) {
	var reviewPoint float64
	for _, review := range d.Review {
		switch v := review.(type) {
		case string:
			if v == "Excellent" {
				reviewPoint += 5
			} else if v == "Good" {
				reviewPoint += 4
			} else if v == "Fair" {
				reviewPoint += 3
			} else if v == "Poor" {
				reviewPoint += 2
			} else if v == "Unsatisfactory" {
				reviewPoint++
			} else {
				panic(errors.New("invalid review"))
			}
		case int:
			if v > 5 || v < 1 {
				panic(errors.New("invalid review score"))
			}
			reviewPoint += float64(v)
		case float64:
			if v > 5 || v < 1 {
				panic(errors.New("invalid review score"))
			}
			reviewPoint += v
		}
	}
	var averageReviewPoint float64
	if len(d.Review) == 0 {
		averageReviewPoint = 0
	} else {
		averageReviewPoint = reviewPoint / float64(len(d.Review))
	}
	fmt.Printf(
		"%s %s got a review rating of %.2f\n",
		d.Individual.FirstName, d.Individual.LastName, averageReviewPoint,
	)
}
