package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/ayushka11/LibraryManagerMVC/pkg/models"
	"github.com/ayushka11/LibraryManagerMVC/pkg/types"
	"github.com/ayushka11/LibraryManagerMVC/pkg/views"
)

func ApproveAdminRequest(writer http.ResponseWriter, request *http.Request) {
	userIdStr := request.FormValue("userId")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		fmt.Println(err)
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}

	message, err := models.VerifyAdminRequest(userId, 'A')
	if err != nil {
		fmt.Println(err)
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}

	data := types.PgMessage{Message: message}

	files := views.ViewFileNames()
	t := views.Render(files.AdminHome)
	t.Execute(writer, data)
}

func RejectAdminRequest(writer http.ResponseWriter, request *http.Request) {
	userIdStr := request.FormValue("userId")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		fmt.Println(err)
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}

	message, err := models.VerifyAdminRequest(userId, 'R')
	if err != nil {
		fmt.Println(err)
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}

	data := types.PgMessage{Message: message}

	files := views.ViewFileNames()
	t := views.Render(files.AdminHome)
	t.Execute(writer, data)
}