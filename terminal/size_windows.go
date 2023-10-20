//go:build windows
// +build windows

package terminal

import (
	"os"
	"unsafe"
)

// GetSize could get the size of the terminal window
func GetSize() (windowSize size) {
	// if in testing, return virtual size directly
	flag.Parse()
	if *testEnv == "GithubWorkflow" {
		return size{110, 30}
	}

	// in Windows system, another approach to acquire a handle for the terminal file
	// is utilizing os.OpenFile("CONOUT$") or os.OpenFile("CONIN$")
	f := os.Stdout

	var ci consoleScreenBufferInfo
	res, _, _ := procGetConsoleScreenBufferInfo.Call(f.Fd(), uintptr(unsafe.Pointer(&ci)))
	if res == 0 {
		return
	}
	windowSize.Width = uint16(ci.window.right - ci.window.left + 1)
	windowSize.Height = uint16(ci.window.bottom - ci.window.top + 1)

	return windowSize
}
