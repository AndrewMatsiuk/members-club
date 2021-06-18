package handlers

import (
	"errors"
	"html/template"
	"net/http"
	"net/mail"
	"path"
	"regexp"

	"members-club/services"
)

func New(service Service) *handler {
	return &handler{
		service: service,
	}
}

type Service interface {
	Add(email, name string) error
	List() services.ListResponse
}

type handler struct {
	service Service
}

func (h *handler) Add(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	name := r.FormValue("name")

	err := validate(email, name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.Add(email, name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	members := h.service.List()

	pathToFile := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(pathToFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, members); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func validate(email, name string) error {
	ok, err := regexp.Match("^[a-zA-Z|.| ]*$", []byte(name))
	if err != nil {
		return err
	}

	if !ok {
		return errors.New("-----")
	}

	_, err = mail.ParseAddress(email)
	if err != nil {
		return err
	}

	return nil
}
