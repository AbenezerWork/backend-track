# Console-Based Library Management System

## Objective

The goal is to create a simple console-based library management system in Go, demonstrating the use of structs, interfaces, methods, slices, and maps.

## Requirements

### Structs

#### Book

Defines a `Book` struct with the following fields:
- `ID` (int)
- `Title` (string)
- `Author` (string)
- `Status` (string) // can be "Available" or "Borrowed"

#### Member

Defines a `Member` struct with the following fields:
- `ID` (int)
- `Name` (string)
- `BorrowedBooks` ([]Book) // a slice to hold borrowed books

### Interfaces

#### LibraryManager

Defines a `LibraryManager` interface with the following methods:
- `AddBook(book Book)`
- `RemoveBook(bookID int)`
- `BorrowBook(bookID int, memberID int) error`
- `ReturnBook(bookID int, memberID int) error`
- `ListAvailableBooks() []Book`
- `ListBorrowedBooks(memberID int) []Book`

### Implementation

Implements the `LibraryManager` interface in a `Library` struct. The `Library` struct should have fields to store all books (use a map with book ID as the key) and members (use a map with member ID as the key).

#### Methods

Implements the methods defined in the `LibraryManager` interface:
- `AddBook`: Adds a new book to the library.
- `RemoveBook`: Removes a book from the library by its ID.
- `BorrowBook`: Allows a member to borrow a book if it is available.
- `ReturnBook`: Allows a member to return a borrowed book.
- `ListAvailableBooks`: Lists all available books in the library.
- `ListBorrowedBooks`: Lists all books borrowed by a specific member.

### Console Interaction

Creates a simple console interface to interact with the library management system. Implements functions to:
- Add a new book.
- Remove an existing book.
- Borrow a book.
- Return a book.
- List all available books.
- List all borrowed books by a member.

## Folder Structure

The folder structure for this task is as follows:

library_management/
├── main.go
├── controllers/
│ └── library_controller.go
├── models/
│ └── book.go
│ └── member.go
├── services/
│ └── library_service.go
├── docs/
│ └── documentation.md
└── go.mod

### Files

#### `main.go`

Entry point of the application.

#### `controllers/library_controller.go`

Handles console input and invokes the appropriate service methods.

#### `models/book.go`

Defines the `Book` struct.

#### `models/member.go`

Defines the `Member` struct.

#### `services/library_service.go`

Contains business logic and data manipulation functions.

#### `docs/documentation.md`

Contains system documentation and other related information.

## Getting Started

1. Clone the repository.
2. Navigate to the project directory.
3. Run the application using `go run .`.

## Usage

Follow the on-screen prompts to interact with the library management system. The options include adding books, removing books, borrowing books, returning books, listing available books, and listing borrowed books by members.

## Error Handling

The system includes error handling for scenarios such as:
- Book or member not found.
- Book already borrowed.
- Book not borrowed by the member attempting to return it.

## Conclusion

This project provides a console-based library management system in Go, showcasing the use of core Go features such as structs, interfaces, methods, slices, and maps. The system includes all essential functionalities, appropriate error handling, and a well-organized code structure.
