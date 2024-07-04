package controller

import (
	"net/http"
	"github.com/ayushka11/LibraryManagerMVC/pkg/views"
)

func UserHome (writer http.ResponseWriter, request *http.Request) {
	files := views.ViewFileNames()
	t := views.Render(files.UserHome)
	t.Execute(writer, nil)
}
