package controller

import (
	"net/http"
	"fmt"
	"strconv"
	"github.com/ayushka11/LibraryManagerMVC/pkg/models"
	"github.com/ayushka11/LibraryManagerMVC/pkg/views"
	"github.com/ayushka11/LibraryManagerMVC/pkg/types"
)

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
	t := views.Render(files.AdminHome)
	t.Execute(writer, data)
}