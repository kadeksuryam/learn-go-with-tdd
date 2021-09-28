package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Surya")
	want := "Hello, Surya"

	if got != want {
		t.Errorf("got %q want %q ", got, want)
	}
}
