package models

import (
	"github.com/ayushka11/LibraryManagerMVC/pkg/types"
	"fmt"
	"github.com/ayushka11/LibraryManagerMVC/pkg/utils"
)

func GetCheckedOutBooksByUser(id int) ([]types.Checkouts, error){
	db, err := utils.Connection()
	if err != nil {
		return nil, err
	}

	checkoutsquery := `
		SELECT b.id, b.title, b.author, c.checkout_date, c.due_date
		FROM books b
		JOIN checkouts c ON b.id = c.book_id
		WHERE c.user_id = ? AND ((c.status = 'approved' AND c.type = 'checkout') OR (c.status = 'pending' AND c.type = 'checkin'));
	`
	rows, err := db.Query(checkoutsquery, id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var checkouts []types.Checkouts
	for rows.Next() {
		var checkout types.Checkouts
		err := rows.Scan(&checkout.BookId, &checkout.Title, &checkout.Author, &checkout.CheckoutDate, &checkout.DueDate)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		checkouts = append(checkouts, checkout)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return checkouts, nil
}