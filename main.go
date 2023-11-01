package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

const BUFFER_SIZE = 8_000

func main() {
	fizzBuzz(os.Stdout, 1_000_000_000)
}

func fizzBuzz(w io.Writer, n int) {
	bufOut := bufio.NewWriterSize(w, BUFFER_SIZE)
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			bufOut.WriteString("FizzBuzz\n")
			continue
		}
		if i%3 == 0 {
			bufOut.WriteString("Fizz\n")
			continue
		}
		if i%5 == 0 {
			bufOut.WriteString("Buzz\n")
			continue
		}
		bufOut.WriteString(strconv.Itoa(i))
		bufOut.WriteRune('\n')
	}
	bufOut.Flush()
}
