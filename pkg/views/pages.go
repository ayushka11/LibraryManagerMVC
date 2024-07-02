package views

import (
	"html/template"
	"net/http"
)

func Render(w http.ResponseWriter, tmpl string, data interface{}) error {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		return err
	}
	err = t.Execute(w, data)
	if err != nil {
		return err
	}
	return nil
}