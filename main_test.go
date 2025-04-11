package main

import "testing"

func TestDoMathAdd(t *testing.T) {
	actual := doMath(5, 5, add)
	expected := 10
	if actual != expected {
		t.Errorf("Sum was incorrect, got: %v, want: %v.", actual, 10)
	}
}

func TestAdd(t *testing.T) {
	actual := add(9, 3)
	expected := 12
	if actual != expected {
		t.Errorf("Sum was incorrect, got: %v, want: %v.", actual, 10)
	}
}

func TestDoMathSubtract(t *testing.T) {
	actual := doMath(15, 5, subtract)
	expected := 10
	if actual != expected {
		t.Errorf("Sum was incorrect, got: %v, want: %v.", actual, 10)
	}
}

func TestSubtract(t *testing.T) {
	actual := subtract(9, 6)
	expected := 3
	if actual != expected {
		t.Errorf("Sum was incorrect, got: %v, want: %v.", actual, 10)
	}
}
