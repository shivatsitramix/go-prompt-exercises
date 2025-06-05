package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"expense-tracker-backend/models"
)

var (
	fileLocks = make(map[string]*sync.Mutex)
	lockMutex sync.Mutex
)

// GetFileLock returns a mutex for the given token
func GetFileLock(token string) *sync.Mutex {
	lockMutex.Lock()
	defer lockMutex.Unlock()

	if lock, exists := fileLocks[token]; exists {
		return lock
	}

	lock := &sync.Mutex{}
	fileLocks[token] = lock
	return lock
}

// SaveExpenses saves expenses to a JSON file for the given token
func SaveExpenses(token string, expenses []models.Expense) error {
	path := filepath.Join("data", fmt.Sprintf("data_%s.json", token))
	data, err := json.MarshalIndent(expenses, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// LoadExpenses loads expenses from a JSON file for the given token
func LoadExpenses(token string) ([]models.Expense, error) {
	path := filepath.Join("data", fmt.Sprintf("data_%s.json", token))
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return []models.Expense{}, nil
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var expenses []models.Expense
	if err := json.Unmarshal(data, &expenses); err != nil {
		return nil, err
	}
	return expenses, nil
}

// EnsureDataDir ensures the data directory exists
func EnsureDataDir() error {
	return os.MkdirAll("data", 0755)
}
