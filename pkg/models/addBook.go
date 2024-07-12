package models

import (
	"github.com/ayushka11/LibraryManagerMVC/pkg/types"
	"database/sql"
)

func AddBook(title string, author string, quantity int) (string, error){
	db, err := Connection()
	if err != nil {
		return "", err
	}

	bookExists, book, err := bookExists(title, author)
	if err != nil {
		return "", err
	}

	if (bookExists) {
		totalQuantity := quantity + book.Quantity
		totalAvailable := quantity + book.Available

		updatequery := `UPDATE books SET quantity = ?, available = ? WHERE title = ?;`
		_, err := db.Exec(updatequery, totalQuantity, totalAvailable, book.Title)
		if err != nil {
			return "", err
		}
		return "Quantity updates of book " + title, nil
	} else {
		insertquery := `INSERT INTO books(title, author, quantity, available) VALUES (?, ?, ?, ?);`
		_, err := db.Exec(insertquery, title, author, quantity, quantity)
		if err != nil {
			return "", err
		}
		return "Added the book " + title, nil
	}
}

func bookExists(title string, author string) (bool, types.Book, error) {
	var book types.Book

	db, err := Connection()
	if err != nil {
		return false, book, err
	}

	selectquery := `SELECT * FROM books WHERE title = ? AND author = ?;`
	err = db.QueryRow(selectquery, title, author).Scan(&book.BookId, &book.Title, &book.Author, &book.Available, &book.Quantity)
	if err != nil {
		if err != sql.ErrNoRows {
			return false, book, err
		}
		return false, book, nil
	}
	return true, book, nil
}

func RemoveBooks(id int, remQuantity int) (string, error) {
	db, err := Connection()
	if err != nil {
		return "", err
	}

	checkquery := `SELECT available, quantity FROM books WHERE id = ?;`
	var available, oldQuantity int
	err = db.QueryRow(checkquery, id).Scan(&available, &oldQuantity)
	if err != nil {
		return "", err
	}
	if (remQuantity > available) {
		return "Not enough books available", nil
	}

	newQuantity := oldQuantity - remQuantity
	newAvailable := available - remQuantity

	updatequery1 := `UPDATE books SET available = ? WHERE id = ?;`
	_, err = db.Exec(updatequery1, newAvailable, id)
	if err != nil {
		return "", err
	}
	updatequery2 := `UPDATE books SET quantity = ? WHERE id = ?;`
	_, err = db.Exec(updatequery2, newQuantity, id)
	if err != nil {
		return "", err
	}
	return "Books removed successfully", nil
}

func DeleteBook(id int) (string, error) {
	db, err := Connection()
	if err != nil {
		return "", err
	}

	checkPending := `SELECT COUNT(*) FROM checkouts WHERE book_id = ? AND status = 'pending';`
	var pending int
	err = db.QueryRow(checkPending, id).Scan(&pending)
	if err != nil {
		return "", err
	}
	if (pending > 0) {
		return "Cannot delete book with pending requests", nil
	}

	checkAvailable := `SELECT available and quantity FROM books WHERE id = ?;`
	var available, quantity int
	err = db.QueryRow(checkAvailable, id).Scan(&available, &quantity)
	if err != nil {
		return "", err
	}
	if (available != quantity) {
		return "Cannot delete book with pending requests", nil
	}

	deletequery := `DELETE FROM books WHERE id = ?;`
	_, error := db.Exec(deletequery, id)
	if error != nil {
		return "", err
	}

	return "Book deleted successfully", nil
}