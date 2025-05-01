package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Dan")

	got := buffer.String()
	want := "Hello, Dan"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
