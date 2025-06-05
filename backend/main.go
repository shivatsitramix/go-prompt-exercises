package main

import (
	"log"
	"net/http"

	"expense-tracker-backend/middleware"
	"expense-tracker-backend/routes"
	"expense-tracker-backend/utils"
)

func main() {
	// Ensure data directory exists
	if err := utils.EnsureDataDir(); err != nil {
		log.Fatal("Failed to create data directory:", err)
	}

	// Create a new mux router
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/sync", routes.SyncHandler)
	mux.HandleFunc("/expenses", routes.GetExpensesHandler)
	mux.HandleFunc("/expenses/delete", routes.DeleteExpenseHandler)

	// Apply CORS middleware
	handler := middleware.CORSMiddleware(mux)

	// Start the server
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
