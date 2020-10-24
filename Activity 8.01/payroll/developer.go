package payroll

import (
	"errors"
	"fmt"
)

type Developer struct {
	Individual        Employee
	HourlyRate        float64
	HoursWorkedInYear float64
	Review            map[string]interface{}
}

func ReviewDeveloper(d Developer) {
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

func (d Developer) pay() (string, float64) {
	name := d.Individual.FirstName + " " + d.Individual.LastName
	yearlyPay := d.HourlyRate * d.HoursWorkedInYear
	return name, yearlyPay
}
