package views

import (
	"html/template"
	"github.com/ayushka11/LibraryManagerMVC/pkg/types"
)

func ViewFileNames() types.FileName {
	return types.FileName{
		AdminHome:           "adminHome",
		UserHome:            "userHome",
		Login:               "login",
		PageNotFound:        "pageNotFound",
		UnauthorizedAccess:  "unauthorized",
		InternalServerError: "internalServerError",
		Signup:              "signup",
		AddBook:             "addBook",
		ViewBooks:           "viewBooks",
		AvailableBooks:      "availableBooks",
		BookRequests:        "bookRequests",
		AdminRequests:       "adminRequests",
		ViewHistory:         "viewHistory",
	}
}

func Render(page string) *template.Template {
	file := "templates/" + page + ".html"
	temp := template.Must(template.ParseFiles(file))
	return temp
}