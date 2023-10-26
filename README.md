# single-line-print [![](https://camo.githubusercontent.com/315a8800fc96d3c5b32e13227b10500ef850688793cc6664418d018980eb3cb4/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f676f737572692f75696c6976653f7374617475732e737667)](https://pkg.go.dev/github.com/lianggaoqiang/single-line-print) [![https://github.com/lianggaoqiang/single-line-print/blob/main/LICENSE](https://img.shields.io/badge/license-MIT-red.svg)](https://github.com/lianggaoqiang/single-line-print/blob/main/LICENSE) [![](https://github.com/lianggaoqiang/single-line-print/actions/workflows/ci.yml/badge.svg)](https://github.com/lianggaoqiang/single-line-print/actions/workflows/ci.yml)

Single-Line-Print is a cross-platform terminal single line printing program implemented with Go. 

It features the following benefits:
+ Ease of use and pretty good cross-platform performance (Linux, Windows, MacOS and so on).
+ Custom output and run mode. With this ability, you can disable input during the printing process, hide cursor, and even maintain stable output when dynamically modifying the terminal window.
+ Available two styles: Simple Printer and Writer. With this ability, you can print directly with Printer, or you can use buffering function of Writer to process extra long text.

<br>

## Usage

```go
import slp "github.com/lianggaoqiang/single-line-print"

func main() {
	p := slp.NewPrinter()
	for i := 0; i <= 100; i++ {
		p.Print(fmt.Sprintf("> Downloading from remote: %d%%\n\n", i))
		time.Sleep(time.Millisecond * 30)
	}
	p.Stop()
}

```

<img src="https://github.com/lianggaoqiang/single-line-print/blob/main/doc/simple-example.gif" style="width:70%" />

<br>

## Why need p.Stop?

1. When initializing with either `NewPrinter` or `NewWriter`, the resulting instance will default to disabling terminal input and hiding the cursor during the printing process.To change this behavior, see [Use Flag](#Use-Flag)
<br>- After printing is completed, if you want to restore terminal input and redisplay the cursor, you can call Stop to achieve this goal.

2. There can only be one instance of printer or writer globally, and before creating another instance, the existing instance must be registered using Stop to avoid the program from crashing and throwing an error.

<br>

## Use Writer

If you want to utilize it as a writer, you can obtain an instance that implements io.Writer through NewWriter.

```go
import slp "github.com/lianggaoqiang/single-line-print"

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
```

<br>

## Use Flag

NewPrinter and NewWriter default to disable terminal input and hide the cursor. If you want to change this behavior, you can use `NewPrinterWithFlag` or `NewWriterWithFlag` instead. There are three flag bits available, and their functions are as follows:
+ HideCursor: if set, the cursor will be hidden during printing or writing.
+ DisableInput: if set, input will be disabled during printing or writing.
+ ResizeReactively: if set, terminal window size will be got before each printing or writing. This will ensure more stable terminal output when the window is resized dynamically during printing.

```go
import slp "github.com/lianggaoqiang/single-line-print"

func main() {
	p := slp.NewPrinterWithFlag(
		slp.HideCursor | slp.DisableInput | slp.ResizeReactively,
	)
	// do some prints
	p.Stop()
}
```

<br>

## FAQ

1. Previous output contents was accidentally cleared?
<br>If you print some contents between two single line prints without using Stop, this may happen. To avoid this situation, you can use `Reload` before reusing an existing instance.

<br>

## Installation
```shell
go get github.com/lianggaoqiang/single-line-print
```
