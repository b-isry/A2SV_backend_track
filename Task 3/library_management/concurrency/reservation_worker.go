package concurrency

import (
	"Library_Management/services"
	"fmt"
	"sync"
)

type ReservationHandler struct {
	library          *services.Library
	reservationQueue chan struct{ memberID, bookID int }
	workerCount      int
	wg               sync.WaitGroup
}

func NewReservationHandler(library *services.Library, workerCount int) *ReservationHandler {
	return &ReservationHandler{
		library:          library,
		reservationQueue: make(chan struct{ memberID, bookID int }, 100),
		workerCount:      workerCount,
	}
}

func (h *ReservationHandler) Start() {
	for i := 0; i < h.workerCount; i++ {
		h.wg.Add(1)
		go h.worker()
	}
}

func (h *ReservationHandler) Stop() {
	close(h.reservationQueue)
	h.wg.Wait()
}

func (h *ReservationHandler) SubmitReservation(memberID, bookID int) {
	h.reservationQueue <- struct{ memberID, bookID int }{memberID, bookID}
}

func (h *ReservationHandler) worker() {
	defer h.wg.Done()

	for reservation := range h.reservationQueue {
		err := h.library.ReserveBook(reservation.memberID, reservation.bookID)
		if err != nil {
			fmt.Printf("Failed to process reservation: %v\n", err)
		}
	}
}
