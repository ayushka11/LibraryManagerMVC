package controller

import (
	"net/http"
	"fmt"
	"github.com/ayushka11/LibraryManagerMVC/pkg/models"
	"github.com/ayushka11/LibraryManagerMVC/pkg/types"
	"github.com/ayushka11/LibraryManagerMVC/pkg/views"
)

func ViewAdminRequests(writer http.ResponseWriter, request *http.Request) {
	adminRequests, err := models.GetAdminRequests()
	if err != nil {
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}

	files := views.ViewFileNames()
	t := views.AdminRender(files.AdminRequests)

	data := struct {
		AdminRequests []types.AdminRequest
	}{
		AdminRequests: adminRequests,
	}

	error := t.Execute(writer, data)
	if error != nil {
		fmt.Println(error)
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}
}