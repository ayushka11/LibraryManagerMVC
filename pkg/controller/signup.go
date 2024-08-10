package controller

import (
	"log"
	"net/http"
	"fmt"
	"github.com/ayushka11/LibraryManagerMVC/pkg/models"
	"github.com/ayushka11/LibraryManagerMVC/pkg/types"
	"github.com/ayushka11/LibraryManagerMVC/pkg/views"
	"github.com/ayushka11/LibraryManagerMVC/pkg/utils"

)

func Signup(writer http.ResponseWriter, request *http.Request) {
	files := views.ViewFileNames()
	t := views.Render(files.Signup)
	t.Execute(writer, nil)
}

func AddUser(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("AddUser")
	var RequestUser types.RequestUser
	RequestUser.UserName = request.FormValue("username")
	RequestUser.Password = request.FormValue("password")
	RequestUser.ConfirmPassword = request.FormValue("passwordC")

	if RequestUser.Password != RequestUser.ConfirmPassword {
		http.Redirect(writer, request, "/signup?message=passwords do not match", http.StatusSeeOther)
		return
	} else if RequestUser.UserName == "" || RequestUser.Password == "" || RequestUser.ConfirmPassword == "" {
		http.Redirect(writer, request, "/signup?message=all fields are mandatory", http.StatusSeeOther)
		return
	} else {
		db, err := utils.Connection()
		if err != nil {
			log.Println(err)
			http.Redirect(writer, request, "/500", http.StatusSeeOther)
			return
		}
		userExists, _, err := models.UserExists(db, RequestUser.UserName)
		if err != nil {
			log.Println(err)
			http.Redirect(writer, request, "/500", http.StatusSeeOther)
			return
		}

		if userExists {
			http.Redirect(writer, request, "/signup?message=user already exists", http.StatusSeeOther)
			return
		} else {
			password, err := utils.HashPassword(RequestUser.Password)
			if err != nil {
				log.Println(err)
				http.Redirect(writer, request, "/500", http.StatusSeeOther)
				return
			}
			err = models.AddUser(RequestUser.UserName, password)
			if err != nil {
				log.Println(err)
				http.Redirect(writer, request, "/500", http.StatusSeeOther)
				return
			}
		}

	}
	http.Redirect(writer, request, "/login", http.StatusSeeOther)
}