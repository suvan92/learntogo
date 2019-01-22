package main

import "testing"

func TestFactorial(t *testing.T) {
	// Single test
	v := factorial(4)
	if v != 24 {
		t.Errorf("Expected 24, got %d", v)
	}

	// Table test
	type test struct {
		input  int
		answer int
	}

	tests := []test{
		test{input: 0, answer: 1},
		test{input: 1, answer: 1},
		test{input: 2, answer: 2},
		test{input: 3, answer: 6},
		test{input: 4, answer: 24},
	}

	for _, v := range tests {
		x := factorial(v.input)
		if x != v.answer {
			t.Error("Expected", v.answer, "got", x)
		}
	}
}
