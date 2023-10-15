//go:build unix || darwin || linux
// +build unix darwin linux

package terminal

import (
	"os"
	"syscall"
	"unsafe"
)

// ChangeTerminalInputMode changes the mode of Stdin, it's responsible for following tasks:
// 1. hide the echo of Standard Input
// 2. disable the buffering function of Standard Input
func ChangeTerminalInputMode() error {
	var (
		f          = os.Stdin
		terminalOs = syscall.Termios{}
	)

	// get terminal mode
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, f.Fd(), SYS_IOCTL_GET, uintptr(unsafe.Pointer(&terminalOs)))
	if errno != 0 {
		return getModeError
	} else {
		// store old mode
		mode := terminalOs.Lflag
		oriModeState = modeStage{
			echo: mode&syscall.ECHO == syscall.ECHO,
			buf:  mode&syscall.ICANON == syscall.ICANON,
		}

		// change and set terminal mode
		flag := ^(syscall.ECHO | syscall.ICANON)
		terminalOs.Lflag &= lFlag(flag)
		_, _, errno = syscall.Syscall(syscall.SYS_IOCTL, f.Fd(), SYS_IOCTL_SET, uintptr(unsafe.Pointer(&terminalOs)))
		if errno != 0 {
			return setModeError
		}
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
		f          = os.Stdin
		terminalOs = syscall.Termios{}
	)

	// get terminal mode
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, f.Fd(), SYS_IOCTL_GET, uintptr(unsafe.Pointer(&terminalOs)))
	if errno != 0 {
		return getModeError
	} else {
		// restore terminal mode
		mode := terminalOs.Lflag
		if oriModeState.echo {
			mode |= syscall.ECHO
		}
		if oriModeState.buf {
			mode |= syscall.ICANON
		}

		// change and set terminal mode
		terminalOs.Lflag = mode
		_, _, errno = syscall.Syscall(syscall.SYS_IOCTL, f.Fd(), SYS_IOCTL_SET, uintptr(unsafe.Pointer(&terminalOs)))
		if errno != 0 {
			return setModeError
		}
	}

	return nil
}
