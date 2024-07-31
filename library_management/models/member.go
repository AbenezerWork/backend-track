package models

type Member struct {
	ID            int
	Name          string
	BorrowedBooks []Book
}

func (m *Member) Borrow(book Book) {
	m.BorrowedBooks = append(m.BorrowedBooks, book)
}

func (m *Member) Return(book Book) {
	//find the book in the BorrowedBooks slice
	idx := 0
	for i, curr := range m.BorrowedBooks {
		if book.ID == curr.ID {
			idx = i
		}
	}

	//remove item idx from the slice
	m.BorrowedBooks = append(m.BorrowedBooks[:idx], m.BorrowedBooks[idx+1:]...)
}
