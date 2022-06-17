package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Naczaaja")
	want := "Hello, Naczaaja"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
