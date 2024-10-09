package bookstore

import (
	"errors"
	"fmt"
)

type Book struct {
	Title  string
	Author string
	Copies int
	ID     int
}

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("No copies left.")
	}
	b.Copies -= 1
	return b, nil
}

func GetAllBooks(c []Book) []Book {
	return c
}

func GetBook(catalog map[int]Book, bookID int) (Book, error) {
	book, ok := catalog[bookID]
	if !ok {
		return Book{}, fmt.Errorf("ID %d doesnt exixst", bookID)
	}
	return book, nil
}
