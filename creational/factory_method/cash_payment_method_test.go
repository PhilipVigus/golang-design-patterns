package factory_method

import "testing"

func TestCashPaymentMethod_Pay(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)

	if err != nil {
		t.Fatal(err)
	}

	if payment == nil {
		t.Fatal("payment method is nil")
	}

	paymentResponse := payment.Pay(20)

	if paymentResponse != "20.00 payed by cash" {
		t.Error("payment response is wrong")
	}
}
