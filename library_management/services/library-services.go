package services

import (
	"errors"
	"fmt"
	"library_management/models"
	"strings"
)

type LibraryManager interface {
	AddBook(string, string)
	RemoveBook(int)
	BorrowBook(int, int)
	ReturnBook(int)
	ListAvailableBooks()
	ListBorrowedBooks()
}

type Library struct {
	books    map[int]models.Book
	members  map[int]models.Member
	lastID   int
	LastUser int
}

func MakeLibrary() *Library {
	l := Library{}
	l.books = make(map[int]models.Book)
	l.members = make(map[int]models.Member)
	return &l
}

func (l *Library) AddMemeber() int {
	m := models.Member{ID: l.LastUser}
	l.LastUser++
	return m.ID
}

func (l *Library) AddBook(title, author string) {
	b := models.Book{Title: title, Author: author, Status: "Available", ID: l.lastID + 1}
	l.lastID++
	l.books[b.ID] = b
}

func (l *Library) RemoveBook(id int) {
	delete(l.books, id)
}

func (l *Library) BorrowBook(id, userId int) error {
	curr, exists := l.books[id]
	if !exists {
		return errors.New("The books you're looking for doesn't exist!")
	}
	curr.Borrow()
	borrowedBooks := l.members[userId]
	borrowedBooks.Borrow(curr)
	l.books[curr.ID] = curr
	return nil
}

func (l *Library) ReturnBook(id int, userId int) error {
	curr, exists := l.books[id]
	if !exists {
		return errors.New("The book you're looking for doesn't exist!")
	}
	curr.Return()
	borrowedBooks := l.members[userId]
	borrowedBooks.Return(curr)
	l.books[curr.ID] = curr
	return nil
}

func (l Library) ListAvailableBooks() {
	listBooksByStatus(&l.books, "Available")
}

func (l Library) ListBorrowedBooks() {
	listBooksByStatus(&l.books, "Borrowed")
}

func listBooksByStatus(books *map[int]models.Book, status string) {
	fmt.Printf("Books with status '%s':\n", status)
	fmt.Println(strings.Repeat("-", 50))
	fmt.Printf("%-5s %-25s %-20s\n", "ID", "Title", "Author")
	fmt.Println(strings.Repeat("-", 50))

	found := false
	for _, book := range *books {
		if book.Status == status {
			fmt.Printf("%-5d %-25s %-20s\n", book.ID, book.Title, book.Author)
			found = true
		}
	}

	if !found {
		fmt.Println("No books found with the given status.")
	}

	fmt.Println(strings.Repeat("-", 50))
}
