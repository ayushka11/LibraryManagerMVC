package models

import (
	"fmt"
	"github.com/ayushka11/LibraryManagerMVC/pkg/types"

)

func GetPendingRequests() ([]types.BookRequest, error){
	db, err := Connection()
	if err != nil {
		return nil, err
	}
	requestquery := `SELECT c.id, c.user_id, c.book_id, c.checkout_date, c.due_date, c.type, b.title, u.username
		FROM checkouts c
		JOIN books b on b.id = c.book_id
		JOIN users u on u.id = c.user_id 
		WHERE c.status = 'pending'
	`

	rows, err := db.Query(requestquery)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var bookRequests []types.BookRequest
	for rows.Next() {
		var bookRequest types.BookRequest
		err := rows.Scan(&bookRequest.RequestId, &bookRequest.BookId, &bookRequest.UserId, &bookRequest.Type, &bookRequest.CheckoutDate, &bookRequest.DueDate, &bookRequest.Book, &bookRequest.Username)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		bookRequests = append(bookRequests, bookRequest)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return bookRequests, nil

}