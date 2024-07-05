package models

import (
	"fmt"
)

func RequestAdmin(userid int) (string, error) {
	db, err := Connection()
	if err != nil {
		return "", err
	}

	requestquery := `UPDATE users SET admin_request_status = "pending" WHERE id = ?`
	_, error := db.Exec(requestquery, userid)
	if error != nil {
		fmt.Println(error)
		return "", err
	}

	return "Admin request sent successfully", nil
}