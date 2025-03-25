// test the library controller
package Test

import (
	"Library_Management/concurrency"
	"Library_Management/models"
	"Library_Management/services"
	"sync"
	"testing"
	"time"
)

func TestBookController(t *testing.T) {
	library := &services.Library{
		Books:   make(map[int]models.Books),
		Members: make(map[int]models.Member),
	}

	//test add book
	book := models.Books{
		Id:     1,
		Title:  "Test Book",
		Author: "Test Author",
		Status: "available",
	}
	library.AddBook(book)

	//test add member
	member := models.Member{
		Id:            1,
		Name:          "Test Member",
		BorrowedBooks: []models.Books{},
	}
	library.AddMember(member)

	//test remove book
	library.RemoveBook(1)

	//test borrow book
	library.BorrowBook(1, 1)

	//test return book
	library.ReturnBook(1, 1)

	//test reserve book
	library.ReserveBook(1, 1)

	//test list available books
	library.ListAvailableBooks()

	//test list borrowed books
	library.ListBorrowedBooks()
}

func TestAsynchronousBorrowing(t *testing.T) {
	// Initialize library
	library := &services.Library{
		Books:   make(map[int]models.Books),
		Members: make(map[int]models.Member),
	}

	// Add test book
	book := models.Books{
		Id:     1,
		Title:  "Test Book",
		Author: "Test Author",
		Status: "available",
	}
	library.AddBook(book)

	// Add test member
	member := models.Member{
		Id:            1,
		Name:          "Test Member",
		BorrowedBooks: []models.Books{},
		ReservedBooks: []models.Books{},
	}
	library.AddMember(member)

	// Initialize reservation handler
	reservationHandler := concurrency.NewReservationHandler(library, 1)
	reservationHandler.Start()
	defer reservationHandler.Stop()

	// Submit a reservation
	reservationHandler.SubmitReservation(1, 1)

	// Wait for the reservation to be processed and book to be borrowed
	time.Sleep(3 * time.Second)

	// Verify the book status changed to borrowed
	if book, exists := library.Books[1]; exists {
		if book.Status != "borrowed" {
			t.Errorf("Expected book status to be 'borrowed', got '%s'", book.Status)
		}
	}

	// Verify the book is in member's borrowed books
	if member, exists := library.Members[1]; exists {
		found := false
		for _, borrowedBook := range member.BorrowedBooks {
			if borrowedBook.Id == 1 {
				found = true
				break
			}
		}
		if !found {
			t.Error("Book not found in member's borrowed books")
		}
	}

	// Test reservation of already reserved book
	err := library.ReserveBook(1, 1)
	if err == nil {
		t.Error("Expected error when reserving already reserved book, got nil")
	}
}

func TestConcurrentReservations(t *testing.T) {
	// Initialize library
	library := &services.Library{
		Books:   make(map[int]models.Books),
		Members: make(map[int]models.Member),
	}

	// Add test books
	for i := 1; i <= 5; i++ {
		book := models.Books{
			Id:     i,
			Title:  "Test Book " + string(rune(i)),
			Author: "Test Author",
			Status: "available",
		}
		library.AddBook(book)
	}

	// Add test members
	for i := 1; i <= 10; i++ {
		member := models.Member{
			Id:            i,
			Name:          "Test Member " + string(rune(i)),
			BorrowedBooks: []models.Books{},
			ReservedBooks: []models.Books{},
		}
		library.AddMember(member)
	}

	// Initialize reservation handler
	reservationHandler := concurrency.NewReservationHandler(library, 5)
	reservationHandler.Start()
	defer reservationHandler.Stop()

	// Simulate concurrent reservations
	var wg sync.WaitGroup
	concurrentReservations := 20 // Number of concurrent reservation attempts

	// Start concurrent reservation attempts
	for i := 0; i < concurrentReservations; i++ {
		wg.Add(1)
		go func(reservationID int) {
			defer wg.Done()
			// Use modulo to cycle through books and members
			memberID := (reservationID % 10) + 1
			bookID := (reservationID % 5) + 1
			reservationHandler.SubmitReservation(memberID, bookID)
		}(i)
	}

	// Wait for all reservations to be submitted
	wg.Wait()

	// Give some time for reservations to be processed
	time.Sleep(6 * time.Second)

	// Verify results
	for bookID, book := range library.Books {
		if book.Status == "reserved" {
			// Check if the book is in any member's reserved books
			found := false
			for _, member := range library.Members {
				for _, reservedBook := range member.ReservedBooks {
					if reservedBook.Id == bookID {
						found = true
						break
					}
				}
				if found {
					break
				}
			}
			if !found {
				t.Errorf("Book %d is marked as reserved but not found in any member's reserved books", bookID)
			}
		}
	}

	// Verify that no book is reserved by multiple members
	for bookID, book := range library.Books {
		if book.Status == "reserved" {
			reservationCount := 0
			for _, member := range library.Members {
				for _, reservedBook := range member.ReservedBooks {
					if reservedBook.Id == bookID {
						reservationCount++
					}
				}
			}
			if reservationCount > 1 {
				t.Errorf("Book %d is reserved by multiple members (%d)", bookID, reservationCount)
			}
		}
	}
}
