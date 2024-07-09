package controller

import (
	"net/http"
	"fmt"
	"strconv"
	"github.com/ayushka11/LibraryManagerMVC/pkg/models"
	"github.com/ayushka11/LibraryManagerMVC/pkg/types"
)

func CheckinBook(writer http.ResponseWriter, Request *http.Request) {
	idstr := Request.FormValue("bookId")

	bookId, err := strconv.Atoi(idstr)
	if err != nil {
		fmt.Println(err)
		http.Redirect(writer, Request, "/500", http.StatusSeeOther)
		return
	}

	userId, ok := Request.Context().Value(types.UserIdContextKey).(int)
    if !ok {
        http.Error(writer, "User not authenticated", http.StatusUnauthorized)
        return
    }

	message, err := models.RequestCheckin(bookId, userId)
	if err != nil {
		fmt.Println(err)
		http.Redirect(writer, Request, "/500", http.StatusSeeOther)
		return
	}

	showMessage(writer, Request, message)
}

func CheckoutBook(writer http.ResponseWriter, Request *http.Request) {
	idstr := Request.FormValue("bookId")

	bookId, err := strconv.Atoi(idstr)
	if err != nil {
		fmt.Println(err)
		http.Redirect(writer, Request, "/500", http.StatusSeeOther)
		return
	}

	userId, ok := Request.Context().Value(types.UserIdContextKey).(int)
    if !ok {
        http.Error(writer, "User not authenticated", http.StatusUnauthorized)
        return
    }

	message, err := models.RequestCheckout(bookId, userId)
	if err != nil {
		fmt.Println(err)
		http.Redirect(writer, Request, "/500", http.StatusSeeOther)
		return
	}

	showMessage(writer, Request, message)
}