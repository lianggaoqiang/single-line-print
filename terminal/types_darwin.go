//go:build darwin
// +build darwin

package terminal

import "syscall"

type lFlag = uint64

// define syscall number of getting and setting terminal attributes
//
//goland:noinspection GoSnakeCaseUsage
const (
	SYS_IOCTL_GET = syscall.TIOCGETA
	SYS_IOCTL_SET = syscall.TIOCSETA
)
