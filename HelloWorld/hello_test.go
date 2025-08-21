package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Genki")
	want := "Hello, Genki"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
