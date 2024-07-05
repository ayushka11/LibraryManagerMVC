package controller

import (
	"fmt"
	"net/http"
	"github.com/ayushka11/LibraryManagerMVC/pkg/models"
	"github.com/ayushka11/LibraryManagerMVC/pkg/types"
	"github.com/ayushka11/LibraryManagerMVC/pkg/views"
)

func ViewHistory(writer http.ResponseWriter, request *http.Request) {
	history, err := models.GetHistory()
	if err != nil {
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}

	files := views.ViewFileNames()
	t := views.Render(files.ViewHistory)

	data := struct {
		History   []types.History
	}{
		History:   history,
	}

	error := t.Execute(writer, data)
	if error != nil {
		fmt.Println(error)
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}
}