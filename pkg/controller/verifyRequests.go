package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/ayushka11/LibraryManagerMVC/pkg/models"
	"github.com/ayushka11/LibraryManagerMVC/pkg/types"
	"github.com/ayushka11/LibraryManagerMVC/pkg/views"
)

func ApproveRequest(writer http.ResponseWriter, request *http.Request){
	requestidstr := request.FormValue("request_id")
	requestid, err := strconv.Atoi(requestidstr)

	if err != nil {
		fmt.Println(err)
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}

	message, err := models.ApproveRequest(requestid)
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

func DeclineRequest(writer http.ResponseWriter, request *http.Request){
	requestidstr := request.FormValue("request_id")
	requestid, err := strconv.Atoi(requestidstr)

	if err != nil {
		fmt.Println(err)
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}

	message, err := models.DeclineRequest(requestid)
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