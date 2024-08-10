package models

import (
	"fmt"
	"github.com/ayushka11/LibraryManagerMVC/pkg/utils"
)

func RequestAdmin(userid int) (string, error) {
	db, err := utils.Connection()
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