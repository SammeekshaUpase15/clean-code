package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Book struct {
	ID, Title, Author string
}

type User struct {
	ID, Name, Borrowed string
}

type Library struct {
	Books []Book
	Users []User
}

func (lib *Library) AddBook(id, title, author string) {
	lib.Books = append(lib.Books, Book{ID: id, Title: title, Author: author})
	fmt.Printf("Book added: %s\n", title)
}

func (lib *Library) RemoveBook(id string) {
	for i, book := range lib.Books {
		if book.ID == id {
			lib.Books = append(lib.Books[:i], lib.Books[i+1:]...)
			fmt.Printf("Book removed: %s\n", book.Title)
			return
		}
	}
	fmt.Println("Book not found")
}

func (lib *Library) ListBooks() {
	fmt.Println("List of Books:")
	for _, book := range lib.Books {
		fmt.Printf("ID: %s, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}

func (lib *Library) AddUser(id, name string) {
	lib.Users = append(lib.Users, User{ID: id, Name: name})
	fmt.Printf("User added: %s\n", name)
}

func (lib *Library) BorrowBook(userID, bookID string) {
	user, err := lib.findUser(userID)
	if err != nil {
		fmt.Println(err)
		return
	}
	book, err := lib.findBook(bookID)
	if err != nil {
		fmt.Println(err)
		return
	}
	if user.Borrowed != "" {
		fmt.Println("User already has a borrowed book")
		return
	}
	user.Borrowed = book.ID
	fmt.Printf("Book borrowed: %s by %s\n", book.Title, user.Name)
}

func (lib *Library) ReturnBook(userID, bookID string) {
	user, err := lib.findUser(userID)
	if err != nil {
		fmt.Println(err)
		return
	}
	if user.Borrowed != bookID {
		fmt.Println("Mismatch: User did not borrow this book")
		return
	}
	user.Borrowed = ""
	fmt.Printf("Book returned: %s by %s\n", bookID, user.Name)
}

func (lib *Library) ListUsers() {
	fmt.Println("Users in the system:")
	for _, user := range lib.Users {
		borrowed := "None"
		if user.Borrowed != "" {
			borrowed = user.Borrowed
		}
		fmt.Printf("ID: %s, Name: %s, Borrowed Book ID: %s\n", user.ID, user.Name, borrowed)
	}
}

func (lib *Library) findUser(id string) (*User, error) {
	for i := range lib.Users {
		if lib.Users[i].ID == id {
			return &lib.Users[i], nil
		}
	}
	return nil, errors.New("user not found")
}

func (lib *Library) findBook(id string) (*Book, error) {
	for i := range lib.Books {
		if lib.Books[i].ID == id {
			return &lib.Books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func getUserInput(prompt string) string {
	fmt.Print(prompt + " ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func main() {
	library := Library{}
	for {
		fmt.Println("\nLibrary Menu:")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. List Books")
		fmt.Println("4. Add User")
		fmt.Println("5. Borrow Book")
		fmt.Println("6. Return Book")
		fmt.Println("7. List Users")
		fmt.Println("8. Exit")

		switch getUserInput("Enter your choice:") {
		case "1":
			library.AddBook(getUserInput("Enter Book ID:"), getUserInput("Enter Book Title:"), getUserInput("Enter Author Name:"))
		case "2":
			library.RemoveBook(getUserInput("Enter Book ID to remove:"))
		case "3":
			library.ListBooks()
		case "4":
			library.AddUser(getUserInput("Enter User ID:"), getUserInput("Enter User Name:"))
		case "5":
			library.BorrowBook(getUserInput("Enter User ID:"), getUserInput("Enter Book ID to borrow:"))
		case "6":
			library.ReturnBook(getUserInput("Enter User ID:"), getUserInput("Enter Book ID to return:"))
		case "7":
			library.ListUsers()
		case "8":
			fmt.Println("Exiting Library System.")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
