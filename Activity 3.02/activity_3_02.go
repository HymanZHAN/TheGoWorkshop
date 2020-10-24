package main

import (
	"errors"
	"fmt"
)

type applicant struct {
	name          string
	creditScore   int
	monthlyIncome float64
}

type loan struct {
	amount float64
	term   int
	rate   float64
}

type loanApplicationResult struct {
	monthlyPayment float64
	totalCost      float64
	approved       bool
}

func main() {
	u1 := applicant{
		name:          "1",
		creditScore:   500,
		monthlyIncome: 1000,
	}
	l1 := loan{
		amount: 1000,
		rate:   15,
		term:   24,
	}
	r1, err := calculateLoan(u1, l1)

	if err != nil {
		fmt.Print("Error:", err)
	} else {
		printResult(u1, l1, r1)
	}

	u2 := applicant{
		name:          "2",
		creditScore:   350,
		monthlyIncome: 1000,
	}
	l2 := loan{
		amount: 10000,
		rate:   20,
		term:   12,
	}
	r2, err := calculateLoan(u2, l2)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		printResult(u2, l2, r2)
	}
}

func calculateLoan(u applicant, l loan) (loanApplicationResult, error) {
	var totalCost float64
	var monthlyPayment float64
	var approved bool

	if l.term%12 != 0 {
		return loanApplicationResult{}, errors.New("Loan term not divisible by 12")
	}
	totalCost = l.amount * l.rate / 100
	monthlyPayment = (l.amount + totalCost) / float64(l.term)

	if (u.creditScore >= 450 && monthlyPayment/u.monthlyIncome > 0.2) || (u.creditScore < 450 && monthlyPayment/u.monthlyIncome > 0.1) {
		approved = false
	} else {
		approved = true
	}

	return loanApplicationResult{
		monthlyPayment: monthlyPayment,
		totalCost:      totalCost,
		approved:       approved,
	}, nil
}

func printResult(u applicant, l loan, r loanApplicationResult) {
	fmt.Println("Applicant", u.name)
	fmt.Println("-----------")
	fmt.Println("Credit Score    :", u.creditScore)
	fmt.Println("Income          :", u.monthlyIncome)
	fmt.Println("Loan Amount     :", l.amount)
	fmt.Println("Loan Term       :", l.term)
	fmt.Println("Monthly Payment :", r.monthlyPayment)
	fmt.Println("Rate            :", l.rate)
	fmt.Println("Total Cost      :", r.totalCost)
	fmt.Println("Approved        :", r.approved)
	fmt.Println()
}
