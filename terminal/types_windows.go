//go:build windows
// +build windows

package terminal

import (
	"errors"
	"syscall"
)

// learn more details about these structs:
// https://learn.microsoft.com/en-us/windows/console/console-screen-buffer-info-str
//
//goland:noinspection SpellCheckingInspection
type coord struct {
	x int16
	y int16
}
type smallRect struct {
	left   int16
	top    int16
	right  int16
	bottom int16
}
type consoleScreenBufferInfo struct {
	size              coord
	cursorPosition    coord
	attributes        uint16
	window            smallRect
	maximumWindowSize coord
}

// define software sdk calls
var (
	kernel32 = syscall.NewLazyDLL("kernel32.dll")

	procGetConsoleMode             = kernel32.NewProc("GetConsoleMode")
	procSetConsoleMode             = kernel32.NewProc("SetConsoleMode")
	procSetConsoleCursorPosition   = kernel32.NewProc("SetConsoleCursorPosition")
	procGetConsoleScreenBufferInfo = kernel32.NewProc("GetConsoleScreenBufferInfo")
	procFillConsoleOutputCharacter = kernel32.NewProc("FillConsoleOutputCharacterW")
)

// errors
var (
	getModeError = errors.New("get console mode fail")
	setModeError = errors.New("set console mode fail")
)

// stdout flags
//
//goland:noinspection GoSnakeCaseUsage
const (
	ENABLE_VIRTUAL_TERMINAL_PROCESSING uint32 = 0x0004
)

// stdin flags
//
//goland:noinspection GoSnakeCaseUsage
const (
	ENABLE_ECHO_INPUT uint32 = 0x0004
	ENABLE_LINE_INPUT uint32 = 0x0002
)
