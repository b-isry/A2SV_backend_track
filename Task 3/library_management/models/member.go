package models

type Member struct {
	Id            int
	Name          string
	BorrowedBooks []Books
	ReservedBooks []Books
}
