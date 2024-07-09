package models

import (
	"fmt"
	"github.com/ayushka11/LibraryManagerMVC/pkg/types"
)

func GetBooks() ([]types.Book, error) {
	db, err := Connection()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var books []types.Book
	for rows.Next() {
		var book types.Book
		err := rows.Scan(&book.BookId, &book.Title, &book.Author, &book.Quantity, &book.Available)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		books = append(books, book)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return books, nil
}

func GetAvailableBooks () ([]types.Book, error){
	db, err := Connection()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM books WHERE available > 0")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var books []types.Book
	for rows.Next() {
		var book types.Book
		err := rows.Scan(&book.BookId, &book.Title, &book.Author, &book.Quantity, &book.Available)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		books = append(books, book)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return books, nil
}