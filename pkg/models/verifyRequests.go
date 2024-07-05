package models

import (

)

func ApproveRequest(requestid int) (string, error) {
	db, err := Connection()
	if err != nil {
		return "", err
	}

	approvequery := `UPDATE checkouts SET status = 'approved' WHERE id = ?`
	_, error := db.Exec(approvequery, requestid)
	if error != nil {
		return "", err
	}

	return "Request approved successfully", nil
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