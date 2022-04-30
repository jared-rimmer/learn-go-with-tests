package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Jazzy")
	want := "Hello, Jazzy"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
