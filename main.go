package main

import (
	"bufio"
	"io"
	"os"
)

const BUFFER_SIZE = 8_000

const (
	FIZZ     = "Fizz\n"
	BUZZ     = "Buzz\n"
	FIZZBUZZ = "FizzBuzz\n"
	NEWLINE  = '\n'
)

func main() {
	fizzBuzz(os.Stdout, 1_000_000_000)
}

func fizzBuzz(w io.Writer, n int) {
	bufOut := bufio.NewWriterSize(w, BUFFER_SIZE)
	// Intermediate buffer for writing integers. The max int64 is 19 digits in base 10.
	var a [19]byte
	i := 1
	for i+15 <= n {
		writeInt(i, bufOut, &a)
		bufOut.WriteRune(NEWLINE)
		i++
		writeInt(i, bufOut, &a)
		bufOut.WriteRune(NEWLINE)
		bufOut.WriteString(FIZZ)
		i += 2
		writeInt(i, bufOut, &a)
		bufOut.WriteRune(NEWLINE)
		bufOut.WriteString(BUZZ)
		bufOut.WriteString(FIZZ)
		i += 3
		writeInt(i, bufOut, &a)
		bufOut.WriteRune(NEWLINE)
		i++
		writeInt(i, bufOut, &a)
		bufOut.WriteRune(NEWLINE)
		bufOut.WriteString(FIZZ)
		bufOut.WriteString(BUZZ)
		i += 3
		writeInt(i, bufOut, &a)
		bufOut.WriteRune(NEWLINE)
		bufOut.WriteString(FIZZ)
		i += 2
		writeInt(i, bufOut, &a)
		bufOut.WriteRune(NEWLINE)
		i++
		writeInt(i, bufOut, &a)
		bufOut.WriteRune(NEWLINE)
		bufOut.WriteString(FIZZBUZZ)
		i += 2
	}
	for i <= n {
		if i%15 == 0 {
			bufOut.WriteString(FIZZBUZZ)
		} else if i%3 == 0 {
			bufOut.WriteString(FIZZ)
		} else if i%5 == 0 {
			bufOut.WriteString(BUZZ)
		} else {
			writeInt(i, bufOut, &a)
			bufOut.WriteRune(NEWLINE)
		}
		i++
	}
	bufOut.Flush()
}

const smallsString = "00010203040506070809" +
	"10111213141516171819" +
	"20212223242526272829" +
	"30313233343536373839" +
	"40414243444546474849" +
	"50515253545556575859" +
	"60616263646566676869" +
	"70717273747576777879" +
	"80818283848586878889" +
	"90919293949596979899"

// writeInt writes u in base 10 in b, using a as an intermediate buffer
func writeInt(u int, b *bufio.Writer, a *[19]byte) {
	i := 19

	for u >= 100 {
		is := u % 100 * 2
		u /= 100
		i -= 2
		a[i+1] = smallsString[is+1]
		a[i+0] = smallsString[is+0]
	}

	// us < 100
	is := u * 2
	i--
	a[i] = smallsString[is+1]
	if u >= 10 {
		i--
		a[i] = smallsString[is]
	}

	b.Write(a[i:])
}
