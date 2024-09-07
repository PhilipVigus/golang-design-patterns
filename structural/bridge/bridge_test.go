package bridge

import (
	"errors"
	"strings"
	"testing"
)

type TestWriter struct {
	Msg string
}

func (t *TestWriter) Write(msg []byte) (n int, err error) {
	n = len(msg)
	if n > 0 {
		t.Msg = string(msg)
		return n, nil
	}
	err = errors.New("content received on Writer is empty")
	return
}

func TestPrinterImpl1_Print(t *testing.T) {
	api1 := PrinterImpl1{}

	err := api1.Print("Hello World")
	if err != nil {
		t.Errorf("Error while printing with API 1: %v", err)
	}
}

func TestPrinterImpl2_Print(t *testing.T) {
	api2 := PrinterImpl2{}

	err := api2.Print("Hello")
	if err != nil {
		expectedErrorMessage := "you need to pass an io.Writer to PrinterImpl2"
		if !strings.Contains(err.Error(), expectedErrorMessage) {
			t.Errorf("Error message was not correct.\nActual: %s\nExpected: %s\n", err.Error(), expectedErrorMessage)
		}
	}

	testWriter := TestWriter{}
	api2 = PrinterImpl2{Writer: &testWriter}
	expectedMessage := "Hello\n"
	err = api2.Print("Hello")
	if err != nil {
		t.Errorf("Error while printing with API 2: %v", err.Error())
	}

	if testWriter.Msg != expectedMessage {
		t.Errorf("Error message was not correct. Actual: %s Expected: %s", testWriter.Msg, expectedMessage)
	}
}

func TestNormalPrinter_Print(t *testing.T) {
	expectedMessage := "Hello io.Writer\n"

	normal := NormalPrinter{
		Msg:     "Hello io.Writer",
		Printer: &PrinterImpl1{},
	}

	err := normal.Print()
	if err != nil {
		t.Errorf(err.Error())
	}
	testWriter := TestWriter{}
	normal = NormalPrinter{
		Msg: "Hello io.Writer",
		Printer: &PrinterImpl2{
			Writer: &testWriter,
		},
	}

	err = normal.Print()
	if err != nil {
		t.Error(err.Error())
	}

	if testWriter.Msg != expectedMessage {
		t.Errorf("The expected message on the io.Writer doesn't match actual. Actual: %s Expected: %s", testWriter.Msg, expectedMessage)
	}
}

func TestPacktPrinter_Print(t *testing.T) {
	expectedMessage := "From Packt: Hello io.Writer \n"

	packt := PacktPrinter{
		Msg:     "Hello io.Writer",
		Printer: &PrinterImpl1{},
	}

	err := packt.Print()
	if err != nil {
		t.Errorf(err.Error())
	}

	testWriter := TestWriter{}
	packt = PacktPrinter{
		Msg: "Hello io.Writer",
		Printer: &PrinterImpl2{
			Writer: &testWriter,
		},
	}

	err = packt.Print()
	if err != nil {
		t.Error(err.Error())
	}

	if testWriter.Msg != expectedMessage {
		t.Errorf("The expected message on the io.Writer doesn't match actual.\n  Actual: %s\nExpected: %s\n", testWriter.Msg, expectedMessage)
	}
}
