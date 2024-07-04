package models

func DeleteBook(id int) (string, error) {
	db, err := Connection()
	if err != nil {
		return "", err
	}

	deletequery := `DELETE FROM books WHERE id = ?;`
	_, error := db.Exec(deletequery, id)
	if error != nil {
		return "", err
	}

	return "Book deleted successfully", nil
}