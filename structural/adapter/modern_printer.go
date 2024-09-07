package adapter

import "fmt"

type ModernPrinter interface {
	PrintStored() string
}

type PrinterAdaptor struct {
	OldPrinter LegacyPrinter
	Msg        string
}

func (p *PrinterAdaptor) PrintStored() (newMsg string) {
	if p.OldPrinter != nil {
		newMsg = fmt.Sprintf("Adapter: %s", p.Msg)
		newMsg = p.OldPrinter.Print(newMsg)
	} else {
		newMsg = p.Msg
	}
	return
}
