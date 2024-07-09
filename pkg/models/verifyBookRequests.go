package models

import (
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
	} 

	return "Request updated successfully", nil
}

func DeclineRequest(requestid int) (string, error) {
	db, err := Connection()
	if err != nil {
		return "", err
	}

	approvequery := `UPDATE checkouts SET status = 'declined' WHERE id = ?`
	_, error := db.Exec(approvequery, requestid)
	if error != nil {
		return "", err
	}

	return "Request declined successfully", nil
}

