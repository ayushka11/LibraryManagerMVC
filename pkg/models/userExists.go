package models

import (
	"database/sql"
	"github.com/ayushka11/LibraryManagerMVC/pkg/types"
)

func UserExists(db *sql.DB ,username string) (bool, types.User, error) {
	var user types.User

	err := db.QueryRow("SELECT * FROM users WHERE username = $1", username).Scan(&user.UserId, &user.UserName, &user.IsAdmin, &user.Password, &user.AdminRequestStatus)
	if err != nil {
		return false, user, err
	}
	return true, user, nil
}