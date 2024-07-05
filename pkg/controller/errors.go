package controller

import (
	"net/http"
	"github.com/ayushka11/LibraryManagerMVC/pkg/views"
)

func Unauthorized(writer http.ResponseWriter, request *http.Request) {
	files := views.ViewFileNames()
	t := views.Render(files.UnauthorizedAccess)
	t.Execute(writer, nil)
}

func InternalServerError(writer http.ResponseWriter, request *http.Request) {
	files := views.ViewFileNames()
	t := views.Render(files.InternalServerError)
	t.Execute(writer, nil)
}