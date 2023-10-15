package main

import (
	"fmt"
	slp "github.com/lianggaoqiang/single-line-print"
	"time"
)

func main() {
	p := slp.NewPrinter()
	for i := 0; i <= 100; i++ {
		p.Print(fmt.Sprintf("> Downloading from remote: %d%%", i))
		time.Sleep(time.Millisecond * 3000)
	}
	p.Stop()
}
