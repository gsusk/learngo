package creditcard_test

import (
	"creditcard"
	"testing"
)

func TestNew(t *testing.T) {
	want := "123456789"
	cc, err := creditcard.New(want)
	if err != nil {
		t.Fatal(err)
	}
	got := cc.Number()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
