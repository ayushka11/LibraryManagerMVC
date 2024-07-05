package controller

import (
	"net/http"
	"github.com/ayushka11/LibraryManagerMVC/pkg/views"
	"github.com/ayushka11/LibraryManagerMVC/pkg/types"
	"github.com/ayushka11/LibraryManagerMVC/pkg/models"
)

func UserHome (writer http.ResponseWriter, request *http.Request) {
	userId, ok := request.Context().Value(types.UserIdContextKey).(int)
    if !ok {
        http.Error(writer, "User not authenticated", http.StatusUnauthorized)
        return
    }
    
	checkouts, err := models.GetCheckedOutBooksByUser(userId)
    if err != nil {
        http.Redirect(writer, request, "/500", http.StatusSeeOther)
        return
    }

	files := views.ViewFileNames()
	t := views.Render(files.UserHome)

	data := struct {
        Message string
        Checkouts   []types.Checkouts
    }{
        Message: "", 
        Checkouts: checkouts,
    }

    err = t.Execute(writer, data)
    if err != nil {
        http.Redirect(writer, request, "/500", http.StatusSeeOther)
        return
    }
}

