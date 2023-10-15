package main

import (
	"bufio"
	"fmt"
	slp "github.com/lianggaoqiang/single-line-print"
	"time"
)

func main() {
	w := slp.NewWriter()
	bw := bufio.NewWriter(w)
	for i := 0; i <= 100; i++ {
		s := fmt.Sprintf("now the number is %d!\n", i)
		bw.Write([]byte(s))
		bw.Flush()
		time.Sleep(time.Millisecond * 10)
	}
	w.Stop()
}
