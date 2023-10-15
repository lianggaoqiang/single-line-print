//go:build windows
// +build windows

package terminal

import (
	"os"
	"unsafe"
)

// EnableVirtualTerminal could enable Virtual Terminal Sequence in Windows System
// learn more details about Virtual Terminal Sequence:
// https://learn.microsoft.com/en-us/windows/console/console-virtual-terminal-sequences
func EnableVirtualTerminal() error {
	var (
		mode uint32
		f    = os.Stdout
	)

	// get console mode
	res, _, _ := procGetConsoleMode.Call(f.Fd(), uintptr(unsafe.Pointer(&mode)))
	if res == 0 {
		return getModeError
	}

	// change mode and set console mode
	mode |= ENABLE_VIRTUAL_TERMINAL_PROCESSING
	res, _, _ = procSetConsoleMode.Call(f.Fd(), uintptr(mode))
	if res == 0 {
		return setModeError
	}

	return nil
}

// DisableVirtualTerminal is the inverse method to EnableVirtualTerminal
func DisableVirtualTerminal() error {
	var (
		mode uint32
		f    = os.Stdout
	)

	// get console mode
	res, _, _ := procGetConsoleMode.Call(f.Fd(), uintptr(unsafe.Pointer(&mode)))
	if res == 0 {
		return getModeError
	}

	// change mode and set console mode
	mode &= ^ENABLE_VIRTUAL_TERMINAL_PROCESSING
	res, _, _ = procSetConsoleMode.Call(f.Fd(), uintptr(mode))
	if res == 0 {
		return setModeError
	}

	return nil
}

// ChangeTerminalInputMode changes the mode of Stdin, it's responsible for following tasks:
// 1. hide the echo of Standard Input
// 2. disable the buffering function of Standard Input
func ChangeTerminalInputMode() error {
	var (
		mode uint32
		f    = os.Stdin
	)

	// get console mode
	res, _, _ := procGetConsoleMode.Call(f.Fd(), uintptr(unsafe.Pointer(&mode)))
	if res == 0 {
		return getModeError
	}

	// storage old mode
	oriModeState = modeStage{
		echo: mode&ENABLE_ECHO_INPUT == ENABLE_ECHO_INPUT,
		buf:  mode&ENABLE_LINE_INPUT == ENABLE_LINE_INPUT,
	}

	// change mode and set console mode
	mode &= ^(ENABLE_ECHO_INPUT | ENABLE_LINE_INPUT)
	res, _, _ = procSetConsoleMode.Call(f.Fd(), uintptr(mode))
	if res == 0 {
		return setModeError
	}

	go startScanning()
	scanSignal <- 1
	return nil
}

// RestoreTerminalInputMode is the inverse method to ChangeTerminalInputMode,
// and it's responsible for following tasks:
// 1. restore the echo of Standard Input
// 2. restore the buffering function of Standard Input
func RestoreTerminalInputMode() error {
	var (
		mode uint32
		f    = os.Stdin
	)

	// get console mode
	res, _, _ := procGetConsoleMode.Call(f.Fd(), uintptr(unsafe.Pointer(&mode)))
	if res == 0 {
		return getModeError
	}

	// restore console mode
	if oriModeState.echo {
		mode |= ENABLE_ECHO_INPUT
	}
	if oriModeState.buf {
		mode |= ENABLE_LINE_INPUT
	}

	// set console mode
	res, _, _ = procSetConsoleMode.Call(f.Fd(), uintptr(mode))
	if res == 0 {
		return setModeError
	}

	return nil
}
