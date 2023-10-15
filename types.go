package single_line_print

import (
	"fmt"
	"github.com/lianggaoqiang/single-line-print/terminal"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// struct of printer and writer
type ins struct {
	lineCount int
	termWidth int
	closed    bool
	noPrint   bool
	mode      uint8
	kind      string
}

type printer ins
type writer ins
type insPointer = *ins
type printerPointer = *printer

// define package state
var (
	kind       string
	active     = false
	mtx        = sync.Mutex{}
	stopSignal = make(chan bool)
)

// define the flags of mode
const (
	HideCursor       uint8 = 0x01 // if set, the cursor will be hidden during printing or writing
	DisableInput     uint8 = 0x02 // if set, input will be disabled during printing or writing
	ResizeReactively uint8 = 0x04 // if set, terminal window size will be got before each printing or writing
)

// generate new instance of ins
func defaultIns(k string) ins {
	pkgCheck(k)
	defaultMode := HideCursor | DisableInput
	setMode(defaultMode)
	return ins{
		kind:      k,
		noPrint:   true,
		mode:      defaultMode,
		termWidth: int(terminal.GetSize().Width),
	}
}

// generate new instance of ins with flag
func defaultInsWithFlag(k string, f uint8) ins {
	pkgCheck(k)
	setMode(f)
	return ins{
		kind:      k,
		mode:      f,
		noPrint:   true,
		termWidth: int(terminal.GetSize().Width),
	}
}

// check package state
func pkgCheck(k string) {
	mtx.Lock()
	defer mtx.Unlock()
	if active {
		panic(fmt.Sprintf("a %s is actived, please use Stop to close it before calling this method", kind))
	} else {
		kind = k
		active = true
	}
	go listen()
}

// set instance mode
func setMode(f uint8) {
	if f&HideCursor == HideCursor {
		os.Stdout.Write([]byte(esc("?25l")))
	}
	if f&DisableInput == DisableInput {
		terminal.ChangeTerminalInputMode()
	}
}

// Generate multiple ANSI control characters
func esc(suffix ...string) (ansis string) {
	for _, s := range suffix {
		ansis += fmt.Sprintf("%c[%s", 033, s)
	}
	return
}

// count the number of message lines to be printed in the terminal
func (i *ins) countLine(s string) {
	if i.mode&ResizeReactively == ResizeReactively {
		i.termWidth = int(terminal.GetSize().Width)
	}
	l, r := 0, 0
	for _, c := range []byte(s) {
		if c == '\n' {
			l++
			r = 0
			continue
		}
		if r == i.termWidth {
			l++
			r = 0
			continue
		}
		r++
	}
	i.lineCount = l
}

// Stop is the implement of Stop in print.go and writer.go
func (i *ins) Stop() {
	i.istCheck()
	mtx.Lock()
	defer mtx.Unlock()
	active = false
	i.closed = true
	restoreTerminalMode(i.mode)
	stopSignal <- true
}
func restoreTerminalMode(mode uint8) {
	if mode&HideCursor == HideCursor {
		os.Stdout.Write([]byte(esc("?25h")))
	}
	if mode&DisableInput == DisableInput {
		terminal.RestoreTerminalInputMode()
	}
}

// Reload is the implement of Reload in print.go and writer.go
func (i *ins) Reload() {
	i.istCheck()
	mtx.Lock()
	defer mtx.Unlock()
	i.noPrint = true
}

// check instance state
func (i *ins) istCheck() {
	mtx.Lock()
	defer mtx.Unlock()
	if i.closed {
		panic(fmt.Sprintf("%s is already in closed state", i.kind))
	}
}

// program listening
func listen() {
	osc := make(chan os.Signal, 1)
	signal.Notify(osc, syscall.SIGTERM, syscall.SIGINT)
	select {
	case <-osc:
		terminal.RestoreTerminalInputMode()
		os.Stdout.Write([]byte(esc("?25h")))
		os.Exit(0)
	case <-stopSignal:
		return
	}
}
