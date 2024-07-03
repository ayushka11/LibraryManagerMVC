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

	bookExists, book, err := bookExists(title)
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

func bookExists(title string) (bool, types.Book, error) {
	var book types.Book

	db, err := Connection()
	if err != nil {
		return false, book, err
	}

	selectquery := `SELECT * FROM books WHERE title = ?;`
	err = db.QueryRow(selectquery, title).Scan(&book.BookId, &book.Title, &book.Author, &book.Available, &book.Quantity)
	if err != nil {
		if err != sql.ErrNoRows {
			return false, book, err
		}
		return false, book, nil
	}
	return true, book, nil
}