package main

import (
	"Library_Management/concurrency"
	"Library_Management/controllers"
	"Library_Management/models"
	"Library_Management/services"
	"fmt"
)

func main() {
	library := &services.Library{
		Books:   make(map[int]models.Books),
		Members: make(map[int]models.Member),
	}


	reservationHandler := concurrency.NewReservationHandler(library, 5)
	reservationHandler.Start()
	defer reservationHandler.Stop()

	fmt.Println("========== Library Management System ==========")
	controllers.BookController(library, reservationHandler)
}
