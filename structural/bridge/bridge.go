package bridge

import (
	"errors"
	"fmt"
	"io"
)

// ////////////////////////////
// Implementations of the PrintApi interface
// ////////////////////////////

type PrinterApi interface {
	Print(msg string) error
}

// PrinterImpl1 is the first implementation
type PrinterImpl1 struct{}

func (p *PrinterImpl1) Print(msg string) error {
	fmt.Println("PrinterImpl1 prints ", msg)
	return nil
}

// PrinterImpl2 is the second implementation
type PrinterImpl2 struct {
	Writer io.Writer
}

func (p *PrinterImpl2) Print(msg string) error {
	if p.Writer == nil {
		return errors.New("you need to pass an io.Writer to PrinterImpl2")
	}
	_, err := fmt.Fprintln(p.Writer, msg)
	if err != nil {
		return err
	}
	return nil
}

// ////////////////////////////
// Abstractions
//
// These use the two printApi implementations to do the printing
// The key is that the implementations are interchangeable as far
// as NormalPrinter and PacktPrinter are concerned
// ////////////////////////////

// NormalPrinter contains a msg and printer
// It uses the printer to print its message
type NormalPrinter struct {
	Msg     string
	Printer PrinterApi
}

// Print prints the message using a NormalPrinter
func (p *NormalPrinter) Print() error {
	err := p.Printer.Print(p.Msg)
	if err != nil {
		return err
	}
	return nil
}

// PacktPrinter contains a msg and printer
// It uses the printer to print its message
type PacktPrinter struct {
	Msg     string
	Printer PrinterApi
}

// Print prints the message using a PacktPrinter
func (p *PacktPrinter) Print() error {
	err := p.Printer.Print(fmt.Sprintf("From Packt: %s ", p.Msg))
	if err != nil {
		return err
	}
	return nil
}
