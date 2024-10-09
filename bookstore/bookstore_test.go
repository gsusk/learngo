package bookstore_test

import (
	"bookstore"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
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

func TestGetAllBook(t *testing.T) {
	t.Parallel()
	catalog := []bookstore.Book{
		{Title: "Spark Joy", Author: "Marie Kondo", Copies: 1},
		{Title: "Niggastory", Author: "BLC", Copies: 4},
		{Title: "Respify", Author: "yupi", Copies: 99},
	}
	want := []bookstore.Book{
		{Title: "Spark Joy", Author: "Marie Kondo", Copies: 1},
		{Title: "Niggastory", Author: "BLC", Copies: 4},
		{Title: "Respify", Author: "yupi", Copies: 99},
	}
	got := bookstore.GetAllBooks(catalog)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBook(t *testing.T) {
	catalog := []bookstore.Book{
		{Title: "For the Love of Go", ID: 1},
		{Title: "Spark Joy", ID: 2},
	}
	want := bookstore.Book{Title: "Spark Joy", ID: 2}
	ID := 2
	got := bookstore.GetBook(catalog, ID)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
