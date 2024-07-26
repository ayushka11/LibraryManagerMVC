package models

import (
	"testing"
	"time"
	"github.com/ayushka11/LibraryManagerMVC/pkg/models"
)

func TestCalculateFine(t *testing.T) {
	tests := []struct {
		dueDate     time.Time
		returnDate  time.Time
		expectedFine int
	}{
		{time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC), time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC), 0},
		{time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC), time.Date(2024, 7, 2, 0, 0, 0, 0, time.UTC), 5},
		{time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC), time.Date(2024, 7, 5, 0, 0, 0, 0, time.UTC), 20},
		{time.Date(2024, 7, 5, 0, 0, 0, 0, time.UTC), time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC), 0},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			fine := models.CalculateFine(tt.dueDate, tt.returnDate)
			if fine != tt.expectedFine {
				t.Errorf("CalculateFine(%v, %v) = %d; expected %d", tt.dueDate, tt.returnDate, fine, tt.expectedFine)
			}
		})
	}
}
