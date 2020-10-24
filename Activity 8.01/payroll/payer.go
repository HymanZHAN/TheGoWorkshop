package payroll

import "fmt"

type payer interface {
	pay() (string, float64)
}

func PayDetails(p payer) {
	fullName, yearlyPay := p.pay()
	fmt.Printf("%s got pad %.2f for the year\n", fullName, yearlyPay)
}
