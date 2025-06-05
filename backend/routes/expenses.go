package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"expense-tracker-backend/models"
	"expense-tracker-backend/utils"
)

// GetTokenFromHeader extracts the token from the Authorization header
func GetTokenFromHeader(r *http.Request) (string, error) {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		return "", fmt.Errorf("missing authorization header")
	}
	parts := strings.Split(auth, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", fmt.Errorf("invalid authorization header format")
	}
	return parts[1], nil
}

// SyncHandler handles the /sync endpoint
func SyncHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	token, err := GetTokenFromHeader(r)
	if err != nil {
		http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
		return
	}

	var newExpenses []models.Expense
	if err := json.NewDecoder(r.Body).Decode(&newExpenses); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	lock := utils.GetFileLock(token)
	lock.Lock()
	defer lock.Unlock()

	if err := utils.SaveExpenses(token, newExpenses); err != nil {
		http.Error(w, "Failed to save expenses: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

// GetExpensesHandler handles the /expenses endpoint
func GetExpensesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	token, err := GetTokenFromHeader(r)
	if err != nil {
		http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
		return
	}

	lock := utils.GetFileLock(token)
	lock.Lock()
	defer lock.Unlock()

	expenses, err := utils.LoadExpenses(token)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(expenses)
}

// DeleteExpenseHandler handles the /expenses/delete endpoint
func DeleteExpenseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	token, err := GetTokenFromHeader(r)
	if err != nil {
		http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing expense ID", http.StatusBadRequest)
		return
	}

	lock := utils.GetFileLock(token)
	lock.Lock()
	defer lock.Unlock()

	expenses, err := utils.LoadExpenses(token)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Filter out the expense with the given ID
	var updatedExpenses []models.Expense
	for _, exp := range expenses {
		if fmt.Sprintf("%d", exp.ID) != id {
			updatedExpenses = append(updatedExpenses, exp)
		}
	}

	if err := utils.SaveExpenses(token, updatedExpenses); err != nil {
		http.Error(w, "Failed to save expenses: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}
