//go:build unix || darwin || linux
// +build unix darwin linux

package terminal

import "errors"

// errors
var (
	getModeError = errors.New("get terminal mode fail")
	setModeError = errors.New("set terminal mode fail")
)

// EnableVirtualTerminal exists solely for compatibility with mode_windows.go
func EnableVirtualTerminal() error {
	return errors.New("EnableVirtualTerminal: this method is only defined under Windows System")
}
