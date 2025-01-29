package controllers

import (
	"go_web/models"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("views/index.html"))


func ContactIndex(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Contacts []models.Contact
	}{
		Contacts: models.GetAllContacts(),
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func ContactCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		email := r.FormValue("email")
		role := r.FormValue("role")

		newContact := models.Contact{
			Name:  name,
			Email: email,
			Role:  role,
		}

		models.AddContact(newContact)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
