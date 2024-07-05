package models

import (
	"fmt"
	"github.com/ayushka11/LibraryManagerMVC/pkg/types"
)

func GetAdminRequests() ([]types.AdminRequest, error){
	db, err := Connection()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT id, username FROM users WHERE admin_request_status = 'pending'")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var adminRequests []types.AdminRequest
	for rows.Next() {
		var adminRequest types.AdminRequest
		err := rows.Scan(&adminRequest.UserId, &adminRequest.Username)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		adminRequests = append(adminRequests, adminRequest)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return adminRequests, nil
}