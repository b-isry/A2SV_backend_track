package services

import (
	"Library_Management/models"
	"fmt"
	"sync"
	"time"
)

type Library struct {
	Books   map[int]models.Books
	Members map[int]models.Member
	mu      sync.Mutex
}

type LibraryManager interface {
	AddBook(book models.Books)
	AddMember(member models.Member)
	RemoveBook(bookId int)
	BorrowBook(memberID int, bookID int)
	ReturnBook(memberID int, bookID int)
	ReserveBook(memberID int, bookID int) error
	ListAvailableBooks() []models.Books
	ListBorrowedBooks() []models.Books
}

func (l *Library) AddBook(book models.Books) {
	l.Books[book.Id] = book
	book.Status = "available"
	fmt.Println("Book added successfully")
}

func (l *Library) AddMember(member models.Member) {
	l.Members[member.Id] = member
	fmt.Println("Member added successfully")
}

func (l *Library) RemoveBook(bookId int) {
	delete(l.Books, bookId)
	fmt.Println("Book removed successfully")
}

func (l *Library) BorrowBook(memberID int, bookID int) {
	member, ok := l.Members[memberID]
	if !ok {
		fmt.Println("Member not found")
		return
	}
	book, ok := l.Books[bookID]
	if !ok {
		fmt.Println("Book not found")
		return
	}
	if book.Status != "available" {
		fmt.Println("Book is not available")
		return
	}
	book.Status = "borrowed"
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.Books[bookID] = book

	fmt.Printf(" %s borrowed %s successfully\n", member.Name, book.Title)
}

func (l *Library) ReturnBook(memberID int, bookID int) {
	member, ok := l.Members[memberID]
	if !ok {
		fmt.Println("Member not found")
		return
	}
	book, ok := l.Books[bookID]
	if !ok {
		fmt.Println("Book not found")
		return
	}
	if book.Status != "borrowed" {
		fmt.Println("Book is not borrowed")
		return
	}
	book.Status = "available"
	member.BorrowedBooks = removeFromSlice(member.BorrowedBooks, book)
	l.Books[bookID] = book
	fmt.Printf(" %s returned %s successfully\n", member.Name, book.Title)
}

func (l *Library) ReserveBook(memberID int, bookID int) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	member, ok := l.Members[memberID]
	if !ok {
		return fmt.Errorf("member not found")
	}
	book, ok := l.Books[bookID]
	if !ok {
		return fmt.Errorf("book not found")
	}
	if book.Status != "available" {
		return fmt.Errorf("book is not available")
	}

	book.Status = "reserved"
	l.Books[bookID] = book
	member.ReservedBooks = append(member.ReservedBooks, book)
	l.Members[memberID] = member

	// Start a goroutine to handle the reservation and borrowing process
	go func() {
		// Wait for 2 seconds to simulate processing time
		time.Sleep(2 * time.Second)

		l.mu.Lock()
		defer l.mu.Unlock()

		// Check if the book is still reserved
		if book, exists := l.Books[bookID]; exists && book.Status == "reserved" {
			// Process the borrowing
			book.Status = "borrowed"
			l.Books[bookID] = book

			// Update member's books
			if member, exists := l.Members[memberID]; exists {
				member.ReservedBooks = removeFromSlice(member.ReservedBooks, book)
				member.BorrowedBooks = append(member.BorrowedBooks, book)
				l.Members[memberID] = member
			}

			fmt.Printf("Book %s has been borrowed by %s\n", book.Title, member.Name)
		}
	}()

	// Start a goroutine to auto-cancel the reservation after 5 seconds if not borrowed
	go func() {
		time.Sleep(5 * time.Second)
		l.mu.Lock()
		defer l.mu.Unlock()

		// Check if the book is still reserved
		if book, exists := l.Books[bookID]; exists && book.Status == "reserved" {
			book.Status = "available"
			l.Books[bookID] = book
			// Remove from member's reserved books
			if member, exists := l.Members[memberID]; exists {
				member.ReservedBooks = removeFromSlice(member.ReservedBooks, book)
				l.Members[memberID] = member
			}
			fmt.Printf("Reservation for book %s automatically cancelled\n", book.Title)
		}
	}()

	fmt.Printf("%s reserved %s successfully\n", member.Name, book.Title)
	return nil
}

func (l *Library) ListAvailableBooks() []models.Books {
	availableBooks := []models.Books{}
	for _, book := range l.Books {
		if book.Status == "available" {
			availableBooks = append(availableBooks, book)
		}
	}
	for _, book := range availableBooks {
		fmt.Println(book.Title)
	}
	return availableBooks
}

func (l *Library) ListBorrowedBooks() []models.Books {
	borrowedBooks := []models.Books{}
	for _, book := range l.Books {
		if book.Status == "borrowed" {
			borrowedBooks = append(borrowedBooks, book)
		}
	}
	for _, book := range borrowedBooks {
		fmt.Println(book.Title)
	}
	return borrowedBooks
}

func removeFromSlice(slice []models.Books, book models.Books) []models.Books {
	for i, b := range slice {
		if b.Id == book.Id {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
