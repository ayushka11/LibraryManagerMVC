package controller

import (
	"net/http"
	"fmt"
	"github.com/ayushka11/LibraryManagerMVC/pkg/models"
	"github.com/ayushka11/LibraryManagerMVC/pkg/types"
)

func RequestAdmin(writer http.ResponseWriter, Request *http.Request) {
	userId, ok := Request.Context().Value(types.UserIdContextKey).(int)
    if !ok {
        http.Error(writer, "User not authenticated", http.StatusUnauthorized)
        return
    }

	message, err := models.RequestAdmin(userId)
	if err != nil {
		fmt.Println(err)
		http.Redirect(writer, Request, "/500", http.StatusSeeOther)
		return
	}

	showMessage(writer, Request, message)
}