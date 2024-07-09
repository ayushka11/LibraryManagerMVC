package controller

import (
	"net/http"
	"fmt"
	"github.com/ayushka11/LibraryManagerMVC/pkg/models"
	"github.com/ayushka11/LibraryManagerMVC/pkg/types"
	"github.com/ayushka11/LibraryManagerMVC/pkg/views"
)

func AvailableBooks(writer http.ResponseWriter, request *http.Request) {
	books, err := models.GetAvailableBooks()
	if err != nil {
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}

	files := views.ViewFileNames()
	t := views.UserRender(files.AvailableBooks)

	data := struct {
		Books   []types.Book
	}{
		Books:   books,
	}

	error := t.Execute(writer, data)
	if error != nil {
		fmt.Println(error)
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}
}