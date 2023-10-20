package main

import (
	slp "github.com/lianggaoqiang/single-line-print"
	"testing"
)

var panicMsg = "printer or writer is repeatedly instantiate but without panic"

// check duplicate instantiation no.1
func TestDuplicateInstantiation1(t *testing.T) {
	p := slp.NewPrinter()
	defer func() {
		if ok := recover(); ok == nil {
			t.Error(panicMsg)
		}
		p.Stop()
	}()
	slp.NewPrinter()
}

// check duplicate instantiation no.2
func TestDuplicateInstantiation2(t *testing.T) {
	p := slp.NewPrinter()
	defer func() {
		if ok := recover(); ok == nil {
			t.Error(panicMsg)
		}
		p.Stop()
	}()
	slp.NewPrinterWithFlag(0)
}

// check duplicate instantiation no.3
func TestDuplicateInstantiation3(t *testing.T) {
	p := slp.NewWriter()
	defer func() {
		if ok := recover(); ok == nil {
			t.Error(panicMsg)
		}
		p.Stop()
	}()
	slp.NewPrinter()
}

// check duplicate instantiation no.4
func TestDuplicateInstantiation4(t *testing.T) {
	p := slp.NewPrinterWithFlag(0)
	defer func() {
		if ok := recover(); ok == nil {
			t.Error(panicMsg)
		}
		p.Stop()
	}()
	slp.NewWriterWithFlag(0)
}
