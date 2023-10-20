package util

import "fmt"

// ESC generates multiple ANSI control characters
func ESC(suffix ...string) (ansis string) {
	for _, s := range suffix {
		ansis += fmt.Sprintf("%c[%s", 033, s)
	}
	return
}
