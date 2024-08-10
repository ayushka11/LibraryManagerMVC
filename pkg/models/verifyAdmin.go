package models

import (
	"database/sql"
	"github.com/ayushka11/LibraryManagerMVC/pkg/utils"
)

func VerifyAdmin(userId int) (bool, error) {
	db, err := utils.Connection()
	if err != nil {
		return false, err
	}

	var isAdmin bool
	err = db.QueryRow("SELECT isAdmin FROM users WHERE id = ?", userId).Scan(&isAdmin)
	if err != nil {
		if err != sql.ErrNoRows {
			return false , err
		}
		return isAdmin, nil
	}
	return isAdmin, nil
}