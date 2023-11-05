package main

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	content, err := os.ReadFile("expected")
	if err != nil {
		panic(err)
	}
	expected := string(content)
	var b strings.Builder
	fizzBuzz(&b, 20)
	actual := b.String()
	if expected != actual {
		panic(fmt.Sprintf("unexpected output:\n%s", actual))
	}
}

func BenchmarkFizzBuzz100k(b *testing.B) {
	f, err := os.OpenFile(os.DevNull, os.O_WRONLY, fs.ModeAppend)
	if err != nil {
		panic(err)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fizzBuzz(f, 100_000)
	}
}
