package controllers

import (
	"Library_Management/concurrency"
	"Library_Management/models"
	"Library_Management/services"
	"fmt"
)

func BookController(library *services.Library, reservationHandler *concurrency.ReservationHandler) {
	fmt.Println("1. Add Book")
	fmt.Println("2. Add Member")
	fmt.Println("3. Remove Book")
	fmt.Println("4. Borrow Book")
	fmt.Println("5. Return Book")
	fmt.Println("6. List Available Books")
	fmt.Println("7. List Borrowed Books")
	fmt.Println("8. Reserve Book")
	fmt.Println("9. Exit")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		var book models.Books
		fmt.Println("Enter book ID: ")
		fmt.Scan(&book.Id)
		fmt.Println("Enter book title: ")
		fmt.Scan(&book.Title)
		fmt.Println("Enter book author: ")
		fmt.Scan(&book.Author)
		book.Status = "available"
		library.AddBook(book)
	case 2:
		var member models.Member
		fmt.Println("Enter member ID: ")
		fmt.Scan(&member.Id)
		fmt.Println("Enter member name: ")
		fmt.Scan(&member.Name)
		member.BorrowedBooks = []models.Books{}
		member.ReservedBooks = []models.Books{}
		library.AddMember(member)
	case 3:
		var bookId int
		fmt.Println("Enter book id: ")
		fmt.Scan(&bookId)
		library.RemoveBook(bookId)
	case 4:
		var memberID int
		var bookID int
		fmt.Println("Enter member id: ")
		fmt.Scan(&memberID)
		fmt.Println("Enter book id: ")
		fmt.Scan(&bookID)
		library.BorrowBook(memberID, bookID)
	case 5:
		var memberID int
		var bookID int
		fmt.Println("Enter member id: ")
		fmt.Scan(&memberID)
		fmt.Println("Enter book id: ")
		fmt.Scan(&bookID)
		library.ReturnBook(memberID, bookID)
	case 6:
		fmt.Println("Available books:")
		library.ListAvailableBooks()
	case 7:
		fmt.Println("Borrowed books: ")
		library.ListBorrowedBooks()
	case 8:
		var memberID int
		var bookID int
		fmt.Println("Enter member id: ")
		fmt.Scan(&memberID)
		fmt.Println("Enter book id: ")
		fmt.Scan(&bookID)

		reservationHandler.SubmitReservation(memberID, bookID)
		fmt.Println("Reservation request submitted")
	case 9:
		fmt.Println("Thank you for using the library management system")
		return
	default:
		fmt.Println("Invalid choice")
	}

	fmt.Println("Do you want to continue? (y/n)")
	var continueChoice string
	fmt.Scan(&continueChoice)
	if continueChoice == "y" {
		BookController(library, reservationHandler)
	} else {
		fmt.Println("Thank you for using the library management system")
	}
}
