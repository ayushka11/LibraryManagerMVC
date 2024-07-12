package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ayushka11/LibraryManagerMVC/pkg/models"
	"github.com/ayushka11/LibraryManagerMVC/pkg/types"
	"github.com/ayushka11/LibraryManagerMVC/pkg/views"
)

func AddBookPage (writer http.ResponseWriter, request *http.Request) {
	files := views.ViewFileNames()
	t := views.AdminRender(files.AddBook)
	t.Execute(writer, nil)
}

func AddBook (writer http.ResponseWriter, request *http.Request) {
	title := request.FormValue("title")
	author := request.FormValue("author")
	quantityStr := request.FormValue("quantity")

	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		fmt.Println(err)
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}

	message, err := models.AddBook(title, author, quantity)
	if err != nil {
		fmt.Println(err)
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}

	data := types.PgMessage{Message: message}

	files := views.ViewFileNames()
	t := views.AdminRender(files.AdminHome)
	t.Execute(writer, data)
}

func RemoveBooks(writer http.ResponseWriter, Request *http.Request) {
	idstr := Request.FormValue("id")
	removequantitystr := Request.FormValue("removeQuantity")

	id, err := strconv.Atoi(idstr)
	if err != nil {
		fmt.Println(err)
		http.Redirect(writer, Request, "/500", http.StatusSeeOther)
		return
	}

	removequantity, err := strconv.Atoi(removequantitystr)
	if err != nil {
		fmt.Println(err)
		http.Redirect(writer, Request, "/500", http.StatusSeeOther)
		return
	}

	message, err := models.RemoveBooks(id, removequantity)
	if err != nil {
		fmt.Println(err)
		http.Redirect(writer, Request, "/500", http.StatusSeeOther)
		return
	}

	data := types.PgMessage{Message: message}

	files := views.ViewFileNames()
	t := views.AdminRender(files.AdminHome)
	t.Execute(writer, data)
}

func DeleteBook(writer http.ResponseWriter, Request *http.Request) {
	idstr := Request.FormValue("id")

	id, err := strconv.Atoi(idstr)
	if err != nil {
		fmt.Println(err)
		http.Redirect(writer, Request, "/500", http.StatusSeeOther)
		return
	}

	message, err := models.DeleteBook(id)
	if err != nil {
		fmt.Println(err)
		http.Redirect(writer, Request, "/500", http.StatusSeeOther)
		return
	}

	data := types.PgMessage{Message: message}

	files := views.ViewFileNames()
	t := views.AdminRender(files.AdminHome)
	t.Execute(writer, data)
}