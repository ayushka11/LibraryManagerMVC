package models

import (

)

func AddUser(username string, password string) error {
	db, err := Connection()
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", username, password)
	if err != nil {
		return err
	}

	var id int
	err = db.QueryRow("SELECT id FROM users WHERE username = $1", username).Scan(&id)
	if err != nil {
		return err
	}

	if (id == 1) {
		_, err = db.Exec("UPDATE users SET is_admin = true WHERE id = 1")
		if err != nil {
			return err
		}
	}

	return nil
}