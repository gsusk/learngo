package mytypes_test

import (
	"mytypes"
	"strings"
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

func TestStringBuild(t *testing.T) {
	t.Parallel()
	var sb strings.Builder
	sb.WriteString("Hello, ")
	sb.WriteString("Gophers!")
	want := "Hello, Gophers!"
	got := sb.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
	wantLen := 15
	gotLen := sb.Len()
	if wantLen != gotLen {
		t.Errorf("%q: want len %d, got %d", sb.String(), wantLen, gotLen)
	}
}
