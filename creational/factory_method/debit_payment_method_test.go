package factory_method

import "testing"

func TestDebitPaymentMethod_Pay(t *testing.T) {
	payment, err := GetPaymentMethod(DebitCard)

	if err != nil {
		t.Fatal(err)
	}

	if payment == nil {
		t.Fatal("payment method is nil")
	}

	paymentResponse := payment.Pay(30)

	if paymentResponse != "30.00 payed by debit card" {
		t.Error("payment response is wrong")
	}
}
