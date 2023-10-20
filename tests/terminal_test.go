package main

import (
	"github.com/lianggaoqiang/single-line-print/terminal"
	"testing"
)

func TestGetTerminalSize(t *testing.T) {
	size := terminal.GetSize()
	if size.Width == 0 || size.Height == 0 {
		t.Error("can not get the size of terminal")
	}
}
