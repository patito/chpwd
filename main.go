package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-ldap/ldap"
)

type userConfig struct {
	username        string
	password        string
	newPassword     string
	confirmPassword string
}

type ldapConfig struct {
	baseDN  string
	port    int
	address string
}

func changeLdapPassword(user *userConfig) error {

	l := &ldapConfig{"base_dn", 389, "127.0.0.1"}

	conn, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", l.address, l.port))
	if err != nil {
		return err
	}
	defer conn.Close()

	err = conn.Bind(l.baseDN, user.password)
	if err != nil {
		return err
	}

	passwordModifyRequest := ldap.NewPasswordModifyRequest("", user.password, user.newPassword)
	_, err = conn.PasswordModify(passwordModifyRequest)
	if err != nil {
		return err
	}

	return nil
}

func login(w http.ResponseWriter, r *http.Request) error {

	fmt.Println("method:", r.Method)
	if r.Method == "POST" {
		r.ParseForm()

		username := r.Form.Get("username")
		password := r.Form.Get("username")
		newPassword := r.Form.Get("username")
		confirmPassword := r.Form.Get("username")

		if newPassword == confirmPassword {
			user := &userConfig{username, password, newPassword, confirmPassword}
			err := changeLdapPassword(user)
			if err != nil {
				return err
			}
		}
	}

	templ := template.Must(template.ParseFiles(filepath.Join("templates", "index.html")))
	templ.Execute(w, nil)
}

func main() {

	http.HandleFunc("/", login)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAnsServe:", err)
	}
}
