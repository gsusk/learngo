package bookstore

import (
	"errors"
	"fmt"
)

type Category int

const (
	CategoryAutobiography Category = iota
	CategoryLargePrintRomance
	CategoryParticlePhysics
)

var validCategory = map[Category]bool{
	CategoryAutobiography:     true,
	CategoryLargePrintRomance: true,
	CategoryParticlePhysics:   true,
}

type Book struct {
	Title           string
	Author          string
	Copies          int
	ID              int
	PriceCents      float64
	DiscountPercent float64
	category        Category
}

func (b Book) NetPriceCents() float64 {
	return b.PriceCents * (1 - (b.DiscountPercent / 100))
}

type Catalog map[int]Book

func (c Catalog) GetBook(id int) (Book, error) {
	book, ok := c[id]
	if !ok {
		return Book{}, fmt.Errorf("Book not found, ID: %d", id)
	}
	return book, nil
}

func (c Catalog) GetAllBooks() []Book {
	var data []Book
	for _, book := range c {
		data = append(data, book)
	}
	return data
}

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("No copies left.")
	}
	b.Copies -= 1
	return b, nil
}

func GetBook(catalog map[int]Book, bookID int) (Book, error) {
	book, ok := catalog[bookID]
	if !ok {
		return Book{}, fmt.Errorf("ID %d doesnt exixst", bookID)
	}
	return book, nil
}

func NetPriceCents(b Book) float64 {
	return b.PriceCents * (1 - (b.DiscountPercent / 100))
}

func (b *Book) SetPriceCents(p int) error {
	if p < 0 {
		return fmt.Errorf("Invalid Price : %d", p)
	}
	b.PriceCents = float64(p)
	return nil
}

func (b *Book) SetCategory(cat Category) error {
	if !validCategory[cat] {
		return fmt.Errorf("unknown category %q", cat)
	}
	b.category = cat
	return nil
}

func (b Book) Category() Category {
	return b.category
}
