package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Bisrat/task-manager/data"
	"github.com/Bisrat/task-manager/router"
)

func main() {
	// Initialize MongoDB connection
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	if err := data.InitDB(mongoURI); err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer data.CloseDB()

	// Setup router
	r := router.SetupRouter()

	// Start server in a goroutine
	go func() {
		if err := r.Run(); err != nil {
			log.Printf("Server error: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
}
