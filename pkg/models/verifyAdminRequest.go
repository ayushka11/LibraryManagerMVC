package models

import (
	"github.com/ayushka11/LibraryManagerMVC/pkg/utils"
)

func VerifyAdminRequest (userId int, action rune) (string, error) {
	db, err := utils.Connection()
	if err != nil {
		return "", err
	}

	var status string
	if action == 'A' {
		status = "approved"
	} else {
		status = "rejected"
	}

	requestquery := `UPDATE users SET admin_request_status = ? WHERE id = ?`
	_, error := db.Exec(requestquery, status, userId)
	if error != nil {
		return "", err
	}

	if (status == "approved") {
		updateQuery := `UPDATE users SET isAdmin = true WHERE id = ?`
		_, err = db.Exec(updateQuery, userId)
		if err != nil {
			return "", err
		}
	}

	return "Admin request " + status, nil
}