package bookstore_test

import (
	"bookstore"
	"fmt"
	"testing"
)

func TestBook(t *testing.T) {
	_ = bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
	}
}

func TestBuy(t *testing.T) {
	t.Parallel()
	book := bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
	}
	want := 1
	got, err := bookstore.Buy(book)
	if err != nil {
		t.Fatal(err)
	}
	if got.Copies != want {
		fmt.Errorf("want %v, got %v", want, got)
	}
}

func TestBuyErrorsIfNoCopiesLeft(t *testing.T) {
	t.Parallel()
	book := bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 0,
	}
	_, err := bookstore.Buy(book)
	if err == nil {
		t.Error("want error buying from zero copies, got nil")
	}
}
