package main

import (
	"bufio"
	"os"
	"strconv"
)

const BUFFER_SIZE = 8_000

func main() {
	bufStdout := bufio.NewWriterSize(os.Stdout, BUFFER_SIZE)
	for i := 1; i <= 1_000_000_000; i++ {
		if i%3 == 0 && i%5 == 0 {
			bufStdout.WriteString("FizzBuzz\n")
			continue
		}
		if i%3 == 0 {
			bufStdout.WriteString("Fizz\n")
			continue
		}
		if i%5 == 0 {
			bufStdout.WriteString("Buzz\n")
			continue
		}
		bufStdout.WriteString(strconv.Itoa(i))
		bufStdout.WriteRune('\n')
	}
	bufStdout.Flush()
}
