package models

import (

)

func RequestCheckout(bookid int, userid int) (string, error) {
	db, err := Connection()
	if err != nil {
		return "", err
	}

	requestquery := `INSERT INTO checkouts (book_id, user_id, status, type) VALUES (?, ?, ?, ?)`
	_, error := db.Exec(requestquery, bookid, userid, "pending", "checkout")
	if error != nil {
		return "", err
	}

	return "Checkout request sent successfully", nil
}

func RequestCheckin(bookid int, userid int) (string, error) {
	db, err := Connection()
	if err != nil {
		return "", err
	}

	requestquery := `INSERT INTO checkouts (book_id, user_id, status, type) VALUES (?, ?, ?, ?)`
	_, error := db.Exec(requestquery, bookid, userid, "pending", "checkin")
	if error != nil {
		return "", err
	}

	return "Checkin request sent successfully", nil
}