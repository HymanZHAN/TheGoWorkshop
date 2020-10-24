package payroll

type Manager struct {
	Individual     Employee
	Salary         float64
	CommissionRate float64
}

func (m Manager) pay() (string, float64) {
	name := m.Individual.FirstName + " " + m.Individual.LastName
	yearlyPay := m.CommissionRate*m.Salary + m.Salary
	return name, yearlyPay
}
