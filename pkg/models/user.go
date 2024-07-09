package models

import (
	"fmt"
	"database/sql"

	"github.com/ayushka11/LibraryManagerMVC/pkg/types"
)

func AddUser(username string, password string) error {
	db, err := Connection()
	if err != nil {
		return err
	}

	isAdmin := false

	_, err = db.Exec("INSERT INTO users (username, password, isAdmin) VALUES (?, ?, ?)", username, password, isAdmin)
	if err != nil {
		return err
	}
	fmt.Println("User added successfully")
	var id int
	err = db.QueryRow("SELECT id FROM users WHERE username = ?", username).Scan(&id)
	if err != nil {
		return err
	}

	if (id == 1) {
		_, err = db.Exec("UPDATE users SET isAdmin = true WHERE id = 1")
		if err != nil {
			return err
		}
	}

	return nil
}

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