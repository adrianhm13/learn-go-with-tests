package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Adrian")
	want := "Hello, Adrian"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
