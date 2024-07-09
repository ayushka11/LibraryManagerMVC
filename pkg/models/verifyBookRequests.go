package models

import (
	"fmt"
	"time"
)

func ApproveRequest(requestid int) (string, error) {
	db, err := Connection()
	if err != nil {
		return "", err
	}

	var requestType string
	err = db.QueryRow("SELECT type FROM checkouts WHERE id = ?", requestid).Scan(&requestType)
	if err != nil {
		return "", err
	}

	var updateQuery string
	if requestType == "checkout" {
		checkoutDate := time.Now()
		dueDate := checkoutDate.AddDate(0, 0, 14)
		fine := 0

		updateQuery = `UPDATE checkouts SET status = 'approved', checkout_date = ?, due_date = ?, fine = ? WHERE id = ?`

		_, err = db.Exec(updateQuery, checkoutDate, dueDate, fine, requestid)
		if err != nil {
			return "", err
		}

		getBookIDQuery := `SELECT book_id FROM checkouts WHERE id = ?`

		var bookID int
		err = db.QueryRow(getBookIDQuery, requestid).Scan(&bookID)
		if err != nil {
			return "", err
		}

		updateBooksQuery := `UPDATE books SET available = available - 1 WHERE id = ?`

		_, err = db.Exec(updateBooksQuery, bookID)
		if err != nil {
			return "", err
		}
	} else if requestType == "checkin" {
		returnDate := time.Now()
		
		getDueDateQuery := `SELECT due_date FROM checkouts WHERE id = ?`
		
		var dueDateStr string
		err = db.QueryRow(getDueDateQuery, requestid).Scan(&dueDateStr)
		if err != nil {
			return "", err
		}
		
		dueDate, err := time.Parse("2006-01-02", dueDateStr)
		if err != nil {
			return "", err
		}
	
		fine := calculateFine(dueDate, returnDate)
	
		updateQuery = `UPDATE checkouts SET status = 'approved', return_date = ?, fine = ? WHERE id = ?`
	
		_, err = db.Exec(updateQuery, returnDate, fine, requestid)
		if err != nil {
			return "", err
		}
		
	}

	return "Request updated successfully", nil
}

func calculateFine(dueDate time.Time, returnDate time.Time) int {
	fine := 0
	if returnDate.After(dueDate) {
		daysLate := returnDate.Sub(dueDate).Hours() / 24
		fine = int(daysLate) * 5 
	}
	return fine
}

func DeclineRequest(requestid int) (string, error) {
	db, err := Connection()
	if err != nil {
		return "", err
	}

	declinequery := `UPDATE checkouts SET status = 'rejected' WHERE id = ?`
	_, error := db.Exec(declinequery, requestid)
	if error != nil {
		fmt.Println(error)
		return "", err
	}

	return "Request declined successfully", nil
}

