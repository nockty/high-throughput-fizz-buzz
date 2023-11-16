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
	i := 1
	for i+15 <= n {
		bufOut.WriteString(strconv.Itoa(i))
		bufOut.WriteRune('\n')
		i++
		bufOut.WriteString(strconv.Itoa(i))
		bufOut.WriteRune('\n')
		bufOut.WriteString("Fizz\n")
		i += 2
		bufOut.WriteString(strconv.Itoa(i))
		bufOut.WriteRune('\n')
		bufOut.WriteString("Buzz\n")
		bufOut.WriteString("Fizz\n")
		i += 3
		bufOut.WriteString(strconv.Itoa(i))
		bufOut.WriteRune('\n')
		i++
		bufOut.WriteString(strconv.Itoa(i))
		bufOut.WriteRune('\n')
		bufOut.WriteString("Fizz\n")
		bufOut.WriteString("Buzz\n")
		i += 3
		bufOut.WriteString(strconv.Itoa(i))
		bufOut.WriteRune('\n')
		bufOut.WriteString("Fizz\n")
		i += 2
		bufOut.WriteString(strconv.Itoa(i))
		bufOut.WriteRune('\n')
		i++
		bufOut.WriteString(strconv.Itoa(i))
		bufOut.WriteRune('\n')
		bufOut.WriteString("FizzBuzz\n")
		i += 2
	}
	for i <= n {
		if i%15 == 0 {
			bufOut.WriteString("FizzBuzz\n")
		} else if i%3 == 0 {
			bufOut.WriteString("Fizz\n")
		} else if i%5 == 0 {
			bufOut.WriteString("Buzz\n")
		} else {
			bufOut.WriteString(strconv.Itoa(i))
			bufOut.WriteRune('\n')
		}
		i++
	}
	bufOut.Flush()
}
