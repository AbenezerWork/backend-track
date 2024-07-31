package models

import "fmt"

type Book struct {
	ID     int
	Title  string
	Author string
	Status string
}

func (book Book) PrettyPrintBook() {
	fmt.Println("Book Details")
	fmt.Println("------------")
	fmt.Printf("ID:     %d\n", book.ID)
	fmt.Printf("Title:  %s\n", book.Title)
	fmt.Printf("Author: %s\n", book.Author)
	fmt.Printf("Status: %s\n", book.Status)
	fmt.Println("------------")
}

func (b *Book) Borrow() {
	b.Status = "Borrowed"
}

func (b *Book) Return() {
	b.Status = "Available"
}
