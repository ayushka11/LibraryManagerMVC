package controller

import (
	"net/http"
	"html/template"
)

func showMessage(writer http.ResponseWriter, request *http.Request, message string) {
    t, err := template.ParseFiles("templates/message.html")
    if err != nil {
        http.Redirect(writer, request, "/500", http.StatusSeeOther)
        return
    }

    data := struct {
        Message string
    }{
        Message: message,
    }

    err = t.Execute(writer, data)
    if err != nil {
        http.Redirect(writer, request, "/500", http.StatusSeeOther)
        return
    }
}
