package controller

import (
	"log"
	"net/http"
	"fmt"
	"github.com/ayushka11/LibraryManagerMVC/pkg/models"
	"github.com/ayushka11/LibraryManagerMVC/pkg/types"
	"github.com/ayushka11/LibraryManagerMVC/pkg/views"
)

func Signup(writer http.ResponseWriter, request *http.Request) {
	//fmt.Println("Signup")
	views.Render(writer, "signup.html", nil)
}

func AddUser(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("AddUser")
	var RequestUser types.RequestUser
	RequestUser.UserName = request.FormValue("username")
	RequestUser.Password = request.FormValue("password")
	RequestUser.ConfirmPassword = request.FormValue("passwordC")

	if RequestUser.Password != RequestUser.ConfirmPassword {
		views.Render(writer, "signup.html", "Password and Confirm Password do not match")
		return
	} else if (RequestUser.UserName == "" || RequestUser.Password == "" || RequestUser.ConfirmPassword == "") {
		views.Render(writer, "signup.html", "All fields are mandatory")
		return
	} else {
		db, err := models.Connection()
		if err != nil {
			log.Println(err)
			views.Render(writer, "signup.html", "Internal Server Error")
			return
		}
		userExists, _, err := models.UserExists(db, RequestUser.UserName)
		if err != nil {
			log.Println(err)
			views.Render(writer, "signup.html", "Internal Server Error")
			return
		}

		if userExists {
			views.Render(writer, "signup.html", "User already exists")
			return
		} else {
			password, err := models.HashPassword(RequestUser.Password)
			if err != nil {
				log.Println(err)
				views.Render(writer, "signup.html", "Internal Server Error")
				return
			}
			err = models.AddUser(RequestUser.UserName, password)
			if err != nil {
				log.Println(err)
				views.Render(writer, "signup.html", "Internal Server Error")
				return
			}
			//fmt.Println("User Added")
		}

	}
	http.Redirect(writer, request, "/login", http.StatusSeeOther)
}