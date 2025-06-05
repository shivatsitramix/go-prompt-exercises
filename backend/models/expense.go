package models

import (
	"encoding/json"
	"time"
)

type Expense struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Amount   float64   `json:"amount"`
	Category string    `json:"category"`
	Date     time.Time `json:"date"`
}

// UnmarshalJSON implements custom JSON unmarshaling for Expense
func (e *Expense) UnmarshalJSON(data []byte) error {
	type Alias Expense
	aux := &struct {
		Date string `json:"date"`
		*Alias
	}{
		Alias: (*Alias)(e),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	// Parse the date string
	t, err := time.Parse(time.RFC3339, aux.Date)
	if err != nil {
		// Try parsing without timezone
		t, err = time.Parse("2006-01-02T15:04:05.999", aux.Date)
		if err != nil {
			return err
		}
	}
	e.Date = t
	return nil
}

// MarshalJSON implements custom JSON marshaling for Expense
func (e Expense) MarshalJSON() ([]byte, error) {
	type Alias Expense
	return json.Marshal(&struct {
		Date string `json:"date"`
		*Alias
	}{
		Date:  e.Date.Format(time.RFC3339),
		Alias: (*Alias)(&e),
	})
}
