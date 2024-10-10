package bookstore_test

import (
	"bookstore"
	"fmt"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
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

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{
		1: {Title: "Spark Joy", Author: "Marie Kondo", Copies: 1, ID: 1},
		2: {Title: "Niggastory", Author: "BLC", Copies: 4, ID: 2},
		3: {Title: "Respify", Author: "yupi", Copies: 99, ID: 3},
	}
	want := []bookstore.Book{
		{Title: "Spark Joy", Author: "Marie Kondo", Copies: 1, ID: 1},
		{Title: "Niggastory", Author: "BLC", Copies: 4, ID: 2},
		{Title: "Respify", Author: "yupi", Copies: 99, ID: 3},
	}
	got := catalog.GetAllBooks()

	sort.Slice(got, func(i, j int) bool {
		return got[i].ID < got[j].ID
	})
	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBook(t *testing.T) {
	catalog := bookstore.Catalog{
		1: {Title: "For the Love of Go", ID: 1},
		2: {Title: "Spark Joy", ID: 2},
	}
	want := bookstore.Book{Title: "Spark Joy", ID: 2}
	ID := 2
	got, err := catalog.GetBook(ID)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBookBadIDReturnsError(t *testing.T) {
	t.Parallel()
	catalog := map[int]bookstore.Book{}
	_, err := bookstore.GetBook(catalog, 999)
	if err == nil {
		t.Fatal("want error for non-existent ID, got nil")
	}
}

func TestNetPriceCents(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:           "For the love of go",
		ID:              1,
		Copies:          3,
		Author:          "Marie Kondo",
		PriceCents:      4000,
		DiscountPercent: 25,
	}
	var want float64 = 3000
	got := b.NetPriceCents()
	if got != want {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestSetPriceCents(t *testing.T) {
	t.Parallel()
	book := bookstore.Book{
		Title:           "For the love of go",
		ID:              1,
		Copies:          3,
		Author:          "Marie Kondo",
		PriceCents:      0,
		DiscountPercent: 25,
	}
	want := 2000
	err := book.SetPriceCents(want)
	if err != nil {
		t.Fatal(err)
	}
	got := int(book.PriceCents)
	if want != got {
		t.Errorf("want updated price %d, got %d", want, got)
	}
}

func TestSetPriceCentsInvalid(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:      "For the Love of Go",
		PriceCents: 4000,
	}
	err := b.SetPriceCents(-1)
	if err == nil {
		t.Fatal("want error setting invalid price -1, got nil")
	}
}

func TestSetCategory(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:      "For the Love of Go",
		PriceCents: 4000,
	}
	err := b.SetCategory("Drama")
	if err == nil {
		t.Fatal("want error for invalid category, got nil")
	}
}

func TestInvalidSetC(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:      "For the Love of Go",
		PriceCents: 4000,
	}
	want := "Autobiography"
	err := b.SetCategory(want)
	if err != nil {
		t.Fatal(err)
	}
	got := b.Category()

	if want != got {
		t.Errorf("want category %q, got %q", want, got)
	}
}
