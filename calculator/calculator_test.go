package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(
	t *testing.T,
) {
	t.Parallel()
	var want float64 = 4
	got := calculator.Add(2, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

/*
To a test to work-out
you need
- A file that ends in _test.go
- functions starting with Test
- A param that is of type *testing.T
- returns nothing
*/

func TestSubtract(t *testing.T) {
	t.Parallel()
	var want float64 = 2
	got := calculator.Subtract(4, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}
