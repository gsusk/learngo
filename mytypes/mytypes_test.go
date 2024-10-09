package mytypes_test

import (
	"mytypes"
	"testing"
)

func TestTwice(t *testing.T) {
	input := mytypes.MyInt(9)
	want := mytypes.MyInt(18)
	got := input.Twice()
	if want != got {
		t.Errorf("twice: want %d, got %d", input, want)
	}
}

func TestMyString(t *testing.T) {
	input := mytypes.MyString("this is my string")
	want := 17
	got := input.Len()
	if want != got {
		t.Errorf("%q: want: %d, got: %d", input, want, got)
	}
}
