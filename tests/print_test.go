package main

import (
	"fmt"
	slp "github.com/lianggaoqiang/single-line-print"
	"github.com/lianggaoqiang/single-line-print/terminal"
	"github.com/lianggaoqiang/single-line-print/util"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

var (
	filePath, expect string
	file             *os.File
	oldStdout        = os.Stdout
)

func init() {
	// get the path of "output" file
	curPath, err := os.Getwd()
	if err != nil {
		panic("get current path fail: " + err.Error())
	}
	filePath = filepath.Join(curPath, "output")
}

// check one line that with line break
func TestOneLineWithWrap(t *testing.T) {
	createFile()
	p := slp.NewPrinter()
	for i := 0; i <= 100; i++ {
		s := strconv.Itoa(i) + "\n"
		p.Print(s)
		if i == 0 {
			expect += esc("?25l") + s
		} else {
			expect += fmt.Sprintf("%s%d\n", esc("1A", "0K"), i)
		}
	}
	if readFile() != expect {
		t.Error("result does not match the expectation")
	}
	p.Stop()
}

// check one line that without line break
func TestOneLineWithoutWrap(t *testing.T) {
	emptyFile()
	p := slp.NewPrinter()
	for i := 0; i <= 100; i++ {
		s := strconv.Itoa(i)
		p.Print(s)
		if i == 0 {
			expect += esc("?25l") + s
		} else {
			lastLen := strconv.Itoa(len(strconv.Itoa(i - 1)))
			expect += fmt.Sprintf("%s%d", esc(lastLen+"D", "0K"), i)
		}
	}
	if readFile() != expect {
		t.Error("result does not match the expectation")
	}
	p.Stop()
}

// check multi lines that with line break
func TestMultiLinesWithWrap(t *testing.T) {
	emptyFile()
	p := slp.NewPrinter()
	termSize := terminal.GetSize()
	s := strings.Repeat("-", int(termSize.Width+1)) + "\n"
	for i := 0; i < 3; i++ {
		p.Print(s)
		if i == 0 {
			expect += esc("?25l")
		} else {
			expect += esc("1A", "0K", "1A", "0K")
		}
		expect += s
	}
	if readFile() != expect {
		t.Error("result does not match the expectation")
	}
	p.Stop()
}

// check multi lines that without line break
func TestMultiLinesWithoutWrap(t *testing.T) {
	emptyFile()
	p := slp.NewPrinter()
	termSize := terminal.GetSize()
	s := strings.Repeat("-", int(termSize.Width+1))
	for i := 0; i < 3; i++ {
		p.Print(s)
		if i == 0 {
			expect += esc("?25l")
		} else {
			expect += esc("1D", "1A", "0K")
		}
		expect += s
	}
	if readFile() != expect {
		t.Error("result does not match the expectation")
	}
	p.Stop()
}

// check Chinese that with multiple line breaks
func TestOneChineseLineWithoutWrap(t *testing.T) {
	emptyFile()
	p := slp.NewPrinter()
	s := "中\n文\ntest测试"
	for i := 0; i < 3; i++ {
		p.Print(s)
		if i == 0 {
			expect += esc("?25l")
		} else {
			expect += esc("8D", "1A", "0K", "1A", "0K")
		}
		expect += s
	}
	if readFile() != expect {
		t.Error("result does not math the expectation")
	}
	p.Stop()
	removeFile()
}

// create the "output" file
func createFile() {
	var err error
	file, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		panic("can not create file: " + err.Error())
	}
	os.Stdout = file
}

// remove the "output" file
func removeFile() {
	err := os.Remove(filePath)
	if err != nil {
		panic("can not remove file: " + err.Error())
	}
	expect = ""
}

// empty the "output" file
func emptyFile() {
	removeFile()
	createFile()
}

// read the "output" file
func readFile() string {
	createFile()
	content, err := io.ReadAll(file)
	if err != nil {
		panic("can not read file: " + err.Error())
	}
	os.Stdout = oldStdout
	return string(content)
}

// the reference of util.ESC
func esc(suffix ...string) string {
	return util.ESC(suffix...)
}
