package main

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected, '%d' but got '%d'", expected, sum)
	}
}

func TestMultiplier(t *testing.T) {
	mul := Multiply(2, 3)
	expected := 6

	if mul != expected {
		t.Errorf("expected, '%d' but got '%d'", expected, mul)
	}
}
