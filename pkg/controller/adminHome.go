package controller

import (
	"net/http"
	"github.com/ayushka11/LibraryManagerMVC/pkg/views"
)

func AdminHome (writer http.ResponseWriter, request *http.Request) {
	files := views.ViewFileNames()
	t := views.Render(files.AdminHome)
	t.Execute(writer, nil)
}
