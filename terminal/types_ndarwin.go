//go:build (unix || linux) && !darwin
// +build unix linux
// +build !darwin

package terminal

import "syscall"

type lFlag = uint32

// define syscall number of getting and setting attributes of terminal
//
//goland:noinspection GoSnakeCaseUsage
const (
	SYS_IOCTL_GET = syscall.TCGETS
	SYS_IOCTL_SET = syscall.TCSETS
)
