package models

import (
	"database/sql"

	"github.com/ayushka11/LibraryManagerMVC/pkg/types"
)

func UserExists(db *sql.DB ,username string) (bool, types.User, error) {
	var user types.User

	err := db.QueryRow("SELECT * FROM users WHERE username = ?", username).Scan(&user.UserId, &user.UserName, &user.Password, &user.IsAdmin, &user.AdminRequestStatus)
	if err != nil {
		if err != sql.ErrNoRows {
			return false, user, err
		}
        return false, user, nil
	}
	return true, user, nil
}