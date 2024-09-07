package adapter

import "testing"

func TestPrinterAdaptor(t *testing.T) {
	msg := "Hello World"

	adapter := PrinterAdaptor{OldPrinter: &MyLegacyPrinter{}, Msg: msg}

	returnedMsg := adapter.PrintStored()

	if returnedMsg != "Legacy Printer: Adapter: Hello World" {
		t.Errorf("Message didn't match :%s\n", returnedMsg)
	}
}
