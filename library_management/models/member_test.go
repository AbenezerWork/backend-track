package models

import "testing"

func TestReturn(T *testing.T) {
	m := Member{BorrowedBooks: []Book{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}}}
	m.Return(Book{ID: 2})

	ex := Member{BorrowedBooks: []Book{{ID: 1}, {ID: 3}, {ID: 4}}}
	if m.BorrowedBooks[1] != ex.BorrowedBooks[1] {
		T.Error(ex, m)
	}
}
