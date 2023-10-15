//go:build !windows
// +build !windows

package single_line_print

import (
	"fmt"
)

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
