package main

import (
	"fmt"
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
