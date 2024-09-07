package factory_method

import "fmt"

type DebitCardPaymentMethod struct {
}

func (m *DebitCardPaymentMethod) Pay(amount float32) string {
	return fmt.Sprintf("%.2f payed by debit card", amount)
}
