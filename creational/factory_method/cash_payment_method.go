package factory_method

import "fmt"

type CashPaymentMethod struct {
}

func (m *CashPaymentMethod) Pay(amount float32) string {
	return fmt.Sprintf("%.2f payed by cash", amount)
}
