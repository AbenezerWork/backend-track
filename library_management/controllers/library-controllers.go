package controllers

import (
	"bufio"
	"errors"
	"fmt"
	"library_management/services"
	"os"
	"strconv"
	"strings"
)

func getInt(b *bufio.Scanner) (int, error) {
	b.Scan()
	line := b.Text()
	id, err := strconv.Atoi(line)

	if err != nil {
		return -1, errors.New("Trouble parsing the input make sure ")
	}
	return id, nil
}

func ListAllAvailableBooks(Library *services.Library) {
	Library.ListAvailableBooks()
}

func ListAllBorrowedBooks(Library *services.Library) {
	Library.ListBorrowedBooks()
}

func ReturnBook(b *bufio.Scanner, Library *services.Library, userID int) error {
	fmt.Print("Enter the ID of the book you want to remove; ")
	line := b.Text()

	id, err := strconv.Atoi(line)

	if err != nil {
		return errors.New("Trouble parsing the input make sure ")
	}
	Library.ReturnBook(id, userID)
	return nil

}

func AddBook(b *bufio.Scanner, Library *services.Library) error {
	fmt.Print("Enter the name of the book you will be adding: ")
	line := b.Text()
	title := strings.TrimSpace(line)

	fmt.Print("Enter the name of the author of the book you will be adding: ")
	line = b.Text()
	author := strings.TrimSpace(line)

	if len(title) == 0 || len(author) == 0 || len(title) > 50 || len(author) > 50 {
		return errors.New("Invalid string length")
	}

	Library.AddBook(title, author)
	return nil
}

func RemoveBook(b *bufio.Scanner, Library *services.Library) error {
	fmt.Print("Enter the ID of the book you want to remove; ")
	line := b.Text()
	id, err := strconv.Atoi(line)

	if err != nil {
		return errors.New("Trouble parsing the input make sure ")
	}
	Library.RemoveBook(id)
	return nil
}

func BorrowBook(b *bufio.Scanner, Library *services.Library, userID int) error {
	fmt.Print("Enter the ID of the book you want to remove; ")
	line := b.Text()

	id, err := strconv.Atoi(line)

	if err != nil {
		return errors.New("Trouble parsing the input make sure ")
	}
	Library.BorrowBook(id, userID)
	return nil
}

func Route() {
	library := services.MakeLibrary()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("WELCOME to our library!!")
	var input int
	for true {
		fmt.Println("Are you a return cutomer or a new one: ")
		fmt.Print("Enter 1 if you're a returing customer, 2 if you're new: ")
		input, err := getInt(scanner)
		if err != nil {
			fmt.Println(err)
			continue
		} else if input != 1 && input != 2 {
			fmt.Println("You have entered an invalid number.")
			continue
		}
		break
	}
	var id int
	if input == 2 {
		id = library.AddMemeber()
		fmt.Println("Your ID is ", id, " you will use this for further interactions.")
	}

	for {
		fmt.Println("Library Management System")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("7. Exit")
		fmt.Print("Enter choice: ")

		if !scanner.Scan() {
			fmt.Println("Failed to read input")
			return
		}

		choice, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid choice, please enter a number")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Enter book title: ")
			scanner.Scan()
			bookTitle := scanner.Text()

			fmt.Print("Enter book author: ")
			scanner.Scan()
			bookAuthor := scanner.Text()

			library.AddBook(bookTitle, bookAuthor)

		case 2:
			fmt.Print("Enter book ID: ")
			scanner.Scan()
			bookID, _ := strconv.Atoi(scanner.Text())
			library.RemoveBook(bookID)

		case 3:
			fmt.Print("Enter book ID: ")
			scanner.Scan()
			bookID, _ := strconv.Atoi(scanner.Text())

			fmt.Print("Enter member ID: ")
			scanner.Scan()
			memberID, _ := strconv.Atoi(scanner.Text())

			if err := library.BorrowBook(bookID, memberID); err != nil {
				fmt.Println("Error:", err)
			}

		case 4:
			fmt.Print("Enter book ID: ")
			scanner.Scan()
			bookID, _ := strconv.Atoi(scanner.Text())

			fmt.Print("Enter member ID: ")
			scanner.Scan()
			memberID, _ := strconv.Atoi(scanner.Text())

			if err := library.ReturnBook(bookID, memberID); err != nil {
				fmt.Println("Error:", err)
			}

		case 5:
			library.ListAvailableBooks()

		case 6:
			library.ListBorrowedBooks()

		case 7:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice, please try again")
		}
	}
}
