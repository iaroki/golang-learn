package main

import "testing"

func TestHelloFunc(t *testing.T) {
	got := HelloFunc("Max")
	want := "Hello, Max"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
