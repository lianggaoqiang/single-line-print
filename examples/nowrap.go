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
		var suffix string
		if i != 100 {
			suffix = "..."
		} else {
			suffix = "Done!"
		}
		p.Print(fmt.Sprintf("> Downloading from remote: %d%%\n%s\n", i, suffix))
		time.Sleep(time.Millisecond * 10)
	}
	p.Stop()
}
