package factory_method

import (
	"errors"
	"fmt"
)

type PaymentMethodType int

type PaymentMethod interface {
	Pay(amount float32) string
}

const (
	Cash PaymentMethodType = iota
	DebitCard
)

func GetPaymentMethod(paymentMethod PaymentMethodType) (PaymentMethod, error) {
	switch paymentMethod {
	case Cash:
		return new(CashPaymentMethod), nil
	case DebitCard:
		return new(DebitCardPaymentMethod), nil
	default:
		return nil, errors.New(fmt.Sprintf("Payment method not supported: %v", paymentMethod))
	}
}
