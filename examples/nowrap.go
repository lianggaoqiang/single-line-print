package main

import (
	"fmt"
	slp "github.com/lianggaoqiang/single-line-print"
	"strings"
	"time"
)

func main() {
	p := slp.NewPrinter()
	for i := 0; i <= 100; i++ {
		p.Print(fmt.Sprintf("> Downloading from remote: %d%%", i))
		time.Sleep(time.Millisecond * 10)
	}

	fmt.Println("\n", strings.Repeat("-", 100))

	p.Reload()
	for i := 0; i <= 100; i++ {
		if i != 100 {
			p.Print(fmt.Sprintf("> Downloading from remote: %d%%\n...\n", i))
		} else {
			p.Print(fmt.Sprintf("> Downloading from remote: %d%%\nDone!\n", i))
		}
		time.Sleep(time.Millisecond * 10)
	}
	p.Stop()
}
