package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Aron")

	got := buffer.String()
	want := "Hello, Aron"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
