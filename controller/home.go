package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/pluralsight/webservice/viewmodel"
)

type home struct {
	homeTemplate  *template.Template
	loginTemplate *template.Template
}

func (h home) registerRoutes() {
	http.HandleFunc("/home", h.handleHome)
	http.HandleFunc("/", h.handleHome)
	http.HandleFunc("/login", h.handleLogin)
}

func (h home) handleHome(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewBase()
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Println(fmt.Errorf("error: %v", err))
		}
		fmt.Println(r.Form)
	}
	h.homeTemplate.Execute(w, vm)
}

func (h home) handleLogin(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewLogin()
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Println(fmt.Errorf("error: %v", err))
		}
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		if username == "quinn" && password == "password1" {
			http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
			return
		} else {
			vm.Username = username
			vm.Password = password
		}
	}
	h.loginTemplate.Execute(w, vm)
}
