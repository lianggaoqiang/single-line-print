package terminal

import (
	"os"
)

// struct of the terminal window size
type size struct {
	Height, Width uint16
}

// struct which storage the terminal modes temporarily before each change
// when ChangeTerminalInputMode is called, set corresponding property base on old terminal mode flag
// when RestoreTerminalInputMode is called, restore terminal modes base on the properties of this struct
type modeStage struct {
	echo, buf bool
}

var (
	oriModeState = modeStage{}    // instance of modeStage
	scanSignal   = make(chan int) // signal of scanning from terminal, finish scanning when it is 2
)

// When ChangeTerminalInputMode is called, start scanning to disable input
func startScanning() {
	for {
		code := <-scanSignal
		if code == 1 {
			abandon := make([]byte, 1)
			os.Stdin.Read(abandon)
			go func() {
				scanSignal <- 1
			}()
		} else {
			return
		}
	}
}
