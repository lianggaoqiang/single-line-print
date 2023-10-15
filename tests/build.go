package main

import slp "github.com/lianggaoqiang/single-line-print"

// This file is used for testing the build process of this program
func main() {
	// normal printer
	p := slp.NewPrinter()
	for i := 0; i < 1000; i++ {
		p.Print("")
	}
	p.Stop()

	// printer with flag
	p = slp.NewPrinterWithFlag(slp.DisableInput | slp.HideCursor | slp.ResizeReactively)
	for i := 0; i < 1000; i++ {
		p.Print("")
	}
	p.Reload()
	p.Stop()
}
