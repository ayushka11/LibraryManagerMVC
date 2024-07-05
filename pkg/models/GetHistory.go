package models

import (
	"fmt"
	"github.com/ayushka11/LibraryManagerMVC/pkg/types"
)

func GetHistory() ([]types.History, error) {
	db, err := Connection()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT c.book_id, b.title, b.author, c.checkout_date, c.due_date, c.return_date, c.fine, c.status, c.type FROM checkouts c JOIN books b ON c.book_id = b.id;")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var history []types.History
	for rows.Next() {
		var h types.History
		err := rows.Scan(&h.BookId, &h.Title, &h.Author, &h.CheckOutDate, &h.DueDate, &h.ReturnDate, &h.Fine, &h.Status, &h.Type)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		history = append(history, h)
	}

	return history, nil
}