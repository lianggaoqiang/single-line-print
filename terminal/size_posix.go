//go:build !windows
// +build !windows

package terminal

import (
	"flag"
	"os"
	"runtime"
	"syscall"
	"unsafe"
)

// GetSize could get the size of the terminal window
func GetSize() (windowSize size) {
	// if in testing, return virtual size directly
	flag.Parse()
	if *testEnv == "GithubWorkflow" {
		return size{110, 30}
	}

	var (
		f   *os.File
		err error
	)

	// get the descriptor of terminal file
	if runtime.GOOS == "openbsd" {
		f, err = os.OpenFile("/dev/tty", os.O_RDWR, 0)
		if err != nil {
			return
		}
	} else {
		f, err = os.OpenFile("/dev/tty", os.O_WRONLY, 0)
		if err != nil {
			return
		}
	}
	defer f.Close()

	// system call
	syscall.Syscall(
		syscall.SYS_IOCTL,
		f.Fd(),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(&windowSize)),
	)

	return windowSize
}
