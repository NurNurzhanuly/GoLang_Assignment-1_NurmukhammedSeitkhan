package library

import "fmt"

type Book struct {
	ID         string
	Title      string
	Author     string
	IsBorrowed bool
}

type Library struct {
	Books map[string]Book
}

func NewLibrary() *Library {
	return &Library{Books: make(map[string]Book)}
}

func (lib *Library) AddBook(book Book) {
	lib.Books[book.ID] = book
}

func (lib *Library) BorrowBook(id string) {
	if book, exists := lib.Books[id]; exists && !book.IsBorrowed {
		book.IsBorrowed = true
		lib.Books[id] = book
		fmt.Println("Book borrowed:", book.Title)
	} else {
		fmt.Println("Book not available.")
	}
}

func (lib *Library) ReturnBook(id string) {
	if book, exists := lib.Books[id]; exists && book.IsBorrowed {
		book.IsBorrowed = false
		lib.Books[id] = book
		fmt.Println("Book returned:", book.Title)
	} else {
		fmt.Println("Book not borrowed.")
	}
}

func (lib *Library) ListBooks() {
	for _, book := range lib.Books {
		if !book.IsBorrowed {
			fmt.Printf("ID: %s, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
		}
	}
}
