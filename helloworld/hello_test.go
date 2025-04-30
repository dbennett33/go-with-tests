package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Dan")
	want := "Hello, Dan"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
