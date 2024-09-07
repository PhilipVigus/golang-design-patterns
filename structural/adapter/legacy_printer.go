package adapter

import "fmt"

type LegacyPrinter interface {
	Print(s string) string
}

type MyLegacyPrinter struct{}

func (p *MyLegacyPrinter) Print(s string) string {
	msg := fmt.Sprintf("Legacy Printer: %s", s)
	println(msg)
	return msg
}
