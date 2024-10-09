package bookstore

import "errors"

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

func GetBook(catalog map[int]Book, bookID int) Book {
	return catalog[bookID]
}
