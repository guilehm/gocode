package main

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("a", 4)
	expected := "aaaa"

	if got != expected {
		t.Errorf("expected %q but got %q", expected, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 4)
	}
}
