package main

import (
	"errors"
	"fmt"
)

type directDeposit struct {
	lastName      string
	firstName     string
	bankName      string
	routingNumber int
	accountNumber int
}

var (
	ErrInvalidRoutingNum = errors.New("invalid routing number")
	ErrInvalidLastName   = errors.New("invalid last name")
)

func (d directDeposit) validateRoutingNumber() error {
	if d.routingNumber < 100 {
		return ErrInvalidRoutingNum
	}
	return nil
}

func (d directDeposit) validateLastName() error {
	if d.lastName == "" {
		return ErrInvalidLastName
	}
	return nil
}

func (d directDeposit) report() {
	fmt.Println("*************************")
	fmt.Println("Last Name:", d.lastName)
	fmt.Println("First Name:", d.firstName)
	fmt.Println("Bank Name", d.bankName)
	fmt.Println("Routing Number", d.routingNumber)
	fmt.Println("Account Number", d.accountNumber)
}

func main() {
	deposit := directDeposit{
		lastName:      "",
		firstName:     "Abe",
		bankName:      "XYZ Inc",
		routingNumber: 17,
		accountNumber: 1809,
	}
	fmt.Println(deposit.validateRoutingNumber())
	fmt.Println(deposit.validateLastName())
	deposit.report()
}
