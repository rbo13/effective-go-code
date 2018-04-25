package main

import (
	"bytes"
	"testing"
)

func TestGreeter(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Richard")

	got := buffer.String()
	want := "Hello, Richard"

	if got != want {
		t.Errorf("Got '%s' want '%s'", got, want)
	}
}
