package main

import (
	"bufio"
	"io/fs"
	"os"
	"strings"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	content, err := os.ReadFile("golden100")
	if err != nil {
		t.Error(err)
	}
	expected := string(content)
	var b strings.Builder
	fizzBuzz(&b, 100)
	actual := b.String()
	if expected != actual {
		t.Errorf("unexpected output:\n%s", actual)
	}
}

type testWriter struct {
	t            *testing.T
	goldenReader *bufio.Reader
}

func (w *testWriter) Write(p []byte) (n int, err error) {
	for i, b := range p {
		golden, err := w.goldenReader.ReadByte()
		if err != nil {
			return 0, err
		}
		if b != golden {
			w.t.Errorf("unexpected output (byte #%d): %s != %s, \n%s", i, string(golden), string(b), p)
		}
	}
	return len(p), nil
}

func TestFizzBuzzHigh(t *testing.T) {
	golden, err := os.Open("golden10m")
	if err != nil {
		t.Error(err)
	}
	tw := &testWriter{
		t:            t,
		goldenReader: bufio.NewReader(golden),
	}

	fizzBuzz(tw, 10_000_000)
}

func TestFillInt(t *testing.T) {
	var a [19]byte
	buf := make([]byte, 100)

	testCases := []struct {
		u        int
		offset   int
		want     int
		expected string
	}{
		{0, 0, 2, "0\n"},
		{42, 0, 3, "42\n"},
		{1337, 42, 5, "1337\n"},
		{100_001, 0, 7, "100001\n"},
		{255_255, 0, 7, "255255\n"},
		{1_010_001, 0, 8, "1010001\n"},
		{3_333_333, 0, 8, "3333333\n"},
		{10_000_000, 0, 9, "10000000\n"},
		{42_042_420, 0, 9, "42042420\n"},
		{100_000_016, 0, 10, "100000016\n"},
		{1_000_000_000, 0, 11, "1000000000\n"},
		{10_010_001_100, 0, 12, "10010001100\n"},
	}
	for _, tc := range testCases {
		n := fillInt(tc.u, buf, tc.offset, &a)
		if n != tc.want {
			t.Errorf("fillInt(%d, _, %d, _) = %d, want %d", tc.u, tc.offset, n, tc.want)
		}
		if string(buf[tc.offset:tc.offset+n]) != tc.expected {
			t.Errorf("fillInt(%d, _, %d, _) writes %s, want %s", tc.u, tc.offset, string(buf[tc.offset:tc.offset+n]), tc.expected)
		}
	}
}

func BenchmarkFizzBuzz100k(b *testing.B) {
	f, err := os.OpenFile(os.DevNull, os.O_WRONLY, fs.ModeAppend)
	if err != nil {
		b.Error(err)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fizzBuzz(f, 100_000)
	}
}
