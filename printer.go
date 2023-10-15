package single_line_print

import (
	"fmt"
	"os"
	"strconv"
)

// NewPrinter returns a new instance of printer
func NewPrinter() *printer {
	p := printer(defaultIns("printer"))
	return &p
}

// NewPrinterWithFlag returns a new instance of printer with flag
func NewPrinterWithFlag(f uint8) *printer {
	p := printer(defaultInsWithFlag("printer", f))
	return &p
}

// Print outputs strings as a single line to the terminal
func (pt *printer) Print(s string) (n int, err error) {
	// convert *pt to *ins, and run instance checking
	asIns := insPointer(pt)
	asIns.istCheck()

	// ANSI control character that clean pt.lineCount lines
	if pt.noPrint {
		pt.noPrint = false
	} else {
		transLeftAnsi := strconv.Itoa(pt.cursorOffset) + "D"
		if pt.lineCount == 0 {
			fmt.Fprint(os.Stdout, esc(transLeftAnsi, "0K"))
		} else {
			for i := 0; i < pt.lineCount; i++ {
				fmt.Fprint(os.Stdout, esc("1A", "0K"))
			}
		}
	}
	asIns.countLine(s)

	return fmt.Fprint(os.Stdout, s)
}

// Reload resets the state of an existing printer or writer instance
func (pt *printer) Reload() {
	insPointer(pt).Reload()
}

// Stop will set printer or writer to closed state, and eliminate the impact of flags on terminal
func (pt *printer) Stop() {
	insPointer(pt).Stop()
}
