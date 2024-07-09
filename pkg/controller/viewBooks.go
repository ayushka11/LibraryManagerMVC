package controller

import (
	"fmt"
	"net/http"

	"github.com/ayushka11/LibraryManagerMVC/pkg/models"
	"github.com/ayushka11/LibraryManagerMVC/pkg/types"
	"github.com/ayushka11/LibraryManagerMVC/pkg/views"
)

func ViewBooks(writer http.ResponseWriter, request *http.Request) {
	books, err := models.GetBooks()
	if err != nil {
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}

	files := views.ViewFileNames()
	t := views.AdminRender(files.ViewBooks)

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