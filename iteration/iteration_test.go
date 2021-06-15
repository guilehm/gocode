package main

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("a")
	expected := "aaaa"

	if got != expected {
		t.Errorf("expected %q but got %q", expected, got)
	}
}
